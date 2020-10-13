package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lemon/cloud"
	"lemon/cloud/serviceregistry"
	"lemon/util"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestConsulServiceRegistry(t *testing.T) {
	host := "127.0.0.1"
	port := 8500
	token := ""
	registryDiscoveryClient, err := serviceregistry.NewConsulServiceRegistry(host, port, token)

	ip, err := util.GetLocalIP()
	if err != nil {
		t.Error(err)
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
	server := &http.Server{Addr: ":8010", Handler: r}
	server.RegisterOnShutdown(func() {
		fmt.Println(" shutdown ...")
		registryDiscoveryClient.Deregister()
	})
	err = server.ListenAndServe()
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

	t.Log(registryDiscoveryClient.GetInstances("go-user-server"))
}
