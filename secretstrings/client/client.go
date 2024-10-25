package main

import (
	//	"net/rpc"
	"bufio"
	"flag"
	"net/rpc"
	"os"

	//	"bufio"
	//	"os"
	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"

	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func makeCall(client *rpc.Client, message string) {
	request := stubs.Request{Message: message}
	response := new(stubs.Response)
	// client.Call(stubs.ReverseHandler, request, response)
	client.Call(stubs.PremiumReverseHandler, request, response)
	fmt.Println("Responded: " + response.Message)
}

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	// request := stubs.Request{Message: "Hello"}
	// response := new(stubs.Response)
	// // client.Call(stubs.ReverseHandler, request, response)
	// client.Call(stubs.PremiumReverseHandler, request, response)
	// fmt.Println("Responded: " + response.Message)

	file, _ := os.Open("wordlist")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Println("Called: " + t)
		makeCall(client, t)
	}

}
