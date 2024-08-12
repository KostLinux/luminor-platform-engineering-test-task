package model

type Payload struct {
	Ts         string  `json:"ts"`
	Sender     string  `json:"sender"`
	Message    Message `json:"message"`
	SentFromIP string  `json:"sent_from_ip"`
}

type Message struct {
	Input string `json:"input"`
	Hash  string `json:"hash"`
}
