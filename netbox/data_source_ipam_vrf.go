package netbox

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

func dataSourceIpamVrf() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamVrfRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"results": {
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
					},
				},
			},
		},
	}
}

func dataSourceIpamVrfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	params := &ipam.IpamVrfsListParams{
		Context: ctx,
	}

	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		params.Name = &name
	}


	resp, err := c.Ipam.IpamVrfsList(params, nil)
	if err != nil {
		return diag.Errorf("Unable to get Vrfs: %v", err)
	}

	//lintignore:R017
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("results", flattenIpamVrfResults(resp.Payload.Results))

	return diags
}

func flattenIpamVrfResults(input []*models.VRF) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["id"] = item.ID
		values["name"] = item.Name

		result = append(result, values)
	}

	return result
}
