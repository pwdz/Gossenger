package main

import(
	"fmt"
	"net"
	"bufio"
)

type client struct{
	conn net.Conn
}

func newClient() *client{
	return &client{
		conn: nil,
	}
}

func (client *client) connect(address string){
	fmt.Printf("[$] Connecting to host %s ...\n", address)
	conn, err := net.Dial("tcp4", address)

	if err != nil{
		fmt.Println("[$ERROR] Unable to Dial host: " + err.Error())
	}

	fmt.Println("[$] Connected")
	client.conn = conn

	go readInput(conn)
	client.runConsole()
}

func (client *client) runConsole(){
	fmt.Println("[$] Starting Console...")
	client.conn.Write([]byte("salammmmmmmmmm"))
	for true{

	}
}

func readInput(conn net.Conn){//from server
	fmt.Println("[$] Listening to server...")
	if conn == nil{
		fmt.Println("koskeshhhhhhhhhh")
	}
	// var input []byte
	for true{
		input, err := bufio.NewReader(conn).ReadString('\n')
		if(err != nil){
			fmt.Println("[#ERROR] read: " + err.Error())
			continue
		}
		// if byteCount > 0{ 
		// 	fmt.Println("::", string(input))
		// }

		fmt.Println(input)

	}
}
