package handler

import(
	"fmt"
	"bufio"
	"Gossenger/command"
	"Gossenger/command/types"
	"Gossenger/utils"
	"Gossenger/constants"
	"io/ioutil"
	"strings"
)

func (client *client) readInput(){//from server
	fmt.Println("[$] Listening to server...")
	for true{

		input, err := bufio.NewReader(client.conn).ReadBytes(constants.Delimiter)

		if err != nil{

		}
		client.in <- input 

	}
}
func (client *client) startReadChannel(){
	for inputBytes := range client.in{
		cmd := utils.FromBase64(inputBytes[0:len(inputBytes)-1])

		fmt.Println("[",len(inputBytes) ,"] type:", cmd.CmdType)
		
		switch cmd.CmdType{
		case types.EnterUsername, types.ServerMsg, types.MsgTo:
			fmt.Println(string(cmd.Data))
		case types.RegisterSuccess:
			client.usernameSuccess(cmd)
		case types.LoginSuccess:
			client.usernameSuccess(cmd)
		case types.FileTo:
			client.receiveFile(cmd)
		case types.GetUsersList:
			client.receiveUsersList(cmd)
		}

	}
}





func (client *client) usernameSuccess(cmd command.Command){
	client.username = string(cmd.Data)
	fmt.Println("[*] username set successfully: ", client.username)
}

func (client *client) receiveFile(cmd command.Command){
	err := ioutil.WriteFile("/home/pwdz/cache/"+cmd.Filename, cmd.Data, 0666)
	if err != nil { 
		// handle error
		fmt.Println(err.Error())
	}

}
func (client *client) receiveUsersList(cmd command.Command){
	fmt.Println("kioni",string(cmd.Data))
	clientNames := string(cmd.Data)
	for index,username := range strings.Split(clientNames,"\n"){
		if username!=""{
			fmt.Println(index,">",username)
		}
	}

}