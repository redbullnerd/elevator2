package network

// In this part, remote elevators are pinged through UDP
// pings are also received, so that we may keep track of who is alive and who isn't

import(
	"fmt"
	"net"
	"time"
	"os"
)


func UDPHandler(communicator CommChannels) { //goroutine that keeps track of who is alive and who isn't
	aliveMap := make(map[string]time.Time)
	for{
		select{
		case ip := <- internal.isAlivechan:
			_, exists := aliveMap[ip]
			if exists {
				aliveMap[ip] = time.Now()
			} else	{			
			aliveMap[ip] = time.Now()
			internal.newIPchan <- ip
			}
		case <- time.After(30*time.Millisecond):
			for ip, lasttime := range aliveMap {
				if time.Now().After(lasttime.Add(toleratedLosses * sleepduration * time.Millisecond)) {
					fmt.Println("someone missed UDP deadline, and is terminated from aliveMap")
					internal.isDeadchan <- ip
					delete(aliveMap, ip)
				}
			}
		// elevator pack might want alives to count who will give cost
		case <- communicator.GiveMeCurrentAlives:
			communicator.GetCurrentAlives <- aliveMap
		}
	}
}

func sendImAlive() {
	service := broadcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println("net.ResolveUDPAddr error in sendImAlive: ", err)
		internal.setupFail <- true
	}

	isalivesocket, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println("net.DialUDP error in sendImAlive: ", err)
		internal.setupFail <- true
	}
	isAliveMessage := []byte("ping")
	
	for {
		select {
		case <- internal.quitsendImAlive:
			return
		default:
			_, err := isalivesocket.Write(isAliveMessage)
			if err != nil {
				fmt.Println("Write error in sendImAlive: ", err)
			}
			time.Sleep(sleepduration * time.Millisecond)
		}
	}
}

func listenImAlive() {
	service := broadcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println("net.ResolveUDPAddr error in listenImAlive: ", err)
		internal.setupFail <- true
	}

	isalivesocket, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("net.ListenUDP error in listenImAlive: ", err)
		internal.setupFail <- true
	}
	var data [512]byte
	
	for {
		select {
		case <- internal.quitlistenImAlive:
			return
		default:
			_, senderAddr, err := isalivesocket.ReadFromUDP(data[0:])
			if err != nil {
				fmt.Println("ReadFromUDP error in listenImAlive: ", err)
			}
			if localIP != senderAddr.IP.String(){
				if err != nil {
					fmt.Println("read error in listenImAlive")
				} else {
					remoteElev := senderAddr.IP.String()
					internal.isAlivechan <- remoteElev
				}
			}
		}
	}
}

func errorhandler(err error){ // tidies up code. will be replaced by individualized error handling for each error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
