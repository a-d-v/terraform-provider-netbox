package netbox

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamPrefixes_basic(t *testing.T) {
	prefix := fmt.Sprintf("192.168.%d.0/24", rand.Intn(255))

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIpamPrefixesConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ipam_prefixes.test", "results.0.prefix", prefix),
				),
			},
		},
	})
}

func TestAccDataSourceIpamPrefixes_Family(t *testing.T) {
	prefix := fmt.Sprintf("192.168.%d.0/24", rand.Intn(255))

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIpamPrefixesFamilyConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ipam_prefixes.familytest", "results.0.family.0.value", "4"),
				),
			},
		},
	})
}

func testAccDataSourceIpamPrefixesConfig(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_prefix" "test" {
  prefix = "%s"
  status = "active"
}

data "netbox_ipam_prefixes" "test" {
  prefix = netbox_ipam_prefix.test.prefix
  depends_on = [
    netbox_ipam_prefix.test,
  ]
}
`, prefix)
}

func testAccDataSourceIpamPrefixesFamilyConfig(prefix string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_prefix" "familytest" {
  prefix = "%s"
  status = "active"
}

data "netbox_ipam_prefixes" "familytest" {
  family = 4
  prefix = netbox_ipam_prefix.familytest.prefix
  limit = 500
  depends_on = [
    netbox_ipam_prefix.familytest,
  ]
}
`, prefix)
}
