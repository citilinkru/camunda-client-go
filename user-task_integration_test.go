package camunda_client_go

import (
	"testing"
	"time"
)

func TestIdentityLinksIntegration(t *testing.T) {
	client := NewClient(ClientOptions{
		EndpointUrl: "http://s83:8090/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	client.UserTask.IdentiyLink.GetList(client, "493a6a95-46ea-11ed-b40f-0242ac13001a")
}

func TestIdentityLinksIntegrationCreate(t *testing.T) {
	client := NewClient(ClientOptions{
		EndpointUrl: "http://s83:8090/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	client.UserTask.IdentiyLink.Add(client, "493a6a95-46ea-11ed-b40f-0242ac13001a", ReqIdentityLinkCreate{
		UserId: "userid1",
		Type:   "assignee",
	})
}
