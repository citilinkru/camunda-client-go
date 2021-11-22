package camunda_client_go

// Tenant a client for Tenant
type Tenant struct {
	client *Client
}

// Create a new tenant.
// `id` - The id of the tenant.
// `name` - The name of the tenant.
func (p *Tenant) Create(id, name string) (err error) {
	req := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{
		Id:   id,
		Name: name,
	}
	_, err = p.client.doPostJson("/tenant/create", map[string]string{}, &req)
	return
}
