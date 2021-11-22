package camunda_client_go

import "io/ioutil"

type History struct {
	client *Client
}

// ReqHistoryProcessInstanceQuery a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqHistoryProcessInstanceQuery struct {
	// Filter by process instance id.
	ProcessInstanceId *string `json:"processInstanceId"`
	// Filter by a list of process instance ids. Must be a JSON array of Strings.
	ProcessInstanceIds *[]string `json:"processInstanceIds"`
	// Filter by process instance business key.
	BusinessKey *string `json:"processInstanceBusinessKey"`
	// Filter by process instance business key that the parameter is a substring of.
	BusinessKeyLike *string `json:"processInstanceBusinessKeyLike"`
	// Filter by case instance id.
	CaseInstanceId *string `json:"caseInstanceId"`
	// Filter by the process definition the instances run on.
	ProcessDefinitionId *string `json:"processDefinitionId"`
	// Filter by the key of the process definition the instances run on.
	ProcessDefinitionKey *string `json:"processDefinitionKey"`
	// Filter by a list of process definition keys. A process instance must have one of the
	// given process definition keys. Must be a JSON array of Strings.
	ProcessDefinitionKeyIn *[]string `json:"processDefinitionKeyIn"`
	// Exclude instances by a list of process definition keys. A process instance must not have one of the
	// given process definition keys. Must be a JSON array of Strings.
	ProcessDefinitionKeyNotIn *[]string `json:"processDefinitionKeyNotIn"`
	// Filter by the name of the process definition the instances run on.
	ProcessDefinitionName *string `json:"processDefinitionName"`
	// Filter by process definition names that the parameter is a substring of.
	ProcessDefinitionNameLike *string `json:"processDefinitionNameLike"`
	// Filter by the deployment the id belongs to.
	DeploymentId *string `json:"deploymentId"`
	// Restrict query to all process instances that are sub process instances of the given process instance.
	// Takes a process instance id.
	SuperProcessInstance *string `json:"superProcessInstance"`
	// Restrict query to all process instances that have the given process instance as a sub process instance.
	// Takes a process instance id.
	SubProcessInstance *string `json:"subProcessInstance"`
	// Restrict query to all process instances that are sub process instances of the given case instance.
	// Takes a case instance id.
	SuperCaseInstance *string `json:"superCaseInstance"`
	// Restrict query to all process instances that have the given case instance as a sub-case instance.
	// Takes a case instance id.
	SubCaseInstance *string `json:"subCaseInstance"`
	// Only include active process instances. Value may only be true, as false is the default behavior.
	Active *bool `json:"active"`
	// Only include finished process instances. This flag includes all process instances that are completed or terminated.
	// Value may only be true, as false is the default behavior.
	Finished *bool `json:"finished"`
	// Only include unfinished process instances. Value may only be true, as false is the default behavior.
	Unfinished *bool `json:"unfinished"`
	// Only include suspended process instances. Value may only be true, as false is the default behavior.
	Suspended *bool `json:"suspended"`
	// Restrict to instance that is externally terminated
	ExternallyTerminated *bool `json:"externallyTerminated"`
	// Restrict to instance that is internally terminated
	InternallyTerminated *bool `json:"internallyTerminated"`
	// Filter by presence of incidents. Selects only process instances that have an incident.
	WithIncidents *bool `json:"withIncidents"`
	// Only include process instances which have a root incident. Value may only be true, as false is the default behavior.
	WithRootIncidents *bool `json:"withRootIncidents"`
	// Filter by the incident type.
	IncidentType *string `json:"incidentType"`
	// Only include process instances which have an incident in status either open or resolved.
	// To get all process instances, use the query parameter withIncidents.
	IncidentStatus *bool `json:"incidentStatus"`
	// Filter by the incident message.
	IncidentMessage *string `json:"incidentMessage"`
	// Filter by the incident message that the parameter is a substring of.
	IncidentMessageLike *string `json:"incidentMessageLike"`
	// Filter by a list of tenant ids. A process instance must have one of the given tenant ids.
	// Must be a JSON array of Strings.
	TenantIdIn *[]string `json:"tenantIdIn"`
	// Only include process instances which belong to no tenant. Value may only be true, as false is the default behavior.
	WithoutTenantId *bool `json:"withoutTenantId"`
	// Filter by a list of activity ids. A process instance must currently wait in a leaf activity with one of the given activity ids.
	ActivityIdIn *[]string `json:"activityIdIn"`
	// Restrict to instance that executed an activity with one of given ids.
	ExecutedActivityIdIn *[]string `json:"executedActivityIdIn"`
	// Restrict the query to all process instances that are top level process instances.
	RootProcessInstances *bool `json:"rootProcessInstances"`
	// Restrict the query to all process instances that are leaf instances. (i.e. don't have any sub instances)
	LeafProcessInstances *bool `json:"leafProcessInstances"`
	// Only include process instances which process definition has no tenant id.
	ProcessDefinitionWithoutTenantId *bool `json:"processDefinitionWithoutTenantId"`
	// A JSON array to only include process instances that have variables with certain values.
	Variables *[]ReqProcessVariableQuery `json:"variables"`
	// Match all variable names in this query case-insensitively.
	// If set to true variable-Name and variable-name are treated as equal.
	VariableNamesIgnoreCase *bool `json:"variableNamesIgnoreCase"`
	// Match all variable values in this query case-insensitively.
	// If set to true variable-Value and variable-value are treated as equal.
	VariableValuesIgnoreCase *bool `json:"variableValuesIgnoreCase"`
	// A JSON array of nested process instance queries with OR semantics.
	// A process instance matches a nested query if it fulfills at least one of the query's predicates.
	// With multiple nested queries, a process instance must fulfill at least one predicate of each query.
	// All process instance query properties can be used except for: sorting.
	OrQueries *[]ReqProcessInstanceQuery `json:"orQueries"`
	// A JSON array of criteria to sort the result by.
	// Each element of the array is a JSON object that specifies one ordering.
	// The position in the array identifies the rank of an ordering, i.e., whether it is primary, secondary, etc.
	Sorting *[]ReqSort `json:"sorting"`
	// Restrict to instance that was started before the given date.
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	StartedBefore *string `json:"startedBefore"`
	// Restrict to instance that was started after the given date.
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	StartedAfter *string `json:"startedAfter"`
	// Restrict to instance that was finished before the given date.
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	FinishedBefore *string `json:"finishedBefore"`
	// Restrict to instance that was finished after the given date.
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	FinishedAfter *string `json:"finishedAfter"`
	// Restrict to instance that executed an activity before the given date (inclusive).
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedActivityBefore *string `json:"executedActivityBefore"`
	// Restrict to instance that executed an activity after the given date (inclusive).
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedActivityAfter *string `json:"executedActivityAfter"`
	// Restrict to instance that executed a job before the given date (inclusive).
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedJobBefore *string `json:"executedJobBefore"`
	// Restrict to instance that executed a job after the given date (inclusive).
	// By default, the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200.
	ExecutedJobAfter *string `json:"executedJobAfter"`
}

// ReqHistoryDeleteProcessInstance a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqHistoryDeleteProcessInstance struct {
	// A list process instance ids to delete.
	HistoricProcessInstanceIds *[]string `json:"historicProcessInstanceIds,omitempty"`
	// A process instance query
	HistoricProcessInstanceQuery *ReqProcessInstanceQuery `json:"historicProcessInstanceQuery,omitempty"`
	// A string with delete reason.
	DeleteReason *string `json:"deleteReason,omitempty"`
}

// ReqHistoryVariableInstanceQuery a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqHistoryVariableInstanceQuery struct {
	VariableName         *string     `json:"variableName"`
	VariableNameLike     *string     `json:"variableNameLike"`
	VariableValue        interface{} `json:"variableValue"`
	ProcessInstanceId    *string     `json:"processInstanceId"`
	ProcessInstanceIdIn  *[]string   `json:"processInstanceIdIn"`
	ProcessDefinitionId  *string     `json:"process_definition_id"`
	ExecutionIdIn        *[]string   `json:"executionIdIn"`
	CaseInstanceId       *string     `json:"caseInstanceId"`
	CaseExecutionIdIn    *[]string   `json:"caseExecutionIdIn"`
	CaseActivityIdIn     *[]string   `json:"caseActivityIdIn"`
	TaskIdIn             *[]string   `json:"taskIdIn"`
	ActivityInstanceIdIn *[]string   `json:"activityInstanceIdIn"`
	TenantIdIn           *[]string   `json:"tenantIdIn"`
}

// ResHistoryProcessInstance a response object for process instance
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

// ResHistoryProcessInstanceDurationReport a response object for history process instance duration report
type ResHistoryProcessInstanceDurationReport struct {
	// Specifies a timespan within a year.
	// The period must be interpreted in conjunction with the returned periodUnit.
	Period int `json:"period"`
	// The unit of the given period. Possible values are MONTH and QUARTER.
	PeriodUnit string `json:"periodUnit"`
	// The greatest duration in milliseconds of all completed process instances which were started in the given period.
	Maximum int64 `json:"maximum"`
	// The smallest duration in milliseconds of all completed process instances which were started in the given period.
	Minimum int64 `json:"minimum"`
	// The average duration in milliseconds of all completed process instances which were started in the given period.
	Average int64 `json:"average"`
}

// ResHistoryVariableInstance a response object for history variable instance
type ResHistoryVariableInstance struct {
	// The id of the variable instance.
	Id string `json:"id"`
	// The name of the variable instance.
	Name string `json:"name"`
	// The value type of the variable. Can be String/Number/Boolean/Object
	Type string `json:"type"`
	// The variable's value. Value differs depending on the variable's type and on the deserializeValues parameter.
	Value interface{} `json:"value"`
	// A JSON object containing additional, value-type-dependent properties.
	ValueInfo ResProcessVariableValueInfo `json:"valueInfo"`
	// The key of the process definition the variable instance belongs to.
	ProcessDefinitionKey string `json:"processDefinitionKey"`
	// The id of the process definition the variable instance belongs to.
	ProcessDefinitionId string `json:"processDefinitionId"`
	// The process instance id the variable instance belongs to.
	ProcessInstanceId string `json:"processInstanceId"`
	// The execution id the variable instance belongs to.
	ExecutionId string `json:"executionId"`
	// The id of the activity instance in which the variable is valid.
	ActivityInstanceId string `json:"activityInstanceId"`
	// The key of the case definition the variable instance belongs to.
	CaseDefinitionKey string `json:"caseDefinitionKey"`
	// The id of the case definition the variable instance belongs to.
	CaseDefinitionId string `json:"caseDefinitionId"`
	// The case instance id the variable instance belongs to.
	CaseInstanceId string `json:"caseInstanceId"`
	// The case execution id the variable instance belongs to.
	CaseExecutionId string `json:"caseExecutionId"`
	// The id of the task the variable instance belongs to.
	TaskId string `json:"taskId"`
	// The id of the tenant that this variable instance belongs to.
	TenantId string `json:"tenantId"`
	// An error message in case a Java Serialized Object could not be de-serialized.
	ErrorMessage string `json:"errorMessage"`
}

// GetProcessInstanceCount queries for the number of historic process instances that fulfill the given parameters.
// https://docs.camunda.org/manual/latest/reference/rest/history/process-instance/get-process-instance-query-count/#query-parameters
func (h *History) GetProcessInstanceCount(query map[string]string) (count int, err error) {
	resCount := ResCount{}
	res, err := h.client.doGet("/history/process-instance/count", query)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetProcessInstanceList queries for historic process instances that fulfill the given parameters.
// https://docs.camunda.org/manual/latest/reference/rest/history/process-instance/get-process-instance-query/#query-parameters
func (h *History) GetProcessInstanceList(query map[string]string) (processInstances []*ResHistoryProcessInstance, err error) {
	res, err := h.client.doGet("/history/process-instance", query)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &processInstances)
	return
}

// GetProcessInstance Retrieves a historic process instance by id, according to the HistoricProcessInstance interface in the engine.
func (h *History) GetProcessInstance(id string) (processInstance *ResHistoryProcessInstance, err error) {
	processInstance = &ResHistoryProcessInstance{}
	res, err := h.client.doGet("/history/process-instance/"+id, nil)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &processInstance)
	return
}

// GetProcessInstanceCountPost queries for the number of historic process instances that fulfill the given parameters.
func (h *History) GetProcessInstanceCountPost(req ReqHistoryProcessInstanceQuery) (count int, err error) {
	resCount := ResCount{}
	res, err := h.client.doPostJson("/history/process-instance/count", nil, req)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetProcessInstanceListPost queries for historic process instances that fulfill given parameters through a JSON object.
// https://docs.camunda.org/manual/latest/reference/rest/history/process-instance/post-process-instance-query/#query-parameters
func (h *History) GetProcessInstanceListPost(query map[string]string, req ReqHistoryProcessInstanceQuery) (processInstances []*ResHistoryProcessInstance, err error) {
	res, err := h.client.doPostJson("/history/process-instance", query, req)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &processInstances)
	return
}

// DeleteProcessInstance deletes a historic process instance by id, according to the HistoricProcessInstance interface in the engine.
func (h *History) DeleteProcessInstance(id string) error {
	return h.client.doDelete("/history/process-instance/"+id, nil)
}

// DeleteProcessInstanceAsync deletes multiple history process instances asynchronously (batch).
func (h *History) DeleteProcessInstanceAsync(req ReqHistoryDeleteProcessInstance) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := h.client.doPostJson("/history/process-instance/delete", nil, req)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, batch)
	return
}

// GetProcessInstanceDurationReport retrieves a report about the duration of completed process instances, grouped by a period.
// These reports include the maximum, minimum and average duration of all completed process instances which were started in a given period.
// This only includes historic data.
// https://docs.camunda.org/manual/latest/reference/rest/history/process-instance/get-duration-report/#query-parameters
func (h *History) GetProcessInstanceDurationReport(query map[string]string) (reports []*ResHistoryProcessInstanceDurationReport, err error) {
	res, err := h.client.doGet("/history/process-instance/report?reportType=duration", query)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &reports)
	return
}

// GetVariableInstanceCount queries for the number of historic variable instances that fulfill the given parameters.
// https://docs.camunda.org/manual/latest/reference/rest/history/variable-instance/get-variable-instance-query/#query-parameters
func (h *History) GetVariableInstanceCount(query map[string]string) (count int, err error) {
	resCount := ResCount{}
	res, err := h.client.doGet("/history/variable-instance/count", query)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetVariableInstanceList queries for historic variable instances that fulfill the given parameters.
// https://docs.camunda.org/manual/latest/reference/rest/history/variable-instance/get-variable-instance-query/#query-parameters
func (h *History) GetVariableInstanceList(query map[string]string) (variableInstances []*ResHistoryVariableInstance, err error) {
	res, err := h.client.doGet("/history/variable-instance", query)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &variableInstances)
	return
}

// GetVariableInstance retrieves a historic variable by id.
// https://docs.camunda.org/manual/latest/reference/rest/history/variable-instance/get-variable-instance/#query-parameters
func (h *History) GetVariableInstance(id string, query map[string]string) (variableInstance *ResHistoryVariableInstance, err error) {
	variableInstance = &ResHistoryVariableInstance{}
	res, err := h.client.doGet("/history/variable-instance/"+id, query)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &variableInstance)
	return
}

// GetVariableInstanceBinaryData retrieves the content of a historic variable by id. Applicable for variables
// that are serialized as binary data.
func (h *History) GetVariableInstanceBinaryData(id string) (data []byte, err error) {
	res, err := h.client.doGet("/history/variable-instance/"+id+"/data", nil)
	if err != nil {
		return
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// GetVariableInstanceCountPost queries for historic variable instances that fulfill the given parameters.
func (h *History) GetVariableInstanceCountPost(req ReqHistoryVariableInstanceQuery) (count int, err error) {
	resCount := ResCount{}
	res, err := h.client.doPostJson("/history/variable-instance/count", nil, req)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetVariableInstanceListPost queries for historic variable instances that fulfill the given parameters.
// https://docs.camunda.org/manual/latest/reference/rest/history/variable-instance/post-variable-instance-query/#query-parameters
func (h *History) GetVariableInstanceListPost(query map[string]string, req ReqHistoryVariableInstanceQuery) (variableInstances []*ResHistoryVariableInstance, err error) {
	res, err := h.client.doPostJson("/history/variable-instance", query, req)
	if err != nil {
		return
	}

	err = h.client.readJsonResponse(res, &variableInstances)
	return
}
