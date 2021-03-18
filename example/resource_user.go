package example

import (
	//"http://github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"https://github.com/hashicorp/terraform-plugin-sdk/tree/master/helper/schema"
)

func resourceItem() *schema.Resource{
	return &schema.Resource{
		Create:resourceUserAdd,
		Importer:&schema.ResourceImporter{
			State:resourceUSerImportState,
		}
		Schema:map[string]*schema.Schema{
			"email":&schema.Schema{
				Type:schema.TypeString,
				Description:"The email id of new user",
				Required:true,
			},
			"first_name":&schema.Schema{
				Type:schema.TypeString,
				Description:"first name of new user",
				Required:true,
			},
			"last_name":&schema.Schema{
				Type:schema.TypeString,
				Description:"last name of new user",
				Required:true,
			},
			

		}
	}

}

func resourceUSerImportState(d *schema.ResourceData,meta interface{}) ([]*schema.ResourceData,error){
	fmt.Println("hello world")
}