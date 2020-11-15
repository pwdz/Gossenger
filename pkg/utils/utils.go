package utils

import(
	"fmt"	
	"Gossenger/command"
	"encoding/base64"
    "encoding/gob"
    "bytes"
)
//ToBase64 convert structs to bytes
func ToBase64(in interface{}) []byte {
    buff := bytes.Buffer{}
    encoder := gob.NewEncoder(&buff)
    err := encoder.Encode(in)
    if err != nil { 
		fmt.Println(`failed gob Encode`, err) 
	}
    return []byte(base64.StdEncoding.EncodeToString(buff.Bytes()))
}

// FromBase64 to request struct
func FromBase64(data []byte) command.Command {
	req := command.Command{}
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