package controllers_customer

import (
	"coffe-shop-be-golang/src/lib"
	"coffe-shop-be-golang/src/models"
	"coffe-shop-be-golang/src/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

type FormReset struct {
	Email           string `form:"email"`
	Otp             string `form:"otp"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=Password"`
}

// func Login(c *gin.Context) {
// 	form := models.User{}
// 	err := c.ShouldBind(&form)

// 	found, err := models.FindOneUsersByEmail(form.Email)

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, &service.ResponseOnly{
// 			Success: false,
// 			Message: "wrong email",
// 		})
// 		return
// 	}

// 	decode, err := argonize.DecodeHashStr(found.Password)

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, &service.ResponseOnly{
// 			Success: false,
// 			Message: "pass error",
// 		})
// 		return
// 	}

// 	token, err := service.GenerateToken(found.Id, found.Role)

// 	plain := []byte(form.Password)
// 	if decode.IsValidPassword(plain) {
// 		c.JSON(http.StatusOK, &service.Response{
// 			Success: true,
// 			Message: "Login success",
// 			Results: token,
// 		})
// 		return
// 	} else {
// 		c.JSON(http.StatusUnauthorized, &service.ResponseOnly{
// 			Success: false,
// 			Message: "wrong password",
// 		})
// 	}

// }

func Register(c *gin.Context) {
	form := models.UserForm{}
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}


	defaultRole := "customer"
	form.Role = &defaultRole
	form.Picture = ""

	plain := []byte(form.Password)
	hash, _ := argonize.Hash(plain)
	form.Password = hash.String()

	result, err := models.CreateUser(form)

	if err != nil {
		fmt.Println(err)
		if strings.HasSuffix(err.Error(), `unique constraint "users_email_key"`) {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "email already registered. . . please login",
			})
			return
		}
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: "Register failed",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Please check your email to confirm your account!",
		Results: result,
	})
}

func ForgotPassword(c *gin.Context) {
	form := FormReset{}
	c.ShouldBind(&form)

	if form.Email != "" {
		found, _ := models.FindOneUsersByEmail(form.Email)

		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "email not registered. . . . please use another email",
			})
			return
		}

		otp := lib.RandomNumberStr(6)
		FormReset := models.ForgotPassword{
			Otp:   otp,
			Email: form.Email,
		}
		models.CreateForgotPassword(FormReset)
		// START SEND EMAIL
		lib.Mail(
			found.Email,
			found.FullName,
			otp,
			"enter the 6-digit code below to create a new password",
			"create new password",
			"http://localhost:5173/create-new-password",
			"Thank you for entrusting us to safeguard your account security.",
			"Here is your OTP code",
		)
		fmt.Println(otp)
		// END SEND EMAIL
		c.JSON(http.StatusOK, &service.ResponseOnly{
			Success: true,
			Message: "OTP has been sent to your email",
		})
		return
	}

	if form.Otp != "" {
		found, _ := models.FindOneByOtp(form.Otp)
		if found.Id == 0 {
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: "invalid OTP code. . . please enter the correct code",
			})
			return
		}

		// if form.Password != form.ConfirmPassword{
		// 	c.JSON(http.StatusBadRequest, &service.ResponseOnly{
		// 		Success: false,
		// 		Message: "Confirm password does not match",
		// 	})
		// 	return
		// }

		foundUser, err := models.FindOneUsersByEmail(found.Email)
		if err != nil{
			fmt.Println(err)
			return
		}

		data := models.UserForm{}

		plain := []byte(form.Password)
		hash, _ := argonize.Hash(plain)
		data.Password = hash.String()
		data.Id = foundUser.Id

		updated, err := models.UpdateUser(data)
		fmt.Println(err, updated)
		if err != nil {
			fmt.Println(err,updated)
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
		}

		models.DeleteForgotPassword(found.Id)
		message := fmt.Sprintf("Reset password for %v success", *updated.Email)
		c.JSON(http.StatusOK, &service.ResponseOnly{
			Success: true,
			Message: message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
		Success: false,
		Message: "Internal server error",
	})
}




func ConfirmAccount(c *gin.Context){
	form := models.ConfirmAccount{}
	err := c.ShouldBind(&form)
	if err != nil{
		c.JSON(http.StatusOK, &service.ResponseOnly{
			Success: true,
			Message: err.Error(),
		})
	}


	result := lib.Mail(
		*form.Email,
		*form.FullName,
		"",
		"Welcome to Coffee Shop Web App, We're very excited to have you on board",
		"confirm your account",
		"http://localhost:5173/login",
		"Let's start your coffee journey.",
		"confirm account",
	)

	if result {
		c.JSON(http.StatusOK, &service.ResponseOnly{
			Success: true,
			Message: "Confirm Password Success",
		})
	}else{
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: "Sorry, we couldn't find an account. Please provide a valid email address",
		})
	}
}