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

func resourceVrfGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVrfGroupUpdate,
		ReadContext:   resourceVrfGroupRead,
		UpdateContext: resourceVrfGroupUpdate,
		DeleteContext: resourceVrfGroupDelete,
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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"buildings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVrfGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	log.Println(fmt.Sprintf("[DEBUG] buildings : %s", d.Get("buildings")))

	buildings := make([]string, len(d.Get("buildings").([]interface{})))

	for i, v := range d.Get("buildings").([]interface{}) {
		buildings[i] = fmt.Sprint(v)
	}
	vrfGroup, err := c.SetVrfGroup(&device42.VrfGroup{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Buildings:   buildings,
		Groups:      d.Get("groups").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create vrf group with name " + d.Get("name").(string),
			Detail:   err.Error(),
		})
	}

	log.Println(fmt.Sprintf("[DEBUG] vrf group : %v", vrfGroup))

	d.SetId(strconv.Itoa(vrfGroup.ID))

	resourceVrfGroupRead(ctx, d, m)

	return diags
}

func resourceVrfGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vrfGroupId, err := strconv.Atoi(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read id",
			Detail:   err.Error(),
		})
		return diags
	}
	vrfGroup, err := c.GetVrfGroupById(vrfGroupId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get vrf group with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	log.Println(fmt.Sprintf("[DEBUG] vrf group : %v", vrfGroup))

	d.Set("name", vrfGroup.Name)
	d.Set("description", vrfGroup.Description)
	d.Set("buildings", vrfGroup.Buildings)
	d.Set("groups", vrfGroup.Groups)

	return diags
}

// delete vrf group
func resourceVrfGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var id int
	_, err := fmt.Sscan(d.Id(), &id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get vrf group id",
			Detail:   err.Error(),
		})
		return diags
	}

	err = c.DeleteVrfGroup(id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete vrf group with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId("")

	return diags
}
