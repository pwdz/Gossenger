package handler

import(
	"fmt"
	"net"
	"strconv"
	"strings"
	"Gossenger/command"
	"bufio"
	"os"
)
const(
	port = 9000
)

type server struct{
	groups map[string]*group //gpName, group struct
	clients map[string]*Client//clientName,Client struct 
	commands chan *command.Command
}

func NewServer() *server{
	return &server{
		groups: make(map[string]*group),
		clients: make(map[string]*Client),
		commands: make(chan *command.Command, 50),
	}
}

func (server *server) StartListening(){
	go server.runConsole()

	listener, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	if err != nil{
		fmt.Println("[#ERROR] Unable to listen: " + err.Error())
	}

	fmt.Printf("[#] Listening on: %s\n", listener.Addr().String());
	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err!=nil{
			fmt.Println("[#ERROR] Failed to accept connecton: " + err.Error())
			continue
		}
		go server.newConn(conn)
	}	
}
func (server *server) newConn(conn net.Conn){
	fmt.Printf("[#] New connection. addr:%s\n", conn.RemoteAddr().String());

	newGuest := NewClient(conn)

	newGuest.greetings()
	go newGuest.readInput()
	go newGuest.startListenChannel(server)
	go newGuest.startWriteChannel()
}
func (server *server) runConsole(){
	fmt.Println("[$] Starting Console...")
	// input := ""
	reader := bufio.NewReader(os.Stdin)
	for true{
		input,_ := reader.ReadString('\n')

		cmdStr := strings.Split(input, " ")[0]
		cmdStr = strings.Trim(cmdStr, "\n\r ")

		input = strings.Replace(input, cmdStr+" ", "", 1)
		input = strings.Replace(input, "\n", "", 1)
		
		if cmdStr == "/get"{
			fmt.Println("##################Online Users##################")
			index := 1;
			for username := range server.clients{
				fmt.Println(index,">",username)
				index++
			}
		}
		
	}
}