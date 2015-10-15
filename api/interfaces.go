package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

// XMLClient a client capable of making XML requests to a vcloud API
type XMLClient interface {
	XMLRequest(string, string, string, RequestBody, interface{}) error
}

// DefaultConfig the default api client config, filled out with variables from the environment
func DefaultConfig() *Config {
	baseURL := os.Getenv("VCLOUDAIR_BASEURL")
	if baseURL == "" {
		baseURL = "https://vca.vmware.com/api"
	}
	return &Config{
		HTTP:       http.DefaultClient,
		Username:   os.Getenv("VCLOUDAIR_USERNAME"),
		Password:   os.Getenv("VCLOUDAIR_PASSWORD"),
		Debug:      os.Getenv("VCLOUDAIR_DEBUG") != "",
		BaseURL:    baseURL,
		APIVersion: "5.11",
	}
}

// Config is the client config for the vCloud Air API
type Config struct {
	// Override the default http client
	HTTP *http.Client
	// Username the username to use when authenticating
	Username string
	// Password the username to use when authenticating
	Password string
	// Debug, when true this will dump requests and responses with ALL parameters to the std logger.
	// All parameters also includes things like passwords etc, so be careful when you turn this on for live systems
	// because it's a security hole.
	Debug bool

	// BaseURL is the base url to use when talking to vCloud Air api's. Normal usage would not need to customize this URL.
	// This should be hugely useful in unit tests and stuff.
	BaseURL string

	// Token retrieved as result of the authentication flow
	Token string

	// APIVersion the api version currently in use
	APIVersion string
}

// NewRequestBody creates a new request body interface implementation
func NewRequestBody(contentType string, payload interface{}) RequestBody {
	return &defaultRequestBody{contentType, payload}
}

type defaultRequestBody struct {
	contentType string
	payload     interface{}
}

func (d *defaultRequestBody) Payload() interface{} {
	return d.payload
}

func (d *defaultRequestBody) ContentType() string {
	return d.contentType
}

// RequestBody an interface representing a request body
type RequestBody interface {
	Payload() interface{}
	ContentType() string
}

// XMLRequest uses the context to make XML based HTTP requests
func XMLRequest(context *Config, method, url, tpe string, body RequestBody, result interface{}) error {
	if context == nil {
		return fmt.Errorf("context needs to be provided")
	}

	r, _ := http.NewRequest(method, url, nil)
	if body != nil {
		buf := bytes.NewBuffer(nil)
		enc := xml.NewEncoder(buf)
		if err := enc.Encode(body.Payload()); err != nil {
			return err
		}
		r, _ = http.NewRequest(method, url, buf)
	}

	r.Header.Set("Accept", tpe+";version="+context.APIVersion)
	if body != nil {
		r.Header.Set("Content-Type", body.ContentType())
	}

	if context.Token != "" {
		r.Header.Set("X-Vcloud-Authorization", context.Token)
	}

	if context.Debug {
		dr, _ := httputil.DumpRequestOut(r, true)
		fmt.Println(string(dr))
	}

	resp, err := context.HTTP.Do(r)
	if err != nil {
		return err
	}
	if context.Debug {
		dr, _ := httputil.DumpResponse(resp, true)
		fmt.Println(string(dr))
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		log.Fatalf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(result); err != nil {
		return err
	}

	return nil
}
