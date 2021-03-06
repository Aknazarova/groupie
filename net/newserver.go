package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

type client chan<- string // Канал исходящих сообщений

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // Все входящие сообщения клиента
)

func broadcaster() {

	clients := make(map[client]bool) // Все подключенные клиенты
	for {
		select {
		case msg := <-messages:
			// fmt.Println("received", msg)
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)

		}
	}

}

func handleConn(conn net.Conn) {
	img, err := ioutil.ReadFile("img.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ch := make(chan string) //Исходящие сообщения клиентов
	user := ""
	go clientWriter(conn, ch)

	io.WriteString(conn, string(img))
	io.WriteString(conn, "\r\n[ENTER YOUR NAME]:")

	client := []string{}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		name := scanner.Text()
		if name == "" {
			fmt.Fprintf(conn, "%s", "[ENTER YOUR NAME]: ")
			continue
		}

		user = name
		client = append(client, user)
		break
	}
	if user == "" {
		return
	} else {
		fmt.Fprintf(conn, "[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), user)
	}
	apuser := []string{}
	newuser := ""
	for scanner.Scan() {
		newuser = scanner.Text()
		if user == newuser {
			fmt.Fprintf(conn, "%s", "Данное имя уже используется")
		}

		apuser = append(apuser, newuser)
		break
	}

	enter := "\n" + user + " has joined our chat..." + "\n"

	select {
	case messages <- enter:
		// messages<-  "\n" + "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": "
		// messages <- enter
		messages <- "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": "

	}
	entering <- ch

	// output := ""

	for scanner.Scan() {
		letter := scanner.Text()

		// broadcastMessage := fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), user, letter)
		// nameHandler := fmt.Sprintf("[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), user)
		output := "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": " 

		for _, c := range client {
			for _, v := range apuser {
				if c != v {
					fmt.Fprintf(conn, output)
				}else{
					fmt.Fprintf(conn, "%s", "Данное имя уже используется")
				}
			}
			// fmt.Println(client)
			// if c == user {

			// 	fmt.Fprintf(conn, nameHandler)
			// 	// fmt.Fprintf(conn, broadcastMessage,nameHandler,letter)
			// } else {
			// 	// fmt.Fprintf(conn, nameHandler)
			// 	messages <-" выходи уже f"
			// return
			// messages <- new

		}

		// output = "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": " + letter

		// output = letter

		messages <- letter

		new := "\n" + "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": "
		messages <- new

		// messages <- str
		// output = "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": " + letter
	}

	messages <- "\n" + user + " has left our chat"
	messages <- "\n" + "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": "
	leaving <- ch
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

func main() {

	arg := os.Args[1:]

	port := ":" + arg[0]

	listener, err := net.Listen("tcp", port)
	if len(arg) == 0 {
		// arg[0] = "8989"
		fmt.Printf("Listening on the port " + port)
	} else if len(arg) == 1 {
		fmt.Printf("Listening on the port " + port)
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			break
		}
		go handleConn(conn)

	}
}
