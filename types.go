package camunda_client_go

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
