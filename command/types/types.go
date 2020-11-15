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

	NewUserJoined
	Quit 

	GetUsersList
	ConnToUser
	// MsgToUser
	// FileToUser

	CreateGp
	AddMember
	ConnToGp
	// MsgToGp
	// FileToGp 

	Msg
	File
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