---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "device42_dynamic_subnet Resource - terraform-provider-device42"
subcategory: ""
description: |-
  device42_dynamic_subnet resource can be used to generate a new subnet.
---

# device42_dynamic_subnet (Resource)

`device42_dynamic_subnet` resource can be used to generate a new subnet.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **mask_bits** (Number) The `mask_bits` of the dynamic subnet.
- **parent_subnet_id** (Number) The `parent_subnet_id` of the dynamic subnet.

### Optional

- **id** (String) The ID of this resource.
- **name** (String) The `name` of the dynamic subnet.
- **tags** (List of String) The `tags` of the dynamic subnet.

### Read-Only

- **last_updated** (String) The last time this resource was updated.
- **network** (String) The `network` of this dynamic subnet.
- **vrf_group_id** (Number) The `vrf_group_id` of this dynamic subnet.

