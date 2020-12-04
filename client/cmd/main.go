package main

import(
	"Gossenger/client/handler"
)
		
func main()  {
	client := handler.NewClient()
	client.Connect()
}	