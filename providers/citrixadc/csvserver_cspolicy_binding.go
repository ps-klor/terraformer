package citrixadc

import (
	"log"
	"fmt"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type CsvServerCsPolicyBindingGenerator struct {
	CitrixService
}

func (g *CsvServerCsPolicyBindingGenerator) createLbserverRewritePolicyBindingGenerator(client *service.NitroClient) error {
	cs, err := client.FindAllResources(service.Csvserver.Type())
	if err != nil {
		return err
	}
	for _, c := range cs {
		name := c["name"].(string)
		findParams := service.FindParams{
			ResourceType: service.Csvserver_cspolicy_binding.Type(),
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
			policyname := t["policyname"].(string)
			id := fmt.Sprintf("%s,%s", name, policyname)
			
			g.Resources = append(g.Resources, terraformutils.NewResource(
				id,
				normalizeResourceName(id),
				"citrixadc_csvserver_cspolicy_binding",
				g.ProviderName,
				map[string]string{},
				[]string{""},
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *CsvServerCsPolicyBindingGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createLbserverRewritePolicyBindingGenerator(client); err != nil {
		return err
	}
	return nil
}