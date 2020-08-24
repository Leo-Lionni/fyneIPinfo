///
package main
import (
	"errors"
	"fmt"
	"net"
	"strings"
)
func main(){
	addr ,_:=	GetLocalAddr()
	fmt.Println(addr,"ok!")
	var s ,_ = LocalIPv4s()
	fmt.Println(s)
}
func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}



func GetLocalAddr() (net.IP, error) {
	fmt.Println(`config.HostName is empty(""), choose one IP for listening automatically.`)
	//接口们
	infs, err := net.Interfaces()
	if err != nil {
		panic("fail to get net interfaces")
	}

	for _, inf := range infs {
		if inf.Flags&4 != 4 && !strings.Contains(inf.Name, "docker") {
			//拿到每一个接口的地址们
			addrs, err := inf.Addrs()
			if err != nil {
				panic("fail to get addrs of interface")
			}
			//拿到地址；
			for _, addr := range addrs {
				//反射，判定addr类型
				switch v := addr.(type) {
				case *net.IPAddr:
					if !strings.Contains(v.IP.String(), ":") {
						return v.IP, nil
					}
				case *net.IPNet:
					if !strings.Contains(v.IP.String(), ":") {
						return v.IP, nil
					}
				}
			}
		}
	}

	return nil, errors.New("no addr found")
}

