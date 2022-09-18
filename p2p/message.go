package p2p

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	// Message 종류 유형
	Kind    MessageKind
	Payload []byte
}
