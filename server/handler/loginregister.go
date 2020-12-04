package handler

import(
	"fmt"
	"Gossenger/constants"
	"Gossenger/command"
	"Gossenger/command/types"
	"strings"


)

func (server *server) checkUsername(cmd command.Command, client *Client){
	fmt.Printf("[#] Checking username... [Conn: %s]", client.conn.RemoteAddr())

	username := string(cmd.Data)
	username = strings.Trim(username,"\n\r ")

	respMsg := "" 
	if server.db.DoesExist(username){//Username already exists
		client.isGuest = false;
		client.username = username 
		respMsg = "Welcome back "+username+". Please enter password. >/password $$password$$"
	}else{//New Username
		client.isGuest = true;
		client.username = username 
		respMsg = "Welcome "+username+". Please set your password. >/password $$password$$"
	}

	client.sendMsg(respMsg)
}

func (server *server) checkPassword(cmd command.Command, client *Client){
	fmt.Println("[#] Checking password...")
	password := string(cmd.Data)

	respCmd := command.NewCommand(-1, []byte{}, constants.ServerName ,client.username)	

	if client.isGuest{
		//Guest choosed a password
		client.isLoggedIn = true
		client.isGuest = false
		
		respCmd.CmdType = types.RegisterSuccess
		respCmd.Data = []byte(client.username)

		//Save username and password		
		server.db.AddData(client.username, password)

		server.loginSuccess(client)
	}else{
		//Check if password is correct
		//Send the proper message

		if password == server.db.GetPassword(client.username){
			respCmd.CmdType = types.LoginSuccess
			respCmd.Data = []byte(client.username)
			
			server.loginSuccess(client)
			client.isLoggedIn = true
			client.isGuest = false
	

		}else{
			respCmd.CmdType = types.Failure
			respCmd.Data = []byte("")		
		}
	}

	client.send(respCmd)
}
func (server *server) loginSuccess(client *Client){
	for _,user := range server.clients{
		user.sendMsg(client.username+" Joined the server!")
	}
	server.clients[client.username] = client
}