package main

import (
	"fmt"
	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

//Connect to Unixsocket 
func connect() (container lxd.ContainerServer) {

	container, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//Create container 
func create(container lxd.ContainerServer) {

	//container creation request
	req := api.ContainersPost{
		Name: "container-name",
		Source: api.ContainerSource{
			Type: "image",
			Alias: "debian",
		},
	}

	op, err := container.CreateContainer(req)
	if err != nil {
		fmt.Println(err)
	}

	err = op.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

//Start Container 
func start(name string, container lxd.ContainerServer) {
	reqState := api.ContainerStatePut{
		Action: "start",
		Timeout: -1,
	}

	op, err := container.UpdateContainerState(name, reqState, "")
	if err != nil {
		fmt.Println(err)
	}

	err = op.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

//Stop Container
func stop(name string, container lxd.ContainerServer) {
	reqState := api.ContainerStatePut{
		Action: "stop",
		Timeout: -1,
	}

	op, err := container.UpdateContainerState(name, reqState, "")
	if err != nil {
		fmt.Println(err)
	}

	err = op.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

//Delete Container
//if doesnt stop the container , it will be panic
func delete(name string, container lxd.ContainerServer) {
	op, err := container.DeleteContainer(name)
	if err != nil {
		fmt.Println(err)
	}

	err = op.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

//Get Container Status
func status(name string, container lxd.ContainerServer) {
	var stat *api.ContainerState

	stat, str, err := container.GetContainerState(name)
	if err != nil {
		fmt.Println(err)
		fmt.Println(str)
	}
	fmt.Println(*stat)
}

func main() {
	c  := connect()
	//create(c)
	//start("container-name", c)
	//stop("container-name", c)
	status("debian", c)
	//delete("container-name", c)
}
