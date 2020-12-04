package handler

import(
	"fmt"
	"bufio"
	"Gossenger/command"
	"Gossenger/command/types"
	"Gossenger/utils"
	"Gossenger/constants"
)

func (client *client) readInput(){//from server
	fmt.Println("[$] Listening to server...")
	for true{

		input, err := bufio.NewReader(client.conn).ReadBytes(constants.Delimiter)

		if err != nil{

		}

		fmt.Println("[",len(input) ,"]")
		client.in <- input 


	}
}
func (client *client) startReadChannel(){
	for inputBytes := range client.in{
		cmd := utils.FromBase64(inputBytes[0:len(inputBytes)-1])
		// fmt.Println("type:", cmd.CmdType)
		// fmt.Println(string(cmd.Data), cmd.CmdType)
		// fmt.Println("[",len(inputBytes) ,"] type:", cmd.CmdType, "follow:",cmd.FollowPackets)
		switch cmd.CmdType{
		case types.RegisterSuccess:
			client.usernameSuccess(cmd)
		case types.LoginSuccess:
			client.usernameSuccess(cmd)
		}

	}
}





func (client *client) usernameSuccess(cmd command.Command){
	client.username = string(cmd.Data)
	fmt.Println("[*] username set successfully: ", client.username)
}