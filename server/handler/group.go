package handler


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