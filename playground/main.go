package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vmware/govcloudair/api/v57"
	"github.com/vmware/govcloudair/schemas/vcloud"
)

func main() {

	ses, err := v57.NewAuthenticatedSession(nil)
	if err != nil {
		log.Fatalln(err)
	}

	orgList, err := ses.OrgList()
	if err != nil {
		log.Fatalln(err)
	}

	org, err := orgList.FirstOrg(ses)
	if err != nil {
		log.Fatalln(err)
	}

	catalog, err := org.RetrievePublicCatalog(ses)
	if err != nil {
		log.Fatalln(err)
	}

	catalogItem, err := catalog.ItemForName("VMware Photon OS - Tech Preview 2", ses)
	if err != nil {
		log.Fatalln(err)
	}

	vdc, err := org.FindVDC("VDC1", ses)
	if err != nil {
		log.Fatalln(err)
	}

	vAppTemplate, err := catalogItem.VAppTemplate(ses)
	if err != nil {
		log.Fatalln(err)
	}

	inst := vcloud.NewInstantiateVAppTemplateParams()
	inst.Deploy = true
	inst.PowerOn = true
	inst.Name = "Linux FTP Server"
	inst.AllEULAsAccepted = true
	params := new(vcloud.InstantiationParams)
	nc := new(vcloud.VAppNetworkConfiguration)
	nc.NetworkName = "vAppNetwork"
	nc.Configuration = new(vcloud.NetworkConfiguration)
	nc.Configuration.ParentNetwork = vdc.AvailableNetworks[0]
	nc.Configuration.FenceMode = "bridged"
	ncs := new(vcloud.NetworkConfigSection)
	ncs.Info = "Configuration parameters for logical networks"
	ncs.NetworkConfig = nc
	params.NetworkConfigSection = ncs
	inst.InstantiationParams = params
	inst.Source = vAppTemplate.Ref()

	vapp, err := vdc.InstantiateVAppTemplate(inst, ses)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.MarshalIndent(vapp, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

}
