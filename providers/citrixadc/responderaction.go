package citrixadc

import (
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ResponderActionGenerator struct {
	CitrixService
}

func (g *ResponderActionGenerator) createResponderAction(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Responderaction.Type())
	if err != nil {
		return err
	}
	for _, t := range server {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_responderaction",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *ResponderActionGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createResponderAction(client); err != nil {
		return err
	}
	return nil
}