package vcloud

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
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
	fixedCatalogXML := rewriteXML(publicCatalogXML, serv.URL)
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
		rw.Header().Set("Content-Type", MimeVAppTemplate)
		rw.WriteHeader(200)
		rw.Write([]byte(vappTemplateXML))
	}))
	defer serv.Close()

	tc := newTestXMLClient(serv.URL)
	fixedCatalogItemXML := rewriteXML(catalogItemExample, serv.URL)
	var ci CatalogItem
	if err := xml.Unmarshal([]byte(fixedCatalogItemXML), &ci); assert.NoError(t, err) {
		templ, err := ci.VAppTemplate(tc)
		if assert.NoError(t, err) {
			assert.Len(t, templ.Links, 8)
			assert.Equal(t, "id: VMW-PHOTON-TP2-64BIT", templ.Description)
			if assert.NotNil(t, templ.Owner) && assert.NotNil(t, templ.Owner.User) {
				usr := templ.Owner.User
				assert.Equal(t, "system", usr.Name)
			}
			if assert.NotNil(t, templ.NetworkConfigSection) {
				ncs := templ.NetworkConfigSection
				// TODO: make this actually work, doesn't seem to grab the info from the ovf:Info element
				//assert.Equal(t, "The configuration parameters for logical networks", ncs.Info)
				if assert.NotNil(t, ncs.NetworkConfig) {
					assert.Equal(t, "This is a special place-holder used for disconnected network interfaces.", ncs.NetworkConfig.Description)
					assert.False(t, ncs.NetworkConfig.IsDeployed)
					assert.Equal(t, "none", ncs.NetworkConfig.NetworkName)
					if assert.NotNil(t, ncs.NetworkConfig.Configuration) {
						ncfg := ncs.NetworkConfig.Configuration
						assert.Equal(t, "isolated", ncfg.FenceMode)
						if assert.NotNil(t, ncfg.IPScopes) && assert.NotEmpty(t, ncfg.IPScopes.IPScope) {
							ipsc := ncfg.IPScopes.IPScope[0]
							assert.False(t, ipsc.IsInherited)
							assert.Equal(t, "196.254.254.254", ipsc.Gateway)
							assert.Equal(t, "255.255.0.0", ipsc.Netmask)
							assert.Equal(t, "196.254.254.254", ipsc.DNS1)
						}
					}
				}
			}
		}

		if assert.Len(t, templ.Children, 1) {
			vm := templ.Children[0]
			assert.False(t, vm.GoldMaster)
			assert.EqualValues(t, 8, vm.Status)
			assert.Equal(t, "Photon", vm.Name)
			assert.Equal(t, "urn:vcloud:vm:vm-uuid-goes-here", vm.ID)
			assert.Equal(t, "https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here", vm.HREF)
			assert.Len(t, vm.Links, 4)
			assert.Equal(t, "vm", vm.VAppScopedLocalID)

			if assert.NotNil(t, vm.NetworkConnectionSection) {
				nwcs := vm.NetworkConnectionSection
				//assert.Equal(t, "Specifies the available VM network connections", nwcs.Info)
				assert.EqualValues(t, 0, nwcs.PrimaryNetworkConnectionIndex)
				if assert.NotNil(t, nwcs.NetworkConnection) {
					nc := nwcs.NetworkConnection
					assert.EqualValues(t, 0, nc.NetworkConnectionIndex)
					assert.Equal(t, "none", nc.Network)
					assert.True(t, nc.NeedsCustomization)
					assert.Equal(t, "00:50:56:1d:b4:c5", nc.MACAddress)
					assert.Equal(t, "NONE", nc.IPAddressAllocationMode)
					assert.Equal(t, "VMXNET3", nc.NetworkAdapterType)
				}
			}

			if assert.NotNil(t, vm.GuestCustomizationSection) {
				gcs := vm.GuestCustomizationSection
				assert.True(t, gcs.Enabled)
				assert.False(t, gcs.ChangeSid)
				assert.Equal(t, "vm-uuid-goes-here", gcs.VirtualMachineID)
				assert.False(t, gcs.JoinDomainEnabled)
				assert.False(t, gcs.UseOrgSettings)
				assert.True(t, gcs.AdminPasswordEnabled)
				assert.True(t, gcs.AdminPasswordAuto)
				assert.False(t, gcs.AdminAutoLogonEnabled)
				assert.EqualValues(t, 0, gcs.AdminAutoLogonCount)
				assert.Equal(t, "Photon-001", gcs.ComputerName)
			}
		}

	}
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
