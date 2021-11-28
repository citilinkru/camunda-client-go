package camunda_client_go

import "io/ioutil"

// ProcessDefinition a client for ProcessDefinition
type ProcessDefinition struct {
	client *Client
}

// ResProcessDefinition a JSON object corresponding to the ProcessDefinition interface in the engine
type ResProcessDefinition struct {
	// The id of the process definition
	Id string `json:"id"`
	// The key of the process definition, i.e., the id of the BPMN 2.0 XML process definition
	Key string `json:"key"`
	// The category of the process definition
	Category string `json:"category"`
	// The description of the process definition
	Description string `json:"description"`
	// The name of the process definition
	Name string `json:"name"`
	// The version of the process definition that the engine assigned to it
	Version int `json:"Version"`
	// The file name of the process definition
	Resource string `json:"resource"`
	// The deployment id of the process definition
	DeploymentId string `json:"deploymentId"`
	// The file name of the process definition diagram, if it exists
	Diagram string `json:"diagram"`
	// A flag indicating whether the definition is suspended or not
	Suspended bool `json:"suspended"`
	// The tenant id of the process definition
	TenantId string `json:"tenantId"`
	// The version tag of the process definition
	VersionTag string `json:"versionTag"`
	// History time to live value of the process definition. Is used within History cleanup
	HistoryTimeToLive int `json:"historyTimeToLive"`
	// A flag indicating whether the process definition is startable in Tasklist or not
	StartableInTasklist bool `json:"startableInTasklist"`
}

// ResActivityInstanceStatistics a JSON array containing statistics results per activity
type ResActivityInstanceStatistics struct {
	// The id of the activity the results are aggregated for
	Id string `json:"id"`
	// The total number of running instances of this activity
	Instances int `json:"instances"`
	// Number	The total number of failed jobs for the running instances.
	// Note: Will be 0 (not null), if failed jobs were excluded
	FailedJobs int `json:"failedJobs"`
	// Each item in the resulting array is an object which contains the following properties
	Incidents []ResActivityInstanceStatisticsIncident `json:"incidents"`
}

// ResInstanceStatistics a JSON array containing statistics results per process definition
type ResInstanceStatistics struct {
	// The id of the activity the results are aggregated for
	Id string `json:"id"`
	// The total number of running instances of this activity
	Instances int `json:"instances"`
	// Number	The total number of failed jobs for the running instances.
	// Note: Will be 0 (not null), if failed jobs were excluded
	FailedJobs int `json:"failedJobs"`
	// The process definition with the properties as described in the get single definition method
	Definition ResProcessDefinition `json:"definition"`
	// Each item in the resulting array is an object which contains the following properties
	Incidents []ResActivityInstanceStatisticsIncident `json:"incidents"`
}

// ResActivityInstanceStatisticsIncident a statistics incident
type ResActivityInstanceStatisticsIncident struct {
	// The type of the incident the number of incidents is aggregated for.
	// See the User Guide for a list of incident types
	IncidentType string `json:"incidentType"`
	// The total number of incidents for the corresponding incident type
	IncidentCount int `json:"incidentCount"`
}

// ReqStartInstance a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqStartInstance struct {
	// A JSON object containing the variables the process is to be initialized with
	Variables *map[string]Variable `json:"variables,omitempty"`
	// The business key the process instance is to be initialized with.
	// The business key uniquely identifies the process instance in the context of the given process definition
	BusinessKey *string `json:"businessKey,omitempty"`
	// The case instance id the process instance is to be initialized with
	CaseInstanceId *string `json:"caseInstanceId,omitempty"`
	// Optional. A JSON array of instructions that specify which activities to start the process instance at.
	// If this property is omitted, the process instance starts at its default blank start event
	StartInstructions []ReqStartInstructions `json:"startInstructions,omitempty"`
	// Skip execution listener invocation for activities that are started or ended as part of this request
	// Note: This option is currently only respected when start instructions are submitted via
	// the startInstructions property
	SkipCustomListeners *bool `json:"skipCustomListeners,omitempty"`
	// Skip execution of input/output variable mappings for activities that are started or ended as part of this request
	// Note: This option is currently only respected when start instructions are submitted via
	// the startInstructions property
	SkipIoMappings *bool `json:"skipIoMappings,omitempty"`
	// Indicates if the variables, which was used by the process instance during execution, should be returned. Default value: false
	WithVariablesInReturn *bool `json:"withVariablesInReturn,omitempty"`
}

// ReqRestartInstance a request to restart instance
type ReqRestartInstance struct {
	// A list of process instance ids to restart
	ProcessInstanceIds *string `json:"processInstanceIds,omitempty"`
	// A historic process instance query like the request body described by POST /history/process-instance
	HistoricProcessInstanceQuery *string `json:"historicProcessInstanceQuery,omitempty"`
	// Optional. A JSON array of instructions that specify which activities to start the process instance at.
	// If this property is omitted, the process instance starts at its default blank start event
	StartInstructions []ReqStartInstructions `json:"startInstructions,omitempty"`
	// Skip execution listener invocation for activities that are started or ended as part of this request
	// Note: This option is currently only respected when start instructions are submitted via
	// the startInstructions property
	SkipCustomListeners *bool `json:"skipCustomListeners,omitempty"`
	// Skip execution of input/output variable mappings for activities that are started or ended as part of this request
	// Note: This option is currently only respected when start instructions are submitted via
	// the startInstructions property
	SkipIoMappings *bool `json:"skipIoMappings,omitempty"`
	// Set the initial set of variables during restart. By default, the last set of variables is used
	InitialVariables *bool `json:"initialVariables,omitempty"`
	// Do not take over the business key of the historic process instance.
	WithoutBusinessKey *bool `json:"withoutBusinessKey,omitempty"`
}

// ReqStartInstructions a JSON array of instructions that specify which activities to start the process instance at
type ReqStartInstructions struct {
	// Mandatory. One of the following values: startBeforeActivity, startAfterActivity, startTransition.
	// A startBeforeActivity instruction requests to start execution before entering a given activity.
	// A startAfterActivity instruction requests to start at the single outgoing sequence flow of a given activity.
	// A startTransition instruction requests to execute a specific sequence flow
	Type string `json:"type"`
	// Can be used with instructions of types startBeforeActivity and startAfterActivity.
	// Specifies the activity the instruction targets
	ActivityId *string `json:"activityId,omitempty"`
	// Can be used with instructions of types startTransition. Specifies the sequence flow to start
	TransitionId *string `json:"transitionId,omitempty"`
	// Can be used with instructions of type startBeforeActivity, startAfterActivity, and startTransition.
	// A JSON object containing variable key-value pairs
	Variables *map[string]VariableSet `json:"variables,omitempty"`
}

// QueryProcessDefinitionBy path builder
type QueryProcessDefinitionBy struct {
	Id       *string
	Key      *string
	TenantId *string
}

// ResGetStartFormKey a response from GetStartFormKey method
type ResGetStartFormKey struct {
	// The form key for the process definition
	Key string `json:"key"`
	// The context path of the process application
	ContextPath string `json:"contextPath"`
}

// ResBPMNProcessDefinition a JSON object containing the id of the definition and the BPMN 2.0 XML
type ResBPMNProcessDefinition struct {
	// The id of the process definition
	Id string `json:"id"`
	// An escaped XML string containing the XML that this definition was deployed with.
	// Carriage returns, line feeds and quotation marks are escaped
	Bpmn20Xml string `json:"bpmn20Xml"`
}

// String a build path part
func (q *QueryProcessDefinitionBy) String() string {
	if q.Key != nil && q.TenantId != nil {
		return "key/" + *q.Key + "/tenant-id/" + *q.TenantId
	} else if q.Key != nil {
		return "key/" + *q.Key
	}

	return *q.Id
}

// ResStartedProcessDefinition ProcessDefinition for started
type ResStartedProcessDefinition struct {
	// The id of the process definition
	Id string `json:"id"`
	// The id of the process definition
	DefinitionId string `json:"definitionId"`
	// The business key of the process instance
	BusinessKey string `json:"businessKey"`
	// The case instance id of the process instance
	CaseInstanceId string `json:"caseInstanceId"`
	// The tenant id of the process instance
	TenantId string `json:"tenantId"`
	// A flag indicating whether the instance is still running or not
	Ended bool `json:"ended"`
	// A flag indicating whether the instance is suspended or not
	Suspended bool `json:"suspended"`
	// A JSON array containing links to interact with the instance
	Links []ResLink `json:"links"`
	// A JSON object containing a property for each of the latest variables
	Variables map[string]Variable `json:"variables"`
}

// ReqSubmitStartForm request a SubmitStartForm
type ReqSubmitStartForm struct {
	// A JSON object containing the variables the process is to be initialized with.
	// Each key corresponds to a variable name and each value to a variable value
	Variables map[string]Variable `json:"variables"`
	// A JSON object containing the business key the process is to be initialized with.
	// The business key uniquely identifies the process instance in the context of the given process definition
	BusinessKey string `json:"businessKey"`
}

// ReqSubmitStartForm response rrom SubmitStartForm method
type ResSubmitStartForm struct {
	Links        []ResLink `json:"links"`
	Id           string    `json:"id"`
	DefinitionId string    `json:"definitionId"`
	BusinessKey  string    `json:"businessKey"`
	Ended        bool      `json:"ended"`
	Suspended    bool      `json:"suspended"`
}

// ReqActivateOrSuspendById response ActivateOrSuspendById
type ReqActivateOrSuspendById struct {
	// A Boolean value which indicates whether to activate or suspend a given process definition. When the value
	// is set to true, the given process definition will be suspended and when the value is set to false,
	// the given process definition will be activated
	Suspended *bool `json:"suspended,omitempty"`
	// A Boolean value which indicates whether to activate or suspend also all process instances of the given process
	// definition. When the value is set to true, all process instances of the provided process definition will be
	// activated or suspended and when the value is set to false, the suspension state of all process instances of
	// the provided process definition will not be updated
	IncludeProcessInstances *bool `json:"includeProcessInstances,omitempty"`
	// The date on which the given process definition will be activated or suspended. If null, the suspension state
	// of the given process definition is updated immediately. The date must have the format yyyy-MM-dd'T'HH:mm:ss,
	// e.g., 2013-01-23T14:42:45
	ExecutionDate *Time `json:"executionDate,omitempty"`
}

// ReqActivateOrSuspendByKey response ActivateOrSuspendByKey
type ReqActivateOrSuspendByKey struct {
	// The key of the process definitions to activate or suspend
	ProcessDefinitionKey string `json:"processDefinitionKey"`
	// A Boolean value which indicates whether to activate or suspend a given process definition. When the value
	// is set to true, the given process definition will be suspended and when the value is set to false,
	// the given process definition will be activated
	Suspended *bool `json:"suspended,omitempty"`
	// A Boolean value which indicates whether to activate or suspend also all process instances of the given process
	// definition. When the value is set to true, all process instances of the provided process definition will be
	// activated or suspended and when the value is set to false, the suspension state of all process instances of
	// the provided process definition will not be updated
	IncludeProcessInstances *bool `json:"includeProcessInstances,omitempty"`
	// The date on which the given process definition will be activated or suspended. If null, the suspension state
	// of the given process definition is updated immediately. The date must have the format yyyy-MM-dd'T'HH:mm:ss,
	// e.g., 2013-01-23T14:42:45
	ExecutionDate *Time `json:"executionDate,omitempty"`
}

// GetActivityInstanceStatistics retrieves runtime statistics of a given process definition, grouped by activities.
// These statistics include the number of running activity instances, optionally the number of failed jobs
// and also optionally the number of incidents either grouped by incident types or for a specific incident type.
// Note: This does not include historic data
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/get-activity-statistics/#query-parameters
func (p *ProcessDefinition) GetActivityInstanceStatistics(by QueryProcessDefinitionBy, query map[string]string) (statistic []*ResActivityInstanceStatistics, err error) {
	res, err := p.client.doGet("/process-definition/"+by.String()+"/statistics", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &statistic)
	return
}

// GetDiagram retrieves the diagram of a process definition.
// If the process definition’s deployment contains an image resource with the same file name as the process definition,
// the deployed image will be returned by the Get Diagram endpoint. Example: someProcess.bpmn and someProcess.png.
// Supported file extentions for the image are: svg, png, jpg, and gif
func (p *ProcessDefinition) GetDiagram(by QueryProcessDefinitionBy) (data []byte, err error) {
	res, err := p.client.doGet("/process-definition/"+by.String()+"/diagram", map[string]string{})
	if err != nil {
		return
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// GetStartFormVariables Retrieves the start form variables for a process definition
// (only if they are defined via the Generated Task Form approach). The start form variables take form data specified
// on the start event into account. If form fields are defined, the variable types and default values of the form
// fields are taken into account
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/get-form-variables/#query-parameters
func (p *ProcessDefinition) GetStartFormVariables(by QueryProcessDefinitionBy, query map[string]string) (variables map[string]Variable, err error) {
	res, err := p.client.doGet("/process-definition/"+by.String()+"/form-variables", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &variables)
	return
}

// GetListCount requests the number of process definitions that fulfill the query criteria.
// Takes the same filtering parameters as the Get Definitions method
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/get-query-count/#query-parameters
func (p *ProcessDefinition) GetListCount(query map[string]string) (count int, err error) {
	resCount := ResCount{}
	res, err := p.client.doGet("/process-definition/count", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetList queries for process definitions that fulfill given parameters.
// Parameters may be the properties of process definitions, such as the name, key or version.
// The size of the result set can be retrieved by using the Get Definition Count method
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/get-query/#query-parameters
func (p *ProcessDefinition) GetList(query map[string]string) (processDefinitions []*ResProcessDefinition, err error) {
	res, err := p.client.doGet("/process-definition", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &processDefinitions)
	return
}

// GetRenderedStartForm retrieves the rendered form for a process definition.
// This method can be used for getting the HTML rendering of a Generated Task Form
func (p *ProcessDefinition) GetRenderedStartForm(by QueryProcessDefinitionBy) (htmlForm string, err error) {
	res, err := p.client.doGet("/process-definition/"+by.String()+"/rendered-form", map[string]string{})
	if err != nil {
		return
	}

	defer res.Body.Close()
	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return string(rawData), nil
}

// GetStartFormKey retrieves the key of the start form for a process definition.
// The form key corresponds to the FormData#formKey property in the engine
func (p *ProcessDefinition) GetStartFormKey(by QueryProcessDefinitionBy) (resp *ResGetStartFormKey, err error) {
	resp = &ResGetStartFormKey{}
	res, err := p.client.doGet("/process-definition/"+by.String()+"/startForm", map[string]string{})
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &resp)
	return
}

// GetProcessInstanceStatistics retrieves runtime statistics of the process engine, grouped by process definitions.
// These statistics include the number of running process instances, optionally the number of failed jobs and also optionally the number of incidents either grouped by incident types or for a specific incident type.
// Note: This does not include historic data
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/get-statistics/#query-parameters
func (p *ProcessDefinition) GetProcessInstanceStatistics(query map[string]string) (statistic []*ResInstanceStatistics, err error) {
	res, err := p.client.doGet("/process-definition/statistics", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &statistic)
	return
}

// GetXML retrieves the BPMN 2.0 XML of a process definition
func (p *ProcessDefinition) GetXML(by QueryProcessDefinitionBy) (resp *ResBPMNProcessDefinition, err error) {
	resp = &ResBPMNProcessDefinition{}
	res, err := p.client.doGet("/process-definition/"+by.String()+"/xml", map[string]string{})
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &resp)
	return
}

// Get retrieves a process definition according to the ProcessDefinition interface in the engine
func (p *ProcessDefinition) Get(by QueryProcessDefinitionBy) (processDefinition *ResProcessDefinition, err error) {
	processDefinition = &ResProcessDefinition{}
	res, err := p.client.doGet("/process-definition/"+by.String(), map[string]string{})
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &processDefinition)
	return
}

// StartInstance instantiates a given process definition. Process variables and business key may be supplied
// in the request body
func (p *ProcessDefinition) StartInstance(by QueryProcessDefinitionBy, req ReqStartInstance) (processDefinition *ResStartedProcessDefinition, err error) {
	processDefinition = &ResStartedProcessDefinition{}
	res, err := p.client.doPostJson("/process-definition/"+by.String()+"/start", map[string]string{}, &req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, processDefinition)
	return
}

// SubmitStartForm starts a process instance using a set of process variables and the business key.
// If the start event has Form Field Metadata defined, the process engine will perform backend validation for any form
// fields which have validators defined. See Documentation on Generated Task Forms
func (p *ProcessDefinition) SubmitStartForm(by QueryProcessDefinitionBy, req ReqSubmitStartForm) (reps *ResSubmitStartForm, err error) {
	reps = &ResSubmitStartForm{}
	res, err := p.client.doPostJson("/process-definition/"+by.String()+"/submit-form", map[string]string{}, &req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, reps)
	return
}

// ActivateOrSuspendById activates or suspends a given process definition by id or by latest version
// of process definition key
func (p *ProcessDefinition) ActivateOrSuspendById(by QueryProcessDefinitionBy, req ReqActivateOrSuspendById) error {
	return p.client.doPutJson("/process-definition/"+by.String()+"/suspended", map[string]string{}, &req)
}

// ActivateOrSuspendByKey activates or suspends process definitions with the given process definition key
func (p *ProcessDefinition) ActivateOrSuspendByKey(req ReqActivateOrSuspendByKey) error {
	return p.client.doPutJson("/process-definition/suspended", map[string]string{}, &req)
}

// UpdateHistoryTimeToLive updates history time to live for process definition.
// The field is used within History cleanup
func (p *ProcessDefinition) UpdateHistoryTimeToLive(by QueryProcessDefinitionBy, historyTimeToLive int) error {
	return p.client.doPutJson("/process-definition/"+by.String()+"/history-time-to-live", map[string]string{}, &map[string]int{"historyTimeToLive": historyTimeToLive})
}

// Delete deletes a process definition from a deployment by id
// https://docs.camunda.org/manual/latest/reference/rest/process-definition/delete-process-definition/#query-parameters
func (p *ProcessDefinition) Delete(by QueryProcessDefinitionBy, query map[string]string) error {
	err := p.client.doDelete("/process-definition/"+by.String(), query)
	return err
}

// GetDeployedStartForm retrieves the deployed form that can be referenced from a start event. For further information please refer to User Guide
func (p *ProcessDefinition) GetDeployedStartForm(by QueryProcessDefinitionBy) (htmlForm string, err error) {
	res, err := p.client.doGet("/process-definition/"+by.String()+"/deployed-start-form", map[string]string{})
	if err != nil {
		return
	}

	defer res.Body.Close()
	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return string(rawData), nil
}

// RestartProcessInstance restarts process instances that were canceled or terminated synchronously.
// To execute the restart asynchronously, use the Restart Process Instance Async method
// For more information about the difference between synchronous and asynchronous execution,
// please refer to the related section of the user guide
func (p *ProcessDefinition) RestartProcessInstance(id string, req ReqRestartInstance) error {
	_, err := p.client.doPostJson("/process-definition/"+id+"/restart", map[string]string{}, &req)
	return err
}

// RestartProcessInstanceAsync restarts process instances that were canceled or terminated asynchronously.
// To execute the restart synchronously, use the Restart Process Instance method
// For more information about the difference between synchronous and asynchronous execution,
// please refer to the related section of the user guide
func (p *ProcessDefinition) RestartProcessInstanceAsync(id string, req ReqRestartInstance) (resp *ResBatch, err error) {
	resp = &ResBatch{}
	res, err := p.client.doPostJson("/process-definition/"+id+"/restart-async", map[string]string{}, &req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, resp)
	return
}
