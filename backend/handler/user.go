package handler

import (
	"net/http"

	"github.com/yoshi429/test/model"
)

func (h Handler) RegisterAccount(w http.ResponseWriter, r *http.Request) {

	var user model.User
	err := h.Context.UnmarshalFromRequest(r, &user)
	if err != nil {
		h.Context.Logger.Fatalln(err)
	}

	err = user.Insert(h.Context.Db.PSQLDB)
	if err != nil {
		h.Context.Logger.Fatalln(err)
	}

	h.Context.JSON(w, http.StatusOK, user)
}
