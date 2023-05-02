package citrixadc

import (
	"log"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGroupGenerator struct {
	CitrixService
}

func (g *ServiceGroupGenerator) createServiceGroup(client *service.NitroClient) error {
	sg, err := client.FindAllResources(service.Servicegroup.Type())
	if err != nil {
		return err
	}
	for _, t := range sg {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			t["servicegroupname"].(string),
			normalizeResourceName(t["servicegroupname"].(string)),
			"citrixadc_servicegroup",
			g.ProviderName,
			map[string]string{
				"servicetype":		t["servicetype"].(string),
				"servicegroupname": t["servicegroupname"].(string),
			},
			[]string{""},
			map[string]interface{}{},
		))
	}

	return nil
}

func (g *ServiceGroupGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createServiceGroup(client); err != nil {
		return err
	}
	return nil
}