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
	AddMembers
	RemoveMembers

	MsgTo
	FileTo
	FileACK

	ServerMsg
	ServerErr
)