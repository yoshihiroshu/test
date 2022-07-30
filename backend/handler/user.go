package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yoshi429/test/model"
	"golang.org/x/crypto/bcrypt"
)

func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) error {

	var user model.User
	err := h.Context.UnmarshalFromRequest(r, &user)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	user.SetBcryptPassword()

	err = user.Insert(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusInternalServerError, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) error {

	var u model.User
	users, err := u.GetAll(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, users)
}

func (h Handler) GetUserBYEmail(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	email := vars["email"]

	user := &model.User{
		Email: email,
	}

	err := user.GetByEmail(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

func (h Handler) GetUserBYID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user := &model.User{
		ID: uuid.MustParse(id),
	}

	err := user.GetByUUID(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	return h.Context.JSON(w, http.StatusOK, user)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) error {

	var user model.User
	err := h.Context.UnmarshalFromRequest(r, &user)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	// get password by user.email
	hashedPassword, err := user.GetPasswordByEmail(h.Context.Db.PSQLDB)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

	// compare password and crypt password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return h.Context.JSON(w, http.StatusNotFound, "password is mistaken")
	}
	// set cookie
	return h.Context.JSON(w, http.StatusOK, "LOGIN!!!")
}
