package main

import(
	"net"
)

type group struct{
	name string
	members map[net.Addr]*client //addr, clientObj
}

func (gp *group) publishMsg(sender *client, msg string){
	for addr,client := range gp.members{
		if sender.conn.RemoteAddr() != addr{
			client.sendMsg(msg)
		}
	}
}

//TODO
func (gp *group) publishFile(){

}