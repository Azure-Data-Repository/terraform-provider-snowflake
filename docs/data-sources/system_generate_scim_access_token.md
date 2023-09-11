---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "snowflake_system_generate_scim_access_token Data Source - terraform-provider-snowflake"
subcategory: ""
description: |-
  
---

# snowflake_system_generate_scim_access_token (Data Source)



## Example Usage

```terraform
data "snowflake_system_generate_scim_access_token" "scim" {
  integration_name = "AAD_PROVISIONING"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `integration_name` (String) SCIM Integration Name

### Read-Only

- `access_token` (String) SCIM Access Token
- `id` (String) The ID of this resource.

