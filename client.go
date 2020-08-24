package camundaclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const PackageVersion = "{{version}}"
const DefaultUserAgent = "CamundaClientGo/" + PackageVersion
const DefaultEndpointUrl = "http://localhost:8080/engine-rest"
const DefaultTimeoutSec = 60
const DefaultDateTimeFormat = "2006-01-02T15:04:05.000-0700"

// ClientOptions a client options
type ClientOptions struct {
	UserAgent   string
	EndpointUrl string
	Timeout     time.Duration
	ApiUser     string
	ApiPassword string
}

// Client a client for Camunda API
type Client struct {
	httpClient  *http.Client
	endpointUrl string
	userAgent   string
	apiUser     string
	apiPassword string

	ExternalTask      *ExternalTask
	Deployment        *Deployment
	ProcessDefinition *ProcessDefinition
	UserTask          *userTaskApi
}

var ErrorNotFound = &Error{
	Type:    "NotFound",
	Message: "Not found",
}

// Error a custom error type
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Error error message
func (e *Error) Error() string {
	return e.Message
}

// Time a custom time format
type Time struct {
	time.Time
}

// UnmarshalJSON
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse(DefaultDateTimeFormat, strings.Trim(string(b), "\""))
	return
}

// MarshalJSON
func (t *Time) MarshalJSON() ([]byte, error) {
	timeStr := t.Time.Format(DefaultDateTimeFormat)
	return []byte("\"" + timeStr + "\""), nil
}

// toCamundaTime return time formatted for camunda
func toCamundaTime(dt time.Time) string {
	if dt.IsZero() {
		return ""
	}

	return dt.Format(DefaultDateTimeFormat)
}

// NewClient a create new instance Client
func NewClient(options ClientOptions) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: time.Second * DefaultTimeoutSec,
		},
		endpointUrl: DefaultEndpointUrl,
		userAgent:   DefaultUserAgent,
		apiUser:     options.ApiUser,
		apiPassword: options.ApiPassword,
	}

	if options.EndpointUrl != "" {
		client.endpointUrl = options.EndpointUrl
	}

	if options.UserAgent != "" {
		client.userAgent = options.UserAgent
	}

	if options.Timeout.Nanoseconds() != 0 {
		client.httpClient.Timeout = options.Timeout
	}

	client.ExternalTask = &ExternalTask{client: client}
	client.Deployment = &Deployment{client: client}
	client.ProcessDefinition = &ProcessDefinition{client: client}
	client.UserTask = &userTaskApi{client: client}

	return client
}

// SetCustomTransport set new custom transport
func (c *Client) SetCustomTransport(customHTTPTransport http.RoundTripper) {
	if c.httpClient != nil {
		c.httpClient.Transport = customHTTPTransport
	}
}

func (c *Client) doPostJson(path string, query map[string]string, v interface{}) (res *http.Response, err error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(v); err != nil {
		return nil, err
	}

	res, err = c.do(http.MethodPost, path, query, body, "application/json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) doPutJson(path string, query map[string]string, v interface{}) (res *http.Response, err error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(v); err != nil {
		return nil, err
	}

	res, err = c.do(http.MethodPut, path, query, body, "application/json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) doDelete(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodDelete, path, query, nil, "")
}

func (c *Client) doPost(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodPost, path, query, nil, "")
}

func (c *Client) doPut(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodPut, path, query, nil, "")
}

func (c *Client) do(method, path string, query map[string]string, body io.Reader, contentType string) (res *http.Response, err error) {
	url, err := c.buildUrl(path, query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	req.SetBasicAuth(c.apiUser, c.apiPassword)

	res, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := c.checkResponse(res); err != nil {
		return nil, err
	}

	return
}

func (c *Client) doGet(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodGet, path, query, nil, "")
}

func (c *Client) checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	defer res.Body.Close()

	if res.Header.Get("Content-Type") == "application/json" {
		if res.StatusCode == 404 {
			return ErrorNotFound
		}

		jsonErr := &Error{}
		err := json.NewDecoder(res.Body).Decode(jsonErr)
		if err != nil {
			return fmt.Errorf("response error with status code %d: failed unmarshal error response: %s", res.StatusCode, err)
		}

		return jsonErr
	}

	errText, err := ioutil.ReadAll(res.Body)
	if err == nil {
		return fmt.Errorf("response error with status code %d: %s", res.StatusCode, string(errText))
	}

	return fmt.Errorf("response error with status code %d", res.StatusCode)
}

func (c *Client) readJsonResponse(res *http.Response, v interface{}) error {
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) buildUrl(path string, query map[string]string) (string, error) {
	if len(query) == 0 {
		return c.endpointUrl + path, nil
	}
	url, err := url.Parse(c.endpointUrl + path)
	if err != nil {
		return "", err
	}

	q := url.Query()
	for k, v := range query {
		q.Set(k, v)
	}

	url.RawQuery = q.Encode()
	return url.String(), nil
}
