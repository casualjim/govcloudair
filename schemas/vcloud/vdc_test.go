package vcloud

var vdcXML = `
<?xml version="1.0" encoding="UTF-8"?>
<Vdc xmlns="http://www.vmware.com/vcloud/v1.5" status="1" name="VDC1" id="urn:vcloud:vdc:ffa862b3-abe5-4b13-9f91-138c18d11357" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357" type="application/vnd.vmware.vcloud.vdc+xml" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd">
    <Link rel="up" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/org/64d2185d-3ef6-42f3-a277-92250d9c96a9" type="application/vnd.vmware.vcloud.org+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/metadata" type="application/vnd.vmware.vcloud.metadata+xml"/>
    <Link rel="edit" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357" type="application/vnd.vmware.vcloud.vdc+xml"/>
    <Link rel="remove" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/uploadVAppTemplate" type="application/vnd.vmware.vcloud.uploadVAppTemplateParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/media" type="application/vnd.vmware.vcloud.media+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/instantiateOvf" type="application/vnd.vmware.vcloud.instantiateOvfParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/instantiateVAppTemplate" type="application/vnd.vmware.vcloud.instantiateVAppTemplateParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/cloneVApp" type="application/vnd.vmware.vcloud.cloneVAppParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/cloneVAppTemplate" type="application/vnd.vmware.vcloud.cloneVAppTemplateParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/cloneMedia" type="application/vnd.vmware.vcloud.cloneMediaParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/captureVApp" type="application/vnd.vmware.vcloud.captureVAppParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/action/composeVApp" type="application/vnd.vmware.vcloud.composeVAppParams+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/disk" type="application/vnd.vmware.vcloud.diskCreateParams+xml"/>
    <Link rel="edgeGateways" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/edgeGateways" type="application/vnd.vmware.vcloud.query.records+xml"/>
    <Link rel="add" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/networks" type="application/vnd.vmware.vcloud.orgVdcNetwork+xml"/>
    <Link rel="orgVdcNetworks" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357/networks" type="application/vnd.vmware.vcloud.query.records+xml"/>
    <Link rel="alternate" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/vdc/ffa862b3-abe5-4b13-9f91-138c18d11357" type="application/vnd.vmware.admin.vdc+xml"/>
    <Link rel="vchs:edgeGateways" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vchs/query?type=edgeGateway&amp;vdcId=ffa862b3-abe5-4b13-9f91-138c18d11357" type="application/vnd.vmware.vchs.query.records+xml"/>
    <Link rel="vchs:orgVdcNetworks" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vchs/query?type=orgVdcNetwork&amp;vdcId=ffa862b3-abe5-4b13-9f91-138c18d11357" type="application/vnd.vmware.vchs.query.records+xml"/>
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
        <Network href="https://us-california-1-3.vchs.vmware.com/api/compute/api/network/77f2dfcd-cd66-4e57-9dd0-8ae942a324e8" name="default-routed-network" type="application/vnd.vmware.vcloud.network+xml"/>
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
        <VdcStorageProfile href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcStorageProfile/57ff8080-14bf-4e68-a683-39a02dfa86b1" name="SSD-Accelerated" type="application/vnd.vmware.vcloud.vdcStorageProfile+xml"/>
        <VdcStorageProfile href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcStorageProfile/82a52275-f394-4546-9263-2de1c065004d" name="Standard" type="application/vnd.vmware.vcloud.vdcStorageProfile+xml"/>
    </VdcStorageProfiles>
    <VCpuInMhz2>2600</VCpuInMhz2>
</Vdc>
`
