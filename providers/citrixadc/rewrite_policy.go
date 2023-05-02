package citrixadc

import (
	"log"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RewritePolicyGenerator struct {
	CitrixService
}

func (g *RewritePolicyGenerator) createRewritePolicy(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Rewritepolicy.Type())
	if err != nil {
		return err
	}
	for _, t := range server {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_rewritepolicy",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *RewritePolicyGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createRewritePolicy(client); err != nil {
		return err
	}
	return nil
}