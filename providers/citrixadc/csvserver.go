package citrixadc

import (
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type CsvServereGenerator struct {
	CitrixService
}

func (g *CsvServereGenerator) createCsvServer(client *service.NitroClient) error {
	cs, err := client.FindAllResources(service.Csvserver.Type())
	if err != nil {
		return err
	}
	for _, t := range cs {
		name := t["name"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			name,
			normalizeResourceName(name),
			"citrixadc_csvserver",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *CsvServereGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createCsvServer(client); err != nil {
		return err
	}
	return nil
}