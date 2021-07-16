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
			"building_ids": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
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

	buildings := make([]int, len(d.Get("building_ids").([]interface{})))

	for i, v := range d.Get("building_ids").([]interface{}) {
		b, err := c.GetBuildingById(v.(int))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to get building with id " + strconv.Itoa(v.(int)),
				Detail:   err.Error(),
			})
			return diags
		}
		buildings[i] = (*b)[0].BuildingID
	}

	d.Set("name", vrfGroup.Name)
	d.Set("description", vrfGroup.Description)
	d.Set("building_ids", buildings)

	d.SetId(strconv.Itoa(vrfGroupId))

	return diags
}
