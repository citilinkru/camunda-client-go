package camunda_client_go

// ResDecisionRequirementsDefinition a JSON object corresponding to the DecisionRequirementsDefinition
// interface in the engine
type ResDecisionRequirementsDefinition struct {
	// The id of the decision requirements definition
	Id string `json:"id"`
	// The key of the decision requirements definition, i.e., the id of the DMN 1.1 XML decision definition
	Key string `json:"key"`
	// The category of the decision requirements definition
	Category string `json:"category"`
	// The name of the decision requirements definition
	Name string `json:"name"`
	// The version of the decision requirements definition that the engine assigned to it
	Version int `json:"Version"`
	// The file name of the decision requirements definition
	Resource string `json:"resource"`
	// The deployment id of the decision requirements definition
	DeploymentId string `json:"deploymentId"`
	// The tenant id of the decision requirements definition
	TenantId string `json:"tenantId"`
}
