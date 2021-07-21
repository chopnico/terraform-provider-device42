resource "device42_dynamic_ip" "arnolds_room_security_camera_1"{
  label = "Arnold's Room - Security - Camera 1"
  subnet_id = data.device42_subnet.arnolds_room_security.id
  mask_bits = 29
}

resource "device42_dynamic_ip" "ps_118_history_class_security_camera_1"{
  label = "P.S. 118 - History Class - Security - Camera 1"
  subnet_id = data.device42_subnet.ps_118_history_class_security.id
  mask_bits = 29
}

data "device42_ip" "arnolds_room_security_camera_1" {
  id = resource.device42_dynamic_ip.arnolds_room_security_camera_1.id
}

data "device42_ip" "ps_118_history_class_security_camera_1" {
  id = resource.device42_dynamic_ip.ps_118_history_class_security_camera_1.id
}

output "arnolds_room_security_camera_1_id" {
  value = data.device42_ip.arnolds_room_security_camera_1.id
}

output "ps_118_history_class_security_camera_id" {
  value = data.device42_ip.ps_118_history_class_security_camera_1.id
}

output "arnolds_room_security_camera_1_ip_address" {
  value = data.device42_ip.arnolds_room_security_camera_1.address
}

output "ps_118_history_class_security_camera_ip_address" {
  value = data.device42_ip.ps_118_history_class_security_camera_1.address
}
