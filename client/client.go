package main

import(
	"fmt"
	"net"
	"bufio"
	"Gossenger/constants"
	"Gossenger/pkg/utils"
	"strconv"
	"strings"
	"os"
	"Gossenger/command"
	"Gossenger/command/types"
)

type client struct{
	conn net.Conn
	username string
}

func newClient() *client{
	return &client{
		conn: nil,
		username: "",
	}
}

func (client *client) connect(){
	addr := ":" + strconv.Itoa(constants.Port)
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

func (client *client) readInput(){//from server
	fmt.Println("[$] Listening to server...")
	for true{

		input, err := bufio.NewReader(client.conn).ReadBytes(constants.Delimiter)

		if err != nil{

		}

		cmd := utils.FromBase64(input[0:len(input)-1])
		// fmt.Println("type:", cmd.CmdType)
		fmt.Println(string(cmd.Data))

	}
}
func (client *client) runConsole(){
	fmt.Println("[$] Starting Console...")
	// input := ""
	reader := bufio.NewReader(os.Stdin)
	for true{
		input,_ := reader.ReadString('\n')

		cmdStr := strings.Split(input, " ")[0]
		cmdStr = strings.Trim(cmdStr, "\n\r ")

		input = strings.Replace(input, cmdStr+" ", "", 1)
		input = strings.Replace(input, "\n", "", 1)
		fmt.Println(input)


		switch cmdStr{
		case "/username":
			client.sendUsername(input)
		case "/password":
			client.sendPassword(input)
		}
		
		// input =""
	}
}
func (client *client) send(cmd command.Command){
	encodedData := utils.ToBase64(cmd)
	encodedData = append(encodedData, constants.Delimiter)

	bytesCount,err := client.conn.Write(encodedData)
	if err != nil{
		fmt.Println("[#ERROR] Failed to write data to socket")
	}

	fmt.Printf("[#] sent bytes count: %d\n", bytesCount)
}

func (client *client) sendUsername(username string){
	cmd := command.NewCommand(types.EnterUsername, []byte(username), "", constants.ServerName)

	client.send(*cmd)
}

func (client *client) sendPassword(password string){
	cmd := command.NewCommand(types.Password, []byte(password), client.username, constants.ServerName)
	client.send(*cmd)
}