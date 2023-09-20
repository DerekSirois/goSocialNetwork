package app

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"goSocialNetwork/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var SECRET_KEY = []byte("thisIsAVerySecureKey")

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

type Response struct {
	Msg string
}

type ResponseToken struct {
	Token string
}

func New() (*App, error) {
	cs := "host=localhost user=dev password=abcde dbname=goSocial sslmode=disable"
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &App{
		DB:     db,
		Router: mux.NewRouter(),
	}, nil
}

func (a *App) Run() {
	log.Println("Serving on models 8000")
	log.Fatalln(http.ListenAndServe(":8000", a.Router))
}

func (a *App) Migrate() error {
	err := a.DB.AutoMigrate(&models.Post{}, &models.User{})
	return err
}

func (a *App) Routes() {
	a.Router.HandleFunc("/", index()).Methods("GET")
	a.Router.HandleFunc("/api/post", a.verifyJWT(a.GetAllPostHandler())).Methods("GET")
	a.Router.HandleFunc("/api/post", a.verifyJWT(a.CreatePostHandler())).Methods("POST")
	a.Router.HandleFunc("/api/post/{id:[0-9]+}", a.verifyJWT(a.GetPostByIdHandler())).Methods("GET")
	a.Router.HandleFunc("/api/post/{id:[0-9]+}", a.verifyJWT(a.UpdatePostHandler())).Methods("PUT")
	a.Router.HandleFunc("/api/post/{id:[0-9]+}", a.verifyJWT(a.DeletePostHandler())).Methods("DELETE")

	a.Router.HandleFunc("/api/register", a.Register()).Methods("POST")
	a.Router.HandleFunc("/api/login", a.Login()).Methods("POST")
}

func index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "Welcome to the Go social network")
	}
}

func (a *App) respond(writer http.ResponseWriter, _ *http.Request, data interface{}, statusCode int) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if data == nil {
		return
	}
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		log.Printf("Cannot format json err=%v\n", err)
	}
}

func (a *App) handleError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	log.Println(err)
	a.respond(w, r, &Response{
		Msg: err.Error(),
	}, statusCode)
}

type UserClaim struct {
	jwt.RegisteredClaims
	ID       int
	UserName string
}

func CreateJWTToken(id int, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10))},
		ID:               id,
		UserName:         name,
	})

	// Create the actual JWT token
	signedString, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}

	return signedString, nil
}

func (a *App) verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			var jwtToken = request.Header["Token"][0]
			var userClaim UserClaim
			token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
				return SECRET_KEY, nil
			})
			if err != nil {
				a.respond(writer, request, &Response{Msg: err.Error()}, http.StatusBadRequest)
				return
			}
			if !token.Valid {
				a.respond(writer, request, &Response{Msg: "Invalid token"}, http.StatusBadRequest)
				return
			}
			endpointHandler(writer, request)
		} else {
			a.respond(writer, request, &Response{Msg: "Missing token"}, http.StatusBadRequest)
			return
		}
	})
}
