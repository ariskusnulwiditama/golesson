package connectnetwork

import (
	"fmt"
	"net"
	"testing"
)

func TestConnect(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Println(interf.Name)

		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPAddr); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}
}
