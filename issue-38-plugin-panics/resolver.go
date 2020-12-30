package main

import (
	"context"
	"errors"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/web-ridge/utils-go/boilergql/v3"

	"database/sql"

	"github.com/rs/zerolog/log"

	. "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/helpers"

	dm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/models"

	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/graphql_models"
)

type Resolver struct {
	db *sql.DB
}

const inputKey = "input"

const publicAdditiveCreateError = "could not create additive"

func (r *mutationResolver) CreateAdditive(ctx context.Context, input fm.AdditiveCreateInput) (*fm.AdditivePayload, error) {

	m := AdditiveCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAdditiveCreateError)
		return nil, errors.New(publicAdditiveCreateError)
	}

	mods := GetAdditivePreloadModsWithLevel(ctx, AdditivePayloadPreloadLevels.Additive)
	mods = append(mods, dm.AdditiveWhere.ID.EQ(m.ID))
	pM, err := dm.Additives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveCreateError)
		return nil, errors.New(publicAdditiveCreateError)
	}
	return &fm.AdditivePayload{
		Additive: AdditiveToGraphQL(pM),
	}, nil
}

const publicAdditiveUpdateError = "could not update additive"

func (r *mutationResolver) UpdateAdditive(ctx context.Context, id string, input fm.AdditiveUpdateInput) (*fm.AdditivePayload, error) {
	m := AdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AdditiveID(id)
	if _, err := dm.Additives(
		dm.AdditiveWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAdditiveUpdateError)
		return nil, errors.New(publicAdditiveUpdateError)
	}

	mods := GetAdditivePreloadModsWithLevel(ctx, AdditivePayloadPreloadLevels.Additive)
	mods = append(mods, dm.AdditiveWhere.ID.EQ(dbID))

	pM, err := dm.Additives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveUpdateError)
		return nil, errors.New(publicAdditiveUpdateError)
	}
	return &fm.AdditivePayload{
		Additive: AdditiveToGraphQL(pM),
	}, nil
}

const publicAdditiveBatchUpdateError = "could not update additives"

func (r *mutationResolver) UpdateAdditives(ctx context.Context, filter *fm.AdditiveFilter, input fm.AdditiveUpdateInput) (*fm.AdditivesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AdditiveFilterToMods(filter)...)

	m := AdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Additives(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAdditiveBatchUpdateError)
		return nil, errors.New(publicAdditiveBatchUpdateError)
	}

	return &fm.AdditivesUpdatePayload{
		Ok: true,
	}, nil
}

const publicAdditiveDeleteError = "could not delete additive"

func (r *mutationResolver) DeleteAdditive(ctx context.Context, id string) (*fm.AdditiveDeletePayload, error) {
	dbID := AdditiveID(id)
	mods := []qm.QueryMod{
		dm.AdditiveWhere.ID.EQ(dbID),
	}
	if _, err := dm.Additives(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAdditiveDeleteError)
		return nil, errors.New(publicAdditiveDeleteError)
	}

	return &fm.AdditiveDeletePayload{
		ID: id,
	}, nil
}

const publicAdditiveBatchDeleteError = "could not delete additives"

func (r *mutationResolver) DeleteAdditives(ctx context.Context, filter *fm.AdditiveFilter) (*fm.AdditivesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AdditiveFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AdditiveColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Additive))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.Additives(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAdditiveBatchDeleteError)
		return nil, errors.New(publicAdditiveBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.Additives(dm.AdditiveWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAdditiveBatchDeleteError)
		return nil, errors.New(publicAdditiveBatchDeleteError)
	}

	return &fm.AdditivesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.Additive),
	}, nil
}

const publicAdditiveInventoryCreateError = "could not create additiveInventory"

func (r *mutationResolver) CreateAdditiveInventory(ctx context.Context, input fm.AdditiveInventoryCreateInput) (*fm.AdditiveInventoryPayload, error) {

	m := AdditiveInventoryCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryCreateError)
		return nil, errors.New(publicAdditiveInventoryCreateError)
	}

	mods := GetAdditiveInventoryPreloadModsWithLevel(ctx, AdditiveInventoryPayloadPreloadLevels.AdditiveInventory)
	mods = append(mods, dm.AdditiveInventoryWhere.ID.EQ(m.ID))
	pM, err := dm.AdditiveInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryCreateError)
		return nil, errors.New(publicAdditiveInventoryCreateError)
	}
	return &fm.AdditiveInventoryPayload{
		AdditiveInventory: AdditiveInventoryToGraphQL(pM),
	}, nil
}

const publicAdditiveInventoryUpdateError = "could not update additiveInventory"

func (r *mutationResolver) UpdateAdditiveInventory(ctx context.Context, id string, input fm.AdditiveInventoryUpdateInput) (*fm.AdditiveInventoryPayload, error) {
	m := AdditiveInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AdditiveInventoryID(id)
	if _, err := dm.AdditiveInventories(
		dm.AdditiveInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryUpdateError)
		return nil, errors.New(publicAdditiveInventoryUpdateError)
	}

	mods := GetAdditiveInventoryPreloadModsWithLevel(ctx, AdditiveInventoryPayloadPreloadLevels.AdditiveInventory)
	mods = append(mods, dm.AdditiveInventoryWhere.ID.EQ(dbID))

	pM, err := dm.AdditiveInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryUpdateError)
		return nil, errors.New(publicAdditiveInventoryUpdateError)
	}
	return &fm.AdditiveInventoryPayload{
		AdditiveInventory: AdditiveInventoryToGraphQL(pM),
	}, nil
}

const publicAdditiveInventoryBatchUpdateError = "could not update additiveInventories"

func (r *mutationResolver) UpdateAdditiveInventories(ctx context.Context, filter *fm.AdditiveInventoryFilter, input fm.AdditiveInventoryUpdateInput) (*fm.AdditiveInventoriesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AdditiveInventoryFilterToMods(filter)...)

	m := AdditiveInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.AdditiveInventories(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryBatchUpdateError)
		return nil, errors.New(publicAdditiveInventoryBatchUpdateError)
	}

	return &fm.AdditiveInventoriesUpdatePayload{
		Ok: true,
	}, nil
}

const publicAdditiveInventoryDeleteError = "could not delete additiveInventory"

func (r *mutationResolver) DeleteAdditiveInventory(ctx context.Context, id string) (*fm.AdditiveInventoryDeletePayload, error) {
	dbID := AdditiveInventoryID(id)
	mods := []qm.QueryMod{
		dm.AdditiveInventoryWhere.ID.EQ(dbID),
	}
	if _, err := dm.AdditiveInventories(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryDeleteError)
		return nil, errors.New(publicAdditiveInventoryDeleteError)
	}

	return &fm.AdditiveInventoryDeletePayload{
		ID: id,
	}, nil
}

const publicAdditiveInventoryBatchDeleteError = "could not delete additiveInventories"

func (r *mutationResolver) DeleteAdditiveInventories(ctx context.Context, filter *fm.AdditiveInventoryFilter) (*fm.AdditiveInventoriesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AdditiveInventoryFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AdditiveInventoryColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AdditiveInventory))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.AdditiveInventories(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryBatchDeleteError)
		return nil, errors.New(publicAdditiveInventoryBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.AdditiveInventories(dm.AdditiveInventoryWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryBatchDeleteError)
		return nil, errors.New(publicAdditiveInventoryBatchDeleteError)
	}

	return &fm.AdditiveInventoriesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AdditiveInventory),
	}, nil
}

const publicAuthGroupCreateError = "could not create authGroup"

func (r *mutationResolver) CreateAuthGroup(ctx context.Context, input fm.AuthGroupCreateInput) (*fm.AuthGroupPayload, error) {

	m := AuthGroupCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupCreateError)
		return nil, errors.New(publicAuthGroupCreateError)
	}

	mods := GetAuthGroupPreloadModsWithLevel(ctx, AuthGroupPayloadPreloadLevels.AuthGroup)
	mods = append(mods, dm.AuthGroupWhere.ID.EQ(m.ID))
	pM, err := dm.AuthGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupCreateError)
		return nil, errors.New(publicAuthGroupCreateError)
	}
	return &fm.AuthGroupPayload{
		AuthGroup: AuthGroupToGraphQL(pM),
	}, nil
}

const publicAuthGroupUpdateError = "could not update authGroup"

func (r *mutationResolver) UpdateAuthGroup(ctx context.Context, id string, input fm.AuthGroupUpdateInput) (*fm.AuthGroupPayload, error) {
	m := AuthGroupUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthGroupID(id)
	if _, err := dm.AuthGroups(
		dm.AuthGroupWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupUpdateError)
		return nil, errors.New(publicAuthGroupUpdateError)
	}

	mods := GetAuthGroupPreloadModsWithLevel(ctx, AuthGroupPayloadPreloadLevels.AuthGroup)
	mods = append(mods, dm.AuthGroupWhere.ID.EQ(dbID))

	pM, err := dm.AuthGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupUpdateError)
		return nil, errors.New(publicAuthGroupUpdateError)
	}
	return &fm.AuthGroupPayload{
		AuthGroup: AuthGroupToGraphQL(pM),
	}, nil
}

const publicAuthGroupBatchUpdateError = "could not update authGroups"

func (r *mutationResolver) UpdateAuthGroups(ctx context.Context, filter *fm.AuthGroupFilter, input fm.AuthGroupUpdateInput) (*fm.AuthGroupsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthGroupFilterToMods(filter)...)

	m := AuthGroupUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.AuthGroups(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupBatchUpdateError)
		return nil, errors.New(publicAuthGroupBatchUpdateError)
	}

	return &fm.AuthGroupsUpdatePayload{
		Ok: true,
	}, nil
}

const publicAuthGroupDeleteError = "could not delete authGroup"

func (r *mutationResolver) DeleteAuthGroup(ctx context.Context, id string) (*fm.AuthGroupDeletePayload, error) {
	dbID := AuthGroupID(id)
	mods := []qm.QueryMod{
		dm.AuthGroupWhere.ID.EQ(dbID),
	}
	if _, err := dm.AuthGroups(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupDeleteError)
		return nil, errors.New(publicAuthGroupDeleteError)
	}

	return &fm.AuthGroupDeletePayload{
		ID: id,
	}, nil
}

const publicAuthGroupBatchDeleteError = "could not delete authGroups"

func (r *mutationResolver) DeleteAuthGroups(ctx context.Context, filter *fm.AuthGroupFilter) (*fm.AuthGroupsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthGroupFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthGroupColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthGroup))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.AuthGroups(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupBatchDeleteError)
		return nil, errors.New(publicAuthGroupBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.AuthGroups(dm.AuthGroupWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupBatchDeleteError)
		return nil, errors.New(publicAuthGroupBatchDeleteError)
	}

	return &fm.AuthGroupsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthGroup),
	}, nil
}

const publicAuthGroupPermissionCreateError = "could not create authGroupPermission"

func (r *mutationResolver) CreateAuthGroupPermission(ctx context.Context, input fm.AuthGroupPermissionCreateInput) (*fm.AuthGroupPermissionPayload, error) {

	m := AuthGroupPermissionCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionCreateError)
		return nil, errors.New(publicAuthGroupPermissionCreateError)
	}

	mods := GetAuthGroupPermissionPreloadModsWithLevel(ctx, AuthGroupPermissionPayloadPreloadLevels.AuthGroupPermission)
	mods = append(mods, dm.AuthGroupPermissionWhere.ID.EQ(m.ID))
	pM, err := dm.AuthGroupPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionCreateError)
		return nil, errors.New(publicAuthGroupPermissionCreateError)
	}
	return &fm.AuthGroupPermissionPayload{
		AuthGroupPermission: AuthGroupPermissionToGraphQL(pM),
	}, nil
}

const publicAuthGroupPermissionUpdateError = "could not update authGroupPermission"

func (r *mutationResolver) UpdateAuthGroupPermission(ctx context.Context, id string, input fm.AuthGroupPermissionUpdateInput) (*fm.AuthGroupPermissionPayload, error) {
	m := AuthGroupPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthGroupPermissionID(id)
	if _, err := dm.AuthGroupPermissions(
		dm.AuthGroupPermissionWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionUpdateError)
		return nil, errors.New(publicAuthGroupPermissionUpdateError)
	}

	mods := GetAuthGroupPermissionPreloadModsWithLevel(ctx, AuthGroupPermissionPayloadPreloadLevels.AuthGroupPermission)
	mods = append(mods, dm.AuthGroupPermissionWhere.ID.EQ(dbID))

	pM, err := dm.AuthGroupPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionUpdateError)
		return nil, errors.New(publicAuthGroupPermissionUpdateError)
	}
	return &fm.AuthGroupPermissionPayload{
		AuthGroupPermission: AuthGroupPermissionToGraphQL(pM),
	}, nil
}

const publicAuthGroupPermissionBatchUpdateError = "could not update authGroupPermissions"

func (r *mutationResolver) UpdateAuthGroupPermissions(ctx context.Context, filter *fm.AuthGroupPermissionFilter, input fm.AuthGroupPermissionUpdateInput) (*fm.AuthGroupPermissionsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthGroupPermissionFilterToMods(filter)...)

	m := AuthGroupPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.AuthGroupPermissions(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionBatchUpdateError)
		return nil, errors.New(publicAuthGroupPermissionBatchUpdateError)
	}

	return &fm.AuthGroupPermissionsUpdatePayload{
		Ok: true,
	}, nil
}

const publicAuthGroupPermissionDeleteError = "could not delete authGroupPermission"

func (r *mutationResolver) DeleteAuthGroupPermission(ctx context.Context, id string) (*fm.AuthGroupPermissionDeletePayload, error) {
	dbID := AuthGroupPermissionID(id)
	mods := []qm.QueryMod{
		dm.AuthGroupPermissionWhere.ID.EQ(dbID),
	}
	if _, err := dm.AuthGroupPermissions(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionDeleteError)
		return nil, errors.New(publicAuthGroupPermissionDeleteError)
	}

	return &fm.AuthGroupPermissionDeletePayload{
		ID: id,
	}, nil
}

const publicAuthGroupPermissionBatchDeleteError = "could not delete authGroupPermissions"

func (r *mutationResolver) DeleteAuthGroupPermissions(ctx context.Context, filter *fm.AuthGroupPermissionFilter) (*fm.AuthGroupPermissionsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthGroupPermissionFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthGroupPermissionColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthGroupPermissions))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.AuthGroupPermissions(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionBatchDeleteError)
		return nil, errors.New(publicAuthGroupPermissionBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.AuthGroupPermissions(dm.AuthGroupPermissionWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionBatchDeleteError)
		return nil, errors.New(publicAuthGroupPermissionBatchDeleteError)
	}

	return &fm.AuthGroupPermissionsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthGroupPermissions),
	}, nil
}

const publicAuthPermissionCreateError = "could not create authPermission"

func (r *mutationResolver) CreateAuthPermission(ctx context.Context, input fm.AuthPermissionCreateInput) (*fm.AuthPermissionPayload, error) {

	m := AuthPermissionCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionCreateError)
		return nil, errors.New(publicAuthPermissionCreateError)
	}

	mods := GetAuthPermissionPreloadModsWithLevel(ctx, AuthPermissionPayloadPreloadLevels.AuthPermission)
	mods = append(mods, dm.AuthPermissionWhere.ID.EQ(m.ID))
	pM, err := dm.AuthPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionCreateError)
		return nil, errors.New(publicAuthPermissionCreateError)
	}
	return &fm.AuthPermissionPayload{
		AuthPermission: AuthPermissionToGraphQL(pM),
	}, nil
}

const publicAuthPermissionUpdateError = "could not update authPermission"

func (r *mutationResolver) UpdateAuthPermission(ctx context.Context, id string, input fm.AuthPermissionUpdateInput) (*fm.AuthPermissionPayload, error) {
	m := AuthPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthPermissionID(id)
	if _, err := dm.AuthPermissions(
		dm.AuthPermissionWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionUpdateError)
		return nil, errors.New(publicAuthPermissionUpdateError)
	}

	mods := GetAuthPermissionPreloadModsWithLevel(ctx, AuthPermissionPayloadPreloadLevels.AuthPermission)
	mods = append(mods, dm.AuthPermissionWhere.ID.EQ(dbID))

	pM, err := dm.AuthPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionUpdateError)
		return nil, errors.New(publicAuthPermissionUpdateError)
	}
	return &fm.AuthPermissionPayload{
		AuthPermission: AuthPermissionToGraphQL(pM),
	}, nil
}

const publicAuthPermissionBatchUpdateError = "could not update authPermissions"

func (r *mutationResolver) UpdateAuthPermissions(ctx context.Context, filter *fm.AuthPermissionFilter, input fm.AuthPermissionUpdateInput) (*fm.AuthPermissionsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthPermissionFilterToMods(filter)...)

	m := AuthPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.AuthPermissions(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionBatchUpdateError)
		return nil, errors.New(publicAuthPermissionBatchUpdateError)
	}

	return &fm.AuthPermissionsUpdatePayload{
		Ok: true,
	}, nil
}

const publicAuthPermissionDeleteError = "could not delete authPermission"

func (r *mutationResolver) DeleteAuthPermission(ctx context.Context, id string) (*fm.AuthPermissionDeletePayload, error) {
	dbID := AuthPermissionID(id)
	mods := []qm.QueryMod{
		dm.AuthPermissionWhere.ID.EQ(dbID),
	}
	if _, err := dm.AuthPermissions(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionDeleteError)
		return nil, errors.New(publicAuthPermissionDeleteError)
	}

	return &fm.AuthPermissionDeletePayload{
		ID: id,
	}, nil
}

const publicAuthPermissionBatchDeleteError = "could not delete authPermissions"

func (r *mutationResolver) DeleteAuthPermissions(ctx context.Context, filter *fm.AuthPermissionFilter) (*fm.AuthPermissionsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthPermissionFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthPermissionColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthPermission))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.AuthPermissions(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionBatchDeleteError)
		return nil, errors.New(publicAuthPermissionBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.AuthPermissions(dm.AuthPermissionWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionBatchDeleteError)
		return nil, errors.New(publicAuthPermissionBatchDeleteError)
	}

	return &fm.AuthPermissionsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthPermission),
	}, nil
}

const publicAuthUserCreateError = "could not create authUser"

func (r *mutationResolver) CreateAuthUser(ctx context.Context, input fm.AuthUserCreateInput) (*fm.AuthUserPayload, error) {

	m := AuthUserCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAuthUserCreateError)
		return nil, errors.New(publicAuthUserCreateError)
	}

	mods := GetAuthUserPreloadModsWithLevel(ctx, AuthUserPayloadPreloadLevels.AuthUser)
	mods = append(mods, dm.AuthUserWhere.ID.EQ(m.ID))
	pM, err := dm.AuthUsers(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserCreateError)
		return nil, errors.New(publicAuthUserCreateError)
	}
	return &fm.AuthUserPayload{
		AuthUser: AuthUserToGraphQL(pM),
	}, nil
}

const publicAuthUserUpdateError = "could not update authUser"

func (r *mutationResolver) UpdateAuthUser(ctx context.Context, id string, input fm.AuthUserUpdateInput) (*fm.AuthUserPayload, error) {
	m := AuthUserUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthUserID(id)
	if _, err := dm.AuthUsers(
		dm.AuthUserWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUpdateError)
		return nil, errors.New(publicAuthUserUpdateError)
	}

	mods := GetAuthUserPreloadModsWithLevel(ctx, AuthUserPayloadPreloadLevels.AuthUser)
	mods = append(mods, dm.AuthUserWhere.ID.EQ(dbID))

	pM, err := dm.AuthUsers(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUpdateError)
		return nil, errors.New(publicAuthUserUpdateError)
	}
	return &fm.AuthUserPayload{
		AuthUser: AuthUserToGraphQL(pM),
	}, nil
}

const publicAuthUserBatchUpdateError = "could not update authUsers"

func (r *mutationResolver) UpdateAuthUsers(ctx context.Context, filter *fm.AuthUserFilter, input fm.AuthUserUpdateInput) (*fm.AuthUsersUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthUserFilterToMods(filter)...)

	m := AuthUserUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.AuthUsers(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserBatchUpdateError)
		return nil, errors.New(publicAuthUserBatchUpdateError)
	}

	return &fm.AuthUsersUpdatePayload{
		Ok: true,
	}, nil
}

const publicAuthUserDeleteError = "could not delete authUser"

func (r *mutationResolver) DeleteAuthUser(ctx context.Context, id string) (*fm.AuthUserDeletePayload, error) {
	dbID := AuthUserID(id)
	mods := []qm.QueryMod{
		dm.AuthUserWhere.ID.EQ(dbID),
	}
	if _, err := dm.AuthUsers(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserDeleteError)
		return nil, errors.New(publicAuthUserDeleteError)
	}

	return &fm.AuthUserDeletePayload{
		ID: id,
	}, nil
}

const publicAuthUserBatchDeleteError = "could not delete authUsers"

func (r *mutationResolver) DeleteAuthUsers(ctx context.Context, filter *fm.AuthUserFilter) (*fm.AuthUsersDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthUserFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthUserColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthUser))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.AuthUsers(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAuthUserBatchDeleteError)
		return nil, errors.New(publicAuthUserBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.AuthUsers(dm.AuthUserWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserBatchDeleteError)
		return nil, errors.New(publicAuthUserBatchDeleteError)
	}

	return &fm.AuthUsersDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthUser),
	}, nil
}

const publicAuthUserGroupCreateError = "could not create authUserGroup"

func (r *mutationResolver) CreateAuthUserGroup(ctx context.Context, input fm.AuthUserGroupCreateInput) (*fm.AuthUserGroupPayload, error) {

	m := AuthUserGroupCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupCreateError)
		return nil, errors.New(publicAuthUserGroupCreateError)
	}

	mods := GetAuthUserGroupPreloadModsWithLevel(ctx, AuthUserGroupPayloadPreloadLevels.AuthUserGroup)
	mods = append(mods, dm.AuthUserGroupWhere.ID.EQ(m.ID))
	pM, err := dm.AuthUserGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupCreateError)
		return nil, errors.New(publicAuthUserGroupCreateError)
	}
	return &fm.AuthUserGroupPayload{
		AuthUserGroup: AuthUserGroupToGraphQL(pM),
	}, nil
}

const publicAuthUserGroupUpdateError = "could not update authUserGroup"

func (r *mutationResolver) UpdateAuthUserGroup(ctx context.Context, id string, input fm.AuthUserGroupUpdateInput) (*fm.AuthUserGroupPayload, error) {
	m := AuthUserGroupUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthUserGroupID(id)
	if _, err := dm.AuthUserGroups(
		dm.AuthUserGroupWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupUpdateError)
		return nil, errors.New(publicAuthUserGroupUpdateError)
	}

	mods := GetAuthUserGroupPreloadModsWithLevel(ctx, AuthUserGroupPayloadPreloadLevels.AuthUserGroup)
	mods = append(mods, dm.AuthUserGroupWhere.ID.EQ(dbID))

	pM, err := dm.AuthUserGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupUpdateError)
		return nil, errors.New(publicAuthUserGroupUpdateError)
	}
	return &fm.AuthUserGroupPayload{
		AuthUserGroup: AuthUserGroupToGraphQL(pM),
	}, nil
}

const publicAuthUserGroupBatchUpdateError = "could not update authUserGroups"

func (r *mutationResolver) UpdateAuthUserGroups(ctx context.Context, filter *fm.AuthUserGroupFilter, input fm.AuthUserGroupUpdateInput) (*fm.AuthUserGroupsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthUserGroupFilterToMods(filter)...)

	m := AuthUserGroupUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.AuthUserGroups(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupBatchUpdateError)
		return nil, errors.New(publicAuthUserGroupBatchUpdateError)
	}

	return &fm.AuthUserGroupsUpdatePayload{
		Ok: true,
	}, nil
}

const publicAuthUserGroupDeleteError = "could not delete authUserGroup"

func (r *mutationResolver) DeleteAuthUserGroup(ctx context.Context, id string) (*fm.AuthUserGroupDeletePayload, error) {
	dbID := AuthUserGroupID(id)
	mods := []qm.QueryMod{
		dm.AuthUserGroupWhere.ID.EQ(dbID),
	}
	if _, err := dm.AuthUserGroups(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupDeleteError)
		return nil, errors.New(publicAuthUserGroupDeleteError)
	}

	return &fm.AuthUserGroupDeletePayload{
		ID: id,
	}, nil
}

const publicAuthUserGroupBatchDeleteError = "could not delete authUserGroups"

func (r *mutationResolver) DeleteAuthUserGroups(ctx context.Context, filter *fm.AuthUserGroupFilter) (*fm.AuthUserGroupsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthUserGroupFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthUserGroupColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthUserGroups))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.AuthUserGroups(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupBatchDeleteError)
		return nil, errors.New(publicAuthUserGroupBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.AuthUserGroups(dm.AuthUserGroupWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupBatchDeleteError)
		return nil, errors.New(publicAuthUserGroupBatchDeleteError)
	}

	return &fm.AuthUserGroupsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthUserGroups),
	}, nil
}

const publicAuthUserUserPermissionCreateError = "could not create authUserUserPermission"

func (r *mutationResolver) CreateAuthUserUserPermission(ctx context.Context, input fm.AuthUserUserPermissionCreateInput) (*fm.AuthUserUserPermissionPayload, error) {

	m := AuthUserUserPermissionCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionCreateError)
		return nil, errors.New(publicAuthUserUserPermissionCreateError)
	}

	mods := GetAuthUserUserPermissionPreloadModsWithLevel(ctx, AuthUserUserPermissionPayloadPreloadLevels.AuthUserUserPermission)
	mods = append(mods, dm.AuthUserUserPermissionWhere.ID.EQ(m.ID))
	pM, err := dm.AuthUserUserPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionCreateError)
		return nil, errors.New(publicAuthUserUserPermissionCreateError)
	}
	return &fm.AuthUserUserPermissionPayload{
		AuthUserUserPermission: AuthUserUserPermissionToGraphQL(pM),
	}, nil
}

const publicAuthUserUserPermissionUpdateError = "could not update authUserUserPermission"

func (r *mutationResolver) UpdateAuthUserUserPermission(ctx context.Context, id string, input fm.AuthUserUserPermissionUpdateInput) (*fm.AuthUserUserPermissionPayload, error) {
	m := AuthUserUserPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthUserUserPermissionID(id)
	if _, err := dm.AuthUserUserPermissions(
		dm.AuthUserUserPermissionWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionUpdateError)
		return nil, errors.New(publicAuthUserUserPermissionUpdateError)
	}

	mods := GetAuthUserUserPermissionPreloadModsWithLevel(ctx, AuthUserUserPermissionPayloadPreloadLevels.AuthUserUserPermission)
	mods = append(mods, dm.AuthUserUserPermissionWhere.ID.EQ(dbID))

	pM, err := dm.AuthUserUserPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionUpdateError)
		return nil, errors.New(publicAuthUserUserPermissionUpdateError)
	}
	return &fm.AuthUserUserPermissionPayload{
		AuthUserUserPermission: AuthUserUserPermissionToGraphQL(pM),
	}, nil
}

const publicAuthUserUserPermissionBatchUpdateError = "could not update authUserUserPermissions"

func (r *mutationResolver) UpdateAuthUserUserPermissions(ctx context.Context, filter *fm.AuthUserUserPermissionFilter, input fm.AuthUserUserPermissionUpdateInput) (*fm.AuthUserUserPermissionsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthUserUserPermissionFilterToMods(filter)...)

	m := AuthUserUserPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.AuthUserUserPermissions(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionBatchUpdateError)
		return nil, errors.New(publicAuthUserUserPermissionBatchUpdateError)
	}

	return &fm.AuthUserUserPermissionsUpdatePayload{
		Ok: true,
	}, nil
}

const publicAuthUserUserPermissionDeleteError = "could not delete authUserUserPermission"

func (r *mutationResolver) DeleteAuthUserUserPermission(ctx context.Context, id string) (*fm.AuthUserUserPermissionDeletePayload, error) {
	dbID := AuthUserUserPermissionID(id)
	mods := []qm.QueryMod{
		dm.AuthUserUserPermissionWhere.ID.EQ(dbID),
	}
	if _, err := dm.AuthUserUserPermissions(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionDeleteError)
		return nil, errors.New(publicAuthUserUserPermissionDeleteError)
	}

	return &fm.AuthUserUserPermissionDeletePayload{
		ID: id,
	}, nil
}

const publicAuthUserUserPermissionBatchDeleteError = "could not delete authUserUserPermissions"

func (r *mutationResolver) DeleteAuthUserUserPermissions(ctx context.Context, filter *fm.AuthUserUserPermissionFilter) (*fm.AuthUserUserPermissionsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, AuthUserUserPermissionFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthUserUserPermissionColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthUserUserPermissions))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.AuthUserUserPermissions(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionBatchDeleteError)
		return nil, errors.New(publicAuthUserUserPermissionBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.AuthUserUserPermissions(dm.AuthUserUserPermissionWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionBatchDeleteError)
		return nil, errors.New(publicAuthUserUserPermissionBatchDeleteError)
	}

	return &fm.AuthUserUserPermissionsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthUserUserPermissions),
	}, nil
}

const publicFragranceCreateError = "could not create fragrance"

func (r *mutationResolver) CreateFragrance(ctx context.Context, input fm.FragranceCreateInput) (*fm.FragrancePayload, error) {

	m := FragranceCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicFragranceCreateError)
		return nil, errors.New(publicFragranceCreateError)
	}

	mods := GetFragrancePreloadModsWithLevel(ctx, FragrancePayloadPreloadLevels.Fragrance)
	mods = append(mods, dm.FragranceWhere.ID.EQ(m.ID))
	pM, err := dm.Fragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceCreateError)
		return nil, errors.New(publicFragranceCreateError)
	}
	return &fm.FragrancePayload{
		Fragrance: FragranceToGraphQL(pM),
	}, nil
}

const publicFragranceUpdateError = "could not update fragrance"

func (r *mutationResolver) UpdateFragrance(ctx context.Context, id string, input fm.FragranceUpdateInput) (*fm.FragrancePayload, error) {
	m := FragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := FragranceID(id)
	if _, err := dm.Fragrances(
		dm.FragranceWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFragranceUpdateError)
		return nil, errors.New(publicFragranceUpdateError)
	}

	mods := GetFragrancePreloadModsWithLevel(ctx, FragrancePayloadPreloadLevels.Fragrance)
	mods = append(mods, dm.FragranceWhere.ID.EQ(dbID))

	pM, err := dm.Fragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceUpdateError)
		return nil, errors.New(publicFragranceUpdateError)
	}
	return &fm.FragrancePayload{
		Fragrance: FragranceToGraphQL(pM),
	}, nil
}

const publicFragranceBatchUpdateError = "could not update fragrances"

func (r *mutationResolver) UpdateFragrances(ctx context.Context, filter *fm.FragranceFilter, input fm.FragranceUpdateInput) (*fm.FragrancesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, FragranceFilterToMods(filter)...)

	m := FragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Fragrances(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFragranceBatchUpdateError)
		return nil, errors.New(publicFragranceBatchUpdateError)
	}

	return &fm.FragrancesUpdatePayload{
		Ok: true,
	}, nil
}

const publicFragranceDeleteError = "could not delete fragrance"

func (r *mutationResolver) DeleteFragrance(ctx context.Context, id string) (*fm.FragranceDeletePayload, error) {
	dbID := FragranceID(id)
	mods := []qm.QueryMod{
		dm.FragranceWhere.ID.EQ(dbID),
	}
	if _, err := dm.Fragrances(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicFragranceDeleteError)
		return nil, errors.New(publicFragranceDeleteError)
	}

	return &fm.FragranceDeletePayload{
		ID: id,
	}, nil
}

const publicFragranceBatchDeleteError = "could not delete fragrances"

func (r *mutationResolver) DeleteFragrances(ctx context.Context, filter *fm.FragranceFilter) (*fm.FragrancesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, FragranceFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.FragranceColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Fragrance))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.Fragrances(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicFragranceBatchDeleteError)
		return nil, errors.New(publicFragranceBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.Fragrances(dm.FragranceWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicFragranceBatchDeleteError)
		return nil, errors.New(publicFragranceBatchDeleteError)
	}

	return &fm.FragrancesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.Fragrance),
	}, nil
}

const publicFragranceInventoryCreateError = "could not create fragranceInventory"

func (r *mutationResolver) CreateFragranceInventory(ctx context.Context, input fm.FragranceInventoryCreateInput) (*fm.FragranceInventoryPayload, error) {

	m := FragranceInventoryCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryCreateError)
		return nil, errors.New(publicFragranceInventoryCreateError)
	}

	mods := GetFragranceInventoryPreloadModsWithLevel(ctx, FragranceInventoryPayloadPreloadLevels.FragranceInventory)
	mods = append(mods, dm.FragranceInventoryWhere.ID.EQ(m.ID))
	pM, err := dm.FragranceInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryCreateError)
		return nil, errors.New(publicFragranceInventoryCreateError)
	}
	return &fm.FragranceInventoryPayload{
		FragranceInventory: FragranceInventoryToGraphQL(pM),
	}, nil
}

const publicFragranceInventoryUpdateError = "could not update fragranceInventory"

func (r *mutationResolver) UpdateFragranceInventory(ctx context.Context, id string, input fm.FragranceInventoryUpdateInput) (*fm.FragranceInventoryPayload, error) {
	m := FragranceInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := FragranceInventoryID(id)
	if _, err := dm.FragranceInventories(
		dm.FragranceInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryUpdateError)
		return nil, errors.New(publicFragranceInventoryUpdateError)
	}

	mods := GetFragranceInventoryPreloadModsWithLevel(ctx, FragranceInventoryPayloadPreloadLevels.FragranceInventory)
	mods = append(mods, dm.FragranceInventoryWhere.ID.EQ(dbID))

	pM, err := dm.FragranceInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryUpdateError)
		return nil, errors.New(publicFragranceInventoryUpdateError)
	}
	return &fm.FragranceInventoryPayload{
		FragranceInventory: FragranceInventoryToGraphQL(pM),
	}, nil
}

const publicFragranceInventoryBatchUpdateError = "could not update fragranceInventories"

func (r *mutationResolver) UpdateFragranceInventories(ctx context.Context, filter *fm.FragranceInventoryFilter, input fm.FragranceInventoryUpdateInput) (*fm.FragranceInventoriesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, FragranceInventoryFilterToMods(filter)...)

	m := FragranceInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.FragranceInventories(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryBatchUpdateError)
		return nil, errors.New(publicFragranceInventoryBatchUpdateError)
	}

	return &fm.FragranceInventoriesUpdatePayload{
		Ok: true,
	}, nil
}

const publicFragranceInventoryDeleteError = "could not delete fragranceInventory"

func (r *mutationResolver) DeleteFragranceInventory(ctx context.Context, id string) (*fm.FragranceInventoryDeletePayload, error) {
	dbID := FragranceInventoryID(id)
	mods := []qm.QueryMod{
		dm.FragranceInventoryWhere.ID.EQ(dbID),
	}
	if _, err := dm.FragranceInventories(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryDeleteError)
		return nil, errors.New(publicFragranceInventoryDeleteError)
	}

	return &fm.FragranceInventoryDeletePayload{
		ID: id,
	}, nil
}

const publicFragranceInventoryBatchDeleteError = "could not delete fragranceInventories"

func (r *mutationResolver) DeleteFragranceInventories(ctx context.Context, filter *fm.FragranceInventoryFilter) (*fm.FragranceInventoriesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, FragranceInventoryFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.FragranceInventoryColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.FragranceInventory))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.FragranceInventories(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryBatchDeleteError)
		return nil, errors.New(publicFragranceInventoryBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.FragranceInventories(dm.FragranceInventoryWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryBatchDeleteError)
		return nil, errors.New(publicFragranceInventoryBatchDeleteError)
	}

	return &fm.FragranceInventoriesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.FragranceInventory),
	}, nil
}

const publicLipidCreateError = "could not create lipid"

func (r *mutationResolver) CreateLipid(ctx context.Context, input fm.LipidCreateInput) (*fm.LipidPayload, error) {

	m := LipidCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicLipidCreateError)
		return nil, errors.New(publicLipidCreateError)
	}

	mods := GetLipidPreloadModsWithLevel(ctx, LipidPayloadPreloadLevels.Lipid)
	mods = append(mods, dm.LipidWhere.ID.EQ(m.ID))
	pM, err := dm.Lipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidCreateError)
		return nil, errors.New(publicLipidCreateError)
	}
	return &fm.LipidPayload{
		Lipid: LipidToGraphQL(pM),
	}, nil
}

const publicLipidUpdateError = "could not update lipid"

func (r *mutationResolver) UpdateLipid(ctx context.Context, id string, input fm.LipidUpdateInput) (*fm.LipidPayload, error) {
	m := LipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LipidID(id)
	if _, err := dm.Lipids(
		dm.LipidWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLipidUpdateError)
		return nil, errors.New(publicLipidUpdateError)
	}

	mods := GetLipidPreloadModsWithLevel(ctx, LipidPayloadPreloadLevels.Lipid)
	mods = append(mods, dm.LipidWhere.ID.EQ(dbID))

	pM, err := dm.Lipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidUpdateError)
		return nil, errors.New(publicLipidUpdateError)
	}
	return &fm.LipidPayload{
		Lipid: LipidToGraphQL(pM),
	}, nil
}

const publicLipidBatchUpdateError = "could not update lipids"

func (r *mutationResolver) UpdateLipids(ctx context.Context, filter *fm.LipidFilter, input fm.LipidUpdateInput) (*fm.LipidsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LipidFilterToMods(filter)...)

	m := LipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Lipids(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLipidBatchUpdateError)
		return nil, errors.New(publicLipidBatchUpdateError)
	}

	return &fm.LipidsUpdatePayload{
		Ok: true,
	}, nil
}

const publicLipidDeleteError = "could not delete lipid"

func (r *mutationResolver) DeleteLipid(ctx context.Context, id string) (*fm.LipidDeletePayload, error) {
	dbID := LipidID(id)
	mods := []qm.QueryMod{
		dm.LipidWhere.ID.EQ(dbID),
	}
	if _, err := dm.Lipids(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLipidDeleteError)
		return nil, errors.New(publicLipidDeleteError)
	}

	return &fm.LipidDeletePayload{
		ID: id,
	}, nil
}

const publicLipidBatchDeleteError = "could not delete lipids"

func (r *mutationResolver) DeleteLipids(ctx context.Context, filter *fm.LipidFilter) (*fm.LipidsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LipidFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.LipidColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Lipid))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.Lipids(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicLipidBatchDeleteError)
		return nil, errors.New(publicLipidBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.Lipids(dm.LipidWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLipidBatchDeleteError)
		return nil, errors.New(publicLipidBatchDeleteError)
	}

	return &fm.LipidsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.Lipid),
	}, nil
}

const publicLipidInventoryCreateError = "could not create lipidInventory"

func (r *mutationResolver) CreateLipidInventory(ctx context.Context, input fm.LipidInventoryCreateInput) (*fm.LipidInventoryPayload, error) {

	m := LipidInventoryCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryCreateError)
		return nil, errors.New(publicLipidInventoryCreateError)
	}

	mods := GetLipidInventoryPreloadModsWithLevel(ctx, LipidInventoryPayloadPreloadLevels.LipidInventory)
	mods = append(mods, dm.LipidInventoryWhere.ID.EQ(m.ID))
	pM, err := dm.LipidInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryCreateError)
		return nil, errors.New(publicLipidInventoryCreateError)
	}
	return &fm.LipidInventoryPayload{
		LipidInventory: LipidInventoryToGraphQL(pM),
	}, nil
}

const publicLipidInventoryUpdateError = "could not update lipidInventory"

func (r *mutationResolver) UpdateLipidInventory(ctx context.Context, id string, input fm.LipidInventoryUpdateInput) (*fm.LipidInventoryPayload, error) {
	m := LipidInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LipidInventoryID(id)
	if _, err := dm.LipidInventories(
		dm.LipidInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryUpdateError)
		return nil, errors.New(publicLipidInventoryUpdateError)
	}

	mods := GetLipidInventoryPreloadModsWithLevel(ctx, LipidInventoryPayloadPreloadLevels.LipidInventory)
	mods = append(mods, dm.LipidInventoryWhere.ID.EQ(dbID))

	pM, err := dm.LipidInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryUpdateError)
		return nil, errors.New(publicLipidInventoryUpdateError)
	}
	return &fm.LipidInventoryPayload{
		LipidInventory: LipidInventoryToGraphQL(pM),
	}, nil
}

const publicLipidInventoryBatchUpdateError = "could not update lipidInventories"

func (r *mutationResolver) UpdateLipidInventories(ctx context.Context, filter *fm.LipidInventoryFilter, input fm.LipidInventoryUpdateInput) (*fm.LipidInventoriesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LipidInventoryFilterToMods(filter)...)

	m := LipidInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.LipidInventories(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryBatchUpdateError)
		return nil, errors.New(publicLipidInventoryBatchUpdateError)
	}

	return &fm.LipidInventoriesUpdatePayload{
		Ok: true,
	}, nil
}

const publicLipidInventoryDeleteError = "could not delete lipidInventory"

func (r *mutationResolver) DeleteLipidInventory(ctx context.Context, id string) (*fm.LipidInventoryDeletePayload, error) {
	dbID := LipidInventoryID(id)
	mods := []qm.QueryMod{
		dm.LipidInventoryWhere.ID.EQ(dbID),
	}
	if _, err := dm.LipidInventories(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryDeleteError)
		return nil, errors.New(publicLipidInventoryDeleteError)
	}

	return &fm.LipidInventoryDeletePayload{
		ID: id,
	}, nil
}

const publicLipidInventoryBatchDeleteError = "could not delete lipidInventories"

func (r *mutationResolver) DeleteLipidInventories(ctx context.Context, filter *fm.LipidInventoryFilter) (*fm.LipidInventoriesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LipidInventoryFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.LipidInventoryColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.LipidInventory))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.LipidInventories(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryBatchDeleteError)
		return nil, errors.New(publicLipidInventoryBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.LipidInventories(dm.LipidInventoryWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryBatchDeleteError)
		return nil, errors.New(publicLipidInventoryBatchDeleteError)
	}

	return &fm.LipidInventoriesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.LipidInventory),
	}, nil
}

const publicLyeCreateError = "could not create lye"

func (r *mutationResolver) CreateLye(ctx context.Context, input fm.LyeCreateInput) (*fm.LyePayload, error) {

	m := LyeCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicLyeCreateError)
		return nil, errors.New(publicLyeCreateError)
	}

	mods := GetLyePreloadModsWithLevel(ctx, LyePayloadPreloadLevels.Lye)
	mods = append(mods, dm.LyeWhere.ID.EQ(m.ID))
	pM, err := dm.Lyes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeCreateError)
		return nil, errors.New(publicLyeCreateError)
	}
	return &fm.LyePayload{
		Lye: LyeToGraphQL(pM),
	}, nil
}

const publicLyeUpdateError = "could not update lye"

func (r *mutationResolver) UpdateLye(ctx context.Context, id string, input fm.LyeUpdateInput) (*fm.LyePayload, error) {
	m := LyeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LyeID(id)
	if _, err := dm.Lyes(
		dm.LyeWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLyeUpdateError)
		return nil, errors.New(publicLyeUpdateError)
	}

	mods := GetLyePreloadModsWithLevel(ctx, LyePayloadPreloadLevels.Lye)
	mods = append(mods, dm.LyeWhere.ID.EQ(dbID))

	pM, err := dm.Lyes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeUpdateError)
		return nil, errors.New(publicLyeUpdateError)
	}
	return &fm.LyePayload{
		Lye: LyeToGraphQL(pM),
	}, nil
}

const publicLyeBatchUpdateError = "could not update lyes"

func (r *mutationResolver) UpdateLyes(ctx context.Context, filter *fm.LyeFilter, input fm.LyeUpdateInput) (*fm.LyesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LyeFilterToMods(filter)...)

	m := LyeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Lyes(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLyeBatchUpdateError)
		return nil, errors.New(publicLyeBatchUpdateError)
	}

	return &fm.LyesUpdatePayload{
		Ok: true,
	}, nil
}

const publicLyeDeleteError = "could not delete lye"

func (r *mutationResolver) DeleteLye(ctx context.Context, id string) (*fm.LyeDeletePayload, error) {
	dbID := LyeID(id)
	mods := []qm.QueryMod{
		dm.LyeWhere.ID.EQ(dbID),
	}
	if _, err := dm.Lyes(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLyeDeleteError)
		return nil, errors.New(publicLyeDeleteError)
	}

	return &fm.LyeDeletePayload{
		ID: id,
	}, nil
}

const publicLyeBatchDeleteError = "could not delete lyes"

func (r *mutationResolver) DeleteLyes(ctx context.Context, filter *fm.LyeFilter) (*fm.LyesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LyeFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.LyeColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Lye))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.Lyes(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicLyeBatchDeleteError)
		return nil, errors.New(publicLyeBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.Lyes(dm.LyeWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLyeBatchDeleteError)
		return nil, errors.New(publicLyeBatchDeleteError)
	}

	return &fm.LyesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.Lye),
	}, nil
}

const publicLyeInventoryCreateError = "could not create lyeInventory"

func (r *mutationResolver) CreateLyeInventory(ctx context.Context, input fm.LyeInventoryCreateInput) (*fm.LyeInventoryPayload, error) {

	m := LyeInventoryCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryCreateError)
		return nil, errors.New(publicLyeInventoryCreateError)
	}

	mods := GetLyeInventoryPreloadModsWithLevel(ctx, LyeInventoryPayloadPreloadLevels.LyeInventory)
	mods = append(mods, dm.LyeInventoryWhere.ID.EQ(m.ID))
	pM, err := dm.LyeInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryCreateError)
		return nil, errors.New(publicLyeInventoryCreateError)
	}
	return &fm.LyeInventoryPayload{
		LyeInventory: LyeInventoryToGraphQL(pM),
	}, nil
}

const publicLyeInventoryUpdateError = "could not update lyeInventory"

func (r *mutationResolver) UpdateLyeInventory(ctx context.Context, id string, input fm.LyeInventoryUpdateInput) (*fm.LyeInventoryPayload, error) {
	m := LyeInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LyeInventoryID(id)
	if _, err := dm.LyeInventories(
		dm.LyeInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryUpdateError)
		return nil, errors.New(publicLyeInventoryUpdateError)
	}

	mods := GetLyeInventoryPreloadModsWithLevel(ctx, LyeInventoryPayloadPreloadLevels.LyeInventory)
	mods = append(mods, dm.LyeInventoryWhere.ID.EQ(dbID))

	pM, err := dm.LyeInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryUpdateError)
		return nil, errors.New(publicLyeInventoryUpdateError)
	}
	return &fm.LyeInventoryPayload{
		LyeInventory: LyeInventoryToGraphQL(pM),
	}, nil
}

const publicLyeInventoryBatchUpdateError = "could not update lyeInventories"

func (r *mutationResolver) UpdateLyeInventories(ctx context.Context, filter *fm.LyeInventoryFilter, input fm.LyeInventoryUpdateInput) (*fm.LyeInventoriesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LyeInventoryFilterToMods(filter)...)

	m := LyeInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.LyeInventories(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryBatchUpdateError)
		return nil, errors.New(publicLyeInventoryBatchUpdateError)
	}

	return &fm.LyeInventoriesUpdatePayload{
		Ok: true,
	}, nil
}

const publicLyeInventoryDeleteError = "could not delete lyeInventory"

func (r *mutationResolver) DeleteLyeInventory(ctx context.Context, id string) (*fm.LyeInventoryDeletePayload, error) {
	dbID := LyeInventoryID(id)
	mods := []qm.QueryMod{
		dm.LyeInventoryWhere.ID.EQ(dbID),
	}
	if _, err := dm.LyeInventories(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryDeleteError)
		return nil, errors.New(publicLyeInventoryDeleteError)
	}

	return &fm.LyeInventoryDeletePayload{
		ID: id,
	}, nil
}

const publicLyeInventoryBatchDeleteError = "could not delete lyeInventories"

func (r *mutationResolver) DeleteLyeInventories(ctx context.Context, filter *fm.LyeInventoryFilter) (*fm.LyeInventoriesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, LyeInventoryFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.LyeInventoryColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.LyeInventory))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.LyeInventories(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryBatchDeleteError)
		return nil, errors.New(publicLyeInventoryBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.LyeInventories(dm.LyeInventoryWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryBatchDeleteError)
		return nil, errors.New(publicLyeInventoryBatchDeleteError)
	}

	return &fm.LyeInventoriesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.LyeInventory),
	}, nil
}

const publicRecipeCreateError = "could not create recipe"

func (r *mutationResolver) CreateRecipe(ctx context.Context, input fm.RecipeCreateInput) (*fm.RecipePayload, error) {

	m := RecipeCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeCreateError)
		return nil, errors.New(publicRecipeCreateError)
	}

	mods := GetRecipePreloadModsWithLevel(ctx, RecipePayloadPreloadLevels.Recipe)
	mods = append(mods, dm.RecipeWhere.ID.EQ(m.ID))
	pM, err := dm.Recipes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeCreateError)
		return nil, errors.New(publicRecipeCreateError)
	}
	return &fm.RecipePayload{
		Recipe: RecipeToGraphQL(pM),
	}, nil
}

const publicRecipeUpdateError = "could not update recipe"

func (r *mutationResolver) UpdateRecipe(ctx context.Context, id string, input fm.RecipeUpdateInput) (*fm.RecipePayload, error) {
	m := RecipeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeID(id)
	if _, err := dm.Recipes(
		dm.RecipeWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeUpdateError)
		return nil, errors.New(publicRecipeUpdateError)
	}

	mods := GetRecipePreloadModsWithLevel(ctx, RecipePayloadPreloadLevels.Recipe)
	mods = append(mods, dm.RecipeWhere.ID.EQ(dbID))

	pM, err := dm.Recipes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeUpdateError)
		return nil, errors.New(publicRecipeUpdateError)
	}
	return &fm.RecipePayload{
		Recipe: RecipeToGraphQL(pM),
	}, nil
}

const publicRecipeBatchUpdateError = "could not update recipes"

func (r *mutationResolver) UpdateRecipes(ctx context.Context, filter *fm.RecipeFilter, input fm.RecipeUpdateInput) (*fm.RecipesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeFilterToMods(filter)...)

	m := RecipeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Recipes(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchUpdateError)
		return nil, errors.New(publicRecipeBatchUpdateError)
	}

	return &fm.RecipesUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeDeleteError = "could not delete recipe"

func (r *mutationResolver) DeleteRecipe(ctx context.Context, id string) (*fm.RecipeDeletePayload, error) {
	dbID := RecipeID(id)
	mods := []qm.QueryMod{
		dm.RecipeWhere.ID.EQ(dbID),
	}
	if _, err := dm.Recipes(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeDeleteError)
		return nil, errors.New(publicRecipeDeleteError)
	}

	return &fm.RecipeDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchDeleteError = "could not delete recipes"

func (r *mutationResolver) DeleteRecipes(ctx context.Context, filter *fm.RecipeFilter) (*fm.RecipesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Recipe))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.Recipes(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchDeleteError)
		return nil, errors.New(publicRecipeBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.Recipes(dm.RecipeWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchDeleteError)
		return nil, errors.New(publicRecipeBatchDeleteError)
	}

	return &fm.RecipesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.Recipe),
	}, nil
}

const publicRecipeAdditiveCreateError = "could not create recipeAdditive"

func (r *mutationResolver) CreateRecipeAdditive(ctx context.Context, input fm.RecipeAdditiveCreateInput) (*fm.RecipeAdditivePayload, error) {

	m := RecipeAdditiveCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveCreateError)
		return nil, errors.New(publicRecipeAdditiveCreateError)
	}

	mods := GetRecipeAdditivePreloadModsWithLevel(ctx, RecipeAdditivePayloadPreloadLevels.RecipeAdditive)
	mods = append(mods, dm.RecipeAdditiveWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeAdditives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveCreateError)
		return nil, errors.New(publicRecipeAdditiveCreateError)
	}
	return &fm.RecipeAdditivePayload{
		RecipeAdditive: RecipeAdditiveToGraphQL(pM),
	}, nil
}

const publicRecipeAdditiveUpdateError = "could not update recipeAdditive"

func (r *mutationResolver) UpdateRecipeAdditive(ctx context.Context, id string, input fm.RecipeAdditiveUpdateInput) (*fm.RecipeAdditivePayload, error) {
	m := RecipeAdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeAdditiveID(id)
	if _, err := dm.RecipeAdditives(
		dm.RecipeAdditiveWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveUpdateError)
		return nil, errors.New(publicRecipeAdditiveUpdateError)
	}

	mods := GetRecipeAdditivePreloadModsWithLevel(ctx, RecipeAdditivePayloadPreloadLevels.RecipeAdditive)
	mods = append(mods, dm.RecipeAdditiveWhere.ID.EQ(dbID))

	pM, err := dm.RecipeAdditives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveUpdateError)
		return nil, errors.New(publicRecipeAdditiveUpdateError)
	}
	return &fm.RecipeAdditivePayload{
		RecipeAdditive: RecipeAdditiveToGraphQL(pM),
	}, nil
}

const publicRecipeAdditiveBatchUpdateError = "could not update recipeAdditives"

func (r *mutationResolver) UpdateRecipeAdditives(ctx context.Context, filter *fm.RecipeAdditiveFilter, input fm.RecipeAdditiveUpdateInput) (*fm.RecipeAdditivesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeAdditiveFilterToMods(filter)...)

	m := RecipeAdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeAdditives(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveBatchUpdateError)
		return nil, errors.New(publicRecipeAdditiveBatchUpdateError)
	}

	return &fm.RecipeAdditivesUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeAdditiveDeleteError = "could not delete recipeAdditive"

func (r *mutationResolver) DeleteRecipeAdditive(ctx context.Context, id string) (*fm.RecipeAdditiveDeletePayload, error) {
	dbID := RecipeAdditiveID(id)
	mods := []qm.QueryMod{
		dm.RecipeAdditiveWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeAdditives(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveDeleteError)
		return nil, errors.New(publicRecipeAdditiveDeleteError)
	}

	return &fm.RecipeAdditiveDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeAdditiveBatchDeleteError = "could not delete recipeAdditives"

func (r *mutationResolver) DeleteRecipeAdditives(ctx context.Context, filter *fm.RecipeAdditiveFilter) (*fm.RecipeAdditivesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeAdditiveFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeAdditiveColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeAdditive))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeAdditives(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveBatchDeleteError)
		return nil, errors.New(publicRecipeAdditiveBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeAdditives(dm.RecipeAdditiveWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveBatchDeleteError)
		return nil, errors.New(publicRecipeAdditiveBatchDeleteError)
	}

	return &fm.RecipeAdditivesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeAdditive),
	}, nil
}

const publicOneRecipeBatchCreateError = "could not create recipeBatch"

func (r *mutationResolver) CreateRecipeBatch(ctx context.Context, input fm.RecipeBatchCreateInput) (*fm.RecipeBatchPayload, error) {

	m := RecipeBatchCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicOneRecipeBatchCreateError)
		return nil, errors.New(publicOneRecipeBatchCreateError)
	}

	mods := GetRecipeBatchPreloadModsWithLevel(ctx, RecipeBatchPayloadPreloadLevels.RecipeBatch)
	mods = append(mods, dm.RecipeBatchWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeBatches(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicOneRecipeBatchCreateError)
		return nil, errors.New(publicOneRecipeBatchCreateError)
	}
	return &fm.RecipeBatchPayload{
		RecipeBatch: RecipeBatchToGraphQL(pM),
	}, nil
}

const publicOneRecipeBatchUpdateError = "could not update recipeBatch"

func (r *mutationResolver) UpdateRecipeBatch(ctx context.Context, id string, input fm.RecipeBatchUpdateInput) (*fm.RecipeBatchPayload, error) {
	m := RecipeBatchUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchID(id)
	if _, err := dm.RecipeBatches(
		dm.RecipeBatchWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicOneRecipeBatchUpdateError)
		return nil, errors.New(publicOneRecipeBatchUpdateError)
	}

	mods := GetRecipeBatchPreloadModsWithLevel(ctx, RecipeBatchPayloadPreloadLevels.RecipeBatch)
	mods = append(mods, dm.RecipeBatchWhere.ID.EQ(dbID))

	pM, err := dm.RecipeBatches(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicOneRecipeBatchUpdateError)
		return nil, errors.New(publicOneRecipeBatchUpdateError)
	}
	return &fm.RecipeBatchPayload{
		RecipeBatch: RecipeBatchToGraphQL(pM),
	}, nil
}

const publicRecipeBatchBatchUpdateError = "could not update recipeBatches"

func (r *mutationResolver) UpdateRecipeBatches(ctx context.Context, filter *fm.RecipeBatchFilter, input fm.RecipeBatchUpdateInput) (*fm.RecipeBatchesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchFilterToMods(filter)...)

	m := RecipeBatchUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeBatches(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchBatchUpdateError)
		return nil, errors.New(publicRecipeBatchBatchUpdateError)
	}

	return &fm.RecipeBatchesUpdatePayload{
		Ok: true,
	}, nil
}

const publicOneRecipeBatchDeleteError = "could not delete recipeBatch"

func (r *mutationResolver) DeleteRecipeBatch(ctx context.Context, id string) (*fm.RecipeBatchDeletePayload, error) {
	dbID := RecipeBatchID(id)
	mods := []qm.QueryMod{
		dm.RecipeBatchWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeBatches(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicOneRecipeBatchDeleteError)
		return nil, errors.New(publicOneRecipeBatchDeleteError)
	}

	return &fm.RecipeBatchDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchBatchDeleteError = "could not delete recipeBatches"

func (r *mutationResolver) DeleteRecipeBatches(ctx context.Context, filter *fm.RecipeBatchFilter) (*fm.RecipeBatchesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeBatchColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeBatch))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeBatches(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchBatchDeleteError)
		return nil, errors.New(publicRecipeBatchBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeBatches(dm.RecipeBatchWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchBatchDeleteError)
		return nil, errors.New(publicRecipeBatchBatchDeleteError)
	}

	return &fm.RecipeBatchesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeBatch),
	}, nil
}

const publicRecipeBatchAdditiveCreateError = "could not create recipeBatchAdditive"

func (r *mutationResolver) CreateRecipeBatchAdditive(ctx context.Context, input fm.RecipeBatchAdditiveCreateInput) (*fm.RecipeBatchAdditivePayload, error) {

	m := RecipeBatchAdditiveCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveCreateError)
		return nil, errors.New(publicRecipeBatchAdditiveCreateError)
	}

	mods := GetRecipeBatchAdditivePreloadModsWithLevel(ctx, RecipeBatchAdditivePayloadPreloadLevels.RecipeBatchAdditive)
	mods = append(mods, dm.RecipeBatchAdditiveWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeBatchAdditives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveCreateError)
		return nil, errors.New(publicRecipeBatchAdditiveCreateError)
	}
	return &fm.RecipeBatchAdditivePayload{
		RecipeBatchAdditive: RecipeBatchAdditiveToGraphQL(pM),
	}, nil
}

const publicRecipeBatchAdditiveUpdateError = "could not update recipeBatchAdditive"

func (r *mutationResolver) UpdateRecipeBatchAdditive(ctx context.Context, id string, input fm.RecipeBatchAdditiveUpdateInput) (*fm.RecipeBatchAdditivePayload, error) {
	m := RecipeBatchAdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchAdditiveID(id)
	if _, err := dm.RecipeBatchAdditives(
		dm.RecipeBatchAdditiveWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveUpdateError)
		return nil, errors.New(publicRecipeBatchAdditiveUpdateError)
	}

	mods := GetRecipeBatchAdditivePreloadModsWithLevel(ctx, RecipeBatchAdditivePayloadPreloadLevels.RecipeBatchAdditive)
	mods = append(mods, dm.RecipeBatchAdditiveWhere.ID.EQ(dbID))

	pM, err := dm.RecipeBatchAdditives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveUpdateError)
		return nil, errors.New(publicRecipeBatchAdditiveUpdateError)
	}
	return &fm.RecipeBatchAdditivePayload{
		RecipeBatchAdditive: RecipeBatchAdditiveToGraphQL(pM),
	}, nil
}

const publicRecipeBatchAdditiveBatchUpdateError = "could not update recipeBatchAdditives"

func (r *mutationResolver) UpdateRecipeBatchAdditives(ctx context.Context, filter *fm.RecipeBatchAdditiveFilter, input fm.RecipeBatchAdditiveUpdateInput) (*fm.RecipeBatchAdditivesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchAdditiveFilterToMods(filter)...)

	m := RecipeBatchAdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeBatchAdditives(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveBatchUpdateError)
		return nil, errors.New(publicRecipeBatchAdditiveBatchUpdateError)
	}

	return &fm.RecipeBatchAdditivesUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeBatchAdditiveDeleteError = "could not delete recipeBatchAdditive"

func (r *mutationResolver) DeleteRecipeBatchAdditive(ctx context.Context, id string) (*fm.RecipeBatchAdditiveDeletePayload, error) {
	dbID := RecipeBatchAdditiveID(id)
	mods := []qm.QueryMod{
		dm.RecipeBatchAdditiveWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeBatchAdditives(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveDeleteError)
		return nil, errors.New(publicRecipeBatchAdditiveDeleteError)
	}

	return &fm.RecipeBatchAdditiveDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchAdditiveBatchDeleteError = "could not delete recipeBatchAdditives"

func (r *mutationResolver) DeleteRecipeBatchAdditives(ctx context.Context, filter *fm.RecipeBatchAdditiveFilter) (*fm.RecipeBatchAdditivesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchAdditiveFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeBatchAdditiveColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeBatchAdditive))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeBatchAdditives(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveBatchDeleteError)
		return nil, errors.New(publicRecipeBatchAdditiveBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeBatchAdditives(dm.RecipeBatchAdditiveWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveBatchDeleteError)
		return nil, errors.New(publicRecipeBatchAdditiveBatchDeleteError)
	}

	return &fm.RecipeBatchAdditivesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeBatchAdditive),
	}, nil
}

const publicRecipeBatchFragranceCreateError = "could not create recipeBatchFragrance"

func (r *mutationResolver) CreateRecipeBatchFragrance(ctx context.Context, input fm.RecipeBatchFragranceCreateInput) (*fm.RecipeBatchFragrancePayload, error) {

	m := RecipeBatchFragranceCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceCreateError)
		return nil, errors.New(publicRecipeBatchFragranceCreateError)
	}

	mods := GetRecipeBatchFragrancePreloadModsWithLevel(ctx, RecipeBatchFragrancePayloadPreloadLevels.RecipeBatchFragrance)
	mods = append(mods, dm.RecipeBatchFragranceWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeBatchFragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceCreateError)
		return nil, errors.New(publicRecipeBatchFragranceCreateError)
	}
	return &fm.RecipeBatchFragrancePayload{
		RecipeBatchFragrance: RecipeBatchFragranceToGraphQL(pM),
	}, nil
}

const publicRecipeBatchFragranceUpdateError = "could not update recipeBatchFragrance"

func (r *mutationResolver) UpdateRecipeBatchFragrance(ctx context.Context, id string, input fm.RecipeBatchFragranceUpdateInput) (*fm.RecipeBatchFragrancePayload, error) {
	m := RecipeBatchFragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchFragranceID(id)
	if _, err := dm.RecipeBatchFragrances(
		dm.RecipeBatchFragranceWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceUpdateError)
		return nil, errors.New(publicRecipeBatchFragranceUpdateError)
	}

	mods := GetRecipeBatchFragrancePreloadModsWithLevel(ctx, RecipeBatchFragrancePayloadPreloadLevels.RecipeBatchFragrance)
	mods = append(mods, dm.RecipeBatchFragranceWhere.ID.EQ(dbID))

	pM, err := dm.RecipeBatchFragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceUpdateError)
		return nil, errors.New(publicRecipeBatchFragranceUpdateError)
	}
	return &fm.RecipeBatchFragrancePayload{
		RecipeBatchFragrance: RecipeBatchFragranceToGraphQL(pM),
	}, nil
}

const publicRecipeBatchFragranceBatchUpdateError = "could not update recipeBatchFragrances"

func (r *mutationResolver) UpdateRecipeBatchFragrances(ctx context.Context, filter *fm.RecipeBatchFragranceFilter, input fm.RecipeBatchFragranceUpdateInput) (*fm.RecipeBatchFragrancesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchFragranceFilterToMods(filter)...)

	m := RecipeBatchFragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeBatchFragrances(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceBatchUpdateError)
		return nil, errors.New(publicRecipeBatchFragranceBatchUpdateError)
	}

	return &fm.RecipeBatchFragrancesUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeBatchFragranceDeleteError = "could not delete recipeBatchFragrance"

func (r *mutationResolver) DeleteRecipeBatchFragrance(ctx context.Context, id string) (*fm.RecipeBatchFragranceDeletePayload, error) {
	dbID := RecipeBatchFragranceID(id)
	mods := []qm.QueryMod{
		dm.RecipeBatchFragranceWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeBatchFragrances(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceDeleteError)
		return nil, errors.New(publicRecipeBatchFragranceDeleteError)
	}

	return &fm.RecipeBatchFragranceDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchFragranceBatchDeleteError = "could not delete recipeBatchFragrances"

func (r *mutationResolver) DeleteRecipeBatchFragrances(ctx context.Context, filter *fm.RecipeBatchFragranceFilter) (*fm.RecipeBatchFragrancesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchFragranceFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeBatchFragranceColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeBatchFragrance))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeBatchFragrances(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceBatchDeleteError)
		return nil, errors.New(publicRecipeBatchFragranceBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeBatchFragrances(dm.RecipeBatchFragranceWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceBatchDeleteError)
		return nil, errors.New(publicRecipeBatchFragranceBatchDeleteError)
	}

	return &fm.RecipeBatchFragrancesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeBatchFragrance),
	}, nil
}

const publicRecipeBatchLipidCreateError = "could not create recipeBatchLipid"

func (r *mutationResolver) CreateRecipeBatchLipid(ctx context.Context, input fm.RecipeBatchLipidCreateInput) (*fm.RecipeBatchLipidPayload, error) {

	m := RecipeBatchLipidCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidCreateError)
		return nil, errors.New(publicRecipeBatchLipidCreateError)
	}

	mods := GetRecipeBatchLipidPreloadModsWithLevel(ctx, RecipeBatchLipidPayloadPreloadLevels.RecipeBatchLipid)
	mods = append(mods, dm.RecipeBatchLipidWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeBatchLipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidCreateError)
		return nil, errors.New(publicRecipeBatchLipidCreateError)
	}
	return &fm.RecipeBatchLipidPayload{
		RecipeBatchLipid: RecipeBatchLipidToGraphQL(pM),
	}, nil
}

const publicRecipeBatchLipidUpdateError = "could not update recipeBatchLipid"

func (r *mutationResolver) UpdateRecipeBatchLipid(ctx context.Context, id string, input fm.RecipeBatchLipidUpdateInput) (*fm.RecipeBatchLipidPayload, error) {
	m := RecipeBatchLipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchLipidID(id)
	if _, err := dm.RecipeBatchLipids(
		dm.RecipeBatchLipidWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidUpdateError)
		return nil, errors.New(publicRecipeBatchLipidUpdateError)
	}

	mods := GetRecipeBatchLipidPreloadModsWithLevel(ctx, RecipeBatchLipidPayloadPreloadLevels.RecipeBatchLipid)
	mods = append(mods, dm.RecipeBatchLipidWhere.ID.EQ(dbID))

	pM, err := dm.RecipeBatchLipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidUpdateError)
		return nil, errors.New(publicRecipeBatchLipidUpdateError)
	}
	return &fm.RecipeBatchLipidPayload{
		RecipeBatchLipid: RecipeBatchLipidToGraphQL(pM),
	}, nil
}

const publicRecipeBatchLipidBatchUpdateError = "could not update recipeBatchLipids"

func (r *mutationResolver) UpdateRecipeBatchLipids(ctx context.Context, filter *fm.RecipeBatchLipidFilter, input fm.RecipeBatchLipidUpdateInput) (*fm.RecipeBatchLipidsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchLipidFilterToMods(filter)...)

	m := RecipeBatchLipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeBatchLipids(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidBatchUpdateError)
		return nil, errors.New(publicRecipeBatchLipidBatchUpdateError)
	}

	return &fm.RecipeBatchLipidsUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeBatchLipidDeleteError = "could not delete recipeBatchLipid"

func (r *mutationResolver) DeleteRecipeBatchLipid(ctx context.Context, id string) (*fm.RecipeBatchLipidDeletePayload, error) {
	dbID := RecipeBatchLipidID(id)
	mods := []qm.QueryMod{
		dm.RecipeBatchLipidWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeBatchLipids(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidDeleteError)
		return nil, errors.New(publicRecipeBatchLipidDeleteError)
	}

	return &fm.RecipeBatchLipidDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchLipidBatchDeleteError = "could not delete recipeBatchLipids"

func (r *mutationResolver) DeleteRecipeBatchLipids(ctx context.Context, filter *fm.RecipeBatchLipidFilter) (*fm.RecipeBatchLipidsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchLipidFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeBatchLipidColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeBatchLipid))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeBatchLipids(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidBatchDeleteError)
		return nil, errors.New(publicRecipeBatchLipidBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeBatchLipids(dm.RecipeBatchLipidWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidBatchDeleteError)
		return nil, errors.New(publicRecipeBatchLipidBatchDeleteError)
	}

	return &fm.RecipeBatchLipidsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeBatchLipid),
	}, nil
}

const publicRecipeBatchLyeCreateError = "could not create recipeBatchLye"

func (r *mutationResolver) CreateRecipeBatchLye(ctx context.Context, input fm.RecipeBatchLyeCreateInput) (*fm.RecipeBatchLyePayload, error) {

	m := RecipeBatchLyeCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeCreateError)
		return nil, errors.New(publicRecipeBatchLyeCreateError)
	}

	mods := GetRecipeBatchLyePreloadModsWithLevel(ctx, RecipeBatchLyePayloadPreloadLevels.RecipeBatchLye)
	mods = append(mods, dm.RecipeBatchLyeWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeBatchLyes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeCreateError)
		return nil, errors.New(publicRecipeBatchLyeCreateError)
	}
	return &fm.RecipeBatchLyePayload{
		RecipeBatchLye: RecipeBatchLyeToGraphQL(pM),
	}, nil
}

const publicRecipeBatchLyeUpdateError = "could not update recipeBatchLye"

func (r *mutationResolver) UpdateRecipeBatchLye(ctx context.Context, id string, input fm.RecipeBatchLyeUpdateInput) (*fm.RecipeBatchLyePayload, error) {
	m := RecipeBatchLyeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchLyeID(id)
	if _, err := dm.RecipeBatchLyes(
		dm.RecipeBatchLyeWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeUpdateError)
		return nil, errors.New(publicRecipeBatchLyeUpdateError)
	}

	mods := GetRecipeBatchLyePreloadModsWithLevel(ctx, RecipeBatchLyePayloadPreloadLevels.RecipeBatchLye)
	mods = append(mods, dm.RecipeBatchLyeWhere.ID.EQ(dbID))

	pM, err := dm.RecipeBatchLyes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeUpdateError)
		return nil, errors.New(publicRecipeBatchLyeUpdateError)
	}
	return &fm.RecipeBatchLyePayload{
		RecipeBatchLye: RecipeBatchLyeToGraphQL(pM),
	}, nil
}

const publicRecipeBatchLyeBatchUpdateError = "could not update recipeBatchLyes"

func (r *mutationResolver) UpdateRecipeBatchLyes(ctx context.Context, filter *fm.RecipeBatchLyeFilter, input fm.RecipeBatchLyeUpdateInput) (*fm.RecipeBatchLyesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchLyeFilterToMods(filter)...)

	m := RecipeBatchLyeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeBatchLyes(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeBatchUpdateError)
		return nil, errors.New(publicRecipeBatchLyeBatchUpdateError)
	}

	return &fm.RecipeBatchLyesUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeBatchLyeDeleteError = "could not delete recipeBatchLye"

func (r *mutationResolver) DeleteRecipeBatchLye(ctx context.Context, id string) (*fm.RecipeBatchLyeDeletePayload, error) {
	dbID := RecipeBatchLyeID(id)
	mods := []qm.QueryMod{
		dm.RecipeBatchLyeWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeBatchLyes(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeDeleteError)
		return nil, errors.New(publicRecipeBatchLyeDeleteError)
	}

	return &fm.RecipeBatchLyeDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchLyeBatchDeleteError = "could not delete recipeBatchLyes"

func (r *mutationResolver) DeleteRecipeBatchLyes(ctx context.Context, filter *fm.RecipeBatchLyeFilter) (*fm.RecipeBatchLyesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchLyeFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeBatchLyeColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeBatchLye))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeBatchLyes(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeBatchDeleteError)
		return nil, errors.New(publicRecipeBatchLyeBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeBatchLyes(dm.RecipeBatchLyeWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeBatchDeleteError)
		return nil, errors.New(publicRecipeBatchLyeBatchDeleteError)
	}

	return &fm.RecipeBatchLyesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeBatchLye),
	}, nil
}

const publicRecipeBatchNoteCreateError = "could not create recipeBatchNote"

func (r *mutationResolver) CreateRecipeBatchNote(ctx context.Context, input fm.RecipeBatchNoteCreateInput) (*fm.RecipeBatchNotePayload, error) {

	m := RecipeBatchNoteCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteCreateError)
		return nil, errors.New(publicRecipeBatchNoteCreateError)
	}

	mods := GetRecipeBatchNotePreloadModsWithLevel(ctx, RecipeBatchNotePayloadPreloadLevels.RecipeBatchNote)
	mods = append(mods, dm.RecipeBatchNoteWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeBatchNotes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteCreateError)
		return nil, errors.New(publicRecipeBatchNoteCreateError)
	}
	return &fm.RecipeBatchNotePayload{
		RecipeBatchNote: RecipeBatchNoteToGraphQL(pM),
	}, nil
}

const publicRecipeBatchNoteUpdateError = "could not update recipeBatchNote"

func (r *mutationResolver) UpdateRecipeBatchNote(ctx context.Context, id string, input fm.RecipeBatchNoteUpdateInput) (*fm.RecipeBatchNotePayload, error) {
	m := RecipeBatchNoteUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchNoteID(id)
	if _, err := dm.RecipeBatchNotes(
		dm.RecipeBatchNoteWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteUpdateError)
		return nil, errors.New(publicRecipeBatchNoteUpdateError)
	}

	mods := GetRecipeBatchNotePreloadModsWithLevel(ctx, RecipeBatchNotePayloadPreloadLevels.RecipeBatchNote)
	mods = append(mods, dm.RecipeBatchNoteWhere.ID.EQ(dbID))

	pM, err := dm.RecipeBatchNotes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteUpdateError)
		return nil, errors.New(publicRecipeBatchNoteUpdateError)
	}
	return &fm.RecipeBatchNotePayload{
		RecipeBatchNote: RecipeBatchNoteToGraphQL(pM),
	}, nil
}

const publicRecipeBatchNoteBatchUpdateError = "could not update recipeBatchNotes"

func (r *mutationResolver) UpdateRecipeBatchNotes(ctx context.Context, filter *fm.RecipeBatchNoteFilter, input fm.RecipeBatchNoteUpdateInput) (*fm.RecipeBatchNotesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchNoteFilterToMods(filter)...)

	m := RecipeBatchNoteUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeBatchNotes(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteBatchUpdateError)
		return nil, errors.New(publicRecipeBatchNoteBatchUpdateError)
	}

	return &fm.RecipeBatchNotesUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeBatchNoteDeleteError = "could not delete recipeBatchNote"

func (r *mutationResolver) DeleteRecipeBatchNote(ctx context.Context, id string) (*fm.RecipeBatchNoteDeletePayload, error) {
	dbID := RecipeBatchNoteID(id)
	mods := []qm.QueryMod{
		dm.RecipeBatchNoteWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeBatchNotes(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteDeleteError)
		return nil, errors.New(publicRecipeBatchNoteDeleteError)
	}

	return &fm.RecipeBatchNoteDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchNoteBatchDeleteError = "could not delete recipeBatchNotes"

func (r *mutationResolver) DeleteRecipeBatchNotes(ctx context.Context, filter *fm.RecipeBatchNoteFilter) (*fm.RecipeBatchNotesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeBatchNoteFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeBatchNoteColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeBatchNote))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeBatchNotes(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteBatchDeleteError)
		return nil, errors.New(publicRecipeBatchNoteBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeBatchNotes(dm.RecipeBatchNoteWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteBatchDeleteError)
		return nil, errors.New(publicRecipeBatchNoteBatchDeleteError)
	}

	return &fm.RecipeBatchNotesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeBatchNote),
	}, nil
}

const publicRecipeFragranceCreateError = "could not create recipeFragrance"

func (r *mutationResolver) CreateRecipeFragrance(ctx context.Context, input fm.RecipeFragranceCreateInput) (*fm.RecipeFragrancePayload, error) {

	m := RecipeFragranceCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceCreateError)
		return nil, errors.New(publicRecipeFragranceCreateError)
	}

	mods := GetRecipeFragrancePreloadModsWithLevel(ctx, RecipeFragrancePayloadPreloadLevels.RecipeFragrance)
	mods = append(mods, dm.RecipeFragranceWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeFragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceCreateError)
		return nil, errors.New(publicRecipeFragranceCreateError)
	}
	return &fm.RecipeFragrancePayload{
		RecipeFragrance: RecipeFragranceToGraphQL(pM),
	}, nil
}

const publicRecipeFragranceUpdateError = "could not update recipeFragrance"

func (r *mutationResolver) UpdateRecipeFragrance(ctx context.Context, id string, input fm.RecipeFragranceUpdateInput) (*fm.RecipeFragrancePayload, error) {
	m := RecipeFragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeFragranceID(id)
	if _, err := dm.RecipeFragrances(
		dm.RecipeFragranceWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceUpdateError)
		return nil, errors.New(publicRecipeFragranceUpdateError)
	}

	mods := GetRecipeFragrancePreloadModsWithLevel(ctx, RecipeFragrancePayloadPreloadLevels.RecipeFragrance)
	mods = append(mods, dm.RecipeFragranceWhere.ID.EQ(dbID))

	pM, err := dm.RecipeFragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceUpdateError)
		return nil, errors.New(publicRecipeFragranceUpdateError)
	}
	return &fm.RecipeFragrancePayload{
		RecipeFragrance: RecipeFragranceToGraphQL(pM),
	}, nil
}

const publicRecipeFragranceBatchUpdateError = "could not update recipeFragrances"

func (r *mutationResolver) UpdateRecipeFragrances(ctx context.Context, filter *fm.RecipeFragranceFilter, input fm.RecipeFragranceUpdateInput) (*fm.RecipeFragrancesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeFragranceFilterToMods(filter)...)

	m := RecipeFragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeFragrances(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceBatchUpdateError)
		return nil, errors.New(publicRecipeFragranceBatchUpdateError)
	}

	return &fm.RecipeFragrancesUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeFragranceDeleteError = "could not delete recipeFragrance"

func (r *mutationResolver) DeleteRecipeFragrance(ctx context.Context, id string) (*fm.RecipeFragranceDeletePayload, error) {
	dbID := RecipeFragranceID(id)
	mods := []qm.QueryMod{
		dm.RecipeFragranceWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeFragrances(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceDeleteError)
		return nil, errors.New(publicRecipeFragranceDeleteError)
	}

	return &fm.RecipeFragranceDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeFragranceBatchDeleteError = "could not delete recipeFragrances"

func (r *mutationResolver) DeleteRecipeFragrances(ctx context.Context, filter *fm.RecipeFragranceFilter) (*fm.RecipeFragrancesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeFragranceFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeFragranceColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeFragrance))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeFragrances(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceBatchDeleteError)
		return nil, errors.New(publicRecipeFragranceBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeFragrances(dm.RecipeFragranceWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceBatchDeleteError)
		return nil, errors.New(publicRecipeFragranceBatchDeleteError)
	}

	return &fm.RecipeFragrancesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeFragrance),
	}, nil
}

const publicRecipeLipidCreateError = "could not create recipeLipid"

func (r *mutationResolver) CreateRecipeLipid(ctx context.Context, input fm.RecipeLipidCreateInput) (*fm.RecipeLipidPayload, error) {

	m := RecipeLipidCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidCreateError)
		return nil, errors.New(publicRecipeLipidCreateError)
	}

	mods := GetRecipeLipidPreloadModsWithLevel(ctx, RecipeLipidPayloadPreloadLevels.RecipeLipid)
	mods = append(mods, dm.RecipeLipidWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeLipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidCreateError)
		return nil, errors.New(publicRecipeLipidCreateError)
	}
	return &fm.RecipeLipidPayload{
		RecipeLipid: RecipeLipidToGraphQL(pM),
	}, nil
}

const publicRecipeLipidUpdateError = "could not update recipeLipid"

func (r *mutationResolver) UpdateRecipeLipid(ctx context.Context, id string, input fm.RecipeLipidUpdateInput) (*fm.RecipeLipidPayload, error) {
	m := RecipeLipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeLipidID(id)
	if _, err := dm.RecipeLipids(
		dm.RecipeLipidWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidUpdateError)
		return nil, errors.New(publicRecipeLipidUpdateError)
	}

	mods := GetRecipeLipidPreloadModsWithLevel(ctx, RecipeLipidPayloadPreloadLevels.RecipeLipid)
	mods = append(mods, dm.RecipeLipidWhere.ID.EQ(dbID))

	pM, err := dm.RecipeLipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidUpdateError)
		return nil, errors.New(publicRecipeLipidUpdateError)
	}
	return &fm.RecipeLipidPayload{
		RecipeLipid: RecipeLipidToGraphQL(pM),
	}, nil
}

const publicRecipeLipidBatchUpdateError = "could not update recipeLipids"

func (r *mutationResolver) UpdateRecipeLipids(ctx context.Context, filter *fm.RecipeLipidFilter, input fm.RecipeLipidUpdateInput) (*fm.RecipeLipidsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeLipidFilterToMods(filter)...)

	m := RecipeLipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeLipids(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidBatchUpdateError)
		return nil, errors.New(publicRecipeLipidBatchUpdateError)
	}

	return &fm.RecipeLipidsUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeLipidDeleteError = "could not delete recipeLipid"

func (r *mutationResolver) DeleteRecipeLipid(ctx context.Context, id string) (*fm.RecipeLipidDeletePayload, error) {
	dbID := RecipeLipidID(id)
	mods := []qm.QueryMod{
		dm.RecipeLipidWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeLipids(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidDeleteError)
		return nil, errors.New(publicRecipeLipidDeleteError)
	}

	return &fm.RecipeLipidDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeLipidBatchDeleteError = "could not delete recipeLipids"

func (r *mutationResolver) DeleteRecipeLipids(ctx context.Context, filter *fm.RecipeLipidFilter) (*fm.RecipeLipidsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeLipidFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeLipidColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeLipid))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeLipids(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidBatchDeleteError)
		return nil, errors.New(publicRecipeLipidBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeLipids(dm.RecipeLipidWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidBatchDeleteError)
		return nil, errors.New(publicRecipeLipidBatchDeleteError)
	}

	return &fm.RecipeLipidsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeLipid),
	}, nil
}

const publicRecipeStepCreateError = "could not create recipeStep"

func (r *mutationResolver) CreateRecipeStep(ctx context.Context, input fm.RecipeStepCreateInput) (*fm.RecipeStepPayload, error) {

	m := RecipeStepCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepCreateError)
		return nil, errors.New(publicRecipeStepCreateError)
	}

	mods := GetRecipeStepPreloadModsWithLevel(ctx, RecipeStepPayloadPreloadLevels.RecipeStep)
	mods = append(mods, dm.RecipeStepWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeSteps(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeStepCreateError)
		return nil, errors.New(publicRecipeStepCreateError)
	}
	return &fm.RecipeStepPayload{
		RecipeStep: RecipeStepToGraphQL(pM),
	}, nil
}

const publicRecipeStepUpdateError = "could not update recipeStep"

func (r *mutationResolver) UpdateRecipeStep(ctx context.Context, id string, input fm.RecipeStepUpdateInput) (*fm.RecipeStepPayload, error) {
	m := RecipeStepUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeStepID(id)
	if _, err := dm.RecipeSteps(
		dm.RecipeStepWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepUpdateError)
		return nil, errors.New(publicRecipeStepUpdateError)
	}

	mods := GetRecipeStepPreloadModsWithLevel(ctx, RecipeStepPayloadPreloadLevels.RecipeStep)
	mods = append(mods, dm.RecipeStepWhere.ID.EQ(dbID))

	pM, err := dm.RecipeSteps(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeStepUpdateError)
		return nil, errors.New(publicRecipeStepUpdateError)
	}
	return &fm.RecipeStepPayload{
		RecipeStep: RecipeStepToGraphQL(pM),
	}, nil
}

const publicRecipeStepBatchUpdateError = "could not update recipeSteps"

func (r *mutationResolver) UpdateRecipeSteps(ctx context.Context, filter *fm.RecipeStepFilter, input fm.RecipeStepUpdateInput) (*fm.RecipeStepsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeStepFilterToMods(filter)...)

	m := RecipeStepUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.RecipeSteps(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepBatchUpdateError)
		return nil, errors.New(publicRecipeStepBatchUpdateError)
	}

	return &fm.RecipeStepsUpdatePayload{
		Ok: true,
	}, nil
}

const publicRecipeStepDeleteError = "could not delete recipeStep"

func (r *mutationResolver) DeleteRecipeStep(ctx context.Context, id string) (*fm.RecipeStepDeletePayload, error) {
	dbID := RecipeStepID(id)
	mods := []qm.QueryMod{
		dm.RecipeStepWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeSteps(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepDeleteError)
		return nil, errors.New(publicRecipeStepDeleteError)
	}

	return &fm.RecipeStepDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeStepBatchDeleteError = "could not delete recipeSteps"

func (r *mutationResolver) DeleteRecipeSteps(ctx context.Context, filter *fm.RecipeStepFilter) (*fm.RecipeStepsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, RecipeStepFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.RecipeStepColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.RecipeStep))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.RecipeSteps(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepBatchDeleteError)
		return nil, errors.New(publicRecipeStepBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.RecipeSteps(dm.RecipeStepWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepBatchDeleteError)
		return nil, errors.New(publicRecipeStepBatchDeleteError)
	}

	return &fm.RecipeStepsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.RecipeStep),
	}, nil
}

const publicSupplierCreateError = "could not create supplier"

func (r *mutationResolver) CreateSupplier(ctx context.Context, input fm.SupplierCreateInput) (*fm.SupplierPayload, error) {

	m := SupplierCreateInputToBoiler(&input)

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		log.Error().Err(err).Msg(publicSupplierCreateError)
		return nil, errors.New(publicSupplierCreateError)
	}

	mods := GetSupplierPreloadModsWithLevel(ctx, SupplierPayloadPreloadLevels.Supplier)
	mods = append(mods, dm.SupplierWhere.ID.EQ(m.ID))
	pM, err := dm.Suppliers(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicSupplierCreateError)
		return nil, errors.New(publicSupplierCreateError)
	}
	return &fm.SupplierPayload{
		Supplier: SupplierToGraphQL(pM),
	}, nil
}

const publicSupplierUpdateError = "could not update supplier"

func (r *mutationResolver) UpdateSupplier(ctx context.Context, id string, input fm.SupplierUpdateInput) (*fm.SupplierPayload, error) {
	m := SupplierUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := SupplierID(id)
	if _, err := dm.Suppliers(
		dm.SupplierWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicSupplierUpdateError)
		return nil, errors.New(publicSupplierUpdateError)
	}

	mods := GetSupplierPreloadModsWithLevel(ctx, SupplierPayloadPreloadLevels.Supplier)
	mods = append(mods, dm.SupplierWhere.ID.EQ(dbID))

	pM, err := dm.Suppliers(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicSupplierUpdateError)
		return nil, errors.New(publicSupplierUpdateError)
	}
	return &fm.SupplierPayload{
		Supplier: SupplierToGraphQL(pM),
	}, nil
}

const publicSupplierBatchUpdateError = "could not update suppliers"

func (r *mutationResolver) UpdateSuppliers(ctx context.Context, filter *fm.SupplierFilter, input fm.SupplierUpdateInput) (*fm.SuppliersUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, SupplierFilterToMods(filter)...)

	m := SupplierUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Suppliers(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicSupplierBatchUpdateError)
		return nil, errors.New(publicSupplierBatchUpdateError)
	}

	return &fm.SuppliersUpdatePayload{
		Ok: true,
	}, nil
}

const publicSupplierDeleteError = "could not delete supplier"

func (r *mutationResolver) DeleteSupplier(ctx context.Context, id string) (*fm.SupplierDeletePayload, error) {
	dbID := SupplierID(id)
	mods := []qm.QueryMod{
		dm.SupplierWhere.ID.EQ(dbID),
	}
	if _, err := dm.Suppliers(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicSupplierDeleteError)
		return nil, errors.New(publicSupplierDeleteError)
	}

	return &fm.SupplierDeletePayload{
		ID: id,
	}, nil
}

const publicSupplierBatchDeleteError = "could not delete suppliers"

func (r *mutationResolver) DeleteSuppliers(ctx context.Context, filter *fm.SupplierFilter) (*fm.SuppliersDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, SupplierFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.SupplierColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Supplier))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.Suppliers(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicSupplierBatchDeleteError)
		return nil, errors.New(publicSupplierBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.Suppliers(dm.SupplierWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicSupplierBatchDeleteError)
		return nil, errors.New(publicSupplierBatchDeleteError)
	}

	return &fm.SuppliersDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.Supplier),
	}, nil
}

const publicAdditiveSingleError = "could not get additive"

func (r *queryResolver) Additive(ctx context.Context, id string) (*fm.Additive, error) {
	dbID := AdditiveID(id)
	mods := GetAdditivePreloadMods(ctx)
	mods = append(mods, dm.AdditiveWhere.ID.EQ(dbID))

	m, err := dm.Additives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveSingleError)
		return nil, errors.New(publicAdditiveSingleError)
	}
	return AdditiveToGraphQL(m), nil
}

const publicAdditiveListError = "could not list additives"

func (r *queryResolver) Additives(ctx context.Context, first int, after *string, ordering []*fm.AdditiveOrdering, filter *fm.AdditiveFilter) (*fm.AdditiveConnection, error) {
	mods := GetAdditiveNodePreloadMods(ctx)

	mods = append(mods, AdditiveFilterToMods(filter)...)
	connection, err := AdditiveConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveListError)
		return nil, errors.New(publicAdditiveListError)
	}
	return connection, nil
}

const publicAdditiveInventorySingleError = "could not get additiveInventory"

func (r *queryResolver) AdditiveInventory(ctx context.Context, id string) (*fm.AdditiveInventory, error) {
	dbID := AdditiveInventoryID(id)
	mods := GetAdditiveInventoryPreloadMods(ctx)
	mods = append(mods, dm.AdditiveInventoryWhere.ID.EQ(dbID))

	m, err := dm.AdditiveInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventorySingleError)
		return nil, errors.New(publicAdditiveInventorySingleError)
	}
	return AdditiveInventoryToGraphQL(m), nil
}

const publicAdditiveInventoryListError = "could not list additiveInventories"

func (r *queryResolver) AdditiveInventories(ctx context.Context, first int, after *string, ordering []*fm.AdditiveInventoryOrdering, filter *fm.AdditiveInventoryFilter) (*fm.AdditiveInventoryConnection, error) {
	mods := GetAdditiveInventoryNodePreloadMods(ctx)

	mods = append(mods, AdditiveInventoryFilterToMods(filter)...)
	connection, err := AdditiveInventoryConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryListError)
		return nil, errors.New(publicAdditiveInventoryListError)
	}
	return connection, nil
}

const publicAuthGroupSingleError = "could not get authGroup"

func (r *queryResolver) AuthGroup(ctx context.Context, id string) (*fm.AuthGroup, error) {
	dbID := AuthGroupID(id)
	mods := GetAuthGroupPreloadMods(ctx)
	mods = append(mods, dm.AuthGroupWhere.ID.EQ(dbID))

	m, err := dm.AuthGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupSingleError)
		return nil, errors.New(publicAuthGroupSingleError)
	}
	return AuthGroupToGraphQL(m), nil
}

const publicAuthGroupListError = "could not list authGroups"

func (r *queryResolver) AuthGroups(ctx context.Context, first int, after *string, ordering []*fm.AuthGroupOrdering, filter *fm.AuthGroupFilter) (*fm.AuthGroupConnection, error) {
	mods := GetAuthGroupNodePreloadMods(ctx)

	mods = append(mods, AuthGroupFilterToMods(filter)...)
	connection, err := AuthGroupConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupListError)
		return nil, errors.New(publicAuthGroupListError)
	}
	return connection, nil
}

const publicAuthGroupPermissionSingleError = "could not get authGroupPermission"

func (r *queryResolver) AuthGroupPermission(ctx context.Context, id string) (*fm.AuthGroupPermission, error) {
	dbID := AuthGroupPermissionID(id)
	mods := GetAuthGroupPermissionPreloadMods(ctx)
	mods = append(mods, dm.AuthGroupPermissionWhere.ID.EQ(dbID))

	m, err := dm.AuthGroupPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionSingleError)
		return nil, errors.New(publicAuthGroupPermissionSingleError)
	}
	return AuthGroupPermissionToGraphQL(m), nil
}

const publicAuthGroupPermissionListError = "could not list authGroupPermissions"

func (r *queryResolver) AuthGroupPermissions(ctx context.Context, first int, after *string, ordering []*fm.AuthGroupPermissionOrdering, filter *fm.AuthGroupPermissionFilter) (*fm.AuthGroupPermissionConnection, error) {
	mods := GetAuthGroupPermissionNodePreloadMods(ctx)

	mods = append(mods, AuthGroupPermissionFilterToMods(filter)...)
	connection, err := AuthGroupPermissionConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionListError)
		return nil, errors.New(publicAuthGroupPermissionListError)
	}
	return connection, nil
}

const publicAuthPermissionSingleError = "could not get authPermission"

func (r *queryResolver) AuthPermission(ctx context.Context, id string) (*fm.AuthPermission, error) {
	dbID := AuthPermissionID(id)
	mods := GetAuthPermissionPreloadMods(ctx)
	mods = append(mods, dm.AuthPermissionWhere.ID.EQ(dbID))

	m, err := dm.AuthPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionSingleError)
		return nil, errors.New(publicAuthPermissionSingleError)
	}
	return AuthPermissionToGraphQL(m), nil
}

const publicAuthPermissionListError = "could not list authPermissions"

func (r *queryResolver) AuthPermissions(ctx context.Context, first int, after *string, ordering []*fm.AuthPermissionOrdering, filter *fm.AuthPermissionFilter) (*fm.AuthPermissionConnection, error) {
	mods := GetAuthPermissionNodePreloadMods(ctx)

	mods = append(mods, AuthPermissionFilterToMods(filter)...)
	connection, err := AuthPermissionConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionListError)
		return nil, errors.New(publicAuthPermissionListError)
	}
	return connection, nil
}

const publicAuthUserSingleError = "could not get authUser"

func (r *queryResolver) AuthUser(ctx context.Context, id string) (*fm.AuthUser, error) {
	dbID := AuthUserID(id)
	mods := GetAuthUserPreloadMods(ctx)
	mods = append(mods, dm.AuthUserWhere.ID.EQ(dbID))

	m, err := dm.AuthUsers(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserSingleError)
		return nil, errors.New(publicAuthUserSingleError)
	}
	return AuthUserToGraphQL(m), nil
}

const publicAuthUserListError = "could not list authUsers"

func (r *queryResolver) AuthUsers(ctx context.Context, first int, after *string, ordering []*fm.AuthUserOrdering, filter *fm.AuthUserFilter) (*fm.AuthUserConnection, error) {
	mods := GetAuthUserNodePreloadMods(ctx)

	mods = append(mods, AuthUserFilterToMods(filter)...)
	connection, err := AuthUserConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserListError)
		return nil, errors.New(publicAuthUserListError)
	}
	return connection, nil
}

const publicAuthUserGroupSingleError = "could not get authUserGroup"

func (r *queryResolver) AuthUserGroup(ctx context.Context, id string) (*fm.AuthUserGroup, error) {
	dbID := AuthUserGroupID(id)
	mods := GetAuthUserGroupPreloadMods(ctx)
	mods = append(mods, dm.AuthUserGroupWhere.ID.EQ(dbID))

	m, err := dm.AuthUserGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupSingleError)
		return nil, errors.New(publicAuthUserGroupSingleError)
	}
	return AuthUserGroupToGraphQL(m), nil
}

const publicAuthUserGroupListError = "could not list authUserGroups"

func (r *queryResolver) AuthUserGroups(ctx context.Context, first int, after *string, ordering []*fm.AuthUserGroupOrdering, filter *fm.AuthUserGroupFilter) (*fm.AuthUserGroupConnection, error) {
	mods := GetAuthUserGroupNodePreloadMods(ctx)

	mods = append(mods, AuthUserGroupFilterToMods(filter)...)
	connection, err := AuthUserGroupConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupListError)
		return nil, errors.New(publicAuthUserGroupListError)
	}
	return connection, nil
}

const publicAuthUserUserPermissionSingleError = "could not get authUserUserPermission"

func (r *queryResolver) AuthUserUserPermission(ctx context.Context, id string) (*fm.AuthUserUserPermission, error) {
	dbID := AuthUserUserPermissionID(id)
	mods := GetAuthUserUserPermissionPreloadMods(ctx)
	mods = append(mods, dm.AuthUserUserPermissionWhere.ID.EQ(dbID))

	m, err := dm.AuthUserUserPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionSingleError)
		return nil, errors.New(publicAuthUserUserPermissionSingleError)
	}
	return AuthUserUserPermissionToGraphQL(m), nil
}

const publicAuthUserUserPermissionListError = "could not list authUserUserPermissions"

func (r *queryResolver) AuthUserUserPermissions(ctx context.Context, first int, after *string, ordering []*fm.AuthUserUserPermissionOrdering, filter *fm.AuthUserUserPermissionFilter) (*fm.AuthUserUserPermissionConnection, error) {
	mods := GetAuthUserUserPermissionNodePreloadMods(ctx)

	mods = append(mods, AuthUserUserPermissionFilterToMods(filter)...)
	connection, err := AuthUserUserPermissionConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionListError)
		return nil, errors.New(publicAuthUserUserPermissionListError)
	}
	return connection, nil
}

const publicFragranceSingleError = "could not get fragrance"

func (r *queryResolver) Fragrance(ctx context.Context, id string) (*fm.Fragrance, error) {
	dbID := FragranceID(id)
	mods := GetFragrancePreloadMods(ctx)
	mods = append(mods, dm.FragranceWhere.ID.EQ(dbID))

	m, err := dm.Fragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceSingleError)
		return nil, errors.New(publicFragranceSingleError)
	}
	return FragranceToGraphQL(m), nil
}

const publicFragranceListError = "could not list fragrances"

func (r *queryResolver) Fragrances(ctx context.Context, first int, after *string, ordering []*fm.FragranceOrdering, filter *fm.FragranceFilter) (*fm.FragranceConnection, error) {
	mods := GetFragranceNodePreloadMods(ctx)

	mods = append(mods, FragranceFilterToMods(filter)...)
	connection, err := FragranceConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceListError)
		return nil, errors.New(publicFragranceListError)
	}
	return connection, nil
}

const publicFragranceInventorySingleError = "could not get fragranceInventory"

func (r *queryResolver) FragranceInventory(ctx context.Context, id string) (*fm.FragranceInventory, error) {
	dbID := FragranceInventoryID(id)
	mods := GetFragranceInventoryPreloadMods(ctx)
	mods = append(mods, dm.FragranceInventoryWhere.ID.EQ(dbID))

	m, err := dm.FragranceInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceInventorySingleError)
		return nil, errors.New(publicFragranceInventorySingleError)
	}
	return FragranceInventoryToGraphQL(m), nil
}

const publicFragranceInventoryListError = "could not list fragranceInventories"

func (r *queryResolver) FragranceInventories(ctx context.Context, first int, after *string, ordering []*fm.FragranceInventoryOrdering, filter *fm.FragranceInventoryFilter) (*fm.FragranceInventoryConnection, error) {
	mods := GetFragranceInventoryNodePreloadMods(ctx)

	mods = append(mods, FragranceInventoryFilterToMods(filter)...)
	connection, err := FragranceInventoryConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryListError)
		return nil, errors.New(publicFragranceInventoryListError)
	}
	return connection, nil
}

const publicLipidSingleError = "could not get lipid"

func (r *queryResolver) Lipid(ctx context.Context, id string) (*fm.Lipid, error) {
	dbID := LipidID(id)
	mods := GetLipidPreloadMods(ctx)
	mods = append(mods, dm.LipidWhere.ID.EQ(dbID))

	m, err := dm.Lipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidSingleError)
		return nil, errors.New(publicLipidSingleError)
	}
	return LipidToGraphQL(m), nil
}

const publicLipidListError = "could not list lipids"

func (r *queryResolver) Lipids(ctx context.Context, first int, after *string, ordering []*fm.LipidOrdering, filter *fm.LipidFilter) (*fm.LipidConnection, error) {
	mods := GetLipidNodePreloadMods(ctx)

	mods = append(mods, LipidFilterToMods(filter)...)
	connection, err := LipidConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidListError)
		return nil, errors.New(publicLipidListError)
	}
	return connection, nil
}

const publicLipidInventorySingleError = "could not get lipidInventory"

func (r *queryResolver) LipidInventory(ctx context.Context, id string) (*fm.LipidInventory, error) {
	dbID := LipidInventoryID(id)
	mods := GetLipidInventoryPreloadMods(ctx)
	mods = append(mods, dm.LipidInventoryWhere.ID.EQ(dbID))

	m, err := dm.LipidInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidInventorySingleError)
		return nil, errors.New(publicLipidInventorySingleError)
	}
	return LipidInventoryToGraphQL(m), nil
}

const publicLipidInventoryListError = "could not list lipidInventories"

func (r *queryResolver) LipidInventories(ctx context.Context, first int, after *string, ordering []*fm.LipidInventoryOrdering, filter *fm.LipidInventoryFilter) (*fm.LipidInventoryConnection, error) {
	mods := GetLipidInventoryNodePreloadMods(ctx)

	mods = append(mods, LipidInventoryFilterToMods(filter)...)
	connection, err := LipidInventoryConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryListError)
		return nil, errors.New(publicLipidInventoryListError)
	}
	return connection, nil
}

const publicLyeSingleError = "could not get lye"

func (r *queryResolver) Lye(ctx context.Context, id string) (*fm.Lye, error) {
	dbID := LyeID(id)
	mods := GetLyePreloadMods(ctx)
	mods = append(mods, dm.LyeWhere.ID.EQ(dbID))

	m, err := dm.Lyes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeSingleError)
		return nil, errors.New(publicLyeSingleError)
	}
	return LyeToGraphQL(m), nil
}

const publicLyeListError = "could not list lyes"

func (r *queryResolver) Lyes(ctx context.Context, first int, after *string, ordering []*fm.LyeOrdering, filter *fm.LyeFilter) (*fm.LyeConnection, error) {
	mods := GetLyeNodePreloadMods(ctx)

	mods = append(mods, LyeFilterToMods(filter)...)
	connection, err := LyeConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeListError)
		return nil, errors.New(publicLyeListError)
	}
	return connection, nil
}

const publicLyeInventorySingleError = "could not get lyeInventory"

func (r *queryResolver) LyeInventory(ctx context.Context, id string) (*fm.LyeInventory, error) {
	dbID := LyeInventoryID(id)
	mods := GetLyeInventoryPreloadMods(ctx)
	mods = append(mods, dm.LyeInventoryWhere.ID.EQ(dbID))

	m, err := dm.LyeInventories(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeInventorySingleError)
		return nil, errors.New(publicLyeInventorySingleError)
	}
	return LyeInventoryToGraphQL(m), nil
}

const publicLyeInventoryListError = "could not list lyeInventories"

func (r *queryResolver) LyeInventories(ctx context.Context, first int, after *string, ordering []*fm.LyeInventoryOrdering, filter *fm.LyeInventoryFilter) (*fm.LyeInventoryConnection, error) {
	mods := GetLyeInventoryNodePreloadMods(ctx)

	mods = append(mods, LyeInventoryFilterToMods(filter)...)
	connection, err := LyeInventoryConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryListError)
		return nil, errors.New(publicLyeInventoryListError)
	}
	return connection, nil
}

const publicRecipeSingleError = "could not get recipe"

func (r *queryResolver) Recipe(ctx context.Context, id string) (*fm.Recipe, error) {
	dbID := RecipeID(id)
	mods := GetRecipePreloadMods(ctx)
	mods = append(mods, dm.RecipeWhere.ID.EQ(dbID))

	m, err := dm.Recipes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeSingleError)
		return nil, errors.New(publicRecipeSingleError)
	}
	return RecipeToGraphQL(m), nil
}

const publicRecipeListError = "could not list recipes"

func (r *queryResolver) Recipes(ctx context.Context, first int, after *string, ordering []*fm.RecipeOrdering, filter *fm.RecipeFilter) (*fm.RecipeConnection, error) {
	mods := GetRecipeNodePreloadMods(ctx)

	mods = append(mods, RecipeFilterToMods(filter)...)
	connection, err := RecipeConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeListError)
		return nil, errors.New(publicRecipeListError)
	}
	return connection, nil
}

const publicRecipeAdditiveSingleError = "could not get recipeAdditive"

func (r *queryResolver) RecipeAdditive(ctx context.Context, id string) (*fm.RecipeAdditive, error) {
	dbID := RecipeAdditiveID(id)
	mods := GetRecipeAdditivePreloadMods(ctx)
	mods = append(mods, dm.RecipeAdditiveWhere.ID.EQ(dbID))

	m, err := dm.RecipeAdditives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveSingleError)
		return nil, errors.New(publicRecipeAdditiveSingleError)
	}
	return RecipeAdditiveToGraphQL(m), nil
}

const publicRecipeAdditiveListError = "could not list recipeAdditives"

func (r *queryResolver) RecipeAdditives(ctx context.Context, first int, after *string, ordering []*fm.RecipeAdditiveOrdering, filter *fm.RecipeAdditiveFilter) (*fm.RecipeAdditiveConnection, error) {
	mods := GetRecipeAdditiveNodePreloadMods(ctx)

	mods = append(mods, RecipeAdditiveFilterToMods(filter)...)
	connection, err := RecipeAdditiveConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveListError)
		return nil, errors.New(publicRecipeAdditiveListError)
	}
	return connection, nil
}

const publicRecipeBatchSingleError = "could not get recipeBatch"

func (r *queryResolver) RecipeBatch(ctx context.Context, id string) (*fm.RecipeBatch, error) {
	dbID := RecipeBatchID(id)
	mods := GetRecipeBatchPreloadMods(ctx)
	mods = append(mods, dm.RecipeBatchWhere.ID.EQ(dbID))

	m, err := dm.RecipeBatches(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchSingleError)
		return nil, errors.New(publicRecipeBatchSingleError)
	}
	return RecipeBatchToGraphQL(m), nil
}

const publicRecipeBatchListError = "could not list recipeBatches"

func (r *queryResolver) RecipeBatches(ctx context.Context, first int, after *string, ordering []*fm.RecipeBatchOrdering, filter *fm.RecipeBatchFilter) (*fm.RecipeBatchConnection, error) {
	mods := GetRecipeBatchNodePreloadMods(ctx)

	mods = append(mods, RecipeBatchFilterToMods(filter)...)
	connection, err := RecipeBatchConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchListError)
		return nil, errors.New(publicRecipeBatchListError)
	}
	return connection, nil
}

const publicRecipeBatchAdditiveSingleError = "could not get recipeBatchAdditive"

func (r *queryResolver) RecipeBatchAdditive(ctx context.Context, id string) (*fm.RecipeBatchAdditive, error) {
	dbID := RecipeBatchAdditiveID(id)
	mods := GetRecipeBatchAdditivePreloadMods(ctx)
	mods = append(mods, dm.RecipeBatchAdditiveWhere.ID.EQ(dbID))

	m, err := dm.RecipeBatchAdditives(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveSingleError)
		return nil, errors.New(publicRecipeBatchAdditiveSingleError)
	}
	return RecipeBatchAdditiveToGraphQL(m), nil
}

const publicRecipeBatchAdditiveListError = "could not list recipeBatchAdditives"

func (r *queryResolver) RecipeBatchAdditives(ctx context.Context, first int, after *string, ordering []*fm.RecipeBatchAdditiveOrdering, filter *fm.RecipeBatchAdditiveFilter) (*fm.RecipeBatchAdditiveConnection, error) {
	mods := GetRecipeBatchAdditiveNodePreloadMods(ctx)

	mods = append(mods, RecipeBatchAdditiveFilterToMods(filter)...)
	connection, err := RecipeBatchAdditiveConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveListError)
		return nil, errors.New(publicRecipeBatchAdditiveListError)
	}
	return connection, nil
}

const publicRecipeBatchFragranceSingleError = "could not get recipeBatchFragrance"

func (r *queryResolver) RecipeBatchFragrance(ctx context.Context, id string) (*fm.RecipeBatchFragrance, error) {
	dbID := RecipeBatchFragranceID(id)
	mods := GetRecipeBatchFragrancePreloadMods(ctx)
	mods = append(mods, dm.RecipeBatchFragranceWhere.ID.EQ(dbID))

	m, err := dm.RecipeBatchFragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceSingleError)
		return nil, errors.New(publicRecipeBatchFragranceSingleError)
	}
	return RecipeBatchFragranceToGraphQL(m), nil
}

const publicRecipeBatchFragranceListError = "could not list recipeBatchFragrances"

func (r *queryResolver) RecipeBatchFragrances(ctx context.Context, first int, after *string, ordering []*fm.RecipeBatchFragranceOrdering, filter *fm.RecipeBatchFragranceFilter) (*fm.RecipeBatchFragranceConnection, error) {
	mods := GetRecipeBatchFragranceNodePreloadMods(ctx)

	mods = append(mods, RecipeBatchFragranceFilterToMods(filter)...)
	connection, err := RecipeBatchFragranceConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceListError)
		return nil, errors.New(publicRecipeBatchFragranceListError)
	}
	return connection, nil
}

const publicRecipeBatchLipidSingleError = "could not get recipeBatchLipid"

func (r *queryResolver) RecipeBatchLipid(ctx context.Context, id string) (*fm.RecipeBatchLipid, error) {
	dbID := RecipeBatchLipidID(id)
	mods := GetRecipeBatchLipidPreloadMods(ctx)
	mods = append(mods, dm.RecipeBatchLipidWhere.ID.EQ(dbID))

	m, err := dm.RecipeBatchLipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidSingleError)
		return nil, errors.New(publicRecipeBatchLipidSingleError)
	}
	return RecipeBatchLipidToGraphQL(m), nil
}

const publicRecipeBatchLipidListError = "could not list recipeBatchLipids"

func (r *queryResolver) RecipeBatchLipids(ctx context.Context, first int, after *string, ordering []*fm.RecipeBatchLipidOrdering, filter *fm.RecipeBatchLipidFilter) (*fm.RecipeBatchLipidConnection, error) {
	mods := GetRecipeBatchLipidNodePreloadMods(ctx)

	mods = append(mods, RecipeBatchLipidFilterToMods(filter)...)
	connection, err := RecipeBatchLipidConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidListError)
		return nil, errors.New(publicRecipeBatchLipidListError)
	}
	return connection, nil
}

const publicRecipeBatchLyeSingleError = "could not get recipeBatchLye"

func (r *queryResolver) RecipeBatchLye(ctx context.Context, id string) (*fm.RecipeBatchLye, error) {
	dbID := RecipeBatchLyeID(id)
	mods := GetRecipeBatchLyePreloadMods(ctx)
	mods = append(mods, dm.RecipeBatchLyeWhere.ID.EQ(dbID))

	m, err := dm.RecipeBatchLyes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeSingleError)
		return nil, errors.New(publicRecipeBatchLyeSingleError)
	}
	return RecipeBatchLyeToGraphQL(m), nil
}

const publicRecipeBatchLyeListError = "could not list recipeBatchLyes"

func (r *queryResolver) RecipeBatchLyes(ctx context.Context, first int, after *string, ordering []*fm.RecipeBatchLyeOrdering, filter *fm.RecipeBatchLyeFilter) (*fm.RecipeBatchLyeConnection, error) {
	mods := GetRecipeBatchLyeNodePreloadMods(ctx)

	mods = append(mods, RecipeBatchLyeFilterToMods(filter)...)
	connection, err := RecipeBatchLyeConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeListError)
		return nil, errors.New(publicRecipeBatchLyeListError)
	}
	return connection, nil
}

const publicRecipeBatchNoteSingleError = "could not get recipeBatchNote"

func (r *queryResolver) RecipeBatchNote(ctx context.Context, id string) (*fm.RecipeBatchNote, error) {
	dbID := RecipeBatchNoteID(id)
	mods := GetRecipeBatchNotePreloadMods(ctx)
	mods = append(mods, dm.RecipeBatchNoteWhere.ID.EQ(dbID))

	m, err := dm.RecipeBatchNotes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteSingleError)
		return nil, errors.New(publicRecipeBatchNoteSingleError)
	}
	return RecipeBatchNoteToGraphQL(m), nil
}

const publicRecipeBatchNoteListError = "could not list recipeBatchNotes"

func (r *queryResolver) RecipeBatchNotes(ctx context.Context, first int, after *string, ordering []*fm.RecipeBatchNoteOrdering, filter *fm.RecipeBatchNoteFilter) (*fm.RecipeBatchNoteConnection, error) {
	mods := GetRecipeBatchNoteNodePreloadMods(ctx)

	mods = append(mods, RecipeBatchNoteFilterToMods(filter)...)
	connection, err := RecipeBatchNoteConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteListError)
		return nil, errors.New(publicRecipeBatchNoteListError)
	}
	return connection, nil
}

const publicRecipeFragranceSingleError = "could not get recipeFragrance"

func (r *queryResolver) RecipeFragrance(ctx context.Context, id string) (*fm.RecipeFragrance, error) {
	dbID := RecipeFragranceID(id)
	mods := GetRecipeFragrancePreloadMods(ctx)
	mods = append(mods, dm.RecipeFragranceWhere.ID.EQ(dbID))

	m, err := dm.RecipeFragrances(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceSingleError)
		return nil, errors.New(publicRecipeFragranceSingleError)
	}
	return RecipeFragranceToGraphQL(m), nil
}

const publicRecipeFragranceListError = "could not list recipeFragrances"

func (r *queryResolver) RecipeFragrances(ctx context.Context, first int, after *string, ordering []*fm.RecipeFragranceOrdering, filter *fm.RecipeFragranceFilter) (*fm.RecipeFragranceConnection, error) {
	mods := GetRecipeFragranceNodePreloadMods(ctx)

	mods = append(mods, RecipeFragranceFilterToMods(filter)...)
	connection, err := RecipeFragranceConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceListError)
		return nil, errors.New(publicRecipeFragranceListError)
	}
	return connection, nil
}

const publicRecipeLipidSingleError = "could not get recipeLipid"

func (r *queryResolver) RecipeLipid(ctx context.Context, id string) (*fm.RecipeLipid, error) {
	dbID := RecipeLipidID(id)
	mods := GetRecipeLipidPreloadMods(ctx)
	mods = append(mods, dm.RecipeLipidWhere.ID.EQ(dbID))

	m, err := dm.RecipeLipids(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidSingleError)
		return nil, errors.New(publicRecipeLipidSingleError)
	}
	return RecipeLipidToGraphQL(m), nil
}

const publicRecipeLipidListError = "could not list recipeLipids"

func (r *queryResolver) RecipeLipids(ctx context.Context, first int, after *string, ordering []*fm.RecipeLipidOrdering, filter *fm.RecipeLipidFilter) (*fm.RecipeLipidConnection, error) {
	mods := GetRecipeLipidNodePreloadMods(ctx)

	mods = append(mods, RecipeLipidFilterToMods(filter)...)
	connection, err := RecipeLipidConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidListError)
		return nil, errors.New(publicRecipeLipidListError)
	}
	return connection, nil
}

const publicRecipeStepSingleError = "could not get recipeStep"

func (r *queryResolver) RecipeStep(ctx context.Context, id string) (*fm.RecipeStep, error) {
	dbID := RecipeStepID(id)
	mods := GetRecipeStepPreloadMods(ctx)
	mods = append(mods, dm.RecipeStepWhere.ID.EQ(dbID))

	m, err := dm.RecipeSteps(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeStepSingleError)
		return nil, errors.New(publicRecipeStepSingleError)
	}
	return RecipeStepToGraphQL(m), nil
}

const publicRecipeStepListError = "could not list recipeSteps"

func (r *queryResolver) RecipeSteps(ctx context.Context, first int, after *string, ordering []*fm.RecipeStepOrdering, filter *fm.RecipeStepFilter) (*fm.RecipeStepConnection, error) {
	mods := GetRecipeStepNodePreloadMods(ctx)

	mods = append(mods, RecipeStepFilterToMods(filter)...)
	connection, err := RecipeStepConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeStepListError)
		return nil, errors.New(publicRecipeStepListError)
	}
	return connection, nil
}

const publicSupplierSingleError = "could not get supplier"

func (r *queryResolver) Supplier(ctx context.Context, id string) (*fm.Supplier, error) {
	dbID := SupplierID(id)
	mods := GetSupplierPreloadMods(ctx)
	mods = append(mods, dm.SupplierWhere.ID.EQ(dbID))

	m, err := dm.Suppliers(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicSupplierSingleError)
		return nil, errors.New(publicSupplierSingleError)
	}
	return SupplierToGraphQL(m), nil
}

const publicSupplierListError = "could not list suppliers"

func (r *queryResolver) Suppliers(ctx context.Context, first int, after *string, ordering []*fm.SupplierOrdering, filter *fm.SupplierFilter) (*fm.SupplierConnection, error) {
	mods := GetSupplierNodePreloadMods(ctx)

	mods = append(mods, SupplierFilterToMods(filter)...)
	connection, err := SupplierConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicSupplierListError)
		return nil, errors.New(publicSupplierListError)
	}
	return connection, nil
}

func (r *queryResolver) Node(ctx context.Context, globalGraphID string) (fm.Node, error) {
	splitID := strings.SplitN(globalGraphID, "-", 1)
	if len(splitID) != 2 {
		return nil, errors.New("could not parse id")
	}

	model := splitID[0]
	switch model {
	case "Additive":
		return r.Additive(ctx, globalGraphID)
	case "AdditiveInventory":
		return r.AdditiveInventory(ctx, globalGraphID)
	case "AuthGroup":
		return r.AuthGroup(ctx, globalGraphID)
	case "AuthGroupPermission":
		return r.AuthGroupPermission(ctx, globalGraphID)
	case "AuthPermission":
		return r.AuthPermission(ctx, globalGraphID)
	case "AuthUser":
		return r.AuthUser(ctx, globalGraphID)
	case "AuthUserGroup":
		return r.AuthUserGroup(ctx, globalGraphID)
	case "AuthUserUserPermission":
		return r.AuthUserUserPermission(ctx, globalGraphID)
	case "Fragrance":
		return r.Fragrance(ctx, globalGraphID)
	case "FragranceInventory":
		return r.FragranceInventory(ctx, globalGraphID)
	case "Lipid":
		return r.Lipid(ctx, globalGraphID)
	case "LipidInventory":
		return r.LipidInventory(ctx, globalGraphID)
	case "Lye":
		return r.Lye(ctx, globalGraphID)
	case "LyeInventory":
		return r.LyeInventory(ctx, globalGraphID)
	case "Recipe":
		return r.Recipe(ctx, globalGraphID)
	case "RecipeAdditive":
		return r.RecipeAdditive(ctx, globalGraphID)
	case "RecipeBatch":
		return r.RecipeBatch(ctx, globalGraphID)
	case "RecipeBatchAdditive":
		return r.RecipeBatchAdditive(ctx, globalGraphID)
	case "RecipeBatchFragrance":
		return r.RecipeBatchFragrance(ctx, globalGraphID)
	case "RecipeBatchLipid":
		return r.RecipeBatchLipid(ctx, globalGraphID)
	case "RecipeBatchLye":
		return r.RecipeBatchLye(ctx, globalGraphID)
	case "RecipeBatchNote":
		return r.RecipeBatchNote(ctx, globalGraphID)
	case "RecipeFragrance":
		return r.RecipeFragrance(ctx, globalGraphID)
	case "RecipeLipid":
		return r.RecipeLipid(ctx, globalGraphID)
	case "RecipeStep":
		return r.RecipeStep(ctx, globalGraphID)
	case "Supplier":
		return r.Supplier(ctx, globalGraphID)
	default:
		return nil, errors.New("could not find corresponding model for id")
	}
}

func (r *Resolver) Mutation() fm.MutationResolver	{ return &mutationResolver{r} }
func (r *Resolver) Query() fm.QueryResolver		{ return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
