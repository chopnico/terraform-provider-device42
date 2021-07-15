package device42

import (
	"context"
	"fmt"
	"strconv"

	device42 "github.com/chopnico/device42-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSubnet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSubnetRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mask_bits": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vrf_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

// get a building by id
func dataSourceSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	var diags diag.Diagnostics

	subnetId := d.Get("id").(int)

	subnet, err := c.GetSubnetById(subnetId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get subnet with id " + strconv.Itoa(subnetId),
			Detail:   err.Error(),
		})
		return diags
	}

	c.WriteToDebugLog(fmt.Sprintf("%v", subnet))

	d.Set("name", subnet.Name)
	d.Set("network", subnet.Network)
	d.Set("mask_bits", subnet.MaskBits)
	d.Set("vrf_group_id", subnet.VrfGroupID)

	d.SetId(strconv.Itoa(subnetId))

	return diags
}
