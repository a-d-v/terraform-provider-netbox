package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDcimSite_basic(t *testing.T) {
	sitename := "testsite"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDcimSiteConfig(sitename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_dcim_site.test", "name", sitename),
				),
			},
		},
	})
}

func testAccDataSourceDcimSiteConfig(name string) string {
	return fmt.Sprintf(`
resource "netbox_dcim_site" "test" {
  name = "%s"
  slug = "%s"
}

data "netbox_dcim_site" "test" {
  id = netbox_dcim_site.test.id
}
`, name, name)
}
