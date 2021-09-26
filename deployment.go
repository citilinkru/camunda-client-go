package camunda_client_go

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

// Deployment a client for Deployment API
type Deployment struct {
	client *Client
}

// ResDeployment a JSON array of deployment objects
type ResDeployment struct {
	// The id of the deployment
	Id string `json:"id"`
	// The name of the deployment
	Name string `json:"name"`
	// The source of the deployment
	Source string `json:"source"`
	// The tenant id of the deployment
	TenantId string `json:"tenantId"`
	// The date and time of the deployment.
	DeploymentTime Time `json:"deploymentTime"`
}

// ResDeploymentCreate a JSON object corresponding to the DeploymentWithDefinitions interface in the engine
type ResDeploymentCreate struct {
	// The id of the deployment
	Id string `json:"id"`
	// The name of the deployment
	Name string `json:"name"`
	// The source of the deployment
	Source string `json:"source"`
	// The tenant id of the deployment
	TenantId string `json:"tenant_id"`
	// The time when the deployment was created
	DeploymentTime Time `json:"deployment_time"`
	// Link to the newly created deployment with method, href and rel
	Links []ResLink `json:"links"`
	// A JSON Object containing a property for each of the process definitions,
	// which are successfully deployed with that deployment
	DeployedProcessDefinitions map[string]ResProcessDefinition `json:"deployedProcessDefinitions"`
	// A JSON Object containing a property for each of the case definitions,
	// which are successfully deployed with that deployment
	DeployedCaseDefinitions map[string]ResCaseDefinition `json:"deployedCaseDefinitions"`
	// A JSON Object containing a property for each of the decision definitions,
	// which are successfully deployed with that deployment
	DeployedDecisionDefinitions map[string]ResDecisionDefinition `json:"deployedDecisionDefinitions"`
	// A JSON Object containing a property for each of the decision requirements definitions,
	// which are successfully deployed with that deployment
	DeployedDecisionRequirementsDefinitions map[string]ResDecisionRequirementsDefinition `json:"deployedDecisionRequirementsDefinitions"`
}

// ReqDeploymentCreate a request to deployment create
type ReqDeploymentCreate struct {
	DeploymentName           string
	EnableDuplicateFiltering *bool
	DeployChangedOnly        *bool
	DeploymentSource         *string
	TenantId                 *string
	Resources                map[string]interface{}
}

// ReqRedeploy a request to redeploy
type ReqRedeploy struct {
	// A list of deployment resource ids to re-deploy
	ResourceIds *string `json:"resourceIds,omitempty"`
	// A list of deployment resource names to re-deploy
	ResourceNames *string `json:"resourceNames,omitempty"`
	// Sets the source of the deployment
	Source *string `json:"source,omitempty"`
}

// ResDeploymentResource a JSON array containing all deployment resources of the given deployment
type ResDeploymentResource struct {
	// The id of the deployment resource
	Id string `json:"id"`
	// The name of the deployment resource
	Name string `json:"name"`
	// The id of the deployment
	DeploymentId string `json:"deploymentId"`
}

// GetList a queries for deployments that fulfill given parameters. Parameters may be the properties of deployments,
// such as the id or name or a range of the deployment time. The size of the result set can be retrieved by using
// the Get Deployment count method.
// Query parameters described in the documentation:
// https://docs.camunda.org/manual/latest/reference/rest/deployment/get-query/#query-parameters
func (d *Deployment) GetList(query map[string]string) (deployments []*ResDeployment, err error) {
	res, err := d.client.doGet("/deployment", query)
	if err != nil {
		return
	}

	err = d.client.readJsonResponse(res, &deployments)
	return
}

// GetListCount a queries for the number of deployments that fulfill given parameters.
// Takes the same parameters as the Get Deployments method
func (d *Deployment) GetListCount(query map[string]string) (count int, err error) {
	res, err := d.client.doGet("/deployment/count", query)
	if err != nil {
		return
	}

	resCount := ResCount{}
	err = d.client.readJsonResponse(res, &resCount)
	return resCount.Count, err
}

// Get retrieves a deployment by id, according to the Deployment interface of the engine
func (d *Deployment) Get(id string) (deployment ResDeployment, err error) {
	res, err := d.client.doGet("/deployment/"+id, map[string]string{})
	if err != nil {
		return
	}

	err = d.client.readJsonResponse(res, &deployment)
	return
}

// Create creates a deployment
func (d *Deployment) Create(deploymentCreate ReqDeploymentCreate) (deployment *ResDeploymentCreate, err error) {
	deployment = &ResDeploymentCreate{}
	var data []byte
	body := bytes.NewBuffer(data)
	w := multipart.NewWriter(body)

	if err = w.WriteField("deployment-name", deploymentCreate.DeploymentName); err != nil {
		return nil, err
	}

	if deploymentCreate.EnableDuplicateFiltering != nil {
		if err = w.WriteField("enable-duplicate-filtering", strconv.FormatBool(*deploymentCreate.EnableDuplicateFiltering)); err != nil {
			return nil, err
		}
	}

	if deploymentCreate.DeployChangedOnly != nil {
		if err = w.WriteField("deploy-changed-only", strconv.FormatBool(*deploymentCreate.DeployChangedOnly)); err != nil {
			return nil, err
		}
	}

	if deploymentCreate.DeploymentSource != nil {
		if err = w.WriteField("deployment-source", *deploymentCreate.DeploymentSource); err != nil {
			return nil, err
		}
	}

	if deploymentCreate.TenantId != nil {
		if err = w.WriteField("tenant-id", *deploymentCreate.TenantId); err != nil {
			return nil, err
		}
	}

	for key, resource := range deploymentCreate.Resources {
		var fw io.Writer

		if x, ok := resource.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := resource.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}

		if r, ok := resource.(io.Reader); ok {
			if _, err = io.Copy(fw, r); err != nil {
				return nil, err
			}
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	res, err := d.client.do(http.MethodPost, "/deployment/create", map[string]string{}, body, w.FormDataContentType())
	if err != nil {
		return nil, err
	}

	err = d.client.readJsonResponse(res, deployment)

	return deployment, err
}

// Redeploy a re-deploys an existing deployment.
// The deployment resources to re-deploy can be restricted by using the properties resourceIds or resourceNames.
// If no deployment resources to re-deploy are passed then all existing resources of the given deployment
// are re-deployed
func (d *Deployment) Redeploy(id string, req ReqRedeploy) (deployment *ResDeploymentCreate, err error) {
	deployment = &ResDeploymentCreate{}
	res, err := d.client.doPostJson("/deployment/"+id+"/redeploy", map[string]string{}, &req)
	if err != nil {
		return
	}

	err = d.client.readJsonResponse(res, deployment)
	return
}

// GetResources retrieves all deployment resources of a given deployment
func (d *Deployment) GetResources(id string) (resources []*ResDeploymentResource, err error) {
	res, err := d.client.doGet("/deployment/"+id+"/resources", map[string]string{})
	if err != nil {
		return
	}

	err = d.client.readJsonResponse(res, &resources)
	return
}

// GetResource retrieves a deployment resource by resource id for the given deployment
func (d *Deployment) GetResource(id, resourceId string) (resource *ResDeploymentResource, err error) {
	resource = &ResDeploymentResource{}
	res, err := d.client.doGet("/deployment/"+id+"/resources/"+resourceId, map[string]string{})
	if err != nil {
		return
	}

	err = d.client.readJsonResponse(res, &resource)
	return
}

// GetResourceBinary retrieves the binary content of a deployment resource for the given deployment by id
func (d *Deployment) GetResourceBinary(id, resourceId string) (data []byte, err error) {
	res, err := d.client.doGet("/deployment/"+id+"/resources/"+resourceId+"/data", map[string]string{})
	if err != nil {
		return
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// Delete deletes a deployment by id
func (d *Deployment) Delete(id string, query map[string]string) error {
	err := d.client.doDelete("/deployment/"+id, query)
	return err
}
