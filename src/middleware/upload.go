package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// // tanpa cloudinary
// func Upload(c *gin.Context, field string, dest string) (string, error){

// 	file, err := c.FormFile(field)

// 	if err != nil{
// 		fmt.Println(err)
// 		return "", errors.New("field file not found")
// 	}

// 	ext := map[string]string{
// 		"image/jpeg": ".jpg",
// 		"image/png": ".png",
// 	}

// 	fileType := file.Header["Content-Type"][0]
// 	extension := ext[fileType]
// 	if extension == ""{
// 		fmt.Println(fileType)
// 		return "", errors.New("only jpeg, jpg, and png files allowed")
// 	}

// 	fileSize := file.Size
// 	if fileSize > 2 * 1024 * 1024{
// 		fmt.Println(fileSize)
// 		return "", errors.New("the file size exceeds the maximum limit of 2 MB")
// 	}

// 	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), extension)

// 	c.SaveUploadedFile(file, fileName)

// 	return fileName, nil
// }

// menggunakan cloudinary
func Upload(c *gin.Context, field string, dest string) (string, error){
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	if err != nil{
		fmt.Println(err)
		return "", errors.New(err.Error())
	}

	file, err := c.FormFile(field)

	if err != nil{
		fmt.Println(err)
		return "", errors.New("field file not found")
	}

	ext := map[string]string{
		"image/jpeg": ".jpg",
		"image/png": ".png",
		"image/jpg": ".jpg",
	}


	fileType := file.Header["Content-Type"][0]
	extension := ext[fileType]
	if extension == ""{
		fmt.Println(fileType)
		return "", errors.New("only jpeg, jpg, and png files allowed")
	}

	fileSize := file.Size
	if fileSize > 2 * 1024 * 1024{
		fmt.Println(fileSize)
		return "", errors.New("the file size exceeds the maximum limit of 2 MB")
	}

	// fileName := fmt.Sprintf("coffee-shop-be/%v/%v%v", dest, uuid.NewString(), extension)

	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: uuid.NewString(),
		Folder: fmt.Sprintf("coffee-shop-be/%v", dest),
	})
	
	if err != nil{
		fmt.Println("err", err)
		return "", errors.New(err.Error())
	}

	fmt.Println("resutl", result)
	return result.SecureURL, nil
}