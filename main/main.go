package main

import router2 "github.com/ItsWewin/the-general-public/router"

func main() {
	router := router2.Router()
	router.Run(":8080")
}
