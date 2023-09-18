package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"goSocialNetwork/post"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
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
	log.Println("Serving on post 8000")
	log.Fatalln(http.ListenAndServe(":8000", a.Router))
}

func (a *App) Migrate() error {
	err := a.DB.AutoMigrate(&post.Post{})
	return err
}

func (a *App) Routes() {
	a.Router.HandleFunc("/", index()).Methods("GET")
}

func index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Welcome to the Go social network")
	}
}
