resource "tanzu-mission-control_backup" "sample-full" {
  name                    = "full-cluster"
  management_cluster_name = "MGMT_CLS_NAME"
  provisioner_name        = "PROVISIONER_NAME"
  cluster_name            = "CLS_NAME"

  backup_scope = "FULL_CLUSTER"

  spec {
    backup_ttl = "2592000s"
    excluded_namespaces = [
      "app-01",
      "app-02",
      "app-03",
      "app-04"
    ]
    excluded_resources = [
      "secrets",
      "configmaps"
    ]

    storage_location = "TARGET_LOCATION_NAME"
  }
}

