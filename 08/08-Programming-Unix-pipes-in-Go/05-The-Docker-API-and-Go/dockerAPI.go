package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// type Container struct {
// 	ID         string `json:"Id"`
// 	Names      []string
// 	Image      string
// 	ImageID    string
// 	Command    string
// 	Created    int64
// 	Ports      []Port
// 	SizeRw     int64 `json:",omitempty"`
// 	SizeRootFs int64 `json:",omitempty"`
// 	Labels     map[string]string
// 	State      string
// 	Status     string
// 	HostConfig struct {
// 		NetworkMode string `json:",omitempty"`
// 	}
// 	NetworkSettings *SummaryNetworkSettings
// 	Mounts          []MountPoint
// }

func listContainers() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err
	}

	for _, container := range containers {
		fmt.Println("Images:", container.Image, "with ID:", container.ID)
	}

	return nil
}

// type ImageSummary struct {

// 	// containers
// 	// Required: true
// 	Containers int64 `json:"Containers"`

// 	// created
// 	// Required: true
// 	Created int64 `json:"Created"`

// 	// Id
// 	// Required: true
// 	ID string `json:"Id"`

// 	// labels
// 	// Required: true
// 	Labels map[string]string `json:"Labels"`

// 	// parent Id
// 	// Required: true
// 	ParentID string `json:"ParentId"`

// 	// repo digests
// 	// Required: true
// 	RepoDigests []string `json:"RepoDigests"`

// 	// repo tags
// 	// Required: true
// 	RepoTags []string `json:"RepoTags"`

// 	// shared size
// 	// Required: true
// 	SharedSize int64 `json:"SharedSize"`

// 	// size
// 	// Required: true
// 	Size int64 `json:"Size"`

// 	// virtual size
// 	// Required: true
// 	VirtualSize int64 `json:"VirtualSize"`
// }

func listImages() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	imgs, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return err
	}

	for _, img := range imgs {
		fmt.Printf("Images %s with size %d\n", img.RepoTags, img.Size)
	}

	return nil
}

func main() {

	fmt.Println("The available images are:")
	if err := listImages(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("The running Containers are:")
	if err := listContainers(); err != nil {
		fmt.Println(err)
	}

}
