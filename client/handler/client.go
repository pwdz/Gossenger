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
	"math"
)

type client struct{
	conn net.Conn
	username string
	in chan []byte
}

func NewClient() *client{
	return &client{
		conn: nil,
		username: "",
		in: make(chan []byte, 500),
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
		fmt.Println(input)


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
		}
		
		// input =""
	}
}
func (client *client) send(cmd command.Command){
	encodedData := utils.ToBase64(cmd)
	encodedData = append(encodedData, constants.Delimiter)

	bytesCount,err := client.conn.Write(encodedData)
	if err != nil{
		fmt.Println("[#ERROR] Failed to write data to socket")
	}

	fmt.Printf("[#] sent bytes count: %d\n", bytesCount)
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
	if err != nil {
		fmt.Println("file read riiiiiiiid", err.Error())
		return
	}

	fileLength := len(data)
	var packetsCount int = int(math.Ceil(float64(fileLength)/float64(constants.MaxStreamSize)))
	var packetSize int = constants.MaxStreamSize
	for i := 0; i < packetsCount; i++ {
		fmt.Println("[*][FILE] sending packet",i)
		beginIndex := i*packetSize
		
		var endIndex int
		if fileLength <= (i+1)*packetSize{
			 endIndex = fileLength
		}else{
			 endIndex = (i+1)*packetSize
		}
		
		cmd := command.NewCommand(types.FileTo, data[beginIndex:endIndex], client.username, constants.ServerName)
		cmd.FollowPackets = packetsCount - 1 - i
		client.send(*cmd)
	}


}