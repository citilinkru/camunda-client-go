package camunda_client_go

// ResCaseDefinition a JSON object corresponding to the CaseDefinition interface in the engine
type ResCaseDefinition struct {
	// The id of the case definition
	Id string `json:"id"`
	// The key of the case definition, i.e., the id of the CMMN 2.0 XML case definition
	Key string `json:"key"`
	// The category of the case definition
	Category string `json:"category"`
	// The name of the case definition
	Name string `json:"name"`
	// The version of the case definition that the engine assigned to it
	Version int `json:"Version"`
	// The file name of the case definition
	Resource string `json:"resource"`
	// The deployment id of the case definition
	DeploymentId string `json:"deploymentId"`
	// The tenant id of the case definition
	TenantId string `json:"tenantId"`
	// History time to live value of the case definition. Is used within History cleanup
	HistoryTimeToLive int `json:"historyTimeToLive"`
}
