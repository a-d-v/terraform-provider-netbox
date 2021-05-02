package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamAggregates_basic(t *testing.T) {
	prefix := "10.0.0.0/8"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIpamAggregatesConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ipam_aggregates.test", "prefix", prefix),
				),
			},
		},
	})
}

func testAccDataSourceIpamAggregatesConfig(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_rir" "test" {
	name        = "Test"
	slug        = "test"
}
resource "netbox_ipam_aggregates" "test" {
	prefix      = "%s"
	rir_id      = netbox_ipam_rir.test.id
}
data "netbox_ipam_aggregates" "test" {
  family = 4
  prefix = "%s"
  limit = 500
}
`, prefix, prefix)
}
