package main

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
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
	// proto := "tcp"

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
		fmt.Println(" Remote Network details for " + ip + ": \n")
		ps := &PortScanner{
			ip:   ip,
			lock: semaphore.NewWeighted(1024),
		}
		ps.Start(1, 65535, 500*time.Millisecond)
	}
}

// *************************************************************************

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			fmt.Println(port, "closed")
		}
		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(port)
	}
}

// func main() {
// 	ps := &PortScanner{
// 		ip:   "127.0.0.1",
// 		lock: semaphore.NewWeighted(Ulimit()),
// 	}
// 	ps.Start(1, 65535, 500*time.Millisecond)
// }
