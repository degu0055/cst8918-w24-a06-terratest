# Define config variables
variable "labelPrefix" {
  type        = string
  description = "Your college username. This will form the beginning of various resource names."
}

variable "region" {
  type        = string
  default     = "canadacentral"
  description = "The Azure region where resources will be created."
}

variable "admin_username" {
  type        = string
  default     = "azureadmin"
  description = "The username for the local user account on the VM."
}

variable "subscription_id" {
  type        = string
  description = "The Azure subscription ID where resources will be deployed."
}

