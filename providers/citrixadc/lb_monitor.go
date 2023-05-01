package citrixadc

import (
	"log"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LbMonitorGenerator struct {
	CitrixService
}

func (g *LbMonitorGenerator) createLbMonitor(client *service.NitroClient) error {
	sg, err := client.FindAllResources(service.Lbmonitor.Type())
	if err != nil {
		return err
	}
	for _, t := range sg {
		monitorname := t["monitorname"].(string)
		g.Resources = append(g.Resources, terraformutils.NewResource(
			monitorname,
			monitorname,
			"citrixadc_lbmonitor",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{},
		))
	}

	return nil
}

func (g *LbMonitorGenerator) InitResources() error {
	log.Printf("creating lb_monitor")
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createLbMonitor(client); err != nil {
		return err
	}
	return nil
}