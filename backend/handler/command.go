package handler

import (
	"bytes"
	"net/http"
	"os/exec"
)

func (h Handler) Command(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("pwd")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		h.Context.Logger.Fatalln(err)
	}

	h.Context.JSON(w, http.StatusOK, out.String())
}
