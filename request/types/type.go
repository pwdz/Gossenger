package types

//TypeID is for request types
type TypeID int

const(
	//Request types
	EnterUsername TypeID = iota 
	Password
	ChangeUsername
	GetUsersList
	ConnToUser
	ConnToGp
	MsgToUser
	FileToUser
	MsgToGp
	FileToGp 
	Quit 
)