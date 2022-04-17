package camunda_client_go

// Message a client for Message API
type Message struct {
	client *Client
}

// ReqMessage a request to send a message
type ReqMessage struct {
	MessageName      string               `json:"messageName"`
	BusinessKey      string               `json:"businessKey"`
	ProcessVariables *map[string]Variable `json:"processVariables,omitempty"`
}

// SendMessage sends message to a process
func (m *Message) SendMessage(query *ReqMessage) error {
	res, err := m.client.doPostJson("/message", map[string]string{}, query)
	if res != nil {
		res.Body.Close()
	}
	return err
}
