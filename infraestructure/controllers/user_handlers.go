package controllers

import (
	"encoding/json"
	"fmt"
	"hexagonal02/domain/entities"
	"hexagonal02/domain/usecases"
	"hexagonal02/infraestructure/controllers/dto"
	"hexagonal02/infraestructure/helpers/logging"
	"io"
	"net/http"
	"strconv"
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

    responseError, err := uh.UserService.CreateUser(user)
    if err != nil {        
        logger.Errorw("Failed to create user", logging.KeyID, req, logging.KeyErr, err)
        http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)        
        //http.Error(w,"Failed to create user", http.StatusInternalServerError)
        return
    }
    if responseError!= nil && len(responseError.Error) > 0 {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(responseError)    
        return
    } 
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(req)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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
    id, err := strconv.ParseInt(req.ID, 10, 64)
    if err != nil {
        logger.Errorw("Failed to Update user", logging.KeyID, req, logging.KeyErr, err)
        http.Error(w, fmt.Sprintf("Failed to Update user: %v", err), http.StatusInternalServerError)                
        return
    }

    user := entities.NewUserID(req.Name, req.Email, req.Password,id)

    responseError, err := uh.UserService.UpdateUser(user)
    if err != nil {        
        logger.Errorw("Failed to Update user", logging.KeyID, req, logging.KeyErr, err)
        http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)        
        return
    }
    if responseError!= nil && len(responseError.Error) > 0 {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(responseError)    
        return
    }
    req.Password = "" 
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(req)
}