resource "device42_building" "sunset_arms" {
  name = "Sunset Arms"
  address = "4040 Vine Street"
  notes = "Arnold's house"
}

resource "device42_building" "ps_118" {
  name = "P.S. 118"
  address = "unknown"
  notes = "Arnold's school"
}

data "device42_building" "sunset_arms" {
  name = resource.device42_building.sunset_arms.name
}

data "device42_building" "ps_118" {
  name = resource.device42_building.ps_118.name
}

output "sunset_arms_building_id" {
  value = data.device42_building.sunset_arms.id
}

output "ps_118_building_id" {
  value = data.device42_building.ps_118.id
}
