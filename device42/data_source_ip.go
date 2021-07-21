package device42

import (
	"context"
	"fmt"
	"log"
	"strconv"

	device42 "github.com/chopnico/device42-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:         schema.TypeString,
				Computed:     true,
				RequiredWith: []string{"subnet_id"},
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": &schema.Schema{
				Type:         schema.TypeInt,
				RequiredWith: []string{"address"},
				Computed:     true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// get a building by id
func dataSourceIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.API)

	var diags diag.Diagnostics
	var err error

	ipID := d.Get("id").(int)
	ipAddress := d.Get("address").(string)
	ipSubnetID := d.Get("subnet_id").(int)
	ip := &device42.IP{}

	if ipID != 0 {
		log.Printf("[DEBUG] ip id: %d\n", ipID)

		ip, err = c.GetIPByID(ipID)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to get ip with id " + strconv.Itoa(ipID),
				Detail:   err.Error(),
			})
			return diags
		}
	} else if ipAddress != "" && ipSubnetID != 0 {
		log.Printf("[DEBUG] ip address: %s\n", ipAddress)

		ip, err = c.GetIPByAddressWithSubnetID(ipAddress, ipSubnetID)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to get subnet with address " + ipAddress + "and subnet id " + strconv.Itoa(ipSubnetID),
				Detail:   err.Error(),
			})
			return diags
		}
	}

	c.WriteToDebugLog(fmt.Sprintf("ip : %v", ip))

	_ = d.Set("address", ip.Address)
	_ = d.Set("subnet", ip.Subnet)
	_ = d.Set("subnet_id", ip.SubnetID)
	_ = d.Set("label", ip.Label)
	_ = d.Set("mac_address", ip.MacAddress)

	d.SetId(strconv.Itoa(ip.ID))

	return diags
}
