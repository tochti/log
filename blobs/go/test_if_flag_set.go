package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, iface := range l {
		fmt.Printf("--- %v ---\n", iface.Name)
		// Hier wird mittels dem UND Operator getestet ob
		// die Broadcast Flag für das Netzwerk-Interface gesetzt ist
		// siehe dazu auch bit_operation_use_for_flags.go
		// und
		// https://golang.org/src/net/interface.go?s=1590:1620#L36
		if (iface.Flags & net.FlagBroadcast) == 0 {
			fmt.Println("No broadcast")
			continue
		}

		fmt.Println("\tFlags", iface.Flags)
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println(err)
			continue
		}

		for _, addr := range addrs {
			fmt.Printf("\t%v\n", addr.String())
		}

	}
}
