package citrixadc

import (
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ResponderPolicyGenerator struct {
	CitrixService
}

func (g *ResponderPolicyGenerator) createResponderPolicy(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Responderpolicy.Type())
	if err != nil {
		return err
	}
	for _, t := range server {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_responderpolicy",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *ResponderPolicyGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createResponderPolicy(client); err != nil {
		return err
	}
	return nil
}