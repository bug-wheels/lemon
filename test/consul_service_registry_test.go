package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon/cloud"
	"lemon/cloud/serviceregistry"
	"lemon/cloud/util"
	"math/rand"
	"testing"
	"time"
)

func TestConsulServiceRegistry(t *testing.T) {
	host := "127.0.0.1"
	port := 8500
	token := ""
	registryDiscoveryClient, err := serviceregistry.NewConsulServiceRegistry(host, port, token)

	ip, err := util.FindFirstNonLoopbackIP()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	fmt.Println(ip)
	rand.Seed(time.Now().UnixNano())

	si, _ := cloud.NewDefaultServiceInstance("go-user-server", "", 8010,
		false, map[string]string{"user": "zyn2"}, "")

	registryDiscoveryClient.Register(si)

	r := gin.Default()
	r.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err = r.Run(":8010")
	if err != nil {
		registryDiscoveryClient.Deregister()
	}
}

func TestConsulServiceDiscovery(t *testing.T) {
	host := "127.0.0.1"
	port := 8500
	token := ""
	registryDiscoveryClient, err := serviceregistry.NewConsulServiceRegistry(host, port, token)
	if err != nil {
		panic(err)
	}

	t.Log(registryDiscoveryClient.GetServices())

	t.Log(registryDiscoveryClient.GetInstances("ecm-monitor"))
}
