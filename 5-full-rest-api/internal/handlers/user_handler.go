package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/dto"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/entity"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/repository"
)

type UserHandler struct {
	UserRepo     repository.UserRepositoryInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(repo repository.UserRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserRepo:     repo,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

// CreateUser 	godoc
// @Summary 		Create a new user
// @Description Create a new user
// @Tags 				users
// @Accept  		json
// @Produce  		json
// @Param 			request 	body 			dto.CreateUserInput true "User request"
// @Success 		201
// @Failure 		500 			{object} 	dto.Error
// @Router 			/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserRepo.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Auth 	godoc
// @Summary 		Auth a user
// @Description Auth a user
// @Tags 				users
// @Accept  		json
// @Produce  		json
// @Param 			request 	body 			dto.AuthUserInput true "User auth"
// @Success 		200 			{object} 	dto.JWTOutput
// @Failure 		401
// @Router 			/users/auth [post]
func (h *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user dto.AuthUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserRepo.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := dto.JWTOutput{AccessToken: token}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(accessToken)
}
