package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpamVrf() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamVrfCreate,
		ReadContext:   resourceIpamVrfRead,
		UpdateContext: resourceIpamVrfUpdate,
		DeleteContext: resourceIpamVrfDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIpamVrfCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	params := &ipam.IpamVrfsCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableVRF{
		Name: &name,
	}

	resp, err := c.Ipam.IpamVrfsCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create vrf: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceIpamVrfRead(ctx, d, m)

	return diags
}

func resourceIpamVrfRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	vrfID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamVrfsReadParams{
		Context: ctx,
		ID:      vrfID,
	}

	resp, err := c.Ipam.IpamVrfsRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get vrf: %v", err)
	}
	d.Set("name", resp.Payload.Name)

	return diags
}

func resourceIpamVrfUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	vrfID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)

	params := &ipam.IpamVrfsPartialUpdateParams{
		Context: ctx,
		ID:      vrfID,
	}

	params.Data = &models.WritableVRF{
		Name: &name,
	}

	_, err = c.Ipam.IpamVrfsPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update vrf: %v", err)
	}

	return resourceIpamVrfRead(ctx, d, m)
}

func resourceIpamVrfDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	vrfID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &ipam.IpamVrfsDeleteParams{
		Context: ctx,
		ID:      vrfID,
	}

	_, err = c.Ipam.IpamVrfsDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete vrf: %v", err)
	}

	d.SetId("")

	return diags
}
