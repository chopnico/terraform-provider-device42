package device42

import (
	"context"
	"fmt"
	"strconv"

	device42 "github.com/chopnico/device42-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVrfGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVrfGroupRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"buildings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// get a vrf group by id
func dataSourceVrfGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	var diags diag.Diagnostics

	vrfGroupId := d.Get("id").(int)

	vrfGroup, err := c.GetVrfGroupById(vrfGroupId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get vrf group with id " + strconv.Itoa(vrfGroupId),
			Detail:   err.Error(),
		})
		return diags
	}

	c.WriteToDebugLog(fmt.Sprintf("%v", vrfGroup))

	d.Set("name", vrfGroup.Name)
	d.Set("description", vrfGroup.Description)
	d.Set("buildings", vrfGroup.Buildings)
	d.Set("groups", vrfGroup.Groups)

	d.SetId(strconv.Itoa(vrfGroupId))

	return diags
}
