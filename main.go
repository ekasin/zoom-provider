package main
import (
	//"fmt"
	//"github.com/hashicorp/terraform/tree/main/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	//"https://github.com/hashicorp/terraform-plugin-sdk/tree/master/plugin"
	//"https://github.com/ekasin/eGit/tree/master"

	"C:\Program Files\Go\src\terraform-provider-example\example"
)

func main(){
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc:master.Provider,
	})
}