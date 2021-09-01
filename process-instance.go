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
}

// ReqProcessVariable a JSON object with the following properties: (at least an empty JSON object {}
// or an empty request body)
type ReqProcessVariable struct {
	// The variable's value. For variables of type Object, the serialized value has to be submitted as a String value.
	// For variables of type File the value has to be submitted as Base64 encoded string.
	Value *interface{} `json:"value,omitempty"`
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
	_, err := p.client.doDelete(by.String(), nil)
	return err
}

// GetBinaryProcessVariable retrieves the content of a Process Variable by the Process Instance id and the
// Process Variable name. Applicable for byte array or file Process Variables.
func (p *ProcessInstance) GetBinaryProcessVariable(by QueryProcessInstanceVariableBy) (data []byte, err error) {
	res, err := p.client.doGet(by.String()+"/data", nil)
	if err != nil {
		return
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// GetProcessVariable retrieves a variable of a given process instance by id.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/variables/get-variable/#query-parameters
func (p *ProcessInstance) GetProcessVariable(by QueryProcessInstanceVariableBy, query map[string]string) (variable *ResProcessVariable, err error) {
	variable = &ResProcessVariable{}
	res, err := p.client.doGet(by.String(), query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, variable)
	return
}

// GetProcessVariables retrieves all variables of a given process instance by id.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/variables/get-variables/#query-parameters
func (p *ProcessInstance) GetProcessVariables(id string, query map[string]string) (variables map[string]*ResProcessVariable, err error) {
	res, err := p.client.doGet("/process-instance/"+id, query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &variables)
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

// DeleteProcessInstance deletes a running process instance by id.
// https://docs.camunda.org/manual/latest/reference/rest/process-instance/delete/#query-parameters
func (p *ProcessInstance) DeleteProcessInstance(id string, query map[string]string) error {
	_, err := p.client.doDelete("/process-instance/"+id, query)
	return err
}
