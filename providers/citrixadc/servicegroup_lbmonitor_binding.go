package citrixadc

import (
	"log"
	"fmt"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGroupLbMonitorBindingGenerator struct {
	CitrixService
}

func (g *ServiceGroupLbMonitorBindingGenerator) createServiceGroupLbMonitorBinding(client *service.NitroClient) error {
	serviceGroups, err := client.FindAllResources(service.Servicegroup.Type())
	if err != nil {
		return err
	}
	for _, sg := range serviceGroups {
		servicegroupname := sg["servicegroupname"].(string)
		findParams := service.FindParams{
			ResourceType: service.Servicegroup_lbmonitor_binding.Type(),
			ResourceName: servicegroupname,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if len(dataArr) == 0 {
			continue
		}
		for _, t := range dataArr {
			monitorname := t["monitor_name"].(string)
			id := fmt.Sprintf("%s,%s", servicegroupname, monitorname)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				id,
				normalizeResourceName(id),
				"citrixadc_servicegroup_lbmonitor_binding",
				g.ProviderName,
				map[string]string{
					"servicegroupname": servicegroupname,
					"monitorname": monitorname,
				},
				[]string{""},
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *ServiceGroupLbMonitorBindingGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createServiceGroupLbMonitorBinding(client); err != nil {
		return err
	}
	return nil
}