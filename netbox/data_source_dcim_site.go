package netbox

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func dataSourceDcimSite() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDcimSiteRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tenant": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"custom_fields": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceDcimSiteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &dcim.DcimSitesReadParams{
		Context: ctx,
		ID:      int64(d.Get("id").(int)),
	}

	resp, err := c.Dcim.DcimSitesRead(params, nil)
	if err != nil {
		return diag.Errorf("Unable to get site: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))
	d.Set("name", resp.Payload.Name)
	d.Set("tenant", flattenDcimSiteTenant(resp.Payload.Tenant))
	d.Set("status", flattenDcimSiteStatus(resp.Payload.Status))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func flattenDcimSiteTenant(input *models.NestedTenant) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["id"] = input.ID
	values["name"] = input.Name
	values["slug"] = input.Slug

	return []interface{}{values}
}

func flattenDcimSiteStatus(input *models.SiteStatus) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["label"] = input.Label
	values["value"] = input.Value

	return []interface{}{values}
}
