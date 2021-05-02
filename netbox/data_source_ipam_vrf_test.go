package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamVrf_basic(t *testing.T) {
	vrf := "test"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIpamVrfConfig(vrf),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ipam_vrf.test", "name", vrf),
				),
			},
		},
	})
}

func testAccDataSourceIpamVrfConfig(vrf string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_vrf" "test" {
  name = "%s"
}

data "netbox_ipam_vrf" "test" {
  name = "test"
}
`, vrf)
}
