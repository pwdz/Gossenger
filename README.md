# Gossenger
Go-Messenger [Client Server App] 🚶‍♂️🚶‍♂️🚶‍♂️🚶‍♂️
 First run server
 ```
 go run server/cmd/main.go
 ```
 Then run clients
 ```
 go run client/cmd/main.go
 ```
Encoding ➡️ Base64  
Database ➡️ SQLite    

Currently Supported Commands:    
- **/username [username]**
- **/password [password]**
- **/changeusername [newUsername]**
- **/connect [UserID/GroupID]**
- **/send [message]**
- **/file [path]**
- **/creategp [gpName]** (Default Admin: Group creator) 
- **/addmembers [UserID1] [UserID2] ...** (Permission: Group Member)
- **/removemembers [UserID1] [UserID2] ...** (Permission: Group Admin)  

After registeration, username and password will be saved in database
  
TODO:  
- Save messages
- Save files
- save groups
- Add GUI
