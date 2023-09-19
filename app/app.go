package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goSocialNetwork/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

type Response struct {
	Msg string
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
	a.Router.HandleFunc("/api/post", a.GetAllPostHandler()).Methods("GET")
	a.Router.HandleFunc("/api/post", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/post/{id:[0-9]+}", a.GetPostByIdHandler()).Methods("GET")
	a.Router.HandleFunc("/api/post/{id:[0-9]+}", a.UpdatePostHandler()).Methods("PUT")
	a.Router.HandleFunc("/api/post/{id:[0-9]+}", a.DeletePostHandler()).Methods("DELETE")

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
