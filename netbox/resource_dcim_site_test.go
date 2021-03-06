package netbox

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

func TestAccDcimSite_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDcimSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDcimSiteConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcimSiteExists("netbox_dcim_site.newsite"),
				),
			},
		},
	})
}

func testAccCheckDcimSiteDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_dcim_site" {
			continue
		}

		siteID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimSitesReadParams{
			Context: context.Background(),
			ID:      siteID,
		}

		resp, err := c.Dcim.DcimSitesRead(params, nil)
		if err != nil {
			if err.(*runtime.APIError).Code == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Site ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckDcimSiteExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Site ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		siteID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimSitesReadParams{
			Context: context.Background(),
			ID:      siteID,
		}

		_, err = c.Dcim.DcimSitesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckDcimSiteConfigBasic() string {
	return `
resource "netbox_dcim_site" "newsite" {
  name        = "NewSite"
  slug        = "newsitet"
}
`
}
