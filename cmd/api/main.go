package main

import "github.com/fathlon/household_grant/cmd/api/handler"

var port = 8080

func main() {
	handler.StartServer(port)
}
