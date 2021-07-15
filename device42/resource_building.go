package device42

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/chopnico/device42-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBuilding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBuildingUpdate,
		ReadContext:   resourceBuildingRead,
		UpdateContext: resourceBuildingUpdate,
		DeleteContext: resourceBuildingDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceBuildingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	buildings, err := c.SetBuilding(&device42.Building{
		Name:    d.Get("name").(string),
		Address: d.Get("name").(string),
		Notes:   d.Get("notes").(string),
	})

	building := (*buildings)[0]

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create building with name " + d.Get("name").(string),
			Detail:   err.Error(),
		})
	}

	log.Println(fmt.Sprintf("[DEBUG] building : %v", building))

	d.SetId(strconv.Itoa(building.BuildingID))

	resourceBuildingRead(ctx, d, m)

	return diags
}

func resourceBuildingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	buildingId, err := strconv.Atoi(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read building id",
			Detail:   err.Error(),
		})
		return diags
	}
	buildings, err := c.GetBuildingById(buildingId)
	building := (*buildings)[0]
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get building with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	log.Println(fmt.Sprintf("[DEBUG] building : %v", building))

	d.Set("name", building.Name)
	d.Set("address", building.Address)
	d.Set("notes", building.Notes)

	return diags
}

func resourceBuildingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var id int
	_, err := fmt.Sscan(d.Id(), &id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get building id",
			Detail:   err.Error(),
		})
		return diags
	}

	err = c.DeleteBuilding(id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete building with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId("")

	return diags
}
