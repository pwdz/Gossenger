package handler

import(
	"fmt"
	"Gossenger/request"
	"Gossenger/request/types"
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
	
	in chan<- *request.Request
	out chan<- *request.Request

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
		
		in: make(chan *request.Request, 50),
		out: make(chan *request.Request, 50),
		
		reader: bufio.NewReader(conn),

		targetGp: nil,
		targetClient: nil,

		isLoggedIn: false,
		isGuest: true,
	}
}

func (client *Client) send(req request.Request){
	fmt.Printf("[#] Sending to: %s ...\n", client.conn.RemoteAddr().String())
	
	encodedData := utils.ToBase64(req)
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
		req := utils.FromBase64(buffer)

		if err != nil{	
			fmt.Println("[#ERROR] Failed to read socket data:", err)
			continue
		}
		
		if !client.isLoggedIn{
			switch req.ReqType{
			case types.EnterUsername:
				client.checkUsername(req.Data)
			case types.Password:
				client.checkPassword(req.Data)
			}
		}else{
			switch req.ReqType{
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