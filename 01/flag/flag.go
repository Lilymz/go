package main

import (
	"flag"
	"fmt"
)

func main() {
	var ip string
	var port int
	var timeout int
	flag.StringVar(&ip, "ip", "127.0.0.1", "this is ip")
	flag.IntVar(&port, "port", 8080, "this is port")
	flag.IntVar(&timeout, "timeout", 3000, "this is timeout")
	flag.Parse()
	nFlag := flag.NFlag()
	nArg := flag.NArg()
	fmt.Println(nFlag)
	fmt.Println(nArg)
	fmt.Println("ip 有值", ip)
	flag.Set("timeout", "5000")
	fmt.Println("port 有值", port)
	fmt.Println("timeout 有值", timeout)
	flag.PrintDefaults()
}
