terraform {
  required_providers {
    device42 = {
        version = "~> 0.1.0"
        source = "github.com/chopnico/device42"
    }
  }
}

provider "device42" {
  ignore_ssl = true
  username   = "${yamldecode(file("private.yml"))["username"]}"
  password   = "${yamldecode(file("private.yml"))["password"]}"
  host       = "${yamldecode(file("private.yml"))["host"]}"
  proxy      = "http://127.0.0.1:8080"
}
