package main

import "C"
import (
	identify "github.com/libp2p/go-libp2p/p2p/protocol/identify"

	p2pc "github.com/libp2p/go-libp2p-daemon/p2pclient"
)

func main() {
	identify.ClientVersion = "p2pc/0.1"
	config := p2pc.Initialize()
	p2pc.Start(config)
}

//export startClient
func startClient(args *C.char) {
	argsGoString := C.GoString(args)
	config := p2pc.ProcessArgs(&argsGoString)
	p2pc.Start(config)
}
