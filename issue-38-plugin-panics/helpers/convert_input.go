package helpers

import (
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func AdditiveCreateInputsToBoiler(am []*graphql_models.AdditiveCreateInput) []*models.Additive {
	ar := make([]*models.Additive, len(am))
	for i, m := range am {
		ar[i] = AdditiveCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AdditiveCreateInputToBoiler(
	m *graphql_models.AdditiveCreateInput,
) *models.Additive {
	if m == nil {
		return nil
	}

	r := &models.Additive{
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func AdditiveCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AdditiveCreateInput,
) models.M {
	model := AdditiveCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.AdditiveColumns.Name] = model.Name
		case "note":
			modelM[models.AdditiveColumns.Note] = model.Note
		case "createdAt":
			modelM[models.AdditiveColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.AdditiveColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.AdditiveColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func AdditiveCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AdditiveInventoryCreateInputsToBoiler(am []*graphql_models.AdditiveInventoryCreateInput) []*models.AdditiveInventory {
	ar := make([]*models.AdditiveInventory, len(am))
	for i, m := range am {
		ar[i] = AdditiveInventoryCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AdditiveInventoryCreateInputToBoiler(
	m *graphql_models.AdditiveInventoryCreateInput,
) *models.AdditiveInventory {
	if m == nil {
		return nil
	}

	r := &models.AdditiveInventory{
		PurchaseDate:	boilergql.IntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.IntToTimeDotTime(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,
		AdditiveID:	int(boilergql.IDToBoiler(m.AdditiveID)),
		SupplierID:	int(boilergql.IDToBoiler(m.SupplierID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func AdditiveInventoryCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AdditiveInventoryCreateInput,
) models.M {
	model := AdditiveInventoryCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.AdditiveInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.AdditiveInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.AdditiveInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.AdditiveInventoryColumns.Weight] = model.Weight
		case "additiveId":
			modelM[models.AdditiveInventoryColumns.AdditiveID] = model.AdditiveID
		case "supplierId":
			modelM[models.AdditiveInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.AdditiveInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.AdditiveInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.AdditiveInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func AdditiveInventoryCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.Weight)
		case "additiveId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.AdditiveID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AdditiveInventoryUpdateInputsToBoiler(am []*graphql_models.AdditiveInventoryUpdateInput) []*models.AdditiveInventory {
	ar := make([]*models.AdditiveInventory, len(am))
	for i, m := range am {
		ar[i] = AdditiveInventoryUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AdditiveInventoryUpdateInputToBoiler(
	m *graphql_models.AdditiveInventoryUpdateInput,
) *models.AdditiveInventory {
	if m == nil {
		return nil
	}

	r := &models.AdditiveInventory{
		PurchaseDate:	boilergql.PointerIntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.PointerIntToTimeDotTime(m.ExpiryDate),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		AdditiveID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.AdditiveID))),
		SupplierID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.SupplierID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func AdditiveInventoryUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AdditiveInventoryUpdateInput,
) models.M {
	model := AdditiveInventoryUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.AdditiveInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.AdditiveInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.AdditiveInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.AdditiveInventoryColumns.Weight] = model.Weight
		case "additiveId":
			modelM[models.AdditiveInventoryColumns.AdditiveID] = model.AdditiveID
		case "supplierId":
			modelM[models.AdditiveInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.AdditiveInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.AdditiveInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.AdditiveInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func AdditiveInventoryUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.Weight)
		case "additiveId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.AdditiveID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AdditiveUpdateInputsToBoiler(am []*graphql_models.AdditiveUpdateInput) []*models.Additive {
	ar := make([]*models.Additive, len(am))
	for i, m := range am {
		ar[i] = AdditiveUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AdditiveUpdateInputToBoiler(
	m *graphql_models.AdditiveUpdateInput,
) *models.Additive {
	if m == nil {
		return nil
	}

	r := &models.Additive{
		Name:		boilergql.PointerStringToString(m.Name),
		Note:		boilergql.PointerStringToString(m.Note),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func AdditiveUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AdditiveUpdateInput,
) models.M {
	model := AdditiveUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.AdditiveColumns.Name] = model.Name
		case "note":
			modelM[models.AdditiveColumns.Note] = model.Note
		case "createdAt":
			modelM[models.AdditiveColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.AdditiveColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.AdditiveColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func AdditiveUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AdditiveColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthGroupCreateInputsToBoiler(am []*graphql_models.AuthGroupCreateInput) []*models.AuthGroup {
	ar := make([]*models.AuthGroup, len(am))
	for i, m := range am {
		ar[i] = AuthGroupCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthGroupCreateInputToBoiler(
	m *graphql_models.AuthGroupCreateInput,
) *models.AuthGroup {
	if m == nil {
		return nil
	}

	r := &models.AuthGroup{
		Name: m.Name,
	}
	return r
}

func AuthGroupCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthGroupCreateInput,
) models.M {
	model := AuthGroupCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.AuthGroupColumns.Name] = model.Name
		}
	}
	return modelM
}

func AuthGroupCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthGroupColumns.Name)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthGroupPermissionCreateInputsToBoiler(am []*graphql_models.AuthGroupPermissionCreateInput) []*models.AuthGroupPermission {
	ar := make([]*models.AuthGroupPermission, len(am))
	for i, m := range am {
		ar[i] = AuthGroupPermissionCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthGroupPermissionCreateInputToBoiler(
	m *graphql_models.AuthGroupPermissionCreateInput,
) *models.AuthGroupPermission {
	if m == nil {
		return nil
	}

	r := &models.AuthGroupPermission{
		GroupID:	int(boilergql.IDToBoiler(m.GroupID)),
		PermissionID:	int(boilergql.IDToBoiler(m.PermissionID)),
	}
	return r
}

func AuthGroupPermissionCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthGroupPermissionCreateInput,
) models.M {
	model := AuthGroupPermissionCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "groupId":
			modelM[models.AuthGroupPermissionColumns.GroupID] = model.GroupID
		case "permissionId":
			modelM[models.AuthGroupPermissionColumns.PermissionID] = model.PermissionID
		}
	}
	return modelM
}

func AuthGroupPermissionCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "groupId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthGroupPermissionColumns.GroupID)
		case "permissionId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthGroupPermissionColumns.PermissionID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthGroupPermissionUpdateInputsToBoiler(am []*graphql_models.AuthGroupPermissionUpdateInput) []*models.AuthGroupPermission {
	ar := make([]*models.AuthGroupPermission, len(am))
	for i, m := range am {
		ar[i] = AuthGroupPermissionUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthGroupPermissionUpdateInputToBoiler(
	m *graphql_models.AuthGroupPermissionUpdateInput,
) *models.AuthGroupPermission {
	if m == nil {
		return nil
	}

	r := &models.AuthGroupPermission{
		GroupID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.GroupID))),
		PermissionID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.PermissionID))),
	}
	return r
}

func AuthGroupPermissionUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthGroupPermissionUpdateInput,
) models.M {
	model := AuthGroupPermissionUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "groupId":
			modelM[models.AuthGroupPermissionColumns.GroupID] = model.GroupID
		case "permissionId":
			modelM[models.AuthGroupPermissionColumns.PermissionID] = model.PermissionID
		}
	}
	return modelM
}

func AuthGroupPermissionUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "groupId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthGroupPermissionColumns.GroupID)
		case "permissionId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthGroupPermissionColumns.PermissionID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthGroupUpdateInputsToBoiler(am []*graphql_models.AuthGroupUpdateInput) []*models.AuthGroup {
	ar := make([]*models.AuthGroup, len(am))
	for i, m := range am {
		ar[i] = AuthGroupUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthGroupUpdateInputToBoiler(
	m *graphql_models.AuthGroupUpdateInput,
) *models.AuthGroup {
	if m == nil {
		return nil
	}

	r := &models.AuthGroup{
		Name: boilergql.PointerStringToString(m.Name),
	}
	return r
}

func AuthGroupUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthGroupUpdateInput,
) models.M {
	model := AuthGroupUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.AuthGroupColumns.Name] = model.Name
		}
	}
	return modelM
}

func AuthGroupUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthGroupColumns.Name)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthPermissionCreateInputsToBoiler(am []*graphql_models.AuthPermissionCreateInput) []*models.AuthPermission {
	ar := make([]*models.AuthPermission, len(am))
	for i, m := range am {
		ar[i] = AuthPermissionCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthPermissionCreateInputToBoiler(
	m *graphql_models.AuthPermissionCreateInput,
) *models.AuthPermission {
	if m == nil {
		return nil
	}

	r := &models.AuthPermission{
		Name:		m.Name,
		ContentTypeID:	boilergql.StringToInt(m.ContentTypeID),
		Codename:	m.Codename,
	}
	return r
}

func AuthPermissionCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthPermissionCreateInput,
) models.M {
	model := AuthPermissionCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.AuthPermissionColumns.Name] = model.Name
		case "contentTypeId":
			modelM[models.AuthPermissionColumns.ContentTypeID] = model.ContentTypeID
		case "codename":
			modelM[models.AuthPermissionColumns.Codename] = model.Codename
		}
	}
	return modelM
}

func AuthPermissionCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthPermissionColumns.Name)
		case "contentTypeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthPermissionColumns.ContentTypeID)
		case "codename":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthPermissionColumns.Codename)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthPermissionUpdateInputsToBoiler(am []*graphql_models.AuthPermissionUpdateInput) []*models.AuthPermission {
	ar := make([]*models.AuthPermission, len(am))
	for i, m := range am {
		ar[i] = AuthPermissionUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthPermissionUpdateInputToBoiler(
	m *graphql_models.AuthPermissionUpdateInput,
) *models.AuthPermission {
	if m == nil {
		return nil
	}

	r := &models.AuthPermission{
		Name:		boilergql.PointerStringToString(m.Name),
		ContentTypeID:	boilergql.PointerStringToInt(m.ContentTypeID),
		Codename:	boilergql.PointerStringToString(m.Codename),
	}
	return r
}

func AuthPermissionUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthPermissionUpdateInput,
) models.M {
	model := AuthPermissionUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.AuthPermissionColumns.Name] = model.Name
		case "contentTypeId":
			modelM[models.AuthPermissionColumns.ContentTypeID] = model.ContentTypeID
		case "codename":
			modelM[models.AuthPermissionColumns.Codename] = model.Codename
		}
	}
	return modelM
}

func AuthPermissionUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthPermissionColumns.Name)
		case "contentTypeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthPermissionColumns.ContentTypeID)
		case "codename":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthPermissionColumns.Codename)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthUserCreateInputsToBoiler(am []*graphql_models.AuthUserCreateInput) []*models.AuthUser {
	ar := make([]*models.AuthUser, len(am))
	for i, m := range am {
		ar[i] = AuthUserCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthUserCreateInputToBoiler(
	m *graphql_models.AuthUserCreateInput,
) *models.AuthUser {
	if m == nil {
		return nil
	}

	r := &models.AuthUser{
		Password:	m.Password,
		LastLogin:	boilergql.PointerIntToNullDotTime(m.LastLogin),
		IsSuperuser:	m.IsSuperuser,
		Username:	m.Username,
		FirstName:	m.FirstName,
		LastName:	m.LastName,
		Email:		m.Email,
		IsStaff:	m.IsStaff,
		IsActive:	m.IsActive,
		DateJoined:	boilergql.IntToTimeDotTime(m.DateJoined),
	}
	return r
}

func AuthUserCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthUserCreateInput,
) models.M {
	model := AuthUserCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "password":
			modelM[models.AuthUserColumns.Password] = model.Password
		case "lastLogin":
			modelM[models.AuthUserColumns.LastLogin] = model.LastLogin
		case "isSuperuser":
			modelM[models.AuthUserColumns.IsSuperuser] = model.IsSuperuser
		case "username":
			modelM[models.AuthUserColumns.Username] = model.Username
		case "firstName":
			modelM[models.AuthUserColumns.FirstName] = model.FirstName
		case "lastName":
			modelM[models.AuthUserColumns.LastName] = model.LastName
		case "email":
			modelM[models.AuthUserColumns.Email] = model.Email
		case "isStaff":
			modelM[models.AuthUserColumns.IsStaff] = model.IsStaff
		case "isActive":
			modelM[models.AuthUserColumns.IsActive] = model.IsActive
		case "dateJoined":
			modelM[models.AuthUserColumns.DateJoined] = model.DateJoined
		}
	}
	return modelM
}

func AuthUserCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "password":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.Password)
		case "lastLogin":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.LastLogin)
		case "isSuperuser":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.IsSuperuser)
		case "username":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.Username)
		case "firstName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.FirstName)
		case "lastName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.LastName)
		case "email":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.Email)
		case "isStaff":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.IsStaff)
		case "isActive":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.IsActive)
		case "dateJoined":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.DateJoined)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthUserGroupCreateInputsToBoiler(am []*graphql_models.AuthUserGroupCreateInput) []*models.AuthUserGroup {
	ar := make([]*models.AuthUserGroup, len(am))
	for i, m := range am {
		ar[i] = AuthUserGroupCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthUserGroupCreateInputToBoiler(
	m *graphql_models.AuthUserGroupCreateInput,
) *models.AuthUserGroup {
	if m == nil {
		return nil
	}

	r := &models.AuthUserGroup{
		UserID:		int(boilergql.IDToBoiler(m.UserID)),
		GroupID:	int(boilergql.IDToBoiler(m.GroupID)),
	}
	return r
}

func AuthUserGroupCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthUserGroupCreateInput,
) models.M {
	model := AuthUserGroupCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "userId":
			modelM[models.AuthUserGroupColumns.UserID] = model.UserID
		case "groupId":
			modelM[models.AuthUserGroupColumns.GroupID] = model.GroupID
		}
	}
	return modelM
}

func AuthUserGroupCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "userId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserGroupColumns.UserID)
		case "groupId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserGroupColumns.GroupID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthUserGroupUpdateInputsToBoiler(am []*graphql_models.AuthUserGroupUpdateInput) []*models.AuthUserGroup {
	ar := make([]*models.AuthUserGroup, len(am))
	for i, m := range am {
		ar[i] = AuthUserGroupUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthUserGroupUpdateInputToBoiler(
	m *graphql_models.AuthUserGroupUpdateInput,
) *models.AuthUserGroup {
	if m == nil {
		return nil
	}

	r := &models.AuthUserGroup{
		UserID:		int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.UserID))),
		GroupID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.GroupID))),
	}
	return r
}

func AuthUserGroupUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthUserGroupUpdateInput,
) models.M {
	model := AuthUserGroupUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "userId":
			modelM[models.AuthUserGroupColumns.UserID] = model.UserID
		case "groupId":
			modelM[models.AuthUserGroupColumns.GroupID] = model.GroupID
		}
	}
	return modelM
}

func AuthUserGroupUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "userId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserGroupColumns.UserID)
		case "groupId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserGroupColumns.GroupID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthUserUpdateInputsToBoiler(am []*graphql_models.AuthUserUpdateInput) []*models.AuthUser {
	ar := make([]*models.AuthUser, len(am))
	for i, m := range am {
		ar[i] = AuthUserUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthUserUpdateInputToBoiler(
	m *graphql_models.AuthUserUpdateInput,
) *models.AuthUser {
	if m == nil {
		return nil
	}

	r := &models.AuthUser{
		Password:	boilergql.PointerStringToString(m.Password),
		LastLogin:	boilergql.PointerIntToNullDotTime(m.LastLogin),
		IsSuperuser:	boilergql.PointerBoolToBool(m.IsSuperuser),
		Username:	boilergql.PointerStringToString(m.Username),
		FirstName:	boilergql.PointerStringToString(m.FirstName),
		LastName:	boilergql.PointerStringToString(m.LastName),
		Email:		boilergql.PointerStringToString(m.Email),
		IsStaff:	boilergql.PointerBoolToBool(m.IsStaff),
		IsActive:	boilergql.PointerBoolToBool(m.IsActive),
		DateJoined:	boilergql.PointerIntToTimeDotTime(m.DateJoined),
	}
	return r
}

func AuthUserUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthUserUpdateInput,
) models.M {
	model := AuthUserUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "password":
			modelM[models.AuthUserColumns.Password] = model.Password
		case "lastLogin":
			modelM[models.AuthUserColumns.LastLogin] = model.LastLogin
		case "isSuperuser":
			modelM[models.AuthUserColumns.IsSuperuser] = model.IsSuperuser
		case "username":
			modelM[models.AuthUserColumns.Username] = model.Username
		case "firstName":
			modelM[models.AuthUserColumns.FirstName] = model.FirstName
		case "lastName":
			modelM[models.AuthUserColumns.LastName] = model.LastName
		case "email":
			modelM[models.AuthUserColumns.Email] = model.Email
		case "isStaff":
			modelM[models.AuthUserColumns.IsStaff] = model.IsStaff
		case "isActive":
			modelM[models.AuthUserColumns.IsActive] = model.IsActive
		case "dateJoined":
			modelM[models.AuthUserColumns.DateJoined] = model.DateJoined
		}
	}
	return modelM
}

func AuthUserUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "password":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.Password)
		case "lastLogin":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.LastLogin)
		case "isSuperuser":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.IsSuperuser)
		case "username":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.Username)
		case "firstName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.FirstName)
		case "lastName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.LastName)
		case "email":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.Email)
		case "isStaff":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.IsStaff)
		case "isActive":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.IsActive)
		case "dateJoined":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserColumns.DateJoined)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthUserUserPermissionCreateInputsToBoiler(am []*graphql_models.AuthUserUserPermissionCreateInput) []*models.AuthUserUserPermission {
	ar := make([]*models.AuthUserUserPermission, len(am))
	for i, m := range am {
		ar[i] = AuthUserUserPermissionCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthUserUserPermissionCreateInputToBoiler(
	m *graphql_models.AuthUserUserPermissionCreateInput,
) *models.AuthUserUserPermission {
	if m == nil {
		return nil
	}

	r := &models.AuthUserUserPermission{
		UserID:		int(boilergql.IDToBoiler(m.UserID)),
		PermissionID:	int(boilergql.IDToBoiler(m.PermissionID)),
	}
	return r
}

func AuthUserUserPermissionCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthUserUserPermissionCreateInput,
) models.M {
	model := AuthUserUserPermissionCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "userId":
			modelM[models.AuthUserUserPermissionColumns.UserID] = model.UserID
		case "permissionId":
			modelM[models.AuthUserUserPermissionColumns.PermissionID] = model.PermissionID
		}
	}
	return modelM
}

func AuthUserUserPermissionCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "userId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserUserPermissionColumns.UserID)
		case "permissionId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserUserPermissionColumns.PermissionID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func AuthUserUserPermissionUpdateInputsToBoiler(am []*graphql_models.AuthUserUserPermissionUpdateInput) []*models.AuthUserUserPermission {
	ar := make([]*models.AuthUserUserPermission, len(am))
	for i, m := range am {
		ar[i] = AuthUserUserPermissionUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func AuthUserUserPermissionUpdateInputToBoiler(
	m *graphql_models.AuthUserUserPermissionUpdateInput,
) *models.AuthUserUserPermission {
	if m == nil {
		return nil
	}

	r := &models.AuthUserUserPermission{
		UserID:		int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.UserID))),
		PermissionID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.PermissionID))),
	}
	return r
}

func AuthUserUserPermissionUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.AuthUserUserPermissionUpdateInput,
) models.M {
	model := AuthUserUserPermissionUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "userId":
			modelM[models.AuthUserUserPermissionColumns.UserID] = model.UserID
		case "permissionId":
			modelM[models.AuthUserUserPermissionColumns.PermissionID] = model.PermissionID
		}
	}
	return modelM
}

func AuthUserUserPermissionUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "userId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserUserPermissionColumns.UserID)
		case "permissionId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.AuthUserUserPermissionColumns.PermissionID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FragranceCreateInputsToBoiler(am []*graphql_models.FragranceCreateInput) []*models.Fragrance {
	ar := make([]*models.Fragrance, len(am))
	for i, m := range am {
		ar[i] = FragranceCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func FragranceCreateInputToBoiler(
	m *graphql_models.FragranceCreateInput,
) *models.Fragrance {
	if m == nil {
		return nil
	}

	r := &models.Fragrance{
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func FragranceCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.FragranceCreateInput,
) models.M {
	model := FragranceCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.FragranceColumns.Name] = model.Name
		case "note":
			modelM[models.FragranceColumns.Note] = model.Note
		case "createdAt":
			modelM[models.FragranceColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.FragranceColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.FragranceColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func FragranceCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FragranceInventoryCreateInputsToBoiler(am []*graphql_models.FragranceInventoryCreateInput) []*models.FragranceInventory {
	ar := make([]*models.FragranceInventory, len(am))
	for i, m := range am {
		ar[i] = FragranceInventoryCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func FragranceInventoryCreateInputToBoiler(
	m *graphql_models.FragranceInventoryCreateInput,
) *models.FragranceInventory {
	if m == nil {
		return nil
	}

	r := &models.FragranceInventory{
		PurchaseDate:	boilergql.IntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.IntToTimeDotTime(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,
		FragranceID:	int(boilergql.IDToBoiler(m.FragranceID)),
		SupplierID:	int(boilergql.IDToBoiler(m.SupplierID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func FragranceInventoryCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.FragranceInventoryCreateInput,
) models.M {
	model := FragranceInventoryCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.FragranceInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.FragranceInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.FragranceInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.FragranceInventoryColumns.Weight] = model.Weight
		case "fragranceId":
			modelM[models.FragranceInventoryColumns.FragranceID] = model.FragranceID
		case "supplierId":
			modelM[models.FragranceInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.FragranceInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.FragranceInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.FragranceInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func FragranceInventoryCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.Weight)
		case "fragranceId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.FragranceID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FragranceInventoryUpdateInputsToBoiler(am []*graphql_models.FragranceInventoryUpdateInput) []*models.FragranceInventory {
	ar := make([]*models.FragranceInventory, len(am))
	for i, m := range am {
		ar[i] = FragranceInventoryUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func FragranceInventoryUpdateInputToBoiler(
	m *graphql_models.FragranceInventoryUpdateInput,
) *models.FragranceInventory {
	if m == nil {
		return nil
	}

	r := &models.FragranceInventory{
		PurchaseDate:	boilergql.PointerIntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.PointerIntToTimeDotTime(m.ExpiryDate),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		FragranceID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.FragranceID))),
		SupplierID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.SupplierID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func FragranceInventoryUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.FragranceInventoryUpdateInput,
) models.M {
	model := FragranceInventoryUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.FragranceInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.FragranceInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.FragranceInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.FragranceInventoryColumns.Weight] = model.Weight
		case "fragranceId":
			modelM[models.FragranceInventoryColumns.FragranceID] = model.FragranceID
		case "supplierId":
			modelM[models.FragranceInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.FragranceInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.FragranceInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.FragranceInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func FragranceInventoryUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.Weight)
		case "fragranceId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.FragranceID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FragranceUpdateInputsToBoiler(am []*graphql_models.FragranceUpdateInput) []*models.Fragrance {
	ar := make([]*models.Fragrance, len(am))
	for i, m := range am {
		ar[i] = FragranceUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func FragranceUpdateInputToBoiler(
	m *graphql_models.FragranceUpdateInput,
) *models.Fragrance {
	if m == nil {
		return nil
	}

	r := &models.Fragrance{
		Name:		boilergql.PointerStringToString(m.Name),
		Note:		boilergql.PointerStringToString(m.Note),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func FragranceUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.FragranceUpdateInput,
) models.M {
	model := FragranceUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.FragranceColumns.Name] = model.Name
		case "note":
			modelM[models.FragranceColumns.Note] = model.Note
		case "createdAt":
			modelM[models.FragranceColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.FragranceColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.FragranceColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func FragranceUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FragranceColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LipidCreateInputsToBoiler(am []*graphql_models.LipidCreateInput) []*models.Lipid {
	ar := make([]*models.Lipid, len(am))
	for i, m := range am {
		ar[i] = LipidCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func LipidCreateInputToBoiler(
	m *graphql_models.LipidCreateInput,
) *models.Lipid {
	if m == nil {
		return nil
	}

	r := &models.Lipid{
		Name:		m.Name,
		Lauric:		m.Lauric,
		Myristic:	m.Myristic,
		Palmitic:	m.Palmitic,
		Stearic:	m.Stearic,
		Ricinoleic:	m.Ricinoleic,
		Oleic:		m.Oleic,
		Linoleic:	m.Linoleic,
		Linolenic:	m.Linolenic,
		Hardness:	m.Hardness,
		Cleansing:	m.Cleansing,
		Conditioning:	m.Conditioning,
		Bubbly:		m.Bubbly,
		Creamy:		m.Creamy,
		Iodine:		m.Iodine,
		Ins:		m.Ins,
		InciName:	m.InciName,
		Family:		m.Family,
		Naoh:		m.Naoh,
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LipidCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.LipidCreateInput,
) models.M {
	model := LipidCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.LipidColumns.Name] = model.Name
		case "lauric":
			modelM[models.LipidColumns.Lauric] = model.Lauric
		case "myristic":
			modelM[models.LipidColumns.Myristic] = model.Myristic
		case "palmitic":
			modelM[models.LipidColumns.Palmitic] = model.Palmitic
		case "stearic":
			modelM[models.LipidColumns.Stearic] = model.Stearic
		case "ricinoleic":
			modelM[models.LipidColumns.Ricinoleic] = model.Ricinoleic
		case "oleic":
			modelM[models.LipidColumns.Oleic] = model.Oleic
		case "linoleic":
			modelM[models.LipidColumns.Linoleic] = model.Linoleic
		case "linolenic":
			modelM[models.LipidColumns.Linolenic] = model.Linolenic
		case "hardness":
			modelM[models.LipidColumns.Hardness] = model.Hardness
		case "cleansing":
			modelM[models.LipidColumns.Cleansing] = model.Cleansing
		case "conditioning":
			modelM[models.LipidColumns.Conditioning] = model.Conditioning
		case "bubbly":
			modelM[models.LipidColumns.Bubbly] = model.Bubbly
		case "creamy":
			modelM[models.LipidColumns.Creamy] = model.Creamy
		case "iodine":
			modelM[models.LipidColumns.Iodine] = model.Iodine
		case "ins":
			modelM[models.LipidColumns.Ins] = model.Ins
		case "inciName":
			modelM[models.LipidColumns.InciName] = model.InciName
		case "family":
			modelM[models.LipidColumns.Family] = model.Family
		case "naoh":
			modelM[models.LipidColumns.Naoh] = model.Naoh
		case "createdAt":
			modelM[models.LipidColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LipidColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LipidColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LipidCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Name)
		case "lauric":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Lauric)
		case "myristic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Myristic)
		case "palmitic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Palmitic)
		case "stearic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Stearic)
		case "ricinoleic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Ricinoleic)
		case "oleic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Oleic)
		case "linoleic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Linoleic)
		case "linolenic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Linolenic)
		case "hardness":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Hardness)
		case "cleansing":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Cleansing)
		case "conditioning":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Conditioning)
		case "bubbly":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Bubbly)
		case "creamy":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Creamy)
		case "iodine":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Iodine)
		case "ins":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Ins)
		case "inciName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.InciName)
		case "family":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Family)
		case "naoh":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Naoh)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LipidInventoryCreateInputsToBoiler(am []*graphql_models.LipidInventoryCreateInput) []*models.LipidInventory {
	ar := make([]*models.LipidInventory, len(am))
	for i, m := range am {
		ar[i] = LipidInventoryCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func LipidInventoryCreateInputToBoiler(
	m *graphql_models.LipidInventoryCreateInput,
) *models.LipidInventory {
	if m == nil {
		return nil
	}

	r := &models.LipidInventory{
		PurchaseDate:	boilergql.IntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.IntToTimeDotTime(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,
		Sap:		m.Sap,
		Naoh:		m.Naoh,
		Koh:		m.Koh,
		GramsPerLiter:	m.GramsPerLiter,
		LipidID:	int(boilergql.IDToBoiler(m.LipidID)),
		SupplierID:	int(boilergql.IDToBoiler(m.SupplierID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LipidInventoryCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.LipidInventoryCreateInput,
) models.M {
	model := LipidInventoryCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.LipidInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.LipidInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.LipidInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.LipidInventoryColumns.Weight] = model.Weight
		case "sap":
			modelM[models.LipidInventoryColumns.Sap] = model.Sap
		case "naoh":
			modelM[models.LipidInventoryColumns.Naoh] = model.Naoh
		case "koh":
			modelM[models.LipidInventoryColumns.Koh] = model.Koh
		case "gramsPerLiter":
			modelM[models.LipidInventoryColumns.GramsPerLiter] = model.GramsPerLiter
		case "lipidId":
			modelM[models.LipidInventoryColumns.LipidID] = model.LipidID
		case "supplierId":
			modelM[models.LipidInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.LipidInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LipidInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LipidInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LipidInventoryCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Weight)
		case "sap":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Sap)
		case "naoh":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Naoh)
		case "koh":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Koh)
		case "gramsPerLiter":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.GramsPerLiter)
		case "lipidId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.LipidID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LipidInventoryUpdateInputsToBoiler(am []*graphql_models.LipidInventoryUpdateInput) []*models.LipidInventory {
	ar := make([]*models.LipidInventory, len(am))
	for i, m := range am {
		ar[i] = LipidInventoryUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func LipidInventoryUpdateInputToBoiler(
	m *graphql_models.LipidInventoryUpdateInput,
) *models.LipidInventory {
	if m == nil {
		return nil
	}

	r := &models.LipidInventory{
		PurchaseDate:	boilergql.PointerIntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.PointerIntToTimeDotTime(m.ExpiryDate),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		Sap:		boilergql.PointerFloat64ToFloat64(m.Sap),
		Naoh:		boilergql.PointerFloat64ToFloat64(m.Naoh),
		Koh:		boilergql.PointerFloat64ToFloat64(m.Koh),
		GramsPerLiter:	boilergql.PointerFloat64ToFloat64(m.GramsPerLiter),
		LipidID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.LipidID))),
		SupplierID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.SupplierID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LipidInventoryUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.LipidInventoryUpdateInput,
) models.M {
	model := LipidInventoryUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.LipidInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.LipidInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.LipidInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.LipidInventoryColumns.Weight] = model.Weight
		case "sap":
			modelM[models.LipidInventoryColumns.Sap] = model.Sap
		case "naoh":
			modelM[models.LipidInventoryColumns.Naoh] = model.Naoh
		case "koh":
			modelM[models.LipidInventoryColumns.Koh] = model.Koh
		case "gramsPerLiter":
			modelM[models.LipidInventoryColumns.GramsPerLiter] = model.GramsPerLiter
		case "lipidId":
			modelM[models.LipidInventoryColumns.LipidID] = model.LipidID
		case "supplierId":
			modelM[models.LipidInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.LipidInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LipidInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LipidInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LipidInventoryUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Weight)
		case "sap":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Sap)
		case "naoh":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Naoh)
		case "koh":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.Koh)
		case "gramsPerLiter":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.GramsPerLiter)
		case "lipidId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.LipidID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LipidUpdateInputsToBoiler(am []*graphql_models.LipidUpdateInput) []*models.Lipid {
	ar := make([]*models.Lipid, len(am))
	for i, m := range am {
		ar[i] = LipidUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func LipidUpdateInputToBoiler(
	m *graphql_models.LipidUpdateInput,
) *models.Lipid {
	if m == nil {
		return nil
	}

	r := &models.Lipid{
		Name:		boilergql.PointerStringToString(m.Name),
		Lauric:		boilergql.PointerIntToInt(m.Lauric),
		Myristic:	boilergql.PointerIntToInt(m.Myristic),
		Palmitic:	boilergql.PointerIntToInt(m.Palmitic),
		Stearic:	boilergql.PointerIntToInt(m.Stearic),
		Ricinoleic:	boilergql.PointerIntToInt(m.Ricinoleic),
		Oleic:		boilergql.PointerIntToInt(m.Oleic),
		Linoleic:	boilergql.PointerIntToInt(m.Linoleic),
		Linolenic:	boilergql.PointerIntToInt(m.Linolenic),
		Hardness:	boilergql.PointerIntToInt(m.Hardness),
		Cleansing:	boilergql.PointerIntToInt(m.Cleansing),
		Conditioning:	boilergql.PointerIntToInt(m.Conditioning),
		Bubbly:		boilergql.PointerIntToInt(m.Bubbly),
		Creamy:		boilergql.PointerIntToInt(m.Creamy),
		Iodine:		boilergql.PointerIntToInt(m.Iodine),
		Ins:		boilergql.PointerIntToInt(m.Ins),
		InciName:	boilergql.PointerStringToString(m.InciName),
		Family:		boilergql.PointerStringToString(m.Family),
		Naoh:		boilergql.PointerFloat64ToFloat64(m.Naoh),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LipidUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.LipidUpdateInput,
) models.M {
	model := LipidUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.LipidColumns.Name] = model.Name
		case "lauric":
			modelM[models.LipidColumns.Lauric] = model.Lauric
		case "myristic":
			modelM[models.LipidColumns.Myristic] = model.Myristic
		case "palmitic":
			modelM[models.LipidColumns.Palmitic] = model.Palmitic
		case "stearic":
			modelM[models.LipidColumns.Stearic] = model.Stearic
		case "ricinoleic":
			modelM[models.LipidColumns.Ricinoleic] = model.Ricinoleic
		case "oleic":
			modelM[models.LipidColumns.Oleic] = model.Oleic
		case "linoleic":
			modelM[models.LipidColumns.Linoleic] = model.Linoleic
		case "linolenic":
			modelM[models.LipidColumns.Linolenic] = model.Linolenic
		case "hardness":
			modelM[models.LipidColumns.Hardness] = model.Hardness
		case "cleansing":
			modelM[models.LipidColumns.Cleansing] = model.Cleansing
		case "conditioning":
			modelM[models.LipidColumns.Conditioning] = model.Conditioning
		case "bubbly":
			modelM[models.LipidColumns.Bubbly] = model.Bubbly
		case "creamy":
			modelM[models.LipidColumns.Creamy] = model.Creamy
		case "iodine":
			modelM[models.LipidColumns.Iodine] = model.Iodine
		case "ins":
			modelM[models.LipidColumns.Ins] = model.Ins
		case "inciName":
			modelM[models.LipidColumns.InciName] = model.InciName
		case "family":
			modelM[models.LipidColumns.Family] = model.Family
		case "naoh":
			modelM[models.LipidColumns.Naoh] = model.Naoh
		case "createdAt":
			modelM[models.LipidColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LipidColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LipidColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LipidUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Name)
		case "lauric":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Lauric)
		case "myristic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Myristic)
		case "palmitic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Palmitic)
		case "stearic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Stearic)
		case "ricinoleic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Ricinoleic)
		case "oleic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Oleic)
		case "linoleic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Linoleic)
		case "linolenic":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Linolenic)
		case "hardness":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Hardness)
		case "cleansing":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Cleansing)
		case "conditioning":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Conditioning)
		case "bubbly":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Bubbly)
		case "creamy":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Creamy)
		case "iodine":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Iodine)
		case "ins":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Ins)
		case "inciName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.InciName)
		case "family":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Family)
		case "naoh":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.Naoh)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LipidColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LyeCreateInputsToBoiler(am []*graphql_models.LyeCreateInput) []*models.Lye {
	ar := make([]*models.Lye, len(am))
	for i, m := range am {
		ar[i] = LyeCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func LyeCreateInputToBoiler(
	m *graphql_models.LyeCreateInput,
) *models.Lye {
	if m == nil {
		return nil
	}

	r := &models.Lye{
		Kind:		m.Kind,
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LyeCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.LyeCreateInput,
) models.M {
	model := LyeCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "kind":
			modelM[models.LyeColumns.Kind] = model.Kind
		case "name":
			modelM[models.LyeColumns.Name] = model.Name
		case "note":
			modelM[models.LyeColumns.Note] = model.Note
		case "createdAt":
			modelM[models.LyeColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LyeColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LyeColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LyeCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "kind":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.Kind)
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LyeInventoryCreateInputsToBoiler(am []*graphql_models.LyeInventoryCreateInput) []*models.LyeInventory {
	ar := make([]*models.LyeInventory, len(am))
	for i, m := range am {
		ar[i] = LyeInventoryCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func LyeInventoryCreateInputToBoiler(
	m *graphql_models.LyeInventoryCreateInput,
) *models.LyeInventory {
	if m == nil {
		return nil
	}

	r := &models.LyeInventory{
		PurchaseDate:	boilergql.IntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.IntToTimeDotTime(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,
		Concentration:	m.Concentration,
		LyeID:		int(boilergql.IDToBoiler(m.LyeID)),
		SupplierID:	int(boilergql.IDToBoiler(m.SupplierID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LyeInventoryCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.LyeInventoryCreateInput,
) models.M {
	model := LyeInventoryCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.LyeInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.LyeInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.LyeInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.LyeInventoryColumns.Weight] = model.Weight
		case "concentration":
			modelM[models.LyeInventoryColumns.Concentration] = model.Concentration
		case "lyeId":
			modelM[models.LyeInventoryColumns.LyeID] = model.LyeID
		case "supplierId":
			modelM[models.LyeInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.LyeInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LyeInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LyeInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LyeInventoryCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.Weight)
		case "concentration":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.Concentration)
		case "lyeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.LyeID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LyeInventoryUpdateInputsToBoiler(am []*graphql_models.LyeInventoryUpdateInput) []*models.LyeInventory {
	ar := make([]*models.LyeInventory, len(am))
	for i, m := range am {
		ar[i] = LyeInventoryUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func LyeInventoryUpdateInputToBoiler(
	m *graphql_models.LyeInventoryUpdateInput,
) *models.LyeInventory {
	if m == nil {
		return nil
	}

	r := &models.LyeInventory{
		PurchaseDate:	boilergql.PointerIntToTimeDotTime(m.PurchaseDate),
		ExpiryDate:	boilergql.PointerIntToTimeDotTime(m.ExpiryDate),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		Concentration:	boilergql.PointerFloat64ToFloat64(m.Concentration),
		LyeID:		int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.LyeID))),
		SupplierID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.SupplierID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LyeInventoryUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.LyeInventoryUpdateInput,
) models.M {
	model := LyeInventoryUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "purchaseDate":
			modelM[models.LyeInventoryColumns.PurchaseDate] = model.PurchaseDate
		case "expiryDate":
			modelM[models.LyeInventoryColumns.ExpiryDate] = model.ExpiryDate
		case "cost":
			modelM[models.LyeInventoryColumns.Cost] = model.Cost
		case "weight":
			modelM[models.LyeInventoryColumns.Weight] = model.Weight
		case "concentration":
			modelM[models.LyeInventoryColumns.Concentration] = model.Concentration
		case "lyeId":
			modelM[models.LyeInventoryColumns.LyeID] = model.LyeID
		case "supplierId":
			modelM[models.LyeInventoryColumns.SupplierID] = model.SupplierID
		case "createdAt":
			modelM[models.LyeInventoryColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LyeInventoryColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LyeInventoryColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LyeInventoryUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "purchaseDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.PurchaseDate)
		case "expiryDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.ExpiryDate)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.Cost)
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.Weight)
		case "concentration":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.Concentration)
		case "lyeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.LyeID)
		case "supplierId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.SupplierID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeInventoryColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LyeUpdateInputsToBoiler(am []*graphql_models.LyeUpdateInput) []*models.Lye {
	ar := make([]*models.Lye, len(am))
	for i, m := range am {
		ar[i] = LyeUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func LyeUpdateInputToBoiler(
	m *graphql_models.LyeUpdateInput,
) *models.Lye {
	if m == nil {
		return nil
	}

	r := &models.Lye{
		Kind:		boilergql.PointerStringToString(m.Kind),
		Name:		boilergql.PointerStringToString(m.Name),
		Note:		boilergql.PointerStringToString(m.Note),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func LyeUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.LyeUpdateInput,
) models.M {
	model := LyeUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "kind":
			modelM[models.LyeColumns.Kind] = model.Kind
		case "name":
			modelM[models.LyeColumns.Name] = model.Name
		case "note":
			modelM[models.LyeColumns.Note] = model.Note
		case "createdAt":
			modelM[models.LyeColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.LyeColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.LyeColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func LyeUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "kind":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.Kind)
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LyeColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeAdditiveCreateInputsToBoiler(am []*graphql_models.RecipeAdditiveCreateInput) []*models.RecipeAdditive {
	ar := make([]*models.RecipeAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeAdditiveCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeAdditiveCreateInputToBoiler(
	m *graphql_models.RecipeAdditiveCreateInput,
) *models.RecipeAdditive {
	if m == nil {
		return nil
	}

	r := &models.RecipeAdditive{
		Percentage:	m.Percentage,
		AdditiveID:	int(boilergql.IDToBoiler(m.AdditiveID)),
		RecipeID:	int(boilergql.IDToBoiler(m.RecipeID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeAdditiveCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeAdditiveCreateInput,
) models.M {
	model := RecipeAdditiveCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "percentage":
			modelM[models.RecipeAdditiveColumns.Percentage] = model.Percentage
		case "additiveId":
			modelM[models.RecipeAdditiveColumns.AdditiveID] = model.AdditiveID
		case "recipeId":
			modelM[models.RecipeAdditiveColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeAdditiveColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeAdditiveColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeAdditiveColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeAdditiveCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "percentage":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.Percentage)
		case "additiveId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.AdditiveID)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeAdditiveUpdateInputsToBoiler(am []*graphql_models.RecipeAdditiveUpdateInput) []*models.RecipeAdditive {
	ar := make([]*models.RecipeAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeAdditiveUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeAdditiveUpdateInputToBoiler(
	m *graphql_models.RecipeAdditiveUpdateInput,
) *models.RecipeAdditive {
	if m == nil {
		return nil
	}

	r := &models.RecipeAdditive{
		Percentage:	boilergql.PointerFloat64ToFloat64(m.Percentage),
		AdditiveID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.AdditiveID))),
		RecipeID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.RecipeID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeAdditiveUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeAdditiveUpdateInput,
) models.M {
	model := RecipeAdditiveUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "percentage":
			modelM[models.RecipeAdditiveColumns.Percentage] = model.Percentage
		case "additiveId":
			modelM[models.RecipeAdditiveColumns.AdditiveID] = model.AdditiveID
		case "recipeId":
			modelM[models.RecipeAdditiveColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeAdditiveColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeAdditiveColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeAdditiveColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeAdditiveUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "percentage":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.Percentage)
		case "additiveId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.AdditiveID)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeAdditiveColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchAdditiveCreateInputsToBoiler(am []*graphql_models.RecipeBatchAdditiveCreateInput) []*models.RecipeBatchAdditive {
	ar := make([]*models.RecipeBatchAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchAdditiveCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchAdditiveCreateInputToBoiler(
	m *graphql_models.RecipeBatchAdditiveCreateInput,
) *models.RecipeBatchAdditive {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchAdditive{
		Weight:		m.Weight,
		Cost:		m.Cost,
		AdditiveID:	int(boilergql.IDToBoiler(m.AdditiveID)),
		BatchID:	int(boilergql.IDToBoiler(m.BatchID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchAdditiveCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchAdditiveCreateInput,
) models.M {
	model := RecipeBatchAdditiveCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchAdditiveColumns.Weight] = model.Weight
		case "cost":
			modelM[models.RecipeBatchAdditiveColumns.Cost] = model.Cost
		case "additiveId":
			modelM[models.RecipeBatchAdditiveColumns.AdditiveID] = model.AdditiveID
		case "batchId":
			modelM[models.RecipeBatchAdditiveColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchAdditiveColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchAdditiveColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchAdditiveColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchAdditiveCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.Weight)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.Cost)
		case "additiveId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.AdditiveID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchAdditiveUpdateInputsToBoiler(am []*graphql_models.RecipeBatchAdditiveUpdateInput) []*models.RecipeBatchAdditive {
	ar := make([]*models.RecipeBatchAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchAdditiveUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchAdditiveUpdateInputToBoiler(
	m *graphql_models.RecipeBatchAdditiveUpdateInput,
) *models.RecipeBatchAdditive {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchAdditive{
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		AdditiveID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.AdditiveID))),
		BatchID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.BatchID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchAdditiveUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchAdditiveUpdateInput,
) models.M {
	model := RecipeBatchAdditiveUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchAdditiveColumns.Weight] = model.Weight
		case "cost":
			modelM[models.RecipeBatchAdditiveColumns.Cost] = model.Cost
		case "additiveId":
			modelM[models.RecipeBatchAdditiveColumns.AdditiveID] = model.AdditiveID
		case "batchId":
			modelM[models.RecipeBatchAdditiveColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchAdditiveColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchAdditiveColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchAdditiveColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchAdditiveUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.Weight)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.Cost)
		case "additiveId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.AdditiveID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchAdditiveColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchCreateInputsToBoiler(am []*graphql_models.RecipeBatchCreateInput) []*models.RecipeBatch {
	ar := make([]*models.RecipeBatch, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchCreateInputToBoiler(
	m *graphql_models.RecipeBatchCreateInput,
) *models.RecipeBatch {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatch{
		Tag:			m.Tag,
		ProductionDate:		boilergql.IntToTimeDotTime(m.ProductionDate),
		SellableDate:		boilergql.IntToTimeDotTime(m.SellableDate),
		Note:			m.Note,
		LipidWeight:		m.LipidWeight,
		ProductionWeight:	m.ProductionWeight,
		CuredWeight:		m.CuredWeight,
		RecipeID:		int(boilergql.IDToBoiler(m.RecipeID)),
		CreatedAt:		boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:		boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:		boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchCreateInput,
) models.M {
	model := RecipeBatchCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "tag":
			modelM[models.RecipeBatchColumns.Tag] = model.Tag
		case "productionDate":
			modelM[models.RecipeBatchColumns.ProductionDate] = model.ProductionDate
		case "sellableDate":
			modelM[models.RecipeBatchColumns.SellableDate] = model.SellableDate
		case "note":
			modelM[models.RecipeBatchColumns.Note] = model.Note
		case "lipidWeight":
			modelM[models.RecipeBatchColumns.LipidWeight] = model.LipidWeight
		case "productionWeight":
			modelM[models.RecipeBatchColumns.ProductionWeight] = model.ProductionWeight
		case "curedWeight":
			modelM[models.RecipeBatchColumns.CuredWeight] = model.CuredWeight
		case "recipeId":
			modelM[models.RecipeBatchColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeBatchColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "tag":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.Tag)
		case "productionDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.ProductionDate)
		case "sellableDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.SellableDate)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.Note)
		case "lipidWeight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.LipidWeight)
		case "productionWeight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.ProductionWeight)
		case "curedWeight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.CuredWeight)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchFragranceCreateInputsToBoiler(am []*graphql_models.RecipeBatchFragranceCreateInput) []*models.RecipeBatchFragrance {
	ar := make([]*models.RecipeBatchFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchFragranceCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchFragranceCreateInputToBoiler(
	m *graphql_models.RecipeBatchFragranceCreateInput,
) *models.RecipeBatchFragrance {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchFragrance{
		Weight:		m.Weight,
		Cost:		m.Cost,
		FragranceID:	int(boilergql.IDToBoiler(m.FragranceID)),
		BatchID:	int(boilergql.IDToBoiler(m.BatchID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchFragranceCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchFragranceCreateInput,
) models.M {
	model := RecipeBatchFragranceCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchFragranceColumns.Weight] = model.Weight
		case "cost":
			modelM[models.RecipeBatchFragranceColumns.Cost] = model.Cost
		case "fragranceId":
			modelM[models.RecipeBatchFragranceColumns.FragranceID] = model.FragranceID
		case "batchId":
			modelM[models.RecipeBatchFragranceColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchFragranceColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchFragranceColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchFragranceColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchFragranceCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.Weight)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.Cost)
		case "fragranceId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.FragranceID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchFragranceUpdateInputsToBoiler(am []*graphql_models.RecipeBatchFragranceUpdateInput) []*models.RecipeBatchFragrance {
	ar := make([]*models.RecipeBatchFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchFragranceUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchFragranceUpdateInputToBoiler(
	m *graphql_models.RecipeBatchFragranceUpdateInput,
) *models.RecipeBatchFragrance {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchFragrance{
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		FragranceID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.FragranceID))),
		BatchID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.BatchID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchFragranceUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchFragranceUpdateInput,
) models.M {
	model := RecipeBatchFragranceUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchFragranceColumns.Weight] = model.Weight
		case "cost":
			modelM[models.RecipeBatchFragranceColumns.Cost] = model.Cost
		case "fragranceId":
			modelM[models.RecipeBatchFragranceColumns.FragranceID] = model.FragranceID
		case "batchId":
			modelM[models.RecipeBatchFragranceColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchFragranceColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchFragranceColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchFragranceColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchFragranceUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.Weight)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.Cost)
		case "fragranceId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.FragranceID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchFragranceColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchLipidCreateInputsToBoiler(am []*graphql_models.RecipeBatchLipidCreateInput) []*models.RecipeBatchLipid {
	ar := make([]*models.RecipeBatchLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLipidCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchLipidCreateInputToBoiler(
	m *graphql_models.RecipeBatchLipidCreateInput,
) *models.RecipeBatchLipid {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchLipid{
		Weight:		m.Weight,
		Cost:		m.Cost,
		LipidID:	int(boilergql.IDToBoiler(m.LipidID)),
		BatchID:	int(boilergql.IDToBoiler(m.BatchID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchLipidCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchLipidCreateInput,
) models.M {
	model := RecipeBatchLipidCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchLipidColumns.Weight] = model.Weight
		case "cost":
			modelM[models.RecipeBatchLipidColumns.Cost] = model.Cost
		case "lipidId":
			modelM[models.RecipeBatchLipidColumns.LipidID] = model.LipidID
		case "batchId":
			modelM[models.RecipeBatchLipidColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchLipidColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchLipidColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchLipidColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchLipidCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.Weight)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.Cost)
		case "lipidId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.LipidID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchLipidUpdateInputsToBoiler(am []*graphql_models.RecipeBatchLipidUpdateInput) []*models.RecipeBatchLipid {
	ar := make([]*models.RecipeBatchLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLipidUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchLipidUpdateInputToBoiler(
	m *graphql_models.RecipeBatchLipidUpdateInput,
) *models.RecipeBatchLipid {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchLipid{
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		LipidID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.LipidID))),
		BatchID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.BatchID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchLipidUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchLipidUpdateInput,
) models.M {
	model := RecipeBatchLipidUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchLipidColumns.Weight] = model.Weight
		case "cost":
			modelM[models.RecipeBatchLipidColumns.Cost] = model.Cost
		case "lipidId":
			modelM[models.RecipeBatchLipidColumns.LipidID] = model.LipidID
		case "batchId":
			modelM[models.RecipeBatchLipidColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchLipidColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchLipidColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchLipidColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchLipidUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.Weight)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.Cost)
		case "lipidId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.LipidID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLipidColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchLyeCreateInputsToBoiler(am []*graphql_models.RecipeBatchLyeCreateInput) []*models.RecipeBatchLye {
	ar := make([]*models.RecipeBatchLye, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLyeCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchLyeCreateInputToBoiler(
	m *graphql_models.RecipeBatchLyeCreateInput,
) *models.RecipeBatchLye {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchLye{
		Weight:		m.Weight,
		Discount:	m.Discount,
		Cost:		m.Cost,
		LyeID:		int(boilergql.IDToBoiler(m.LyeID)),
		BatchID:	int(boilergql.IDToBoiler(m.BatchID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchLyeCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchLyeCreateInput,
) models.M {
	model := RecipeBatchLyeCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchLyeColumns.Weight] = model.Weight
		case "discount":
			modelM[models.RecipeBatchLyeColumns.Discount] = model.Discount
		case "cost":
			modelM[models.RecipeBatchLyeColumns.Cost] = model.Cost
		case "lyeId":
			modelM[models.RecipeBatchLyeColumns.LyeID] = model.LyeID
		case "batchId":
			modelM[models.RecipeBatchLyeColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchLyeColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchLyeColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchLyeColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchLyeCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.Weight)
		case "discount":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.Discount)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.Cost)
		case "lyeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.LyeID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchLyeUpdateInputsToBoiler(am []*graphql_models.RecipeBatchLyeUpdateInput) []*models.RecipeBatchLye {
	ar := make([]*models.RecipeBatchLye, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLyeUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchLyeUpdateInputToBoiler(
	m *graphql_models.RecipeBatchLyeUpdateInput,
) *models.RecipeBatchLye {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchLye{
		Weight:		boilergql.PointerFloat64ToFloat64(m.Weight),
		Discount:	boilergql.PointerFloat64ToFloat64(m.Discount),
		Cost:		boilergql.PointerFloat64ToFloat64(m.Cost),
		LyeID:		int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.LyeID))),
		BatchID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.BatchID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchLyeUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchLyeUpdateInput,
) models.M {
	model := RecipeBatchLyeUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "weight":
			modelM[models.RecipeBatchLyeColumns.Weight] = model.Weight
		case "discount":
			modelM[models.RecipeBatchLyeColumns.Discount] = model.Discount
		case "cost":
			modelM[models.RecipeBatchLyeColumns.Cost] = model.Cost
		case "lyeId":
			modelM[models.RecipeBatchLyeColumns.LyeID] = model.LyeID
		case "batchId":
			modelM[models.RecipeBatchLyeColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchLyeColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchLyeColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchLyeColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchLyeUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "weight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.Weight)
		case "discount":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.Discount)
		case "cost":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.Cost)
		case "lyeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.LyeID)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchLyeColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchNoteCreateInputsToBoiler(am []*graphql_models.RecipeBatchNoteCreateInput) []*models.RecipeBatchNote {
	ar := make([]*models.RecipeBatchNote, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchNoteCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchNoteCreateInputToBoiler(
	m *graphql_models.RecipeBatchNoteCreateInput,
) *models.RecipeBatchNote {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchNote{
		Note:		m.Note,
		Link:		m.Link,
		BatchID:	int(boilergql.IDToBoiler(m.BatchID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchNoteCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchNoteCreateInput,
) models.M {
	model := RecipeBatchNoteCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "note":
			modelM[models.RecipeBatchNoteColumns.Note] = model.Note
		case "link":
			modelM[models.RecipeBatchNoteColumns.Link] = model.Link
		case "batchId":
			modelM[models.RecipeBatchNoteColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchNoteColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchNoteColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchNoteColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchNoteCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.Note)
		case "link":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.Link)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchNoteUpdateInputsToBoiler(am []*graphql_models.RecipeBatchNoteUpdateInput) []*models.RecipeBatchNote {
	ar := make([]*models.RecipeBatchNote, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchNoteUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchNoteUpdateInputToBoiler(
	m *graphql_models.RecipeBatchNoteUpdateInput,
) *models.RecipeBatchNote {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatchNote{
		Note:		boilergql.PointerStringToString(m.Note),
		Link:		boilergql.PointerStringToString(m.Link),
		BatchID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.BatchID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchNoteUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchNoteUpdateInput,
) models.M {
	model := RecipeBatchNoteUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "note":
			modelM[models.RecipeBatchNoteColumns.Note] = model.Note
		case "link":
			modelM[models.RecipeBatchNoteColumns.Link] = model.Link
		case "batchId":
			modelM[models.RecipeBatchNoteColumns.BatchID] = model.BatchID
		case "createdAt":
			modelM[models.RecipeBatchNoteColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchNoteColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchNoteColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchNoteUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.Note)
		case "link":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.Link)
		case "batchId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.BatchID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchNoteColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeBatchUpdateInputsToBoiler(am []*graphql_models.RecipeBatchUpdateInput) []*models.RecipeBatch {
	ar := make([]*models.RecipeBatch, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeBatchUpdateInputToBoiler(
	m *graphql_models.RecipeBatchUpdateInput,
) *models.RecipeBatch {
	if m == nil {
		return nil
	}

	r := &models.RecipeBatch{
		Tag:			boilergql.PointerStringToString(m.Tag),
		ProductionDate:		boilergql.PointerIntToTimeDotTime(m.ProductionDate),
		SellableDate:		boilergql.PointerIntToTimeDotTime(m.SellableDate),
		Note:			boilergql.PointerStringToString(m.Note),
		LipidWeight:		boilergql.PointerFloat64ToFloat64(m.LipidWeight),
		ProductionWeight:	boilergql.PointerFloat64ToFloat64(m.ProductionWeight),
		CuredWeight:		boilergql.PointerFloat64ToFloat64(m.CuredWeight),
		RecipeID:		int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.RecipeID))),
		CreatedAt:		boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:		boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:		boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeBatchUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeBatchUpdateInput,
) models.M {
	model := RecipeBatchUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "tag":
			modelM[models.RecipeBatchColumns.Tag] = model.Tag
		case "productionDate":
			modelM[models.RecipeBatchColumns.ProductionDate] = model.ProductionDate
		case "sellableDate":
			modelM[models.RecipeBatchColumns.SellableDate] = model.SellableDate
		case "note":
			modelM[models.RecipeBatchColumns.Note] = model.Note
		case "lipidWeight":
			modelM[models.RecipeBatchColumns.LipidWeight] = model.LipidWeight
		case "productionWeight":
			modelM[models.RecipeBatchColumns.ProductionWeight] = model.ProductionWeight
		case "curedWeight":
			modelM[models.RecipeBatchColumns.CuredWeight] = model.CuredWeight
		case "recipeId":
			modelM[models.RecipeBatchColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeBatchColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeBatchColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeBatchColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeBatchUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "tag":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.Tag)
		case "productionDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.ProductionDate)
		case "sellableDate":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.SellableDate)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.Note)
		case "lipidWeight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.LipidWeight)
		case "productionWeight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.ProductionWeight)
		case "curedWeight":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.CuredWeight)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeBatchColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeCreateInputsToBoiler(am []*graphql_models.RecipeCreateInput) []*models.Recipe {
	ar := make([]*models.Recipe, len(am))
	for i, m := range am {
		ar[i] = RecipeCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeCreateInputToBoiler(
	m *graphql_models.RecipeCreateInput,
) *models.Recipe {
	if m == nil {
		return nil
	}

	r := &models.Recipe{
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeCreateInput,
) models.M {
	model := RecipeCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.RecipeColumns.Name] = model.Name
		case "note":
			modelM[models.RecipeColumns.Note] = model.Note
		case "createdAt":
			modelM[models.RecipeColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeFragranceCreateInputsToBoiler(am []*graphql_models.RecipeFragranceCreateInput) []*models.RecipeFragrance {
	ar := make([]*models.RecipeFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeFragranceCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeFragranceCreateInputToBoiler(
	m *graphql_models.RecipeFragranceCreateInput,
) *models.RecipeFragrance {
	if m == nil {
		return nil
	}

	r := &models.RecipeFragrance{
		Percentage:	m.Percentage,
		FragranceID:	int(boilergql.IDToBoiler(m.FragranceID)),
		RecipeID:	int(boilergql.IDToBoiler(m.RecipeID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeFragranceCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeFragranceCreateInput,
) models.M {
	model := RecipeFragranceCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "percentage":
			modelM[models.RecipeFragranceColumns.Percentage] = model.Percentage
		case "fragranceId":
			modelM[models.RecipeFragranceColumns.FragranceID] = model.FragranceID
		case "recipeId":
			modelM[models.RecipeFragranceColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeFragranceColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeFragranceColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeFragranceColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeFragranceCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "percentage":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.Percentage)
		case "fragranceId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.FragranceID)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeFragranceUpdateInputsToBoiler(am []*graphql_models.RecipeFragranceUpdateInput) []*models.RecipeFragrance {
	ar := make([]*models.RecipeFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeFragranceUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeFragranceUpdateInputToBoiler(
	m *graphql_models.RecipeFragranceUpdateInput,
) *models.RecipeFragrance {
	if m == nil {
		return nil
	}

	r := &models.RecipeFragrance{
		Percentage:	boilergql.PointerFloat64ToFloat64(m.Percentage),
		FragranceID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.FragranceID))),
		RecipeID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.RecipeID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeFragranceUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeFragranceUpdateInput,
) models.M {
	model := RecipeFragranceUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "percentage":
			modelM[models.RecipeFragranceColumns.Percentage] = model.Percentage
		case "fragranceId":
			modelM[models.RecipeFragranceColumns.FragranceID] = model.FragranceID
		case "recipeId":
			modelM[models.RecipeFragranceColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeFragranceColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeFragranceColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeFragranceColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeFragranceUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "percentage":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.Percentage)
		case "fragranceId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.FragranceID)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeFragranceColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeLipidCreateInputsToBoiler(am []*graphql_models.RecipeLipidCreateInput) []*models.RecipeLipid {
	ar := make([]*models.RecipeLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeLipidCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeLipidCreateInputToBoiler(
	m *graphql_models.RecipeLipidCreateInput,
) *models.RecipeLipid {
	if m == nil {
		return nil
	}

	r := &models.RecipeLipid{
		Percentage:	m.Percentage,
		LipidID:	int(boilergql.IDToBoiler(m.LipidID)),
		RecipeID:	int(boilergql.IDToBoiler(m.RecipeID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeLipidCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeLipidCreateInput,
) models.M {
	model := RecipeLipidCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "percentage":
			modelM[models.RecipeLipidColumns.Percentage] = model.Percentage
		case "lipidId":
			modelM[models.RecipeLipidColumns.LipidID] = model.LipidID
		case "recipeId":
			modelM[models.RecipeLipidColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeLipidColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeLipidColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeLipidColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeLipidCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "percentage":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.Percentage)
		case "lipidId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.LipidID)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeLipidUpdateInputsToBoiler(am []*graphql_models.RecipeLipidUpdateInput) []*models.RecipeLipid {
	ar := make([]*models.RecipeLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeLipidUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeLipidUpdateInputToBoiler(
	m *graphql_models.RecipeLipidUpdateInput,
) *models.RecipeLipid {
	if m == nil {
		return nil
	}

	r := &models.RecipeLipid{
		Percentage:	boilergql.PointerFloat64ToFloat64(m.Percentage),
		LipidID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.LipidID))),
		RecipeID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.RecipeID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeLipidUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeLipidUpdateInput,
) models.M {
	model := RecipeLipidUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "percentage":
			modelM[models.RecipeLipidColumns.Percentage] = model.Percentage
		case "lipidId":
			modelM[models.RecipeLipidColumns.LipidID] = model.LipidID
		case "recipeId":
			modelM[models.RecipeLipidColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeLipidColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeLipidColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeLipidColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeLipidUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "percentage":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.Percentage)
		case "lipidId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.LipidID)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeLipidColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeStepCreateInputsToBoiler(am []*graphql_models.RecipeStepCreateInput) []*models.RecipeStep {
	ar := make([]*models.RecipeStep, len(am))
	for i, m := range am {
		ar[i] = RecipeStepCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeStepCreateInputToBoiler(
	m *graphql_models.RecipeStepCreateInput,
) *models.RecipeStep {
	if m == nil {
		return nil
	}

	r := &models.RecipeStep{
		Num:		m.Num,
		Note:		m.Note,
		RecipeID:	int(boilergql.IDToBoiler(m.RecipeID)),
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeStepCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeStepCreateInput,
) models.M {
	model := RecipeStepCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "num":
			modelM[models.RecipeStepColumns.Num] = model.Num
		case "note":
			modelM[models.RecipeStepColumns.Note] = model.Note
		case "recipeId":
			modelM[models.RecipeStepColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeStepColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeStepColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeStepColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeStepCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "num":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.Num)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.Note)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeStepUpdateInputsToBoiler(am []*graphql_models.RecipeStepUpdateInput) []*models.RecipeStep {
	ar := make([]*models.RecipeStep, len(am))
	for i, m := range am {
		ar[i] = RecipeStepUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeStepUpdateInputToBoiler(
	m *graphql_models.RecipeStepUpdateInput,
) *models.RecipeStep {
	if m == nil {
		return nil
	}

	r := &models.RecipeStep{
		Num:		boilergql.PointerIntToInt(m.Num),
		Note:		boilergql.PointerStringToString(m.Note),
		RecipeID:	int(boilergql.IDToBoiler(boilergql.PointerStringToString(m.RecipeID))),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeStepUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeStepUpdateInput,
) models.M {
	model := RecipeStepUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "num":
			modelM[models.RecipeStepColumns.Num] = model.Num
		case "note":
			modelM[models.RecipeStepColumns.Note] = model.Note
		case "recipeId":
			modelM[models.RecipeStepColumns.RecipeID] = model.RecipeID
		case "createdAt":
			modelM[models.RecipeStepColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeStepColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeStepColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeStepUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "num":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.Num)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.Note)
		case "recipeId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.RecipeID)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeStepColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func RecipeUpdateInputsToBoiler(am []*graphql_models.RecipeUpdateInput) []*models.Recipe {
	ar := make([]*models.Recipe, len(am))
	for i, m := range am {
		ar[i] = RecipeUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func RecipeUpdateInputToBoiler(
	m *graphql_models.RecipeUpdateInput,
) *models.Recipe {
	if m == nil {
		return nil
	}

	r := &models.Recipe{
		Name:		boilergql.PointerStringToString(m.Name),
		Note:		boilergql.PointerStringToString(m.Note),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func RecipeUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.RecipeUpdateInput,
) models.M {
	model := RecipeUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.RecipeColumns.Name] = model.Name
		case "note":
			modelM[models.RecipeColumns.Note] = model.Note
		case "createdAt":
			modelM[models.RecipeColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.RecipeColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.RecipeColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func RecipeUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.Name)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.RecipeColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func SupplierCreateInputsToBoiler(am []*graphql_models.SupplierCreateInput) []*models.Supplier {
	ar := make([]*models.Supplier, len(am))
	for i, m := range am {
		ar[i] = SupplierCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func SupplierCreateInputToBoiler(
	m *graphql_models.SupplierCreateInput,
) *models.Supplier {
	if m == nil {
		return nil
	}

	r := &models.Supplier{
		Name:		m.Name,
		Website:	m.Website,
		Note:		m.Note,
		CreatedAt:	boilergql.IntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.IntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func SupplierCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.SupplierCreateInput,
) models.M {
	model := SupplierCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.SupplierColumns.Name] = model.Name
		case "website":
			modelM[models.SupplierColumns.Website] = model.Website
		case "note":
			modelM[models.SupplierColumns.Note] = model.Note
		case "createdAt":
			modelM[models.SupplierColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.SupplierColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.SupplierColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func SupplierCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.Name)
		case "website":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.Website)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func SupplierUpdateInputsToBoiler(am []*graphql_models.SupplierUpdateInput) []*models.Supplier {
	ar := make([]*models.Supplier, len(am))
	for i, m := range am {
		ar[i] = SupplierUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func SupplierUpdateInputToBoiler(
	m *graphql_models.SupplierUpdateInput,
) *models.Supplier {
	if m == nil {
		return nil
	}

	r := &models.Supplier{
		Name:		boilergql.PointerStringToString(m.Name),
		Website:	boilergql.PointerStringToString(m.Website),
		Note:		boilergql.PointerStringToString(m.Note),
		CreatedAt:	boilergql.PointerIntToTimeDotTime(m.CreatedAt),
		UpdatedAt:	boilergql.PointerIntToTimeDotTime(m.UpdatedAt),
		DeletedAt:	boilergql.PointerIntToNullDotTime(m.DeletedAt),
	}
	return r
}

func SupplierUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.SupplierUpdateInput,
) models.M {
	model := SupplierUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "name":
			modelM[models.SupplierColumns.Name] = model.Name
		case "website":
			modelM[models.SupplierColumns.Website] = model.Website
		case "note":
			modelM[models.SupplierColumns.Note] = model.Note
		case "createdAt":
			modelM[models.SupplierColumns.CreatedAt] = model.CreatedAt
		case "updatedAt":
			modelM[models.SupplierColumns.UpdatedAt] = model.UpdatedAt
		case "deletedAt":
			modelM[models.SupplierColumns.DeletedAt] = model.DeletedAt
		}
	}
	return modelM
}

func SupplierUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "name":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.Name)
		case "website":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.Website)
		case "note":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.Note)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.CreatedAt)
		case "updatedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.UpdatedAt)
		case "deletedAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.SupplierColumns.DeletedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}
