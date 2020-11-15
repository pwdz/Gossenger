package request

import(
	"Gossenger/request/types"
)
//Request is the req struct
type Request struct{
	ReqType types.TypeID
	Data []byte
	From,To string
}

//NewReq Creates a new request
func NewReq(reqType types.TypeID, data []byte, from,to string)*Request{
	return &Request{
		ReqType: reqType,
		Data: data,
		From: from,
		To: to,
	}
}