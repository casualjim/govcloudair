package vcloud

var vappTemplateXML = `
<?xml version="1.0" encoding="UTF-8"?>
<VAppTemplate xmlns="http://www.vmware.com/vcloud/v1.5" xmlns:ovf="http://schemas.dmtf.org/ovf/envelope/1" xmlns:vssd="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData" xmlns:rasd="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData" xmlns:vmw="http://www.vmware.com/schema/ovf" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" goldMaster="false" ovfDescriptorUploaded="true" status="8" name="VMware Photon OS - Tech Preview 2" id="urn:vcloud:vapptemplate:vapptemplate-uuid-goes-her" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her" type="application/vnd.vmware.vcloud.vAppTemplate+xml" xsi:schemaLocation="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_VirtualSystemSettingData.xsd http://www.vmware.com/schema/ovf http://www.vmware.com/schema/ovf http://schemas.dmtf.org/ovf/envelope/1 http://schemas.dmtf.org/ovf/envelope/1/dsp8023_1.1.0.xsd http://www.vmware.com/vcloud/v1.5 http://us-california-1-3.vchs.vmware.com/api/compute/api/v1.5/schema/master.xsd http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_ResourceAllocationSettingData.xsd">
    <Link rel="catalogItem" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/catalogItem/catalogitem-uuid-here" type="application/vnd.vmware.vcloud.catalogItem+xml"/>
    <Link rel="enable" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/action/enableDownload"/>
    <Link rel="disable" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/action/disableDownload"/>
    <Link rel="ovf" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/ovf" type="text/xml"/>
    <Link rel="storageProfile" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcStorageProfile/storage-profile-uuid-goes-here" name="NFS-Catalog Storage Policy" type="application/vnd.vmware.vcloud.vdcStorageProfile+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/owner" type="application/vnd.vmware.vcloud.owner+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/metadata" type="application/vnd.vmware.vcloud.metadata+xml"/>
    <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/productSections/" type="application/vnd.vmware.vcloud.productSections+xml"/>
    <Description>id: VMW-PHOTON-TP2-64BIT</Description>
    <Owner type="application/vnd.vmware.vcloud.owner+xml">
        <User href="https://us-california-1-3.vchs.vmware.com/api/compute/api/admin/user/user-uuid-goes-here" name="system" type="application/vnd.vmware.admin.user+xml"/>
    </Owner>
    <Children>
        <Vm goldMaster="false" status="8" name="Photon" id="urn:vcloud:vm:vm-uuid-goes-here" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here" type="application/vnd.vmware.vcloud.vm+xml">
            <Link rel="up" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her" type="application/vnd.vmware.vcloud.vAppTemplate+xml"/>
            <Link rel="storageProfile" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcStorageProfile/storage-profile-uuid-goes-here" type="application/vnd.vmware.vcloud.vdcStorageProfile+xml"/>
            <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/metadata" type="application/vnd.vmware.vcloud.metadata+xml"/>
            <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/productSections/" type="application/vnd.vmware.vcloud.productSections+xml"/>
            <NetworkConnectionSection href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/networkConnectionSection/" type="application/vnd.vmware.vcloud.networkConnectionSection+xml" ovf:required="false">
                <ovf:Info>Specifies the available VM network connections</ovf:Info>
                <PrimaryNetworkConnectionIndex>0</PrimaryNetworkConnectionIndex>
                <NetworkConnection needsCustomization="true" network="none">
                    <NetworkConnectionIndex>0</NetworkConnectionIndex>
                    <IsConnected>false</IsConnected>
                    <MACAddress>00:50:56:1d:b4:c5</MACAddress>
                    <IpAddressAllocationMode>NONE</IpAddressAllocationMode>
                    <NetworkAdapterType>VMXNET3</NetworkAdapterType>
                </NetworkConnection>
            </NetworkConnectionSection>
            <GuestCustomizationSection href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/guestCustomizationSection/" type="application/vnd.vmware.vcloud.guestCustomizationSection+xml" ovf:required="true">
                <ovf:Info>Specifies Guest OS Customization Settings</ovf:Info>
                <Enabled>true</Enabled>
                <ChangeSid>false</ChangeSid>
                <VirtualMachineId>vm-uuid-goes-here</VirtualMachineId>
                <JoinDomainEnabled>false</JoinDomainEnabled>
                <UseOrgSettings>false</UseOrgSettings>
                <AdminPasswordEnabled>true</AdminPasswordEnabled>
                <AdminPasswordAuto>true</AdminPasswordAuto>
                <AdminAutoLogonEnabled>false</AdminAutoLogonEnabled>
                <AdminAutoLogonCount>0</AdminAutoLogonCount>
                <ResetPasswordRequired>false</ResetPasswordRequired>
                <ComputerName>Photon-001</ComputerName>
            </GuestCustomizationSection>
            <ovf:VirtualHardwareSection xmlns:vcloud="http://www.vmware.com/vcloud/v1.5" ovf:transport="" vcloud:href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/virtualHardwareSection/" vcloud:type="application/vnd.vmware.vcloud.virtualHardwareSection+xml">
                <ovf:Info>Virtual hardware requirements</ovf:Info>
                <ovf:System>
                    <vssd:ElementName>Virtual Hardware Family</vssd:ElementName>
                    <vssd:InstanceID>0</vssd:InstanceID>
                    <vssd:VirtualSystemIdentifier>Photon</vssd:VirtualSystemIdentifier>
                    <vssd:VirtualSystemType>vmx-10</vssd:VirtualSystemType>
                </ovf:System>
                <ovf:Item>
                    <rasd:Address>00:50:56:1d:b4:c5</rasd:Address>
                    <rasd:AddressOnParent>0</rasd:AddressOnParent>
                    <rasd:AutomaticAllocation>false</rasd:AutomaticAllocation>
                    <rasd:Connection vcloud:primaryNetworkConnection="true" vcloud:ipAddressingMode="NONE">none</rasd:Connection>
                    <rasd:Description>Vmxnet3 ethernet adapter on "none"</rasd:Description>
                    <rasd:ElementName>Network adapter 0</rasd:ElementName>
                    <rasd:InstanceID>1</rasd:InstanceID>
                    <rasd:ResourceSubType>VMXNET3</rasd:ResourceSubType>
                    <rasd:ResourceType>10</rasd:ResourceType>
                </ovf:Item>
                <ovf:Item>
                    <rasd:Address>0</rasd:Address>
                    <rasd:Description>SCSI Controller</rasd:Description>
                    <rasd:ElementName>SCSI Controller 0</rasd:ElementName>
                    <rasd:InstanceID>2</rasd:InstanceID>
                    <rasd:ResourceSubType>lsilogic</rasd:ResourceSubType>
                    <rasd:ResourceType>6</rasd:ResourceType>
                </ovf:Item>
                <ovf:Item>
                    <rasd:AddressOnParent>0</rasd:AddressOnParent>
                    <rasd:Description>Hard disk</rasd:Description>
                    <rasd:ElementName>Hard disk 1</rasd:ElementName>
                    <rasd:HostResource vcloud:capacity="16000" vcloud:storageProfileOverrideVmDefault="false" vcloud:busSubType="lsilogic" vcloud:storageProfileHref="https://us-california-1-3.vchs.vmware.com/api/compute/api/vdcStorageProfile/storage-profile-uuid-goes-here" vcloud:busType="6"/>
                    <rasd:InstanceID>2000</rasd:InstanceID>
                    <rasd:Parent>2</rasd:Parent>
                    <rasd:ResourceType>17</rasd:ResourceType>
                    <rasd:VirtualQuantity>16777216000</rasd:VirtualQuantity>
                    <rasd:VirtualQuantityUnits>byte</rasd:VirtualQuantityUnits>
                </ovf:Item>
                <ovf:Item>
                    <rasd:Address>1</rasd:Address>
                    <rasd:Description>IDE Controller</rasd:Description>
                    <rasd:ElementName>IDE Controller 1</rasd:ElementName>
                    <rasd:InstanceID>3</rasd:InstanceID>
                    <rasd:ResourceType>5</rasd:ResourceType>
                </ovf:Item>
                <ovf:Item>
                    <rasd:AddressOnParent>0</rasd:AddressOnParent>
                    <rasd:AutomaticAllocation>false</rasd:AutomaticAllocation>
                    <rasd:Description>CD/DVD Drive</rasd:Description>
                    <rasd:ElementName>CD/DVD Drive 1</rasd:ElementName>
                    <rasd:HostResource/>
                    <rasd:InstanceID>3002</rasd:InstanceID>
                    <rasd:Parent>3</rasd:Parent>
                    <rasd:ResourceType>15</rasd:ResourceType>
                </ovf:Item>
                <ovf:Item>
                    <rasd:AddressOnParent>0</rasd:AddressOnParent>
                    <rasd:AutomaticAllocation>false</rasd:AutomaticAllocation>
                    <rasd:Description>Floppy Drive</rasd:Description>
                    <rasd:ElementName>Floppy Drive 1</rasd:ElementName>
                    <rasd:HostResource/>
                    <rasd:InstanceID>8000</rasd:InstanceID>
                    <rasd:ResourceType>14</rasd:ResourceType>
                </ovf:Item>
                <ovf:Item>
                    <rasd:AllocationUnits>hertz * 10^6</rasd:AllocationUnits>
                    <rasd:Description>Number of Virtual CPUs</rasd:Description>
                    <rasd:ElementName>1 virtual CPU(s)</rasd:ElementName>
                    <rasd:InstanceID>4</rasd:InstanceID>
                    <rasd:Reservation>0</rasd:Reservation>
                    <rasd:ResourceType>3</rasd:ResourceType>
                    <rasd:VirtualQuantity>1</rasd:VirtualQuantity>
                    <rasd:Weight>0</rasd:Weight>
                    <vmw:CoresPerSocket ovf:required="false">1</vmw:CoresPerSocket>
                </ovf:Item>
                <ovf:Item>
                    <rasd:AllocationUnits>byte * 2^20</rasd:AllocationUnits>
                    <rasd:Description>Memory Size</rasd:Description>
                    <rasd:ElementName>2048 MB of memory</rasd:ElementName>
                    <rasd:InstanceID>5</rasd:InstanceID>
                    <rasd:Reservation>0</rasd:Reservation>
                    <rasd:ResourceType>4</rasd:ResourceType>
                    <rasd:VirtualQuantity>2048</rasd:VirtualQuantity>
                    <rasd:Weight>0</rasd:Weight>
                </ovf:Item>
                <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/virtualHardwareSection/cpu" type="application/vnd.vmware.vcloud.rasdItem+xml"/>
                <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/virtualHardwareSection/memory" type="application/vnd.vmware.vcloud.rasdItem+xml"/>
                <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/virtualHardwareSection/disks" type="application/vnd.vmware.vcloud.rasdItemsList+xml"/>
                <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/virtualHardwareSection/media" type="application/vnd.vmware.vcloud.rasdItemsList+xml"/>
                <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/virtualHardwareSection/networkCards" type="application/vnd.vmware.vcloud.rasdItemsList+xml"/>
                <Link rel="down" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vm-vm-uuid-goes-here/virtualHardwareSection/serialPorts" type="application/vnd.vmware.vcloud.rasdItemsList+xml"/>
            </ovf:VirtualHardwareSection>
            <VAppScopedLocalId>vm</VAppScopedLocalId>
            <DateCreated>2015-08-28T19:19:03.320Z</DateCreated>
        </Vm>
    </Children>
    <ovf:NetworkSection xmlns:vcloud="http://www.vmware.com/vcloud/v1.5" vcloud:href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/networkSection/" vcloud:type="application/vnd.vmware.vcloud.networkSection+xml">
        <ovf:Info>The list of logical networks</ovf:Info>
        <ovf:Network ovf:name="none">
            <ovf:Description>This is a special place-holder used for disconnected network interfaces.</ovf:Description>
        </ovf:Network>
    </ovf:NetworkSection>
    <NetworkConfigSection href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/networkConfigSection/" type="application/vnd.vmware.vcloud.networkConfigSection+xml" ovf:required="false">
        <ovf:Info>The configuration parameters for logical networks</ovf:Info>
        <NetworkConfig networkName="none">
            <Description>This is a special place-holder used for disconnected network interfaces.</Description>
            <Configuration>
                <IpScopes>
                    <IpScope>
                        <IsInherited>false</IsInherited>
                        <Gateway>196.254.254.254</Gateway>
                        <Netmask>255.255.0.0</Netmask>
                        <Dns1>196.254.254.254</Dns1>
                    </IpScope>
                </IpScopes>
                <FenceMode>isolated</FenceMode>
            </Configuration>
            <IsDeployed>false</IsDeployed>
        </NetworkConfig>
    </NetworkConfigSection>
    <LeaseSettingsSection href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/leaseSettingsSection/" type="application/vnd.vmware.vcloud.leaseSettingsSection+xml" ovf:required="false">
        <ovf:Info>Lease settings section</ovf:Info>
        <Link rel="edit" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/leaseSettingsSection/" type="application/vnd.vmware.vcloud.leaseSettingsSection+xml"/>
        <StorageLeaseInSeconds>0</StorageLeaseInSeconds>
    </LeaseSettingsSection>
    <CustomizationSection goldMaster="false" href="https://us-california-1-3.vchs.vmware.com/api/compute/api/vAppTemplate/vappTemplate-vapptemplate-uuid-goes-her/customizationSection/" type="application/vnd.vmware.vcloud.customizationSection+xml" ovf:required="false">
        <ovf:Info>VApp template customization section</ovf:Info>
        <CustomizeOnInstantiate>true</CustomizeOnInstantiate>
    </CustomizationSection>
    <DateCreated>2015-08-28T19:19:03.320Z</DateCreated>
</VAppTemplate>
`
