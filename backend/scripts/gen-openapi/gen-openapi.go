package main

import (
	"fmt"

	apihttp "agent-cube/api/http"
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/service"

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
			svc := service.NewService(nil, nil, event.NewEventBus())

			server := apihttp.NewServer("Agent Cube API", "v1", "8080", svc, nil)
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
