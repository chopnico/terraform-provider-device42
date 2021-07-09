resource "device42_vrf_group" "arnolds_room"{
  name = "arnold"
  description = "arnold's room"
  buildings = ["6100"]
  groups = "boarders:no"
}

data "device42_vrf_group" "test" {
  id = 1
}
