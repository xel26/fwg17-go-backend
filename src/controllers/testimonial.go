package controllers

import (
	"coffe-shop-be-golang/src/middleware"
	"coffe-shop-be-golang/src/models"
	"coffe-shop-be-golang/src/service"
	"context"
	"fmt"
	"math"
	"os"
	"strings"

	"net/http"
	"net/url"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)


func ListAllTestimonial(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))
	offset := (page - 1) * limit

	result, err := models.FindAllTestimonial(searchKey, sortBy, order, limit, offset)
	if len(result.Data) == 0 {
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
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

	PageInfo := service.PageInfo{
		CurrentPage: page,
		NextPage: nextPage,
		PrevPage: prevPage,
		Limit: limit,
		TotalPage: totalPage,
		TotalData: result.Count,
	}


	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.ResponseList{
		Success: true,
		Message: "List all testimonial",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailTestimonial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	testimonial, err := models.FindOneTestimonial(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Testimonial not found",
			})
		return
		}

		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Detail testimonial",
		Results: testimonial,
	})
}


func CreateTestimonial(c *gin.Context) {
	data := models.TestimonialForm{}
	err := c.ShouldBind(&data)
	if err != nil{
		fmt.Println(err)
		return
	}

	
	_, err = c.FormFile("image")
	if err == nil {
		file, err := middleware.Upload(c, "image", "testimonial")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		data.Image = file
	}else{
		fmt.Println(err)
		data.Image = ""
	}


	testimonial, err := models.CreateTestimonial(data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Testimonial created successfully",
		Results: testimonial,
	})
}


func UpdateTestimonial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.TestimonialForm{}

	c.ShouldBind(&data)
	data.Id = id

	isExist, err := models.FindOneTestimonial(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Testimonial not found",
		})
	return
	}


	_, err = c.FormFile("image")
	if err == nil {
		//// without cloudinary
		// _ = os.Remove("./" + isExist.Image)


		//with cloudinary
		if isExist.Image != ""{
			cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
			resp, err := cld.Admin.Search(context.Background(), search.Query{
				Expression: url.QueryEscape(isExist.Image),
				MaxResults: 1,
			})
			
			response := resp.Response
			responseMap := response.(*map[string]interface{})
			resources := (*responseMap)["resources"].([]interface{})
			resourcesMap := resources[0].(map[string]interface{})
			publicId := resourcesMap["public_id"].(string)
	
			if err == nil {
				_, err := cld.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: publicId})
				if err != nil {
					c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
						Success: false,
						Message: err.Error(),
					})
					return
				}
			}
		}

		file, err := middleware.Upload(c, "image", "testimonial")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		data.Image = file
	}else{
		fmt.Println(err)
		data.Image = ""
	}


	testimonial, err := models.UpdateTestimonial(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Testimonial not found",
			})
		return
		}
		
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Testimonial updated successfully",
		Results: testimonial,
	})
}


func DeleteTestimonial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneTestimonial(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Testimonial not found",
		})
	return
	}

	if isExist.Image != ""{
		cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
		resp, err := cld.Admin.Search(context.Background(), search.Query{
			Expression: url.QueryEscape(isExist.Image),
			MaxResults: 1,
		})
		
		response := resp.Response
		responseMap := response.(*map[string]interface{})
		resources := (*responseMap)["resources"].([]interface{})
		resourcesMap := resources[0].(map[string]interface{})
		publicId := resourcesMap["public_id"].(string)

		if err == nil {
			_, err := cld.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: publicId})
			if err != nil {
				c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
					Success: false,
					Message: err.Error(),
				})
				return
			}
		}
	}

	testimonial, err := models.DeleteTestimonial(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Delete testimonial successfully",
		Results: testimonial,
	})
}