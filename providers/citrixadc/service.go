package citrixadc

import (
	"log"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGenerator struct {
	CitrixService
}

func (g *ServiceGenerator) createService(client *service.NitroClient) error {
	services, err := client.FindAllResources(service.Service.Type())
	if err != nil {
		return err
	}
	for _, t := range services {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_service",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *ServiceGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createService(client); err != nil {
		return err
	}
	return nil
}