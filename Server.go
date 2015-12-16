package main
import "net"
import "fmt"
import "strings"
//import "github.com/go-mailgun"
import (
	"bufio"
	"strconv"
	"os"
	"io/ioutil"
	"github.com/mailgun/mailgun-go"
)
//import "strings"

func main(){
fmt.Println("Launching server...")
// listen on all interfaces

// run loop forever (or until ctrl-c)
for {
	ln, _ := net.Listen("tcp", ":8081")
	// accept connection on port
	conn, _ := ln.Accept()
	// will listen for message to process ending in newline (\n)
  message, _ := bufio.NewReader(conn).ReadString('\n')
	option,_ := strconv.Atoi(message[0:1])
	data := message[2:len(message)-1]
	switch option{
		case 1:
			writeToFile(data)
			break
		case 2:
			conn.Write([]byte(searchFromFile(data) + "\n"))
			break
		case 3:
			conn.Write([]byte(deleteFromFile(data) + "\n"))
			break
		case 4:

			break
	}
// output message received
 //fmt.Print("Message Received: ", string(message))
// sample process for string received
  //  newmessage := strings.ToUpper(message)
// send new string back to client
 //conn.Write([]byte(newmessage + "\n"))
 }
}


func check(e error) {
	if e != nil {
		fmt.Println(error.Error)
	}
}

func createFile(){
	if _, err := os.Stat("data.txt"); os.IsNotExist(err) {
		file, er := os.Create("data.txt")
		check(er)
		defer file.Close()
	}
}

func writeToFile(data string){
	createFile()
	f,error := os.OpenFile("data.txt", os.O_APPEND, 0666)
	check(error)
	if len(data) > 0{
		f.WriteString(data + "/*splithere*/");
	}
	f.Close()
}

func searchFromFile(username string) string{
	createFile()
	dat,err := ioutil.ReadFile("data.txt")
	check(err)
	dats := string(dat)
	elements := strings.Split(dats,"/*splithere*/")
	for _,element := range elements {
		if strings.Contains(element, username){
			return element
		}
	}

	return "User Not Found"
}

func deleteFromFile(username string) string{
	userToDelete := searchFromFile(username)
	if strings.Compare(userToDelete,"User Not Found") == 0 {
		return "User Not Found"
	}
	dat,err := ioutil.ReadFile("data.txt")
	check(err)
	dats := string(dat)
	elements := strings.Split(dats,"/*splithere*/")
	error := os.Remove("data.txt")
	check(error)
	for _,element := range elements {
		if strings.Compare(element, userToDelete) != 0{
			writeToFile(element)
		}
	}
	return "User Deleted!"
}

/*func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	m := mg.NewMessage(
		"Excited User <mailgun@YOUR_DOMAIN_NAME>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"YOU@YOUR_DOMAIN_NAME",
	)
	_, id, err := mg.Send(m)
	return id, err
}*/
