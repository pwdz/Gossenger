package main

import(
	"fmt"
	"net"
	"bufio"
	"Gossenger/constants"
	"Gossenger/pkg/utils"
)

type client struct{
	conn net.Conn
}

func newClient() *client{
	return &client{
		conn: nil,
	}
}

func (client *client) connect(){
	addr := ":" + string(constants.Port)
	fmt.Printf("[$] Connecting to host %s ...\n", addr)
	conn, err := net.Dial(constants.ConnType, addr)

	if err != nil{
		panic("[$ERROR] Unable to Dial host:"+err.Error());
	}

	fmt.Println("[$] Connected")
	client.conn = conn

	go client.readInput()
	client.runConsole()
}

func (client *client) runConsole(){
	fmt.Println("[$] Starting Console...")
	for true{
		
	}
}

func (client *client) readInput(){//from server
	fmt.Println("[$] Listening to server...")
	for true{

		input, err := bufio.NewReader(client.conn).ReadBytes(constants.Delimiter)

		if err != nil{

		}

		req := utils.FromBase64(input)
		fmt.Println("type:", req.ReqType)
		fmt.Println("msg:", string(req.Data))

	}
}
