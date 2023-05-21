package main

import "EPICS/server"

func main() {
	server.GoogleConfig()
	server.ConnectDB()
	server.Serve()

}
