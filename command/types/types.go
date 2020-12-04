package types

//TypeID is for request types
type TypeID int

const(
	//Request types
	EnterUsername TypeID = iota 
	UsernameSuccess

	Password
	RegisterSuccess
	LoginSuccess

	ChangeUsername
	Success
	Failure

	NewUserJoined
	Quit 

	GetUsersList
	Connect
	// MsgToUser
	// FileToUser

	CreateGp
	AddMember
	// ConnToGp
	// MsgToGp
	// FileToGp 

	MsgTo
	FileTo
	FileACK

	ServerMsg
	ServerErr
)

// var CommandsList []string{"List of available command",
// 						 "/chusername",
// 						 "/getusers",
// 						 "/connuser",
// 						 "/creategp",
// 						 "/conngp",
// 						 "/addmember",
// 						 "/msg",
// 						 "/file",
// 						 "/quit"}