package vcloud

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatalogItem(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", MimeCatalogItem)
		rw.WriteHeader(200)
		rw.Write([]byte(catalogItemExample))
	}))
	defer serv.Close()

	tc := newTestXMLClient(serv.URL)
	fixedCatalogXML := strings.Replace(publicCatalogXML, "https://us-california-1-3.vchs.vmware.com", serv.URL, -1)
	var catalog Catalog
	if err := xml.Unmarshal([]byte(fixedCatalogXML), &catalog); assert.NoError(t, err) {
		ci, err := catalog.ItemForName("VMware Photon OS - Tech Preview 2", tc)
		if assert.NoError(t, err) {
			assert.Len(t, ci.Links, 2)
			assert.Equal(t, "id: VMW-PHOTON-TP2-64BIT", ci.Description)
		}
	}
}

func TestFetchVAppTemplate(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", MimeCatalogItem)
		rw.WriteHeader(200)
		rw.Write([]byte(catalogItemExample))
	}))
	defer serv.Close()

}

var catalogItemExample = `
<?xml version="1.0" encoding="UTF-8"?>
<CatalogItem xmlns="http://www.vmware.com/vcloud/v1.5" size="0" name="VMware Photon OS - Tech Preview 2" id="urn:vcloud:catalogitem:catalog-item-uuid-goes-here" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/catalog-item-uuid-goes-here" type="application/vnd.vmware.vcloud.catalogItem+xml" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd">
    <Link rel="up" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalog/catalog-uuid-goes-here" type="application/vnd.vmware.vcloud.catalog+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/catalog-item-uuid-goes-here/metadata" type="application/vnd.vmware.vcloud.metadata+xml"/>
    <Description>id: VMW-PHOTON-TP2-64BIT</Description>
    <Entity href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-uuid-goes-here" name="VMware Photon OS - Tech Preview 2" type="application/vnd.vmware.vcloud.vAppTemplate+xml"/>
    <DateCreated>2015-08-28T19:19:03.107Z</DateCreated>
    <VersionNumber>2</VersionNumber>
</CatalogItem>
`
