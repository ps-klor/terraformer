// Copyright 2022 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"log"
	"os"
	citrixadc_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/citrixadc"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdCitrixAdcImporter(options ImportOptions) *cobra.Command {
	var endpoint string
	cmd := &cobra.Command{
		Use:   "citrixadc",
		Short: "Import current state to Terraform configuration from CitrixADC",
		Long:  "Import current state to Terraform configuration from CitrixADC",
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint := os.Getenv("NS_URL")
			log.Println("Citrix ADC Endpoint: " + endpoint)
			provider := newCitrixAdcProvider()
			err := Import(provider, options, []string{endpoint})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newCitrixAdcProvider()))
	cmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "env param NS_URL")
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")

	return cmd
}

func newCitrixAdcProvider() terraformutils.ProviderGenerator {
	return &citrixadc_terraforming.Provider{}
}