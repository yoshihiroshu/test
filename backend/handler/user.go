package handler

import (
	"net/http"

	"github.com/yoshi429/test/model"
)

func (h Handler) RegisterAccount(w http.ResponseWriter, r *http.Request) error {

	var user model.User
	err := h.Context.UnmarshalFromRequest(r, &user)
	if err != nil {
		return h.Context.JSON(w, http.StatusBadRequest, err.Error())
	}

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
