package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

/*
    todo: Replace all panic with return statements
	and replace Print statements with some logging.
**/
func ListContainer() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		for _, container := range containers {
			formatSpacedStrings(container.Image, container.ID)
		}
	} else {
		fmt.Println("There are no containers running")
	}
	return nil
}
