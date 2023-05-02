package citrixadc

import (
	"log"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type CsPolicyGenerator struct {
	CitrixService
}

func (g *CsPolicyGenerator) createCsPolicy(client *service.NitroClient) error {
	cp, err := client.FindAllResources(service.Cspolicy.Type())
	if err != nil {
		return err
	}
	for _, t := range cp {
		policy := t["policyname"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			policy,
			normalizeResourceName(policy),
			"citrixadc_cspolicy",
			g.ProviderName,
			map[string]string{
				"policyname": policy,
			},
			[]string{""},
			map[string]interface{}{},
		))
	}

	return nil
}

func (g *CsPolicyGenerator) InitResources() error {
	log.Printf("creating cs_policy")
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createCsPolicy(client); err != nil {
		return err
	}
	return nil
}