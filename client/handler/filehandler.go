package handler

import(
	"io/ioutil"
	"fmt"
)

func ReadFile(path string)([]byte,error){
	fmt.Printf("Reading a file %s\n",path) 
      
    data, err := ioutil.ReadFile(path) 
    if err != nil { 
		// log.Panicf("failed reading data from file: %s", err) 
		return nil, err
    } 
	return data, nil
}