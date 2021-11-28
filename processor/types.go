package processor

import "github.com/citilinkru/camunda-client-go"

// QueryComplete a query for Complete request
type QueryComplete struct {
	// A JSON object containing variable key-value pairs
	Variables *map[string]camunda_client_go.Variable `json:"variables"`
	// A JSON object containing variable key-value pairs.
	// Local variables are set only in the scope of external task
	LocalVariables *map[string]camunda_client_go.Variable `json:"localVariables"`
}

// QueryHandleBPMNError a query for HandleBPMNError request
type QueryHandleBPMNError struct {
	// An error code that indicates the predefined error. Is used to identify the BPMN error handler
	ErrorCode *string `json:"errorCode,omitempty"`
	// An error message that describes the error
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A JSON object containing the variables which will be passed to the execution.
	// Each key corresponds to a variable name and each value to a variable value
	Variables *map[string]camunda_client_go.Variable `json:"variables"`
}

// QueryHandleFailure a query for HandleFailure request
type QueryHandleFailure struct {
	// An message indicating the reason of the failure
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A detailed error description
	ErrorDetails *string `json:"errorDetails,omitempty"`
	// A number of how often the task should be retried.
	// Must be >= 0. If this is 0, an incident is created and the task cannot be fetched anymore unless
	// the retries are increased again. The incident's message is set to the errorMessage parameter
	Retries *int `json:"retries,omitempty"`
	// A timeout in milliseconds before the external task becomes available again for fetching. Must be >= 0
	RetryTimeout *int `json:"retryTimeout,omitempty"`
}
