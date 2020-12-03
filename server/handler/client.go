package handler

import(
	// "fmt"
	"Gossenger/command"
	"Gossenger/command/types"
	"Gossenger/constants"
	"bufio"
	"net"
	// "Gossenger/pkg/utils"
	// "strings"
)

//Client struct
type Client struct{
	conn net.Conn
	username string
	
	in chan *command.Command
	out chan *command.Command

	reader *bufio.Reader

	isTargetGp bool
	targetChatID string

	isLoggedIn bool
	isGuest bool
}

//NewClient returns a new client struct
func NewClient(conn net.Conn)*Client{
	return &Client{
		conn: conn,
		username: "",
		
		in: make(chan *command.Command, 50),
		out: make(chan *command.Command, 50),
		
		reader: bufio.NewReader(conn),

		isTargetGp: false,
		targetChatID: "",

		isLoggedIn: false,
		isGuest: true,
	}
}
 
func (client *Client) greetings(){
	cmd := command.NewCommand(types.EnterUsername, []byte("Welcome. Please Enter Username >/username $$username$$"), constants.ServerName, "")
	client.send(cmd)
} 
