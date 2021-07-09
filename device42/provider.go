package device42

import (
	"context"

	device42 "github.com/chopnico/device42-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEVICE42_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DEVICE42_PASSWORD", nil),
			},
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEVICE42_HOST", nil),
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"device42_vrf_group": resourceVrfGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"device42_vrf_groups": dataSourceVrfGroups(),
			"device42_vrf_group":  dataSourceVrfGroup(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	host := d.Get("host").(string)
	ignoreSsl := d.Get("ignore_ssl").(bool)
	proxy := d.Get("proxy").(string)

	var diags diag.Diagnostics

	if (username != "") && (password != "") && (host != "") {
		c, err := device42.NewApiBasicAuth(username, password, host)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		c.Proxy(proxy)

		if ignoreSsl {
			c.IgnoreSslErrors()
		}

		return c, diags
	} else {
		return nil, diag.Errorf("you must provide a username, a password, and the host of the device42 appliance")
	}
}
