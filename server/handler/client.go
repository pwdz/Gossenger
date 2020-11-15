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
		cmd := utils.FromBase64(buffer)

		if err != nil{	
			fmt.Println("[#ERROR] Failed to read socket data:", err)
			continue
		}
		
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
			case types.MsgToUser:
				// client.commands <- command{
				// 	id: cmdMsgToUser,
				// 	client: client,
				// 	args: args,
				// }
			case types.FileToUser:

			case types.MsgToGp:

			case types.FileToGp:

			case types.Quit:
			default:
			}

		}

	}
} 
func (client *Client) checkUsername(data []byte){
	username := string(data)
	username = strings.Trim(username,"\n\r ")
	//TODO
	//TODO
	//TODO
	if false{//Username already exists
		// client.sendMsg("[SERVER] Please enter password")
		client.isGuest = false;

	}else{//New Username
		client.isGuest = true;
		client.username = username 
		
		// client.sendMsg("[SERVER] Please choose a password")
	}
}
func (client *Client) checkPassword(data []byte){
	if client.isGuest{
		//Guest choosed a password
		client.isLoggedIn = true
		client.isGuest = false

		//Save username and password		

	}else{
		//Check if password is correct
		//Send the proper message
	}
}