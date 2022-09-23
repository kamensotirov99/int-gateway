package main

import (
	"int-gateway/app"
)

func main() {
	a := app.InitializeApp()
 	err := a.Run()
	if err != nil {
		panic(err.Error())
	}

}
