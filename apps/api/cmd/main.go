package main

import "todo_api/server"

func main() {
	app := server.NewAppInit()
	app.Run()
}
