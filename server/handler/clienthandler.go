package handler

import(
	"Gossenger/command"
	"Gossenger/command/types"
	"strings"
	"fmt"
)

func (client *Client) listen(server *server){
	fmt.Println("[#] Listening to channel")
	for cmd := range client.in{
		fmt.Printf("[#] Command in listener for client %s\n", client.username)
		switch cmd.CmdType{
		case types.EnterUsername:
			server.checkUsername(*cmd, client)
		case types.Password:
			server.checkPassword(*cmd, client)
		case types.ChangeUsername:
			// server.connectToUser(cmd)
		case types.GetUsersList:
			// server.connectToGroup()
		case types.ConnToUser:
			// server.sendMessageToUser(cmd)
		case types.ConnToGp:
			// server.sendFileToUser()
		case types.CreateGp:
			// server.sendMessageToGroup()
		case types.AddMember: 
			// server.sendFileToGroup()
		case types.MsgTo: 
			server.sendMsg(cmd)
			// server.quit()
		case types.FileTo: 
		case types.Quit:
		default:
			// server.error()
		}
	}
}

func (server *server) changeUsername(cmd command.Command){
	client, ok := server.clients[cmd.From]
	if !ok{
		//Invalid client!
		return
	}

	newUsername := string(cmd.Data)
	newUsername = strings.Trim(newUsername, "\n\r ")

	// respCmd := command.NewCommand()

	//Check usernames in db
	if false{//Username already exists

	}else{//Username is not taken
		client.username = newUsername
		//save in db
	}

}
func (server *server) sendMsg(cmd *command.Command){
	sender, okSender := server.clients[cmd.From]
	if okSender{
		client, okReceiver := server.clients[cmd.To]
		if okReceiver{
			client.send(cmd)	
		}else{
			group, ok := server.groups[cmd.To]
			if ok{
				group.publish(sender, cmd)
			}else{
				//ridiiiiii!
			}
		}
	}else{

	}



}
func (server *server) getUsersList(cmd command.Command){

}
func (server *server) connectToUser(cmd command.Command){

}
func (server *server) connectToGroup(cmd command.Command){

}
func (server *server) quit(){

}