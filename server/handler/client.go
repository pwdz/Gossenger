package handler

import(
	"fmt"
	"Gossenger/command"
	"Gossenger/command/types"
	"Gossenger/constants"
	"bufio"
	"net"
	"Gossenger/pkg/utils"
	"strings"
)

//Client struct
type Client struct{
	conn net.Conn
	username string
	
	in chan<- *command.Command
	out chan<- *command.Command

	reader *bufio.Reader

	targetGp *group
	targetClient *Client

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

		targetGp: nil,
		targetClient: nil,

		isLoggedIn: false,
		isGuest: true,
	}
}

func (client *Client) send(cmd command.Command){
	fmt.Printf("[#] Sending to: %s ...\n", client.conn.RemoteAddr().String())
	
	encodedData := utils.ToBase64(cmd)
	encodedData = append(encodedData, constants.Delimiter)

	bytesCount,err := client.conn.Write(encodedData)
	if err != nil{
		fmt.Println("[#ERROR] Failed to write data to socket")
	}

	fmt.Printf("[#] sent bytes count: %d\n", bytesCount)
}
func (client *Client) readInput(){
	fmt.Printf("[#] Connection %s reading input...\n", client.conn.RemoteAddr().String())
	for{
		buffer, err := client.reader.ReadBytes(constants.Delimiter)

		if err != nil{	
			fmt.Println("[#ERROR] Failed to read socket data:", err)
			if err.Error() == "EOF"{
				fmt.Println("[#ERROR] Client disconnected unexpectedly!")
				return
			}
			continue
		}

		cmd := utils.FromBase64(buffer[0:len(buffer)-1])
		
		if !client.isLoggedIn{
			switch cmd.CmdType{
			case types.EnterUsername:
				client.checkUsername(cmd.Data)
			case types.Password:
				client.checkPassword(cmd.Data)
			}
		}else{
			switch cmd.CmdType{
			case types.ChangeUsername:
			case types.GetUsersList:
			case types.ConnToUser:
			case types.ConnToGp:
			case types.Msg:
				// client.commands <- command{
				// 	id: cmdMsgToUser,
				// 	client: client,
				// 	args: args,
				// }
			case types.File:

			case types.Quit:
			default:
			}

		}

	}
} 
func (client *Client) greetings(){
	cmd := command.NewCommand(types.EnterUsername, []byte("Welcome. Please Enter Username >/username $$username$$"), constants.ServerName, "")
	client.send(*cmd)
} 
func (client *Client) checkUsername(data []byte){
	fmt.Println("[#] Checking username...")
	username := string(data)
	username = strings.Trim(username,"\n\r ")
	//TODO
	//TODO
	//TODO
	cmd := command.NewCommand(types.UsernameSuccess, []byte{}, constants.ServerName ,username)	


	if false{//Username already exists
		// client.sendMsg("[SERVER] Please enter password")
		client.isGuest = false;
		cmd.Data = []byte("Welcome "+username+". Please enter password. >/password $$password$$")
	}else{//New Username
		client.isGuest = true;
		client.username = username 
		cmd.Data = []byte("Welcome "+username+". Please set your password. >/password $$password$$")
	}

	client.send(*cmd)
}
func (client *Client) checkPassword(data []byte){

	cmd := command.NewCommand(-1, []byte{}, constants.ServerName ,client.username)	
	if client.isGuest{
		//Guest choosed a password
		client.isLoggedIn = true
		client.isGuest = false
		
		cmd.CmdType = types.RegisterSuccess
		cmd.Data = []byte("Registeration Completed! wlc " + client.username + " :))")

		//Save username and password		

	}else{
		//Check if password is correct
		//Send the proper message

		cmd.CmdType = types.LoginSuccess
		cmd.Data = []byte("Login Completed! wlc "+ client.username +" :))")
	}

	client.send(*cmd)
}