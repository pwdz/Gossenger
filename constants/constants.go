package constants

const(
	//Delimiter for reading/writing in client/server
	Delimiter byte = '\n'
	//Port for client/server
	Port int = 9000
	//ConnType for socket
	ConnType string = "tcp4"
	//MaxStreamSize the maximum allowed stream size 
	MaxStreamSize int = 4096
	//ServerName ...
	ServerName string = "SERVER"
	BasePath string = "/home/pwdz/cache/"
	DBBasePath string = "/home/pwdz/GossengerServer/"
	DBName string = "Users.db"
)