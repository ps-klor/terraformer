package citrixadc

import (
	"fmt"
	"log"
	service "github.com/citrix/adc-nitro-go/service"
)

type ServiceGroupGenerator struct {
	CitrixService
}

func (g *ServiceGroupGenerator) createServiceGroup(client *service.NitroClient) error {
	sg, err := client.FindAllResources(service.Lbvserver.Type())
	if err != nil {
		return err
	}
	return nil
}

func (g *ServiceGroupGenerator) InitResources() error {
	log.Printf("creating service_group")
	client, err := g.createClient()
	if err != nil {
		return err
	}
	if err := g.createServiceGroup(client); err != nil {
		return err
	}
	return nil
}