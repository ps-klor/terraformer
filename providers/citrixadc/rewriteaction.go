package citrixadc

import (
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RewriteActionGenerator struct {
	CitrixService
}

func (g *RewriteActionGenerator) createRewriteAction(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Rewriteaction.Type())
	if err != nil {
		return err
	}
	for _, t := range server {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_rewriteaction",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *RewriteActionGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createRewriteAction(client); err != nil {
		return err
	}
	return nil
}