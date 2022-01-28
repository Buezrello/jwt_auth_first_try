package domain

type RolePermissions struct {
	rolePermissions map[string][]string
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"GetAllCustomers", "GetCustomer", "NewAccount", "Newtransaction"},
		"user":  {"GetCustomer", "NewTransaction"},
	}}
}
