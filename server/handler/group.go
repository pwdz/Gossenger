package handler

import(
	"Gossenger/command"
)

type group struct{
	name string
	members map[string]*Client //addr, clientObj
	admin string
}
func NewGroup(gpName,adminName string)*group{
	return &group{
		name: gpName,
		members: make(map[string]*Client, 1),
		admin: adminName,
	}
}
func (gp *group) publish(sender *Client, cmd *command.Command){
	cmd.From = gp.name + ":" + sender.username
	for _,client := range gp.members{
		if sender.username != client.username{
			// cmd.From = gp.name +":"+ cmd.From
			client.send(cmd)
		}
	}
}
func (gp *group) addMembers(sender *Client,clients... *Client){
	// duplicateClients := make([]*Client, 0)
	for _,client := range clients{
		if _,ok := gp.members[client.username]; ok{
			//Is duplicate
			// duplicateClients = append(duplicateClients, client)
			sender.sendMsg("[ERROR] User "+client.username+" is already in group.")
		}else{
			gp.members[client.username] = client
			if sender != client{
				// sender.sendMsg("User "+client.username+" added to the group.")
				// cmd := command.NewCommand()
				// gp.publish(client, )
				client.sendMsg("You're added to group "+gp.name+" by "+sender.username)
			}
		}
	}	
	
}
func (gp *group) removeMembers(sender *Client, usernames... string){
	// notFoundClients := make([]*Client, 0)
	
	for _,username := range usernames{
		if _,ok := gp.members[username]; !ok{
			//not Found

			sender.sendMsg("[ERROR] User "+username+" doen't exist in this group.")
			// notFoundClients = append(notFoundClients, client)
		}else{
			delete(gp.members, username)
			sender.sendMsg("User "+username+" is removed from group.")
		}
		
	}

}
func (gp *group) isAdmin(username string)bool{
	return gp.admin == username
}
func (gp *group) isMember(username string)bool{
	_,ok := gp.members[username]
	return ok
}