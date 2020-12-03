package handler

import(
	"fmt"
	"Gossenger/constants"
	"Gossenger/command"
	"Gossenger/command/types"
	"strings"


)

func (server *server) checkUsername(cmd command.Command, client *Client){
	fmt.Println("[#] Checking username...")

	username := string(cmd.Data)
	username = strings.Trim(username,"\n\r ")
	//TODO
	//TODO
	//TODO
	respMsg := "" 
	fmt.Println("Do this shit")
	if false{//Username already exists
		// client.sendMsg("[SERVER] Please enter password")
		client.isGuest = false;
		respMsg = "Welcome "+username+". Please enter password. >/password $$password$$"
	}else{//New Username
		client.isGuest = true;
		client.username = username 
		respMsg = "Welcome "+username+". Please set your password. >/password $$password$$"
	}

	fmt.Println("Do this shit2")
	client.sendMsg(respMsg)
}

func (server *server) checkPassword(cmd command.Command, client *Client){
	fmt.Println("[#] Checking password...")
	password := string(cmd.Data)

	respCmd := command.NewCommand(-1, []byte{}, constants.ServerName ,client.username)	

	if true{
		//Guest choosed a password
		client.isLoggedIn = true
		client.isGuest = false
		
		respCmd.CmdType = types.RegisterSuccess
		respCmd.Data = []byte("Registeration Completed! wlc " + client.username + " :))")

		//Save username and password		
		server.loginSuccess(client)
	}else{
		//Check if password is correct
		//Send the proper message
		if password == "1234"{
			respCmd.CmdType = types.LoginSuccess
			respCmd.Data = []byte("Login Completed! wlc "+ client.username +" :))")
			
			server.loginSuccess(client)

		}else{
			respCmd.CmdType = types.Failure
			respCmd.Data = []byte("Wrong Password!")
		
		}
	}

	client.send(respCmd)
}
func (server *server) loginSuccess(client *Client){
	server.clients[client.username] = client
}