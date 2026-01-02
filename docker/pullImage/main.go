package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
	}
	out, err := cli.ImagePull(ctx, "alpine:latest", image.PullOptions{})
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	fmt.Println("Image pulled successfully")
}
