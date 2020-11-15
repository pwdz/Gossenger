package main

import(
	// "net"
	// "fmt"
	"Gossenger/server/handler"
)

func main()  {
	
	server := handler.NewServer()	
	server.StartListening()


}