package main

import (
	"fmt"
	"net"
	"strings"
)

func getLocalAddrs() ([]net.IPNet, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var list []net.IPNet
	for _, addr := range addrs {
		v := addr.(*net.IPNet)
		if v.IP.To4() != nil {
			list = append(list, *v)
			// fmt.Printf("addr.(*net.IPNet): %v %v\n", v, v.Mask)
		}
	}
	return list, nil
}

func listLocalAddresses(netw, ip string) {
	proto := "tcp"

	// for stat, i := range result.Stats {
	// 	x := strings.Repeat(" ", 29-len(stat+strconv.Itoa(i)))
	// 	fmt.Println(`*      ` + stat + x + strconv.Itoa(i) + " %")
	// }

	if netw != "" {
		fmt.Println("\n**********************************************************")
		fmt.Println(" List of local Network available adresses: \n")
		list, _ := getLocalAddrs()
		for k, v := range list {
			space := strings.Repeat(" ", 25-len(string(k)+": "+v.String()))
			fmt.Printf("     %v: %v %v %v\n", k, v.String(), space, v.Network())
		}
	}
	if ip != "" {
		fmt.Println("\n**********************************************************")
		fmt.Println(" Remote Network details: \n")

		localAddress, err := net.ResolveTCPAddr(proto, ip)
		check(err)
		// for k, v := range list {
		fmt.Printf("      %v %v %v\n", localAddress, localAddress.Network(), localAddress.String)
		// }
	}
}

// func fwd(src net.Conn, remote string, proto string) {
// 	dst, err := net.Dial(proto, remote)
// 	errHandler(err)
// 	go func() {
// 		_, err = io.Copy(src, dst)
// 		errPrinter(err)
// 	}()
// 	go func() {
// 		_, err = io.Copy(dst, src)
// 		errPrinter(err)
// 	}()
// }

// func errHandler(err error) {
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[Error] %s\n", err.Error())
// 		os.Exit(1)
// 	}
// }

// // TODO: merge error handling functions
// func errPrinter(err error) {
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[Error] %s\n", err.Error())
// 	}
// }

// func tcpStart(from string, to string) {
// 	proto := "tcp"

// 	localAddress, err := net.ResolveTCPAddr(proto, from)
// 	errHandler(err)

// 	remoteAddress, err := net.ResolveTCPAddr(proto, to)
// 	errHandler(err)

// 	listener, err := net.ListenTCP(proto, localAddress)
// 	errHandler(err)

// 	defer listener.Close()

// 	fmt.Printf("Forwarding %s traffic from '%v' to '%v'\n", proto, localAddress, remoteAddress)
// 	fmt.Println("<CTRL+C> to exit")
// 	fmt.Println()

// 	for {
// 		src, err := listener.Accept()
// 		errHandler(err)
// 		fmt.Printf("New connection established from '%v'\n", src.RemoteAddr())
// 		go fwd(src, to, proto)
// 	}
// }

// func udpStart(from string, to string) {
// 	proto := "udp"

// 	localAddress, err := net.ResolveUDPAddr(proto, from)
// 	errHandler(err)

// 	remoteAddress, err := net.ResolveUDPAddr(proto, to)
// 	errHandler(err)

// 	listener, err := net.ListenUDP(proto, localAddress)
// 	errHandler(err)
// 	defer listener.Close()

// 	dst, err := net.DialUDP(proto, nil, remoteAddress)
// 	errHandler(err)
// 	defer dst.Close()

// 	fmt.Printf("Forwarding %s traffic from '%v' to '%v'\n", proto, localAddress, remoteAddress)
// 	fmt.Println("<CTRL+C> to exit")
// 	fmt.Println()

// 	buf := make([]byte, 512)
// 	for {
// 		rnum, err := listener.Read(buf[0:])
// 		errHandler(err)

// 		_, err = dst.Write(buf[:rnum])
// 		errHandler(err)

// 		fmt.Printf("%d bytes forwared\n", rnum)
// 	}
// }

// func ctrlc() {
// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
// 	go func() {
// 		sig := <-sigs
// 		fmt.Println("\nExecution stopped by", sig)
// 		os.Exit(0)
// 	}()
// }
