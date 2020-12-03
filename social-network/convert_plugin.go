// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"

	gbgen "github.com/web-ridge/gqlgen-sqlboiler/v3"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	output := gbgen.Config{
		Directory:   "helpers", // supports root or sub directories
		PackageName: "helpers",
	}
	backend := gbgen.Config{
		Directory:   "models",
		PackageName: "models",
	}
	frontend := gbgen.Config{
		Directory:   "graphql_models",
		PackageName: "graphql_models",
	}

	err = gbgen.SchemaWrite(gbgen.SchemaConfig{
		BoilerModelDirectory: backend,
		GenerateBatchCreate:  true,
		GenerateMutations:    true,
		GenerateBatchDelete:  true,
		GenerateBatchUpdate:  true,
		Directives:           []string{"isAuthenticated"},
	}, "schema.graphql", gbgen.SchemaGenerateConfig{
		MergeSchema: false,
	})

	if err != nil {
		fmt.Println("error while trying to gbgen.SchemaWrite")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
	err = api.Generate(cfg,
		api.AddPlugin(gbgen.NewConvertPlugin(
			output,   // directory where convert.go, convert_input.go and preload.go should live
			backend,  // directory where sqlboiler files are put
			frontend, // directory where gqlgen models live
			gbgen.ConvertPluginConfig{},
		)),
		api.AddPlugin(gbgen.NewResolverPlugin(
			output,
			backend,
			frontend,
			"github.com/web-ridge/gqlgen-sqlboiler-examples/social-network/auth", // leave empty if you don't have auth
		)),
	)
	if err != nil {
		fmt.Println("error!!")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
