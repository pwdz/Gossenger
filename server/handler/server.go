package handler

import(
	"fmt"
	"net"
	"strconv"
	// "strings"
	"Gossenger/command"
	// "Gossenger/command/types"
)
const(
	port = 9000
)

type server struct{
	groups map[string]*group //gpName, group struct
	clients map[string]*Client//clientName,Client struct 
	commands chan *command.Command
}

func NewServer() *server{
	return &server{
		groups: make(map[string]*group),
		clients: make(map[string]*Client),
		commands: make(chan *command.Command, 50),
	}
}

func (server *server) StartListening(){
	// go server.run()

	listener, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	if err != nil{
		fmt.Println("[#ERROR] Unable to listen: " + err.Error())
	}

	fmt.Printf("[#] Listening on: %s\n", listener.Addr().String());
	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err!=nil{
			fmt.Println("[#ERROR] Failed to accept connecton: " + err.Error())
			continue
		}
		go server.newConn(conn)
	}	
}
func (server *server) newConn(conn net.Conn){
	fmt.Printf("[#] New connection. addr:%s\n", conn.RemoteAddr().String());

	newGuest := NewClient(conn, server.commands)

	newGuest.greetings()
	go newGuest.readInput()
	go newGuest.listen(server)
	go newGuest.startWriteChannel()
}
// func (server *server) run(){
// 	fmt.Println("[#] Listening to channel")
// 	for cmd := range server.commands{
// 		switch cmd.CmdType{
// 		case types.EnterUsername:
// 			// server.checkUsername(*cmd)
// 		case types.Password:
// 			// server.checkPassword(*cmd)
// 		case types.ChangeUsername:
// 			// server.connectToUser(cmd)
// 		case types.GetUsersList:
// 			// server.connectToGroup()
// 		case types.ConnToUser:
// 			// server.sendMessageToUser(cmd)
// 		case types.ConnToGp:
// 			// server.sendFileToUser()
// 		case types.CreateGp:
// 			// server.sendMessageToGroup()
// 		case types.AddMember: 
// 			// server.sendFileToGroup()
// 		case types.MsgTo: 
// 			// server.quit()
// 		case types.FileTo: 
// 		case types.Quit:
// 		default:
// 			// server.error()
// 		}
// 	}
// }

