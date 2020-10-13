package test

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	t.Log(time.Now().Unix())
}

func TestIp(t *testing.T) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < len(netInterfaces); i++ {
		i2 := netInterfaces[i]
		if (i2.Flags & net.FlagUp) != 0 {
			adders, _ := i2.Addrs()

			for _, address := range adders {
				if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
					if inet.IP.To4() != nil {
						fmt.Println(inet, i2)
					}
				}
			}
		}
	}

	t.Log("success")
}
