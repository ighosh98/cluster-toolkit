/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

locals {
  # This label allows for billing report tracking based on module.
  labels = merge(var.labels, { ghpc_module = "pbspro-client", ghpc_role = "scheduler" })
}

locals {
  resource_prefix = var.name_prefix != null ? var.name_prefix : "${var.deployment_name}-client"

  user_startup_script_runners = var.startup_script == null ? [] : [
    {
      type        = "shell"
      content     = var.startup_script
      destination = "user_startup_script_pbs_server.sh"
    }
  ]
}

module "pbs_install" {
  source = "../../../../community/modules/scripts/pbspro-install"

  pbs_exec   = var.pbs_exec
  pbs_home   = var.pbs_home
  pbs_server = var.pbs_server

  pbs_role = "client"
  rpm_url  = var.pbs_client_rpm_url
}

module "client_startup_script" {
  source = "../../../../modules/scripts/startup-script"

  deployment_name = var.deployment_name
  project_id      = var.project_id
  region          = var.region
  labels          = local.labels

  runners = flatten([
    local.user_startup_script_runners,
    module.pbs_install.runner,
  ])
}

module "pbs_client" {
  source = "../../../../modules/compute/vm-instance"

  instance_count     = var.instance_count
  provisioning_model = var.spot ? "SPOT" : null

  deployment_name = var.deployment_name
  name_prefix     = local.resource_prefix
  project_id      = var.project_id
  region          = var.region
  zone            = var.zone
  labels          = local.labels

  machine_type          = var.machine_type
  service_account_email = var.service_account.email
  metadata              = var.metadata
  startup_script        = module.client_startup_script.startup_script
  enable_oslogin        = var.enable_oslogin

  instance_image        = var.instance_image
  disk_size_gb          = var.disk_size_gb
  disk_type             = var.disk_type
  auto_delete_boot_disk = var.auto_delete_boot_disk
  local_ssd_count       = var.local_ssd_count
  local_ssd_interface   = var.local_ssd_interface

  disable_public_ips   = !var.enable_public_ips
  network_self_link    = var.network_self_link
  subnetwork_self_link = var.subnetwork_self_link
  network_interfaces   = var.network_interfaces
  bandwidth_tier       = var.bandwidth_tier
  placement_policy     = var.placement_policy
  tags                 = var.tags

  guest_accelerator   = var.guest_accelerator
  on_host_maintenance = var.on_host_maintenance
  threads_per_core    = var.threads_per_core

  network_storage = var.network_storage

}
