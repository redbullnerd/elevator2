package network

import(
	"net"

	"fmt"
)

func getIP() string{
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("ERROR FINDING IP", err)
	}
	
	for _, a := range addr {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			fmt.Println(ipnet.IP.String())
			return ipnet.IP.String()
		}
	}
	return "failure"
}

/*
func getMyIP() string{
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("error occured in IP retrieval")
	}
	
	addr, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("error occured in IP retrieval")
	}
	
	for _, a := range addr {
		fmt.Println("ALL IPs: ", a)
	}
	
	return addr[0]
}

func findmyIP() string{ // this function is weird, and should be looked at. returns ip6 -.- working on better option
	systemIPs, err := net.InterfaceAddrs()
	errorhandler(err)

	tempIPstring := make([]string, len(systemIPs))
	
	for i := range systemIPs{
		temp := systemIPs[i].String()
		ip := strings.Split(temp, "/")
		tempIPstring[i] = ip[0]
	}
	myIP := tempIPstring[2]
	return myIP
}
*/
