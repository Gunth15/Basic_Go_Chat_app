package main

import "fmt"

/*
Need these routes:
- landing page/ GET
- signup page/ GET, POST
- login page/ GET, POST
- update login page/ GET, PUT
- logout button / Remove cookie
- profile page/ GET(should be able to see current chat rooms)
- page so users can find each other(Look to create new rooms)
- rooms(places where users commmunicate to each other using web sockets; when users online, open ws for sending. If the other user is online, open ws for recieving.)

Involve cookies
*/

/*
Tables
- User: id, username, password
- Chats: comment, uid_1, uid_2, time
*/

func main() {
	fmt.Println("Start pproject lol")
}
