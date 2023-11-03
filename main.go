package main

import (
	"log"
	userservice "userService/kitex_gen/userService/userservice"
)

func main() {
	svr := userservice.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
