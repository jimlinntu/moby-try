package main

import (
    "context"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "io/ioutil"
    "fmt"
)

func main(){
    cli, err := client.NewClientWithOpts(client.FromEnv)
    fmt.Printf("%v\n", cli.ClientVersion())
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    response, err := cli.ImagePull(ctx, "hello-world", types.ImagePullOptions{})
    if err != nil {
        fmt.Printf("There is an error in ImagePull! %s\n", err)
        panic(err)
    }
    body, err := ioutil.ReadAll(response)
    fmt.Printf("%s\n", body)
    return
}
