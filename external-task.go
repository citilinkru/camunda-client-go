package camunda_client_go

import (
	"fmt"
)

// ExternalTask a client for ExternalTask API
type ExternalTask struct {
	client *Client
}

// ResExternalTask a ExternalTask response
type ResExternalTask struct {
	// The id of the activity that this external task belongs to
	ActivityId string `json:"activityId"`
	// The id of the activity instance that the external task belongs to
	ActivityInstanceId string `json:"activityInstanceId"`
	// The full error message submitted with the latest reported failure executing this task;
	// null if no failure was reported previously or if no error message was submitted
	ErrorMessage string `json:"errorMessage"`
	// The error details submitted with the latest reported failure executing this task.
	// null if no failure was reported previously or if no error details was submitted
	ErrorDetails string `json:"errorDetails"`
	// The id of the execution that the external task belongs to
	ExecutionId string `json:"executionId"`
	// The id of the external task
	Id string `json:"id"`
	// The date that the task's most recent lock expires or has expired
	LockExpirationTime string `json:"lockExpirationTime"`
	// The id of the process definition the external task is defined in
	ProcessDefinitionId string `json:"processDefinitionId"`
	// The key of the process definition the external task is defined in
	ProcessDefinitionKey string `json:"processDefinitionKey"`
	// The id of the process instance the external task belongs to
	ProcessInstanceId string `json:"processInstanceId"`
	// The id of the tenant the external task belongs to
	TenantId string `json:"tenantId"`
	// The number of retries the task currently has left
	Retries int `json:"retries"`
	// A flag indicating whether the external task is suspended or not
	Suspended bool `json:"suspended"`
	// The id of the worker that possesses or possessed the most recent lock
	WorkerId string `json:"workerId"`
	// The priority of the external task
	Priority int `json:"priority"`
	// The topic name of the external task
	TopicName string `json:"topicName"`
	// The business key of the process instance the external task belongs to
	BusinessKey string `json:"businessKey"`
}

// QueryGetListPost a query for ListPost request
type QueryGetListPost struct {
	// Filter by an external task's id
	ExternalTaskId *string `json:"externalTaskId,omitempty"`
	// Filter by an external task topic
	TopicName *string `json:"topicName,omitempty"`
	// Filter by the id of the worker that the task was most recently locked by
	WorkerId *string `json:"workerId,omitempty"`
	// Only include external tasks that are currently locked (i.e., they have a lock time and it has not expired).
	// Value may only be true, as false matches any external task
	Locked *bool `json:"locked,omitempty"`
	// Only include external tasks that are currently not locked (i.e., they have no lock or it has expired).
	// Value may only be true, as false matches any external task
	NotLocked *bool `json:"notLocked,omitempty"`
	// Only include external tasks that have a positive (> 0) number of retries (or null). Value may only be true,
	// as false matches any external task
	WithRetriesLeft *bool `json:"withRetriesLeft,omitempty"`
	// Only include external tasks that have 0 retries. Value may only be true, as false matches any external task
	NoRetriesLeft *bool `json:"noRetriesLeft,omitempty"`
	// Restrict to external tasks that have a lock that expires after a given date. By default*,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200
	LockExpirationAfter *Time `json:"lockExpirationAfter,omitempty"`
	// Restrict to external tasks that have a lock that expires before a given date. By default*,
	// the date must have the format yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g., 2013-01-23T14:42:45.000+0200
	LockExpirationBefore *Time `json:"lockExpirationBefore,omitempty"`
	// Filter by the id of the activity that an external task is created for
	ActivityId *string `json:"activityId,omitempty"`
	// Filter by the comma-separated list of ids of the activities that an external task is created for
	ActivityIdIn *string `json:"activityIdIn,omitempty"`
	// Filter by the id of the execution that an external task belongs to
	ExecutionId *string `json:"executionId,omitempty"`
	// Filter by the id of the process instance that an external task belongs to
	ProcessInstanceId *string `json:"processInstanceId,omitempty"`
	// Filter by the id of the process definition that an external task belongs to
	ProcessDefinitionId *string `json:"processDefinitionId,omitempty"`
	// Filter by a comma-separated list of tenant ids. An external task must have one of the given tenant ids
	TenantIdIn *string `json:"tenantIdIn,omitempty"`
	// Only include active tasks. Value may only be true, as false matches any external task
	Active *bool `json:"active,omitempty"`
	// Only include suspended tasks. Value may only be true, as false matches any external task
	Suspended *bool `json:"suspended,omitempty"`
	// Only include jobs with a priority higher than or equal to the given value. Value must be a valid long value
	PriorityHigherThanOrEquals *int `json:"priorityHigherThanOrEquals,omitempty"`
	// Only include jobs with a priority lower than or equal to the given value. Value must be a valid long value
	PriorityLowerThanOrEquals *int `json:"priorityLowerThanOrEquals,omitempty"`
	// A JSON array of criteria to sort the result by. Each element of the array is a JSON object
	// that specifies one ordering. The position in the array identifies the rank of an ordering,
	// i.e., whether it is primary, secondary, etc.
	Sorting *QueryListPostSorting `json:"sorting,omitempty"`
}

// QueryFetchAndLock query for FetchAndLock request
type QueryFetchAndLock struct {
	// Mandatory. The id of the worker on which behalf tasks are fetched. The returned tasks are locked
	// for that worker and can only be completed when providing the same worker id
	WorkerId string `json:"workerId"`
	// Mandatory. The maximum number of tasks to return
	MaxTasks int `json:"maxTasks"`
	// A boolean value, which indicates whether the task should be fetched based on its priority or arbitrarily
	UsePriority *bool `json:"usePriority,omitempty"`
	// The Long Polling timeout in milliseconds.
	// Note: The value cannot be set larger than 1.800.000 milliseconds (corresponds to 30 minutes)
	AsyncResponseTimeout *int `json:"asyncResponseTimeout,omitempty"`
	// A JSON array of topic objects for which external tasks should be fetched.
	// The returned tasks may be arbitrarily distributed among these topics
	Topics *[]QueryFetchAndLockTopic `json:"topics,omitempty"`
}

// QueryFetchAndLockTopic a JSON array of topic objects for which external tasks should be fetched
type QueryFetchAndLockTopic struct {
	// Mandatory. The topic's name
	TopicName string `json:"topicName"`
	// Mandatory. The duration to lock the external tasks for in milliseconds
	LockDuration int `json:"lockDuration"`
	// A JSON array of String values that represent variable names. For each result task belonging to this topic,
	// the given variables are returned as well if they are accessible from the external task's execution.
	// If not provided - all variables will be fetched
	Variables *[]string `json:"variables,omitempty"`
	// If true only local variables will be fetched
	LocalVariables *bool `json:"localVariables,omitempty"`
	// A String value which enables the filtering of tasks based on process instance business key
	BusinessKey *string `json:"businessKey,omitempty"`
	// Filter tasks based on process definition id
	ProcessDefinitionId *string `json:"processDefinitionId,omitempty"`
	// Filter tasks based on process definition ids
	ProcessDefinitionIdIn *string `json:"processDefinitionIdIn,omitempty"`
	// Filter tasks based on process definition key
	ProcessDefinitionKey *string `json:"processDefinitionKey,omitempty"`
	// Filter tasks based on process definition keys
	ProcessDefinitionKeyIn *string `json:"processDefinitionKeyIn,omitempty"`
	// 	Filter tasks without tenant id
	WithoutTenantId *string `json:"withoutTenantId,omitempty"`
	// Filter tasks based on tenant ids
	TenantIdIn *string `json:"tenantIdIn,omitempty"`
	// A JSON object used for filtering tasks based on process instance variable values.
	// A property name of the object represents a process variable name, while the property value
	// represents the process variable value to filter tasks by
	ProcessVariables map[string]Variable `json:"processVariables,omitempty"`
	// Determines whether serializable variable values (typically variables that store custom Java objects)
	// should be deserialized on server side (default false).
	// If set to true, a serializable variable will be deserialized on server side and transformed to JSON
	// using Jackson's POJO/bean property introspection feature. Note that this requires the Java classes
	// of the variable value to be on the REST API's classpath.
	// If set to false, a serializable variable will be returned in its serialized format.
	// For example, a variable that is serialized as XML will be returned as a JSON string containing XML
	DeserializeValues *bool `json:"deserializeValues,omitempty"`
}

// QueryListPostSorting a JSON array of criteria to sort the result by
type QueryListPostSorting struct {
	// Mandatory. Sort the results lexicographically by a given criterion. Valid values are id, lockExpirationTime,
	// processInstanceId, processDefinitionId, processDefinitionKey, taskPriority and tenantId
	SortBy string `json:"sortBy"`
	// Mandatory. Sort the results in a given order. Values may be asc for ascending order or desc for descending order
	SortOrder string `json:"sortOrder"`
}

// ResLockedExternalTask a response FetchAndLock method
type ResLockedExternalTask struct {
	// The id of the activity that this external task belongs to
	ActivityId string `json:"activityId"`
	// The id of the activity instance that the external task belongs to
	ActivityInstanceId string `json:"activityInstanceId"`
	// The full error message submitted with the latest reported failure executing this task;
	// null if no failure was reported previously or if no error message was submitted
	ErrorMessage string `json:"errorMessage"`
	// The error details submitted with the latest reported failure executing this task.
	// null if no failure was reported previously or if no error details was submitted
	ErrorDetails string `json:"errorDetails"`
	// The id of the execution that the external task belongs to
	ExecutionId string `json:"executionId"`
	// The id of the external task
	Id string `json:"id"`
	// The date that the task's most recent lock expires or has expired
	LockExpirationTime string `json:"lockExpirationTime"`
	// The id of the process definition the external task is defined in
	ProcessDefinitionId string `json:"processDefinitionId"`
	// The key of the process definition the external task is defined in
	ProcessDefinitionKey string `json:"processDefinitionKey"`
	// The id of the process instance the external task belongs to
	ProcessInstanceId string `json:"processInstanceId"`
	// The id of the tenant the external task belongs to
	TenantId string `json:"tenantId"`
	// The number of retries the task currently has left
	Retries int `json:"retries"`
	// The id of the worker that possesses or possessed the most recent lock
	WorkerId string `json:"workerId"`
	// The priority of the external task
	Priority int `json:"priority"`
	// The topic name of the external task
	TopicName string `json:"topicName"`
	// The business key of the process instance the external task belongs to
	BusinessKey string `json:"businessKey"`
	// A JSON object containing a property for each of the requested variables
	Variables map[string]Variable `json:"variables"`
}

// Variable a variable
type Variable struct {
	// The variable's value
	Value interface{} `json:"value"`
	// The value type of the variable.
	Type string `json:"type"`
	// A JSON object containing additional, value-type-dependent properties
	ValueInfo ValueInfo `json:"valueInfo"`
}

// VariableSet a variable for set
type VariableSet struct {
	// The variable's value
	Value string `json:"value"`
	// The value type of the variable.
	Type string `json:"type"`
	// A JSON object containing additional, value-type-dependent properties
	ValueInfo ValueInfo `json:"valueInfo"`
	// Indicates whether the variable should be a local variable or not. If set to true, the variable becomes a local
	// variable of the execution entering the target activity
	Local bool `json:"local"`
}

// ValueInfo a value info in variable
type ValueInfo struct {
	// A string representation of the object's type name
	ObjectTypeName *string `json:"objectTypeName"`
	// The serialization format used to store the variable.
	SerializationDataFormat *string `json:"serializationDataFormat"`
}

// QueryComplete a query for Complete request
type QueryComplete struct {
	// The id of the worker that completes the task.
	// Must match the id of the worker who has most recently locked the task
	WorkerId *string `json:"workerId,omitempty"`
	// A JSON object containing variable key-value pairs
	Variables *map[string]Variable `json:"variables"`
	// A JSON object containing variable key-value pairs.
	// Local variables are set only in the scope of external task
	LocalVariables *map[string]Variable `json:"localVariables"`
}

// QueryHandleBPMNError a query for HandleBPMNError request
type QueryHandleBPMNError struct {
	// The id of the worker that reports the failure.
	// Must match the id of the worker who has most recently locked the task
	WorkerId *string `json:"workerId,omitempty"`
	// An error code that indicates the predefined error. Is used to identify the BPMN error handler
	ErrorCode *string `json:"errorCode,omitempty"`
	// An error message that describes the error
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A JSON object containing the variables which will be passed to the execution.
	// Each key corresponds to a variable name and each value to a variable value
	Variables *map[string]Variable `json:"variables"`
}

// QueryHandleFailure a query for HandleFailure request
type QueryHandleFailure struct {
	// The id of the worker that reports the failure.
	// Must match the id of the worker who has most recently locked the task
	WorkerId *string `json:"workerId,omitempty"`
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

// QueryExtendLock a query for ExtendLock request
type QueryExtendLock struct {
	//	An amount of time (in milliseconds). This is the new lock duration starting from the current moment
	NewDuration *int `json:"newDuration,omitempty"`
	// The ID of a worker who is locking the external task
	WorkerId *string `json:"workerId,omitempty"`
}

// QuerySetRetriesAsync a query for SetRetriesAsync request
type QuerySetRetriesAsync struct {
	// The number of retries to set for the external task. Must be >= 0. If this is 0, an incident is created and the
	// task cannot be fetched anymore unless the retries are increased again. Can not be null
	Retries int `json:"retries"`
	// The ids of the external tasks to set the number of retries for
	ExternalTaskIds *string `json:"externalTaskIds,omitempty"`
	// The ids of process instances containing the tasks to set the number of retries for
	ProcessInstanceIds *string `json:"processInstanceIds,omitempty"`
	// Query for the external tasks to set the number of retries for
	ExternalTaskQuery *string `json:"externalTaskQuery,omitempty"`
	// Query for the process instances containing the tasks to set the number of retries for
	ProcessInstanceQuery *string `json:"processInstanceQuery,omitempty"`
	// Query for the historic process instances containing the tasks to set the number of retries for
	HistoricProcessInstanceQuery *string `json:"historicProcessInstanceQuery,omitempty"`
}

// ResBatch a JSON object corresponding to the Batch interface in the engine
type ResBatch struct {
	// The id of the created batch
	Id string `json:"id"`
	// The type of the created batch
	Type string `json:"type"`
	// The total jobs of a batch is the number of batch execution jobs required to complete the batch
	TotalJobs int `json:"totalJobs"`
	// The number of batch execution jobs created per seed job invocation. The batch seed job is invoked until
	// it has created all batch execution jobs required by the batch (see totalJobs property)
	BatchJobsPerSeed int `json:"batchJobsPerSeed"`
	// Every batch execution job invokes the command executed by the batch invocationsPerBatchJob times.
	// E.g., for a process instance migration batch this specifies the number of process instances which are
	// migrated per batch execution job
	InvocationsPerBatchJob int `json:"invocationsPerBatchJob"`
	// The job definition id for the seed jobs of this batch
	SeedJobDefinitionId string `json:"seedJobDefinitionId"`
	// The job definition id for the monitor jobs of this batch
	MonitorJobDefinitionId string `json:"monitorJobDefinitionId"`
	// The job definition id for the batch execution jobs of this batch
	BatchJobDefinitionId string `json:"batchJobDefinitionId"`
	// The tenant id of the batch
	TenantId string `json:"tenantId"`
}

// QuerySetRetriesSync a query for SetRetriesSync request
type QuerySetRetriesSync struct {
	// The number of retries to set for the external task. Must be >= 0. If this is 0, an incident is created
	// and the task cannot be fetched anymore unless the retries are increased again. Can not be null
	Retries int `json:"retries"`
	// The ids of the external tasks to set the number of retries for
	ExternalTaskIds *string `json:"externalTaskIds,omitempty"`
	// The ids of process instances containing the tasks to set the number of retries for
	ProcessInstanceIds *string `json:"processInstanceIds,omitempty"`
	// Query for the external tasks to set the number of retries for
	ExternalTaskQuery *string `json:"externalTaskQuery,omitempty"`
	// Query for the process instances containing the tasks to set the number of retries for
	ProcessInstanceQuery *string `json:"processInstanceQuery,omitempty"`
	// Query for the historic process instances containing the tasks to set the number of retries for
	HistoricProcessInstanceQuery *string `json:"historicProcessInstanceQuery,omitempty"`
}

// Get retrieves an external task by id, corresponding to the ExternalTask interface in the engine
func (e *ExternalTask) Get(id string) (*ResExternalTask, error) {
	resp := &ResExternalTask{}
	res, err := e.client.doGet(
		"/external-task/"+id,
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	if err := e.client.readJsonResponse(res, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetList queries for the external tasks that fulfill given parameters.
// Parameters may be static as well as dynamic runtime properties of executions
// Query parameters described in the documentation:
// https://docs.camunda.org/manual/latest/reference/rest/external-task/get-query/#query-parameters
func (e *ExternalTask) GetList(query map[string]string) ([]*ResExternalTask, error) {
	resp := []*ResExternalTask{}
	res, err := e.client.doGet(
		"/external-task",
		query,
	)
	if err != nil {
		return nil, err
	}

	if err := e.client.readJsonResponse(res, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetListCount queries for the number of external tasks that fulfill given parameters.
// Takes the same parameters as the Get External Tasks method.
// Query parameters described in the documentation:
// https://docs.camunda.org/manual/latest/reference/rest/external-task/get-query-count/#query-parameters
func (e *ExternalTask) GetListCount(query map[string]string) (int, error) {
	resCount := ResCount{}
	res, err := e.client.doGet("/external-task/count", query)
	if err != nil {
		return 0, err
	}

	err = e.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// GetListPost queries for external tasks that fulfill given parameters in the form of a JSON object.
// This method is slightly more powerful than the Get External Tasks method
// because it allows to specify a hierarchical result sorting.
func (e *ExternalTask) GetListPost(query QueryGetListPost, firstResult, maxResults int) ([]*ResExternalTask, error) {
	resp := []*ResExternalTask{}
	res, err := e.client.doPostJson(
		"/external-task",
		map[string]string{},
		&query,
	)
	if err != nil {
		return nil, err
	}

	if err := e.client.readJsonResponse(res, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetListPostCount queries for the number of external tasks that fulfill given parameters.
// This method takes the same message body as the Get External Tasks (POST) method
func (e *ExternalTask) GetListPostCount(query QueryGetListPost) (int, error) {
	resCount := ResCount{}
	res, err := e.client.doPostJson(
		"/external-task/count",
		map[string]string{},
		query,
	)
	if err != nil {
		return 0, err
	}

	err = e.client.readJsonResponse(res, resCount)
	return resCount.Count, err
}

// FetchAndLock fetches and locks a specific number of external tasks for execution by a worker.
// Query can be restricted to specific task topics and for each task topic an individual lock time can be provided
func (e *ExternalTask) FetchAndLock(query QueryFetchAndLock) ([]*ResLockedExternalTask, error) {
	resp := []*ResLockedExternalTask{}
	res, err := e.client.doPostJson(
		"/external-task/fetchAndLock",
		map[string]string{},
		&query,
	)
	if err != nil {
		return nil, fmt.Errorf("requsest error: %s", err)
	}

	if err := e.client.readJsonResponse(res, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Complete a completes an external task by id and updates process variables
func (e *ExternalTask) Complete(id string, query QueryComplete) error {
	_, err := e.client.doPostJson("/external-task/"+id+"/complete", map[string]string{}, &query)
	return err
}

// HandleBPMNError reports a business error in the context of a running external task by id.
// The error code must be specified to identify the BPMN error handler
func (e *ExternalTask) HandleBPMNError(id string, query QueryHandleBPMNError) error {
	_, err := e.client.doPostJson("/external-task/"+id+"/bpmnError", map[string]string{}, &query)
	return err
}

// HandleFailure reports a failure to execute an external task by id.
// A number of retries and a timeout until the task can be retried can be specified.
// If retries are set to 0, an incident for this task is created
func (e *ExternalTask) HandleFailure(id string, query QueryHandleFailure) error {
	_, err := e.client.doPostJson("/external-task/"+id+"/failure", map[string]string{}, &query)
	return err
}

// Unlock a unlocks an external task by id. Clears the task’s lock expiration time and worker id
func (e *ExternalTask) Unlock(id string) error {
	_, err := e.client.doPost("/external-task/"+id+"/unlock", map[string]string{})
	return err
}

// ExtendLock a extends the timeout of the lock by a given amount of time
func (e *ExternalTask) ExtendLock(id string, query QueryExtendLock) error {
	_, err := e.client.doPostJson("/external-task/"+id+"/extendLock", map[string]string{}, &query)
	return err
}

// SetPriority a sets the priority of an existing external task by id. The default value of a priority is 0
func (e *ExternalTask) SetPriority(id string, priority int) error {
	_, err := e.client.doPut("/external-task/"+id+"/priority", map[string]string{})
	return err
}

// SetRetries a sets the number of retries left to execute an external task by id. If retries are set to 0,
// an incident is created
func (e *ExternalTask) SetRetries(id string, retries int) error {
	_, err := e.client.doPutJson("/external-task/"+id+"/retries", map[string]string{}, map[string]int{
		"retries": retries,
	})
	return err
}

// SetRetriesAsync a set Retries For Multiple External Tasks Async (Batch).
// Sets the number of retries left to execute external tasks by id asynchronously.
// If retries are set to 0, an incident is created
func (e *ExternalTask) SetRetriesAsync(id string, query QuerySetRetriesAsync) (*ResBatch, error) {
	resp := ResBatch{}
	res, err := e.client.doPostJson(
		"/external-task/retries-async",
		map[string]string{},
		&query,
	)
	if err != nil {
		return nil, err
	}

	if err := e.client.readJsonResponse(res, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetRetriesSync a set Retries For Multiple External Tasks Sync.
// Sets the number of retries left to execute external tasks by id synchronously.
// If retries are set to 0, an incident is created
func (e *ExternalTask) SetRetriesSync(id string, query QuerySetRetriesSync) error {
	_, err := e.client.doPutJson("/external-task/retries", map[string]string{}, &query)
	return err
}
