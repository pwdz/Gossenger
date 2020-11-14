package model

import (
	// "fmt"
	"net"
	"bufio"
	// "strings"

)

type client struct{
	conn net.Conn
	username string
	in chan<- command  
	out chan<- command
	reader *bufio.Reader
	writer *bufio.Writer
	targetGp *group
	targetClient *client
	isLoggedIn bool
	isGuest bool
}
func NewClient(conn net.Conn)*client{
	return &client{
		conn: conn,
		username: "",
		in: make(chan )



	}
}
