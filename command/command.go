package command

import(
	"Gossenger/command/types"
)
//Request is the req struct
type Command struct{
	CmdType types.TypeID
	Data []byte
	From,To string
	Filename string
}

//NewReq Creates a new request
func NewCommand(cmdType types.TypeID, data []byte, from,to string)*Command{
	return &Command{
		CmdType: cmdType,
		Data: data,
		From: from,
		To: to,
		Filename: "",
	}
}