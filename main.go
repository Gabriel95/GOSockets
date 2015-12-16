package main
import "fmt"
import	"strconv"
import	"bufio"
import	"os"
import	"strings"
import	"net"

func main(){
	/*reader := bufio.NewReader(os.Stdin)
   fmt.Print("Enter text: ")
   text, _ := reader.ReadString('\n')
   fmt.Println(text)
   /*
   for i:= 0; i < 5; i++{
	   fmt.Println("Hellow World!")
   }
   fmt.Println("Hellow World!")
   */
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for{

		fmt.Print("1. Add User\n")
		fmt.Print("2. Show User in Console\n")
		fmt.Print("3. Delete User\n")
		fmt.Print("4. Send User to Email\n")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		option,_ :=  strconv.Atoi(strings.TrimSpace(text))

		switch option {
		case 1:
			fmt.Print("Name: ")
			name,_ := reader.ReadString('\n')
			fmt.Print("Username: ")
			username,_ := reader.ReadString('\n')
			fmt.Print("Email: ")
			email,_ := reader.ReadString('\n')
			fmt.Print("Id: ")
			id,_ := reader.ReadString('\n')
			fmt.Print("Birthday: ")
			birthday,_ := reader.ReadString('\n')
			fmt.Print("Foto: ")
			foto,_ := reader.ReadString('\n')

			name = strings.TrimSpace(name)
			username =  strings.TrimSpace(username)
			email =  strings.TrimSpace(email)
			id = strings.TrimSpace(id)
			birthday =  strings.TrimSpace(birthday)
			foto =  strings.TrimSpace(foto)
			client := "1 " + name + ", " + username + ", " + email + ", " + id + ", " + birthday + ", " + foto
			fmt.Fprintf(conn, client + "\n")
			break
		case 2:
			fmt.Print("Username: ")
			username,_ := reader.ReadString('\n')
			username =  strings.TrimSpace(username)
			data := "2 " + username
			fmt.Fprintf(conn, data + "\n")
			response, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print(response)
			break
		case 3:
			fmt.Print("Username: ")
			username,_ := reader.ReadString('\n')
			username =  strings.TrimSpace(username)
			data := "3 " + username
			fmt.Fprintf(conn, data + "\n")
			response, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print(response)
			break
		case 4:
			fmt.Print("d\n")
			break
		}
	}
}