package network

import(
	"net"
)

const(
	sleepduration = 1000 //interval between alivemessages given in milliseconds
	toleratedLosses = 4

	isAlive = 1
	dead = 0
)

var(
	localIP = findmyIP()
	broadcast = "235.241.187.255" //må se nærmere på adressen
	
	UDPport = "8769"
	TCPport = " 8770"

)

var(
	
	updateTCPmap chan TCPconnection
	IPshareChan chan string
	isDeadchan chan string
	isAlivechan chan int

	commChan chan string
)

type TCPconnection struct {
	socket net.Conn
	IP string
}
