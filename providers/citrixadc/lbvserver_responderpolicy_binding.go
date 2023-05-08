package citrixadc

import (
	"log"
	"fmt"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LbserverResponderPolicyBindingGenerator struct {
	CitrixService
}

func (g *LbserverResponderPolicyBindingGenerator) createLbserverResponderPolicyBindingGenerator(client *service.NitroClient) error {
	server, err := client.FindAllResources(service.Lbvserver.Type())
	if err != nil {
		return err
	}
	for _, t := range server {
		name := t["name"].(string)
		findParams := service.FindParams{
			ResourceType: service.Lbvserver_responderpolicy_binding.Type(),
			ResourceName: name,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if len(dataArr) == 0 {
			continue
		}
		for _, t := range dataArr {
			policyname := t["policyname"].(string)
			id := fmt.Sprintf("%s,%s", name, policyname)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				id,
				normalizeResourceName(id),
				"citrixadc_lbvserver_responderpolicy_binding",
				g.ProviderName,
				map[string]string{},
				[]string{""},
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *LbserverResponderPolicyBindingGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createLbserverResponderPolicyBindingGenerator(client); err != nil {
		return err
	}
	return nil
}