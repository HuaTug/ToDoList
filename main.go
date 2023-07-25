package main

import (
	"ToDoList/conf"
	"ToDoList/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
