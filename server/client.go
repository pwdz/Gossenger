package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

type client struct{
	conn net.Conn
	username string
	commands chan<- command  
	gp *group
	targetClient *client
	isLoggedIn bool
	isGuest bool
}

func (client *client) sendMsg(input string){
	fmt.Printf("[#] Sending message to: %s|Msg: %s\n", client.conn.RemoteAddr().String(), input)
	client.conn.Write([]byte(input+"\n"))
}
func (client *client) sendErr(err error) {
	client.conn.Write([]byte("err: " + err.Error() + "\n"))
}
func (client *client) readInput(){
	fmt.Printf("[#] Connection %s reading input...\n", client.conn.RemoteAddr().String())
	for{

		// client.conn.Read();
		buffer, err := bufio.NewReader(client.conn).ReadBytes('\n')
		input := string(buffer)
		if err != nil{
			return
		}
		fmt.Printf("[#] Input from user:%s\n", input)

		input = strings.Trim(input,"\r\n")
		args := strings.Split(input, " ")
		cmd := strings.TrimSpace(args[0])
		input = strings.Join(args[1:len(args)]," ")
	
		if !client.isLoggedIn{
			switch cmd{
			case "/username":
				client.checkUsername(input)
			case "/password":
				client.checkPassword(input)
			}
		}else{
			switch cmd{
			case "/changeusername":

			case "/getuserslist":
				client.commands <- command{
					id: cmdGetUsersList,
					client: client,
				}
			case "/conntouser":
				client.commands <- command{
					id: cmdConnToUser,
					client: client,
					args: args,
				}

			case "/conntogp":

			case "/msgtouser":
				client.commands <- command{
					id: cmdMsgToUser,
					client: client,
					args: args,
				}
			case "/filetouser":

			case "/msgtogp":

			case "/filetogp":

			case "/quit":
			default:
				client.sendErr(fmt.Errorf("[SERVER] Unkown command: %s", cmd))
			}

		}

	}
} 
func (client *client) checkUsername(username string){
	username = strings.TrimSpace(username)
	//TODO
	//TODO
	//TODO
	if false{//Username already exists
		client.sendMsg("[SERVER] Please enter password")
		client.isGuest = false;

	}else{//New Username
		client.isGuest = true;
		client.username = username 
		
		client.sendMsg("[SERVER] Please choose a password")
	}
}
func (client *client) checkPassword(password string){
	if client.isGuest{
		//Guest choosed a password
		client.isLoggedIn = true
		client.isGuest = false

		//Save username and password		

	}else{
		//Check if password is correct
		//Send the proper message
	}
}