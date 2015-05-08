package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)
//Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
            "url": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("COUCHDB_URL", nil),
                Description: "The URL for couchdb.",
            },
		},

		ResourcesMap: map[string]*schema.Resource{
			"couchdb_database": resourceCouchDB(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
        URL: d.Get("url").(string),
	}

	return config.Client()
}
