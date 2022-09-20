package http

import (
	"encoding/json"
	url_service "github.com/cingozr/go-tiny-url/internal/tinyurl"
	"github.com/cingozr/go-tiny-url/internal/user"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router         *mux.Router
	TinyUrlService *url_service.TinyUrlService
	UserService    *user.UserService
}

type ResponseModel struct {
	Message string
	Error   string
}

func NewHandler(service *url_service.TinyUrlService, userService *user.UserService) *Handler {
	return &Handler{
		TinyUrlService: service,
		UserService:    userService,
	}
}

func (h *Handler) SetupRoutes() {
	log.Info("Setting Up Routes")
	h.Router = mux.NewRouter()

	//if there is any middleware you can use
	h.Router.Use(LoggingMiddleware)

	//TinyUrl  Routes Start
	h.Router.HandleFunc("/api/tinyurl", JwtAuth(h.PostUrl)).Methods("POST")
	h.Router.HandleFunc("/api/tinyurl/getallurl", JwtAuth(h.GetAllUrl)).Methods("GET")
	h.Router.HandleFunc("/api/tinyurl/{id}", JwtAuth(h.GetUrl)).Methods("GET")
	h.Router.HandleFunc("/api/tinyurl/{id}", JwtAuth(h.DeleteUrl)).Methods("DELETE")
	//TinyUrl  Routes End

	//User  Routes Start
	h.Router.HandleFunc("/api/user", h.PostUser).Methods("POST")
	h.Router.HandleFunc("/api/user/login", h.GetUser).Methods("POST")
	//User  Routes End

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := SendOkResponse(w, ResponseModel{Message: "I'm Alive"}); err != nil {
			panic(err)
		}
	})
}

//SendOkResponse - Retrieve error message handler
func SendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

//SendErrorResponse - Retrieve error message handler
func SendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(ResponseModel{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
