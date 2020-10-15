// Generated with https://github.com/web-ridge/gqlgen-sqlboiler.
package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/graphql_models"
	. "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/helpers"
	dm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/models"
	"github.com/web-ridge/utils-go/boilergql"
)

type Resolver struct {
	db *sql.DB
}

const inputKey = "input"

const publicAdditiveCreateError = "Could not create additive"

func (r *mutationResolver) CreateAdditive(ctx context.Context, input fm.AdditiveCreateInput) (*fm.AdditivePayload, error) {

	m := AdditiveCreateInputToBoiler(&input)

	whiteList := AdditiveCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAdditiveCreateError)
		return nil, errors.New(publicAdditiveCreateError)
	}

	// resolve requested fields after creating
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

const publicAdditiveBatchCreateError = "Could not create additives"

func (r *mutationResolver) CreateAdditives(ctx context.Context, input fm.AdditivesCreateInput) (*fm.AdditivesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAdditiveUpdateError = "Could not update additive"

func (r *mutationResolver) UpdateAdditive(ctx context.Context, id string, input fm.AdditiveUpdateInput) (*fm.AdditivePayload, error) {
	m := AdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AdditiveID(id)

	if _, err := dm.Additives(
		dm.AdditiveWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAdditiveUpdateError)
		return nil, errors.New(publicAdditiveUpdateError)
	}

	// resolve requested fields after updating
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

const publicAdditiveBatchUpdateError = "Could not update additives"

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

const publicAdditiveDeleteError = "Could not delete additive"

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

const publicAdditiveBatchDeleteError = "Could not delete additives"

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

const publicAdditiveInventoryCreateError = "Could not create additiveInventory"

func (r *mutationResolver) CreateAdditiveInventory(ctx context.Context, input fm.AdditiveInventoryCreateInput) (*fm.AdditiveInventoryPayload, error) {

	m := AdditiveInventoryCreateInputToBoiler(&input)

	whiteList := AdditiveInventoryCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryCreateError)
		return nil, errors.New(publicAdditiveInventoryCreateError)
	}

	// resolve requested fields after creating
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

const publicAdditiveInventoryBatchCreateError = "Could not create additiveInventories"

func (r *mutationResolver) CreateAdditiveInventories(ctx context.Context, input fm.AdditiveInventoriesCreateInput) (*fm.AdditiveInventoriesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAdditiveInventoryUpdateError = "Could not update additiveInventory"

func (r *mutationResolver) UpdateAdditiveInventory(ctx context.Context, id string, input fm.AdditiveInventoryUpdateInput) (*fm.AdditiveInventoryPayload, error) {
	m := AdditiveInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AdditiveInventoryID(id)

	if _, err := dm.AdditiveInventories(
		dm.AdditiveInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryUpdateError)
		return nil, errors.New(publicAdditiveInventoryUpdateError)
	}

	// resolve requested fields after updating
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

const publicAdditiveInventoryBatchUpdateError = "Could not update additiveInventories"

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

const publicAdditiveInventoryDeleteError = "Could not delete additiveInventory"

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

const publicAdditiveInventoryBatchDeleteError = "Could not delete additiveInventories"

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

const publicAuthGroupCreateError = "Could not create authGroup"

func (r *mutationResolver) CreateAuthGroup(ctx context.Context, input fm.AuthGroupCreateInput) (*fm.AuthGroupPayload, error) {

	m := AuthGroupCreateInputToBoiler(&input)

	whiteList := AuthGroupCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupCreateError)
		return nil, errors.New(publicAuthGroupCreateError)
	}

	// resolve requested fields after creating
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

const publicAuthGroupBatchCreateError = "Could not create authGroups"

func (r *mutationResolver) CreateAuthGroups(ctx context.Context, input fm.AuthGroupsCreateInput) (*fm.AuthGroupsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAuthGroupUpdateError = "Could not update authGroup"

func (r *mutationResolver) UpdateAuthGroup(ctx context.Context, id string, input fm.AuthGroupUpdateInput) (*fm.AuthGroupPayload, error) {
	m := AuthGroupUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthGroupID(id)

	if _, err := dm.AuthGroups(
		dm.AuthGroupWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupUpdateError)
		return nil, errors.New(publicAuthGroupUpdateError)
	}

	// resolve requested fields after updating
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

const publicAuthGroupBatchUpdateError = "Could not update authGroups"

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

const publicAuthGroupDeleteError = "Could not delete authGroup"

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

const publicAuthGroupBatchDeleteError = "Could not delete authGroups"

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

const publicAuthGroupPermissionCreateError = "Could not create authGroupPermission"

func (r *mutationResolver) CreateAuthGroupPermission(ctx context.Context, input fm.AuthGroupPermissionCreateInput) (*fm.AuthGroupPermissionPayload, error) {

	m := AuthGroupPermissionCreateInputToBoiler(&input)

	whiteList := AuthGroupPermissionCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionCreateError)
		return nil, errors.New(publicAuthGroupPermissionCreateError)
	}

	// resolve requested fields after creating
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

const publicAuthGroupPermissionBatchCreateError = "Could not create authGroupPermissions"

func (r *mutationResolver) CreateAuthGroupPermissions(ctx context.Context, input fm.AuthGroupPermissionsCreateInput) (*fm.AuthGroupPermissionsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAuthGroupPermissionUpdateError = "Could not update authGroupPermission"

func (r *mutationResolver) UpdateAuthGroupPermission(ctx context.Context, id string, input fm.AuthGroupPermissionUpdateInput) (*fm.AuthGroupPermissionPayload, error) {
	m := AuthGroupPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthGroupPermissionID(id)

	if _, err := dm.AuthGroupPermissions(
		dm.AuthGroupPermissionWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionUpdateError)
		return nil, errors.New(publicAuthGroupPermissionUpdateError)
	}

	// resolve requested fields after updating
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

const publicAuthGroupPermissionBatchUpdateError = "Could not update authGroupPermissions"

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

const publicAuthGroupPermissionDeleteError = "Could not delete authGroupPermission"

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

const publicAuthGroupPermissionBatchDeleteError = "Could not delete authGroupPermissions"

func (r *mutationResolver) DeleteAuthGroupPermissions(ctx context.Context, filter *fm.AuthGroupPermissionFilter) (*fm.AuthGroupPermissionsDeletePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, AuthGroupPermissionFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthGroupPermissionColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthGroupPermission))
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
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthGroupPermission),
	}, nil
}

const publicAuthPermissionCreateError = "Could not create authPermission"

func (r *mutationResolver) CreateAuthPermission(ctx context.Context, input fm.AuthPermissionCreateInput) (*fm.AuthPermissionPayload, error) {

	m := AuthPermissionCreateInputToBoiler(&input)

	whiteList := AuthPermissionCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionCreateError)
		return nil, errors.New(publicAuthPermissionCreateError)
	}

	// resolve requested fields after creating
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

const publicAuthPermissionBatchCreateError = "Could not create authPermissions"

func (r *mutationResolver) CreateAuthPermissions(ctx context.Context, input fm.AuthPermissionsCreateInput) (*fm.AuthPermissionsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAuthPermissionUpdateError = "Could not update authPermission"

func (r *mutationResolver) UpdateAuthPermission(ctx context.Context, id string, input fm.AuthPermissionUpdateInput) (*fm.AuthPermissionPayload, error) {
	m := AuthPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthPermissionID(id)

	if _, err := dm.AuthPermissions(
		dm.AuthPermissionWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionUpdateError)
		return nil, errors.New(publicAuthPermissionUpdateError)
	}

	// resolve requested fields after updating
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

const publicAuthPermissionBatchUpdateError = "Could not update authPermissions"

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

const publicAuthPermissionDeleteError = "Could not delete authPermission"

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

const publicAuthPermissionBatchDeleteError = "Could not delete authPermissions"

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

const publicAuthUserCreateError = "Could not create authUser"

func (r *mutationResolver) CreateAuthUser(ctx context.Context, input fm.AuthUserCreateInput) (*fm.AuthUserPayload, error) {

	m := AuthUserCreateInputToBoiler(&input)

	whiteList := AuthUserCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAuthUserCreateError)
		return nil, errors.New(publicAuthUserCreateError)
	}

	// resolve requested fields after creating
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

const publicAuthUserBatchCreateError = "Could not create authUsers"

func (r *mutationResolver) CreateAuthUsers(ctx context.Context, input fm.AuthUsersCreateInput) (*fm.AuthUsersPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAuthUserUpdateError = "Could not update authUser"

func (r *mutationResolver) UpdateAuthUser(ctx context.Context, id string, input fm.AuthUserUpdateInput) (*fm.AuthUserPayload, error) {
	m := AuthUserUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthUserID(id)

	if _, err := dm.AuthUsers(
		dm.AuthUserWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUpdateError)
		return nil, errors.New(publicAuthUserUpdateError)
	}

	// resolve requested fields after updating
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

const publicAuthUserBatchUpdateError = "Could not update authUsers"

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

const publicAuthUserDeleteError = "Could not delete authUser"

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

const publicAuthUserBatchDeleteError = "Could not delete authUsers"

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

const publicAuthUserGroupCreateError = "Could not create authUserGroup"

func (r *mutationResolver) CreateAuthUserGroup(ctx context.Context, input fm.AuthUserGroupCreateInput) (*fm.AuthUserGroupPayload, error) {

	m := AuthUserGroupCreateInputToBoiler(&input)

	m.UserID = auth.UserIDFromContext(ctx)

	whiteList := AuthUserGroupCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
		dm.AuthUserGroupColumns.UserID,
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupCreateError)
		return nil, errors.New(publicAuthUserGroupCreateError)
	}

	// resolve requested fields after creating
	mods := GetAuthUserGroupPreloadModsWithLevel(ctx, AuthUserGroupPayloadPreloadLevels.AuthUserGroup)
	mods = append(mods, dm.AuthUserGroupWhere.ID.EQ(m.ID))
	mods = append(mods, dm.AuthUserGroupWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	pM, err := dm.AuthUserGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupCreateError)
		return nil, errors.New(publicAuthUserGroupCreateError)
	}
	return &fm.AuthUserGroupPayload{
		AuthUserGroup: AuthUserGroupToGraphQL(pM),
	}, nil
}

const publicAuthUserGroupBatchCreateError = "Could not create authUserGroups"

func (r *mutationResolver) CreateAuthUserGroups(ctx context.Context, input fm.AuthUserGroupsCreateInput) (*fm.AuthUserGroupsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAuthUserGroupUpdateError = "Could not update authUserGroup"

func (r *mutationResolver) UpdateAuthUserGroup(ctx context.Context, id string, input fm.AuthUserGroupUpdateInput) (*fm.AuthUserGroupPayload, error) {
	m := AuthUserGroupUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthUserGroupID(id)

	if _, err := dm.AuthUserGroups(
		dm.AuthUserGroupWhere.ID.EQ(dbID),
		dm.AuthUserGroupWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupUpdateError)
		return nil, errors.New(publicAuthUserGroupUpdateError)
	}

	// resolve requested fields after updating
	mods := GetAuthUserGroupPreloadModsWithLevel(ctx, AuthUserGroupPayloadPreloadLevels.AuthUserGroup)
	mods = append(mods, dm.AuthUserGroupWhere.ID.EQ(dbID))
	mods = append(mods, dm.AuthUserGroupWhere.UserID.EQ(auth.UserIDFromContext(ctx)))

	pM, err := dm.AuthUserGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupUpdateError)
		return nil, errors.New(publicAuthUserGroupUpdateError)
	}
	return &fm.AuthUserGroupPayload{
		AuthUserGroup: AuthUserGroupToGraphQL(pM),
	}, nil
}

const publicAuthUserGroupBatchUpdateError = "Could not update authUserGroups"

func (r *mutationResolver) UpdateAuthUserGroups(ctx context.Context, filter *fm.AuthUserGroupFilter, input fm.AuthUserGroupUpdateInput) (*fm.AuthUserGroupsUpdatePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, dm.AuthUserGroupWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
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

const publicAuthUserGroupDeleteError = "Could not delete authUserGroup"

func (r *mutationResolver) DeleteAuthUserGroup(ctx context.Context, id string) (*fm.AuthUserGroupDeletePayload, error) {
	dbID := AuthUserGroupID(id)

	mods := []qm.QueryMod{
		dm.AuthUserGroupWhere.ID.EQ(dbID),
		dm.AuthUserGroupWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	}
	if _, err := dm.AuthUserGroups(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupDeleteError)
		return nil, errors.New(publicAuthUserGroupDeleteError)
	}

	return &fm.AuthUserGroupDeletePayload{
		ID: id,
	}, nil
}

const publicAuthUserGroupBatchDeleteError = "Could not delete authUserGroups"

func (r *mutationResolver) DeleteAuthUserGroups(ctx context.Context, filter *fm.AuthUserGroupFilter) (*fm.AuthUserGroupsDeletePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, dm.AuthUserGroupWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, AuthUserGroupFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthUserGroupColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthUserGroup))
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
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthUserGroup),
	}, nil
}

const publicAuthUserUserPermissionCreateError = "Could not create authUserUserPermission"

func (r *mutationResolver) CreateAuthUserUserPermission(ctx context.Context, input fm.AuthUserUserPermissionCreateInput) (*fm.AuthUserUserPermissionPayload, error) {

	m := AuthUserUserPermissionCreateInputToBoiler(&input)

	m.UserID = auth.UserIDFromContext(ctx)

	whiteList := AuthUserUserPermissionCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
		dm.AuthUserUserPermissionColumns.UserID,
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionCreateError)
		return nil, errors.New(publicAuthUserUserPermissionCreateError)
	}

	// resolve requested fields after creating
	mods := GetAuthUserUserPermissionPreloadModsWithLevel(ctx, AuthUserUserPermissionPayloadPreloadLevels.AuthUserUserPermission)
	mods = append(mods, dm.AuthUserUserPermissionWhere.ID.EQ(m.ID))
	mods = append(mods, dm.AuthUserUserPermissionWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	pM, err := dm.AuthUserUserPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionCreateError)
		return nil, errors.New(publicAuthUserUserPermissionCreateError)
	}
	return &fm.AuthUserUserPermissionPayload{
		AuthUserUserPermission: AuthUserUserPermissionToGraphQL(pM),
	}, nil
}

const publicAuthUserUserPermissionBatchCreateError = "Could not create authUserUserPermissions"

func (r *mutationResolver) CreateAuthUserUserPermissions(ctx context.Context, input fm.AuthUserUserPermissionsCreateInput) (*fm.AuthUserUserPermissionsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicAuthUserUserPermissionUpdateError = "Could not update authUserUserPermission"

func (r *mutationResolver) UpdateAuthUserUserPermission(ctx context.Context, id string, input fm.AuthUserUserPermissionUpdateInput) (*fm.AuthUserUserPermissionPayload, error) {
	m := AuthUserUserPermissionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := AuthUserUserPermissionID(id)

	if _, err := dm.AuthUserUserPermissions(
		dm.AuthUserUserPermissionWhere.ID.EQ(dbID),
		dm.AuthUserUserPermissionWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionUpdateError)
		return nil, errors.New(publicAuthUserUserPermissionUpdateError)
	}

	// resolve requested fields after updating
	mods := GetAuthUserUserPermissionPreloadModsWithLevel(ctx, AuthUserUserPermissionPayloadPreloadLevels.AuthUserUserPermission)
	mods = append(mods, dm.AuthUserUserPermissionWhere.ID.EQ(dbID))
	mods = append(mods, dm.AuthUserUserPermissionWhere.UserID.EQ(auth.UserIDFromContext(ctx)))

	pM, err := dm.AuthUserUserPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionUpdateError)
		return nil, errors.New(publicAuthUserUserPermissionUpdateError)
	}
	return &fm.AuthUserUserPermissionPayload{
		AuthUserUserPermission: AuthUserUserPermissionToGraphQL(pM),
	}, nil
}

const publicAuthUserUserPermissionBatchUpdateError = "Could not update authUserUserPermissions"

func (r *mutationResolver) UpdateAuthUserUserPermissions(ctx context.Context, filter *fm.AuthUserUserPermissionFilter, input fm.AuthUserUserPermissionUpdateInput) (*fm.AuthUserUserPermissionsUpdatePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, dm.AuthUserUserPermissionWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
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

const publicAuthUserUserPermissionDeleteError = "Could not delete authUserUserPermission"

func (r *mutationResolver) DeleteAuthUserUserPermission(ctx context.Context, id string) (*fm.AuthUserUserPermissionDeletePayload, error) {
	dbID := AuthUserUserPermissionID(id)

	mods := []qm.QueryMod{
		dm.AuthUserUserPermissionWhere.ID.EQ(dbID),
		dm.AuthUserUserPermissionWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	}
	if _, err := dm.AuthUserUserPermissions(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionDeleteError)
		return nil, errors.New(publicAuthUserUserPermissionDeleteError)
	}

	return &fm.AuthUserUserPermissionDeletePayload{
		ID: id,
	}, nil
}

const publicAuthUserUserPermissionBatchDeleteError = "Could not delete authUserUserPermissions"

func (r *mutationResolver) DeleteAuthUserUserPermissions(ctx context.Context, filter *fm.AuthUserUserPermissionFilter) (*fm.AuthUserUserPermissionsDeletePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, dm.AuthUserUserPermissionWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, AuthUserUserPermissionFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.AuthUserUserPermissionColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.AuthUserUserPermission))
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
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.AuthUserUserPermission),
	}, nil
}

const publicDjangoAdminLogCreateError = "Could not create djangoAdminLog"

func (r *mutationResolver) CreateDjangoAdminLog(ctx context.Context, input fm.DjangoAdminLogCreateInput) (*fm.DjangoAdminLogPayload, error) {

	m := DjangoAdminLogCreateInputToBoiler(&input)

	m.UserID = auth.UserIDFromContext(ctx)

	whiteList := DjangoAdminLogCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
		dm.DjangoAdminLogColumns.UserID,
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogCreateError)
		return nil, errors.New(publicDjangoAdminLogCreateError)
	}

	// resolve requested fields after creating
	mods := GetDjangoAdminLogPreloadModsWithLevel(ctx, DjangoAdminLogPayloadPreloadLevels.DjangoAdminLog)
	mods = append(mods, dm.DjangoAdminLogWhere.ID.EQ(m.ID))
	mods = append(mods, dm.DjangoAdminLogWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	pM, err := dm.DjangoAdminLogs(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogCreateError)
		return nil, errors.New(publicDjangoAdminLogCreateError)
	}
	return &fm.DjangoAdminLogPayload{
		DjangoAdminLog: DjangoAdminLogToGraphQL(pM),
	}, nil
}

const publicDjangoAdminLogBatchCreateError = "Could not create djangoAdminLogs"

func (r *mutationResolver) CreateDjangoAdminLogs(ctx context.Context, input fm.DjangoAdminLogsCreateInput) (*fm.DjangoAdminLogsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicDjangoAdminLogUpdateError = "Could not update djangoAdminLog"

func (r *mutationResolver) UpdateDjangoAdminLog(ctx context.Context, id string, input fm.DjangoAdminLogUpdateInput) (*fm.DjangoAdminLogPayload, error) {
	m := DjangoAdminLogUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := DjangoAdminLogID(id)

	if _, err := dm.DjangoAdminLogs(
		dm.DjangoAdminLogWhere.ID.EQ(dbID),
		dm.DjangoAdminLogWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogUpdateError)
		return nil, errors.New(publicDjangoAdminLogUpdateError)
	}

	// resolve requested fields after updating
	mods := GetDjangoAdminLogPreloadModsWithLevel(ctx, DjangoAdminLogPayloadPreloadLevels.DjangoAdminLog)
	mods = append(mods, dm.DjangoAdminLogWhere.ID.EQ(dbID))
	mods = append(mods, dm.DjangoAdminLogWhere.UserID.EQ(auth.UserIDFromContext(ctx)))

	pM, err := dm.DjangoAdminLogs(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogUpdateError)
		return nil, errors.New(publicDjangoAdminLogUpdateError)
	}
	return &fm.DjangoAdminLogPayload{
		DjangoAdminLog: DjangoAdminLogToGraphQL(pM),
	}, nil
}

const publicDjangoAdminLogBatchUpdateError = "Could not update djangoAdminLogs"

func (r *mutationResolver) UpdateDjangoAdminLogs(ctx context.Context, filter *fm.DjangoAdminLogFilter, input fm.DjangoAdminLogUpdateInput) (*fm.DjangoAdminLogsUpdatePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, dm.DjangoAdminLogWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, DjangoAdminLogFilterToMods(filter)...)

	m := DjangoAdminLogUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.DjangoAdminLogs(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogBatchUpdateError)
		return nil, errors.New(publicDjangoAdminLogBatchUpdateError)
	}

	return &fm.DjangoAdminLogsUpdatePayload{
		Ok: true,
	}, nil
}

const publicDjangoAdminLogDeleteError = "Could not delete djangoAdminLog"

func (r *mutationResolver) DeleteDjangoAdminLog(ctx context.Context, id string) (*fm.DjangoAdminLogDeletePayload, error) {
	dbID := DjangoAdminLogID(id)

	mods := []qm.QueryMod{
		dm.DjangoAdminLogWhere.ID.EQ(dbID),
		dm.DjangoAdminLogWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	}
	if _, err := dm.DjangoAdminLogs(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogDeleteError)
		return nil, errors.New(publicDjangoAdminLogDeleteError)
	}

	return &fm.DjangoAdminLogDeletePayload{
		ID: id,
	}, nil
}

const publicDjangoAdminLogBatchDeleteError = "Could not delete djangoAdminLogs"

func (r *mutationResolver) DeleteDjangoAdminLogs(ctx context.Context, filter *fm.DjangoAdminLogFilter) (*fm.DjangoAdminLogsDeletePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, dm.DjangoAdminLogWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, DjangoAdminLogFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.DjangoAdminLogColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.DjangoAdminLog))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.DjangoAdminLogs(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogBatchDeleteError)
		return nil, errors.New(publicDjangoAdminLogBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.DjangoAdminLogs(dm.DjangoAdminLogWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogBatchDeleteError)
		return nil, errors.New(publicDjangoAdminLogBatchDeleteError)
	}
	return &fm.DjangoAdminLogsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.DjangoAdminLog),
	}, nil
}

const publicDjangoContentTypeCreateError = "Could not create djangoContentType"

func (r *mutationResolver) CreateDjangoContentType(ctx context.Context, input fm.DjangoContentTypeCreateInput) (*fm.DjangoContentTypePayload, error) {

	m := DjangoContentTypeCreateInputToBoiler(&input)

	whiteList := DjangoContentTypeCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeCreateError)
		return nil, errors.New(publicDjangoContentTypeCreateError)
	}

	// resolve requested fields after creating
	mods := GetDjangoContentTypePreloadModsWithLevel(ctx, DjangoContentTypePayloadPreloadLevels.DjangoContentType)
	mods = append(mods, dm.DjangoContentTypeWhere.ID.EQ(m.ID))
	pM, err := dm.DjangoContentTypes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeCreateError)
		return nil, errors.New(publicDjangoContentTypeCreateError)
	}
	return &fm.DjangoContentTypePayload{
		DjangoContentType: DjangoContentTypeToGraphQL(pM),
	}, nil
}

const publicDjangoContentTypeBatchCreateError = "Could not create djangoContentTypes"

func (r *mutationResolver) CreateDjangoContentTypes(ctx context.Context, input fm.DjangoContentTypesCreateInput) (*fm.DjangoContentTypesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicDjangoContentTypeUpdateError = "Could not update djangoContentType"

func (r *mutationResolver) UpdateDjangoContentType(ctx context.Context, id string, input fm.DjangoContentTypeUpdateInput) (*fm.DjangoContentTypePayload, error) {
	m := DjangoContentTypeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := DjangoContentTypeID(id)

	if _, err := dm.DjangoContentTypes(
		dm.DjangoContentTypeWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeUpdateError)
		return nil, errors.New(publicDjangoContentTypeUpdateError)
	}

	// resolve requested fields after updating
	mods := GetDjangoContentTypePreloadModsWithLevel(ctx, DjangoContentTypePayloadPreloadLevels.DjangoContentType)
	mods = append(mods, dm.DjangoContentTypeWhere.ID.EQ(dbID))

	pM, err := dm.DjangoContentTypes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeUpdateError)
		return nil, errors.New(publicDjangoContentTypeUpdateError)
	}
	return &fm.DjangoContentTypePayload{
		DjangoContentType: DjangoContentTypeToGraphQL(pM),
	}, nil
}

const publicDjangoContentTypeBatchUpdateError = "Could not update djangoContentTypes"

func (r *mutationResolver) UpdateDjangoContentTypes(ctx context.Context, filter *fm.DjangoContentTypeFilter, input fm.DjangoContentTypeUpdateInput) (*fm.DjangoContentTypesUpdatePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, DjangoContentTypeFilterToMods(filter)...)

	m := DjangoContentTypeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.DjangoContentTypes(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeBatchUpdateError)
		return nil, errors.New(publicDjangoContentTypeBatchUpdateError)
	}

	return &fm.DjangoContentTypesUpdatePayload{
		Ok: true,
	}, nil
}

const publicDjangoContentTypeDeleteError = "Could not delete djangoContentType"

func (r *mutationResolver) DeleteDjangoContentType(ctx context.Context, id string) (*fm.DjangoContentTypeDeletePayload, error) {
	dbID := DjangoContentTypeID(id)

	mods := []qm.QueryMod{
		dm.DjangoContentTypeWhere.ID.EQ(dbID),
	}
	if _, err := dm.DjangoContentTypes(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeDeleteError)
		return nil, errors.New(publicDjangoContentTypeDeleteError)
	}

	return &fm.DjangoContentTypeDeletePayload{
		ID: id,
	}, nil
}

const publicDjangoContentTypeBatchDeleteError = "Could not delete djangoContentTypes"

func (r *mutationResolver) DeleteDjangoContentTypes(ctx context.Context, filter *fm.DjangoContentTypeFilter) (*fm.DjangoContentTypesDeletePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, DjangoContentTypeFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.DjangoContentTypeColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.DjangoContentType))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.DjangoContentTypes(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeBatchDeleteError)
		return nil, errors.New(publicDjangoContentTypeBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.DjangoContentTypes(dm.DjangoContentTypeWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeBatchDeleteError)
		return nil, errors.New(publicDjangoContentTypeBatchDeleteError)
	}
	return &fm.DjangoContentTypesDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.DjangoContentType),
	}, nil
}

const publicDjangoMigrationCreateError = "Could not create djangoMigration"

func (r *mutationResolver) CreateDjangoMigration(ctx context.Context, input fm.DjangoMigrationCreateInput) (*fm.DjangoMigrationPayload, error) {

	m := DjangoMigrationCreateInputToBoiler(&input)

	whiteList := DjangoMigrationCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationCreateError)
		return nil, errors.New(publicDjangoMigrationCreateError)
	}

	// resolve requested fields after creating
	mods := GetDjangoMigrationPreloadModsWithLevel(ctx, DjangoMigrationPayloadPreloadLevels.DjangoMigration)
	mods = append(mods, dm.DjangoMigrationWhere.ID.EQ(m.ID))
	pM, err := dm.DjangoMigrations(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationCreateError)
		return nil, errors.New(publicDjangoMigrationCreateError)
	}
	return &fm.DjangoMigrationPayload{
		DjangoMigration: DjangoMigrationToGraphQL(pM),
	}, nil
}

const publicDjangoMigrationBatchCreateError = "Could not create djangoMigrations"

func (r *mutationResolver) CreateDjangoMigrations(ctx context.Context, input fm.DjangoMigrationsCreateInput) (*fm.DjangoMigrationsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicDjangoMigrationUpdateError = "Could not update djangoMigration"

func (r *mutationResolver) UpdateDjangoMigration(ctx context.Context, id string, input fm.DjangoMigrationUpdateInput) (*fm.DjangoMigrationPayload, error) {
	m := DjangoMigrationUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := DjangoMigrationID(id)

	if _, err := dm.DjangoMigrations(
		dm.DjangoMigrationWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationUpdateError)
		return nil, errors.New(publicDjangoMigrationUpdateError)
	}

	// resolve requested fields after updating
	mods := GetDjangoMigrationPreloadModsWithLevel(ctx, DjangoMigrationPayloadPreloadLevels.DjangoMigration)
	mods = append(mods, dm.DjangoMigrationWhere.ID.EQ(dbID))

	pM, err := dm.DjangoMigrations(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationUpdateError)
		return nil, errors.New(publicDjangoMigrationUpdateError)
	}
	return &fm.DjangoMigrationPayload{
		DjangoMigration: DjangoMigrationToGraphQL(pM),
	}, nil
}

const publicDjangoMigrationBatchUpdateError = "Could not update djangoMigrations"

func (r *mutationResolver) UpdateDjangoMigrations(ctx context.Context, filter *fm.DjangoMigrationFilter, input fm.DjangoMigrationUpdateInput) (*fm.DjangoMigrationsUpdatePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, DjangoMigrationFilterToMods(filter)...)

	m := DjangoMigrationUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.DjangoMigrations(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationBatchUpdateError)
		return nil, errors.New(publicDjangoMigrationBatchUpdateError)
	}

	return &fm.DjangoMigrationsUpdatePayload{
		Ok: true,
	}, nil
}

const publicDjangoMigrationDeleteError = "Could not delete djangoMigration"

func (r *mutationResolver) DeleteDjangoMigration(ctx context.Context, id string) (*fm.DjangoMigrationDeletePayload, error) {
	dbID := DjangoMigrationID(id)

	mods := []qm.QueryMod{
		dm.DjangoMigrationWhere.ID.EQ(dbID),
	}
	if _, err := dm.DjangoMigrations(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationDeleteError)
		return nil, errors.New(publicDjangoMigrationDeleteError)
	}

	return &fm.DjangoMigrationDeletePayload{
		ID: id,
	}, nil
}

const publicDjangoMigrationBatchDeleteError = "Could not delete djangoMigrations"

func (r *mutationResolver) DeleteDjangoMigrations(ctx context.Context, filter *fm.DjangoMigrationFilter) (*fm.DjangoMigrationsDeletePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, DjangoMigrationFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.DjangoMigrationColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.DjangoMigration))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.DjangoMigrations(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationBatchDeleteError)
		return nil, errors.New(publicDjangoMigrationBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerInt(IDsToRemove)
	if _, err := dm.DjangoMigrations(dm.DjangoMigrationWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationBatchDeleteError)
		return nil, errors.New(publicDjangoMigrationBatchDeleteError)
	}
	return &fm.DjangoMigrationsDeletePayload{
		Ids: boilergql.IntIDsToGraphQL(boilerIDs, dm.TableNames.DjangoMigration),
	}, nil
}

const publicDjangoSessionCreateError = "Could not create djangoSession"

func (r *mutationResolver) CreateDjangoSession(ctx context.Context, input fm.DjangoSessionCreateInput) (*fm.DjangoSessionPayload, error) {

	m := DjangoSessionCreateInputToBoiler(&input)

	whiteList := DjangoSessionCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionCreateError)
		return nil, errors.New(publicDjangoSessionCreateError)
	}

	// resolve requested fields after creating
	mods := GetDjangoSessionPreloadModsWithLevel(ctx, DjangoSessionPayloadPreloadLevels.DjangoSession)
	mods = append(mods, dm.DjangoSessionWhere.ID.EQ(m.ID))
	pM, err := dm.DjangoSessions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionCreateError)
		return nil, errors.New(publicDjangoSessionCreateError)
	}
	return &fm.DjangoSessionPayload{
		DjangoSession: DjangoSessionToGraphQL(pM),
	}, nil
}

const publicDjangoSessionBatchCreateError = "Could not create djangoSessions"

func (r *mutationResolver) CreateDjangoSessions(ctx context.Context, input fm.DjangoSessionsCreateInput) (*fm.DjangoSessionsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicDjangoSessionUpdateError = "Could not update djangoSession"

func (r *mutationResolver) UpdateDjangoSession(ctx context.Context, id string, input fm.DjangoSessionUpdateInput) (*fm.DjangoSessionPayload, error) {
	m := DjangoSessionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := DjangoSessionID(id)

	if _, err := dm.DjangoSessions(
		dm.DjangoSessionWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionUpdateError)
		return nil, errors.New(publicDjangoSessionUpdateError)
	}

	// resolve requested fields after updating
	mods := GetDjangoSessionPreloadModsWithLevel(ctx, DjangoSessionPayloadPreloadLevels.DjangoSession)
	mods = append(mods, dm.DjangoSessionWhere.ID.EQ(dbID))

	pM, err := dm.DjangoSessions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionUpdateError)
		return nil, errors.New(publicDjangoSessionUpdateError)
	}
	return &fm.DjangoSessionPayload{
		DjangoSession: DjangoSessionToGraphQL(pM),
	}, nil
}

const publicDjangoSessionBatchUpdateError = "Could not update djangoSessions"

func (r *mutationResolver) UpdateDjangoSessions(ctx context.Context, filter *fm.DjangoSessionFilter, input fm.DjangoSessionUpdateInput) (*fm.DjangoSessionsUpdatePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, DjangoSessionFilterToMods(filter)...)

	m := DjangoSessionUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.DjangoSessions(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionBatchUpdateError)
		return nil, errors.New(publicDjangoSessionBatchUpdateError)
	}

	return &fm.DjangoSessionsUpdatePayload{
		Ok: true,
	}, nil
}

const publicDjangoSessionDeleteError = "Could not delete djangoSession"

func (r *mutationResolver) DeleteDjangoSession(ctx context.Context, id string) (*fm.DjangoSessionDeletePayload, error) {
	dbID := DjangoSessionID(id)

	mods := []qm.QueryMod{
		dm.DjangoSessionWhere.ID.EQ(dbID),
	}
	if _, err := dm.DjangoSessions(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionDeleteError)
		return nil, errors.New(publicDjangoSessionDeleteError)
	}

	return &fm.DjangoSessionDeletePayload{
		ID: id,
	}, nil
}

const publicDjangoSessionBatchDeleteError = "Could not delete djangoSessions"

func (r *mutationResolver) DeleteDjangoSessions(ctx context.Context, filter *fm.DjangoSessionFilter) (*fm.DjangoSessionsDeletePayload, error) {
	var mods []qm.QueryMod
	mods = append(mods, DjangoSessionFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.DjangoSessionColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.DjangoSession))
	var IDsToRemove []boilergql.RemovedID
	if err := dm.DjangoSessions(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionBatchDeleteError)
		return nil, errors.New(publicDjangoSessionBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoiler(IDsToRemove)
	if _, err := dm.DjangoSessions(dm.DjangoSessionWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionBatchDeleteError)
		return nil, errors.New(publicDjangoSessionBatchDeleteError)
	}
	return &fm.DjangoSessionsDeletePayload{
		Ids: boilergql.IDsToGraphQL(boilerIDs, dm.TableNames.DjangoSession),
	}, nil
}

const publicFragranceCreateError = "Could not create fragrance"

func (r *mutationResolver) CreateFragrance(ctx context.Context, input fm.FragranceCreateInput) (*fm.FragrancePayload, error) {

	m := FragranceCreateInputToBoiler(&input)

	whiteList := FragranceCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicFragranceCreateError)
		return nil, errors.New(publicFragranceCreateError)
	}

	// resolve requested fields after creating
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

const publicFragranceBatchCreateError = "Could not create fragrances"

func (r *mutationResolver) CreateFragrances(ctx context.Context, input fm.FragrancesCreateInput) (*fm.FragrancesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicFragranceUpdateError = "Could not update fragrance"

func (r *mutationResolver) UpdateFragrance(ctx context.Context, id string, input fm.FragranceUpdateInput) (*fm.FragrancePayload, error) {
	m := FragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := FragranceID(id)

	if _, err := dm.Fragrances(
		dm.FragranceWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFragranceUpdateError)
		return nil, errors.New(publicFragranceUpdateError)
	}

	// resolve requested fields after updating
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

const publicFragranceBatchUpdateError = "Could not update fragrances"

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

const publicFragranceDeleteError = "Could not delete fragrance"

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

const publicFragranceBatchDeleteError = "Could not delete fragrances"

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

const publicFragranceInventoryCreateError = "Could not create fragranceInventory"

func (r *mutationResolver) CreateFragranceInventory(ctx context.Context, input fm.FragranceInventoryCreateInput) (*fm.FragranceInventoryPayload, error) {

	m := FragranceInventoryCreateInputToBoiler(&input)

	whiteList := FragranceInventoryCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryCreateError)
		return nil, errors.New(publicFragranceInventoryCreateError)
	}

	// resolve requested fields after creating
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

const publicFragranceInventoryBatchCreateError = "Could not create fragranceInventories"

func (r *mutationResolver) CreateFragranceInventories(ctx context.Context, input fm.FragranceInventoriesCreateInput) (*fm.FragranceInventoriesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicFragranceInventoryUpdateError = "Could not update fragranceInventory"

func (r *mutationResolver) UpdateFragranceInventory(ctx context.Context, id string, input fm.FragranceInventoryUpdateInput) (*fm.FragranceInventoryPayload, error) {
	m := FragranceInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := FragranceInventoryID(id)

	if _, err := dm.FragranceInventories(
		dm.FragranceInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryUpdateError)
		return nil, errors.New(publicFragranceInventoryUpdateError)
	}

	// resolve requested fields after updating
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

const publicFragranceInventoryBatchUpdateError = "Could not update fragranceInventories"

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

const publicFragranceInventoryDeleteError = "Could not delete fragranceInventory"

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

const publicFragranceInventoryBatchDeleteError = "Could not delete fragranceInventories"

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

const publicLipidCreateError = "Could not create lipid"

func (r *mutationResolver) CreateLipid(ctx context.Context, input fm.LipidCreateInput) (*fm.LipidPayload, error) {

	m := LipidCreateInputToBoiler(&input)

	whiteList := LipidCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicLipidCreateError)
		return nil, errors.New(publicLipidCreateError)
	}

	// resolve requested fields after creating
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

const publicLipidBatchCreateError = "Could not create lipids"

func (r *mutationResolver) CreateLipids(ctx context.Context, input fm.LipidsCreateInput) (*fm.LipidsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicLipidUpdateError = "Could not update lipid"

func (r *mutationResolver) UpdateLipid(ctx context.Context, id string, input fm.LipidUpdateInput) (*fm.LipidPayload, error) {
	m := LipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LipidID(id)

	if _, err := dm.Lipids(
		dm.LipidWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLipidUpdateError)
		return nil, errors.New(publicLipidUpdateError)
	}

	// resolve requested fields after updating
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

const publicLipidBatchUpdateError = "Could not update lipids"

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

const publicLipidDeleteError = "Could not delete lipid"

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

const publicLipidBatchDeleteError = "Could not delete lipids"

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

const publicLipidInventoryCreateError = "Could not create lipidInventory"

func (r *mutationResolver) CreateLipidInventory(ctx context.Context, input fm.LipidInventoryCreateInput) (*fm.LipidInventoryPayload, error) {

	m := LipidInventoryCreateInputToBoiler(&input)

	whiteList := LipidInventoryCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryCreateError)
		return nil, errors.New(publicLipidInventoryCreateError)
	}

	// resolve requested fields after creating
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

const publicLipidInventoryBatchCreateError = "Could not create lipidInventories"

func (r *mutationResolver) CreateLipidInventories(ctx context.Context, input fm.LipidInventoriesCreateInput) (*fm.LipidInventoriesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicLipidInventoryUpdateError = "Could not update lipidInventory"

func (r *mutationResolver) UpdateLipidInventory(ctx context.Context, id string, input fm.LipidInventoryUpdateInput) (*fm.LipidInventoryPayload, error) {
	m := LipidInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LipidInventoryID(id)

	if _, err := dm.LipidInventories(
		dm.LipidInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryUpdateError)
		return nil, errors.New(publicLipidInventoryUpdateError)
	}

	// resolve requested fields after updating
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

const publicLipidInventoryBatchUpdateError = "Could not update lipidInventories"

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

const publicLipidInventoryDeleteError = "Could not delete lipidInventory"

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

const publicLipidInventoryBatchDeleteError = "Could not delete lipidInventories"

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

const publicLyeCreateError = "Could not create lye"

func (r *mutationResolver) CreateLye(ctx context.Context, input fm.LyeCreateInput) (*fm.LyePayload, error) {

	m := LyeCreateInputToBoiler(&input)

	whiteList := LyeCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicLyeCreateError)
		return nil, errors.New(publicLyeCreateError)
	}

	// resolve requested fields after creating
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

const publicLyeBatchCreateError = "Could not create lyes"

func (r *mutationResolver) CreateLyes(ctx context.Context, input fm.LyesCreateInput) (*fm.LyesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicLyeUpdateError = "Could not update lye"

func (r *mutationResolver) UpdateLye(ctx context.Context, id string, input fm.LyeUpdateInput) (*fm.LyePayload, error) {
	m := LyeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LyeID(id)

	if _, err := dm.Lyes(
		dm.LyeWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLyeUpdateError)
		return nil, errors.New(publicLyeUpdateError)
	}

	// resolve requested fields after updating
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

const publicLyeBatchUpdateError = "Could not update lyes"

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

const publicLyeDeleteError = "Could not delete lye"

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

const publicLyeBatchDeleteError = "Could not delete lyes"

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

const publicLyeInventoryCreateError = "Could not create lyeInventory"

func (r *mutationResolver) CreateLyeInventory(ctx context.Context, input fm.LyeInventoryCreateInput) (*fm.LyeInventoryPayload, error) {

	m := LyeInventoryCreateInputToBoiler(&input)

	whiteList := LyeInventoryCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryCreateError)
		return nil, errors.New(publicLyeInventoryCreateError)
	}

	// resolve requested fields after creating
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

const publicLyeInventoryBatchCreateError = "Could not create lyeInventories"

func (r *mutationResolver) CreateLyeInventories(ctx context.Context, input fm.LyeInventoriesCreateInput) (*fm.LyeInventoriesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicLyeInventoryUpdateError = "Could not update lyeInventory"

func (r *mutationResolver) UpdateLyeInventory(ctx context.Context, id string, input fm.LyeInventoryUpdateInput) (*fm.LyeInventoryPayload, error) {
	m := LyeInventoryUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := LyeInventoryID(id)

	if _, err := dm.LyeInventories(
		dm.LyeInventoryWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryUpdateError)
		return nil, errors.New(publicLyeInventoryUpdateError)
	}

	// resolve requested fields after updating
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

const publicLyeInventoryBatchUpdateError = "Could not update lyeInventories"

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

const publicLyeInventoryDeleteError = "Could not delete lyeInventory"

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

const publicLyeInventoryBatchDeleteError = "Could not delete lyeInventories"

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

const publicRecipeCreateError = "Could not create recipe"

func (r *mutationResolver) CreateRecipe(ctx context.Context, input fm.RecipeCreateInput) (*fm.RecipePayload, error) {

	m := RecipeCreateInputToBoiler(&input)

	whiteList := RecipeCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeCreateError)
		return nil, errors.New(publicRecipeCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeBatchCreateError = "Could not create recipes"

func (r *mutationResolver) CreateRecipes(ctx context.Context, input fm.RecipesCreateInput) (*fm.RecipesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeUpdateError = "Could not update recipe"

func (r *mutationResolver) UpdateRecipe(ctx context.Context, id string, input fm.RecipeUpdateInput) (*fm.RecipePayload, error) {
	m := RecipeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeID(id)

	if _, err := dm.Recipes(
		dm.RecipeWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeUpdateError)
		return nil, errors.New(publicRecipeUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeBatchUpdateError = "Could not update recipes"

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

const publicRecipeDeleteError = "Could not delete recipe"

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

const publicRecipeBatchDeleteError = "Could not delete recipes"

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

const publicRecipeAdditiveCreateError = "Could not create recipeAdditive"

func (r *mutationResolver) CreateRecipeAdditive(ctx context.Context, input fm.RecipeAdditiveCreateInput) (*fm.RecipeAdditivePayload, error) {

	m := RecipeAdditiveCreateInputToBoiler(&input)

	whiteList := RecipeAdditiveCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveCreateError)
		return nil, errors.New(publicRecipeAdditiveCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeAdditiveBatchCreateError = "Could not create recipeAdditives"

func (r *mutationResolver) CreateRecipeAdditives(ctx context.Context, input fm.RecipeAdditivesCreateInput) (*fm.RecipeAdditivesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeAdditiveUpdateError = "Could not update recipeAdditive"

func (r *mutationResolver) UpdateRecipeAdditive(ctx context.Context, id string, input fm.RecipeAdditiveUpdateInput) (*fm.RecipeAdditivePayload, error) {
	m := RecipeAdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeAdditiveID(id)

	if _, err := dm.RecipeAdditives(
		dm.RecipeAdditiveWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveUpdateError)
		return nil, errors.New(publicRecipeAdditiveUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeAdditiveBatchUpdateError = "Could not update recipeAdditives"

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

const publicRecipeAdditiveDeleteError = "Could not delete recipeAdditive"

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

const publicRecipeAdditiveBatchDeleteError = "Could not delete recipeAdditives"

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

const publicRecipeBatchCreateError = "Could not create recipeBatch"

func (r *mutationResolver) CreateRecipeBatch(ctx context.Context, input fm.RecipeBatchCreateInput) (*fm.RecipeBatchPayload, error) {

	m := RecipeBatchCreateInputToBoiler(&input)

	whiteList := RecipeBatchCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchCreateError)
		return nil, errors.New(publicRecipeBatchCreateError)
	}

	// resolve requested fields after creating
	mods := GetRecipeBatchPreloadModsWithLevel(ctx, RecipeBatchPayloadPreloadLevels.RecipeBatch)
	mods = append(mods, dm.RecipeBatchWhere.ID.EQ(m.ID))
	pM, err := dm.RecipeBatches(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchCreateError)
		return nil, errors.New(publicRecipeBatchCreateError)
	}
	return &fm.RecipeBatchPayload{
		RecipeBatch: RecipeBatchToGraphQL(pM),
	}, nil
}

const publicRecipeBatchBatchCreateError = "Could not create recipeBatches"

func (r *mutationResolver) CreateRecipeBatches(ctx context.Context, input fm.RecipeBatchesCreateInput) (*fm.RecipeBatchesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeBatchUpdateError = "Could not update recipeBatch"

func (r *mutationResolver) UpdateRecipeBatch(ctx context.Context, id string, input fm.RecipeBatchUpdateInput) (*fm.RecipeBatchPayload, error) {
	m := RecipeBatchUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchID(id)

	if _, err := dm.RecipeBatches(
		dm.RecipeBatchWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchUpdateError)
		return nil, errors.New(publicRecipeBatchUpdateError)
	}

	// resolve requested fields after updating
	mods := GetRecipeBatchPreloadModsWithLevel(ctx, RecipeBatchPayloadPreloadLevels.RecipeBatch)
	mods = append(mods, dm.RecipeBatchWhere.ID.EQ(dbID))

	pM, err := dm.RecipeBatches(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchUpdateError)
		return nil, errors.New(publicRecipeBatchUpdateError)
	}
	return &fm.RecipeBatchPayload{
		RecipeBatch: RecipeBatchToGraphQL(pM),
	}, nil
}

const publicRecipeBatchBatchUpdateError = "Could not update recipeBatches"

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

const publicRecipeBatchDeleteError = "Could not delete recipeBatch"

func (r *mutationResolver) DeleteRecipeBatch(ctx context.Context, id string) (*fm.RecipeBatchDeletePayload, error) {
	dbID := RecipeBatchID(id)

	mods := []qm.QueryMod{
		dm.RecipeBatchWhere.ID.EQ(dbID),
	}
	if _, err := dm.RecipeBatches(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchDeleteError)
		return nil, errors.New(publicRecipeBatchDeleteError)
	}

	return &fm.RecipeBatchDeletePayload{
		ID: id,
	}, nil
}

const publicRecipeBatchBatchDeleteError = "Could not delete recipeBatches"

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

const publicRecipeBatchAdditiveCreateError = "Could not create recipeBatchAdditive"

func (r *mutationResolver) CreateRecipeBatchAdditive(ctx context.Context, input fm.RecipeBatchAdditiveCreateInput) (*fm.RecipeBatchAdditivePayload, error) {

	m := RecipeBatchAdditiveCreateInputToBoiler(&input)

	whiteList := RecipeBatchAdditiveCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveCreateError)
		return nil, errors.New(publicRecipeBatchAdditiveCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeBatchAdditiveBatchCreateError = "Could not create recipeBatchAdditives"

func (r *mutationResolver) CreateRecipeBatchAdditives(ctx context.Context, input fm.RecipeBatchAdditivesCreateInput) (*fm.RecipeBatchAdditivesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeBatchAdditiveUpdateError = "Could not update recipeBatchAdditive"

func (r *mutationResolver) UpdateRecipeBatchAdditive(ctx context.Context, id string, input fm.RecipeBatchAdditiveUpdateInput) (*fm.RecipeBatchAdditivePayload, error) {
	m := RecipeBatchAdditiveUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchAdditiveID(id)

	if _, err := dm.RecipeBatchAdditives(
		dm.RecipeBatchAdditiveWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveUpdateError)
		return nil, errors.New(publicRecipeBatchAdditiveUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeBatchAdditiveBatchUpdateError = "Could not update recipeBatchAdditives"

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

const publicRecipeBatchAdditiveDeleteError = "Could not delete recipeBatchAdditive"

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

const publicRecipeBatchAdditiveBatchDeleteError = "Could not delete recipeBatchAdditives"

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

const publicRecipeBatchFragranceCreateError = "Could not create recipeBatchFragrance"

func (r *mutationResolver) CreateRecipeBatchFragrance(ctx context.Context, input fm.RecipeBatchFragranceCreateInput) (*fm.RecipeBatchFragrancePayload, error) {

	m := RecipeBatchFragranceCreateInputToBoiler(&input)

	whiteList := RecipeBatchFragranceCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceCreateError)
		return nil, errors.New(publicRecipeBatchFragranceCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeBatchFragranceBatchCreateError = "Could not create recipeBatchFragrances"

func (r *mutationResolver) CreateRecipeBatchFragrances(ctx context.Context, input fm.RecipeBatchFragrancesCreateInput) (*fm.RecipeBatchFragrancesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeBatchFragranceUpdateError = "Could not update recipeBatchFragrance"

func (r *mutationResolver) UpdateRecipeBatchFragrance(ctx context.Context, id string, input fm.RecipeBatchFragranceUpdateInput) (*fm.RecipeBatchFragrancePayload, error) {
	m := RecipeBatchFragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchFragranceID(id)

	if _, err := dm.RecipeBatchFragrances(
		dm.RecipeBatchFragranceWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceUpdateError)
		return nil, errors.New(publicRecipeBatchFragranceUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeBatchFragranceBatchUpdateError = "Could not update recipeBatchFragrances"

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

const publicRecipeBatchFragranceDeleteError = "Could not delete recipeBatchFragrance"

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

const publicRecipeBatchFragranceBatchDeleteError = "Could not delete recipeBatchFragrances"

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

const publicRecipeBatchLipidCreateError = "Could not create recipeBatchLipid"

func (r *mutationResolver) CreateRecipeBatchLipid(ctx context.Context, input fm.RecipeBatchLipidCreateInput) (*fm.RecipeBatchLipidPayload, error) {

	m := RecipeBatchLipidCreateInputToBoiler(&input)

	whiteList := RecipeBatchLipidCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidCreateError)
		return nil, errors.New(publicRecipeBatchLipidCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeBatchLipidBatchCreateError = "Could not create recipeBatchLipids"

func (r *mutationResolver) CreateRecipeBatchLipids(ctx context.Context, input fm.RecipeBatchLipidsCreateInput) (*fm.RecipeBatchLipidsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeBatchLipidUpdateError = "Could not update recipeBatchLipid"

func (r *mutationResolver) UpdateRecipeBatchLipid(ctx context.Context, id string, input fm.RecipeBatchLipidUpdateInput) (*fm.RecipeBatchLipidPayload, error) {
	m := RecipeBatchLipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchLipidID(id)

	if _, err := dm.RecipeBatchLipids(
		dm.RecipeBatchLipidWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidUpdateError)
		return nil, errors.New(publicRecipeBatchLipidUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeBatchLipidBatchUpdateError = "Could not update recipeBatchLipids"

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

const publicRecipeBatchLipidDeleteError = "Could not delete recipeBatchLipid"

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

const publicRecipeBatchLipidBatchDeleteError = "Could not delete recipeBatchLipids"

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

const publicRecipeBatchLyeCreateError = "Could not create recipeBatchLye"

func (r *mutationResolver) CreateRecipeBatchLye(ctx context.Context, input fm.RecipeBatchLyeCreateInput) (*fm.RecipeBatchLyePayload, error) {

	m := RecipeBatchLyeCreateInputToBoiler(&input)

	whiteList := RecipeBatchLyeCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeCreateError)
		return nil, errors.New(publicRecipeBatchLyeCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeBatchLyeBatchCreateError = "Could not create recipeBatchLyes"

func (r *mutationResolver) CreateRecipeBatchLyes(ctx context.Context, input fm.RecipeBatchLyesCreateInput) (*fm.RecipeBatchLyesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeBatchLyeUpdateError = "Could not update recipeBatchLye"

func (r *mutationResolver) UpdateRecipeBatchLye(ctx context.Context, id string, input fm.RecipeBatchLyeUpdateInput) (*fm.RecipeBatchLyePayload, error) {
	m := RecipeBatchLyeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchLyeID(id)

	if _, err := dm.RecipeBatchLyes(
		dm.RecipeBatchLyeWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeUpdateError)
		return nil, errors.New(publicRecipeBatchLyeUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeBatchLyeBatchUpdateError = "Could not update recipeBatchLyes"

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

const publicRecipeBatchLyeDeleteError = "Could not delete recipeBatchLye"

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

const publicRecipeBatchLyeBatchDeleteError = "Could not delete recipeBatchLyes"

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

const publicRecipeBatchNoteCreateError = "Could not create recipeBatchNote"

func (r *mutationResolver) CreateRecipeBatchNote(ctx context.Context, input fm.RecipeBatchNoteCreateInput) (*fm.RecipeBatchNotePayload, error) {

	m := RecipeBatchNoteCreateInputToBoiler(&input)

	whiteList := RecipeBatchNoteCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteCreateError)
		return nil, errors.New(publicRecipeBatchNoteCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeBatchNoteBatchCreateError = "Could not create recipeBatchNotes"

func (r *mutationResolver) CreateRecipeBatchNotes(ctx context.Context, input fm.RecipeBatchNotesCreateInput) (*fm.RecipeBatchNotesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeBatchNoteUpdateError = "Could not update recipeBatchNote"

func (r *mutationResolver) UpdateRecipeBatchNote(ctx context.Context, id string, input fm.RecipeBatchNoteUpdateInput) (*fm.RecipeBatchNotePayload, error) {
	m := RecipeBatchNoteUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeBatchNoteID(id)

	if _, err := dm.RecipeBatchNotes(
		dm.RecipeBatchNoteWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteUpdateError)
		return nil, errors.New(publicRecipeBatchNoteUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeBatchNoteBatchUpdateError = "Could not update recipeBatchNotes"

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

const publicRecipeBatchNoteDeleteError = "Could not delete recipeBatchNote"

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

const publicRecipeBatchNoteBatchDeleteError = "Could not delete recipeBatchNotes"

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

const publicRecipeFragranceCreateError = "Could not create recipeFragrance"

func (r *mutationResolver) CreateRecipeFragrance(ctx context.Context, input fm.RecipeFragranceCreateInput) (*fm.RecipeFragrancePayload, error) {

	m := RecipeFragranceCreateInputToBoiler(&input)

	whiteList := RecipeFragranceCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceCreateError)
		return nil, errors.New(publicRecipeFragranceCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeFragranceBatchCreateError = "Could not create recipeFragrances"

func (r *mutationResolver) CreateRecipeFragrances(ctx context.Context, input fm.RecipeFragrancesCreateInput) (*fm.RecipeFragrancesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeFragranceUpdateError = "Could not update recipeFragrance"

func (r *mutationResolver) UpdateRecipeFragrance(ctx context.Context, id string, input fm.RecipeFragranceUpdateInput) (*fm.RecipeFragrancePayload, error) {
	m := RecipeFragranceUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeFragranceID(id)

	if _, err := dm.RecipeFragrances(
		dm.RecipeFragranceWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceUpdateError)
		return nil, errors.New(publicRecipeFragranceUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeFragranceBatchUpdateError = "Could not update recipeFragrances"

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

const publicRecipeFragranceDeleteError = "Could not delete recipeFragrance"

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

const publicRecipeFragranceBatchDeleteError = "Could not delete recipeFragrances"

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

const publicRecipeLipidCreateError = "Could not create recipeLipid"

func (r *mutationResolver) CreateRecipeLipid(ctx context.Context, input fm.RecipeLipidCreateInput) (*fm.RecipeLipidPayload, error) {

	m := RecipeLipidCreateInputToBoiler(&input)

	whiteList := RecipeLipidCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidCreateError)
		return nil, errors.New(publicRecipeLipidCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeLipidBatchCreateError = "Could not create recipeLipids"

func (r *mutationResolver) CreateRecipeLipids(ctx context.Context, input fm.RecipeLipidsCreateInput) (*fm.RecipeLipidsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeLipidUpdateError = "Could not update recipeLipid"

func (r *mutationResolver) UpdateRecipeLipid(ctx context.Context, id string, input fm.RecipeLipidUpdateInput) (*fm.RecipeLipidPayload, error) {
	m := RecipeLipidUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeLipidID(id)

	if _, err := dm.RecipeLipids(
		dm.RecipeLipidWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidUpdateError)
		return nil, errors.New(publicRecipeLipidUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeLipidBatchUpdateError = "Could not update recipeLipids"

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

const publicRecipeLipidDeleteError = "Could not delete recipeLipid"

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

const publicRecipeLipidBatchDeleteError = "Could not delete recipeLipids"

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

const publicRecipeStepCreateError = "Could not create recipeStep"

func (r *mutationResolver) CreateRecipeStep(ctx context.Context, input fm.RecipeStepCreateInput) (*fm.RecipeStepPayload, error) {

	m := RecipeStepCreateInputToBoiler(&input)

	whiteList := RecipeStepCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepCreateError)
		return nil, errors.New(publicRecipeStepCreateError)
	}

	// resolve requested fields after creating
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

const publicRecipeStepBatchCreateError = "Could not create recipeSteps"

func (r *mutationResolver) CreateRecipeSteps(ctx context.Context, input fm.RecipeStepsCreateInput) (*fm.RecipeStepsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicRecipeStepUpdateError = "Could not update recipeStep"

func (r *mutationResolver) UpdateRecipeStep(ctx context.Context, id string, input fm.RecipeStepUpdateInput) (*fm.RecipeStepPayload, error) {
	m := RecipeStepUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := RecipeStepID(id)

	if _, err := dm.RecipeSteps(
		dm.RecipeStepWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicRecipeStepUpdateError)
		return nil, errors.New(publicRecipeStepUpdateError)
	}

	// resolve requested fields after updating
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

const publicRecipeStepBatchUpdateError = "Could not update recipeSteps"

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

const publicRecipeStepDeleteError = "Could not delete recipeStep"

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

const publicRecipeStepBatchDeleteError = "Could not delete recipeSteps"

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

const publicSupplierCreateError = "Could not create supplier"

func (r *mutationResolver) CreateSupplier(ctx context.Context, input fm.SupplierCreateInput) (*fm.SupplierPayload, error) {

	m := SupplierCreateInputToBoiler(&input)

	whiteList := SupplierCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicSupplierCreateError)
		return nil, errors.New(publicSupplierCreateError)
	}

	// resolve requested fields after creating
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

const publicSupplierBatchCreateError = "Could not create suppliers"

func (r *mutationResolver) CreateSuppliers(ctx context.Context, input fm.SuppliersCreateInput) (*fm.SuppliersPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicSupplierUpdateError = "Could not update supplier"

func (r *mutationResolver) UpdateSupplier(ctx context.Context, id string, input fm.SupplierUpdateInput) (*fm.SupplierPayload, error) {
	m := SupplierUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := SupplierID(id)

	if _, err := dm.Suppliers(
		dm.SupplierWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicSupplierUpdateError)
		return nil, errors.New(publicSupplierUpdateError)
	}

	// resolve requested fields after updating
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

const publicSupplierBatchUpdateError = "Could not update suppliers"

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

const publicSupplierDeleteError = "Could not delete supplier"

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

const publicSupplierBatchDeleteError = "Could not delete suppliers"

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

const publicAdditiveSingleError = "Could not get additive"

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

const publicAdditiveListError = "Could not list additives"

func (r *queryResolver) Additives(ctx context.Context, filter *fm.AdditiveFilter) ([]*fm.Additive, error) {
	mods := GetAdditivePreloadMods(ctx)
	mods = append(mods, AdditiveFilterToMods(filter)...)
	a, err := dm.Additives(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveListError)
		return nil, errors.New(publicAdditiveListError)
	}
	return AdditivesToGraphQL(a), nil
}

const publicAdditiveInventorySingleError = "Could not get additiveInventory"

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

const publicAdditiveInventoryListError = "Could not list additiveInventories"

func (r *queryResolver) AdditiveInventories(ctx context.Context, filter *fm.AdditiveInventoryFilter) ([]*fm.AdditiveInventory, error) {
	mods := GetAdditiveInventoryPreloadMods(ctx)
	mods = append(mods, AdditiveInventoryFilterToMods(filter)...)
	a, err := dm.AdditiveInventories(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAdditiveInventoryListError)
		return nil, errors.New(publicAdditiveInventoryListError)
	}
	return AdditiveInventoriesToGraphQL(a), nil
}

const publicAuthGroupSingleError = "Could not get authGroup"

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

const publicAuthGroupListError = "Could not list authGroups"

func (r *queryResolver) AuthGroups(ctx context.Context, filter *fm.AuthGroupFilter) ([]*fm.AuthGroup, error) {
	mods := GetAuthGroupPreloadMods(ctx)
	mods = append(mods, AuthGroupFilterToMods(filter)...)
	a, err := dm.AuthGroups(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupListError)
		return nil, errors.New(publicAuthGroupListError)
	}
	return AuthGroupsToGraphQL(a), nil
}

const publicAuthGroupPermissionSingleError = "Could not get authGroupPermission"

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

const publicAuthGroupPermissionListError = "Could not list authGroupPermissions"

func (r *queryResolver) AuthGroupPermissions(ctx context.Context, filter *fm.AuthGroupPermissionFilter) ([]*fm.AuthGroupPermission, error) {
	mods := GetAuthGroupPermissionPreloadMods(ctx)
	mods = append(mods, AuthGroupPermissionFilterToMods(filter)...)
	a, err := dm.AuthGroupPermissions(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthGroupPermissionListError)
		return nil, errors.New(publicAuthGroupPermissionListError)
	}
	return AuthGroupPermissionsToGraphQL(a), nil
}

const publicAuthPermissionSingleError = "Could not get authPermission"

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

const publicAuthPermissionListError = "Could not list authPermissions"

func (r *queryResolver) AuthPermissions(ctx context.Context, filter *fm.AuthPermissionFilter) ([]*fm.AuthPermission, error) {
	mods := GetAuthPermissionPreloadMods(ctx)
	mods = append(mods, AuthPermissionFilterToMods(filter)...)
	a, err := dm.AuthPermissions(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthPermissionListError)
		return nil, errors.New(publicAuthPermissionListError)
	}
	return AuthPermissionsToGraphQL(a), nil
}

const publicAuthUserSingleError = "Could not get authUser"

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

const publicAuthUserListError = "Could not list authUsers"

func (r *queryResolver) AuthUsers(ctx context.Context, filter *fm.AuthUserFilter) ([]*fm.AuthUser, error) {
	mods := GetAuthUserPreloadMods(ctx)
	mods = append(mods, AuthUserFilterToMods(filter)...)
	a, err := dm.AuthUsers(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserListError)
		return nil, errors.New(publicAuthUserListError)
	}
	return AuthUsersToGraphQL(a), nil
}

const publicAuthUserGroupSingleError = "Could not get authUserGroup"

func (r *queryResolver) AuthUserGroup(ctx context.Context, id string) (*fm.AuthUserGroup, error) {
	dbID := AuthUserGroupID(id)

	mods := GetAuthUserGroupPreloadMods(ctx)
	mods = append(mods, dm.AuthUserGroupWhere.ID.EQ(dbID))
	mods = append(mods, dm.AuthUserGroupWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	m, err := dm.AuthUserGroups(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupSingleError)
		return nil, errors.New(publicAuthUserGroupSingleError)
	}
	return AuthUserGroupToGraphQL(m), nil
}

const publicAuthUserGroupListError = "Could not list authUserGroups"

func (r *queryResolver) AuthUserGroups(ctx context.Context, filter *fm.AuthUserGroupFilter) ([]*fm.AuthUserGroup, error) {
	mods := GetAuthUserGroupPreloadMods(ctx)
	mods = append(mods, dm.AuthUserGroupWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, AuthUserGroupFilterToMods(filter)...)
	a, err := dm.AuthUserGroups(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserGroupListError)
		return nil, errors.New(publicAuthUserGroupListError)
	}
	return AuthUserGroupsToGraphQL(a), nil
}

const publicAuthUserUserPermissionSingleError = "Could not get authUserUserPermission"

func (r *queryResolver) AuthUserUserPermission(ctx context.Context, id string) (*fm.AuthUserUserPermission, error) {
	dbID := AuthUserUserPermissionID(id)

	mods := GetAuthUserUserPermissionPreloadMods(ctx)
	mods = append(mods, dm.AuthUserUserPermissionWhere.ID.EQ(dbID))
	mods = append(mods, dm.AuthUserUserPermissionWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	m, err := dm.AuthUserUserPermissions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionSingleError)
		return nil, errors.New(publicAuthUserUserPermissionSingleError)
	}
	return AuthUserUserPermissionToGraphQL(m), nil
}

const publicAuthUserUserPermissionListError = "Could not list authUserUserPermissions"

func (r *queryResolver) AuthUserUserPermissions(ctx context.Context, filter *fm.AuthUserUserPermissionFilter) ([]*fm.AuthUserUserPermission, error) {
	mods := GetAuthUserUserPermissionPreloadMods(ctx)
	mods = append(mods, dm.AuthUserUserPermissionWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, AuthUserUserPermissionFilterToMods(filter)...)
	a, err := dm.AuthUserUserPermissions(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicAuthUserUserPermissionListError)
		return nil, errors.New(publicAuthUserUserPermissionListError)
	}
	return AuthUserUserPermissionsToGraphQL(a), nil
}

const publicDjangoAdminLogSingleError = "Could not get djangoAdminLog"

func (r *queryResolver) DjangoAdminLog(ctx context.Context, id string) (*fm.DjangoAdminLog, error) {
	dbID := DjangoAdminLogID(id)

	mods := GetDjangoAdminLogPreloadMods(ctx)
	mods = append(mods, dm.DjangoAdminLogWhere.ID.EQ(dbID))
	mods = append(mods, dm.DjangoAdminLogWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	m, err := dm.DjangoAdminLogs(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogSingleError)
		return nil, errors.New(publicDjangoAdminLogSingleError)
	}
	return DjangoAdminLogToGraphQL(m), nil
}

const publicDjangoAdminLogListError = "Could not list djangoAdminLogs"

func (r *queryResolver) DjangoAdminLogs(ctx context.Context, filter *fm.DjangoAdminLogFilter) ([]*fm.DjangoAdminLog, error) {
	mods := GetDjangoAdminLogPreloadMods(ctx)
	mods = append(mods, dm.DjangoAdminLogWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, DjangoAdminLogFilterToMods(filter)...)
	a, err := dm.DjangoAdminLogs(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoAdminLogListError)
		return nil, errors.New(publicDjangoAdminLogListError)
	}
	return DjangoAdminLogsToGraphQL(a), nil
}

const publicDjangoContentTypeSingleError = "Could not get djangoContentType"

func (r *queryResolver) DjangoContentType(ctx context.Context, id string) (*fm.DjangoContentType, error) {
	dbID := DjangoContentTypeID(id)

	mods := GetDjangoContentTypePreloadMods(ctx)
	mods = append(mods, dm.DjangoContentTypeWhere.ID.EQ(dbID))
	m, err := dm.DjangoContentTypes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeSingleError)
		return nil, errors.New(publicDjangoContentTypeSingleError)
	}
	return DjangoContentTypeToGraphQL(m), nil
}

const publicDjangoContentTypeListError = "Could not list djangoContentTypes"

func (r *queryResolver) DjangoContentTypes(ctx context.Context, filter *fm.DjangoContentTypeFilter) ([]*fm.DjangoContentType, error) {
	mods := GetDjangoContentTypePreloadMods(ctx)
	mods = append(mods, DjangoContentTypeFilterToMods(filter)...)
	a, err := dm.DjangoContentTypes(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoContentTypeListError)
		return nil, errors.New(publicDjangoContentTypeListError)
	}
	return DjangoContentTypesToGraphQL(a), nil
}

const publicDjangoMigrationSingleError = "Could not get djangoMigration"

func (r *queryResolver) DjangoMigration(ctx context.Context, id string) (*fm.DjangoMigration, error) {
	dbID := DjangoMigrationID(id)

	mods := GetDjangoMigrationPreloadMods(ctx)
	mods = append(mods, dm.DjangoMigrationWhere.ID.EQ(dbID))
	m, err := dm.DjangoMigrations(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationSingleError)
		return nil, errors.New(publicDjangoMigrationSingleError)
	}
	return DjangoMigrationToGraphQL(m), nil
}

const publicDjangoMigrationListError = "Could not list djangoMigrations"

func (r *queryResolver) DjangoMigrations(ctx context.Context, filter *fm.DjangoMigrationFilter) ([]*fm.DjangoMigration, error) {
	mods := GetDjangoMigrationPreloadMods(ctx)
	mods = append(mods, DjangoMigrationFilterToMods(filter)...)
	a, err := dm.DjangoMigrations(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoMigrationListError)
		return nil, errors.New(publicDjangoMigrationListError)
	}
	return DjangoMigrationsToGraphQL(a), nil
}

const publicDjangoSessionSingleError = "Could not get djangoSession"

func (r *queryResolver) DjangoSession(ctx context.Context, id string) (*fm.DjangoSession, error) {
	dbID := DjangoSessionID(id)

	mods := GetDjangoSessionPreloadMods(ctx)
	mods = append(mods, dm.DjangoSessionWhere.ID.EQ(dbID))
	m, err := dm.DjangoSessions(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionSingleError)
		return nil, errors.New(publicDjangoSessionSingleError)
	}
	return DjangoSessionToGraphQL(m), nil
}

const publicDjangoSessionListError = "Could not list djangoSessions"

func (r *queryResolver) DjangoSessions(ctx context.Context, filter *fm.DjangoSessionFilter) ([]*fm.DjangoSession, error) {
	mods := GetDjangoSessionPreloadMods(ctx)
	mods = append(mods, DjangoSessionFilterToMods(filter)...)
	a, err := dm.DjangoSessions(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicDjangoSessionListError)
		return nil, errors.New(publicDjangoSessionListError)
	}
	return DjangoSessionsToGraphQL(a), nil
}

const publicFragranceSingleError = "Could not get fragrance"

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

const publicFragranceListError = "Could not list fragrances"

func (r *queryResolver) Fragrances(ctx context.Context, filter *fm.FragranceFilter) ([]*fm.Fragrance, error) {
	mods := GetFragrancePreloadMods(ctx)
	mods = append(mods, FragranceFilterToMods(filter)...)
	a, err := dm.Fragrances(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceListError)
		return nil, errors.New(publicFragranceListError)
	}
	return FragrancesToGraphQL(a), nil
}

const publicFragranceInventorySingleError = "Could not get fragranceInventory"

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

const publicFragranceInventoryListError = "Could not list fragranceInventories"

func (r *queryResolver) FragranceInventories(ctx context.Context, filter *fm.FragranceInventoryFilter) ([]*fm.FragranceInventory, error) {
	mods := GetFragranceInventoryPreloadMods(ctx)
	mods = append(mods, FragranceInventoryFilterToMods(filter)...)
	a, err := dm.FragranceInventories(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFragranceInventoryListError)
		return nil, errors.New(publicFragranceInventoryListError)
	}
	return FragranceInventoriesToGraphQL(a), nil
}

const publicLipidSingleError = "Could not get lipid"

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

const publicLipidListError = "Could not list lipids"

func (r *queryResolver) Lipids(ctx context.Context, filter *fm.LipidFilter) ([]*fm.Lipid, error) {
	mods := GetLipidPreloadMods(ctx)
	mods = append(mods, LipidFilterToMods(filter)...)
	a, err := dm.Lipids(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidListError)
		return nil, errors.New(publicLipidListError)
	}
	return LipidsToGraphQL(a), nil
}

const publicLipidInventorySingleError = "Could not get lipidInventory"

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

const publicLipidInventoryListError = "Could not list lipidInventories"

func (r *queryResolver) LipidInventories(ctx context.Context, filter *fm.LipidInventoryFilter) ([]*fm.LipidInventory, error) {
	mods := GetLipidInventoryPreloadMods(ctx)
	mods = append(mods, LipidInventoryFilterToMods(filter)...)
	a, err := dm.LipidInventories(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLipidInventoryListError)
		return nil, errors.New(publicLipidInventoryListError)
	}
	return LipidInventoriesToGraphQL(a), nil
}

const publicLyeSingleError = "Could not get lye"

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

const publicLyeListError = "Could not list lyes"

func (r *queryResolver) Lyes(ctx context.Context, filter *fm.LyeFilter) ([]*fm.Lye, error) {
	mods := GetLyePreloadMods(ctx)
	mods = append(mods, LyeFilterToMods(filter)...)
	a, err := dm.Lyes(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeListError)
		return nil, errors.New(publicLyeListError)
	}
	return LyesToGraphQL(a), nil
}

const publicLyeInventorySingleError = "Could not get lyeInventory"

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

const publicLyeInventoryListError = "Could not list lyeInventories"

func (r *queryResolver) LyeInventories(ctx context.Context, filter *fm.LyeInventoryFilter) ([]*fm.LyeInventory, error) {
	mods := GetLyeInventoryPreloadMods(ctx)
	mods = append(mods, LyeInventoryFilterToMods(filter)...)
	a, err := dm.LyeInventories(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLyeInventoryListError)
		return nil, errors.New(publicLyeInventoryListError)
	}
	return LyeInventoriesToGraphQL(a), nil
}

const publicRecipeSingleError = "Could not get recipe"

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

const publicRecipeListError = "Could not list recipes"

func (r *queryResolver) Recipes(ctx context.Context, filter *fm.RecipeFilter) ([]*fm.Recipe, error) {
	mods := GetRecipePreloadMods(ctx)
	mods = append(mods, RecipeFilterToMods(filter)...)
	a, err := dm.Recipes(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeListError)
		return nil, errors.New(publicRecipeListError)
	}
	return RecipesToGraphQL(a), nil
}

const publicRecipeAdditiveSingleError = "Could not get recipeAdditive"

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

const publicRecipeAdditiveListError = "Could not list recipeAdditives"

func (r *queryResolver) RecipeAdditives(ctx context.Context, filter *fm.RecipeAdditiveFilter) ([]*fm.RecipeAdditive, error) {
	mods := GetRecipeAdditivePreloadMods(ctx)
	mods = append(mods, RecipeAdditiveFilterToMods(filter)...)
	a, err := dm.RecipeAdditives(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeAdditiveListError)
		return nil, errors.New(publicRecipeAdditiveListError)
	}
	return RecipeAdditivesToGraphQL(a), nil
}

const publicRecipeBatchSingleError = "Could not get recipeBatch"

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

const publicRecipeBatchListError = "Could not list recipeBatches"

func (r *queryResolver) RecipeBatches(ctx context.Context, filter *fm.RecipeBatchFilter) ([]*fm.RecipeBatch, error) {
	mods := GetRecipeBatchPreloadMods(ctx)
	mods = append(mods, RecipeBatchFilterToMods(filter)...)
	a, err := dm.RecipeBatches(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchListError)
		return nil, errors.New(publicRecipeBatchListError)
	}
	return RecipeBatchesToGraphQL(a), nil
}

const publicRecipeBatchAdditiveSingleError = "Could not get recipeBatchAdditive"

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

const publicRecipeBatchAdditiveListError = "Could not list recipeBatchAdditives"

func (r *queryResolver) RecipeBatchAdditives(ctx context.Context, filter *fm.RecipeBatchAdditiveFilter) ([]*fm.RecipeBatchAdditive, error) {
	mods := GetRecipeBatchAdditivePreloadMods(ctx)
	mods = append(mods, RecipeBatchAdditiveFilterToMods(filter)...)
	a, err := dm.RecipeBatchAdditives(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchAdditiveListError)
		return nil, errors.New(publicRecipeBatchAdditiveListError)
	}
	return RecipeBatchAdditivesToGraphQL(a), nil
}

const publicRecipeBatchFragranceSingleError = "Could not get recipeBatchFragrance"

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

const publicRecipeBatchFragranceListError = "Could not list recipeBatchFragrances"

func (r *queryResolver) RecipeBatchFragrances(ctx context.Context, filter *fm.RecipeBatchFragranceFilter) ([]*fm.RecipeBatchFragrance, error) {
	mods := GetRecipeBatchFragrancePreloadMods(ctx)
	mods = append(mods, RecipeBatchFragranceFilterToMods(filter)...)
	a, err := dm.RecipeBatchFragrances(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchFragranceListError)
		return nil, errors.New(publicRecipeBatchFragranceListError)
	}
	return RecipeBatchFragrancesToGraphQL(a), nil
}

const publicRecipeBatchLipidSingleError = "Could not get recipeBatchLipid"

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

const publicRecipeBatchLipidListError = "Could not list recipeBatchLipids"

func (r *queryResolver) RecipeBatchLipids(ctx context.Context, filter *fm.RecipeBatchLipidFilter) ([]*fm.RecipeBatchLipid, error) {
	mods := GetRecipeBatchLipidPreloadMods(ctx)
	mods = append(mods, RecipeBatchLipidFilterToMods(filter)...)
	a, err := dm.RecipeBatchLipids(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLipidListError)
		return nil, errors.New(publicRecipeBatchLipidListError)
	}
	return RecipeBatchLipidsToGraphQL(a), nil
}

const publicRecipeBatchLyeSingleError = "Could not get recipeBatchLye"

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

const publicRecipeBatchLyeListError = "Could not list recipeBatchLyes"

func (r *queryResolver) RecipeBatchLyes(ctx context.Context, filter *fm.RecipeBatchLyeFilter) ([]*fm.RecipeBatchLye, error) {
	mods := GetRecipeBatchLyePreloadMods(ctx)
	mods = append(mods, RecipeBatchLyeFilterToMods(filter)...)
	a, err := dm.RecipeBatchLyes(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchLyeListError)
		return nil, errors.New(publicRecipeBatchLyeListError)
	}
	return RecipeBatchLyesToGraphQL(a), nil
}

const publicRecipeBatchNoteSingleError = "Could not get recipeBatchNote"

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

const publicRecipeBatchNoteListError = "Could not list recipeBatchNotes"

func (r *queryResolver) RecipeBatchNotes(ctx context.Context, filter *fm.RecipeBatchNoteFilter) ([]*fm.RecipeBatchNote, error) {
	mods := GetRecipeBatchNotePreloadMods(ctx)
	mods = append(mods, RecipeBatchNoteFilterToMods(filter)...)
	a, err := dm.RecipeBatchNotes(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeBatchNoteListError)
		return nil, errors.New(publicRecipeBatchNoteListError)
	}
	return RecipeBatchNotesToGraphQL(a), nil
}

const publicRecipeFragranceSingleError = "Could not get recipeFragrance"

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

const publicRecipeFragranceListError = "Could not list recipeFragrances"

func (r *queryResolver) RecipeFragrances(ctx context.Context, filter *fm.RecipeFragranceFilter) ([]*fm.RecipeFragrance, error) {
	mods := GetRecipeFragrancePreloadMods(ctx)
	mods = append(mods, RecipeFragranceFilterToMods(filter)...)
	a, err := dm.RecipeFragrances(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeFragranceListError)
		return nil, errors.New(publicRecipeFragranceListError)
	}
	return RecipeFragrancesToGraphQL(a), nil
}

const publicRecipeLipidSingleError = "Could not get recipeLipid"

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

const publicRecipeLipidListError = "Could not list recipeLipids"

func (r *queryResolver) RecipeLipids(ctx context.Context, filter *fm.RecipeLipidFilter) ([]*fm.RecipeLipid, error) {
	mods := GetRecipeLipidPreloadMods(ctx)
	mods = append(mods, RecipeLipidFilterToMods(filter)...)
	a, err := dm.RecipeLipids(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeLipidListError)
		return nil, errors.New(publicRecipeLipidListError)
	}
	return RecipeLipidsToGraphQL(a), nil
}

const publicRecipeStepSingleError = "Could not get recipeStep"

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

const publicRecipeStepListError = "Could not list recipeSteps"

func (r *queryResolver) RecipeSteps(ctx context.Context, filter *fm.RecipeStepFilter) ([]*fm.RecipeStep, error) {
	mods := GetRecipeStepPreloadMods(ctx)
	mods = append(mods, RecipeStepFilterToMods(filter)...)
	a, err := dm.RecipeSteps(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicRecipeStepListError)
		return nil, errors.New(publicRecipeStepListError)
	}
	return RecipeStepsToGraphQL(a), nil
}

const publicSupplierSingleError = "Could not get supplier"

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

const publicSupplierListError = "Could not list suppliers"

func (r *queryResolver) Suppliers(ctx context.Context, filter *fm.SupplierFilter) ([]*fm.Supplier, error) {
	mods := GetSupplierPreloadMods(ctx)
	mods = append(mods, SupplierFilterToMods(filter)...)
	a, err := dm.Suppliers(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicSupplierListError)
		return nil, errors.New(publicSupplierListError)
	}
	return SuppliersToGraphQL(a), nil
}

func (r *Resolver) Mutation() fm.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() fm.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
