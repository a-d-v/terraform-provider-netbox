package netbox

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDcimSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDcimSiteCreate,
		ReadContext:   resourceDcimSiteRead,
		UpdateContext: resourceDcimSiteUpdate,
		DeleteContext: resourceDcimSiteDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDcimSiteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimSitesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableSite{
		Name: &name,
		Slug: &slug,
	}

	resp, err := c.Dcim.DcimSitesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create site: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceDcimSiteRead(ctx, d, meta)

	return diags
}

func resourceDcimSiteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	siteID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimSitesReadParams{
		Context: ctx,
		ID:      siteID,
	}

	resp, err := c.Dcim.DcimSitesRead(params, nil)
	if err != nil {
		if err.(*runtime.APIError).Code == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get site: %v", err)
	}
	d.Set("slug", resp.Payload.Slug)

	return diags
}

func resourceDcimSiteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.NetBoxAPI)

	siteID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	params := &dcim.DcimSitesPartialUpdateParams{
		Context: ctx,
		ID:      siteID,
	}

	params.Data = &models.WritableSite{
		Name: &name,
		Slug: &slug,
	}

	_, err = c.Dcim.DcimSitesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update site: %v", err)
	}

	return resourceDcimSiteRead(ctx, d, meta)
}

func resourceDcimSiteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	siteID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimSitesDeleteParams{
		Context: ctx,
		ID:      siteID,
	}

	_, err = c.Dcim.DcimSitesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete site: %v", err)
	}

	d.SetId("")

	return diags
}
