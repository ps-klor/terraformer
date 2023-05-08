package citrixadc

import (
	"log"
	"fmt"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LbserverServiceGroupBindingGenerator struct {
	CitrixService
}

func (g *LbserverServiceGroupBindingGenerator) createLbserverServiceGroupBindingGenerator(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Lbvserver.Type())
	if err != nil {
		return err
	}
	for _, s := range server {
		name := s["name"].(string)
		findParams := service.FindParams{
			ResourceType: service.Lbvserver_servicegroup_binding.Type(),
			ResourceName: name,
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
			servicegroupname := t["servicegroupname"].(string)
			id := fmt.Sprintf("%s,%s", name, servicegroupname)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				id,
				normalizeResourceName(id),
				"citrixadc_lbvserver_servicegroup_binding",
				g.ProviderName,
				map[string]string{},
				[]string{""},
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *LbserverServiceGroupBindingGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createLbserverServiceGroupBindingGenerator(client); err != nil {
		return err
	}
	return nil
}