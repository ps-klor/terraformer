package citrixadc

import (
	"log"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LbvServerGenerator struct {
	CitrixService
}

func (g *LbvServerGenerator) createLbvServer(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Lbvserver.Type())
	if err != nil {
		return err
	}
	for _, t := range server {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_lbvserver",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *LbvServerGenerator) InitResources() error {
	log.Printf("creating lbv_server")
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createLbvServer(client); err != nil {
		return err
	}
	return nil
}