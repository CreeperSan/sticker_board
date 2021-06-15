package main

import (
	"fmt"
)

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

type AccountV1Login struct {
	Account string `json:"account"`
	Password string `json:"password"`
}

func main() {
	fmt.Println("Hello World")

	app := iris.New()

	app.Use(recover.New())

	app.Post("/api/account/v1/login", func(context iris.Context) {
		var loginRequest AccountV1Login
		err := context.ReadJSON(&loginRequest)
		if err != nil {
			context.Writef("Account -> readJson error = ", err)
			return
		}

		context.JSON(iris.Map{
			"code" : 200,
			"account" : loginRequest.Account,
			"passowrd" : loginRequest.Password,
		})

		//context.Writef("Account -> account=%s password=%s", loginRequest.Account, loginRequest.Password)
	})


	app.Post("/api/account/v1/request_register", func(context iris.Context) {
		email := context.Params().Get("email")
		context.Writef("Account -> email=%s", email)
	})

	app.Post("/api/account/v1/register", func(context iris.Context) {
		account := context.Params().Get("account")
		password := context.Params().Get("password")
		email := context.Params().Get("email")
		authKey := context.Params().Get("auth_key")
		authCode := context.Params().Get("auth_code")
		context.Writef("Account -> account=%s password=%s email=%s authKey=%s authCode=%s",
			account, password, email, authKey, authCode)
	})

	app.Post("/api/account/v1/auth", func(context iris.Context) {
		key := context.Params().Get("key")
		context.Writef("Account -> key=%s", key)
	})


	err := app.Listen(":8080")
	if err != nil {
		return 
	}
}
