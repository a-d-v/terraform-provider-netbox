# netbox_dcim_site Data Source

Use this data source to get information about a site.

## Example Usage

```hcl
data "netbox_dcim_site" "example" {
  site_id = 123
}
```

## Argument Reference

* `id` - (Required) The ID of the site.

## Attribute Reference

* `name` - The name of the site.

* `tenant` - A `tenant` block as defined below.

* `status` - A `status` block as defined below.

* `custom_fields` - A mapping of custom fields for the site.

The `tenant` block contains:

* `id` - The ID of the tenant.

* `name` - The name of the tenant.

* `slug` - The tenant slug.

The `status` block contains:

* `value` - A value for the operational status of the site.

* `label` - A label for the operational status of the site.
