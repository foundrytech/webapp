packer {
  required_plugins {
    googlecompute = {
      version = ">= 1.0.0, < 2.0.0"
      source  = "github.com/hashicorp/googlecompute"
    }
  }
}

variable "project_id" {
  type = string
}

variable "source_image_family" {
  type    = string
  default = "centos-stream-8"
}

variable "zone" {
  type    = string
  default = "us-west1-a"
}

variable "disk_size" {
  type    = number
  default = 20
}

variable "disk_type" {
  type    = string
  default = "pd-standard"
}

source "googlecompute" "csye6225-app-custom-image" {
  project_id              = var.project_id
  source_image_family     = var.source_image_family
  zone                    = var.zone
  disk_size               = var.disk_size
  disk_type               = var.disk_type
  image_name              = "csye6225-{{timestamp}}"
  image_family            = "csye6225-app-custom-image"
  image_storage_locations = ["us"]
  ssh_username            = "packer"
}


build {
  name = "packer"
  sources = [
    "source.googlecompute.csye6225-app-custom-image"
  ]

  provisioner "shell" {
    script = "updateOS.sh"
  }

  provisioner "shell" {
    script = "setupGolang.sh"
  }

  provisioner "shell" {
    script = "setupDB.sh"
  }

  provisioner "shell" {
    script = "addUser.sh"
  }

  provisioner "shell" {
    script = "setupApp.sh"
  }

  provisioner "file" {
    source      = "webapp.service"
    destination = "/etc/systemd/system/webapp.service"
  }
}
