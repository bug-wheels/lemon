package test

import (
	"github.com/hashicorp/consul/api"
	"strconv"
	"testing"
)

func TestConsul(t *testing.T) {
	host := "127.0.0.1"
	port := 8500
	token := ""

	config := api.DefaultConfig()
	config.Address = host + ":" + strconv.Itoa(port)
	config.Token = token
	client, _ := api.NewClient(config)
	services, _, _ := client.Catalog().Services(nil)

	for serviceName, _ := range services {
		catalogService, _, _ := client.Catalog().Service(serviceName, "", nil)
		if len(catalogService) > 0 {
			checks, _, _ := client.Health().Checks(serviceName, nil)
			for _, check := range checks {
				if check.Status != "passing" {
					client.Agent().ServiceDeregister(check.ServiceID)
				}
			}
		}
	}
}
