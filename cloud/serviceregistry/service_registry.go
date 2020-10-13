package serviceregistry

import "lemon/cloud"

type ServiceRegistry interface {
	Register(serviceInstance cloud.ServiceInstance) bool

	Deregister()
}
