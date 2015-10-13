package vcloud

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware/govcloudair/api"
)

func newTestXMLClient(url string) *testXMLClient {
	fixedSessionXML := strings.Replace(sessionsXML, "https://us-california-1-3.vchs.vmware.com", url, -1)
	var tc testXMLClient
	if err := xml.Unmarshal([]byte(fixedSessionXML), &tc); err != nil {
		panic(err)
	}
	tc.Config = testConfig
	return &tc
}

type testXMLClient struct {
	XMLName xml.Name `xml:"http://www.vmware.com/vcloud/v1.5 Session"`
	Links   []*Link  `xml:"Link"`
	Config  *api.Config
}

// XMLRequest makes HTTP request that have XML bodies and get XML results
func (s *testXMLClient) XMLRequest(method, url, tpe string, body, result interface{}) error {
	return api.XMLRequest(s.Config, method, url, tpe, body, result)
}

var testConfig = &api.Config{
	HTTP:  http.DefaultClient,
	Debug: true,
}

func TestFetchOrgList(t *testing.T) { // sanity check for serializing the org list
	serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", MimeOrgList)
		rw.WriteHeader(200)
		rw.Write([]byte(orgListXML))
	}))
	defer serv.Close()

	tc := newTestXMLClient(serv.URL)
	ol, err := FetchOrgList(tc.Links, tc)
	if assert.NoError(t, err) {
		assert.Len(t, ol.Orgs, 1)
	}
}

func TestFetchOrg(t *testing.T) { // sanity check for serializing the org
	serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", MimeOrg)
		rw.WriteHeader(200)
		rw.Write([]byte(orgXML))
	}))
	defer serv.Close()

	tc := newTestXMLClient(serv.URL)
	fixedOrgListXML := strings.Replace(orgListXML, "https://us-california-1-3.vchs.vmware.com", serv.URL, -1)
	var orgList OrgList
	if err := xml.Unmarshal([]byte(fixedOrgListXML), &orgList); assert.NoError(t, err) {
		o, err := orgList.FirstOrg(tc)
		if assert.NoError(t, err) {
			assert.Len(t, o.Links, 15)
		}
	}
}

func loadTestOrg(url string) (api.XMLClient, *Org) {
	tc := newTestXMLClient(url)
	fixedOrgXML := strings.Replace(orgXML, "https://us-california-1-3.vchs.vmware.com", url, -1)
	var org Org
	if err := xml.Unmarshal([]byte(fixedOrgXML), &org); err != nil {
		panic(err)
	}
	return tc, &org
}

var orgListXML = `<?xml version="1.0" encoding="UTF-8"?>
<OrgList xmlns="http://www.vmware.com/vcloud/v1.5" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/" type="application/vnd.vmware.vcloud.orgList+xml" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd">
    <Org href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here" name="org-name-uuid-goes-here" type="application/vnd.vmware.vcloud.org+xml"/>
</OrgList>
`

var orgXML = `<?xml version="1.0" encoding="UTF-8"?>
<Org xmlns="http://www.vmware.com/vcloud/v1.5" name="org-name-uuid-goes-here" id="urn:vcloud:org:org-uuid-goes-here" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here" type="application/vnd.vmware.vcloud.org+xml" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd">
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here" name="VDC1" type="application/vnd.vmware.vcloud.vdc+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/tasksList/org-uuid-goes-here" type="application/vnd.vmware.vcloud.tasksList+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/default-catalog-uuid-goes-here" name="default-catalog" type="application/vnd.vmware.vcloud.catalog+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/default-catalog-uuid-goes-here/controlAccess/" type="application/vnd.vmware.vcloud.controlAccess+xml"/>
    <Link rel="controlAccess" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/default-catalog-uuid-goes-here/action/controlAccess" type="application/vnd.vmware.vcloud.controlAccess+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/d666d38a-fcc3-4fae-9a31-83b928921878" name="Public Catalog" type="application/vnd.vmware.vcloud.catalog+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/d666d38a-fcc3-4fae-9a31-83b928921878/controlAccess/" type="application/vnd.vmware.vcloud.controlAccess+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/org/org-uuid-goes-here/catalogs" type="application/vnd.vmware.admin.catalog+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/network/default-routed-network-uuid-goes-here" name="default-routed-network" type="application/vnd.vmware.vcloud.orgNetwork+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/supportedSystemsInfo/" type="application/vnd.vmware.vcloud.supportedSystemsInfo+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here/metadata" type="application/vnd.vmware.vcloud.metadata+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here/hybrid" type="application/vnd.vmware.vcloud.hybridOrg+xml"/>
    <Link rel="alternate" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/org/org-uuid-goes-here" type="application/vnd.vmware.admin.organization+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcTemplates" type="application/vnd.vmware.admin.vdcTemplates+xml"/>
    <Link rel="instantiate" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here/action/instantiate" type="application/vnd.vmware.vcloud.instantiateVdcTemplateParams+xml"/>
    <Description/>
    <FullName>full-org-name-uuid-goes-here (org-name-uuid-goes-here)</FullName>
</Org>
`

var sessionsXML = `<?xml version="1.0" encoding="UTF-8"?>
<Session xmlns="http://www.vmware.com/vcloud/v1.5" org="org-name-uuid-goes-here" roles="Account Administrator" user="someone@somewhere.com" userId="urn:vcloud:user:user-uuid-goes-here" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/session" type="application/vnd.vmware.vcloud.session+xml" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd">
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/" type="application/vnd.vmware.vcloud.orgList+xml"/>
    <Link rel="remove" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/session"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/" type="application/vnd.vmware.admin.vcloud+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here" name="org-name-uuid-goes-here" type="application/vnd.vmware.vcloud.org+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/query" type="application/vnd.vmware.vcloud.query.queryList+xml"/>
    <Link rel="entityResolver" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/entity/" type="application/vnd.vmware.vcloud.entity+xml"/>
    <Link rel="down:extensibility" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/extensibility" type="application/vnd.vmware.vcloud.apiextensibility+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vchs/query?type=edgeGateway" type="application/vnd.vmware.vchs.query.records+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vchs/query?type=orgVdcNetwork" type="application/vnd.vmware.vchs.query.records+xml"/>
</Session>
`
