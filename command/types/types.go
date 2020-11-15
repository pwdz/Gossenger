package types

//TypeID is for request types
type TypeID int

const(
	//Request types
	EnterUsername TypeID = iota 
	Password
	ChangeUsername

	NewUserJoined
	Quit 

	GetUsersList
	ConnToUser
	MsgToUser
	FileToUser

	CreateGp
	AddMember
	ConnToGp
	MsgToGp
	FileToGp 
)