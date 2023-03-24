package controllers

import (
	"encoding/json"
	"fmt"
	"hexagonal02/domain/entities"
	"hexagonal02/domain/usecases"
	"hexagonal02/infraestructure/controllers/dto"
	"hexagonal02/utils/logging"
	"io"
	"net/http"
)

var (
	logger = logging.NewLogger()
)

type UserHandler struct {
	UserService usecases.UserService
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }

    var req dto.UserCreateRequest
    err = json.Unmarshal(body, &req)
    if err != nil {
        http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
        return
    }

    user := entities.NewUser(req.Name, req.Email, req.Password)

    err = uh.UserService.CreateUser(user)
    if err != nil {
        //log.Println("Error:",err)
        logger.Errorw("Failed to create user", logging.KeyID, req, logging.KeyErr, err)
        http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)        
        //http.Error(w,"Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(req)
}