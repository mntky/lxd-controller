package main

import (
	"fmt"
	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func connect() (container lxd.ContainerServer) {

	//connect to lxd over the unix socket
	//c, err := lxd.ConnectLXDUnix("", nil)
	//fmt.Println(c, err)
	//if err != nil {
	//	return err
	//}

	container, err := lxd.ConnectLXDUnix("", nil)
	fmt.Println(err)
	return 

}

func create(container lxd.ContainerServer) {

	//container creation request
	req := api.ContainersPost{
		Name: "container-name",
		Source: api.ContainerSource{
			Type: "image",
			Alias: "debian",
		},
	}
	fmt.Println(req)

	op, err := container.CreateContainer(req)
	fmt.Println(err)

	err = op.Wait()

	fmt.Println(err)
}



func main() {
	fmt.Println("---LXD-API---")
	c  := connect()
	create(c)
}

