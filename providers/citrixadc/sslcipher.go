package citrixadc

import (
	"fmt"
	"strconv"
	service "github.com/citrix/adc-nitro-go/service"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SslCipherGenerator struct {
	CitrixService
}

func (g *SslCipherGenerator) InitResources() error {
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createSslCipher(client); err != nil {
		return err
	}
	return nil
}

func (g *SslCipherGenerator) createSslCipher(client *service.NitroClient) error {
	sg, err := client.FindAllResources(service.Sslcipher.Type())
	if err != nil {
		return err
	}
	for _, t := range sg {
		ciphergroupname := t["ciphergroupname"].(string)
		findParams := service.FindParams{
			ResourceType: service.Sslcipher_sslciphersuite_binding.Type(),
			ResourceName: ciphergroupname,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		bindings := make([]interface{}, len(dataArr))
		for i, _ := range dataArr {
			bindings[i] = make(map[string]interface{})
			bindings[i].(map[string]interface{})["ciphername"] = dataArr[i]["ciphername"].(string)
			bindings[i].(map[string]interface{})["cipherpriority"], _ = strconv.Atoi(dataArr[i]["cipherpriority"].(string))
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			ciphergroupname,
			normalizeResourceName(ciphergroupname),
			"citrixadc_sslcipher",
			g.ProviderName,
			map[string]string{},
			[]string{""},
			map[string]interface{}{
				"ciphersuitebinding": bindings,
			},
		))
	}

	return nil
}