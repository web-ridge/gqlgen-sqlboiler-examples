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

	if err := gbgen.SchemaWrite(gbgen.SchemaConfig{
		BoilerModelDirectory: backend,
		GenerateBatchCreate:  true,
		GenerateMutations:    true,
		GenerateBatchDelete:  true,
		GenerateBatchUpdate:  true,
		Directives:           []string{"isAuthenticated"},
	}, "schema.graphql", gbgen.SchemaGenerateConfig{
		MergeSchema: false,
	}); err != nil {
		fmt.Println("error while trying to generate schema.graphql")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
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
			gbgen.ResolverPluginConfig{
				// Authorization scopes can be used to override e.g. userId, organizationId, tenantId
				// This will be resolved used the provided ScopeResolverName if the result of the AddTrigger=true
				// You would need this if you don't want to require these fields in your schema but you want to add them
				// to the db model.
				// If you do have these fields in your schema but want them authorized you could use a gqlgen directive
				AuthorizationScopes: []*gbgen.AuthorizationScope{
					{
						ImportPath:        "github.com/web-ridge/gqlgen-sqlboiler-examples/social-network/auth",
						ImportAlias:       "auth",
						ScopeResolverName: "UserIDFromContext", // function which is called with the context of the resolver
						BoilerColumnName:  "UserID",

						AddHook: func(model *gbgen.BoilerModel, resolver *gbgen.Resolver, templateKey string) bool {
							// fmt.Println(templateKey)
							// templateKey contains a unique where the resolver tries to add something
							// e.g.
							// most of the time you can ignore this

							for _, field := range model.Fields {
								if field.Name == "UserID" {
									return true
								}
							}
							return false
						},
					},
				},
			},
		)),
	)
	if err != nil {
		fmt.Println("error!!")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
