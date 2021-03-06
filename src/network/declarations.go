package network

import(
	"net"
	"time"
)

// ALL THIS CHANNEL USE MIGHT BE PRONE TO DEADLOCKS. GOOD IDEA TO IMPLEMENT TIMEOUTS?

const(
	sleepduration = 500 //interval between alivemessages given in milliseconds
	toleratedLosses = 4

	isAlive = 1
	isDead = 0
)

var(
	localIP = getIP()
	broadcast = "129.241.187.255" //må se nærmere på adressen

	UDPport = "6574" // randomly chosen ports
	TCPport = "6476"

)

var(
	internal internalchannels
)

type internalchannels struct {
	updateTCPmap chan TCPconnection // new TCP connections are shared over this channel
	newIPchan chan string // new IPs broadcasting UDP are shared here
	isDeadchan chan string // when UDP module detects that someone is dead, their IP is transmitted here
	isAlivechan chan string // for internal use in UDP module. When new ping is received, input to this channel resets death timer
	giveMeCurrentMap chan bool
	getCurrentMap chan map[string]net.Conn
	giveMeConn chan string
	getSingleConn chan net.Conn
	startNewReceivechan chan TCPconnection
	closeConn chan string
	quitsendImAlive chan bool
	quitlistenImAlive chan bool
	setupFail chan bool
	MessageReceivedchan chan encodedMessage
	encodedMessageSendAll chan encodedMessage
	encodedMessageSendOne chan encodedMessage
}

type TCPconnection struct { // inputs to map containing active TCP connections are of this type. IP is key, socket is content
	socket net.Conn
	IP string
}

type CommChannels struct { // collection of channels used for TCP communication
	SendToAll chan DecodedMessage
	SendToOne chan DecodedMessage
	DecodedMessagechan chan DecodedMessage
	getDeadIPchan chan string
	sendDeadIPchan chan string
	GiveMeCurrentAlives chan bool // can both struct have same names?
	GetCurrentAlives chan map[string]time.Time
}

type DecodedMessage struct { // struct to messageHandler for outgoing
	IP string
	Content string
}

type encodedMessage struct { // struct to messageHandler for incoming
	IP string
	Content []byte
}
