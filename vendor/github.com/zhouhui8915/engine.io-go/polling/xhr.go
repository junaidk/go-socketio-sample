package polling

import (
	"github.com/zhouhui8915/engine.io-go/transport"
)

var Creater = transport.Creater{
	Name:      "polling",
	Upgrading: false,
	Server:    NewServer,
	Client:    NewClient,
}
