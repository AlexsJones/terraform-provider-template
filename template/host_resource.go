package template

/*********************************************************************************
*     File Name           :     engineer/host_resource.go
*     Created By          :     jonesax
*     Creation Date       :     [2017-03-17 16:14]
*     Last Modified       :     [2017-03-19 11:38]
*     Description         :
**********************************************************************************/
import (
	"github.com/hashicorp/terraform/helper/schema"
)

func hostResource() *schema.Resource {

	return &schema.Resource{
		Create: hostResourceCreate,
		Read:   hostResourceRead,
		Update: nil,
		Delete: hostResourceDelete,
		//Expected parameters for this Resource
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

// Resource CRUD

func hostResourceCreate(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func hostResourceRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func hostResourceUpdate(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func hostResourceDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}
