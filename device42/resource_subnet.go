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

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSubnetSet,
		ReadContext:   resourceSubnetRead,
		UpdateContext: resourceSubnetSet,
		DeleteContext: resourceSubnetDelete,
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
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask_bits": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vrf_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSubnetSet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	log.Println(fmt.Sprintf("[DEBUG] subnet : %s", d.Get("subnets")))

	subnet, err := c.SetSubnet(&device42.Subnet{
		Name:    d.Get("name").(string),
		Network: d.Get("network").(string),
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create vrf group with name " + d.Get("name").(string),
			Detail:   err.Error(),
		})
	}

	log.Println(fmt.Sprintf("[DEBUG] subnet : %v", subnet))

	d.SetId(strconv.Itoa(subnet.SubnetID))

	resourceVrfGroupRead(ctx, d, m)

	return diags
}

func resourceSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	subnetId, err := strconv.Atoi(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read id",
			Detail:   err.Error(),
		})
		return diags
	}
	subnet, err := c.GetSubnetById(subnetId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get vrf group with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	log.Println(fmt.Sprintf("[DEBUG] subnet : %v", subnet))

	d.Set("name", subnet.Name)
	d.Set("network", subnet.Network)
	d.Set("mask_bits", subnet.MaskBits)
	d.Set("vrf_group_id", subnet.VrfGroupID)

	return diags
}

// delete vrf group
func resourceSubnetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.Api)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var id int
	_, err := fmt.Sscan(d.Id(), &id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get subnet id",
			Detail:   err.Error(),
		})
		return diags
	}

	err = c.DeleteSubnet(id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete subnet with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId("")

	return diags
}
