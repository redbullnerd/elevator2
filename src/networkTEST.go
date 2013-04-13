package main

import(
	"network"
	"time"
	"fmt"
)

func main() {
	var communicator network.CommChannels
	communicator.CommChanInit()
	network.NetworkInit(communicator)

	time.Sleep(time.Second)
	
	go receiveTESTmail(communicator)

	for {
		sendTESTmail(communicator)
		time.Sleep(1000*time.Millisecond)
	}
}

func sendTESTmail(communicator network.CommChannels) {
	testvar := "WE'RE GOING TO FINISH THIS"
	randomstruct := network.DecodedMessage{"129.241.187.148", testvar}
	communicator.SendToOne <- randomstruct
}

func receiveTESTmail(communicator network.CommChannels) {
	for {
		received := <- communicator.DecodedMessagechan
		fmt.Println("received message: ", received.Content)
	}
}
