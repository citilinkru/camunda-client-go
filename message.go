package camunda_client_go

// Message a client for Message API
type Message struct {
	client *Client
}

// ReqMessage a request to send a message
type ReqMessage struct {
	// The name of the message to deliver.
	MessageName string `json:"messageName,omitempty"`
	// Used for correlation of process instances that wait for incoming messages.
	// Will only correlate to executions that belong to a process instance with the provided business key.
	BusinessKey string `json:"businessKey,omitempty"`
	// Used to correlate the message for a tenant with the given id.
	// Will only correlate to executions and process definitions which belong to the tenant.
	// Must not be supplied in conjunction with a withoutTenantId.
	TenantId string `json:"tenantIdIn,omitempty"`
	// A Boolean value that indicates whether the message should only be correlated to executions and process definitions which belong to no tenant or not.
	// Value may only be true, as false is the default behavior.
	// Must not be supplied in conjunction with a tenantId.
	WithoutTenantId bool `json:"withoutTenantId,omitempty"`
	// Used to correlate the message to the process instance with the given id.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`
	// Used for correlation of process instances that wait for incoming messages.
	// Has to be a JSON object containing key-value pairs that are matched against process instance variables during correlation.
	// Each key is a variable name and each value a JSON variable value object with the following properties.
	CorrelationKeys map[string]Variable `json:"correlationKeys,omitempty"`
	// Local variables used for correlation of executions (process instances) that wait for incoming messages.
	// Has to be a JSON object containing key-value pairs that are matched against local variables during correlation.
	// Each key is a variable name and each value a JSON variable value object with the following properties.
	LocalCorrelationKeys map[string]Variable `json:"localCorrelationKeys,omitempty"`
	// A map of variables that is injected into the triggered execution or process instance after the message has been delivered.
	// Each key is a variable name and each value a JSON variable value object with the following properties.
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
