package model 

import(
	// "net"
)

type group struct{
	name string
	members map[net.Addr]*client //addr, clientObj
}