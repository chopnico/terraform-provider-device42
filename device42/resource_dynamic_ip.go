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

func resourceDynamicIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynamicIPSet,
		ReadContext:   resourceDynamicIPRead,
		UpdateContext: resourceDynamicIPSet,
		DeleteContext: resourceDynamicIPDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask_bits": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"subnet": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"subnet_id", "vrf_group", "vrf_group_id"},
			},
			"subnet_id": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"subnet", "vrf_group", "vrf_group_id"},
			},
			"vrf_group": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"subnet", "subnet_id", "vrf_group_id"},
			},
			"vrf_group_id": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"subnet", "subnet_id", "vrf_group"},
			},
		},
	}
}

func resourceDynamicIPSet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.API)

	var diags diag.Diagnostics
	var err error

	log.Println(fmt.Sprintf("[DEBUG] ip : %s", d.Get("ip")))

	ipID := d.Get("id").(string)
	ipMaskBits := d.Get("mask_bits").(int)
	ipSubnet := d.Get("subnet").(string)
	ipSubnetID := d.Get("subnet_id").(int)
	ipVRFGroup := d.Get("vrf_group").(string)
	ipVRFGroupID := d.Get("vrf_group_id").(int)
	ipLabel := d.Get("label").(string)

	ip := &device42.IP{}

	if ipID == "" {
		if ipSubnetID != 0 {
			ip, err = c.SuggestIPWithSubnetID(ipSubnetID, ipMaskBits, true)
		} else if ipSubnet != "" {
			ip, err = c.SuggestIPWithSubnet(ipSubnet, ipMaskBits, true)
		} else if ipVRFGroupID != 0 {
			ip, err = c.SuggestIPWithVRFGroupID(ipVRFGroupID, ipMaskBits, true)
		} else if ipVRFGroup != "" {
			ip, err = c.SuggestIPWithVRFGroup(ipVRFGroup, ipMaskBits, true)
		}

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to suggest ip",
				Detail:   err.Error(),
			})
			return diags
		}
	}

	ip.Label = ipLabel

	log.Println(fmt.Sprintf("[DEBUG] ip : %v", ip))
	ip, err = c.SetIP(ip)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to suggest ip",
			Detail:   err.Error(),
		})
		return diags
	}

	log.Println(fmt.Sprintf("[DEBUG] ip : %v", ip))

	_ = d.Set("label", ip.Label)
	_ = d.Set("address", ip.Address)
	_ = d.Set("subnet_id", ip.SubnetID)
	_ = d.Set("vrf_group_id", ipVRFGroupID)
	_ = d.Set("vrf_group", ip.VRFGroup)

	d.SetId(strconv.Itoa(ip.ID))

	resourceDynamicIPRead(ctx, d, m)

	return diags
}

func resourceDynamicIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.API)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ipID, err := strconv.Atoi(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read id",
			Detail:   err.Error(),
		})
		return diags
	}
	if ipID == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "ip id is 0",
			Detail:   "the current id of this ip 0. not sure why.",
		})
		return diags
	}
	ip, err := c.GetIPByID(ipID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get ip with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	log.Println(fmt.Sprintf("[DEBUG] ip : %v", ip))

	_ = d.Set("label", ip.Label)
	_ = d.Set("address", ip.Address)
	_ = d.Set("subnet", ip.Subnet)
	_ = d.Set("subnet_id", ip.SubnetID)
	_ = d.Set("vrf_group", ip.VRFGroup)

	return diags
}

// delete ip
func resourceDynamicIPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*device42.API)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var id int
	_, err := fmt.Sscan(d.Id(), &id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get ip id",
			Detail:   err.Error(),
		})
		return diags
	}

	err = c.DeleteIP(id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete ip with id " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId("")

	return diags
}
