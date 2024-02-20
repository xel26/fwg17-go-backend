package lib

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/matcornic/hermes/v2"
	gomail "gopkg.in/mail.v2"
)

func Mail(receiver string, name string, otp string, intro string, textAction string, linkAction string, outro string, subject string) bool {
	h := hermes.Hermes{
		Theme: &hermes.Default{},
		Product: hermes.Product{
			Name: "Coffee Shop Web",
			Logo: "https://res.cloudinary.com/dgtv2r5qh/image/upload/v1708235806/coffee-shop-be/email%20activation/oyy81bgdc4ksu66xs0wu.png",
			Copyright: "Copyright © 2019 Coffee Shop Web App. All rights reserved",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				intro, otp,
			},

			Actions: []hermes.Action{
				{Instructions: "please click here",
					Button: hermes.Button{
						Color: "#000000",
						Text:  textAction,
						Link:  linkAction,
					},
				},
			},

			Outros: []string{outro},
		},
	}


	htmlBody, err := h.GenerateHTML(email)
	if err != nil{
		fmt.Println(err)
		return false
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "alert.app.services@gmail.com")
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", subject + otp)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, "alert.app.services@gmail.com", os.Getenv("EMAIL_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
