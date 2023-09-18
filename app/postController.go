package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goSocialNetwork/models"
	"log"
	"net/http"
	"strconv"
)

func (a *App) GetAllPostHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		posts, err := a.GetAllPost()
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusInternalServerError)
			return
		}
		a.respond(writer, request, posts, http.StatusOK)
	}
}

func (a *App) GetPostByIdHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		p, err := a.GetPostById(uint(id))
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusInternalServerError)
			return
		}
		a.respond(writer, request, p, http.StatusOK)
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		p := &models.Post{}
		err := json.NewDecoder(request.Body).Decode(p)
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		err = a.CreatePost(p)
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusInternalServerError)
			return
		}
		a.respond(writer, request, &Response{Msg: "Post created successfully"}, http.StatusOK)
	}
}

func (a *App) UpdatePostHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		p := &models.Post{}
		err := json.NewDecoder(request.Body).Decode(p)
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		vars := mux.Vars(request)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		err = a.UpdatePost(*p, uint(id))
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusInternalServerError)
			return
		}
		a.respond(writer, request, &Response{Msg: "Post updated successfully"}, http.StatusOK)
	}
}

func (a *App) DeletePostHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		err = a.DeletePost(uint(id))
		if err != nil {
			log.Println(err)
			a.respond(writer, request, &Response{
				Msg: err.Error(),
			}, http.StatusInternalServerError)
			return
		}
		a.respond(writer, request, &Response{Msg: "Post deleted successfully"}, http.StatusOK)
	}
}
