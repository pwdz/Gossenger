package handler

import(
	"Gossenger/command"
)

type group struct{
	name string
	members []*Client //addr, clientObj
}
func (gp *group) publish(sender *Client, cmd *command.Command){
	for _,client := range gp.members{
		if sender.username != client.username{
			client.send(cmd)
		}
	}
}