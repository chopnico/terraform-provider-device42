resource "device42_dynamic_subnet" "arnolds_room_security"{
  name = "Arnold's Room - Security"
  parent_subnet_id = resource.device42_subnet.arnolds_room.id
  mask_bits = 29
}

resource "device42_dynamic_subnet" "ps_118_history_class_security"{
  name = "P.S. 118 - History Class - Security"
  parent_subnet_id = resource.device42_subnet.ps_118_history_class.id
  mask_bits = 29
}

data "device42_subnet" "arnolds_room_security" {
  name = resource.device42_dynamic_subnet.arnolds_room_security.name
  network = resource.device42_subnet.arnolds_room.network
}

data "device42_subnet" "ps_118_history_class_security" {
  name = resource.device42_dynamic_subnet.ps_118_history_class_security.name
  network = resource.device42_subnet.ps_118_history_class.network
}

output "arnolds_room_security_subnet_id" {
  value = data.device42_subnet.arnolds_room_security.id
}

output "ps_118_history_class_security_subnet_id" {
  value = data.device42_subnet.ps_118_history_class_security.id
}
