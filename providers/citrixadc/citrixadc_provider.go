package citrixadc

import (
	"errors"
	"os"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type Provider struct {
	terraformutils.Provider
	username string
	password string
	endpoint string
	insecureSkipVerify bool
}

func (p *Provider) Init(args []string) error {
	if username := os.Getenv("NS_LOGIN"); username != "" {
		p.username = os.Getenv("NS_LOGIN")
	}
	if password := os.Getenv("NS_PASSWORD"); password != "" {
		p.password = os.Getenv("NS_PASSWORD")
	}
	if endpoint := os.Getenv("NS_URL"); endpoint != "" {
		p.endpoint = os.Getenv("NS_URL")
	}
	if sslVerify := os.Getenv("NS_SSLVERIFY"); sslVerify != "" {
		p.insecureSkipVerify = (sslVerify == "false")
	}
	return nil
}

func (p *Provider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"username": cty.StringVal(p.username),
		"password": cty.StringVal(p.password),
		"endpoint": cty.StringVal(p.endpoint),
		"insecure_skip_verify": cty.BoolVal(p.insecureSkipVerify),
	})
}

func (p *Provider) GetName() string {
	return "citrixadc"
}

func (p *Provider) InitService(serviceName string, verbose bool) error {
	if service, isSupported := p.GetSupportedService()[serviceName]; isSupported {
		p.Service = service
		p.Service.SetName(serviceName)
		p.Service.SetVerbose(verbose)
		p.Service.SetProviderName(p.GetName())
		p.Service.SetArgs(map[string]interface{}{
			"username": p.username,
			"password": p.password,
			"endpoint": p.endpoint,
		})
		return nil
	}
	return errors.New(p.GetName() + ": " + serviceName + " not supported service")
}

func (p *Provider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"servicegroup": &ServiceGroupGenerator{},
		"cspolicy": &CsPolicyGenerator{},
		"lbmonitor": &LbMonitorGenerator{},
		"servicegroup_lbmonitor_binding": &ServiceGroupLbMonitorBindingGenerator{},
		"server": &ServerGenerator{},
		"servicegroup_servicegroupmember_binding": &ServiceGroupServiceGroupMemberBindingGenerator{},
		"lbvserver": &LbvServerGenerator{},
		"rewritepolicy": &RewritePolicyGenerator{},
		"lbvserver_rewritepolicy_binding": &LbserverRewritePolicyBindingGenerator{},
		"service": &ServiceGenerator{},
		"lbvserver_service_binding": &LbserverServiceBindingGenerator{},
		"csvserver": &CsvServereGenerator{},
		"csvserver_cspolicy_binding": &CsvServerCsPolicyBindingGenerator{},
		"responderpolicy": &ResponderPolicyGenerator{},
		"lbvserver_responderpolicy_binding": &LbserverResponderPolicyBindingGenerator{},
		"lbvserver_servicegroup_binding": &LbserverServiceGroupBindingGenerator{},
		"sslcipher": &SslCipherGenerator{},
		"rewriteaction": &RewriteActionGenerator{},
		"sslvserver": &SslvServerGenerator{},
		"responderaction": &ResponderActionGenerator{},
	}
}

func (p Provider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p Provider) GetProviderData(_ ...string) map[string]interface{} {
	return map[string]interface{}{}
}