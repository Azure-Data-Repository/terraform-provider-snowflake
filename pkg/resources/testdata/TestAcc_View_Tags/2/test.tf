resource "snowflake_tag" "test_tag" {
  name     = var.tag1Name
  database = var.database
  schema   = var.schema
}

resource "snowflake_tag" "test_tag_2" {
  name     = var.tag2Name
  database = var.database
  schema   = var.schema
}

resource "snowflake_view" "test" {
  name        = var.name
  database    = var.database
  schema      = var.schema
  is_secure   = false
  or_replace  = false
  copy_grants = false
  statement   = var.statement

  tag {
    name     = snowflake_tag.test_tag_2.name
    schema   = var.schema
    database = var.database
    value    = "some_value"
  }
}
