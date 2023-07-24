package server

import (
	"GobbleGame/internal/store"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

type Server struct {
	mux   *http.ServeMux
	store *store.Store
}

func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}
	s.mux.HandleFunc("user/register", CorrectRequest(s.HandleUserRegister, "POST"))
	s.mux.HandleFunc("user/unique_username", CorrectRequest(s.HandleUniqueUsername, "POST"))
	s.mux.HandleFunc("user/unique_email", CorrectRequest(s.HandleUniqueEmail, "POST"))
	s.mux.HandleFunc("jwt/login", CorrectRequest(s.HandleJwtLogin, "POST"))
	s.mux.HandleFunc("jwt/refresh", CorrectRequest(s.HandleJwtRefresh, "POST"))
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func Response(w http.ResponseWriter, r *http.Request, code int, data any) {

}

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {

}

func (s *Server) correctRequest(r *http.Request, request any) (int, error) {
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return http.StatusBadRequest, err
	}
	if err := validate.Struct(request); err != nil {
		return http.StatusUnprocessableEntity, err
	}
	return 0, nil
}
