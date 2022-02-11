package main

import "github.com/Thalisonh/star-wars-api/server"

func main() {
	s := server.Init()
	s.Run()
}
