package handler

import(
	"fmt"
	"Gossenger/utils"
	"Gossenger/constants"
	"Gossenger/command"
	"Gossenger/command/types"
)

func (client *Client) startWriteChannel(){

	for cmd := range client.out {
		fmt.Printf("[#] Sending to: %s ...\n", client.conn.RemoteAddr().String())
	
		encodedData := utils.ToBase64(*cmd)
		encodedData = append(encodedData, constants.Delimiter)

		bytesCount,err := client.conn.Write(encodedData)
		if err != nil{
			fmt.Println("[#ERROR] Failed to write data to socket")
		}

		fmt.Printf("[#] sent bytes count: %d\n", bytesCount)

	}
}

func (client *Client) send(cmd *command.Command){
		// fmt.Printf("[#] Sending to: %s ...\n", client.conn.RemoteAddr().String())
	
		// encodedData := utils.ToBase64(*cmd)
		// encodedData = append(encodedData, constants.Delimiter)

		// bytesCount,err := client.conn.Write(encodedData)
		// if err != nil{
		// 	fmt.Println("[#ERROR] Failed to write data to socket")
		// }

		// fmt.Printf("[#] sent bytes count: %d\n", bytesCount)


	client.out <- cmd;
}
func (client *Client) sendErr(err error){
	cmd := command.NewCommand(types.ServerErr, []byte(err.Error()), constants.ServerName, client.username)
	client.send(cmd)
}
func (client *Client) sendMsg(msg string){
	cmd := command.NewCommand(types.ServerMsg, []byte(msg), constants.ServerName, client.username)
	client.send(cmd)
}