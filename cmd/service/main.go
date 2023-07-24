package main

import "GobbleGame/internal/service"

func main() {
	err := service.Run()
	if err != nil {
		panic(err)
	}
}
