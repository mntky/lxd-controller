package main

import (
	"log"
	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func failOnError(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


type ContainerInfo struct {
	Name	string	`json:"name"`
	Image	string	`json:"image"`
	VxlanId	int		`json:"vxlanid"`
}

type Container struct {
	server lxd.ContainerServer
	info ContainerInfo
	stat *api.ContainerState
}

func (c *Container)connect() {
	var err error
	c.server, err = lxd.ConnectLXDUnix("", nil)
	failOnError(err, "fail connect LXD Unix")
	return
}

func (c *Container)create() {
	req := api.ContainersPost{
		Name: c.info.Name,
		Source: api.ContainerSource{
			Type: "image",
			Alias: c.info.Image,
		},
	}

	op, err := c.server.CreateContainer(req)
	failOnError(err, "fail create container")

	err = op.Wait()
	failOnError(err, "fail create container operation")
}

func (c *Container)start() {
	reqState := api.ContainerStatePut {
		Action: "start",
		Timeout: -1,
	}

	op, err := c.server.UpdateContainerState(c.info.Name, reqState, "")
	failOnError(err, "fail update container state")

	err = op.Wait()
	failOnError(err, "fail start container operation")
}

func (c *Container)stop() {
	reqState := api.ContainerStatePut {
		Action: "stop",
		Timeout: -1,
	}

	op, err := c.server.UpdateContainerState(c.info.Name, reqState, "")
	failOnError(err, "Fail update container state")

	err = op.Wait()
	failOnError(err, "Fail stop container operation")
}

func (c *Container)delete() {
	op, err := c.server.DeleteContainer(c.info.Name)
	failOnError(err, "Fail delete container")

	err = op.Wait()
	failOnError(err, "Fail delete container operation")
}

func (c *Container)status() {
	c.stat, str, err := c.server.GetContainerState(c.info.Name)
	failOnError(err, str)
}
