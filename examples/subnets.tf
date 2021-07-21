resource "device42_subnet" "arnolds_room" {
  name = "Arnold's Room"
  network = "172.168.200.0"
  mask_bits = "24"
  vrf_group_id = resource.device42_vrf_group.arnolds_room.id
}

resource "device42_subnet" "ps_118_history_class" {
  name = "P.S. 118 - History Class"
  network = "172.168.201.0"
  mask_bits = "24"
  vrf_group_id = resource.device42_vrf_group.ps_118_history_class.id
}

data "device42_subnet" "arnolds_room" {
  name = resource.device42_subnet.arnolds_room.name
  network = resource.device42_subnet.arnolds_room.network
}

data "device42_subnet" "ps_118_history_class" {
  name = resource.device42_subnet.ps_118_history_class.name
  network = resource.device42_subnet.ps_118_history_class.network
}

output "arnolds_room_subnet_id" {
  value = data.device42_subnet.arnolds_room.id
}

output "ps_118_history_class_subnet_id" {
  value = data.device42_subnet.ps_118_history_class.id
}
