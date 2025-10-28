package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Getting local IP addresses...")

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip := v.IP
				if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
					continue
				}
				if ip.To4() != nil {
					fmt.Printf("Interface: %-10s | IPv4: %s\n", i.Name, ip.String())
				}
			}
		}
	}
}
