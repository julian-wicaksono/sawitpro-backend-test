package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/util"
	"github.com/labstack/echo/v4"
)

// This is endpoint to log in user
// (GET /login)

func (s *Server) Login(ctx echo.Context) error {
	params := generated.LoginRequest{}
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	var resp generated.SuccessResponse
	userData, userId, err := s.Repository.GetUserDataByPhoneNumber(ctx.Request().Context(),
		repository.UserData{
			PhoneNumber: params.PhoneNumber,
		})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: "Error login",
		})
	}

	if util.CompareHash(params.Password, userData.Password) {
		expirationTime := time.Now().Add(15 * time.Minute)
		token, err := util.CreateToken(userId.UserID, expirationTime, s.JWTKey)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "Error login",
			})
		}
		authorizationCookie := new(http.Cookie)
		authorizationCookie.Name = "authorization"
		authorizationCookie.Value = token
		authorizationCookie.Expires = expirationTime
		ctx.SetCookie(authorizationCookie)

		resp.UserId = int(userId.UserID)
		return ctx.JSON(http.StatusOK, resp)
	}

	return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
		Message: "Wrong password",
	})
}

func (s *Server) Registration(ctx echo.Context) error {
	params := generated.RegisterRequest{}
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	var resp generated.SuccessResponse
	userId, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(),
		repository.UserData{
			PhoneNumber: params.PhoneNumber,
		})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	if userId.UserID != 0 {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: "Phone number already exist",
		})
	}

	validationJson, err := ValidateRegister(params)

	if err != nil {
		validationMessage, _ := json.Marshal(validationJson)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: string(validationMessage),
		})
	}

	hashPassword, err := util.GenerateHash(params.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	userId, err = s.Repository.InsertUser(ctx.Request().Context(),
		repository.UserData{
			PhoneNumber: params.PhoneNumber,
			FullName:    params.FullName,
			Password:    hashPassword,
		})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	resp.UserId = int(userId.UserID)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetProfile(ctx echo.Context) error {

	cookie, err := ctx.Cookie("authorization")
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{
			Message: "Please login",
		})
	}
	token, err := util.ParseToken(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	userData, err := s.Repository.GetUserDataById(ctx.Request().Context(), repository.UserId{
		UserID: token.UserId,
	})

	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, generated.GetProfileSuccessResponse{
		FullName:    userData.FullName,
		PhoneNumber: userData.PhoneNumber,
	})
}

func (s *Server) UpdateProfile(ctx echo.Context, params generated.UpdateProfileParams) error {

	cookie, err := ctx.Cookie("authorization")
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{
			Message: "Please login",
		})
	}
	token, err := util.ParseToken(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{
			Message: "Failed to update profile",
		})
	}

	if params.FullName != nil && params.PhoneNumber != nil {
		userId, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(),
			repository.UserData{
				PhoneNumber: *params.PhoneNumber,
			})

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "Error updating user data",
			})
		}

		if userId.UserID != 0 {
			return ctx.JSON(http.StatusConflict, generated.ErrorResponse{
				Message: "Phone number already exist",
			})
		}

		err = s.Repository.UpdateUserData(ctx.Request().Context(),
			repository.UserData{
				PhoneNumber: *params.PhoneNumber,
				FullName:    *params.FullName,
			},
			repository.UserId{
				UserID: token.UserId,
			})

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "Error updating user data",
			})
		}

	} else if params.PhoneNumber != nil {
		userId, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(),
			repository.UserData{
				PhoneNumber: *params.PhoneNumber,
			})

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "Error updating user data",
			})
		}

		if userId.UserID != 0 {
			return ctx.JSON(http.StatusConflict, generated.ErrorResponse{
				Message: "Phone number already exist",
			})
		}

		err = s.Repository.UpdatePhoneNumber(ctx.Request().Context(),
			repository.UserData{
				PhoneNumber: *params.PhoneNumber,
			},
			repository.UserId{
				UserID: token.UserId,
			})

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "Error updating user data",
			})
		}

	} else if params.FullName != nil {
		err = s.Repository.UpdateFullName(ctx.Request().Context(),
			repository.UserData{
				FullName: *params.FullName,
			},
			repository.UserId{
				UserID: token.UserId,
			})

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
				Message: "Error updating user data",
			})
		}
	}

	return ctx.JSON(http.StatusOK, generated.SuccessResponse{
		UserId: int(token.UserId),
	})

}
