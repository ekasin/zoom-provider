package example

import (
	//"http://github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"https://github.com/hashicorp/terraform-plugin-sdk/tree/master/helper/schema"
	"https://github.com/hashicorp/terraform-plugin-sdk/tree/master/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:map[string]*schema.Schema{
			"url_address":&schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				//DefaultFunc: schema.EnvDefaultFunc("URL_ADDRESS", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				//DefaultFunc: schema.EnvDefaultFunc("SERVICE_TOKEN", ""),
			},

		},
		ResourcesMap:map[string]*schema.Resource{
			"example_user":resourceItem()
		},
		//ConfigureFunc: providerConfigure,

	}
}

/*func providerConfigure(d *schema.ResourceData) (interface{},error){
	 address:=d.Get("url_address").(string)
	 token:=d.Get("token").(string)
     url :=address+"?access_token="+token
	 config:=&Config{
		 finalUrlAddress:=finalAddress
	 }
	 return config,nil
}*/