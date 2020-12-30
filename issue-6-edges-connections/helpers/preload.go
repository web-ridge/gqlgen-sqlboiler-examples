package helpers

import (
	"context"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var TablePreloadMap = map[string]map[string]boilergql.ColumnSetting{
	models.TableNames.User: {},
}

func GetUserPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.User, "")
}
func GetUserNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.User, DefaultLevels.EdgesNode)
}
func GetUserPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.User, level)
}

var DefaultLevels = struct {
	EdgesNode string
}{
	EdgesNode: "edges.node",
}
