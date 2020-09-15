package websocket

import (
	"github.com/zhouhui8915/engine.io-go/transport"
)

var Creater = transport.Creater{
	Name:      "websocket",
	Upgrading: true,
	Server:    NewServer,
	Client:    NewClient,
}
