package main



type server struct{
	groups map[string]*group //groupName,groupObj
	clients map[string]*client //username, clientObj
	commands chan command
}

func new() *server{
	return &server{
		groups: make(map[string]*group),
		clients: make(map[int]*client),
		commands: make(chan command),
	}
}

func (server *server) run(){
	for cmd := range server.commands{
		switch cmd.id{
		case cmdEnterUsername: 
		case cmdChangeUsername:
		case cmdJoinServer:
		case cmdGetUsersList:
		case cmdConnToUser:
		case cmdConnToGp:
		case cmdMsgToUser:
		case cmdFileToUser:
		case cmdMsgToGp:
		case cmdFileToGp: 
		case cmdQuit: 
		default:

		}
	}
}

func (server *server) enterUsername(){

}
func (server *server) changeUsername(){

}
func (server *server) joinServer(){

}
func (server *server) getUsersList(){

}
func (server *server) connectToUser(){
	
}
func (server *server) connectToGroup(){

}
func (server *server) sendMessageToUser(){

}
func (server *server) sendFileToUser(){

}
func (server *server) sendMessageToGroup(){

}
func (server *server) sendFileToGroup(){
	
}
func (server *server) quit(){

}