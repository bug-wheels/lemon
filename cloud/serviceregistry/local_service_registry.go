package serviceregistry

import (
	"errors"
	"lemon/cloud"
)

type localServiceRegistry struct {
	ServiceInstances     map[string]map[string]cloud.ServiceInstance
	localServiceInstance cloud.ServiceInstance
}

func NewLocalServiceRegistry() *localServiceRegistry {
	return &localServiceRegistry{ServiceInstances: map[string]map[string]cloud.ServiceInstance{}}
}

func (d localServiceRegistry) Description() string {
	return "localServiceRegistry"
}

func (d localServiceRegistry) GetInstances(serviceId string) ([]cloud.ServiceInstance, error) {
	if ret, ok := d.ServiceInstances[serviceId]; ok {
		var result []cloud.ServiceInstance
		for _, value := range ret {
			result = append(result, value)
		}
		return result, nil
	} else {
		return nil, errors.New("no data")
	}
}

func (d localServiceRegistry) GetServices() ([]string, error) {
	var result []string
	for key, _ := range d.ServiceInstances {
		result = append(result, key)
	}
	return result, nil
}

func (d localServiceRegistry) Register(serviceInstance cloud.DefaultServiceInstance) bool {
	if d.ServiceInstances == nil {
		d.ServiceInstances = map[string]map[string]cloud.ServiceInstance{}
	}

	services := d.ServiceInstances[serviceInstance.GetServiceId()]

	if services == nil {
		services = map[string]cloud.ServiceInstance{}
	}

	services[serviceInstance.InstanceId] = serviceInstance

	d.ServiceInstances[serviceInstance.GetServiceId()] = services

	d.localServiceInstance = serviceInstance

	return true
}

func (d localServiceRegistry) Deregister() bool {
	if d.ServiceInstances == nil {
		return true
	}

	if d.localServiceInstance == nil {
		return true
	}

	services := d.ServiceInstances[d.localServiceInstance.GetServiceId()]

	if services == nil {
		return true
	}

	delete(services, d.localServiceInstance.GetInstanceId())

	if len(services) == 0 {
		delete(d.ServiceInstances, d.localServiceInstance.GetServiceId())
	}

	d.localServiceInstance = nil

	return true
}
