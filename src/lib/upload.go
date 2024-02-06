package lib

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context, field string, dest string) (string, error){

	
	file, err := c.FormFile(field)
	fmt.Println(file)

	if err != nil{
		fmt.Println(err, file)
		return "", errors.New("error")
	}

	ext := map[string]string{
		"image/jpeg": ".jpg",
		"image/png": ".png",
	}

	fileType := file.Header["Content-Type"][0]
	extension := ext[fileType]
	fmt.Println(extension)
	if extension == ""{
		return "", errors.New("file extension not allowed")
	}

	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), extension)

	c.SaveUploadedFile(file, fileName)

	return fileName, nil
}