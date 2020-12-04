package handler

import(
	"Gossenger/command"
	"Gossenger/command/types"
	"strings"
	"fmt"
	"Gossenger/constants"
)

func (client *Client) startListenChannel(server *server){
	fmt.Println("[#] Listening to channel")
	for cmd := range client.in{
		// fmt.Printf("[#] Command in listener for client %s\n", client.conn.RemoteAddr())
		switch cmd.CmdType{
		case types.EnterUsername:
			server.checkUsername(*cmd, client)
		case types.Password:
			server.checkPassword(*cmd, client)
		case types.ChangeUsername:
			// server.connectToUser(cmd)
		case types.GetUsersList:
			server.sendUsersList(*cmd)
		case types.Connect:
			server.connectTo(*cmd)
		// case types.ConnToGp:
			// server.sendFileToUser()
		case types.CreateGp:
			server.createGp(*cmd)
		case types.AddMembers: 
			server.addClientsToGp(*cmd)
		case types.RemoveMembers: 
			server.removeClientsFromGp(*cmd)
		case types.MsgTo: 
			server.sendMsg(cmd)
			// server.quit()
		case types.FileTo: 
			// server.sendFle()
			server.sendMsg(cmd)
		case types.Quit:
			server.quitClient(*cmd)
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

		// fmt.Println("[#] sending msg to", sender.targetChatID, sender.isTargetGp)
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
				//
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
		fmt.Println("[#ERROR] sender 404", cmd.From)
		return
	}	
	targetChatID := string(cmd.Data)
	targetChatID = strings.Trim(targetChatID,"\n\r ")


	_, okClient := server.clients[targetChatID]
	_, okGroup := server.groups[targetChatID]
	if okClient || okGroup{
		fmt.Println(">>>>>>>>"+targetChatID)
		client.targetChatID = targetChatID
		client.isTargetGp = okGroup	
	}else{
		//ridi no chat!
	}
}
func (server *server) sendUsersList(cmd command.Command){
	client, ok := server.clients[cmd.From]
	if !ok{
		fmt.Println("[#ERROR] sender 404", cmd.From)
		return
	}
	var usersList string = ""
	for clientName, user := range server.clients{
		if client != user{
			usersList +=  clientName+"\n"
		}
	}

	respCmd := command.NewCommand(types.GetUsersList, []byte(usersList), constants.ServerName, client.username)
	client.send(respCmd)
}
func (server *server) quitClient(cmd command.Command){
	delete(server.clients, cmd.From)
	fmt.Println("[#] client",cmd.From,"left!")
	for _,user := range server.clients{
		user.sendMsg(cmd.From+" left the server!")
	}

}
func (server *server)createGp(cmd command.Command){
	gpName := string(cmd.Data)

	client, ok := server.clients[cmd.From]
	if !ok{
		fmt.Println("[#ERROR] sender 404", cmd.From)
		return
	}
	if _, ok := server.groups[gpName]; !ok{
		//not a duplicate gp name
		newGroup := NewGroup(gpName, cmd.From)
		newGroup.addMembers(client,client)
		server.groups[gpName] = newGroup 
		fmt.Println("[#] Group",gpName,"Created by",cmd.From,".")
		client.sendMsg("Group "+gpName+" created.")		
	}else{
		//ridi duplicate

		client.sendMsg("[ERROR] A group with this name already exist!")
	}
}
func (server *server) addClientsToGp(cmd command.Command){
	input := string(cmd.Data)
	members := strings.Split(input," ")
	gpName := members[0]
	members = members[1:]

	client, ok := server.clients[cmd.From]
	if !ok{
		fmt.Println("[#ERROR] sender 404", cmd.From)
		return
	}
	if group, ok := server.groups[gpName]; ok{
		if group.isMember(client.username){
			//we can add
			toBeAddedClients := make([]*Client,0)
			for _,username := range members{
				user, ok1 := server.clients[username]
				if !ok1{
					client.sendMsg("[#ERROR] user "+username+" doesn't exist")
					// return
				}else{
					toBeAddedClients = append(toBeAddedClients, user)
				}
			}
			group.addMembers(client, toBeAddedClients...)

		}else{
			client.sendMsg("[#ERROR] You are not a member of this group")
		}
	}else{
		client.sendMsg("[#ERROR] such a group doesn't exist")
	}

}
func (server *server) removeClientsFromGp(cmd command.Command){
	//Only if admin!
	client, ok := server.clients[cmd.From]
	if !ok{
		fmt.Println("[#ERROR] sender 404", cmd.From)
		return
	}

	input := string(cmd.Data)
	members := strings.Split(input," ")
	gpName := members[0]
	members = members[1:]

	if group, ok := server.groups[gpName]; ok{
		if group.isAdmin(client.username){
			//we can remove
			group.removeMembers(client, members...)
		}else{
			client.sendMsg("[#ERROR] You are not the ADMIN of this group")
		}
	}else{
		client.sendMsg("[#ERROR] such a group doesn't exist")
	}



}