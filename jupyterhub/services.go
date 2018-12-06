package jupyterhub

// GetInfo returns the Hub's system information.
func GetServices() (services []map[string]interface{}, err error) {
	_, err = getResult("/services", services)
	return services, err
}
