package camunda_client_go

type History struct {
	client *Client
}

type QueryHistoryProcessInstanceBy struct {
	// The id of the historic process instance to be retrieved.
	Id string
}

func (q QueryHistoryProcessInstanceBy) String() string {
	return q.Id
}

type ResHistoryProcessInstance struct {
	// The id of the process instance
	Id string `json:"id"`
	// The process instance id of the root process instance that initiated the process.
	RootProcessInstanceId string `json:"rootProcessInstanceId"`
	// The id of the parent process instance, if it exists.
	SuperProcessInstanceId string `json:"superProcessInstanceId"`
	// The id of the parent case instance, if it exists.
	SuperCaseInstanceId string `json:"superCaseInstanceId"`
	// The id of the parent case instance, if it exists.
	CaseInstanceId string `json:"caseInstanceId"`
	// The name of the process definition that this process instance belongs to.
	ProcessDefinitionName string `json:"processDefinitionName"`
	// The key of the process definition that this process instance belongs to.
	ProcessDefinitionKey string `json:"processDefinitionKey"`
	// The version of the process definition that this process instance belongs to.
	ProcessDefinitionVersion int `json:"processDefinitionVersion"`
	// The id of the process definition that this process instance belongs to.
	ProcessDefinitionId string `json:"processDefinitionId"`
	// The business key of the process instance.
	BusinessKey string `json:"businessKey"`
	// The time the instance was started. Default format* yyyy-MM-dd’T’HH:mm:ss.SSSZ.
	StartTime string `json:"startTime"`
	// The time the instance ended. Default format* yyyy-MM-dd’T’HH:mm:ss.SSSZ.
	EndTime string `json:"endTime"`
	// The time after which the instance should be removed by the History Cleanup job. Default format* yyyy-MM-dd’T’HH:mm:ss.SSSZ.
	RemovalTime string `json:"removalTime"`
	// The time the instance took to finish (in milliseconds).
	DurationInMillis float32 `json:"durationInMillis"`
	// The id of the user who started the process instance.
	StartUserId string `json:"startUserId"`
	// The id of the initial activity that was executed (e.g., a start event).
	StartActivityId string `json:"startActivityId"`
	// The provided delete reason in case the process instance was canceled during execution.
	DeleteReason string `json:"deleteReason"`
	// The tenant id of the process instance.
	TenantId string `json:"tenantId"`
	// last state of the process instance, possible values are:
	// ACTIVE - running process instance
	// SUSPENDED - suspended process instances
	// COMPLETED - completed through normal end event
	// EXTERNALLY_TERMINATED - terminated externally, for instance through REST API
	// INTERNALLY_TERMINATED - terminated internally, for instance by terminating boundary event
	State string `json:"state"`
}

// GetProcessInstance Retrieves a historic process instance by id, according to the HistoricProcessInstance interface in the engine.
// https://docs.camunda.org/manual/latest/reference/rest/history/process-instance/get-process-instance/
func (h *History) GetProcessInstance(by *QueryHistoryProcessInstanceBy) (processInstance *ResHistoryProcessInstance, err error) {
	processInstance = &ResHistoryProcessInstance{}
	res, err := h.client.doGet("/history/process-instance/"+by.String(), map[string]string{})
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &processInstance)
	return
}
