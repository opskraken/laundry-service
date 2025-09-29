package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	userRepo "github.com/enghasib/laundry_service/repo/user"
	"github.com/enghasib/laundry_service/utils"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id          int    `json:"id"`
	UserName    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Status      string `json:"status" validate:"uppercase"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// decode request body
	var user *User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Failed to decode request body!")
		utils.SendError(w, http.StatusBadRequest, "Failed to decode request body!")
	}
	// validate user input
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Validation failed on field:", err.Field(), "Condition:", err.Tag())
			utils.SendError(w, http.StatusBadRequest, fmt.Sprintf("Validation failed on field: %s Condition:%s", err.Field(), err.Tag()))
			return
		}
	}

	repoUser := &userRepo.User{
		Id:          user.Id,
		UserName:    user.UserName,
		Email:       user.Email,
		Password:    user.Password,
		Status:      user.Status,
		IsShopOwner: user.IsShopOwner,
	}
	createdUser, err := h.usrRepo.Create(*repoUser)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SendResponse(w, http.StatusCreated, *createdUser)
}
