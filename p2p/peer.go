package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	conn *websocket.Conn
}

func (p *peer) read() {
	for {
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			// delete peer
			break
		}
		fmt.Printf("%s", m)
	}
}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		conn,
	}
	go p.read()
	Peers[key] = p
	return p
}
