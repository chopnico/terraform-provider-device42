package device42

import (
	"context"
	"strconv"
	"time"

	device42 "github.com/chopnico/device42-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVrfGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVrfGroupsRead,
		Schema: map[string]*schema.Schema{
			"vrf_groups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
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
				},
			},
		},
	}
}

// get vrf groups
func dataSourceVrfGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	var diags diag.Diagnostics

	vrfGroups, err := c.GetVrfGroups()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get a list of vrf groups",
			Detail:   err.Error(),
		})
		return diags
	}

	vgs := flattenVrfGroupsData(vrfGroups)
	if err := d.Set("vrf_groups", vgs); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to set vrf groups",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

// flatten vrf groups to a map
func flattenVrfGroupsData(vrfGroups *[]device42.VrfGroup) []interface{} {
	if vrfGroups != nil {
		vgs := make([]interface{}, len(*vrfGroups), len(*vrfGroups))

		for i, vrfGroup := range *vrfGroups {
			vg := make(map[string]interface{})

			vg["id"] = vrfGroup.ID
			vg["name"] = vrfGroup.Name
			vg["description"] = vrfGroup.Description
			vg["buildings"] = vrfGroup.Buildings
			vg["groups"] = vrfGroup.Groups

			vgs[i] = vg
		}

		return vgs
	}

	return make([]interface{}, 0)
}
