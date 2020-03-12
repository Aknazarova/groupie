// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {

// 	// Number of people whom ever connected
// 	//
// 	clientCount := 0

// 	// All people who are connected; a map wherein
// 	// the keys are net.Conn objects and the values
// 	// are client "ids", an integer.
// 	//
// 	allClients := make(map[net.Conn]int)

// 	// Channel into which the TCP server will push
// 	// new connections.
// 	//
// 	newConnections := make(chan net.Conn)

// 	// Channel into which we'll push dead connections
// 	// for removal from allClients.
// 	//
// 	deadConnections := make(chan net.Conn)

// 	// Channel into which we'll push messages from
// 	// connected clients so that we can broadcast them
// 	// to every connection in allClients.
// 	//
// 	messages := make(chan string)

// 	// Start the TCP server
// 	//
// 	server, err := net.Listen("tcp", ":6000")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// Tell the server to accept connections forever
// 	// and push new connections into the newConnections channel.
// 	//
// 	go func() {
// 		for {
// 			conn, err := server.Accept()
// 			if err != nil {
// 				fmt.Println(err)
// 				os.Exit(1)
// 			}
// 			newConnections <- conn
// 		}
// 	}()

// 	// Loop endlessly
// 	//
// 	for {

// 		// Handle 1) new connections; 2) dead connections;
// 		// and, 3) broadcast messages.
// 		//
// 		select {

// 		// Accept new clients
// 		//
// 		case conn := <-newConnections:

// 			log.Printf("Accepted new client, #%d", clientCount)

// 			// Add this connection to the `allClients` map
// 			//
// 			allClients[conn] = clientCount
// 			clientCount += 1

// 			// Constantly read incoming messages from this
// 			// client in a goroutine and push those onto
// 			// the messages channel for broadcast to others.
// 			//
// 			go func(conn net.Conn, clientId int) {
// 				reader := bufio.NewReader(conn)
// 				for {
// 					incoming, err := reader.ReadString('\n')
// 					if err != nil {
// 						break
// 					}
// 					messages <- fmt.Sprintf("Client %d > %s", clientId, incoming)
// 				}

// 				// When we encouter `err` reading, send this
// 				// connection to `deadConnections` for removal.
// 				//
// 				deadConnections <- conn

// 			}(conn, allClients[conn])

// 		// Accept messages from connected clients
// 		//
// 		case message := <-messages:

// 			// Loop over all connected clients
// 			//
// 			for conn, _ := range allClients {

// 				// Send them a message in a go-routine
// 				// so that the network operation doesn't block
// 				//
// 				go func(conn net.Conn, message string) {
// 					_, err := conn.Write([]byte(message))

// 					// If there was an error communicating
// 					// with them, the connection is dead.
// 					if err != nil {
// 						deadConnections <- conn
// 					}
// 				}(conn, message)
// 			}
// 			log.Printf("New message: %s", message)
// 			log.Printf("Broadcast to %d clients", len(allClients))

// 		// Remove dead clients
// 		//
// 		case conn := <-deadConnections:
// 			log.Printf("Client %d disconnected", allClients[conn])
// 			delete(allClients, conn)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2) // в группе две горутины
// 	work := func(id int) {
// 		defer wg.Done()
// 		fmt.Printf("Горутина %d начала выполнение \n", id)
// 		time.Sleep(2 * time.Second)
// 		fmt.Printf("Горутина %d завершила выполнение \n", id)
// 	}

// 	// вызываем горутины
// 	go work(1)
// 	go work(2)

// 	wg.Wait() // ожидаем завершения обоих горутин
// 	fmt.Println("Горутины завершили выполнение")
// }

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	// read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		// logic
		if len(fs) < 1 {
			continue
		}
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}
	}
}
