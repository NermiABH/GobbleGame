package server

import (
	"GobbleGame/internal/jwt"
	"GobbleGame/internal/model"
	"GobbleGame/internal/vars"
	"net/http"
)

func (s *Server) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	req := vars.RequestUserRegister{}
	if code, err := s.correctRequest(r, req); err != nil {
		Error(w, r, code, err)
		return
	}
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.store.UserRegister(user); err != nil {
		//if errors.Is(err, ){
		//	Error(w, r, http.StatusUnprocessableEntity, err)
		//	return
		//}
		Error(w, r, http.StatusInternalServerError, vars.ErrInternalServer)
		return
	}
	Response(w, r, http.StatusOK, user.ToResponse())
}

func (s *Server) HandleJwtLogin(w http.ResponseWriter, r *http.Request) {
	req := vars.RequestJwtLogin{}
	if code, err := s.correctRequest(r, req); err != nil {
		Error(w, r, code, err)
		return
	}
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.store.UserExist(user); err != nil {
		//if errors.Is(err, ){
		//	Error(w, r, http.StatusUnprocessableEntity, err)
		//	return
		//}
		Error(w, r, http.StatusInternalServerError, vars.ErrInternalServer)
		return
	}
	Response(w, r, http.StatusOK, jwt.New(user).ToResponse())
}

func (s *Server) HandleJwtRefresh(w http.ResponseWriter, r *http.Request) {
	req := vars.RequestJwtRefresh{}
	if code, err := s.correctRequest(r, req); err != nil {
		Error(w, r, code, err)
		return
	}
	tokens, err := jwt.RecreateTokens(req.Refresh)
	if err != nil {
		Error(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	Response(w, r, http.StatusOK, tokens.ToResponse())
}

func (s *Server) HandleUniqueUsername(w http.ResponseWriter, r *http.Request) {
	req := vars.RequestUserUniqueUsername{}
	if code, err := s.correctRequest(r, req); err != nil {
		Error(w, r, code, err)
		return
	}
	if err := s.store.UserUniqueUsername(req.Username); err != nil {
		//if errors.Is(err, ){
		//	Error(w, r, http.StatusUnprocessableEntity, err)
		//	return
		//}
		Error(w, r, http.StatusInternalServerError, vars.ErrInternalServer)
		return
	}
	Response(w, r, http.StatusOK, nil)
}

func (s *Server) HandleUniqueEmail(w http.ResponseWriter, r *http.Request) {
	req := vars.RequestUserUniqueEmail{}
	if code, err := s.correctRequest(r, req); err != nil {
		Error(w, r, code, err)
		return
	}
	if err := s.store.UserUniqueEmail(req.Email); err != nil {
		//if errors.Is(err, ){
		//	Error(w, r, http.StatusUnprocessableEntity, err)
		//	return
		//}
		Error(w, r, http.StatusInternalServerError, vars.ErrInternalServer)
		return
	}
	Response(w, r, http.StatusOK, nil)
}
