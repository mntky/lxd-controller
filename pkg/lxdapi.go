package lxdpkg

import (
	"fmt"

	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func Connect() (container lxd.ContainerServer) {
	container, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		fmt.Println(err)
	}
	return container
}

func Status(name string, container lxd.ContainerServer) *api.ContainerState {
	var stat *api.ContainerState

	stat, str, err := container.GetContainerState(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return stat
}
