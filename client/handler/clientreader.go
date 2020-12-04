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
	"os"
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
		
		switch cmd.CmdType{
		case types.EnterUsername: 
			client.receiveEnterUsername(cmd)
		case types.ServerMsg: 
			client.receiveServerMsg(cmd)
		case types.MsgTo:
			client.receiveMsgTo(cmd)
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


func (client *client) receiveEnterUsername(cmd command.Command){
	fmt.Println("["+cmd.From+"] "+string(cmd.Data))
}
func (client *client) receiveServerMsg(cmd command.Command){

	fmt.Println("["+cmd.From+"] "+string(cmd.Data))
}
func (client *client) receiveMsgTo(cmd command.Command){

	fmt.Println("["+cmd.From+"] "+string(cmd.Data))
}



func (client *client) usernameSuccess(cmd command.Command){
	client.username = string(cmd.Data)
	fmt.Println("["+cmd.From+"] username set successfully: "+client.username)
}

func (client *client) receiveFile(cmd command.Command){
	fmt.Println("["+cmd.From+"] file>> "+ cmd.Filename)

	dirPath := constants.BasePath + client.username
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}
	
	err := ioutil.WriteFile(dirPath+"/"+cmd.Filename, cmd.Data, 0666)
	if err != nil { 
		// handle error
		fmt.Println(err.Error())
	}

}
func (client *client) receiveUsersList(cmd command.Command){
	clientNames := string(cmd.Data)
	for index,username := range strings.Split(clientNames,"\n"){
		if username!=""{
			fmt.Println("["+cmd.From+"] ",index,":",username)
		}
	}
}