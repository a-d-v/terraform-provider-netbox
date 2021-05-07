# netbox_rir Resource

Creates a RIR.

## Example Usage

```hcl
resource "netbox_rir" "example" {
  name  = "RIPE"
  slug  = "RIPE"
}
```

## Argument Reference

* `name` - (Required) The name of the RIR.

* `slug` - (Optional) The slug name.


## Attribute Reference

* `id` - The ID of the RIR.
