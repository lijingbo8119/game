package server

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Data struct {
	Cmd     DataCmd     `json:"cmd"`
	Params  interface{} `json:"params"`
}

func (r Data) Send(c *websocket.Conn) error {
	j, _ := json.Marshal(r)
	return c.WriteMessage(websocket.TextMessage, j)
}

func (r Data) ParseParams(ptr interface{}) error {
	j, _ := json.Marshal(r.Params)
	return json.Unmarshal(j, ptr)
}

func (r *Data) Parse(payload []byte) (*Data, error) {
	err := json.Unmarshal(payload, &r)
	return r, err
}
