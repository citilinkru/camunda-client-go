package camunda_client_go

import (
	"encoding/json"
	"fmt"
	"time"
)

// IdentityLink camunda IdentityLink, e.g. assignee, candidate groups, of UserTask
type IdentityLink struct {
	// The user id of the assignee.
	UserId string `json:"userId"`
	// The group id of the candidate
	GroupId string `json:"groupId"`
	// the type of the indentity, either candidate or assign
	Type string `json:"type"`
}

// userTaskApi a client for userTaskApi API
type userTaskApi struct {
	client *Client
}

// UserTaskResponse get task response
type UserTaskResponse struct {
	// The id of the task.
	Id string `json:"id"`
	// The tasks name.
	Name string `json:"name"`
	// The user assigned to this task.
	Assignee string `json:"assignee"`
	// The time the task was created.Format yyyy-MM-dd'T'HH:mm:ss.
	Created string `json:"created"`
	// The due date for the task.Format yyyy-MM-dd'T'HH:mm:ss.
	Due string `json:"due"`
	// The follow-up date for the task.Format yyyy-MM-dd'T'HH:mm:ss.
	FollowUp string `json:"followUp"`
	// The delegation state of the task.Corresponds to the DelegationState enum in the engine.Possible values are RESOLVED and PENDING.
	DelegationState string `json:"delegationState"`
	// The task description.
	Description string `json:"description"`
	// The id of the execution the task belongs to.
	ExecutionId string `json:"executionId"`
	// The owner of the task.
	Owner string `json:"owner"`
	// The id of the parent task, if this task is a subtask.
	ParentTaskId string `json:"parentTaskId"`
	// The priority of the task.
	Priority int64 `json:"priority"`
	// The id of the process definition this task belongs to.
	ProcessDefinitionId string `json:"processDefinitionId"`
	// The id of the process instance this task belongs to.
	ProcessInstanceId string `json:"processInstanceId"`
	// The id of the case execution the task belongs to.
	CaseExecutionId string `json:"caseExecutionId"`
	// The id of the case definition the task belongs to.
	CaseDefinitionId string `json:"caseDefinitionId"`
	// The id of the case instance the task belongs to.
	CaseInstanceId string `json:"caseInstanceId"`
	// The task definition key.
	TaskDefinitionKey string `json:"taskDefinitionKey"`
	// Whether the task belongs to a process instance that is suspended.
	Suspended bool `json:"suspended"`
	// If not null, the form key for the task.
	FormKey *string `json:"formKey"`
	// If not null, the tenantId for the task.
	TenantId *string `json:"tenantId"`
}

// UserTask camunda user task
type UserTask struct {
	*UserTaskResponse

	api *userTaskApi

	IdentityLinks []*IdentityLink
}

// Complete complete user task
func (t *UserTask) Complete(query QueryUserTaskComplete) error {
	err := t.api.Complete(t.Id, query)
	if err != nil {
		return fmt.Errorf("can't complete task: %w", err)
	}

	return nil
}

// GetIdentityLinks retrieve IdentityLinks of the UserTask
func (t *UserTask) GetIdentityLinks() (*[]IdentityLink, error) {
	links, err := t.api.GetIdentityLinks(t.Id)
	if err != nil {
		return nil, fmt.Errorf("can't get identity links: %w", err)
	}

	return links, nil
}

// AddIdentityLink add IdentityLink to UserTask
func (t *UserTask) AddIdentityLink(query ReqIdentityLink) error {
	err := t.api.AddIdentityLink(t.Id, query)
	if err != nil {
		return fmt.Errorf("can't add identity link: %w", err)
	}

	return nil
}

// DeleteIdentityLink delete IdentityLink from UserTask
func (t *UserTask) DeleteIdentityLink(query ReqIdentityLink) error {
	err := t.api.DeleteIdentityLink(t.Id, query)
	if err != nil {
		return fmt.Errorf("can't delete identity link: %w", err)
	}

	return nil
}

// delegationState task delegation state
type delegationState string

const (
	DelegationStatePending  = "PENDING"
	DelegationStateResolved = "RESOLVED"
)

// variableFilterExpressionOperator operator for variable filter expression
type variableFilterExpressionOperator string

const (
	VariableFilterExpressionOperatorEqual              = "eq"
	VariableFilterExpressionOperatorNotEqual           = "neq"
	VariableFilterExpressionOperatorGreaterThan        = "gt"
	VariableFilterExpressionOperatorGreaterThanOrEqual = "gteq"
	VariableFilterExpressionOperatorLessThan           = "lt"
	VariableFilterExpressionOperatorLessThanOrEqual    = "lteq"
	VariableFilterExpressionOperatorLike               = "like"
)

// ReqIdentityLink post request for IdentityLink
type ReqIdentityLink struct {
	UserId  string `json:"userId,omitempty"`
	GroupId string `json:"groupId,omitempty"`
	Type    string `json:"type,omitempty"`
}

// VariableFilterExpression filter expression
type VariableFilterExpression struct {
	Name     string                           `json:"name"`
	Operator variableFilterExpressionOperator `json:"operator"`
	Value    string                           `json:"value"`
}

// UserTaskGetListQuery query for GetList,
type UserTaskGetListQuery struct {
	// Restrict to tasks that belong to process instances with the given id.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`
	// Restrict to tasks that belong to process instances with the given business key.
	ProcessInstanceBusinessKey string `json:"processInstanceBusinessKey,omitempty"`
	// Restrict to tasks that belong to process instances with one of the give business keys. The keys need to be in a comma-separated list.
	ProcessInstanceBusinessKeyIn []string `json:"processInstanceBusinessKeyIn,omitempty"`
	// Restrict to tasks that have a process instance business key that has the parameter value as a substring.
	ProcessInstanceBusinessKeyLike string `json:"processInstanceBusinessKeyLike,omitempty"`
	// Restrict to tasks that belong to a process definition with the given id.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
	// Restrict to tasks that belong to a process definition with the given key.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`
	// Restrict to tasks that belong to a process definition with one of the given keys. The keys need to be in a comma-separated list.
	ProcessDefinitionKeyIn []string `json:"processDefinitionKeyIn,omitempty"`
	// Restrict to tasks that belong to a process definition with the given name.
	ProcessDefinitionName string `json:"processDefinitionName,omitempty"`
	// Restrict to tasks that have a process definition name that has the parameter value as a substring.
	ProcessDefinitionNameLike string `json:"processDefinitionNameLike,omitempty"`
	// Restrict to tasks that belong to an execution with the given id.
	ExecutionId string `json:"executionId,omitempty"`
	// Restrict to tasks that belong to case instances with the given id.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`
	// Restrict to tasks that belong to case instances with the given business key.
	CaseInstanceBusinessKey string `json:"caseInstanceBusinessKey,omitempty"`
	// Restrict to tasks that have a case instance business key that has the parameter value as a substring.
	CaseInstanceBusinessKeyLike string `json:"caseInstanceBusinessKeyLike,omitempty"`
	// Restrict to tasks that belong to a case definition with the given id.
	CaseDefinitionId string `json:"caseDefinitionId,omitempty"`
	// Restrict to tasks that belong to a case definition with the given key.
	CaseDefinitionKey string `json:"caseDefinitionKey,omitempty"`
	// Restrict to tasks that belong to a case definition with the given name.
	CaseDefinitionName string `json:"caseDefinitionName,omitempty"`
	// Restrict to tasks that have a case definition name that has the parameter value as a substring.
	CaseDefinitionNameLike string `json:"caseDefinitionNameLike,omitempty"`
	// Restrict to tasks that belong to a case execution with the given id.
	CaseExecutionId string `json:"caseExecutionId,omitempty"`
	// Only include tasks which belong to one of the passed and comma-separated activity instance ids.
	ActivityInstanceIdIn []string `json:"activityInstanceIdIn,omitempty"`
	// Only include tasks which belong to one of the passed and comma-separated tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
	// Only include tasks which belong to no tenant. Value may only be true, as false is the default behavior.
	WithoutTenantId string `json:"withoutTenantId,omitempty"`
	// Restrict to tasks that the given user is assigned to.
	Assignee string `json:"assignee,omitempty"`
	// Restrict to tasks that the user described by the given expression is assigned to. See the user guide for more information on available functions.
	AssigneeExpression string `json:"assigneeExpression,omitempty"`
	// Restrict to tasks that have an assignee that has the parameter value as a substring.
	AssigneeLike string `json:"assigneeLike,omitempty"`
	// Restrict to tasks that have an assignee that has the parameter value described by the given expression as a substring. See the user guide for more information on available functions.
	AssigneeLikeExpression string `json:"assigneeLikeExpression,omitempty"`
	// Restrict to tasks that the given user owns.
	Owner string `json:"owner,omitempty"`
	// Restrict to tasks that the user described by the given expression owns. See the user guide for more information on available functions.
	OwnerExpression string `json:"ownerExpression,omitempty"`
	// Only include tasks that are offered to the given group.
	CandidateGroup string `json:"candidateGroup,omitempty"`
	// Only include tasks that are offered to the group described by the given expression. See the user guide for more information on available functions.
	CandidateGroupExpression string `json:"candidateGroupExpression,omitempty"`
	// Only include tasks that are offered to the given user or to one of his groups.
	CandidateUser string `json:"candidateUser,omitempty"`
	// Only include tasks that are offered to the user described by the given expression. See the user guide for more information on available functions.
	CandidateUserExpression string `json:"candidateUserExpression,omitempty"`
	// Also include tasks that are assigned to users in candidate queries. Default is to only include tasks that are not assigned to any user if you query by candidate user or group(s).
	IncludeAssignedTasks bool `json:"includeAssignedTasks,omitempty"`
	// Only include tasks that the given user is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee).
	InvolvedUser string `json:"involvedUser,omitempty"`
	// Only include tasks that the user described by the given expression is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee). See the user guide for more information on available functions.
	InvolvedUserExpression string `json:"involvedUserExpression,omitempty"`
	// If set to true, restricts the query to all tasks that are assigned.
	Assigned bool `json:"assigned,omitempty"`
	// If set to true, restricts the query to all tasks that are unassigned.
	Unassigned bool `json:"unassigned,omitempty"`
	// Restrict to tasks that have the given key.
	TaskDefinitionKey string `json:"taskDefinitionKey,omitempty"`
	// Restrict to tasks that have one of the given keys. The keys need to be in a comma-separated list.
	TaskDefinitionKeyIn []string `json:"taskDefinitionKeyIn,omitempty"`
	// Restrict to tasks that have a key that has the parameter value as a substring.
	TaskDefinitionKeyLike string `json:"taskDefinitionKeyLike,omitempty"`
	// Restrict to tasks that have the given name.
	Name string `json:"name,omitempty"`
	// Restrict to tasks that do not have the given name.
	NameNotEqual string `json:"nameNotEqual,omitempty"`
	// Restrict to tasks that have a name with the given parameter value as substring.
	NameLike string `json:"nameLike,omitempty"`
	// Restrict to tasks that do not have a name with the given parameter value as substring.
	NameNotLike string `json:"nameNotLike,omitempty"`
	// Restrict to tasks that have the given description.
	Description string `json:"description,omitempty"`
	// Restrict to tasks that have a description that has the parameter value as a substring.
	DescriptionLike string `json:"descriptionLike,omitempty"`
	// Restrict to tasks that have the given priority.
	Priority int64 `json:"priority,omitempty"`
	// Restrict to tasks that have a lower or equal priority.
	MaxPriority int64 `json:"maxPriority,omitempty"`
	// Restrict to tasks that have a higher or equal priority.
	MinPriority int64 `json:"minPriority,omitempty"`
	// Restrict to tasks that are due on the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	DueDate time.Time `json:"dueDate"`
	// Restrict to tasks that are due on the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	DueDateExpression time.Time `json:"dueDateExpression,omitempty"`
	// Restrict to tasks that are due after the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	DueAfter time.Time `json:"dueAfter,omitempty"`
	// Restrict to tasks that are due after the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	DueAfterExpression string `json:"dueAfterExpression,omitempty"`
	// Restrict to tasks that are due before the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	DueBefore time.Time `json:"dueBefore,omitempty"`
	// Restrict to tasks that are due before the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	DueBeforeExpression string `json:"dueBeforeExpression,omitempty"`
	// Restrict to tasks that have a followUp date on the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	FollowUpDate time.Time `json:"followUpDate,omitempty"`
	// Restrict to tasks that have a followUp date on the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	FollowUpDateExpression string `json:"followUpDateExpression,omitempty"`
	// Restrict to tasks that have a followUp date after the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	FollowUpAfter time.Time `json:"followUpAfter,omitempty"`
	// Restrict to tasks that have a followUp date after the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	FollowUpAfterExpression string `json:"followUpAfterExpression,omitempty"`
	// Restrict to tasks that have a followUp date before the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	FollowUpBefore time.Time `json:"followUpBefore,omitempty"`
	// Restrict to tasks that have a followUp date before the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	FollowUpBeforeExpression string `json:"followUpBeforeExpression,omitempty"`
	// Restrict to tasks that have no followUp date or a followUp date before the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45. The typical use case is to query all "active" tasks for a user for a given date.
	FollowUpBeforeOrNotExistent time.Time `json:"followUpBeforeOrNotExistent,omitempty"`
	// Restrict to tasks that have no followUp date or a followUp date before the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	FollowUpBeforeOrNotExistentExpression string `json:"followUpBeforeOrNotExistentExpression,omitempty"`
	// Restrict to tasks that were created on the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45. Note: if the used database saves dates with milliseconds precision this query only will return tasks created on the given timestamp with zero milliseconds.
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// Restrict to tasks that were created on the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	CreatedOnExpression string `json:"createdOnExpression,omitempty"`
	// Restrict to tasks that were created after the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	CreatedAfter time.Time `json:"createdAfter,omitempty"`
	// Restrict to tasks that were created after the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	CreatedAfterExpression string `json:"createdAfterExpression,omitempty"`
	// Restrict to tasks that were created before the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	CreatedBefore time.Time `json:"createdBefore,omitempty"`
	// Restrict to tasks that were created before the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	CreatedBeforeExpression string `json:"createdBeforeExpression,omitempty"`
	// Restrict to tasks that are in the given delegation state. Valid values are PENDING and RESOLVED.
	DelegationState delegationState `json:"delegationState,omitempty"`
	// Restrict to tasks that are offered to any of the given candidate groups.Takes a comma-separated list of group names, so for example developers, support, sales.
	CandidateGroups []string `json:"candidateGroups,omitempty"`
	// Restrict to tasks that are offered to any of the candidate groups described by the given expression.See the user guide for more information on available functions.The expression must evaluate to java.util.List of Strings.
	CandidateGroupsExpression []string `json:"candidateGroupsExpression,omitempty"`
	// Only include tasks which have a candidate group.Value may only be true, as false is the default behavior.
	WithCandidateGroups bool `json:"withCandidateGroups,omitempty"`
	// Only include tasks which have no candidate group.Value may only be true, as false is the default behavior.
	WithoutCandidateGroups bool `json:"withoutCandidateGroups,omitempty"`
	// Only include tasks which have a candidate user.Value may only be true, as false is the default behavior.
	WithCandidateUsers bool `json:"withCandidateUsers,omitempty"`
	// Only include tasks which have no candidate users.Value may only be true, as false is the default behavior.
	WithoutCandidateUsers bool `json:"withoutCandidateUsers,omitempty"`
	// Only include active tasks.Value may only be true, as false is the default behavior.
	Active bool `json:"active,omitempty"`
	// Only include suspended tasks.Value may only be true, as false is the default behavior.
	Suspended bool `json:"suspended,omitempty"`
	// Only include tasks that have variables with certain values.Variable filtering expressions are comma-separated and are structured as follows:
	// A valid parameter value has the form key_operator_value.key is the variable name, operator is the comparison operator to be used and value the variable value.
	// Note: Values are always treated as String objects on server side.
	//
	// Valid operator values are: eq - equal to;
	// neq - not equal to;
	// gt - greater than;
	// gteq - greater than or equal to;
	// lt - lower than;
	// lteq - lower than or equal to;
	// like.
	// key and value may not contain underscore or comma characters.
	TaskVariables []VariableFilterExpression `json:"taskVariables,omitempty"`
	// Only include tasks that belong to process instances that have variables with certain values.Variable filtering expressions are comma-separated and are structured as follows:
	// A valid parameter value has the form key_operator_value.key is the variable name, operator is the comparison operator to be used and value the variable value.
	// Note: Values are always treated as String objects on server side.
	//
	// Valid operator values are: eq - equal to;
	// neq - not equal to;
	// gt - greater than;
	// gteq - greater than or equal to;
	// lt - lower than;
	// lteq - lower than or equal to;
	// like.
	// key and value may not contain underscore or comma characters.
	ProcessVariables []VariableFilterExpression `json:"processVariables,omitempty"`
	// Only include tasks that belong to case instances that have variables with certain values.Variable filtering expressions are comma-separated and are structured as follows:
	// A valid parameter value has the form key_operator_value.key is the variable name, operator is the comparison operator to be used and value the variable value.
	// Note: Values are always treated as String objects on server side.
	//
	// Valid operator values are: eq - equal to;
	// neq - not equal to;
	// gt - greater than;
	// gteq - greater than or equal to;
	// lt - lower than;
	// lteq - lower than or equal to;
	// like.
	// key and value may not contain underscore or comma characters.
	CaseInstanceVariables []VariableFilterExpression `json:"caseInstanceVariables,omitempty"`
	// Restrict query to all tasks that are sub tasks of the given task.Takes a task id.
	ParentTaskId string `json:"parentTaskId,omitempty"`
	// Sort the results lexicographically by a given criterion.Valid values are instanceId, caseInstanceId, dueDate, executionId, caseExecutionId, assignee, created, description, id, name, nameCaseInsensitive and priority.Must be used in conjunction with the sortOrder parameter.
	SortBy string `json:"sortBy,omitempty"`
	// Sort the results in a given order.Values may be asc for ascending order or desc for descending order.Must be used in conjunction with the sortBy parameter.
	SortOrder string `json:"sortOrder,omitempty"`
	// Pagination of results.Specifies the index of the first result to return.
	FirstResult int64 `json:"firstResult,omitempty"`
	// Pagination of results.Specifies the maximum number of results to return.Will return less results if there are no more results left.
	MaxResults int64 `json:"maxResults,omitempty"`
}

// QueryUserTaskComplete a query for Complete user task request
type QueryUserTaskComplete struct {
	// A JSON object containing variable key-value pairs
	Variables map[string]Variable `json:"variables"`
}

// MarshalJSON marshal to json
func (q *UserTaskGetListQuery) MarshalJSON() ([]byte, error) {
	type Alias UserTaskGetListQuery

	return json.Marshal(&struct {
		*Alias

		DueDate                     string `json:"dueDate,omitempty"`
		DueDateExpression           string `json:"dueDateExpression,omitempty"`
		DueAfter                    string `json:"dueAfter,omitempty"`
		DueBefore                   string `json:"dueBefore,omitempty"`
		FollowUpDate                string `json:"followUpDate,omitempty"`
		FollowUpAfter               string `json:"followUpAfter,omitempty"`
		FollowUpBefore              string `json:"followUpBefore,omitempty"`
		FollowUpBeforeOrNotExistent string `json:"followUpBeforeOrNotExistent,omitempty"`
		CreatedOn                   string `json:"createdOn,omitempty"`
		CreatedAfter                string `json:"createdAfter,omitempty"`
		CreatedBefore               string `json:"createdBefore,omitempty"`
	}{
		Alias: (*Alias)(q),

		DueDate:                     toCamundaTime(q.DueDate),
		DueDateExpression:           toCamundaTime(q.DueDateExpression),
		DueAfter:                    toCamundaTime(q.DueAfter),
		DueBefore:                   toCamundaTime(q.DueBefore),
		FollowUpDate:                toCamundaTime(q.FollowUpDate),
		FollowUpAfter:               toCamundaTime(q.FollowUpAfter),
		FollowUpBefore:              toCamundaTime(q.FollowUpBefore),
		FollowUpBeforeOrNotExistent: toCamundaTime(q.FollowUpBeforeOrNotExistent),
		CreatedOn:                   toCamundaTime(q.CreatedOn),
		CreatedAfter:                toCamundaTime(q.CreatedAfter),
		CreatedBefore:               toCamundaTime(q.CreatedBefore),
	})
}

// Get retrieves a task by id
func (t *userTaskApi) Get(id string) (*UserTask, error) {
	res, err := t.client.doGet("/task/"+id, map[string]string{})
	if err != nil {
		return nil, err
	}

	resp := UserTaskResponse{}
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}

	return &UserTask{
		api:              t,
		UserTaskResponse: &resp,
	}, nil
}

// GetList retrieves task list
func (t *userTaskApi) GetList(query *UserTaskGetListQuery) ([]UserTask, error) {
	if query == nil {
		query = &UserTaskGetListQuery{}
	}

	queryParams := map[string]string{}

	if query.MaxResults > 0 {
		queryParams["maxResults"] = fmt.Sprintf("%d", query.MaxResults)
	}

	if query.FirstResult > 0 {
		queryParams["firstResult"] = fmt.Sprintf("%d", query.FirstResult)
	}

	res, err := t.client.doPostJson("/task", queryParams, query)
	if err != nil {
		return nil, err
	}

	var resp []UserTask
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}

	for i := range resp {
		resp[i].api = t
	}

	return resp, nil
}

// GetListCount retrieves task list count
func (t *userTaskApi) GetListCount(query *UserTaskGetListQuery) (int64, error) {
	if query == nil {
		query = &UserTaskGetListQuery{}
	}

	queryParams := map[string]string{}

	res, err := t.client.doPostJson("/task/count", queryParams, query)
	if err != nil {
		return 0, err
	}

	resp := struct {
		Count int64 `json:"count"`
	}{}

	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return 0, fmt.Errorf("can't read json response: %w", err)
	}

	return resp.Count, nil
}

// Complete complete user task by id
func (t *userTaskApi) Complete(id string, query QueryUserTaskComplete) error {
	res, err := t.client.doPostJson("/task/"+id+"/complete", map[string]string{}, query)
	if err != nil {
		return fmt.Errorf("can't post json: %w", err)
	}

	if res != nil {
		res.Body.Close()
	}

	return nil
}

// GetIdentityLinks retrieves IdentityLinks by id
func (t *userTaskApi) GetIdentityLinks(id string) (*[]IdentityLink, error) {
	res, err := t.client.doGet(fmt.Sprintf("/task/%s/identity-links", id),
		map[string]string{})
	if err != nil {
		return nil, err
	}

	resp := []IdentityLink{}
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}

	return &resp, nil
}

// AddIdentityLink add IdentityLink to UserTask
func (t *userTaskApi) AddIdentityLink(id string, query ReqIdentityLink) error {
	res, err := t.client.doPostJson(fmt.Sprintf("/task/%s/identity-links", id),
		map[string]string{}, query)
	if err != nil {
		return fmt.Errorf("can't post json: %w", err)
	}

	if res != nil {
		res.Body.Close()
	}

	return nil
}

// DeleteIdentityLink delete IdentityLink by id
func (t *userTaskApi) DeleteIdentityLink(id string, query ReqIdentityLink) error {
	res, err := t.client.doPostJson(fmt.Sprintf("/task/%s/identity-links/delete", id),
		map[string]string{}, query)
	if err != nil {
		return fmt.Errorf("can't post json: %w", err)
	}

	if res != nil {
		res.Body.Close()
	}

	return nil
}
