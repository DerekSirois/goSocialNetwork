package app

import (
	"encoding/json"
	"goSocialNetwork/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserJson struct {
	Username string
	Password string
	Active   bool
}

func (a *App) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		u := &UserJson{}
		err := json.NewDecoder(request.Body).Decode(u)
		if err != nil {
			a.handleError(writer, request, err, http.StatusBadRequest)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		if err != nil {
			a.handleError(writer, request, err, http.StatusInternalServerError)
			return
		}
		userModel := mapToUser(u, hash)
		err = a.CreateUser(userModel)
		if err != nil {
			a.handleError(writer, request, err, http.StatusInternalServerError)
			return
		}
		a.respond(writer, request, &Response{Msg: "User created successfully"}, http.StatusOK)
	}
}

func mapToUser(u *UserJson, hash []byte) *models.User {
	return &models.User{
		Username: u.Username,
		Password: hash,
		Active:   u.Active,
	}
}
