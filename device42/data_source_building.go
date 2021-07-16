package device42

import (
	"context"
	"fmt"
	"strconv"

	device42 "github.com/chopnico/device42-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBuilding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBuildingRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"id", "name"},
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"id", "name"},
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// get a building by id
func dataSourceBuildingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	var diags diag.Diagnostics
	var err error

	buildingId := d.Get("id").(int)
	buildingName := d.Get("name").(string)
	buildings := &[]device42.Building{}

	if buildingId != 0 {
		buildings, err = c.GetBuildingById(buildingId)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to get building with id " + strconv.Itoa(buildingId),
				Detail:   err.Error(),
			})
			return diags
		}
	} else if buildingName != "" {
		buildings, err = c.GetBuildingByName(buildingName)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to get building with name " + buildingName,
				Detail:   err.Error(),
			})
			return diags
		}
	}

	building := (*buildings)[0]

	c.WriteToDebugLog(fmt.Sprintf("%v", building))

	d.Set("name", building.Name)
	d.Set("address", building.Address)
	d.Set("notes", building.Notes)

	d.SetId(strconv.Itoa(buildingId))

	return diags
}
