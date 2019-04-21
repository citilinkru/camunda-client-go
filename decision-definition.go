package camunda_client_go

// ResDecisionDefinition a JSON object corresponding to the DecisionDefinition interface in the engine
type ResDecisionDefinition struct {
	// The id of the decision definition
	Id string `json:"id"`
	// The key of the decision definition, i.e., the id of the DMN 1.0 XML decision definition
	Key string `json:"key"`
	// The category of the decision definition
	Category string `json:"category"`
	// The name of the decision definition
	Name string `json:"name"`
	// The version of the decision definition that the engine assigned to it
	Version int `json:"Version"`
	// The file name of the decision definition
	Resource string `json:"resource"`
	// The deployment id of the decision definition
	DeploymentId string `json:"deploymentId"`
	// The tenant id of the decision definition
	TenantId string `json:"tenantId"`
	// The id of the decision requirements definition this decision definition belongs to
	DecisionRequirementsDefinitionId string `json:"decisionRequirementsDefinitionId"`
	// The key of the decision requirements definition this decision definition belongs to
	DecisionRequirementsDefinitionKey string `json:"decisionRequirementsDefinitionKey"`
	// The version tag of the process definition
	VersionTag string `json:"versionTag"`
	// History time to live value of the decision definition. Is used within History cleanup
	HistoryTimeToLive int `json:"historyTimeToLive"`
}
