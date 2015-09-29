package v57

import (
	"fmt"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOAuthToken(t *testing.T) {
	authToken := "imagine this is a ridiculoulsly long kind of random string"

	tc := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Accept")
		mt, params, err := mime.ParseMediaType(ct)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(500)
			return
		}

		if mt != "application/json" {
			rw.WriteHeader(406)
			return
		}

		if params["version"] != "5.7" {
			rw.WriteHeader(412)
			return
		}

		if r.URL.Path == LoginPath {
			rw.Header().Set("vchs-authorization", authToken)
			rw.Header().Set("Content-Type", ct)
			rw.WriteHeader(200)
			rw.Write([]byte(`{"serviceGroupIds":["service-group-uuid-goes-here"]}`))
			return
		}

		if r.URL.Path == InstancesPath {
			rw.Header().Set("Content-Type", ct+"; class=com.vmware.vchs.sc.restapi.model.instancelisttype")
			rw.WriteHeader(200)
			rw.Write([]byte(instancesJSON))
			return
		}

		rw.WriteHeader(404)

	}))
	defer tc.Close()

	_, err := NewAuthenticatedSession(Config{
		BaseURL:  tc.URL,
		Debug:    true,
		Username: "some user",
		Password: "some password",
	})
	assert.NoError(t, err)
}

var instancesJSON = `{
    "instances": [
        {
            "apiUrl": "https://storage.googleapis.com",
            "dashboardUrl": "https://us-california-1-3.vchs.vmware.com/os-g/ui/",
            "description": "Highly scalable and durable storage. Create buckets, upload and manage objects.",
            "id": "0000000000001",
            "instanceAttributes": "random-words-hex",
            "instanceVersion": "6",
            "link": [],
            "name": "Object Storage powered by Google",
            "planId": "region:us-california-1-3.vchs.vmware.com:planID:some-uuid-here-1",
            "region": "us-california-1-3.vchs.vmware.com",
            "serviceGroupId": "service-group-uuid-goes-here"
        },
        {
            "apiUrl": "https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here",
            "dashboardUrl": "https://us-california-1-3.vchs.vmware.com/api/compute/compute/ui/index.html?orgName=org-name-uuid-goes-here&serviceInstanceId=org-uuid-goes-here&servicePlan=plan-uuid-goes-here",
            "description": "Create virtual machines, and easily scale up or down as your needs change.",
            "id": "org-uuid-goes-here",
            "instanceAttributes": "{\"orgName\":\"org-name-uuid-goes-here\",\"sessionUri\":\"https://us-california-1-3.vchs.vmware.com/api/compute/api/sessions\",\"apiVersionUri\":\"https://us-california-1-3.vchs.vmware.com/api/compute/api/versions\"}",
            "instanceVersion": "1.0",
            "link": [],
            "name": "Virtual Private Cloud OnDemand",
            "planId": "region:us-california-1-3.vchs.vmware.com:planID:plan-uuid-goes-here",
            "region": "us-california-1-3.vchs.vmware.com",
            "serviceGroupId": "service-group-uuid-goes-here"
        }
    ]
}`
