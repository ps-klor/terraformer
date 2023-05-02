package citrixadc

import (
	"log"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServerGenerator struct {
	CitrixService
}

func (g *ServerGenerator) createServer(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Server.Type())
	if err != nil {
		return err
	}
	for _, t := range server {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_server",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *ServerGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createServer(client); err != nil {
		return err
	}
	return nil
}