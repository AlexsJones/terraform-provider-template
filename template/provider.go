package template

/*********************************************************************************
*     File Name           :     provider.go
*     Created By          :     jonesax
*     Creation Date       :     [2017-03-17 14:59]
*     Last Modified       :     [2017-03-19 11:40]
*     Description         :
**********************************************************************************/
import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "vagrant",
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "vagrant",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"template_host": hostResource(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {

	return nil, nil
}
