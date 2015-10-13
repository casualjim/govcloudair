package vcloud

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindVDC(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", MimeVDC)
		rw.WriteHeader(200)
		rw.Write([]byte(vdcXML))
	}))
	defer serv.Close()

	tc, org := loadTestOrg(serv.URL)
	vdc, err := org.FindVDC("VDC1", tc)
	if assert.NoError(t, err) {
		assert.Len(t, vdc.Links, 20)
		assert.Equal(t, "First VDC", vdc.Description)
		assert.Equal(t, "AllocationVApp", vdc.AllocationModel)
		if assert.NotNil(t, vdc.ComputeCapacity) {
			cc := vdc.ComputeCapacity
			if assert.NotNil(t, cc.CPU) {
				cpu := cc.CPU
				assert.Equal(t, "MHz", cpu.Units)
				assert.EqualValues(t, 0, cpu.Allocated)
				assert.EqualValues(t, 130000, cpu.Limit)
				assert.EqualValues(t, 0, cpu.Reserved)
				assert.EqualValues(t, 0, cpu.Used)
				assert.EqualValues(t, 0, cpu.Overhead)
				assert.False(t, cpu.Required)
			}
			if assert.NotNil(t, cc.Memory) {
				mem := cc.Memory
				assert.Equal(t, "MB", mem.Units)
				assert.EqualValues(t, 0, mem.Allocated)
				assert.EqualValues(t, 102400, mem.Limit)
				assert.EqualValues(t, 0, mem.Reserved)
				assert.EqualValues(t, 0, mem.Used)
				assert.EqualValues(t, 0, mem.Overhead)
				assert.False(t, mem.Required)
			}
		}
		assert.Empty(t, vdc.ResourceEntities)
		if assert.Len(t, vdc.AvailableNetworks, 1) {
			nw := vdc.AvailableNetworks[0]
			if assert.NotNil(t, nw) {
				assert.Equal(t, "default-routed-network", nw.Name)
				assert.Equal(t, "https://us-california-1-3.vchs.vmware.com/api/compute/api/network/network-uuid-goes-here", nw.HREF)
				assert.Equal(t, "application/vnd.vmware.vcloud.network+xml", nw.Type)
			}
		}

		if assert.NotNil(t, vdc.Capabilities) && assert.NotNil(t, vdc.Capabilities.SupportedHardwareVersions) {
			shw := vdc.Capabilities.SupportedHardwareVersions
			assert.Len(t, shw.SupportedHardwareVersion, 5)
		}

		assert.EqualValues(t, 100, vdc.NICQuota)
		assert.EqualValues(t, 100, vdc.NetworkQuota)
		assert.EqualValues(t, 0, vdc.UsedNetworkCount)
		assert.EqualValues(t, 50, vdc.VMQuota)
		assert.True(t, vdc.IsEnabled)
		assert.EqualValues(t, 2600, vdc.VCPUInMHz)

		assert.Len(t, vdc.VdcStorageProfiles, 2)
	}

}

var vdcXML = `
<?xml version="1.0" encoding="UTF-8"?>
<Vdc xmlns="http://www.vmware.com/vcloud/v1.5" status="1" name="VDC1" id="urn:vcloud:vdc:vdc-uuid-goes-here" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here" type="application/vnd.vmware.vcloud.vdc+xml" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd">
    <Link rel="up" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/org-uuid-goes-here" type="application/vnd.vmware.vcloud.org+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/metadata" type="application/vnd.vmware.vcloud.metadata+xml"/>
    <Link rel="edit" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here" type="application/vnd.vmware.vcloud.vdc+xml"/>
    <Link rel="remove" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/uploadVAppTemplate" type="application/vnd.vmware.vcloud.uploadVAppTemplateParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/media" type="application/vnd.vmware.vcloud.media+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/instantiateOvf" type="application/vnd.vmware.vcloud.instantiateOvfParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/instantiateVAppTemplate" type="application/vnd.vmware.vcloud.instantiateVAppTemplateParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/cloneVApp" type="application/vnd.vmware.vcloud.cloneVAppParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/cloneVAppTemplate" type="application/vnd.vmware.vcloud.cloneVAppTemplateParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/cloneMedia" type="application/vnd.vmware.vcloud.cloneMediaParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/captureVApp" type="application/vnd.vmware.vcloud.captureVAppParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/action/composeVApp" type="application/vnd.vmware.vcloud.composeVAppParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/vdc-uuid-goes-here/disk" type="application/vnd.vmware.vcloud.diskCreateParams+xml"/>
    <Link rel="edgeGateways" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/vdc-uuid-goes-here/edgeGateways" type="application/vnd.vmware.vcloud.query.records+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/vdc-uuid-goes-here/networks" type="application/vnd.vmware.vcloud.orgVdcNetwork+xml"/>
    <Link rel="orgVdcNetworks" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/vdc-uuid-goes-here/networks" type="application/vnd.vmware.vcloud.query.records+xml"/>
    <Link rel="alternate" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/vdc-uuid-goes-here" type="application/vnd.vmware.admin.vdc+xml"/>
    <Link rel="vchs:edgeGateways" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vchs/query?type=edgeGateway&amp;vdcId=vdc-uuid-goes-here" type="application/vnd.vmware.vchs.query.records+xml"/>
    <Link rel="vchs:orgVdcNetworks" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vchs/query?type=orgVdcNetwork&amp;vdcId=vdc-uuid-goes-here" type="application/vnd.vmware.vchs.query.records+xml"/>
    <Description>First VDC</Description>
    <AllocationModel>AllocationVApp</AllocationModel>
    <ComputeCapacity>
        <Cpu>
            <Units>MHz</Units>
            <Allocated>0</Allocated>
            <Limit>130000</Limit>
            <Reserved>0</Reserved>
            <Used>0</Used>
            <Overhead>0</Overhead>
        </Cpu>
        <Memory>
            <Units>MB</Units>
            <Allocated>0</Allocated>
            <Limit>102400</Limit>
            <Reserved>0</Reserved>
            <Used>0</Used>
            <Overhead>0</Overhead>
        </Memory>
    </ComputeCapacity>
    <ResourceEntities/>
    <AvailableNetworks>
        <Network href="https://us-california-1-3.vchs.vmware.com/api/compute/api/network/network-uuid-goes-here" name="default-routed-network" type="application/vnd.vmware.vcloud.network+xml"/>
    </AvailableNetworks>
    <Capabilities>
        <SupportedHardwareVersions>
            <SupportedHardwareVersion>vmx-04</SupportedHardwareVersion>
            <SupportedHardwareVersion>vmx-07</SupportedHardwareVersion>
            <SupportedHardwareVersion>vmx-08</SupportedHardwareVersion>
            <SupportedHardwareVersion>vmx-09</SupportedHardwareVersion>
            <SupportedHardwareVersion>vmx-10</SupportedHardwareVersion>
        </SupportedHardwareVersions>
    </Capabilities>
    <NicQuota>100</NicQuota>
    <NetworkQuota>100</NetworkQuota>
    <UsedNetworkCount>0</UsedNetworkCount>
    <VmQuota>50</VmQuota>
    <IsEnabled>true</IsEnabled>
    <VdcStorageProfiles>
        <VdcStorageProfile href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcStorageProfile/ssd-storage-profile-uuid-here" name="SSD-Accelerated" type="application/vnd.vmware.vcloud.vdcStorageProfile+xml"/>
        <VdcStorageProfile href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcStorageProfile/standard-storage-profile-uuid-here" name="Standard" type="application/vnd.vmware.vcloud.vdcStorageProfile+xml"/>
    </VdcStorageProfiles>
    <VCpuInMhz2>2600</VCpuInMhz2>
</Vdc>
`
