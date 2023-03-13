package main

import "todo_api/internal/server"

func main() {
	app := server.NewAppInit()
	app.Run(app.Router)
}
