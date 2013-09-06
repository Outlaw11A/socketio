package socketio

import (
	"errors"
	"strconv"
	"strings"
)

type IOMessage struct {
	Type     int
	Id       int
	Endpoint *Endpoint
	Data     string
}

func NewIOMessage(rawMsg string) (*IOMessage, error) {
	if len(rawMsg) == 0 {
		return nil, errors.New("Empty message")
	}

	msgType, err := strconv.Atoi(string(rawMsg[0]))
	if err != nil {
		return nil, err
	}

	switch msgType {
	case 3, 4, 5:
		parts := strings.SplitN(rawMsg, ":", 4)

		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		return &IOMessage{msgType, id, nil, parts[3]}, nil
	default:
		return &IOMessage{Type: msgType}, nil
	}
}

func (m IOMessage) String() string {
	raw := strconv.Itoa(m.Type)

	raw += ":"
	if m.Id != 0 {
		raw += strconv.Itoa(m.Id)
	}

	raw += ":"
	if m.Endpoint != nil {
		raw += m.Endpoint.String()
	}

	if m.Data != "" {
		raw += ":" + m.Data
	}

	return raw
}

func NewDisconnect() *IOMessage {
	return &IOMessage{Type: 0}
}

func NewConnect(endpoint *Endpoint) *IOMessage {
	return &IOMessage{Type: 1, Endpoint: endpoint}
}

func NewHeartbeat() *IOMessage {
	return &IOMessage{Type: 2}
}

func NewMessage(endpoint *Endpoint, data string) *IOMessage {
	return &IOMessage{Type: 3, Endpoint: endpoint, Data: data}
}

func NewJSONMessage(endpoint *Endpoint, data string) *IOMessage {
	return &IOMessage{Type: 4, Endpoint: endpoint, Data: data}
}

func NewEvent(endpoint *Endpoint, data string) *IOMessage {
	return &IOMessage{Type: 5, Endpoint: endpoint, Data: data}
}

func NewACK(endpoint *Endpoint, data string) *IOMessage {
	return &IOMessage{Type: 6, Endpoint: endpoint, Data: data}
}

func NewError(endpoint *Endpoint, data string) *IOMessage {
	return &IOMessage{Type: 7, Endpoint: endpoint, Data: data}
}

func NewNoop() *IOMessage {
	return &IOMessage{Type: 8}
}
