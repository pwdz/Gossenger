package utils

import(
	"fmt"	
	"Gossenger/request"
	"encoding/base64"
    "encoding/gob"
    "bytes"
)
//ToBase64 convert request to bytes
func ToBase64(req request.Request) []byte {
    buff := bytes.Buffer{}
    encoder := gob.NewEncoder(&buff)
    err := encoder.Encode(req)
    if err != nil { 
		fmt.Println(`failed gob Encode`, err) 
	}
    return []byte(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

// FromBase64 to request struct
func FromBase64(data []byte) request.Request {
	req := request.Request{}
	// var decodedBytes
    decodedBytes, err := base64.StdEncoding.DecodeString(string(data))
    if err != nil { 
		fmt.Println(`failed base64 Decode`, err); 
	}
    buff := bytes.Buffer{}
    buff.Write(decodedBytes)
    decoder := gob.NewDecoder(&buff)
    err = decoder.Decode(&req)
    if err != nil { 
		fmt.Println(`failed gob Decode`, err); 
	}
    return req
}