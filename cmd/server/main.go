package main

import(
	// "net"
	// "fmt"
)

func main()  {
	server := newServer()	
	// go server.run()

	server.startListening()

}