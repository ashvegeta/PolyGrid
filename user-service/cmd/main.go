package main

import (
	srv "github.com/ashvegeta/user-service/server"
)

func main() {
	// Start the server
	a := srv.UserServer{}
	defer a.Close()

	a.ConnOps = srv.ConnOps{
		Network: "tcp",
		Addr:    ":8080",
	}

	a.Start()
}
