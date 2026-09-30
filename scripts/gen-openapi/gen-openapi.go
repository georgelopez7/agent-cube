package main

import (
	"fmt"

	xhttp "agent-cube/api/http"
	rubikscube "agent-cube/internal/service/rubikscube"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/spf13/cobra"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8080"`
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print the OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			svc := rubikscube.NewRubiksCubeService(nil, nil, nil, nil)

			server := xhttp.NewServer("Agent Cube API", "v1", "8080", svc, nil)
			server.AddRoutes(server.API)

			b, err := server.API.OpenAPI().YAML()
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))
		},
	})

	cli.Run()
}
