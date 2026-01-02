package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
	}

	authConfig := types.AuthConfig{
		Username: "your_username",
		Password: "your_password",
	}

	encodedAuth, err := json.Marshal(authConfig)
	if err != nil {
		fmt.Println(err)
	}
	authStr := base64.StdEncoding.EncodeToString(encodedAuth)
	out, err := cli.ImagePull(ctx, "swr.ru-moscow-1.hc.vtbcloud.ru/tkp3/alpine:latest", image.PullOptions{RegistryAuth: authStr})
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	fmt.Println("Image pulled successfully")
}
