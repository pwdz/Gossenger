package main

type commandID int

const(
	cmdEnterUsername commandID = iota //0
	cmdChangeUsername
	cmdJoin
	cmdGet
	cmdConnToUser
	cmdConnToGp
	cmdMsgToUser
	cmdFileToUser
	cmdMsgToGp
	cmdFileToGp 
	cmdQuit 
)

type command struct{
	id commandID
	client *client
	args []string
}