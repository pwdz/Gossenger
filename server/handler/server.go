package handler

import(
	"fmt"
	"net"
	"strconv"
	// "strings"
	"Gossenger/request"
	"Gossenger/request/types"
)
const(
	port = 9000
)

type server struct{
	groups map[string]*group 
	clients []*Client 
	requests chan *request.Request
}

func NewServer() *server{
	return &server{
		groups: make(map[string]*group),
		clients: make([]*Client, 1),
		requests: make(chan *request.Request, 50),
	}
}

func (server *server) StartListening(){
	go server.run()

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

	newGuest := NewClient(conn)

	t := types.ConnToUser
	// fmt.Println(t)
	msg := "**____*****salamamamamamam arrrrrrrrrrrrrrr"

	req := request.NewReq(t, []byte(msg))

	// newGuest.conn.Write([]byte(msg))
	newGuest.send(*req)
	// newGuest.readInput()
}
func (server *server) run(){
	fmt.Println("[#] Listening to channel")
	// for cmd := range server.commands{
	// 	switch cmd.id{
	// 	case cmdChangeUsername:
	// 		server.changeUsername()
	// 	case cmdGetUsersList:
	// 		server.getUsersList(cmd)
	// 	case cmdConnToUser:
	// 		server.connectToUser(cmd)
	// 	case cmdConnToGp:
	// 		server.connectToGroup()
	// 	case cmdMsgToUser:
	// 		server.sendMessageToUser(cmd)
	// 	case cmdFileToUser:
	// 		server.sendFileToUser()
	// 	case cmdMsgToGp:
	// 		server.sendMessageToGroup()
	// 	case cmdFileToGp: 
	// 		server.sendFileToGroup()
	// 	case cmdQuit: 
	// 		server.quit()
	// 	default:
	// 		server.error()
	// 	}
	// }
}

func (server *server) enterUsername(){

}
func (server *server) changeUsername(){

}
/*
func (server *server) getUsersList(cmd command){
	var clients []string
	
	for _, client := range server.clients{
		clients = append(clients, client.username)
	}

	cmd.client.sendMsg(fmt.Sprintf("[FROM SERVER] Online users:\n %s", strings.Join(clients, "\n")))
}
func (server *server) connectToUser(cmd command){
	
	targetUsername := strings.Join(cmd.args[1:len(cmd.args)], " ")
	targetUsername = strings.TrimSpace(targetUsername)

	targetClient := server.findClientByUsername(targetUsername)

	if targetClient != nil{
		cmd.client.targetClient = targetClient
		cmd.client.sendMsg("[FROM SERVER] connected to " + targetUsername)
	}else{
		cmd.client.sendErr(fmt.Errorf("[FROM SERVER] Not a valid username: %s", targetUsername))
	}	

}
func (server *server) connectToGroup(){

}
func (server *server) sendMessageToUser(cmd command){
	if cmd.client.targetClient != nil{

		msg := strings.Join(cmd.args[1:len(cmd.args)], " ")
		msg = strings.TrimSpace(msg)

		cmd.client.targetClient.sendMsg(cmd.client.username+": " +msg)

	}else{
		cmd.client.sendErr(fmt.Errorf("[FROM SERVER] there is no connection to any user"))
	}
}
func (server *server) sendFileToUser(){

}
func (server *server) sendMessageToGroup(){

}
func (server *server) sendFileToGroup(){
	
}
func (server *server) quit(){

}
func (server *server) error(){

}
func (server *server) findClientByUsername(username string) *client {
	for _, client := range server.clients{
		if client.username == username{
			return client
		}
	}
	return nil
}*/