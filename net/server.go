package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//!+broadcaster
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

			// Broadcast incoming message to all
			// clients' outgoing message channels.
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

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	io.WriteString(conn, "Welcome to TCP-Chat!\n"+
		"\t_nnnn_\t\n"+
		"\tdGGGGMMb\t\n"+
		"\t@p~qp~~qMb\t\n"+
		"\tM|@||@) M|\t\n"+
		"\t@,----.JM|\t\n"+
		"\tJS^'__/  qKL\t\n"+
		"\r\n[ENTER YOUR NAME]:")
		

	ch := make(chan string) //Исходящие сообщения клиентов
	go clientWriter(conn, ch)

	user := ""
	t := time.Now()
	p := fmt.Println
	p(t.Format("2006-01-02 15:04:05"))

	// conn.Write([]byte("[" + t.Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": "))

	messages <- user + " has joined our chat..."
	entering <- ch

	input := bufio.NewReader(conn)
	for {
		letter, err := input.ReadString('\n')
		if err != nil {
			break
		}
		messages <- "[" + t.Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ": " + letter

	}

	messages <- user + " has left our chat"
	leaving <- ch
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

//!-handleConn

//!+main
func main() {

	arg := os.Args[1:]

	port := ":" + arg[0]
	// a := ":" + "8989"

	listener, err := net.Listen("tcp", port)
	if len(arg) == 0 {
		// arg[0] = "8989"
		log.Printf("Listening on the port" + port)
	} else if len(arg) == 1 {
		log.Printf("Listening on the port" + port)
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
