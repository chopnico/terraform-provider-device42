resource "device42_vrf_group" "arnolds_room"{
  name = "Arnold's Room"
  description = "Arnold's room"
  building_ids = [
    resource.device42_building.sunset_arms.id,
  ]
}

resource "device42_vrf_group" "ps_118_history_class"{
  name = "P.S. 118 - History Class"
  description = "vrf group for the history department"
  building_ids = [
    resource.device42_building.ps_118.id,
  ]
}

data "device42_vrf_group" "arnolds_room" {
  name = resource.device42_vrf_group.arnolds_room.name
}

data "device42_vrf_group" "ps_118_history_class" {
  name = resource.device42_vrf_group.ps_118_history_class.name
}

output "arnolds_room_vrf_group_id" {
  value = data.device42_vrf_group.arnolds_room.id
}

output "ps_118_history_class_vrf_group_id" {
  value = data.device42_vrf_group.ps_118_history_class.id
}

