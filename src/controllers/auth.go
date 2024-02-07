package controllers

import (
	"coffe-shop-be-golang/src/lib"
	"coffe-shop-be-golang/src/models"
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
// 		c.JSON(http.StatusUnauthorized, ResponseOnly{
// 			Success: false,
// 			Message: "wrong email",
// 		})
// 		return
// 	}

// 	decode, err := argonize.DecodeHashStr(found.Password)
	
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, ResponseOnly{
// 			Success: false,
// 			Message: "pass error",
// 		})
// 		return
// 	}
	
// 	token, err := service.GenerateToken(found.Id, found.Role)

// 	plain := []byte(form.Password)
// 	if decode.IsValidPassword(plain) {
// 		c.JSON(http.StatusOK, Response{
// 			Success: true,
// 			Message: "Login success",
// 			Results: token,
// 		})
// 		return
// 	} else {
// 		c.JSON(http.StatusUnauthorized, ResponseOnly{
// 			Success: false,
// 			Message: "wrong password",
// 		})
// 	}

// }



func Register(c *gin.Context) {
	form := models.UserForm{}
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	
	defaultRole := "customer"
	form.Role = &defaultRole

	plain := []byte(form.Password)
	hash, err := argonize.Hash(plain)
	form.Password = hash.String()

	result, err := models.CreateUser(form)

	if err != nil {
		if strings.HasSuffix(err.Error(), `unique constraint "users_email_key"`) {
			c.JSON(http.StatusBadRequest, ResponseOnly{
				Success: false,
				Message: "Email already registered",
			})
			return
		}
		c.JSON(http.StatusBadRequest, ResponseOnly{
			Success: false,
			Message: "Register failed",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Register success",
		Results: result,
	})
}



func ForgotPassword(c *gin.Context) {
	form := FormReset{}
	c.ShouldBind(&form)

	if form.Email != ""{
		found, _ := models.FindOneUsersByEmail(form.Email)
		if found.Id != 0{
			FormReset := models.ForgotPassword{
				Otp: lib.RandomNumberStr(6),
				Email: form.Email,
			}
			models.CreateForgotPassword(FormReset)
			// START SEND EMAIL
			fmt.Println(FormReset.Otp)
			// END SEND EMAIL
			c.JSON(http.StatusOK, ResponseOnly{
				Success: true,
				Message: "OTP has been sent to your email",
			})
			return
			
		}else{
			c.JSON(http.StatusBadRequest, ResponseOnly{
				Success: false,
				Message: "email not registered... failed to reset password",
			})
			return
		}
	}

	if form.Otp != ""{
		found, err := models.FindOneByOtp(form.Otp)
		fmt.Println(found.Id, err)
		if found.Id != 0{
			if form.Password == form.ConfirmPassword{
				foundUser, _ := models.FindOneUsersByEmail(found.Email)
				data := models.UserForm{
					Id: foundUser.Id,
				}

				hash, _ := argonize.Hash([]byte(form.Password))
				data.Password = hash.String()

				updated, _ := models.UpdateUser(data)


				models.DeleteForgotPassword(found.Id)
				message := fmt.Sprintf("Reset password for %v success", updated.Email)
				c.JSON(http.StatusOK, ResponseOnly{
					Success: true,
					Message: message,
				})
				return
			}else{
				c.JSON(http.StatusBadRequest, ResponseOnly{
					Success: false,
					Message: "Confirm password does not match",
				})
				return
			}
		}else{
			c.JSON(http.StatusBadRequest, ResponseOnly{
				Success: false,
				Message: "invalid otp code",
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, ResponseOnly{
		Success: false,
		Message: "Internal server error",
	})
}
