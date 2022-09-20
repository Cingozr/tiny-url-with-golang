package http

import (
	"encoding/json"
	"fmt"
	"github.com/cingozr/go-tiny-url/internal/utils"
	"net/http"
	"strconv"

	url_service "github.com/cingozr/go-tiny-url/internal/tinyurl"
	"github.com/gorilla/mux"
)

func (h *Handler) GetAllUrl(w http.ResponseWriter, r *http.Request) {
	tinyUrls, err := h.TinyUrlService.GetAll()
	if err != nil {
		SendErrorResponse(w, "Failed retrieve all tiny urls", err)
		return
	}

	if err := SendOkResponse(w, tinyUrls); err != nil {
		panic(err)
	}
}

func (h *Handler) PostUrl(w http.ResponseWriter, r *http.Request) {
	var tinyUrl url_service.TinyUrl
	if err := json.NewDecoder(r.Body).Decode(&tinyUrl); err != nil {
		SendErrorResponse(w, "Failed to decode Json Body", err)
		return
	}

	tinyUrl.TinyUrl = fmt.Sprintf("%s/%s", r.Host, utils.RandomUrl(8))
	tinyUrl, err := h.TinyUrlService.Create(tinyUrl)
	if err != nil {
		SendErrorResponse(w, "Failed to post new url", err)
		return
	}

	if err := SendOkResponse(w, tinyUrl); err != nil {
		panic(err)
	}

}

func (h *Handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tinyUrlId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		SendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	tinyUrl, err := h.TinyUrlService.GetByUrlId(uint(tinyUrlId))
	if err != nil {
		SendErrorResponse(w, "Error Retrieving url By ID", err)
		return
	}

	if err := SendOkResponse(w, tinyUrl.OriginalUrl); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	deleteId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		SendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	if err = h.TinyUrlService.DeleteByUrlId(uint(deleteId)); err != nil {
		SendErrorResponse(w, "Failed to delete url by id", err)
	}

	if err := SendOkResponse(w, ResponseModel{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}
