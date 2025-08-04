package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// definindo a porta e a interface do servidor
	addr := net.TCPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
		/*
			"0.0.0.0" -> permite que o servidor opere em todas as interfaces disponiveis:
			internet, ethernet, localhost...
		*/
	}

	listener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Println("Erro ao escutar:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor ouvindo em", addr)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Nao foi possivel fazer conexão:", err)
			continue
		}

		// ao invés de seguirmos com o corpo da conexão aqui, vamos paralelizar a continuidade em uma nova
		// função em uma outra gorotine e liberar o "while" para abrir uma nova conexão
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()

	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer) // buffer == &buffer -> true
	if err != nil {
		fmt.Println("Erro ao ler:", err)
		return
	}

	// buffer([]byte) -> string -> stringUpper -> response([]byte)
	bufferUpperCase := strings.ToUpper(string(buffer[:n]))
	response := []byte(bufferUpperCase)
	conn.Write(response)
}
