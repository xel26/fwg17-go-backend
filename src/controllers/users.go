package controllers

import (
	"coffe-shop-be-golang/src/middleware"
	"coffe-shop-be-golang/src/models"
	"fmt"
	"log"
	"math"
	"strings"

	"net/http"
	"strconv"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	TotalPage int `json:"totalPage"`
	NextPage int `json:"nextPage"`
	PrevPage int `json:"prevPage"`
	Limit int `json:"limit"`
	TotalData int `json:"totalData"`
}

type ResponseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"PageInfo"`
	Results  interface{} `json:"results"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type ResponseOnly2 struct {
	Success bool   `json:"success"`
	Message error `json:"message"`
}


func ListAllUsers(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}


	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllUsers(searchKey, sortBy, order, limit, offset)
	if len(result.Data) == 0 {
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "data not found",
		})
		return
	}


	totalPage := int(math.Ceil(float64(result.Count)/float64(limit)))
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = 0
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	PageInfo := PageInfo{
		CurrentPage: page,
		NextPage: nextPage,
		PrevPage: prevPage,
		Limit: limit,
		TotalPage: totalPage,
		TotalData: result.Count,
	}


	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &ResponseList{
		Success: true,
		Message: "List all Users",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailUser(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}



	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.FindOneUsers(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "User not found",
			})
		return
		}

		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Detail user",
		Results: user,
	})
}


func CreateUser(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}



	data := models.User{}
	err := c.ShouldBind(&data)


	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ResponseOnly2{
			Success: false,
			Message: err,
		})
		return
	}
	data.Password =  hash.String()


	user, err := models.CreateUser(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &ResponseOnly2{
			Success: false,
			Message: err,
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "User created successfully",
		Results: user,
	})
}


func UpdateUser(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	isExist, error := models.FindOneUsers(id)
	if error != nil{
		fmt.Println(isExist, error)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "no data found",
		})
	return
	}

	data := models.User{}
	err := c.ShouldBind(&data)

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "failed generate hash",
		})
		return
	}
	data.Password =  hash.String()

	data.Id = id

	user, err := models.UpdateUser(data)

	if err != nil{
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "User updated successfully",
		Results: user,
	})
}



func DeleteUser(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}


	id, _ := strconv.Atoi(c.Param("id"))
	isExist, error := models.FindOneUsers(id)
	if error != nil{
		fmt.Println(isExist, error)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "no data found",
		})
	return
	}

	
	user, err := models.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Delete User Successfully",
		Results: user,
	})
}