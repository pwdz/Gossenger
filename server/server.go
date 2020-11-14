package main

import(
	"fmt"
	"net"
	"strconv"
	"strings"
)
const(
	port = 9000
)

type server struct{
	groups map[string]*group //groupName,groupObj
	clients []*client //clientObj
	commands chan command
}

func newServer() *server{
	fmt.Println("[#] Creating server struct...")
	return &server{
		groups: make(map[string]*group),
		clients: make([]*client, 1),
		commands: make(chan command),
	}
}
func (server *server) startListening(){
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
		conn.Write([]byte("sa;am chaghalll =))\n"))
		go server.newConn(conn)
	}	
}
func (server *server) newConn(conn net.Conn){
	newGuest := &client{
		conn: conn,
		username: "",
		commands: server.commands , 
		gp: nil,
		targetClient: nil,
		isLoggedIn: false,
		isGuest: true,
	}
	fmt.Printf("[#] New connection. addr:%s\n", conn.RemoteAddr().String());

	newGuest.sendMsg("[FROM SERVER] Please enter username")
	newGuest.readInput()
}
func (server *server) run(){
	fmt.Println("[#] Listening to channel")
	for cmd := range server.commands{
		switch cmd.id{
		case cmdChangeUsername:
			server.changeUsername()
		case cmdGetUsersList:
			server.getUsersList(cmd)
		case cmdConnToUser:
			server.connectToUser(cmd)
		case cmdConnToGp:
			server.connectToGroup()
		case cmdMsgToUser:
			server.sendMessageToUser(cmd)
		case cmdFileToUser:
			server.sendFileToUser()
		case cmdMsgToGp:
			server.sendMessageToGroup()
		case cmdFileToGp: 
			server.sendFileToGroup()
		case cmdQuit: 
			server.quit()
		default:
			server.error()
		}
	}
}

func (server *server) enterUsername(){

}
func (server *server) changeUsername(){

}
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
}