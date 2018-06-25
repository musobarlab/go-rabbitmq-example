package sub

import (
	"encoding/json"
)

//Content struct
type Content struct {
	Header string `json:"header"`
	Body string `json:"body"`
}

//Message struct
type Message struct {
	From string `json:"from"`
	Content Content `json:"content"`
}

//JSON function
func (m Message) JSON() ([]byte, error) {
	return json.Marshal(m)
}