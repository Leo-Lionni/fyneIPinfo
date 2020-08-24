package  main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
time.Sleep(time.Millisecond * 1)
conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
if err == nil {
_ = conn.Close()
return true
}

return false
}

func main() {
	hostname := flag.String("hostname","","host to test")
	startPort := flag.Int("start-port",80,"the port to scanning")
	endPort := flag.Int("end-port",100,"the port end ")
	timeout := flag.Duration("timeout",time.Millisecond*200,"timeout")
	flag.Parse()
mutex := &sync.Mutex{}

ports := []int{}

wg := &sync.WaitGroup{}
//timeout := time.Millisecond * 200
for port := *startPort; port < *endPort; port++ {
wg.Add(1)
go func(p int) {
	fmt.Println("i'm ",port,"! running")
opened := isOpen(*hostname, p, *timeout)
if opened {
	mutex.Lock()
ports = append(ports, p)
mutex.Unlock()
}
wg.Done()
	fmt.Println("i'm ",port,"! finished")
}(port)
}

wg.Wait()
fmt.Printf("opened ports: %v\n", ports)
}
