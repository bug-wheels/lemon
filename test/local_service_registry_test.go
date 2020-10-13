package test

import (
	"fmt"
	"lemon/cloud"
	"lemon/cloud/serviceregistry"
	"testing"
)

func TestLocalServiceRegistry(t *testing.T) {
	t.Log("success")

	localRegistry := serviceregistry.NewLocalServiceRegistry()

	si := cloud.DefaultServiceInstance{ServiceId: "user",
		InstanceId: "dddddasd",
		Host:       "192.168.0.23",
		Port:       9090,
		Secure:     true}

	fmt.Println(localRegistry.GetServices())

	localRegistry.Register(si)
	fmt.Println(localRegistry)
	fmt.Println(localRegistry.GetServices())
	fmt.Println(localRegistry.GetInstances("user"))

	localRegistry.Deregister()
	fmt.Println(localRegistry)
	fmt.Println(localRegistry.GetServices())
	fmt.Println(localRegistry.GetInstances("user"))
}
