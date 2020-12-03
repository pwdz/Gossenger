package handler

import (
	"fmt"
	// "Gossenger/command"
	"Gossenger/command/types"
	"Gossenger/constants"
	"Gossenger/utils"
)

func (client *Client) readInput(){
	fmt.Printf("[#] Connection %s reading input...\n", client.conn.RemoteAddr().String())
	for{
		buffer, err := client.reader.ReadBytes(constants.Delimiter)
		
		//Let others KNOW!
		//TODOOOOOO
		if err != nil{	
			fmt.Println("[#ERROR] Failed to read socket data:", err)
			if err.Error() == "EOF"{
				fmt.Println("[#ERROR] Client disconnected unexpectedly!")
				return
			}
			continue
		}

		cmd := utils.FromBase64(buffer[0:len(buffer)-1])
		fmt.Println("[#] Message Received from "+cmd.From+" ",cmd.CmdType)
		
		if !client.isLoggedIn{
			if cmd.CmdType == types.EnterUsername || cmd.CmdType == types.Password{
				fmt.Println("lashi")
				client.in <- &cmd
			}	
		}else{
			client.in <- &cmd
		}

	}
}