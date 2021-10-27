package camunda_client_go

import "io/ioutil"

// ProcessInstance a client for ProcessInstance API
type ProcessInstance struct {
	client *Client
}

// ReqProcessVariableValueInfo a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessVariableValueInfo struct {
	// A string representation of the object's type name.
	ObjectTypeName *string `json:"objectTypeName,omitempty"`
	// The serialization format used to store the variable.
	SerializationDataFormat *string `json:"serializationDataFormat,omitempty"`
	// The name of the file. This is not the variable name but the name that will be used when downloading the file again.
	FileName *string `json:"filename,omitempty"`
	// The MIME type of the file that is being uploaded.
	MimeType *string `json:"mimetype,omitempty"`
	// The encoding of the file that is being uploaded.
	Encoding *string `json:"encoding,omitempty"`
	// Indicates whether the variable should be transient or not. See documentation for more information.
	Transient *bool `json:"transient,omitempty"`
	// Indicates whether the variable should be a local variable or not.
	// If set to true, the variable becomes a local variable of the execution entering the target activity.
	Local *bool `json:"local"`
}

// ReqProcessVariable a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessVariable struct {
	// The variable's value. For variables of type Object, the serialized value has to be submitted as a String value.
	// For variables of type File the value has to be submitted as Base64 encoded string.
	Value interface{} `json:"value,omitempty"`
	// The value type of the variable.
	Type *string `json:"type,omitempty"`
	// A JSON object containing additional, value-type-dependent properties.
	ValueInfo *ReqProcessVariableValueInfo `json:"valueInfo,omitempty"`
}

// ReqModifyProcessVariables a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqModifyProcessVariables struct {
	// A JSON object containing variable key-value pairs. Each key is a variable name and each value a JSON.
	Modifications *map[string]ReqProcessVariable `json:"modifications,omitempty"`
	// An array of String keys of variables to be deleted.
	Deletions *[]string `json:"deletions,omitempty"`
}

// ReqModifyProcessInstanceInstruction a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqModifyProcessInstanceInstruction struct {
	// Mandatory. One of the following values: cancel, startBeforeActivity, startAfterActivity, startTransition.
	// A cancel instruction requests cancellation of a single activity instance or all instances of one activity.
	// A startBeforeActivity instruction requests to enter a given activity.
	// A startAfterActivity instruction requests to execute the single outgoing sequence flow of a given activity.
	// A startTransition instruction requests to execute a specific sequence flow.
	Type string `json:"type"`
	// Can be used with instructions of types startBeforeActivity, startAfterActivity, and cancel.
	// Specifies the activity the instruction targets.
	ActivityId *string `json:"activityId"`
	// Can be used with instructions of types startTransition. Specifies the sequence flow to start.
	TransitionId *string `json:"transitionId"`
	// Can be used with instructions of type cancel. Specifies the activity instance to cancel.
	// Valid values are the activity instance IDs supplied by the Get Activity Instance request.
	ActivityInstanceId *string `json:"activityInstanceId"`
	// Can be used with instructions of type cancel. Specifies the transition instance to cancel.
	// Valid values are the transition instance IDs supplied by the Get Activity Instance request.
	TransitionInstanceId *string `json:"transitionInstanceId"`
	// Can be used with instructions of type startBeforeActivity, startAfterActivity, and startTransition.
	// Valid values are the activity instance IDs supplied by the Get Activity Instance request.
	// If there are multiple parent activity instances of the targeted activity, this specifies the ancestor scope
	// in which hierarchy the activity/transition is to be instantiated.
	//Example: When there are two instances of a subprocess and an activity contained in the subprocess is to be started,
	// this parameter allows specifying under which subprocess instance the activity should be started.
	AncestorActivityInstanceId *string `json:"ancestorActivityInstanceId"`
	// Can be used with instructions of type startBeforeActivity, startAfterActivity, and startTransition.
	// A JSON object containing variable key-value pairs.
	// Each key is a variable name and each value a JSON variable value object.
	Variables *[]ReqProcessVariable `json:"variables"`
}

// ReqModifyProcessInstance a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqModifyProcessInstance struct {
	// Skip execution listener invocation for activities that are started or ended as part of this request.
	SkipCustomListeners *bool `json:"skipCustomListeners"`
	// Skip execution of input/output variable mappings for activities that are started or ended as part of this request.
	SkipIOMappings *bool `json:"skipIoMappings"`
	// A JSON array of modification instructions. The instructions are executed in the order they are in.
	Instructions *[]ReqModifyProcessInstanceInstruction `json:"instructions"`
	// An arbitrary text annotation set by a user for auditing reasons.
	Annotation *string `json:"annotation"`
}

// ReqProcessVariableQuery a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessVariableQuery struct {
	// Process variable name
	Name *string `json:"name"`
	// Valid operator values are: eq - equal to; neq - not equal to; gt - greater than;
	// gteq - greater than or equal to; lt - lower than; lteq - lower than or equal to; like.
	Operator *string `json:"operator"`
	// Process variable value
	Value interface{} `json:"value"`
}

// ReqProcessInstanceQuery a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessInstanceQuery struct {
	// Filter by a list of process instance ids. Must be a JSON array of Strings.
	ProcessInstanceIds *[]string `json:"processInstanceIds"`
	// Filter by process instance business key.
	BusinessKey *string `json:"businessKey"`
	// Filter by process instance business key that the parameter is a substring of.
	BusinessKeyLike *string `json:"businessKeyLike"`
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
	// Only include suspended process instances. Value may only be true, as false is the default behavior.
	Suspended *bool `json:"suspended"`
	// Filter by presence of incidents. Selects only process instances that have an incident.
	WithIncident *bool `json:"withIncident"`
	// Filter by the incident id.
	IncidentId *string `json:"incidentId"`
	// Filter by the incident type.
	IncidentType *string `json:"incidentType"`
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
}

// ReqDeleteProcessInstance a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqDeleteProcessInstance struct {
	// A list process instance ids to delete.
	ProcessInstanceIds *[]string `json:"processInstanceIds"`
	// A process instance query
	ProcessInstanceQuery *ReqProcessInstanceQuery `json:"processInstanceQuery"`
	// A string with delete reason.
	DeleteReason *string `json:"deleteReason"`
	// Skip execution listener invocation for activities that are started or ended as part of this request.
	SkipCustomListeners *bool `json:"skipCustomListeners"`
	// Skip deletion of the subprocesses related to deleted processes as part of this request.
	SkipSubprocesses *bool `json:"skipSubprocesses"`
	// If set to false, the request will still be successful if one or more of the process ids are not found.
	FailIfNotExists *bool `json:"failIfNotExists"`
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

// ReqDeleteHistoryProcessInstance a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqDeleteHistoryProcessInstance struct {
	// A list process instance ids to delete.
	ProcessInstanceIds *[]string `json:"processInstanceIds"`
	// A historic process instance query
	HistoricProcessInstanceQuery *ReqHistoryProcessInstanceQuery `json:"historicProcessInstanceQuery"`
	// A string with delete reason.
	DeleteReason *string `json:"deleteReason"`
	// Skip execution listener invocation for activities that are started or ended as part of this request.
	SkipCustomListeners *bool `json:"skipCustomListeners"`
	// Skip deletion of the subprocesses related to deleted processes as part of this request.
	SkipSubprocesses *bool `json:"skipSubprocesses"`
	// If set to false, the request will still be successful if one or more of the process ids are not found.
	FailIfNotExists *bool `json:"failIfNotExists"`
}

// ReqProcessInstanceJobRetries a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessInstanceJobRetries struct {
	// A list of process instance ids to fetch jobs, for which retries will be set.
	ProcessInstances *[]string `json:"processInstances"`
	// A process instance query
	ProcessInstanceQuery *ReqProcessInstanceQuery `json:"processInstanceQuery"`
	// An integer representing the number of retries. Please note that the value cannot be negative or null.
	Retries int `json:"retries"`
}

// ReqHistoricProcessInstanceJobRetries a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqHistoricProcessInstanceJobRetries struct {
	// A list of process instance ids to fetch jobs, for which retries will be set.
	ProcessInstances *[]string `json:"processInstances"`
	// A process instance query
	HistoricProcessInstanceQuery *ReqHistoryProcessInstanceQuery `json:"historicProcessInstanceQuery"`
	// An integer representing the number of retries. Please note that the value cannot be negative or null.
	Retries int `json:"retries"`
}

// ReqProcessInstanceVariables a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessInstanceVariables struct {
	// A list process instance ids to delete.
	ProcessInstanceIds *[]string `json:"processInstanceIds"`
	// A process instance query
	ProcessInstanceQuery *ReqProcessInstanceQuery `json:"processInstanceQuery"`
	// A historic process instance query
	HistoricProcessInstanceQuery *ReqHistoryProcessInstanceQuery `json:"historicProcessInstanceQuery"`
	// A JSON object containing variable key-value pairs the operation will set in the root scope of the process instances.
	// Each key is a variable name and each value a JSON variable value object.
	Variables *ReqProcessVariable `json:"variables"`
}

// ReqProcessInstanceActivateSuspend a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessInstanceActivateSuspend struct {
	// A list process instance ids to delete.
	ProcessInstanceIds *[]string `json:"processInstanceIds"`
	// A process instance query
	ProcessInstanceQuery *ReqProcessInstanceQuery `json:"processInstanceQuery"`
	// A historic process instance query
	HistoricProcessInstanceQuery *ReqHistoryProcessInstanceQuery `json:"historicProcessInstanceQuery"`
	// The process definition id of the process instances to activate or suspend.
	ProcessDefinitionId *string `json:"processDefinitionId"`
	// The process definition key of the process instances to activate or suspend.
	ProcessDefinitionKey *string `json:"processDefinitionKey"`
	// Only activate or suspend process instances of a process definition which belongs to a tenant with the given id.
	ProcessDefinitionTenantId *string `json:"processDefinitionTenantId"`
	// Only activate or suspend process instances of a process definition which belongs to no tenant.
	// Value may only be true, as false is the default behavior.
	ProcessDefinitionWithoutTenantId *bool `json:"processDefinitionWithoutTenantId"`
	// A Boolean value which indicates whether to activate or suspend a given process instance.
	// When the value is set to true, the given process instance will be suspended and
	// when the value is set to false, the given process instance will be activated.
	Suspended bool `json:"suspended"`
}

// ResProcessVariableValueInfo variable value info
type ResProcessVariableValueInfo struct {
	// A string representation of the object's type name.
	ObjectTypeName string `json:"objectTypeName"`
	// The serialization format used to store the variable.
	SerializationDataFormat string `json:"serializationDataFormat"`
}

// ResProcessVariable a response object for process variable
type ResProcessVariable struct {
	// The variable's value. Value differs depending on the variable's type and on the deserializeValues parameter.
	Value interface{} `json:"value"`
	// The value type of the variable.
	Type string `json:"type"`
	// A JSON object containing additional, value-type-dependent properties.
	ValueInfo ResProcessVariableValueInfo `json:"valueInfo"`
}

// ResActivityInstanceIncident a response object for process activity instance incident
type ResActivityInstanceIncident struct {
	// The id of the incident
	Id string `json:"id"`
	// The activity id in which the incident happened
	ActivityId string `json:"activityId"`
}

// ResProcessTransitionInstance a response object for process transition instance
type ResProcessTransitionInstance struct {
	// The id of the activity instance.
	Id string `json:"id"`
	// The id of the parent activity instance.
	ParentActivityInstanceId string `json:"parentActivityInstanceId"`
	// The id of the activity.
	ActivityId string `json:"activityId"`
	// The name of the activity.
	ActivityName string `json:"activityName"`
	// The type of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job).
	// Corresponds to the XML element name in the BPMN 2.0, e.g., 'userTask'.
	ActivityType string `json:"activityType"`
	// The id of the process instance.
	ProcessInstanceId string `json:"processInstanceId"`
	// The id of the process definition.
	ProcessDefinitionId string `json:"processDefinitionId"`
	// A list of execution ids.
	ExecutionIds []string `json:"executionIds"`
	// A list of incident ids.
	IncidentIds []string `json:"incidentIds"`
	// A list of JSON objects containing incident specific properties.
	Incidents []ResActivityInstanceIncident `json:"incidents"`
}

// ResProcessActivityInstance a response object for process activity instance
type ResProcessActivityInstance struct {
	// The id of the activity instance.
	Id string `json:"id"`
	// The id of the parent activity instance.
	ParentActivityInstanceId string `json:"parentActivityInstanceId"`
	// The id of the activity.
	ActivityId string `json:"activityId"`
	// The name of the activity.
	ActivityName string `json:"activityName"`
	// The type of the activity that this instance enters (asyncBefore job) or leaves (asyncAfter job).
	// Corresponds to the XML element name in the BPMN 2.0, e.g., 'userTask'.
	ActivityType string `json:"activityType"`
	// The id of the process instance.
	ProcessInstanceId string `json:"processInstanceId"`
	// The id of the process definition.
	ProcessDefinitionId string `json:"processDefinitionId"`
	// A list of child activity instances.
	ChildActivityInstances []ResProcessActivityInstance `json:"childActivityInstances"`
	// A list of child transition instances. A transition instance represents an execution
	// waiting in an asynchronous continuation.
	ChildTransitionInstances []ResProcessTransitionInstance `json:"childTransitionInstances"`
	// A list of execution ids.
	ExecutionIds []string `json:"executionIds"`
	// A list of incident ids.
	IncidentIds []string `json:"incidentIds"`
	// A list of JSON objects containing incident specific properties.
	Incidents []ResActivityInstanceIncident `json:"incidents"`
}

// ResProcessInstance a response object for process instance
type ResProcessInstance struct {
	// The id of the process instance.
	Id string `json:"id"`
	// The id of the process definition that this process instance belongs to.
	DefinitionId string `json:"definitionId"`
	// The business key of the process instance.
	BusinessKey string `json:"businessKey"`
	// The id of the case instance associated with the process instance.
	CaseInstanceId string `json:"caseInstanceId"`
	// A flag indicating whether the process instance is suspended or not.
	Suspended bool `json:"suspended"`
	// The tenant id of the process instance.
	TenantId string `json:"tenantId"`
	// A JSON array containing links to interact with the instance
	Links []ResLink `json:"links"`
}

// QueryProcessInstanceVariableBy path builder
type QueryProcessInstanceVariableBy struct {
	Id           *string
	VariableName *string
}

// String a build path part
func (q *QueryProcessInstanceVariableBy) String() string {
	if q.Id != nil && q.VariableName != nil {
		return "/process-instance/" + *q.Id + "/variables/" + *q.VariableName
	} else if q.Id != nil {
		return "/process-instance/" + *q.Id
	}
	return ""
}

// DeleteProcessVariable deletes a variable of a process instance by id.
func (p *ProcessInstance) DeleteProcessVariable(by QueryProcessInstanceVariableBy) error {
	err := p.client.doDelete(by.String(), nil)
	return err
}

// GetBinaryProcessVariableData retrieves the content of a Process Variable by the Process Instance id and the
// Process Variable name. Applicable for byte array or file Process Variables.
func (p *ProcessInstance) GetBinaryProcessVariableData(by QueryProcessInstanceVariableBy) (data []byte, err error) {
	res, err := p.client.doGet(by.String()+"/data", nil)
	if err != nil {
		return
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// GetProcessVariable retrieves a variable of a given process instance by id.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/variables/get-variable/#query-parameters
func (p *ProcessInstance) GetProcessVariable(by QueryProcessInstanceVariableBy, query map[string]string) (processVariable *ResProcessVariable, err error) {
	processVariable = &ResProcessVariable{}
	res, err := p.client.doGet(by.String(), query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, processVariable)
	return
}

// GetProcessVariableList retrieves all variables of a given process instance by id.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/variables/get-variables/#query-parameters
func (p *ProcessInstance) GetProcessVariableList(id string, query map[string]string) (processVariables map[string]*ResProcessVariable, err error) {
	res, err := p.client.doGet("/process-instance/"+id+"/variables", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &processVariables)
	return
}

// ModifyProcessVariables updates or deletes the variables of a process instance by id. Updates precede deletions.
// So, if a variable is updated AND deleted, the deletion overrides the update.
func (p *ProcessInstance) ModifyProcessVariables(id string, req ReqModifyProcessVariables) error {
	_, err := p.client.doPostJson("/process-instance/"+id+"/variables", nil, req)
	return err
}

// UpdateProcessVariable sets a variable of a given process instance by id.
func (p *ProcessInstance) UpdateProcessVariable(by QueryProcessInstanceVariableBy, req ReqProcessVariable) error {
	_, err := p.client.doPostJson(by.String(), nil, req)
	return err
}

// Delete deletes a running process instance by id.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/delete/#query-parameters
func (p *ProcessInstance) Delete(id string, query map[string]string) error {
	err := p.client.doDelete("/process-instance/"+id, query)
	return err
}

// GetActivityInstance retrieves an Activity Instance (Tree) for a given process instance by id.
func (p *ProcessInstance) GetActivityInstance(id string) (instance *ResProcessActivityInstance, err error) {
	instance = &ResProcessActivityInstance{}
	res, err := p.client.doGet("/process-instance/"+id+"/activity-instances", nil)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, instance)
	return
}

// GetCount queries for the number of process instances that fulfill given parameters.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/get-query-count/#query-parameters
func (p *ProcessInstance) GetCount(query map[string]string) (count int, err error) {
	resCount := ResCount{}
	res, err := p.client.doGet("/process-instance/count", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetList queries for process instances that fulfill given parameters.
// Parameters may be static as well as dynamic runtime properties of process instances.
// The size of the result set can be retrieved by using the GetCount method.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/get-query/#query-parameters
func (p *ProcessInstance) GetList(query map[string]string) (processInstances []*ResProcessInstance, err error) {
	res, err := p.client.doGet("/process-instance", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &processInstances)
	return
}

// Get retrieves a process instance by id, according to the ProcessInstance interface in the engine.
func (p *ProcessInstance) Get(id string) (processInstance *ResProcessInstance, err error) {
	processInstance = &ResProcessInstance{}
	res, err := p.client.doGet("/process-instance/"+id, nil)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, processInstance)
	return
}

// Modify submits a list of modification instructions to change a process instance's execution state.
// A modification instruction is one of the following:
//    Starting execution before an activity
//    Starting execution after an activity on its single outgoing sequence flow
//    Starting execution on a specific sequence flow
//    Cancelling an activity instance, transition instance, or all instances (activity or transition) for an activity
// Instructions are executed immediately and in the order they are provided in this request's body.
// Variables can be provided with every starting instruction.
func (p *ProcessInstance) Modify(id string, req ReqModifyProcessInstance) error {
	_, err := p.client.doPostJson("/process-instance/"+id+"/modification", nil, req)
	return err
}

// ModifyAsync submits a list of modification instructions to change a process instance's execution state.
// A modification instruction is one of the following:
//    Starting execution before an activity
//    Starting execution after an activity on its single outgoing sequence flow
//    Starting execution on a specific sequence flow
//    Cancelling an activity instance, transition instance, or all instances (activity or transition) for an activity
// Instructions are executed asynchronous and in the order they are provided in this request's body.
// Variables can be provided with every starting instruction.
func (p *ProcessInstance) ModifyAsync(id string, req ReqModifyProcessInstance) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := p.client.doPostJson("/process-instance/"+id+"/modification-async", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, batch)
	return
}

// DeleteAsync deletes multiple process instances asynchronously (batch).
func (p *ProcessInstance) DeleteAsync(req ReqDeleteProcessInstance) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := p.client.doPostJson("/process-instance/delete", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, batch)
	return
}

// DeleteHistoryAsync deletes a set of process instances asynchronously (batch) based on a historic process instance query.
func (p *ProcessInstance) DeleteHistoryAsync(req ReqDeleteHistoryProcessInstance) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := p.client.doPostJson("/process-instance/delete-historic-query-based", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, batch)
	return
}

// GetCountPost queries for the number of process instances that fulfill the given parameters.
func (p *ProcessInstance) GetCountPost(req ReqProcessInstanceQuery) (count int, err error) {
	resCount := ResCount{}
	res, err := p.client.doPostJson("/process-instance/count", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetListPost queries for process instances that fulfill given parameters through a JSON object.
func (p *ProcessInstance) GetListPost(query map[string]string, req ReqProcessInstanceQuery) (processInstances []*ResProcessInstance, err error) {
	res, err := p.client.doPostJson("/process-instance", query, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &processInstances)
	return
}

// SetJobRetriesAsync creates a batch to set retries of jobs associated with given processes asynchronously.
func (p *ProcessInstance) SetJobRetriesAsync(req ReqProcessInstanceJobRetries) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := p.client.doPostJson("/process-instance/job-retries", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, batch)
	return
}

// SetHistoricJobRetriesAsync creates a batch to set retries of jobs based on a historic process instance query asynchronously.
func (p *ProcessInstance) SetHistoricJobRetriesAsync(req ReqHistoricProcessInstanceJobRetries) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := p.client.doPostJson("/process-instance/job-retries-historic-query-based", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, batch)
	return
}

// SetVariablesAsync updates or creates runtime process variables in the root scope of process instances.
func (p *ProcessInstance) SetVariablesAsync(req ReqProcessInstanceVariables) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := p.client.doPostJson("/process-instance/variables-async", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, batch)
	return
}

// ActivateSuspend activates or suspends a given process instance by id.
func (p *ProcessInstance) ActivateSuspend(id string, req ReqProcessInstanceActivateSuspend) error {
	return p.client.doPutJson("/process-instance/"+id+"/suspended", nil, req)
}

// ActivateSuspendByProcessDefinitionId activates or suspends process instances with the given process definition id.
func (p *ProcessInstance) ActivateSuspendByProcessDefinitionId(req ReqProcessInstanceActivateSuspend) error {
	return p.client.doPutJson("/process-instance/suspended", nil, req)
}

// ActivateSuspendByProcessDefinitionKey activates or suspends process instances with the given process definition key.
func (p *ProcessInstance) ActivateSuspendByProcessDefinitionKey(req ReqProcessInstanceActivateSuspend) error {
	return p.client.doPutJson("/process-instance/suspended", nil, req)
}

// ActivateSuspendInGroup activates or suspends process instances synchronously with a list of process instance ids,
// a process instance query, and/or a historical process instance query
func (p *ProcessInstance) ActivateSuspendInGroup(req ReqProcessInstanceActivateSuspend) error {
	return p.client.doPutJson("/process-instance/suspended", nil, req)
}

// ActivateSuspendInGroupAsync activates or suspends process instances asynchronously with a list of process
// instance ids, a process instance query, and/or a historical process instance query
func (p *ProcessInstance) ActivateSuspendInGroupAsync(req ReqProcessInstanceActivateSuspend) (batch *ResBatch, err error) {
	batch = &ResBatch{}
	res, err := p.client.doPostJson("/process-instance/suspended-async", nil, req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, batch)
	return
}
