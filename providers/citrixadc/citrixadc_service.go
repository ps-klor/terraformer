package citrixadc

import (
	"fmt"
	"log"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	service "github.com/citrix/adc-nitro-go/service"
)

type CitrixService struct {
	terraformutils.Service
}

func (s *CitrixService) createClient() (*service.NitroClient, error) {
	client, err := service.NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
		return nil, fmt.Errorf("unable to initialize Citrix client: %v", err)
	}
	client.Login()
	return client, nil
}