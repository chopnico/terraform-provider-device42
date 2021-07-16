resource "device42_building" "sunset_arms" {
  name = "Sunset Arms"
  address = "4040 Vine Street"
  notes = "arnold's house"
}

resource "device42_building" "ps_118" {
  name = "P.S. 118"
  address = "unknown"
  notes = "arnold's school"
}

data "device42_building" "ps_118" {
  name = "P.S. 118"
}

resource "device42_vrf_group" "arnolds_room"{
  name = "arnold's room"
  description = "arnold's room"
  building_ids = [
    data.device42_building.ps_118.id
  ]
}

data "device42_vrf_group" "test" {
  id = 1
}
