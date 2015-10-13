package vcloud

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCatalog(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", MimeCatalog)
		rw.WriteHeader(200)
		rw.Write([]byte(publicCatalogXML))
	}))
	defer serv.Close()

	fixedSessionXML := strings.Replace(sessionsXML, "https://us-california-1-3.vchs.vmware.com", serv.URL, -1)
	var tc testXMLClient
	if err := xml.Unmarshal([]byte(fixedSessionXML), &tc); assert.NoError(t, err) {
		tc.Config = testConfig
		fixedOrgXML := strings.Replace(orgXML, "https://us-california-1-3.vchs.vmware.com", serv.URL, -1)
		var org Org
		if err := xml.Unmarshal([]byte(fixedOrgXML), &org); assert.NoError(t, err) {
			o, err := org.RetrieveCatalog(PublicCatalog, &tc)
			if assert.NoError(t, err) {
				assert.Len(t, o.Links, 2)
				assert.Len(t, o.CatalogItems, 10)
			}
		}
	}
}

var publicCatalogXML = `<?xml version="1.0" encoding="UTF-8"?>
<Catalog xmlns="http://www.vmware.com/vcloud/v1.5" name="Public Catalog" id="urn:vcloud:catalog:d666d38a-fcc3-4fae-9a31-83b928921878" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/d666d38a-fcc3-4fae-9a31-83b928921878" type="application/vnd.vmware.vcloud.catalog+xml" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd">
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/d666d38a-fcc3-4fae-9a31-83b928921878/metadata" type="application/vnd.vmware.vcloud.metadata+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/d666d38a-fcc3-4fae-9a31-83b928921878/controlAccess/" type="application/vnd.vmware.vcloud.controlAccess+xml"/>
    <Description>vCHS service catalog</Description>
    <CatalogItems>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/00acfc20-afd4-4219-a452-5d3eecfa2b26" id="00acfc20-afd4-4219-a452-5d3eecfa2b26" name="CentOS64-64BIT" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/178c60b8-b42a-4c12-bfdd-ff0a2c8bd721" id="178c60b8-b42a-4c12-bfdd-ff0a2c8bd721" name="CentOS64-32BIT" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/1b868313-3d5b-4e71-af8e-82f7d54bdd1e" id="1b868313-3d5b-4e71-af8e-82f7d54bdd1e" name="Ubuntu Server 12.04 LTS (amd64 20150127)" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/22ba359a-d5a6-4026-9dd5-408277666b49" id="22ba359a-d5a6-4026-9dd5-408277666b49" name="Ubuntu Server 12.04 LTS (i386 20150127)" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/5b0f892d-a1d6-467f-ac53-feb9f6e5cc17" id="5b0f892d-a1d6-467f-ac53-feb9f6e5cc17" name="W2K12-STD-64BIT" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/63daa159-62bf-460d-8c02-907448eeefdf" id="63daa159-62bf-460d-8c02-907448eeefdf" name="CentOS63-64BIT" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/7b1f43ad-a21e-412a-a519-c8d877fe4f92" id="7b1f43ad-a21e-412a-a519-c8d877fe4f92" name="VMware Photon OS - Tech Preview 2" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/81b47b60-5721-4aa8-8663-0f08c3592a51" id="81b47b60-5721-4aa8-8663-0f08c3592a51" name="W2K12-STD-R2-64BIT" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/c2d74058-7bcf-43f3-868e-8a2dda71182f" id="c2d74058-7bcf-43f3-868e-8a2dda71182f" name="W2K8-STD-R2-64BIT" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
        <CatalogItem href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/f03cdae7-cbf3-4241-bde8-9a91bc6b97e2" id="f03cdae7-cbf3-4241-bde8-9a91bc6b97e2" name="CentOS63-32BIT" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
    </CatalogItems>
    <IsPublished>true</IsPublished>
    <DateCreated>2015-03-31T00:08:51.760Z</DateCreated>
    <VersionNumber>159</VersionNumber>
</Catalog>`
