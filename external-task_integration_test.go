// +build integration

package camunda_client_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFetchAndLockIntegration(t *testing.T) {
	processKey := "hello-world-process"

	_, err := client.ProcessDefinition.StartInstance(
		QueryProcessDefinitionBy{Key: &processKey},
		ReqStartInstance{Variables: &map[string]Variable{
			"isWorld": {Value: false, Type: "boolean"},
			"test":    {Value: false, Type: "boolean"},
		}},
	)
	assert.NoError(t, err)

	_, err = client.ProcessDefinition.StartInstance(
		QueryProcessDefinitionBy{Key: &processKey},
		ReqStartInstance{Variables: &map[string]Variable{
			"isWorld": {Value: false, Type: "boolean"},
			"test":    {Value: true, Type: "boolean"},
		}},
	)
	assert.NoError(t, err)

	// wait processing StartInstance in camunda
	time.Sleep(time.Second * 10)

	tasks, err := client.ExternalTask.FetchAndLock(QueryFetchAndLock{
		WorkerId: "test-fetch-and-lock-integration",
		MaxTasks: 10,
		Topics: &[]QueryFetchAndLockTopic{
			{
				LockDuration: 1000,
				TopicName:    "PrintHello",
				ProcessVariables: map[string]interface{}{
					"test": false,
				},
			},
		},
	})

	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.False(t, tasks[0].Variables["test"].Value.(bool))

	tasks, err = client.ExternalTask.FetchAndLock(QueryFetchAndLock{
		WorkerId: "test-fetch-and-lock-integration",
		MaxTasks: 10,
		Topics: &[]QueryFetchAndLockTopic{
			{
				LockDuration: 1000,
				TopicName:    "PrintHello",
				ProcessVariables: map[string]interface{}{
					"test": true,
				},
			},
		},
	})

	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.True(t, tasks[0].Variables["test"].Value.(bool))
}
