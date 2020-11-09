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

func (client *client) sendMsg(msg string){
	client.conn.Write([]byte(msg))
}
func (client *client) sendErr(err error) {
	client.conn.Write([]byte("err: " + err.Error() + "\n"))
}
func (client *client) readInput(){
	for{

		msg, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil{
			return
		}

		msg = strings.Trim(msg,"\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
	
		if !client.isLoggedIn{
			switch cmd{
			case "/username":
				client.checkUsername(strings.Join(args[1:len(args)]," "))
			case "/password":
				client.checkPassword(strings.Join(args[1:len(args)]," "))
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
				client.sendErr(fmt.Errorf("[FROM SERVER] Unkown command: %s", cmd))
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
		client.sendMsg("[FROM SERVER] Please enter password")
		client.isGuest = false;

		client.readInput()
	}else{//New Username
		client.isGuest = true;
		client.username = username 
		
		client.sendMsg("[FROM SERVER] Please choose a password")
		
		client.readInput()
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