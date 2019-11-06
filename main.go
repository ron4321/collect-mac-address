package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	addresses, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}

	display(hostName, addresses)

	waitExit()
}

func getMacAddr() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var macAddresses []string
	for _, i := range interfaces {
		a := i.HardwareAddr.String()
		if a != "" {
			macAddresses = append(macAddresses, a)
		}
	}

	return macAddresses, nil
}

func display(hostName string, macAddresses []string) {
	fmt.Println("------------------------------")
	fmt.Println("- Hostname")
	fmt.Println(hostName)
	fmt.Println("")
	fmt.Println("- MAC Addresses")
	for _, a := range macAddresses {
		fmt.Println(a)
	}
	fmt.Println("------------------------------")
}

func waitExit() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter key to exit.")
	reader.ReadString('\n')
}
