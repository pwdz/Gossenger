# Gossenger
Go-Messenger [Client Server App] (🚶‍♂️🚶‍♂️🚶‍♂️🚶‍♂️ **WIP**)  
 First run server.go
 ```
 go run server/cmd/main.go
 ```
 Then run clients
 ```
 go run client/cmd/main.go
 ```
Encoding > Base64  
Database > SQLite    

Currently Supported Commands:    
- **/username [username]**
- **/password [password]**
- **/changeusername [newUsername]**
- **/connect [UserID/GroupID]**
- **/send [message]**
- **/file [path]**
- **/creategp [gpName]** (Default Admin: Group creator) 
- **/addmembers [UserID1] [UserID2] ...**
- **/removemembers [UserID1] [UserID2] ...** (Permission: Group Admin)
  
TODO:  
- Add GUI
- Save messages in db
- Save files
- save groups
