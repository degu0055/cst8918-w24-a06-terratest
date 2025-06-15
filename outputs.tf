output "resource_group_name" {
  value = azurerm_resource_group.rg.name
}

output "vm_name" {
  value = azurerm_linux_virtual_machine.webserver.name
}

output "nic_name" {
  value = azurerm_network_interface.webserver.name
}

output "public_ip" {
  value = azurerm_public_ip.webserver.ip_address
}

output "admin_ssh_key" {
  value       = file("/Users/romeodeguzmanii/.ssh/id_rsa.pub")
  description = "SSH public key used for VM admin access"
}
