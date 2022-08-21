package main

import (
	"awesomeProject/router"
	"awesomeProject/service"
)

func main() {
	service.InitRedis()
	router.InitRouter()
}
