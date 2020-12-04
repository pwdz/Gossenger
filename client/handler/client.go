package handler

import(
	"fmt"
	"net"
	"bufio"
	"Gossenger/constants"
	"Gossenger/utils"
	"strconv"
	"strings"
	"os"
	"Gossenger/command"
	"Gossenger/command/types"
	// "math"
)

type client struct{
	conn net.Conn
	username string
	in chan []byte
	currFile []byte
}

func NewClient() *client{
	return &client{
		conn: nil,
		username: "",
		in: make(chan []byte, 500),
		currFile: nil,
	}
}

func (client *client) Connect(){
	addr := ":" + strconv.Itoa(constants.Port)
	fmt.Printf("[$] Connecting to host %s ...\n", addr)
	conn, err := net.Dial(constants.ConnType, addr)

	if err != nil{
		panic("[$ERROR] Unable to Dial host:"+err.Error());
	}

	fmt.Println("[$] Connected")
	client.conn = conn

	go client.readInput()
	go client.startReadChannel()
	client.runConsole()
}






func (client *client) runConsole(){
	fmt.Println("[$] Starting Console...")
	// input := ""
	reader := bufio.NewReader(os.Stdin)
	for true{
		input,_ := reader.ReadString('\n')

		cmdStr := strings.Split(input, " ")[0]
		cmdStr = strings.Trim(cmdStr, "\n\r ")

		input = strings.Replace(input, cmdStr+" ", "", 1)
		input = strings.Replace(input, "\n", "", 1)
		// fmt.Println(input)

		switch cmdStr{
		case "/username":
			client.sendUsername(input)
		case "/password":
			client.sendPassword(input)
		case "/connect":
			client.connectToChat(input)
		case "/send":
			client.sendMsg(input)
		case "/file":
			client.sendFile(input)	
		case "/getusers":
			client.getUsers()
		case "/creategp":
			client.createGp(input)
		case "/addmembers":
			client.addMembers(input)
		case "/removemembers":
			client.removeMembers(input)
		case "/changeusername":
			client.changeUsername(input)
		default:
			fmt.Println("[$][ERROR] Invalid command")
		}
		
		// input =""
	}
}

func (client *client) send(cmd command.Command){
	encodedData := utils.ToBase64(cmd)
	encodedData = append(encodedData, constants.Delimiter)

	_,err := client.conn.Write(encodedData)
	if err != nil{
		fmt.Println("[#ERROR] Failed to write data to socket")
	}

	// fmt.Printf("[#] sent bytes count: %d\n", bytesCount)
}

func (client *client) sendUsername(username string){
	cmd := command.NewCommand(types.EnterUsername, []byte(username), "", constants.ServerName)

	client.send(*cmd)
}

func (client *client) sendPassword(password string){
	cmd := command.NewCommand(types.Password, []byte(password), client.username, constants.ServerName)
	client.send(*cmd)
}
func (client *client) connectToChat(chatID string){
	cmd := command.NewCommand(types.Connect, []byte(chatID), client.username, constants.ServerName)
	client.send(*cmd)
}
func (client *client) sendMsg(message string){
	cmd := command.NewCommand(types.MsgTo, []byte(message), client.username, constants.ServerName)
	client.send(*cmd)

}	
func (client *client) sendFile(path string){
	data, err := ReadFile(path)
	parts := strings.Split(path, "/")
	filename := parts[len(parts)-1]
	if err != nil {
		fmt.Println("[*ERROR] open file problem:", err.Error())
		return
	}

	fileLength := len(data)

	cmd := command.NewCommand(types.FileTo, data, client.username, constants.ServerName)
	cmd.Filename = filename
	client.send(*cmd)

}
func (client *client) getUsers(){
	cmd := command.NewCommand(types.GetUsersList, []byte{}, client.username, constants.ServerName)
	client.send(*cmd)
}
func (client *client) createGp(gpName string){

	cmd := command.NewCommand(types.CreateGp, []byte(gpName), client.username, constants.ServerName)
	client.send(*cmd)
	
}
func (client *client) addMembers(members string){
	cmd := command.NewCommand(types.AddMembers, []byte(members), client.username, constants.ServerName)
	client.send(*cmd)
}
func (client *client) removeMembers(members string){
	cmd := command.NewCommand(types.RemoveMembers, []byte(members), client.username, constants.ServerName)
	client.send(*cmd)
}
func (client *client) changeUsername(newname string){
	cmd := command.NewCommand(types.ChangeUsername, []byte(newname), client.username, constants.ServerName)
	client.send(*cmd)
}