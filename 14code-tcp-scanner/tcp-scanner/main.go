package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(j int){
			defer wg.Done()
			addrPort := fmt.Sprintf("192.168.0.106:%d", j)
			
			conn, err := net.Dial("tcp", addrPort)
	
			if err != nil {
				fmt.Printf("%s closed\n", addrPort)
				return
			}
			
			conn.Close()
			fmt.Printf("%s opened\n", addrPort)
		}(i)
	}
	wg.Wait()
}