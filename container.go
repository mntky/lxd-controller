package main

import (
	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

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

func (c *Container)connect() (error){
	var err error
	c.server, err = lxd.ConnectLXDUnix("", nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Container)create() (error){
	req := api.ContainersPost{
		Name: c.info.Name,
		Source: api.ContainerSource{
			Type: "image",
			Alias: c.info.Image,
		},
	}

	op, err := c.server.CreateContainer(req)
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *Container)start() (error) {
	reqState := api.ContainerStatePut {
		Action: "start",
		Timeout: -1,
	}

	op, err := c.server.UpdateContainerState(c.info.Name, reqState, "")
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *Container)stop() (error){
	reqState := api.ContainerStatePut {
		Action: "stop",
		Timeout: -1,
	}

	op, err := c.server.UpdateContainerState(c.info.Name, reqState, "")
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *Container)delete() (error){
	op, err := c.server.DeleteContainer(c.info.Name)
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *Container)status() (string, error){
	var (
		str string
		err error
	)
	c.stat, str, err = c.server.GetContainerState(c.info.Name)
	if err != nil {
		return str, err
	}
	return str, nil
}
