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
	received = make(chan string)
)

func broadcaster() {

	clients := make(map[client]bool) // Все подключенные клиенты
	for {
		select {
		case msg := <-messages:
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
	firstuser:=""
	if user == "" {
		return
		firstuser = user
	} else {
		fmt.Fprintf(conn, "[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), user)
	
	}


		enter := "\n" + user + " has joined our chat..." + "\n"
		if user == firstuser{
		messages<-  "\n" + "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + firstuser + "]" + ": " 
		messages <- enter
		// messages<-  "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": "
		entering <- ch
		}


	for scanner.Scan() {
		letter := scanner.Text()
		// output := ""
		broadcastMessage := fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), user, letter)
		nameHandler := fmt.Sprintf("[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), user)
		for _, c := range client {
			if c != user {

				fmt.Fprintf(conn, broadcastMessage,user)
			} else {
				fmt.Fprintf(conn, nameHandler)
				// return
			}
			
			// output = "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": " + letter
		}

	messages <- letter 

		new := "\n" + "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": "
		messages <- new
		// messages <- str
		// output = "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": " + letter
	}

	messages <- "\n" + user + " has left our chat"
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
