package http

import (
	"encoding/json"
	"github.com/cingozr/go-tiny-url/internal/user"
	"net/http"
)

func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {
	var userModel user.User
	if err := json.NewDecoder(r.Body).Decode(&userModel); err != nil {
		SendErrorResponse(w, "Failed to decode Json Body", err)
		return
	}

	tinyUrl, err := h.UserService.Create(userModel)
	if err != nil {
		SendErrorResponse(w, "Failed to post new user", err)
		return
	}

	if err := SendOkResponse(w, tinyUrl); err != nil {
		panic(err)
	}

}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	var userModel user.User
	if err := json.NewDecoder(r.Body).Decode(&userModel); err != nil {
		SendErrorResponse(w, "Failed to decode Json Body", err)
		return
	}

	err := h.UserService.GetUser(userModel)
	if err != nil {
		SendErrorResponse(w, "Couldn't find user", err)
		return
	}

	token, err := CreateJWT()
	if err != nil {
		SendErrorResponse(w, "Failed to CreateJWT", err)
		return
	}
	SendOkResponse(w, token)
}
