package loadbalance

import (
	"errors"
	"lemon/cloud"
)

type LoadBalance interface {
	choose(serviceId string, instances []cloud.ServiceInstance) (cloud.ServiceInstance, error)
}

type FirstLoadBalance struct {
}

func (f FirstLoadBalance) choose(serviceId string, instances []cloud.ServiceInstance) (cloud.ServiceInstance, error) {
	if instances != nil && len(instances) > 0 {
		return instances[0], nil
	}
	return nil, errors.New("no available instance")
}
