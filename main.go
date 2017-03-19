package main

import (
	"github.com/AlexsJones/terraform-provider-template/template"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: template.Provider,
	})
}
