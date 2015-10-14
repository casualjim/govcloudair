package vcloud

const (
	// PublicCatalog Name
	PublicCatalog = "Public Catalog"

	// DefaultCatalog Name
	DefaultCatalog = "Default Catalog"

	// JSONMimeV57 the json mime for version 5.7 of the API
	JSONMimeV57 = "application/json;version=5.7"
	// AnyXMLMime511 the wildcard xml mime for version 5.11 of the API
	AnyXMLMime511 = "application/*+xml;version=5.11"
	// Version511 the 5.11 version
	Version511 = "5.11"
	// Version is the default version number
	Version = Version511
)

const (
	// NsOvf the ovf xml namespace url
	NsOvf = "http://schemas.dmtf.org/ovf/envelope/1"
	// NsXMLSchema the xml schema namespace url
	NsXMLSchema = "http://www.w3.org/2001/XMLSchema-instance"
	// NsVCloud vcloud xml namespace url
	NsVCloud = "http://www.vmware.com/vcloud/v1.5"
)

const (
	// MimeOrgList mime for org list
	MimeOrgList = "application/vnd.vmware.vcloud.orgList+xml"
	// MimeOrg mime for org
	MimeOrg = "application/vnd.vmware.vcloud.org+xml"
	// MimeCatalog mime for catalog
	MimeCatalog = "application/vnd.vmware.vcloud.catalog+xml"
	// MimeCatalogItem mime for catalog item
	MimeCatalogItem = "application/vnd.vmware.vcloud.catalogItem+xml"
	// MimeVDC mime for a VDC
	MimeVDC = "application/vnd.vmware.vcloud.vdc+xml"
	// MimeVAppTemplate mime for a vapp template
	MimeVAppTemplate = "application/vnd.vmware.vcloud.vAppTemplate+xml"
	// MimeInstantiateVAppTemplate mime fore instantiate VApp template params
	MimeInstantiateVAppTemplate = "application/vnd.vmware.vcloud.instantiateVAppTemplateParams+xml"
	// MimeQueryRecords mime for the query records
	MimeQueryRecords = "application/vnd.vmware.vchs.query.records+xml"
	// MimeAPIExtensibility mime for api extensibility
	MimeAPIExtensibility = "application/vnd.vmware.vcloud.apiextensibility+xml"
	// MimeEntity mime for vcloud entity
	MimeEntity = "application/vnd.vmware.vcloud.entity+xml"
	// MimeQueryList mime for query list
	MimeQueryList = "application/vnd.vmware.vcloud.query.queryList+xml"
	// MimeSession mime for a session
	MimeSession = "application/vnd.vmware.vcloud.session+xml"
	// MimeTask mime for task
	MimeTask = "application/vnd.vmware.vcloud.task+xml"
	// MimeError mime for error
	MimeError = "application/vnd.vmware.vcloud.error+xml"
	// MimeNetwork mime for a network
	MimeNetwork = "application/vnd.vmware.vcloud.network+xml"
)

const (
	RelDown          = "down"
	RelAdd           = "add"
	RelUp            = "up"
	RelEdit          = "edit"
	RelRemove        = "remove"
	RelCopy          = "copy"
	RelMove          = "move"
	RelAlternate     = "alternate"
	RelTaskCancel    = "task:cancel"
	RelDeploy        = "deploy"
	RelUndeploy      = "undeploy"
	RelDiscardState  = "discardState"
	RelPowerOn       = "power:powerOn"
	RelPowerOff      = "power:powerOff"
	RelPowerReset    = "power:reset"
	RelPowerReboot   = "power:reboot"
	RelPowerSuspend  = "power:suspend"
	RelPowerShutdown = "power:shutdown"

	RelScreenThumbnail        = "screen:thumbnail"
	RelScreenAcquireTicket    = "screen:acquireTicket"
	RelScreenAcquireMksTicket = "screen:acquireMksTicket"

	RelMediaInsertMedia = "media:insertMedia"
	RelMediaEjectMedia  = "media:ejectMedia"

	RelDiskAttach = "disk:attach"
	RelDiskDetach = "disk:detach"

	RelUploadDefault   = "upload:default"
	RelUploadAlternate = "upload:alternate"

	RelDownloadDefault   = "download:default"
	RelDownloadAlternate = "download:alternate"
	RelDownloadIdentity  = "download:identity"

	RelSnapshotCreate          = "snapshot:create"
	RelSnapshotRevertToCurrent = "snapshot:revertToCurrent"
	RelSnapshotRemoveAll       = "snapshot:removeAll"

	RelOVF               = "ovf"
	RelOVA               = "ova"
	RelControlAccess     = "controlAccess"
	RelPublish           = "publish"
	RelPublishExternal   = "publishToExternalOrganizations"
	RelSubscribeExternal = "subscribeToExternalCatalog"
	RelExtension         = "extension"
	RelEnable            = "enable"
	RelDisable           = "disable"
	RelMerge             = "merge"
	RelCatalogItem       = "catalogItem"
	RelRecompose         = "recompose"
	RelRegister          = "register"
	RelUnregister        = "unregister"
	RelRepair            = "repair"
	RelReconnect         = "reconnect"
	RelDisconnect        = "disconnect"
	RelUpgrade           = "upgrade"
	RelAnswer            = "answer"
	RelAddOrgs           = "addOrgs"
	RelRemoveOrgs        = "removeOrgs"
	RelSync              = "sync"

	RelVSphereWebClientURL = "vSphereWebClientUrl"
	RelVimServerDvSwitches = "vimServerDvSwitches"

	RelCollaborationResume    = "resume"
	RelCollaborationAbort     = "abort"
	RelCollaborationFail      = "fail"
	RelEnterMaintenanceMode   = "enterMaintenanceMode"
	RelExitMaintenanceMode    = "exitMaintenanceMode"
	RelTask                   = "task"
	RelTaskOwner              = "task:owner"
	RelPreviousPage           = "previousPage"
	RelNextPage               = "nextPage"
	RelFirstPage              = "firstPage"
	RelLastPage               = "lastPage"
	RelInstallVMWareTools     = "installVmwareTools"
	RelConsolidate            = "consolidate"
	RelEntity                 = "entity"
	RelEntityResolver         = "entityResolver"
	RelRelocate               = "relocate"
	RelBlockingTasks          = "blockingTasks"
	RelUpdateProgress         = "updateProgress"
	RelSyncSyslogSettings     = "syncSyslogSettings"
	RelTakeOwnership          = "takeOwnership"
	RelUnlock                 = "unlock"
	RelShadowVMs              = "shadowVms"
	RelTest                   = "test"
	RelUpdateResourcePools    = "update:resourcePools"
	RelRemoveForce            = "remove:force"
	RelStorageClass           = "storageProfile"
	RelRefreshStorageClasses  = "refreshStorageProfile"
	RelRefreshVirtualCenter   = "refreshVirtualCenter"
	RelCheckCompliance        = "checkCompliance"
	RelForceFullCustomization = "customizeAtNextPowerOn"
	RelReloadFromVC           = "reloadFromVc"
	RelMetricsDayView         = "interval:day"
	RelMetricsWeekView        = "interval:week"
	RelMetricsMonthView       = "interval:month"
	RelMetricsYearView        = "interval:year"
	RelMetricsPreviousRange   = "range:previous"
	RelMetricsNextRange       = "range:next"
	RelMetricsLatestRange     = "range:latest"
	RelRights                 = "rights"
	RelMigratVMs              = "migrateVms"
	RelResourcePoolVMList     = "resourcePoolVmList"
	RelCreateEvent            = "event:create"
	RelCreateTask             = "task:create"
	RelUploadBundle           = "bundle:upload"
	RelCleanupBundles         = "bundles:cleanup"
	RelAuthorizationCheck     = "authorization:check"
	RelCleanupRights          = "rights:cleanup"

	RelEdgeGatewayRedeploy           = "edgeGateway:redeploy"
	RelEdgeGatewayReapplyServices    = "edgeGateway:reapplyServices"
	RelEdgeGatewayConfigureServices  = "edgeGateway:configureServices"
	RelEdgeGatewayConfigureSyslog    = "edgeGateway:configureSyslogServerSettings"
	RelEdgeGatewaySyncSyslogSettings = "edgeGateway:syncSyslogSettings"
	RelEdgeGatewayUpgrade            = "edgeGateway:upgrade"
	RelEdgeGatewayUpgradeNetworking  = "edgeGateway:convertToAdvancedNetworking"
	RelVDCManageFirewall             = "manageFirewall"

	RelCertificateUpdate = "certificate:update"
	RelCertificateReset  = "certificate:reset"
	RelTruststoreUpdate  = "truststore:update"
	RelTruststoreReset   = "truststore:reset"
	RelKeyStoreUpdate    = "keystore:update"
	RelKeystoreReset     = "keystore:reset"
	RelKeytabUpdate      = "keytab:update"
	RelKeytabReset       = "keytab:reset"

	RelServiceLinks             = "down:serviceLinks"
	RelAPIFilters               = "down:apiFilters"
	RelResourceClasses          = "down:resourceClasses"
	RelResourceClassActions     = "down:resourceClassActions"
	RelServices                 = "down:services"
	RelACLRules                 = "down:aclRules"
	RelFileDescriptors          = "down:fileDescriptors"
	RelAPIDefinitions           = "down:apiDefinitions"
	RelServiceResources         = "down:serviceResources"
	RelExtensibility            = "down:extensibility"
	RelAPIServiceQuery          = "down:service"
	RelAPIDefinitionsQuery      = "down:apidefinitions"
	RelAPIFilesQuery            = "down:files"
	RelServiceOfferings         = "down:serviceOfferings"
	RelServiceOfferingInstances = "down:serviceOfferingInstances"
	RelHybrid                   = "down:hybrid"

	RelServiceRefresh      = "service:refresh"
	RelServiceAssociate    = "service:associate"
	RelServiceDisassociate = "service:disassociate"

	RelReconfigureVM = "reconfigureVM"

	RelOrgVDCGateways = "edgeGateways"
	RelOrgVDCNetworks = "orgVdcNetworks"

	RelHybridAcquireControlTicket = "hybrid:acquireControlTicket"
	RelHybridAcquireTicket        = "hybrid:acquireTicket"
	RelHybridRefreshTunnel        = "hybrid:refreshTunnel"

	RelMetrics = "metrics"

	RelFederationRegenerateCertificate = "federation:regenerateFederationCertificate"
	RelTemplateInstantiate             = "instantiate"
)

const (
	// HTTPGet the http GET method
	HTTPGet = "GET"
	// HTTPPost the http POST method
	HTTPPost = "POST"
	// HTTPPut the http PUT method
	HTTPPut = "PUT"
	// HTTPPatch the http PATCH method
	HTTPPatch = "PATCH"
	// HTTPDelete the http DELETE method
	HTTPDelete = "DELETE"
)
