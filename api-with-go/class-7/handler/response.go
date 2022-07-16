package handler

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(msgType, msg string, data interface{}) response {
	return response{
		MessageType: msgType,
		Message:     msg,
		Data:        data,
	}
}
