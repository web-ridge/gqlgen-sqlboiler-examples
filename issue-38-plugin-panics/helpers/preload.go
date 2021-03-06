package helpers

import (
	"context"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var TablePreloadMap = map[string]map[string]boilergql.ColumnSetting{
	models.TableNames.Additive: {
		"additiveInventories": {
			Name:			models.AdditiveRels.AdditiveInventories,
			RelationshipModelName:	models.TableNames.AdditiveInventory,
			IDAvailable:		false,
		},
		"recipeAdditive": {
			Name:			models.AdditiveRels.RecipeAdditive,
			RelationshipModelName:	models.TableNames.RecipeAdditive,
			IDAvailable:		true,
		},
		"recipeBatchAdditive": {
			Name:			models.AdditiveRels.RecipeBatchAdditive,
			RelationshipModelName:	models.TableNames.RecipeBatchAdditive,
			IDAvailable:		true,
		},
	},
	models.TableNames.AdditiveInventory: {
		"additive": {
			Name:			models.AdditiveInventoryRels.Additive,
			RelationshipModelName:	models.TableNames.Additive,
			IDAvailable:		true,
		},
		"supplier": {
			Name:			models.AdditiveInventoryRels.Supplier,
			RelationshipModelName:	models.TableNames.Supplier,
			IDAvailable:		true,
		},
	},
	models.TableNames.AuthGroup: {
		"groupAuthGroupPermissions": {
			Name:			models.AuthGroupRels.GroupAuthGroupPermissions,
			RelationshipModelName:	models.TableNames.AuthGroupPermissions,
			IDAvailable:		false,
		},
		"groupAuthUserGroups": {
			Name:			models.AuthGroupRels.GroupAuthUserGroups,
			RelationshipModelName:	models.TableNames.AuthUserGroups,
			IDAvailable:		false,
		},
	},
	models.TableNames.AuthGroupPermissions: {
		"group": {
			Name:			models.AuthGroupPermissionRels.Group,
			RelationshipModelName:	models.TableNames.AuthGroup,
			IDAvailable:		true,
		},
		"permission": {
			Name:			models.AuthGroupPermissionRels.Permission,
			RelationshipModelName:	models.TableNames.AuthPermission,
			IDAvailable:		true,
		},
	},
	models.TableNames.AuthPermission: {
		"permissionAuthGroupPermissions": {
			Name:			models.AuthPermissionRels.PermissionAuthGroupPermissions,
			RelationshipModelName:	models.TableNames.AuthGroupPermissions,
			IDAvailable:		false,
		},
		"permissionAuthUserUserPermissions": {
			Name:			models.AuthPermissionRels.PermissionAuthUserUserPermissions,
			RelationshipModelName:	models.TableNames.AuthUserUserPermissions,
			IDAvailable:		false,
		},
	},
	models.TableNames.AuthUser: {
		"userAuthUserGroups": {
			Name:			models.AuthUserRels.UserAuthUserGroups,
			RelationshipModelName:	models.TableNames.AuthUserGroups,
			IDAvailable:		false,
		},
		"userAuthUserUserPermissions": {
			Name:			models.AuthUserRels.UserAuthUserUserPermissions,
			RelationshipModelName:	models.TableNames.AuthUserUserPermissions,
			IDAvailable:		false,
		},
	},
	models.TableNames.AuthUserGroups: {
		"group": {
			Name:			models.AuthUserGroupRels.Group,
			RelationshipModelName:	models.TableNames.AuthGroup,
			IDAvailable:		true,
		},
		"user": {
			Name:			models.AuthUserGroupRels.User,
			RelationshipModelName:	models.TableNames.AuthUser,
			IDAvailable:		true,
		},
	},
	models.TableNames.AuthUserUserPermissions: {
		"permission": {
			Name:			models.AuthUserUserPermissionRels.Permission,
			RelationshipModelName:	models.TableNames.AuthPermission,
			IDAvailable:		true,
		},
		"user": {
			Name:			models.AuthUserUserPermissionRels.User,
			RelationshipModelName:	models.TableNames.AuthUser,
			IDAvailable:		true,
		},
	},
	models.TableNames.Fragrance: {
		"fragranceInventories": {
			Name:			models.FragranceRels.FragranceInventories,
			RelationshipModelName:	models.TableNames.FragranceInventory,
			IDAvailable:		false,
		},
		"recipeBatchFragrance": {
			Name:			models.FragranceRels.RecipeBatchFragrance,
			RelationshipModelName:	models.TableNames.RecipeBatchFragrance,
			IDAvailable:		true,
		},
		"recipeFragrance": {
			Name:			models.FragranceRels.RecipeFragrance,
			RelationshipModelName:	models.TableNames.RecipeFragrance,
			IDAvailable:		true,
		},
	},
	models.TableNames.FragranceInventory: {
		"fragrance": {
			Name:			models.FragranceInventoryRels.Fragrance,
			RelationshipModelName:	models.TableNames.Fragrance,
			IDAvailable:		true,
		},
		"supplier": {
			Name:			models.FragranceInventoryRels.Supplier,
			RelationshipModelName:	models.TableNames.Supplier,
			IDAvailable:		true,
		},
	},
	models.TableNames.Lipid: {
		"lipidInventories": {
			Name:			models.LipidRels.LipidInventories,
			RelationshipModelName:	models.TableNames.LipidInventory,
			IDAvailable:		false,
		},
		"recipeBatchLipid": {
			Name:			models.LipidRels.RecipeBatchLipid,
			RelationshipModelName:	models.TableNames.RecipeBatchLipid,
			IDAvailable:		true,
		},
		"recipeLipid": {
			Name:			models.LipidRels.RecipeLipid,
			RelationshipModelName:	models.TableNames.RecipeLipid,
			IDAvailable:		true,
		},
	},
	models.TableNames.LipidInventory: {
		"lipid": {
			Name:			models.LipidInventoryRels.Lipid,
			RelationshipModelName:	models.TableNames.Lipid,
			IDAvailable:		true,
		},
		"supplier": {
			Name:			models.LipidInventoryRels.Supplier,
			RelationshipModelName:	models.TableNames.Supplier,
			IDAvailable:		true,
		},
	},
	models.TableNames.Lye: {
		"lyeInventories": {
			Name:			models.LyeRels.LyeInventories,
			RelationshipModelName:	models.TableNames.LyeInventory,
			IDAvailable:		false,
		},
		"recipeBatchLye": {
			Name:			models.LyeRels.RecipeBatchLye,
			RelationshipModelName:	models.TableNames.RecipeBatchLye,
			IDAvailable:		true,
		},
	},
	models.TableNames.LyeInventory: {
		"lye": {
			Name:			models.LyeInventoryRels.Lye,
			RelationshipModelName:	models.TableNames.Lye,
			IDAvailable:		true,
		},
		"supplier": {
			Name:			models.LyeInventoryRels.Supplier,
			RelationshipModelName:	models.TableNames.Supplier,
			IDAvailable:		true,
		},
	},
	models.TableNames.Recipe: {
		"recipeAdditives": {
			Name:			models.RecipeRels.RecipeAdditives,
			RelationshipModelName:	models.TableNames.RecipeAdditive,
			IDAvailable:		false,
		},
		"recipeBatches": {
			Name:			models.RecipeRels.RecipeBatches,
			RelationshipModelName:	models.TableNames.RecipeBatch,
			IDAvailable:		false,
		},
		"recipeFragrances": {
			Name:			models.RecipeRels.RecipeFragrances,
			RelationshipModelName:	models.TableNames.RecipeFragrance,
			IDAvailable:		false,
		},
		"recipeLipids": {
			Name:			models.RecipeRels.RecipeLipids,
			RelationshipModelName:	models.TableNames.RecipeLipid,
			IDAvailable:		false,
		},
		"recipeSteps": {
			Name:			models.RecipeRels.RecipeSteps,
			RelationshipModelName:	models.TableNames.RecipeStep,
			IDAvailable:		false,
		},
	},
	models.TableNames.RecipeAdditive: {
		"additive": {
			Name:			models.RecipeAdditiveRels.Additive,
			RelationshipModelName:	models.TableNames.Additive,
			IDAvailable:		true,
		},
		"recipe": {
			Name:			models.RecipeAdditiveRels.Recipe,
			RelationshipModelName:	models.TableNames.Recipe,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeBatch: {
		"batchRecipeBatchAdditives": {
			Name:			models.RecipeBatchRels.BatchRecipeBatchAdditives,
			RelationshipModelName:	models.TableNames.RecipeBatchAdditive,
			IDAvailable:		false,
		},
		"batchRecipeBatchFragrances": {
			Name:			models.RecipeBatchRels.BatchRecipeBatchFragrances,
			RelationshipModelName:	models.TableNames.RecipeBatchFragrance,
			IDAvailable:		false,
		},
		"batchRecipeBatchLipids": {
			Name:			models.RecipeBatchRels.BatchRecipeBatchLipids,
			RelationshipModelName:	models.TableNames.RecipeBatchLipid,
			IDAvailable:		false,
		},
		"batchRecipeBatchLyes": {
			Name:			models.RecipeBatchRels.BatchRecipeBatchLyes,
			RelationshipModelName:	models.TableNames.RecipeBatchLye,
			IDAvailable:		false,
		},
		"batchRecipeBatchNotes": {
			Name:			models.RecipeBatchRels.BatchRecipeBatchNotes,
			RelationshipModelName:	models.TableNames.RecipeBatchNote,
			IDAvailable:		false,
		},
		"recipe": {
			Name:			models.RecipeBatchRels.Recipe,
			RelationshipModelName:	models.TableNames.Recipe,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeBatchAdditive: {
		"additive": {
			Name:			models.RecipeBatchAdditiveRels.Additive,
			RelationshipModelName:	models.TableNames.Additive,
			IDAvailable:		true,
		},
		"batch": {
			Name:			models.RecipeBatchAdditiveRels.Batch,
			RelationshipModelName:	models.TableNames.RecipeBatch,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeBatchFragrance: {
		"batch": {
			Name:			models.RecipeBatchFragranceRels.Batch,
			RelationshipModelName:	models.TableNames.RecipeBatch,
			IDAvailable:		true,
		},
		"fragrance": {
			Name:			models.RecipeBatchFragranceRels.Fragrance,
			RelationshipModelName:	models.TableNames.Fragrance,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeBatchLipid: {
		"batch": {
			Name:			models.RecipeBatchLipidRels.Batch,
			RelationshipModelName:	models.TableNames.RecipeBatch,
			IDAvailable:		true,
		},
		"lipid": {
			Name:			models.RecipeBatchLipidRels.Lipid,
			RelationshipModelName:	models.TableNames.Lipid,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeBatchLye: {
		"batch": {
			Name:			models.RecipeBatchLyeRels.Batch,
			RelationshipModelName:	models.TableNames.RecipeBatch,
			IDAvailable:		true,
		},
		"lye": {
			Name:			models.RecipeBatchLyeRels.Lye,
			RelationshipModelName:	models.TableNames.Lye,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeBatchNote: {
		"batch": {
			Name:			models.RecipeBatchNoteRels.Batch,
			RelationshipModelName:	models.TableNames.RecipeBatch,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeFragrance: {
		"fragrance": {
			Name:			models.RecipeFragranceRels.Fragrance,
			RelationshipModelName:	models.TableNames.Fragrance,
			IDAvailable:		true,
		},
		"recipe": {
			Name:			models.RecipeFragranceRels.Recipe,
			RelationshipModelName:	models.TableNames.Recipe,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeLipid: {
		"lipid": {
			Name:			models.RecipeLipidRels.Lipid,
			RelationshipModelName:	models.TableNames.Lipid,
			IDAvailable:		true,
		},
		"recipe": {
			Name:			models.RecipeLipidRels.Recipe,
			RelationshipModelName:	models.TableNames.Recipe,
			IDAvailable:		true,
		},
	},
	models.TableNames.RecipeStep: {
		"recipe": {
			Name:			models.RecipeStepRels.Recipe,
			RelationshipModelName:	models.TableNames.Recipe,
			IDAvailable:		true,
		},
	},
	models.TableNames.Supplier: {
		"additiveInventories": {
			Name:			models.SupplierRels.AdditiveInventories,
			RelationshipModelName:	models.TableNames.AdditiveInventory,
			IDAvailable:		false,
		},
		"fragranceInventories": {
			Name:			models.SupplierRels.FragranceInventories,
			RelationshipModelName:	models.TableNames.FragranceInventory,
			IDAvailable:		false,
		},
		"lipidInventories": {
			Name:			models.SupplierRels.LipidInventories,
			RelationshipModelName:	models.TableNames.LipidInventory,
			IDAvailable:		false,
		},
		"lyeInventories": {
			Name:			models.SupplierRels.LyeInventories,
			RelationshipModelName:	models.TableNames.LyeInventory,
			IDAvailable:		false,
		},
	},
}

func GetAdditivePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Additive, "")
}
func GetAdditiveNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Additive, DefaultLevels.EdgesNode)
}
func GetAdditivePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Additive, level)
}

func GetAdditiveInventoryPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AdditiveInventory, "")
}
func GetAdditiveInventoryNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AdditiveInventory, DefaultLevels.EdgesNode)
}
func GetAdditiveInventoryPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AdditiveInventory, level)
}

func GetAuthGroupPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthGroup, "")
}
func GetAuthGroupNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthGroup, DefaultLevels.EdgesNode)
}
func GetAuthGroupPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthGroup, level)
}

func GetAuthGroupPermissionPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthGroupPermissions, "")
}
func GetAuthGroupPermissionNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthGroupPermissions, DefaultLevels.EdgesNode)
}
func GetAuthGroupPermissionPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthGroupPermissions, level)
}

func GetAuthPermissionPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthPermission, "")
}
func GetAuthPermissionNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthPermission, DefaultLevels.EdgesNode)
}
func GetAuthPermissionPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthPermission, level)
}

func GetAuthUserPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUser, "")
}
func GetAuthUserNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUser, DefaultLevels.EdgesNode)
}
func GetAuthUserPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUser, level)
}

func GetAuthUserGroupPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUserGroups, "")
}
func GetAuthUserGroupNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUserGroups, DefaultLevels.EdgesNode)
}
func GetAuthUserGroupPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUserGroups, level)
}

func GetAuthUserUserPermissionPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUserUserPermissions, "")
}
func GetAuthUserUserPermissionNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUserUserPermissions, DefaultLevels.EdgesNode)
}
func GetAuthUserUserPermissionPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.AuthUserUserPermissions, level)
}

func GetFragrancePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Fragrance, "")
}
func GetFragranceNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Fragrance, DefaultLevels.EdgesNode)
}
func GetFragrancePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Fragrance, level)
}

func GetFragranceInventoryPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.FragranceInventory, "")
}
func GetFragranceInventoryNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.FragranceInventory, DefaultLevels.EdgesNode)
}
func GetFragranceInventoryPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.FragranceInventory, level)
}

func GetLipidPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Lipid, "")
}
func GetLipidNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Lipid, DefaultLevels.EdgesNode)
}
func GetLipidPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Lipid, level)
}

func GetLipidInventoryPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.LipidInventory, "")
}
func GetLipidInventoryNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.LipidInventory, DefaultLevels.EdgesNode)
}
func GetLipidInventoryPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.LipidInventory, level)
}

func GetLyePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Lye, "")
}
func GetLyeNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Lye, DefaultLevels.EdgesNode)
}
func GetLyePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Lye, level)
}

func GetLyeInventoryPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.LyeInventory, "")
}
func GetLyeInventoryNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.LyeInventory, DefaultLevels.EdgesNode)
}
func GetLyeInventoryPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.LyeInventory, level)
}

func GetRecipePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Recipe, "")
}
func GetRecipeNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Recipe, DefaultLevels.EdgesNode)
}
func GetRecipePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Recipe, level)
}

func GetRecipeAdditivePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeAdditive, "")
}
func GetRecipeAdditiveNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeAdditive, DefaultLevels.EdgesNode)
}
func GetRecipeAdditivePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeAdditive, level)
}

func GetRecipeBatchPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatch, "")
}
func GetRecipeBatchNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatch, DefaultLevels.EdgesNode)
}
func GetRecipeBatchPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatch, level)
}

func GetRecipeBatchAdditivePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchAdditive, "")
}
func GetRecipeBatchAdditiveNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchAdditive, DefaultLevels.EdgesNode)
}
func GetRecipeBatchAdditivePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchAdditive, level)
}

func GetRecipeBatchFragrancePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchFragrance, "")
}
func GetRecipeBatchFragranceNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchFragrance, DefaultLevels.EdgesNode)
}
func GetRecipeBatchFragrancePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchFragrance, level)
}

func GetRecipeBatchLipidPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchLipid, "")
}
func GetRecipeBatchLipidNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchLipid, DefaultLevels.EdgesNode)
}
func GetRecipeBatchLipidPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchLipid, level)
}

func GetRecipeBatchLyePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchLye, "")
}
func GetRecipeBatchLyeNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchLye, DefaultLevels.EdgesNode)
}
func GetRecipeBatchLyePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchLye, level)
}

func GetRecipeBatchNotePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchNote, "")
}
func GetRecipeBatchNoteNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchNote, DefaultLevels.EdgesNode)
}
func GetRecipeBatchNotePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeBatchNote, level)
}

func GetRecipeFragrancePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeFragrance, "")
}
func GetRecipeFragranceNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeFragrance, DefaultLevels.EdgesNode)
}
func GetRecipeFragrancePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeFragrance, level)
}

func GetRecipeLipidPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeLipid, "")
}
func GetRecipeLipidNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeLipid, DefaultLevels.EdgesNode)
}
func GetRecipeLipidPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeLipid, level)
}

func GetRecipeStepPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeStep, "")
}
func GetRecipeStepNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeStep, DefaultLevels.EdgesNode)
}
func GetRecipeStepPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.RecipeStep, level)
}

func GetSupplierPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Supplier, "")
}
func GetSupplierNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Supplier, DefaultLevels.EdgesNode)
}
func GetSupplierPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Supplier, level)
}

var AdditiveInventoryPayloadPreloadLevels = struct {
	AdditiveInventory string
}{
	AdditiveInventory: "additiveInventory",
}

var AdditivePayloadPreloadLevels = struct {
	Additive string
}{
	Additive: "additive",
}

var AuthGroupPayloadPreloadLevels = struct {
	AuthGroup string
}{
	AuthGroup: "authGroup",
}

var AuthGroupPermissionPayloadPreloadLevels = struct {
	AuthGroupPermission string
}{
	AuthGroupPermission: "authGroupPermission",
}

var AuthPermissionPayloadPreloadLevels = struct {
	AuthPermission string
}{
	AuthPermission: "authPermission",
}

var AuthUserGroupPayloadPreloadLevels = struct {
	AuthUserGroup string
}{
	AuthUserGroup: "authUserGroup",
}

var AuthUserPayloadPreloadLevels = struct {
	AuthUser string
}{
	AuthUser: "authUser",
}

var AuthUserUserPermissionPayloadPreloadLevels = struct {
	AuthUserUserPermission string
}{
	AuthUserUserPermission: "authUserUserPermission",
}

var FragranceInventoryPayloadPreloadLevels = struct {
	FragranceInventory string
}{
	FragranceInventory: "fragranceInventory",
}

var FragrancePayloadPreloadLevels = struct {
	Fragrance string
}{
	Fragrance: "fragrance",
}

var LipidInventoryPayloadPreloadLevels = struct {
	LipidInventory string
}{
	LipidInventory: "lipidInventory",
}

var LipidPayloadPreloadLevels = struct {
	Lipid string
}{
	Lipid: "lipid",
}

var LyeInventoryPayloadPreloadLevels = struct {
	LyeInventory string
}{
	LyeInventory: "lyeInventory",
}

var LyePayloadPreloadLevels = struct {
	Lye string
}{
	Lye: "lye",
}

var RecipeAdditivePayloadPreloadLevels = struct {
	RecipeAdditive string
}{
	RecipeAdditive: "recipeAdditive",
}

var RecipeBatchAdditivePayloadPreloadLevels = struct {
	RecipeBatchAdditive string
}{
	RecipeBatchAdditive: "recipeBatchAdditive",
}

var RecipeBatchFragrancePayloadPreloadLevels = struct {
	RecipeBatchFragrance string
}{
	RecipeBatchFragrance: "recipeBatchFragrance",
}

var RecipeBatchLipidPayloadPreloadLevels = struct {
	RecipeBatchLipid string
}{
	RecipeBatchLipid: "recipeBatchLipid",
}

var RecipeBatchLyePayloadPreloadLevels = struct {
	RecipeBatchLye string
}{
	RecipeBatchLye: "recipeBatchLye",
}

var RecipeBatchNotePayloadPreloadLevels = struct {
	RecipeBatchNote string
}{
	RecipeBatchNote: "recipeBatchNote",
}

var RecipeBatchPayloadPreloadLevels = struct {
	RecipeBatch string
}{
	RecipeBatch: "recipeBatch",
}

var RecipeFragrancePayloadPreloadLevels = struct {
	RecipeFragrance string
}{
	RecipeFragrance: "recipeFragrance",
}

var RecipeLipidPayloadPreloadLevels = struct {
	RecipeLipid string
}{
	RecipeLipid: "recipeLipid",
}

var RecipePayloadPreloadLevels = struct {
	Recipe string
}{
	Recipe: "recipe",
}

var RecipeStepPayloadPreloadLevels = struct {
	RecipeStep string
}{
	RecipeStep: "recipeStep",
}

var SupplierPayloadPreloadLevels = struct {
	Supplier string
}{
	Supplier: "supplier",
}

var DefaultLevels = struct {
	EdgesNode string
}{
	EdgesNode: "edges.node",
}
