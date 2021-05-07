# netbox_site Resource

Creates a Site.

## Example Usage

```hcl
resource "netbox_site" "testsite" {
  name  = "Mysite"
  slug  = "mysite"
}
```

## Argument Reference

* `name` - (Required) The name of the site.

* `slug` - (Optional) The slug name.


## Attribute Reference

* `id` - The ID of the Site.
