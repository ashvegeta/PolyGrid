package main

func main() {
	// Start the server
	a := UserServer{}
	defer a.Close()

	a.ConnOps = ConnOps{
		Network: "tcp",
		Addr:    ":8080",
	}

	a.Start()
}
