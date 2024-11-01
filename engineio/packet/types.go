package packet

import "github.com/init-vlad/socketio/engineio/frame"

type Frame struct {
	FType frame.Type
	Data  []byte
}

type Packet struct {
	FType frame.Type
	PType Type
	Data  []byte
}
