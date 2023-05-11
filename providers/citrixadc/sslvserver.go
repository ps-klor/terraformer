package citrixadc

import (
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SslvServerGenerator struct {
	CitrixService
}

func (g *SslvServerGenerator) createSslvServer(client *service.NitroClient) error {
	services, err := client.FindAllResources(service.Sslvserver.Type())
	if err != nil {
		return err
	}
	for _, t := range services {
		vservername := t["vservername"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			vservername,
			normalizeResourceName(vservername),
			"citrixadc_sslvserver",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *SslvServerGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createSslvServer(client); err != nil {
		return err
	}
	return nil
}