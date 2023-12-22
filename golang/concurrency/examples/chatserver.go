package examples

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 将 client 抽象成一个只写channel
type client chan<- string

// Channels
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func cancelling() {
	// [todo] add cancelling
}

func broadcaster() {
	clients := make(map[client]bool)

	// 所有对clients的rw操作集中在broadcaster-routine
	for {
		select {
		case newCli := <-entering:
			clients[newCli] = true
		case deleteCli := <-leaving:
			delete(clients, deleteCli)
			close(deleteCli)
		case msg := <-messages:
			for cli := range clients {
				// msg 发送到clientCh (另一端为writeToClient)
				cli <- msg
			}
		}
	}
}

func writeToClient(conn *net.Conn, clientCh <-chan string) {
	// clientCh channel 上的消费者
	for msg := range clientCh {
		fmt.Fprintln(*conn, msg) // NOTE: ignoring network errors
	}
}

func connHandler(conn *net.Conn) {
	who := (*conn).RemoteAddr().String()

	clientCh := make(chan string)
	// 注册/启动 clientCh消费者routine，每个client都会拥有自己的 w-goroutine
	go writeToClient(conn, clientCh)
	entering <- clientCh
	clientCh <- "You are" + who
	messages <- who + " has arrived"

	input := bufio.NewScanner(*conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	// 连接关闭，退出了 input.Scan 循环
	leaving <- clientCh
	messages <- who + " has leaved"
	(*conn).Close()
}

func ChatServerMain() {
	// 起后台goroutinue，用于在收到消息时发送给所有聊天室中的人
	go broadcaster()

	// Listen & dispatch
	fmt.Println("Chat Server start!")
	listener, listenCreateErr := net.Listen("tcp", "localhost:8000")
	if listenCreateErr != nil {
		log.Fatal(listenCreateErr)
	}
	for {
		conn, listenErr := listener.Accept()
		if listenErr != nil {
			log.Fatal(listenErr)
			continue
		}
		go connHandler(&conn)
	}
}
