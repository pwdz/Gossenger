package handler

import(
	"Gossenger/command"
	"Gossenger/command/types"
	"strings"
	"fmt"
)

func (client *Client) startListenChannel(server *server){
	fmt.Println("[#] Listening to channel")
	for cmd := range client.in{
		fmt.Printf("[#] Command in listener for client %s\n", client.conn.RemoteAddr())
		switch cmd.CmdType{
		case types.EnterUsername:
			server.checkUsername(*cmd, client)
		case types.Password:
			server.checkPassword(*cmd, client)
		case types.ChangeUsername:
			// server.connectToUser(cmd)
		case types.GetUsersList:
			// server.connectToGroup()
		case types.Connect:
			server.connectTo(*cmd)
		// case types.ConnToGp:
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

		fmt.Println("[#] sending msg to", sender.targetChatID, sender.isTargetGp)
		if !sender.isTargetGp{

			receiver, okReceiver := server.clients[sender.targetChatID]
			if okReceiver{
				receiver.send(cmd)	
			}else{

			}

		}else{
			group, ok := server.groups[sender.targetChatID]
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
func (server *server) connectTo(cmd command.Command){
	client, ok := server.clients[cmd.From]
	if !ok{
		fmt.Println("kiriiiiiiii", cmd.From)
		return
	}	
	targetChatID := string(cmd.Data)
	targetChatID = strings.Trim(targetChatID,"\n\r ")


	_, okClient := server.clients[targetChatID]
	_, okGroup := server.groups[targetChatID]
	if okClient || okGroup{
		fmt.Println("=>>>>>>>>."+targetChatID)
		client.targetChatID = targetChatID
		client.isTargetGp = okGroup	
	}else{
		//ridi no chat!
	}
}
func (server *server) quit(){

}