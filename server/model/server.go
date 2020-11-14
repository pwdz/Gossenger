package model

type server struct{
	groups map[string]*group 
	clients []*client 
	commands chan command
}

func newServer() *server{
	return &server{
		groups: make(map[string]*group),
		clients: make([]*client, 1),
		commands: make(chan command),
	}
}