package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// conectando imediatamente o servidor localhost:8080
	conn, err := net.Dial(
		"tcp",
		"127.0.0.1:8080", // nesse caso o IP e a porta são declarados juntos
	)
	if err != nil {
		fmt.Println("Erro no handshake ao tentar se conectar com o servidor:", err)
		return
	}
	defer conn.Close()

	// escutador para a string de entrada
	reader := bufio.NewReader(os.Stdin)
	mensageByte, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Erro ao tentar ler a mensagem:", err)
		return
	}

	// enviando mensagem ao servidor via TCP
	_, err = conn.Write(mensageByte)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem ao servidor:", err)
		return
	}

	// esperando mensagem resposta do servidor
	buffer := make([]byte, 2048)
	_, err = conn.Read(buffer) // o valor primeiro omitido é o tamanho (n)
	if err != nil {
		fmt.Println("Erro ao ler resposta do servidor:", err)
		return
	}

	fmt.Printf("A resposta foi %s", string(buffer))

}
