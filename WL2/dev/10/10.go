package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	go startServer()
	// даем время серверу запуститься
	time.Sleep(1 * time.Second)

	timeOut := flag.Int("timeout", 10, "Time out flag")
	flag.Parse()
	// коннектимся к серверу с некоторым таймаутом
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8081", time.Duration(*timeOut)*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	// запускаем горутину, которая будет считывать данные из стдин и отрпавлять на сервер.
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				conn.Close()
			}
			fmt.Fprint(conn, text+"\n")
		}
	}()

	for {
		mes, err := bufio.NewReader(conn).ReadString('\n')
		// если соединение закрыто, то выйдет ошибка
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Client: " + mes)
	}

}

func startServer() {
	ln, _ := net.Listen("tcp", "127.0.0.1:8081")

	conn, _ := ln.Accept()

	for {
		// читает сообщеия и возвращает новые
		mes, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Mes server: ", mes)

		mes = "new " + mes

		conn.Write([]byte(mes + "\n"))
	}
}
