package server

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Data struct {
	NetCmd  NetCmd `json:"cmd"`
	Payload any    `json:"payload"`
}

func (r Data) Send(c *websocket.Conn) error {
	j, _ := json.Marshal(r)
	return c.WriteMessage(websocket.TextMessage, j)
}

func (r Data) ParsePayload(ptr any) error {
	j, _ := json.Marshal(r.Payload)
	return json.Unmarshal(j, ptr)
}

func (r *Data) Parse(payload []byte) (*Data, error) {
	err := json.Unmarshal(payload, &r)
	return r, err
}
