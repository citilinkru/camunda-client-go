package camunda_client_go

// ReqSort json representing sort criteria
type ReqSort struct {
	// Mandatory. Sort the results lexicographically by a given criterion.
	// Valid values are instanceId, definitionKey, definitionId, tenantId and businessKey.
	SortBy string `json:"sortBy"`
	// Mandatory. Sort the results in a given order. Values may be asc for ascending order or desc for descending order.
	SortOrder string `json:"sortOrder"`
}

// ResCount a count response
type ResCount struct {
	Count int `json:"count"`
}

// ResLink a link struct
type ResLink struct {
	Method string `json:"method"`
	Href   string `json:"href"`
	Rel    string `json:"rel"`
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
	// The job definition id for the monitor jobs of this batch.
	MonitorJobDefinitionId string `json:"monitorJobDefinitionId"`
	// The job definition id for the batch execution jobs of this batch.
	BatchJobDefinitionId string `json:"batchJobDefinitionId"`
	// Indicates whether this batch is suspended or not.
	Suspended bool `json:"suspended"`
	// The tenant id of the batch.
	TenantId string `json:"tenantId"`
	// The id of the user that created the batch.
	CreateUserId string `json:"createUserId"`
}
