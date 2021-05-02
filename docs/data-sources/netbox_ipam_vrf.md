# netbox_ipam_vrf Data Source

Use this data source to get information about a vrf.

## Example Usage

```hcl
data "netbox_ipam_vrf" "example" {
  name = "test"
}
```

## Argument Reference

* `name` - (Required) - The name of the VRF

## Attribute Reference

* `results` - One or more `results` blocks as defined below.

The `results` block contains:

* `id` - The ID of the vrf.

* `name` - The name of the vrf.
