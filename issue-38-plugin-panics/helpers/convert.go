package helpers

import (
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/null/v8"
)

type RecipeFragranceSort string

const (
	RecipeFragranceSortID		RecipeFragranceSort	= "id"
	RecipeFragranceSortPercentage	RecipeFragranceSort	= "percentage"
	RecipeFragranceSortCreatedAt	RecipeFragranceSort	= "createdAt"
	RecipeFragranceSortUpdatedAt	RecipeFragranceSort	= "updatedAt"
	RecipeFragranceSortDeletedAt	RecipeFragranceSort	= "deletedAt"
)

var RecipeFragranceSortDBValue = map[graphql_models.RecipeFragranceSort]RecipeFragranceSort{
	graphql_models.RecipeFragranceSortID:		RecipeFragranceSortID,
	graphql_models.RecipeFragranceSortPercentage:	RecipeFragranceSortPercentage,
	graphql_models.RecipeFragranceSortCreatedAt:	RecipeFragranceSortCreatedAt,
	graphql_models.RecipeFragranceSortUpdatedAt:	RecipeFragranceSortUpdatedAt,
	graphql_models.RecipeFragranceSortDeletedAt:	RecipeFragranceSortDeletedAt,
}
var RecipeFragranceSortAPIValue = map[RecipeFragranceSort]graphql_models.RecipeFragranceSort{
	RecipeFragranceSortID:		graphql_models.RecipeFragranceSortID,
	RecipeFragranceSortPercentage:	graphql_models.RecipeFragranceSortPercentage,
	RecipeFragranceSortCreatedAt:	graphql_models.RecipeFragranceSortCreatedAt,
	RecipeFragranceSortUpdatedAt:	graphql_models.RecipeFragranceSortUpdatedAt,
	RecipeFragranceSortDeletedAt:	graphql_models.RecipeFragranceSortDeletedAt,
}

func NullDotStringToPointerRecipeFragranceSort(v null.String) *graphql_models.RecipeFragranceSort {
	s := StringToRecipeFragranceSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeFragranceSort(v null.String) graphql_models.RecipeFragranceSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeFragranceSort(v.String)
}

func StringToRecipeFragranceSort(v string) graphql_models.RecipeFragranceSort {
	s := RecipeFragranceSortAPIValue[RecipeFragranceSort(v)]
	return s
}

func StringToPointerRecipeFragranceSort(v string) *graphql_models.RecipeFragranceSort {
	s := StringToRecipeFragranceSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeFragranceSortToString(v *graphql_models.RecipeFragranceSort) string {
	if v == nil {
		return ""
	}
	return RecipeFragranceSortToString(*v)
}

func PointerRecipeFragranceSortToNullDotString(v *graphql_models.RecipeFragranceSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeFragranceSortToNullDotString(*v)
}

func RecipeFragranceSortToNullDotString(v graphql_models.RecipeFragranceSort) null.String {
	s := RecipeFragranceSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeFragranceSortToString(v graphql_models.RecipeFragranceSort) string {
	s := RecipeFragranceSortDBValue[v]
	return string(s)
}

type AuthUserGroupSort string

const (
	AuthUserGroupSortID AuthUserGroupSort = "id"
)

var AuthUserGroupSortDBValue = map[graphql_models.AuthUserGroupSort]AuthUserGroupSort{
	graphql_models.AuthUserGroupSortID: AuthUserGroupSortID,
}
var AuthUserGroupSortAPIValue = map[AuthUserGroupSort]graphql_models.AuthUserGroupSort{
	AuthUserGroupSortID: graphql_models.AuthUserGroupSortID,
}

func NullDotStringToPointerAuthUserGroupSort(v null.String) *graphql_models.AuthUserGroupSort {
	s := StringToAuthUserGroupSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAuthUserGroupSort(v null.String) graphql_models.AuthUserGroupSort {
	if !v.Valid {
		return ""
	}
	return StringToAuthUserGroupSort(v.String)
}

func StringToAuthUserGroupSort(v string) graphql_models.AuthUserGroupSort {
	s := AuthUserGroupSortAPIValue[AuthUserGroupSort(v)]
	return s
}

func StringToPointerAuthUserGroupSort(v string) *graphql_models.AuthUserGroupSort {
	s := StringToAuthUserGroupSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAuthUserGroupSortToString(v *graphql_models.AuthUserGroupSort) string {
	if v == nil {
		return ""
	}
	return AuthUserGroupSortToString(*v)
}

func PointerAuthUserGroupSortToNullDotString(v *graphql_models.AuthUserGroupSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AuthUserGroupSortToNullDotString(*v)
}

func AuthUserGroupSortToNullDotString(v graphql_models.AuthUserGroupSort) null.String {
	s := AuthUserGroupSortToString(v)
	return null.NewString(s, s != "")
}

func AuthUserGroupSortToString(v graphql_models.AuthUserGroupSort) string {
	s := AuthUserGroupSortDBValue[v]
	return string(s)
}

type AuthUserSort string

const (
	AuthUserSortID		AuthUserSort	= "id"
	AuthUserSortPassword	AuthUserSort	= "password"
	AuthUserSortLastLogin	AuthUserSort	= "lastLogin"
	AuthUserSortIsSuperuser	AuthUserSort	= "isSuperuser"
	AuthUserSortUsername	AuthUserSort	= "username"
	AuthUserSortFirstName	AuthUserSort	= "firstName"
	AuthUserSortLastName	AuthUserSort	= "lastName"
	AuthUserSortEmail	AuthUserSort	= "email"
	AuthUserSortIsStaff	AuthUserSort	= "isStaff"
	AuthUserSortIsActive	AuthUserSort	= "isActive"
	AuthUserSortDateJoined	AuthUserSort	= "dateJoined"
)

var AuthUserSortDBValue = map[graphql_models.AuthUserSort]AuthUserSort{
	graphql_models.AuthUserSortID:		AuthUserSortID,
	graphql_models.AuthUserSortPassword:	AuthUserSortPassword,
	graphql_models.AuthUserSortLastLogin:	AuthUserSortLastLogin,
	graphql_models.AuthUserSortIsSuperuser:	AuthUserSortIsSuperuser,
	graphql_models.AuthUserSortUsername:	AuthUserSortUsername,
	graphql_models.AuthUserSortFirstName:	AuthUserSortFirstName,
	graphql_models.AuthUserSortLastName:	AuthUserSortLastName,
	graphql_models.AuthUserSortEmail:	AuthUserSortEmail,
	graphql_models.AuthUserSortIsStaff:	AuthUserSortIsStaff,
	graphql_models.AuthUserSortIsActive:	AuthUserSortIsActive,
	graphql_models.AuthUserSortDateJoined:	AuthUserSortDateJoined,
}
var AuthUserSortAPIValue = map[AuthUserSort]graphql_models.AuthUserSort{
	AuthUserSortID:			graphql_models.AuthUserSortID,
	AuthUserSortPassword:		graphql_models.AuthUserSortPassword,
	AuthUserSortLastLogin:		graphql_models.AuthUserSortLastLogin,
	AuthUserSortIsSuperuser:	graphql_models.AuthUserSortIsSuperuser,
	AuthUserSortUsername:		graphql_models.AuthUserSortUsername,
	AuthUserSortFirstName:		graphql_models.AuthUserSortFirstName,
	AuthUserSortLastName:		graphql_models.AuthUserSortLastName,
	AuthUserSortEmail:		graphql_models.AuthUserSortEmail,
	AuthUserSortIsStaff:		graphql_models.AuthUserSortIsStaff,
	AuthUserSortIsActive:		graphql_models.AuthUserSortIsActive,
	AuthUserSortDateJoined:		graphql_models.AuthUserSortDateJoined,
}

func NullDotStringToPointerAuthUserSort(v null.String) *graphql_models.AuthUserSort {
	s := StringToAuthUserSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAuthUserSort(v null.String) graphql_models.AuthUserSort {
	if !v.Valid {
		return ""
	}
	return StringToAuthUserSort(v.String)
}

func StringToAuthUserSort(v string) graphql_models.AuthUserSort {
	s := AuthUserSortAPIValue[AuthUserSort(v)]
	return s
}

func StringToPointerAuthUserSort(v string) *graphql_models.AuthUserSort {
	s := StringToAuthUserSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAuthUserSortToString(v *graphql_models.AuthUserSort) string {
	if v == nil {
		return ""
	}
	return AuthUserSortToString(*v)
}

func PointerAuthUserSortToNullDotString(v *graphql_models.AuthUserSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AuthUserSortToNullDotString(*v)
}

func AuthUserSortToNullDotString(v graphql_models.AuthUserSort) null.String {
	s := AuthUserSortToString(v)
	return null.NewString(s, s != "")
}

func AuthUserSortToString(v graphql_models.AuthUserSort) string {
	s := AuthUserSortDBValue[v]
	return string(s)
}

type RecipeBatchAdditiveSort string

const (
	RecipeBatchAdditiveSortID		RecipeBatchAdditiveSort	= "id"
	RecipeBatchAdditiveSortWeight		RecipeBatchAdditiveSort	= "weight"
	RecipeBatchAdditiveSortCost		RecipeBatchAdditiveSort	= "cost"
	RecipeBatchAdditiveSortCreatedAt	RecipeBatchAdditiveSort	= "createdAt"
	RecipeBatchAdditiveSortUpdatedAt	RecipeBatchAdditiveSort	= "updatedAt"
	RecipeBatchAdditiveSortDeletedAt	RecipeBatchAdditiveSort	= "deletedAt"
)

var RecipeBatchAdditiveSortDBValue = map[graphql_models.RecipeBatchAdditiveSort]RecipeBatchAdditiveSort{
	graphql_models.RecipeBatchAdditiveSortID:		RecipeBatchAdditiveSortID,
	graphql_models.RecipeBatchAdditiveSortWeight:		RecipeBatchAdditiveSortWeight,
	graphql_models.RecipeBatchAdditiveSortCost:		RecipeBatchAdditiveSortCost,
	graphql_models.RecipeBatchAdditiveSortCreatedAt:	RecipeBatchAdditiveSortCreatedAt,
	graphql_models.RecipeBatchAdditiveSortUpdatedAt:	RecipeBatchAdditiveSortUpdatedAt,
	graphql_models.RecipeBatchAdditiveSortDeletedAt:	RecipeBatchAdditiveSortDeletedAt,
}
var RecipeBatchAdditiveSortAPIValue = map[RecipeBatchAdditiveSort]graphql_models.RecipeBatchAdditiveSort{
	RecipeBatchAdditiveSortID:		graphql_models.RecipeBatchAdditiveSortID,
	RecipeBatchAdditiveSortWeight:		graphql_models.RecipeBatchAdditiveSortWeight,
	RecipeBatchAdditiveSortCost:		graphql_models.RecipeBatchAdditiveSortCost,
	RecipeBatchAdditiveSortCreatedAt:	graphql_models.RecipeBatchAdditiveSortCreatedAt,
	RecipeBatchAdditiveSortUpdatedAt:	graphql_models.RecipeBatchAdditiveSortUpdatedAt,
	RecipeBatchAdditiveSortDeletedAt:	graphql_models.RecipeBatchAdditiveSortDeletedAt,
}

func NullDotStringToPointerRecipeBatchAdditiveSort(v null.String) *graphql_models.RecipeBatchAdditiveSort {
	s := StringToRecipeBatchAdditiveSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeBatchAdditiveSort(v null.String) graphql_models.RecipeBatchAdditiveSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeBatchAdditiveSort(v.String)
}

func StringToRecipeBatchAdditiveSort(v string) graphql_models.RecipeBatchAdditiveSort {
	s := RecipeBatchAdditiveSortAPIValue[RecipeBatchAdditiveSort(v)]
	return s
}

func StringToPointerRecipeBatchAdditiveSort(v string) *graphql_models.RecipeBatchAdditiveSort {
	s := StringToRecipeBatchAdditiveSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeBatchAdditiveSortToString(v *graphql_models.RecipeBatchAdditiveSort) string {
	if v == nil {
		return ""
	}
	return RecipeBatchAdditiveSortToString(*v)
}

func PointerRecipeBatchAdditiveSortToNullDotString(v *graphql_models.RecipeBatchAdditiveSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeBatchAdditiveSortToNullDotString(*v)
}

func RecipeBatchAdditiveSortToNullDotString(v graphql_models.RecipeBatchAdditiveSort) null.String {
	s := RecipeBatchAdditiveSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeBatchAdditiveSortToString(v graphql_models.RecipeBatchAdditiveSort) string {
	s := RecipeBatchAdditiveSortDBValue[v]
	return string(s)
}

type RecipeBatchNoteSort string

const (
	RecipeBatchNoteSortID		RecipeBatchNoteSort	= "id"
	RecipeBatchNoteSortNote		RecipeBatchNoteSort	= "note"
	RecipeBatchNoteSortLink		RecipeBatchNoteSort	= "link"
	RecipeBatchNoteSortCreatedAt	RecipeBatchNoteSort	= "createdAt"
	RecipeBatchNoteSortUpdatedAt	RecipeBatchNoteSort	= "updatedAt"
	RecipeBatchNoteSortDeletedAt	RecipeBatchNoteSort	= "deletedAt"
)

var RecipeBatchNoteSortDBValue = map[graphql_models.RecipeBatchNoteSort]RecipeBatchNoteSort{
	graphql_models.RecipeBatchNoteSortID:		RecipeBatchNoteSortID,
	graphql_models.RecipeBatchNoteSortNote:		RecipeBatchNoteSortNote,
	graphql_models.RecipeBatchNoteSortLink:		RecipeBatchNoteSortLink,
	graphql_models.RecipeBatchNoteSortCreatedAt:	RecipeBatchNoteSortCreatedAt,
	graphql_models.RecipeBatchNoteSortUpdatedAt:	RecipeBatchNoteSortUpdatedAt,
	graphql_models.RecipeBatchNoteSortDeletedAt:	RecipeBatchNoteSortDeletedAt,
}
var RecipeBatchNoteSortAPIValue = map[RecipeBatchNoteSort]graphql_models.RecipeBatchNoteSort{
	RecipeBatchNoteSortID:		graphql_models.RecipeBatchNoteSortID,
	RecipeBatchNoteSortNote:	graphql_models.RecipeBatchNoteSortNote,
	RecipeBatchNoteSortLink:	graphql_models.RecipeBatchNoteSortLink,
	RecipeBatchNoteSortCreatedAt:	graphql_models.RecipeBatchNoteSortCreatedAt,
	RecipeBatchNoteSortUpdatedAt:	graphql_models.RecipeBatchNoteSortUpdatedAt,
	RecipeBatchNoteSortDeletedAt:	graphql_models.RecipeBatchNoteSortDeletedAt,
}

func NullDotStringToPointerRecipeBatchNoteSort(v null.String) *graphql_models.RecipeBatchNoteSort {
	s := StringToRecipeBatchNoteSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeBatchNoteSort(v null.String) graphql_models.RecipeBatchNoteSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeBatchNoteSort(v.String)
}

func StringToRecipeBatchNoteSort(v string) graphql_models.RecipeBatchNoteSort {
	s := RecipeBatchNoteSortAPIValue[RecipeBatchNoteSort(v)]
	return s
}

func StringToPointerRecipeBatchNoteSort(v string) *graphql_models.RecipeBatchNoteSort {
	s := StringToRecipeBatchNoteSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeBatchNoteSortToString(v *graphql_models.RecipeBatchNoteSort) string {
	if v == nil {
		return ""
	}
	return RecipeBatchNoteSortToString(*v)
}

func PointerRecipeBatchNoteSortToNullDotString(v *graphql_models.RecipeBatchNoteSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeBatchNoteSortToNullDotString(*v)
}

func RecipeBatchNoteSortToNullDotString(v graphql_models.RecipeBatchNoteSort) null.String {
	s := RecipeBatchNoteSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeBatchNoteSortToString(v graphql_models.RecipeBatchNoteSort) string {
	s := RecipeBatchNoteSortDBValue[v]
	return string(s)
}

type RecipeBatchSort string

const (
	RecipeBatchSortID		RecipeBatchSort	= "id"
	RecipeBatchSortTag		RecipeBatchSort	= "tag"
	RecipeBatchSortProductionDate	RecipeBatchSort	= "productionDate"
	RecipeBatchSortSellableDate	RecipeBatchSort	= "sellableDate"
	RecipeBatchSortNote		RecipeBatchSort	= "note"
	RecipeBatchSortLipidWeight	RecipeBatchSort	= "lipidWeight"
	RecipeBatchSortProductionWeight	RecipeBatchSort	= "productionWeight"
	RecipeBatchSortCuredWeight	RecipeBatchSort	= "curedWeight"
	RecipeBatchSortCreatedAt	RecipeBatchSort	= "createdAt"
	RecipeBatchSortUpdatedAt	RecipeBatchSort	= "updatedAt"
	RecipeBatchSortDeletedAt	RecipeBatchSort	= "deletedAt"
)

var RecipeBatchSortDBValue = map[graphql_models.RecipeBatchSort]RecipeBatchSort{
	graphql_models.RecipeBatchSortID:		RecipeBatchSortID,
	graphql_models.RecipeBatchSortTag:		RecipeBatchSortTag,
	graphql_models.RecipeBatchSortProductionDate:	RecipeBatchSortProductionDate,
	graphql_models.RecipeBatchSortSellableDate:	RecipeBatchSortSellableDate,
	graphql_models.RecipeBatchSortNote:		RecipeBatchSortNote,
	graphql_models.RecipeBatchSortLipidWeight:	RecipeBatchSortLipidWeight,
	graphql_models.RecipeBatchSortProductionWeight:	RecipeBatchSortProductionWeight,
	graphql_models.RecipeBatchSortCuredWeight:	RecipeBatchSortCuredWeight,
	graphql_models.RecipeBatchSortCreatedAt:	RecipeBatchSortCreatedAt,
	graphql_models.RecipeBatchSortUpdatedAt:	RecipeBatchSortUpdatedAt,
	graphql_models.RecipeBatchSortDeletedAt:	RecipeBatchSortDeletedAt,
}
var RecipeBatchSortAPIValue = map[RecipeBatchSort]graphql_models.RecipeBatchSort{
	RecipeBatchSortID:			graphql_models.RecipeBatchSortID,
	RecipeBatchSortTag:			graphql_models.RecipeBatchSortTag,
	RecipeBatchSortProductionDate:		graphql_models.RecipeBatchSortProductionDate,
	RecipeBatchSortSellableDate:		graphql_models.RecipeBatchSortSellableDate,
	RecipeBatchSortNote:			graphql_models.RecipeBatchSortNote,
	RecipeBatchSortLipidWeight:		graphql_models.RecipeBatchSortLipidWeight,
	RecipeBatchSortProductionWeight:	graphql_models.RecipeBatchSortProductionWeight,
	RecipeBatchSortCuredWeight:		graphql_models.RecipeBatchSortCuredWeight,
	RecipeBatchSortCreatedAt:		graphql_models.RecipeBatchSortCreatedAt,
	RecipeBatchSortUpdatedAt:		graphql_models.RecipeBatchSortUpdatedAt,
	RecipeBatchSortDeletedAt:		graphql_models.RecipeBatchSortDeletedAt,
}

func NullDotStringToPointerRecipeBatchSort(v null.String) *graphql_models.RecipeBatchSort {
	s := StringToRecipeBatchSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeBatchSort(v null.String) graphql_models.RecipeBatchSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeBatchSort(v.String)
}

func StringToRecipeBatchSort(v string) graphql_models.RecipeBatchSort {
	s := RecipeBatchSortAPIValue[RecipeBatchSort(v)]
	return s
}

func StringToPointerRecipeBatchSort(v string) *graphql_models.RecipeBatchSort {
	s := StringToRecipeBatchSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeBatchSortToString(v *graphql_models.RecipeBatchSort) string {
	if v == nil {
		return ""
	}
	return RecipeBatchSortToString(*v)
}

func PointerRecipeBatchSortToNullDotString(v *graphql_models.RecipeBatchSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeBatchSortToNullDotString(*v)
}

func RecipeBatchSortToNullDotString(v graphql_models.RecipeBatchSort) null.String {
	s := RecipeBatchSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeBatchSortToString(v graphql_models.RecipeBatchSort) string {
	s := RecipeBatchSortDBValue[v]
	return string(s)
}

type SupplierSort string

const (
	SupplierSortID		SupplierSort	= "id"
	SupplierSortName	SupplierSort	= "name"
	SupplierSortWebsite	SupplierSort	= "website"
	SupplierSortNote	SupplierSort	= "note"
	SupplierSortCreatedAt	SupplierSort	= "createdAt"
	SupplierSortUpdatedAt	SupplierSort	= "updatedAt"
	SupplierSortDeletedAt	SupplierSort	= "deletedAt"
)

var SupplierSortDBValue = map[graphql_models.SupplierSort]SupplierSort{
	graphql_models.SupplierSortID:		SupplierSortID,
	graphql_models.SupplierSortName:	SupplierSortName,
	graphql_models.SupplierSortWebsite:	SupplierSortWebsite,
	graphql_models.SupplierSortNote:	SupplierSortNote,
	graphql_models.SupplierSortCreatedAt:	SupplierSortCreatedAt,
	graphql_models.SupplierSortUpdatedAt:	SupplierSortUpdatedAt,
	graphql_models.SupplierSortDeletedAt:	SupplierSortDeletedAt,
}
var SupplierSortAPIValue = map[SupplierSort]graphql_models.SupplierSort{
	SupplierSortID:		graphql_models.SupplierSortID,
	SupplierSortName:	graphql_models.SupplierSortName,
	SupplierSortWebsite:	graphql_models.SupplierSortWebsite,
	SupplierSortNote:	graphql_models.SupplierSortNote,
	SupplierSortCreatedAt:	graphql_models.SupplierSortCreatedAt,
	SupplierSortUpdatedAt:	graphql_models.SupplierSortUpdatedAt,
	SupplierSortDeletedAt:	graphql_models.SupplierSortDeletedAt,
}

func NullDotStringToPointerSupplierSort(v null.String) *graphql_models.SupplierSort {
	s := StringToSupplierSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToSupplierSort(v null.String) graphql_models.SupplierSort {
	if !v.Valid {
		return ""
	}
	return StringToSupplierSort(v.String)
}

func StringToSupplierSort(v string) graphql_models.SupplierSort {
	s := SupplierSortAPIValue[SupplierSort(v)]
	return s
}

func StringToPointerSupplierSort(v string) *graphql_models.SupplierSort {
	s := StringToSupplierSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerSupplierSortToString(v *graphql_models.SupplierSort) string {
	if v == nil {
		return ""
	}
	return SupplierSortToString(*v)
}

func PointerSupplierSortToNullDotString(v *graphql_models.SupplierSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return SupplierSortToNullDotString(*v)
}

func SupplierSortToNullDotString(v graphql_models.SupplierSort) null.String {
	s := SupplierSortToString(v)
	return null.NewString(s, s != "")
}

func SupplierSortToString(v graphql_models.SupplierSort) string {
	s := SupplierSortDBValue[v]
	return string(s)
}

type AuthGroupSort string

const (
	AuthGroupSortID		AuthGroupSort	= "id"
	AuthGroupSortName	AuthGroupSort	= "name"
)

var AuthGroupSortDBValue = map[graphql_models.AuthGroupSort]AuthGroupSort{
	graphql_models.AuthGroupSortID:		AuthGroupSortID,
	graphql_models.AuthGroupSortName:	AuthGroupSortName,
}
var AuthGroupSortAPIValue = map[AuthGroupSort]graphql_models.AuthGroupSort{
	AuthGroupSortID:	graphql_models.AuthGroupSortID,
	AuthGroupSortName:	graphql_models.AuthGroupSortName,
}

func NullDotStringToPointerAuthGroupSort(v null.String) *graphql_models.AuthGroupSort {
	s := StringToAuthGroupSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAuthGroupSort(v null.String) graphql_models.AuthGroupSort {
	if !v.Valid {
		return ""
	}
	return StringToAuthGroupSort(v.String)
}

func StringToAuthGroupSort(v string) graphql_models.AuthGroupSort {
	s := AuthGroupSortAPIValue[AuthGroupSort(v)]
	return s
}

func StringToPointerAuthGroupSort(v string) *graphql_models.AuthGroupSort {
	s := StringToAuthGroupSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAuthGroupSortToString(v *graphql_models.AuthGroupSort) string {
	if v == nil {
		return ""
	}
	return AuthGroupSortToString(*v)
}

func PointerAuthGroupSortToNullDotString(v *graphql_models.AuthGroupSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AuthGroupSortToNullDotString(*v)
}

func AuthGroupSortToNullDotString(v graphql_models.AuthGroupSort) null.String {
	s := AuthGroupSortToString(v)
	return null.NewString(s, s != "")
}

func AuthGroupSortToString(v graphql_models.AuthGroupSort) string {
	s := AuthGroupSortDBValue[v]
	return string(s)
}

type LipidInventorySort string

const (
	LipidInventorySortID		LipidInventorySort	= "id"
	LipidInventorySortPurchaseDate	LipidInventorySort	= "purchaseDate"
	LipidInventorySortExpiryDate	LipidInventorySort	= "expiryDate"
	LipidInventorySortCost		LipidInventorySort	= "cost"
	LipidInventorySortWeight	LipidInventorySort	= "weight"
	LipidInventorySortSap		LipidInventorySort	= "sap"
	LipidInventorySortNaoh		LipidInventorySort	= "naoh"
	LipidInventorySortKoh		LipidInventorySort	= "koh"
	LipidInventorySortGramsPerLiter	LipidInventorySort	= "gramsPerLiter"
	LipidInventorySortCreatedAt	LipidInventorySort	= "createdAt"
	LipidInventorySortUpdatedAt	LipidInventorySort	= "updatedAt"
	LipidInventorySortDeletedAt	LipidInventorySort	= "deletedAt"
)

var LipidInventorySortDBValue = map[graphql_models.LipidInventorySort]LipidInventorySort{
	graphql_models.LipidInventorySortID:		LipidInventorySortID,
	graphql_models.LipidInventorySortPurchaseDate:	LipidInventorySortPurchaseDate,
	graphql_models.LipidInventorySortExpiryDate:	LipidInventorySortExpiryDate,
	graphql_models.LipidInventorySortCost:		LipidInventorySortCost,
	graphql_models.LipidInventorySortWeight:	LipidInventorySortWeight,
	graphql_models.LipidInventorySortSap:		LipidInventorySortSap,
	graphql_models.LipidInventorySortNaoh:		LipidInventorySortNaoh,
	graphql_models.LipidInventorySortKoh:		LipidInventorySortKoh,
	graphql_models.LipidInventorySortGramsPerLiter:	LipidInventorySortGramsPerLiter,
	graphql_models.LipidInventorySortCreatedAt:	LipidInventorySortCreatedAt,
	graphql_models.LipidInventorySortUpdatedAt:	LipidInventorySortUpdatedAt,
	graphql_models.LipidInventorySortDeletedAt:	LipidInventorySortDeletedAt,
}
var LipidInventorySortAPIValue = map[LipidInventorySort]graphql_models.LipidInventorySort{
	LipidInventorySortID:			graphql_models.LipidInventorySortID,
	LipidInventorySortPurchaseDate:		graphql_models.LipidInventorySortPurchaseDate,
	LipidInventorySortExpiryDate:		graphql_models.LipidInventorySortExpiryDate,
	LipidInventorySortCost:			graphql_models.LipidInventorySortCost,
	LipidInventorySortWeight:		graphql_models.LipidInventorySortWeight,
	LipidInventorySortSap:			graphql_models.LipidInventorySortSap,
	LipidInventorySortNaoh:			graphql_models.LipidInventorySortNaoh,
	LipidInventorySortKoh:			graphql_models.LipidInventorySortKoh,
	LipidInventorySortGramsPerLiter:	graphql_models.LipidInventorySortGramsPerLiter,
	LipidInventorySortCreatedAt:		graphql_models.LipidInventorySortCreatedAt,
	LipidInventorySortUpdatedAt:		graphql_models.LipidInventorySortUpdatedAt,
	LipidInventorySortDeletedAt:		graphql_models.LipidInventorySortDeletedAt,
}

func NullDotStringToPointerLipidInventorySort(v null.String) *graphql_models.LipidInventorySort {
	s := StringToLipidInventorySort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToLipidInventorySort(v null.String) graphql_models.LipidInventorySort {
	if !v.Valid {
		return ""
	}
	return StringToLipidInventorySort(v.String)
}

func StringToLipidInventorySort(v string) graphql_models.LipidInventorySort {
	s := LipidInventorySortAPIValue[LipidInventorySort(v)]
	return s
}

func StringToPointerLipidInventorySort(v string) *graphql_models.LipidInventorySort {
	s := StringToLipidInventorySort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerLipidInventorySortToString(v *graphql_models.LipidInventorySort) string {
	if v == nil {
		return ""
	}
	return LipidInventorySortToString(*v)
}

func PointerLipidInventorySortToNullDotString(v *graphql_models.LipidInventorySort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return LipidInventorySortToNullDotString(*v)
}

func LipidInventorySortToNullDotString(v graphql_models.LipidInventorySort) null.String {
	s := LipidInventorySortToString(v)
	return null.NewString(s, s != "")
}

func LipidInventorySortToString(v graphql_models.LipidInventorySort) string {
	s := LipidInventorySortDBValue[v]
	return string(s)
}

type LyeInventorySort string

const (
	LyeInventorySortID		LyeInventorySort	= "id"
	LyeInventorySortPurchaseDate	LyeInventorySort	= "purchaseDate"
	LyeInventorySortExpiryDate	LyeInventorySort	= "expiryDate"
	LyeInventorySortCost		LyeInventorySort	= "cost"
	LyeInventorySortWeight		LyeInventorySort	= "weight"
	LyeInventorySortConcentration	LyeInventorySort	= "concentration"
	LyeInventorySortCreatedAt	LyeInventorySort	= "createdAt"
	LyeInventorySortUpdatedAt	LyeInventorySort	= "updatedAt"
	LyeInventorySortDeletedAt	LyeInventorySort	= "deletedAt"
)

var LyeInventorySortDBValue = map[graphql_models.LyeInventorySort]LyeInventorySort{
	graphql_models.LyeInventorySortID:		LyeInventorySortID,
	graphql_models.LyeInventorySortPurchaseDate:	LyeInventorySortPurchaseDate,
	graphql_models.LyeInventorySortExpiryDate:	LyeInventorySortExpiryDate,
	graphql_models.LyeInventorySortCost:		LyeInventorySortCost,
	graphql_models.LyeInventorySortWeight:		LyeInventorySortWeight,
	graphql_models.LyeInventorySortConcentration:	LyeInventorySortConcentration,
	graphql_models.LyeInventorySortCreatedAt:	LyeInventorySortCreatedAt,
	graphql_models.LyeInventorySortUpdatedAt:	LyeInventorySortUpdatedAt,
	graphql_models.LyeInventorySortDeletedAt:	LyeInventorySortDeletedAt,
}
var LyeInventorySortAPIValue = map[LyeInventorySort]graphql_models.LyeInventorySort{
	LyeInventorySortID:		graphql_models.LyeInventorySortID,
	LyeInventorySortPurchaseDate:	graphql_models.LyeInventorySortPurchaseDate,
	LyeInventorySortExpiryDate:	graphql_models.LyeInventorySortExpiryDate,
	LyeInventorySortCost:		graphql_models.LyeInventorySortCost,
	LyeInventorySortWeight:		graphql_models.LyeInventorySortWeight,
	LyeInventorySortConcentration:	graphql_models.LyeInventorySortConcentration,
	LyeInventorySortCreatedAt:	graphql_models.LyeInventorySortCreatedAt,
	LyeInventorySortUpdatedAt:	graphql_models.LyeInventorySortUpdatedAt,
	LyeInventorySortDeletedAt:	graphql_models.LyeInventorySortDeletedAt,
}

func NullDotStringToPointerLyeInventorySort(v null.String) *graphql_models.LyeInventorySort {
	s := StringToLyeInventorySort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToLyeInventorySort(v null.String) graphql_models.LyeInventorySort {
	if !v.Valid {
		return ""
	}
	return StringToLyeInventorySort(v.String)
}

func StringToLyeInventorySort(v string) graphql_models.LyeInventorySort {
	s := LyeInventorySortAPIValue[LyeInventorySort(v)]
	return s
}

func StringToPointerLyeInventorySort(v string) *graphql_models.LyeInventorySort {
	s := StringToLyeInventorySort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerLyeInventorySortToString(v *graphql_models.LyeInventorySort) string {
	if v == nil {
		return ""
	}
	return LyeInventorySortToString(*v)
}

func PointerLyeInventorySortToNullDotString(v *graphql_models.LyeInventorySort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return LyeInventorySortToNullDotString(*v)
}

func LyeInventorySortToNullDotString(v graphql_models.LyeInventorySort) null.String {
	s := LyeInventorySortToString(v)
	return null.NewString(s, s != "")
}

func LyeInventorySortToString(v graphql_models.LyeInventorySort) string {
	s := LyeInventorySortDBValue[v]
	return string(s)
}

type RecipeBatchLyeSort string

const (
	RecipeBatchLyeSortID		RecipeBatchLyeSort	= "id"
	RecipeBatchLyeSortWeight	RecipeBatchLyeSort	= "weight"
	RecipeBatchLyeSortDiscount	RecipeBatchLyeSort	= "discount"
	RecipeBatchLyeSortCost		RecipeBatchLyeSort	= "cost"
	RecipeBatchLyeSortCreatedAt	RecipeBatchLyeSort	= "createdAt"
	RecipeBatchLyeSortUpdatedAt	RecipeBatchLyeSort	= "updatedAt"
	RecipeBatchLyeSortDeletedAt	RecipeBatchLyeSort	= "deletedAt"
)

var RecipeBatchLyeSortDBValue = map[graphql_models.RecipeBatchLyeSort]RecipeBatchLyeSort{
	graphql_models.RecipeBatchLyeSortID:		RecipeBatchLyeSortID,
	graphql_models.RecipeBatchLyeSortWeight:	RecipeBatchLyeSortWeight,
	graphql_models.RecipeBatchLyeSortDiscount:	RecipeBatchLyeSortDiscount,
	graphql_models.RecipeBatchLyeSortCost:		RecipeBatchLyeSortCost,
	graphql_models.RecipeBatchLyeSortCreatedAt:	RecipeBatchLyeSortCreatedAt,
	graphql_models.RecipeBatchLyeSortUpdatedAt:	RecipeBatchLyeSortUpdatedAt,
	graphql_models.RecipeBatchLyeSortDeletedAt:	RecipeBatchLyeSortDeletedAt,
}
var RecipeBatchLyeSortAPIValue = map[RecipeBatchLyeSort]graphql_models.RecipeBatchLyeSort{
	RecipeBatchLyeSortID:		graphql_models.RecipeBatchLyeSortID,
	RecipeBatchLyeSortWeight:	graphql_models.RecipeBatchLyeSortWeight,
	RecipeBatchLyeSortDiscount:	graphql_models.RecipeBatchLyeSortDiscount,
	RecipeBatchLyeSortCost:		graphql_models.RecipeBatchLyeSortCost,
	RecipeBatchLyeSortCreatedAt:	graphql_models.RecipeBatchLyeSortCreatedAt,
	RecipeBatchLyeSortUpdatedAt:	graphql_models.RecipeBatchLyeSortUpdatedAt,
	RecipeBatchLyeSortDeletedAt:	graphql_models.RecipeBatchLyeSortDeletedAt,
}

func NullDotStringToPointerRecipeBatchLyeSort(v null.String) *graphql_models.RecipeBatchLyeSort {
	s := StringToRecipeBatchLyeSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeBatchLyeSort(v null.String) graphql_models.RecipeBatchLyeSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeBatchLyeSort(v.String)
}

func StringToRecipeBatchLyeSort(v string) graphql_models.RecipeBatchLyeSort {
	s := RecipeBatchLyeSortAPIValue[RecipeBatchLyeSort(v)]
	return s
}

func StringToPointerRecipeBatchLyeSort(v string) *graphql_models.RecipeBatchLyeSort {
	s := StringToRecipeBatchLyeSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeBatchLyeSortToString(v *graphql_models.RecipeBatchLyeSort) string {
	if v == nil {
		return ""
	}
	return RecipeBatchLyeSortToString(*v)
}

func PointerRecipeBatchLyeSortToNullDotString(v *graphql_models.RecipeBatchLyeSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeBatchLyeSortToNullDotString(*v)
}

func RecipeBatchLyeSortToNullDotString(v graphql_models.RecipeBatchLyeSort) null.String {
	s := RecipeBatchLyeSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeBatchLyeSortToString(v graphql_models.RecipeBatchLyeSort) string {
	s := RecipeBatchLyeSortDBValue[v]
	return string(s)
}

type RecipeAdditiveSort string

const (
	RecipeAdditiveSortID		RecipeAdditiveSort	= "id"
	RecipeAdditiveSortPercentage	RecipeAdditiveSort	= "percentage"
	RecipeAdditiveSortCreatedAt	RecipeAdditiveSort	= "createdAt"
	RecipeAdditiveSortUpdatedAt	RecipeAdditiveSort	= "updatedAt"
	RecipeAdditiveSortDeletedAt	RecipeAdditiveSort	= "deletedAt"
)

var RecipeAdditiveSortDBValue = map[graphql_models.RecipeAdditiveSort]RecipeAdditiveSort{
	graphql_models.RecipeAdditiveSortID:		RecipeAdditiveSortID,
	graphql_models.RecipeAdditiveSortPercentage:	RecipeAdditiveSortPercentage,
	graphql_models.RecipeAdditiveSortCreatedAt:	RecipeAdditiveSortCreatedAt,
	graphql_models.RecipeAdditiveSortUpdatedAt:	RecipeAdditiveSortUpdatedAt,
	graphql_models.RecipeAdditiveSortDeletedAt:	RecipeAdditiveSortDeletedAt,
}
var RecipeAdditiveSortAPIValue = map[RecipeAdditiveSort]graphql_models.RecipeAdditiveSort{
	RecipeAdditiveSortID:		graphql_models.RecipeAdditiveSortID,
	RecipeAdditiveSortPercentage:	graphql_models.RecipeAdditiveSortPercentage,
	RecipeAdditiveSortCreatedAt:	graphql_models.RecipeAdditiveSortCreatedAt,
	RecipeAdditiveSortUpdatedAt:	graphql_models.RecipeAdditiveSortUpdatedAt,
	RecipeAdditiveSortDeletedAt:	graphql_models.RecipeAdditiveSortDeletedAt,
}

func NullDotStringToPointerRecipeAdditiveSort(v null.String) *graphql_models.RecipeAdditiveSort {
	s := StringToRecipeAdditiveSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeAdditiveSort(v null.String) graphql_models.RecipeAdditiveSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeAdditiveSort(v.String)
}

func StringToRecipeAdditiveSort(v string) graphql_models.RecipeAdditiveSort {
	s := RecipeAdditiveSortAPIValue[RecipeAdditiveSort(v)]
	return s
}

func StringToPointerRecipeAdditiveSort(v string) *graphql_models.RecipeAdditiveSort {
	s := StringToRecipeAdditiveSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeAdditiveSortToString(v *graphql_models.RecipeAdditiveSort) string {
	if v == nil {
		return ""
	}
	return RecipeAdditiveSortToString(*v)
}

func PointerRecipeAdditiveSortToNullDotString(v *graphql_models.RecipeAdditiveSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeAdditiveSortToNullDotString(*v)
}

func RecipeAdditiveSortToNullDotString(v graphql_models.RecipeAdditiveSort) null.String {
	s := RecipeAdditiveSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeAdditiveSortToString(v graphql_models.RecipeAdditiveSort) string {
	s := RecipeAdditiveSortDBValue[v]
	return string(s)
}

type AuthPermissionSort string

const (
	AuthPermissionSortID		AuthPermissionSort	= "id"
	AuthPermissionSortName		AuthPermissionSort	= "name"
	AuthPermissionSortCodename	AuthPermissionSort	= "codename"
)

var AuthPermissionSortDBValue = map[graphql_models.AuthPermissionSort]AuthPermissionSort{
	graphql_models.AuthPermissionSortID:		AuthPermissionSortID,
	graphql_models.AuthPermissionSortName:		AuthPermissionSortName,
	graphql_models.AuthPermissionSortCodename:	AuthPermissionSortCodename,
}
var AuthPermissionSortAPIValue = map[AuthPermissionSort]graphql_models.AuthPermissionSort{
	AuthPermissionSortID:		graphql_models.AuthPermissionSortID,
	AuthPermissionSortName:		graphql_models.AuthPermissionSortName,
	AuthPermissionSortCodename:	graphql_models.AuthPermissionSortCodename,
}

func NullDotStringToPointerAuthPermissionSort(v null.String) *graphql_models.AuthPermissionSort {
	s := StringToAuthPermissionSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAuthPermissionSort(v null.String) graphql_models.AuthPermissionSort {
	if !v.Valid {
		return ""
	}
	return StringToAuthPermissionSort(v.String)
}

func StringToAuthPermissionSort(v string) graphql_models.AuthPermissionSort {
	s := AuthPermissionSortAPIValue[AuthPermissionSort(v)]
	return s
}

func StringToPointerAuthPermissionSort(v string) *graphql_models.AuthPermissionSort {
	s := StringToAuthPermissionSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAuthPermissionSortToString(v *graphql_models.AuthPermissionSort) string {
	if v == nil {
		return ""
	}
	return AuthPermissionSortToString(*v)
}

func PointerAuthPermissionSortToNullDotString(v *graphql_models.AuthPermissionSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AuthPermissionSortToNullDotString(*v)
}

func AuthPermissionSortToNullDotString(v graphql_models.AuthPermissionSort) null.String {
	s := AuthPermissionSortToString(v)
	return null.NewString(s, s != "")
}

func AuthPermissionSortToString(v graphql_models.AuthPermissionSort) string {
	s := AuthPermissionSortDBValue[v]
	return string(s)
}

type RecipeBatchLipidSort string

const (
	RecipeBatchLipidSortID		RecipeBatchLipidSort	= "id"
	RecipeBatchLipidSortWeight	RecipeBatchLipidSort	= "weight"
	RecipeBatchLipidSortCost	RecipeBatchLipidSort	= "cost"
	RecipeBatchLipidSortCreatedAt	RecipeBatchLipidSort	= "createdAt"
	RecipeBatchLipidSortUpdatedAt	RecipeBatchLipidSort	= "updatedAt"
	RecipeBatchLipidSortDeletedAt	RecipeBatchLipidSort	= "deletedAt"
)

var RecipeBatchLipidSortDBValue = map[graphql_models.RecipeBatchLipidSort]RecipeBatchLipidSort{
	graphql_models.RecipeBatchLipidSortID:		RecipeBatchLipidSortID,
	graphql_models.RecipeBatchLipidSortWeight:	RecipeBatchLipidSortWeight,
	graphql_models.RecipeBatchLipidSortCost:	RecipeBatchLipidSortCost,
	graphql_models.RecipeBatchLipidSortCreatedAt:	RecipeBatchLipidSortCreatedAt,
	graphql_models.RecipeBatchLipidSortUpdatedAt:	RecipeBatchLipidSortUpdatedAt,
	graphql_models.RecipeBatchLipidSortDeletedAt:	RecipeBatchLipidSortDeletedAt,
}
var RecipeBatchLipidSortAPIValue = map[RecipeBatchLipidSort]graphql_models.RecipeBatchLipidSort{
	RecipeBatchLipidSortID:		graphql_models.RecipeBatchLipidSortID,
	RecipeBatchLipidSortWeight:	graphql_models.RecipeBatchLipidSortWeight,
	RecipeBatchLipidSortCost:	graphql_models.RecipeBatchLipidSortCost,
	RecipeBatchLipidSortCreatedAt:	graphql_models.RecipeBatchLipidSortCreatedAt,
	RecipeBatchLipidSortUpdatedAt:	graphql_models.RecipeBatchLipidSortUpdatedAt,
	RecipeBatchLipidSortDeletedAt:	graphql_models.RecipeBatchLipidSortDeletedAt,
}

func NullDotStringToPointerRecipeBatchLipidSort(v null.String) *graphql_models.RecipeBatchLipidSort {
	s := StringToRecipeBatchLipidSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeBatchLipidSort(v null.String) graphql_models.RecipeBatchLipidSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeBatchLipidSort(v.String)
}

func StringToRecipeBatchLipidSort(v string) graphql_models.RecipeBatchLipidSort {
	s := RecipeBatchLipidSortAPIValue[RecipeBatchLipidSort(v)]
	return s
}

func StringToPointerRecipeBatchLipidSort(v string) *graphql_models.RecipeBatchLipidSort {
	s := StringToRecipeBatchLipidSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeBatchLipidSortToString(v *graphql_models.RecipeBatchLipidSort) string {
	if v == nil {
		return ""
	}
	return RecipeBatchLipidSortToString(*v)
}

func PointerRecipeBatchLipidSortToNullDotString(v *graphql_models.RecipeBatchLipidSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeBatchLipidSortToNullDotString(*v)
}

func RecipeBatchLipidSortToNullDotString(v graphql_models.RecipeBatchLipidSort) null.String {
	s := RecipeBatchLipidSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeBatchLipidSortToString(v graphql_models.RecipeBatchLipidSort) string {
	s := RecipeBatchLipidSortDBValue[v]
	return string(s)
}

type AuthUserUserPermissionSort string

const (
	AuthUserUserPermissionSortID AuthUserUserPermissionSort = "id"
)

var AuthUserUserPermissionSortDBValue = map[graphql_models.AuthUserUserPermissionSort]AuthUserUserPermissionSort{
	graphql_models.AuthUserUserPermissionSortID: AuthUserUserPermissionSortID,
}
var AuthUserUserPermissionSortAPIValue = map[AuthUserUserPermissionSort]graphql_models.AuthUserUserPermissionSort{
	AuthUserUserPermissionSortID: graphql_models.AuthUserUserPermissionSortID,
}

func NullDotStringToPointerAuthUserUserPermissionSort(v null.String) *graphql_models.AuthUserUserPermissionSort {
	s := StringToAuthUserUserPermissionSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAuthUserUserPermissionSort(v null.String) graphql_models.AuthUserUserPermissionSort {
	if !v.Valid {
		return ""
	}
	return StringToAuthUserUserPermissionSort(v.String)
}

func StringToAuthUserUserPermissionSort(v string) graphql_models.AuthUserUserPermissionSort {
	s := AuthUserUserPermissionSortAPIValue[AuthUserUserPermissionSort(v)]
	return s
}

func StringToPointerAuthUserUserPermissionSort(v string) *graphql_models.AuthUserUserPermissionSort {
	s := StringToAuthUserUserPermissionSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAuthUserUserPermissionSortToString(v *graphql_models.AuthUserUserPermissionSort) string {
	if v == nil {
		return ""
	}
	return AuthUserUserPermissionSortToString(*v)
}

func PointerAuthUserUserPermissionSortToNullDotString(v *graphql_models.AuthUserUserPermissionSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AuthUserUserPermissionSortToNullDotString(*v)
}

func AuthUserUserPermissionSortToNullDotString(v graphql_models.AuthUserUserPermissionSort) null.String {
	s := AuthUserUserPermissionSortToString(v)
	return null.NewString(s, s != "")
}

func AuthUserUserPermissionSortToString(v graphql_models.AuthUserUserPermissionSort) string {
	s := AuthUserUserPermissionSortDBValue[v]
	return string(s)
}

type RecipeLipidSort string

const (
	RecipeLipidSortID		RecipeLipidSort	= "id"
	RecipeLipidSortPercentage	RecipeLipidSort	= "percentage"
	RecipeLipidSortCreatedAt	RecipeLipidSort	= "createdAt"
	RecipeLipidSortUpdatedAt	RecipeLipidSort	= "updatedAt"
	RecipeLipidSortDeletedAt	RecipeLipidSort	= "deletedAt"
)

var RecipeLipidSortDBValue = map[graphql_models.RecipeLipidSort]RecipeLipidSort{
	graphql_models.RecipeLipidSortID:		RecipeLipidSortID,
	graphql_models.RecipeLipidSortPercentage:	RecipeLipidSortPercentage,
	graphql_models.RecipeLipidSortCreatedAt:	RecipeLipidSortCreatedAt,
	graphql_models.RecipeLipidSortUpdatedAt:	RecipeLipidSortUpdatedAt,
	graphql_models.RecipeLipidSortDeletedAt:	RecipeLipidSortDeletedAt,
}
var RecipeLipidSortAPIValue = map[RecipeLipidSort]graphql_models.RecipeLipidSort{
	RecipeLipidSortID:		graphql_models.RecipeLipidSortID,
	RecipeLipidSortPercentage:	graphql_models.RecipeLipidSortPercentage,
	RecipeLipidSortCreatedAt:	graphql_models.RecipeLipidSortCreatedAt,
	RecipeLipidSortUpdatedAt:	graphql_models.RecipeLipidSortUpdatedAt,
	RecipeLipidSortDeletedAt:	graphql_models.RecipeLipidSortDeletedAt,
}

func NullDotStringToPointerRecipeLipidSort(v null.String) *graphql_models.RecipeLipidSort {
	s := StringToRecipeLipidSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeLipidSort(v null.String) graphql_models.RecipeLipidSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeLipidSort(v.String)
}

func StringToRecipeLipidSort(v string) graphql_models.RecipeLipidSort {
	s := RecipeLipidSortAPIValue[RecipeLipidSort(v)]
	return s
}

func StringToPointerRecipeLipidSort(v string) *graphql_models.RecipeLipidSort {
	s := StringToRecipeLipidSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeLipidSortToString(v *graphql_models.RecipeLipidSort) string {
	if v == nil {
		return ""
	}
	return RecipeLipidSortToString(*v)
}

func PointerRecipeLipidSortToNullDotString(v *graphql_models.RecipeLipidSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeLipidSortToNullDotString(*v)
}

func RecipeLipidSortToNullDotString(v graphql_models.RecipeLipidSort) null.String {
	s := RecipeLipidSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeLipidSortToString(v graphql_models.RecipeLipidSort) string {
	s := RecipeLipidSortDBValue[v]
	return string(s)
}

type AdditiveSort string

const (
	AdditiveSortID		AdditiveSort	= "id"
	AdditiveSortName	AdditiveSort	= "name"
	AdditiveSortNote	AdditiveSort	= "note"
	AdditiveSortCreatedAt	AdditiveSort	= "createdAt"
	AdditiveSortUpdatedAt	AdditiveSort	= "updatedAt"
	AdditiveSortDeletedAt	AdditiveSort	= "deletedAt"
)

var AdditiveSortDBValue = map[graphql_models.AdditiveSort]AdditiveSort{
	graphql_models.AdditiveSortID:		AdditiveSortID,
	graphql_models.AdditiveSortName:	AdditiveSortName,
	graphql_models.AdditiveSortNote:	AdditiveSortNote,
	graphql_models.AdditiveSortCreatedAt:	AdditiveSortCreatedAt,
	graphql_models.AdditiveSortUpdatedAt:	AdditiveSortUpdatedAt,
	graphql_models.AdditiveSortDeletedAt:	AdditiveSortDeletedAt,
}
var AdditiveSortAPIValue = map[AdditiveSort]graphql_models.AdditiveSort{
	AdditiveSortID:		graphql_models.AdditiveSortID,
	AdditiveSortName:	graphql_models.AdditiveSortName,
	AdditiveSortNote:	graphql_models.AdditiveSortNote,
	AdditiveSortCreatedAt:	graphql_models.AdditiveSortCreatedAt,
	AdditiveSortUpdatedAt:	graphql_models.AdditiveSortUpdatedAt,
	AdditiveSortDeletedAt:	graphql_models.AdditiveSortDeletedAt,
}

func NullDotStringToPointerAdditiveSort(v null.String) *graphql_models.AdditiveSort {
	s := StringToAdditiveSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAdditiveSort(v null.String) graphql_models.AdditiveSort {
	if !v.Valid {
		return ""
	}
	return StringToAdditiveSort(v.String)
}

func StringToAdditiveSort(v string) graphql_models.AdditiveSort {
	s := AdditiveSortAPIValue[AdditiveSort(v)]
	return s
}

func StringToPointerAdditiveSort(v string) *graphql_models.AdditiveSort {
	s := StringToAdditiveSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAdditiveSortToString(v *graphql_models.AdditiveSort) string {
	if v == nil {
		return ""
	}
	return AdditiveSortToString(*v)
}

func PointerAdditiveSortToNullDotString(v *graphql_models.AdditiveSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AdditiveSortToNullDotString(*v)
}

func AdditiveSortToNullDotString(v graphql_models.AdditiveSort) null.String {
	s := AdditiveSortToString(v)
	return null.NewString(s, s != "")
}

func AdditiveSortToString(v graphql_models.AdditiveSort) string {
	s := AdditiveSortDBValue[v]
	return string(s)
}

type RecipeSort string

const (
	RecipeSortID		RecipeSort	= "id"
	RecipeSortName		RecipeSort	= "name"
	RecipeSortNote		RecipeSort	= "note"
	RecipeSortCreatedAt	RecipeSort	= "createdAt"
	RecipeSortUpdatedAt	RecipeSort	= "updatedAt"
	RecipeSortDeletedAt	RecipeSort	= "deletedAt"
)

var RecipeSortDBValue = map[graphql_models.RecipeSort]RecipeSort{
	graphql_models.RecipeSortID:		RecipeSortID,
	graphql_models.RecipeSortName:		RecipeSortName,
	graphql_models.RecipeSortNote:		RecipeSortNote,
	graphql_models.RecipeSortCreatedAt:	RecipeSortCreatedAt,
	graphql_models.RecipeSortUpdatedAt:	RecipeSortUpdatedAt,
	graphql_models.RecipeSortDeletedAt:	RecipeSortDeletedAt,
}
var RecipeSortAPIValue = map[RecipeSort]graphql_models.RecipeSort{
	RecipeSortID:		graphql_models.RecipeSortID,
	RecipeSortName:		graphql_models.RecipeSortName,
	RecipeSortNote:		graphql_models.RecipeSortNote,
	RecipeSortCreatedAt:	graphql_models.RecipeSortCreatedAt,
	RecipeSortUpdatedAt:	graphql_models.RecipeSortUpdatedAt,
	RecipeSortDeletedAt:	graphql_models.RecipeSortDeletedAt,
}

func NullDotStringToPointerRecipeSort(v null.String) *graphql_models.RecipeSort {
	s := StringToRecipeSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeSort(v null.String) graphql_models.RecipeSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeSort(v.String)
}

func StringToRecipeSort(v string) graphql_models.RecipeSort {
	s := RecipeSortAPIValue[RecipeSort(v)]
	return s
}

func StringToPointerRecipeSort(v string) *graphql_models.RecipeSort {
	s := StringToRecipeSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeSortToString(v *graphql_models.RecipeSort) string {
	if v == nil {
		return ""
	}
	return RecipeSortToString(*v)
}

func PointerRecipeSortToNullDotString(v *graphql_models.RecipeSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeSortToNullDotString(*v)
}

func RecipeSortToNullDotString(v graphql_models.RecipeSort) null.String {
	s := RecipeSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeSortToString(v graphql_models.RecipeSort) string {
	s := RecipeSortDBValue[v]
	return string(s)
}

type AdditiveInventorySort string

const (
	AdditiveInventorySortID			AdditiveInventorySort	= "id"
	AdditiveInventorySortPurchaseDate	AdditiveInventorySort	= "purchaseDate"
	AdditiveInventorySortExpiryDate		AdditiveInventorySort	= "expiryDate"
	AdditiveInventorySortCost		AdditiveInventorySort	= "cost"
	AdditiveInventorySortWeight		AdditiveInventorySort	= "weight"
	AdditiveInventorySortCreatedAt		AdditiveInventorySort	= "createdAt"
	AdditiveInventorySortUpdatedAt		AdditiveInventorySort	= "updatedAt"
	AdditiveInventorySortDeletedAt		AdditiveInventorySort	= "deletedAt"
)

var AdditiveInventorySortDBValue = map[graphql_models.AdditiveInventorySort]AdditiveInventorySort{
	graphql_models.AdditiveInventorySortID:			AdditiveInventorySortID,
	graphql_models.AdditiveInventorySortPurchaseDate:	AdditiveInventorySortPurchaseDate,
	graphql_models.AdditiveInventorySortExpiryDate:		AdditiveInventorySortExpiryDate,
	graphql_models.AdditiveInventorySortCost:		AdditiveInventorySortCost,
	graphql_models.AdditiveInventorySortWeight:		AdditiveInventorySortWeight,
	graphql_models.AdditiveInventorySortCreatedAt:		AdditiveInventorySortCreatedAt,
	graphql_models.AdditiveInventorySortUpdatedAt:		AdditiveInventorySortUpdatedAt,
	graphql_models.AdditiveInventorySortDeletedAt:		AdditiveInventorySortDeletedAt,
}
var AdditiveInventorySortAPIValue = map[AdditiveInventorySort]graphql_models.AdditiveInventorySort{
	AdditiveInventorySortID:		graphql_models.AdditiveInventorySortID,
	AdditiveInventorySortPurchaseDate:	graphql_models.AdditiveInventorySortPurchaseDate,
	AdditiveInventorySortExpiryDate:	graphql_models.AdditiveInventorySortExpiryDate,
	AdditiveInventorySortCost:		graphql_models.AdditiveInventorySortCost,
	AdditiveInventorySortWeight:		graphql_models.AdditiveInventorySortWeight,
	AdditiveInventorySortCreatedAt:		graphql_models.AdditiveInventorySortCreatedAt,
	AdditiveInventorySortUpdatedAt:		graphql_models.AdditiveInventorySortUpdatedAt,
	AdditiveInventorySortDeletedAt:		graphql_models.AdditiveInventorySortDeletedAt,
}

func NullDotStringToPointerAdditiveInventorySort(v null.String) *graphql_models.AdditiveInventorySort {
	s := StringToAdditiveInventorySort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAdditiveInventorySort(v null.String) graphql_models.AdditiveInventorySort {
	if !v.Valid {
		return ""
	}
	return StringToAdditiveInventorySort(v.String)
}

func StringToAdditiveInventorySort(v string) graphql_models.AdditiveInventorySort {
	s := AdditiveInventorySortAPIValue[AdditiveInventorySort(v)]
	return s
}

func StringToPointerAdditiveInventorySort(v string) *graphql_models.AdditiveInventorySort {
	s := StringToAdditiveInventorySort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAdditiveInventorySortToString(v *graphql_models.AdditiveInventorySort) string {
	if v == nil {
		return ""
	}
	return AdditiveInventorySortToString(*v)
}

func PointerAdditiveInventorySortToNullDotString(v *graphql_models.AdditiveInventorySort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AdditiveInventorySortToNullDotString(*v)
}

func AdditiveInventorySortToNullDotString(v graphql_models.AdditiveInventorySort) null.String {
	s := AdditiveInventorySortToString(v)
	return null.NewString(s, s != "")
}

func AdditiveInventorySortToString(v graphql_models.AdditiveInventorySort) string {
	s := AdditiveInventorySortDBValue[v]
	return string(s)
}

type FragranceInventorySort string

const (
	FragranceInventorySortID		FragranceInventorySort	= "id"
	FragranceInventorySortPurchaseDate	FragranceInventorySort	= "purchaseDate"
	FragranceInventorySortExpiryDate	FragranceInventorySort	= "expiryDate"
	FragranceInventorySortCost		FragranceInventorySort	= "cost"
	FragranceInventorySortWeight		FragranceInventorySort	= "weight"
	FragranceInventorySortCreatedAt		FragranceInventorySort	= "createdAt"
	FragranceInventorySortUpdatedAt		FragranceInventorySort	= "updatedAt"
	FragranceInventorySortDeletedAt		FragranceInventorySort	= "deletedAt"
)

var FragranceInventorySortDBValue = map[graphql_models.FragranceInventorySort]FragranceInventorySort{
	graphql_models.FragranceInventorySortID:		FragranceInventorySortID,
	graphql_models.FragranceInventorySortPurchaseDate:	FragranceInventorySortPurchaseDate,
	graphql_models.FragranceInventorySortExpiryDate:	FragranceInventorySortExpiryDate,
	graphql_models.FragranceInventorySortCost:		FragranceInventorySortCost,
	graphql_models.FragranceInventorySortWeight:		FragranceInventorySortWeight,
	graphql_models.FragranceInventorySortCreatedAt:		FragranceInventorySortCreatedAt,
	graphql_models.FragranceInventorySortUpdatedAt:		FragranceInventorySortUpdatedAt,
	graphql_models.FragranceInventorySortDeletedAt:		FragranceInventorySortDeletedAt,
}
var FragranceInventorySortAPIValue = map[FragranceInventorySort]graphql_models.FragranceInventorySort{
	FragranceInventorySortID:		graphql_models.FragranceInventorySortID,
	FragranceInventorySortPurchaseDate:	graphql_models.FragranceInventorySortPurchaseDate,
	FragranceInventorySortExpiryDate:	graphql_models.FragranceInventorySortExpiryDate,
	FragranceInventorySortCost:		graphql_models.FragranceInventorySortCost,
	FragranceInventorySortWeight:		graphql_models.FragranceInventorySortWeight,
	FragranceInventorySortCreatedAt:	graphql_models.FragranceInventorySortCreatedAt,
	FragranceInventorySortUpdatedAt:	graphql_models.FragranceInventorySortUpdatedAt,
	FragranceInventorySortDeletedAt:	graphql_models.FragranceInventorySortDeletedAt,
}

func NullDotStringToPointerFragranceInventorySort(v null.String) *graphql_models.FragranceInventorySort {
	s := StringToFragranceInventorySort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToFragranceInventorySort(v null.String) graphql_models.FragranceInventorySort {
	if !v.Valid {
		return ""
	}
	return StringToFragranceInventorySort(v.String)
}

func StringToFragranceInventorySort(v string) graphql_models.FragranceInventorySort {
	s := FragranceInventorySortAPIValue[FragranceInventorySort(v)]
	return s
}

func StringToPointerFragranceInventorySort(v string) *graphql_models.FragranceInventorySort {
	s := StringToFragranceInventorySort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerFragranceInventorySortToString(v *graphql_models.FragranceInventorySort) string {
	if v == nil {
		return ""
	}
	return FragranceInventorySortToString(*v)
}

func PointerFragranceInventorySortToNullDotString(v *graphql_models.FragranceInventorySort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return FragranceInventorySortToNullDotString(*v)
}

func FragranceInventorySortToNullDotString(v graphql_models.FragranceInventorySort) null.String {
	s := FragranceInventorySortToString(v)
	return null.NewString(s, s != "")
}

func FragranceInventorySortToString(v graphql_models.FragranceInventorySort) string {
	s := FragranceInventorySortDBValue[v]
	return string(s)
}

type FragranceSort string

const (
	FragranceSortID		FragranceSort	= "id"
	FragranceSortName	FragranceSort	= "name"
	FragranceSortNote	FragranceSort	= "note"
	FragranceSortCreatedAt	FragranceSort	= "createdAt"
	FragranceSortUpdatedAt	FragranceSort	= "updatedAt"
	FragranceSortDeletedAt	FragranceSort	= "deletedAt"
)

var FragranceSortDBValue = map[graphql_models.FragranceSort]FragranceSort{
	graphql_models.FragranceSortID:		FragranceSortID,
	graphql_models.FragranceSortName:	FragranceSortName,
	graphql_models.FragranceSortNote:	FragranceSortNote,
	graphql_models.FragranceSortCreatedAt:	FragranceSortCreatedAt,
	graphql_models.FragranceSortUpdatedAt:	FragranceSortUpdatedAt,
	graphql_models.FragranceSortDeletedAt:	FragranceSortDeletedAt,
}
var FragranceSortAPIValue = map[FragranceSort]graphql_models.FragranceSort{
	FragranceSortID:	graphql_models.FragranceSortID,
	FragranceSortName:	graphql_models.FragranceSortName,
	FragranceSortNote:	graphql_models.FragranceSortNote,
	FragranceSortCreatedAt:	graphql_models.FragranceSortCreatedAt,
	FragranceSortUpdatedAt:	graphql_models.FragranceSortUpdatedAt,
	FragranceSortDeletedAt:	graphql_models.FragranceSortDeletedAt,
}

func NullDotStringToPointerFragranceSort(v null.String) *graphql_models.FragranceSort {
	s := StringToFragranceSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToFragranceSort(v null.String) graphql_models.FragranceSort {
	if !v.Valid {
		return ""
	}
	return StringToFragranceSort(v.String)
}

func StringToFragranceSort(v string) graphql_models.FragranceSort {
	s := FragranceSortAPIValue[FragranceSort(v)]
	return s
}

func StringToPointerFragranceSort(v string) *graphql_models.FragranceSort {
	s := StringToFragranceSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerFragranceSortToString(v *graphql_models.FragranceSort) string {
	if v == nil {
		return ""
	}
	return FragranceSortToString(*v)
}

func PointerFragranceSortToNullDotString(v *graphql_models.FragranceSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return FragranceSortToNullDotString(*v)
}

func FragranceSortToNullDotString(v graphql_models.FragranceSort) null.String {
	s := FragranceSortToString(v)
	return null.NewString(s, s != "")
}

func FragranceSortToString(v graphql_models.FragranceSort) string {
	s := FragranceSortDBValue[v]
	return string(s)
}

type LipidSort string

const (
	LipidSortID		LipidSort	= "id"
	LipidSortName		LipidSort	= "name"
	LipidSortLauric		LipidSort	= "lauric"
	LipidSortMyristic	LipidSort	= "myristic"
	LipidSortPalmitic	LipidSort	= "palmitic"
	LipidSortStearic	LipidSort	= "stearic"
	LipidSortRicinoleic	LipidSort	= "ricinoleic"
	LipidSortOleic		LipidSort	= "oleic"
	LipidSortLinoleic	LipidSort	= "linoleic"
	LipidSortLinolenic	LipidSort	= "linolenic"
	LipidSortHardness	LipidSort	= "hardness"
	LipidSortCleansing	LipidSort	= "cleansing"
	LipidSortConditioning	LipidSort	= "conditioning"
	LipidSortBubbly		LipidSort	= "bubbly"
	LipidSortCreamy		LipidSort	= "creamy"
	LipidSortIodine		LipidSort	= "iodine"
	LipidSortIns		LipidSort	= "ins"
	LipidSortInciName	LipidSort	= "inciName"
	LipidSortFamily		LipidSort	= "family"
	LipidSortNaoh		LipidSort	= "naoh"
	LipidSortCreatedAt	LipidSort	= "createdAt"
	LipidSortUpdatedAt	LipidSort	= "updatedAt"
	LipidSortDeletedAt	LipidSort	= "deletedAt"
)

var LipidSortDBValue = map[graphql_models.LipidSort]LipidSort{
	graphql_models.LipidSortID:		LipidSortID,
	graphql_models.LipidSortName:		LipidSortName,
	graphql_models.LipidSortLauric:		LipidSortLauric,
	graphql_models.LipidSortMyristic:	LipidSortMyristic,
	graphql_models.LipidSortPalmitic:	LipidSortPalmitic,
	graphql_models.LipidSortStearic:	LipidSortStearic,
	graphql_models.LipidSortRicinoleic:	LipidSortRicinoleic,
	graphql_models.LipidSortOleic:		LipidSortOleic,
	graphql_models.LipidSortLinoleic:	LipidSortLinoleic,
	graphql_models.LipidSortLinolenic:	LipidSortLinolenic,
	graphql_models.LipidSortHardness:	LipidSortHardness,
	graphql_models.LipidSortCleansing:	LipidSortCleansing,
	graphql_models.LipidSortConditioning:	LipidSortConditioning,
	graphql_models.LipidSortBubbly:		LipidSortBubbly,
	graphql_models.LipidSortCreamy:		LipidSortCreamy,
	graphql_models.LipidSortIodine:		LipidSortIodine,
	graphql_models.LipidSortIns:		LipidSortIns,
	graphql_models.LipidSortInciName:	LipidSortInciName,
	graphql_models.LipidSortFamily:		LipidSortFamily,
	graphql_models.LipidSortNaoh:		LipidSortNaoh,
	graphql_models.LipidSortCreatedAt:	LipidSortCreatedAt,
	graphql_models.LipidSortUpdatedAt:	LipidSortUpdatedAt,
	graphql_models.LipidSortDeletedAt:	LipidSortDeletedAt,
}
var LipidSortAPIValue = map[LipidSort]graphql_models.LipidSort{
	LipidSortID:		graphql_models.LipidSortID,
	LipidSortName:		graphql_models.LipidSortName,
	LipidSortLauric:	graphql_models.LipidSortLauric,
	LipidSortMyristic:	graphql_models.LipidSortMyristic,
	LipidSortPalmitic:	graphql_models.LipidSortPalmitic,
	LipidSortStearic:	graphql_models.LipidSortStearic,
	LipidSortRicinoleic:	graphql_models.LipidSortRicinoleic,
	LipidSortOleic:		graphql_models.LipidSortOleic,
	LipidSortLinoleic:	graphql_models.LipidSortLinoleic,
	LipidSortLinolenic:	graphql_models.LipidSortLinolenic,
	LipidSortHardness:	graphql_models.LipidSortHardness,
	LipidSortCleansing:	graphql_models.LipidSortCleansing,
	LipidSortConditioning:	graphql_models.LipidSortConditioning,
	LipidSortBubbly:	graphql_models.LipidSortBubbly,
	LipidSortCreamy:	graphql_models.LipidSortCreamy,
	LipidSortIodine:	graphql_models.LipidSortIodine,
	LipidSortIns:		graphql_models.LipidSortIns,
	LipidSortInciName:	graphql_models.LipidSortInciName,
	LipidSortFamily:	graphql_models.LipidSortFamily,
	LipidSortNaoh:		graphql_models.LipidSortNaoh,
	LipidSortCreatedAt:	graphql_models.LipidSortCreatedAt,
	LipidSortUpdatedAt:	graphql_models.LipidSortUpdatedAt,
	LipidSortDeletedAt:	graphql_models.LipidSortDeletedAt,
}

func NullDotStringToPointerLipidSort(v null.String) *graphql_models.LipidSort {
	s := StringToLipidSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToLipidSort(v null.String) graphql_models.LipidSort {
	if !v.Valid {
		return ""
	}
	return StringToLipidSort(v.String)
}

func StringToLipidSort(v string) graphql_models.LipidSort {
	s := LipidSortAPIValue[LipidSort(v)]
	return s
}

func StringToPointerLipidSort(v string) *graphql_models.LipidSort {
	s := StringToLipidSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerLipidSortToString(v *graphql_models.LipidSort) string {
	if v == nil {
		return ""
	}
	return LipidSortToString(*v)
}

func PointerLipidSortToNullDotString(v *graphql_models.LipidSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return LipidSortToNullDotString(*v)
}

func LipidSortToNullDotString(v graphql_models.LipidSort) null.String {
	s := LipidSortToString(v)
	return null.NewString(s, s != "")
}

func LipidSortToString(v graphql_models.LipidSort) string {
	s := LipidSortDBValue[v]
	return string(s)
}

type LyeSort string

const (
	LyeSortID		LyeSort	= "id"
	LyeSortKind		LyeSort	= "kind"
	LyeSortName		LyeSort	= "name"
	LyeSortNote		LyeSort	= "note"
	LyeSortCreatedAt	LyeSort	= "createdAt"
	LyeSortUpdatedAt	LyeSort	= "updatedAt"
	LyeSortDeletedAt	LyeSort	= "deletedAt"
)

var LyeSortDBValue = map[graphql_models.LyeSort]LyeSort{
	graphql_models.LyeSortID:		LyeSortID,
	graphql_models.LyeSortKind:		LyeSortKind,
	graphql_models.LyeSortName:		LyeSortName,
	graphql_models.LyeSortNote:		LyeSortNote,
	graphql_models.LyeSortCreatedAt:	LyeSortCreatedAt,
	graphql_models.LyeSortUpdatedAt:	LyeSortUpdatedAt,
	graphql_models.LyeSortDeletedAt:	LyeSortDeletedAt,
}
var LyeSortAPIValue = map[LyeSort]graphql_models.LyeSort{
	LyeSortID:		graphql_models.LyeSortID,
	LyeSortKind:		graphql_models.LyeSortKind,
	LyeSortName:		graphql_models.LyeSortName,
	LyeSortNote:		graphql_models.LyeSortNote,
	LyeSortCreatedAt:	graphql_models.LyeSortCreatedAt,
	LyeSortUpdatedAt:	graphql_models.LyeSortUpdatedAt,
	LyeSortDeletedAt:	graphql_models.LyeSortDeletedAt,
}

func NullDotStringToPointerLyeSort(v null.String) *graphql_models.LyeSort {
	s := StringToLyeSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToLyeSort(v null.String) graphql_models.LyeSort {
	if !v.Valid {
		return ""
	}
	return StringToLyeSort(v.String)
}

func StringToLyeSort(v string) graphql_models.LyeSort {
	s := LyeSortAPIValue[LyeSort(v)]
	return s
}

func StringToPointerLyeSort(v string) *graphql_models.LyeSort {
	s := StringToLyeSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerLyeSortToString(v *graphql_models.LyeSort) string {
	if v == nil {
		return ""
	}
	return LyeSortToString(*v)
}

func PointerLyeSortToNullDotString(v *graphql_models.LyeSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return LyeSortToNullDotString(*v)
}

func LyeSortToNullDotString(v graphql_models.LyeSort) null.String {
	s := LyeSortToString(v)
	return null.NewString(s, s != "")
}

func LyeSortToString(v graphql_models.LyeSort) string {
	s := LyeSortDBValue[v]
	return string(s)
}

type RecipeStepSort string

const (
	RecipeStepSortID	RecipeStepSort	= "id"
	RecipeStepSortNum	RecipeStepSort	= "num"
	RecipeStepSortNote	RecipeStepSort	= "note"
	RecipeStepSortCreatedAt	RecipeStepSort	= "createdAt"
	RecipeStepSortUpdatedAt	RecipeStepSort	= "updatedAt"
	RecipeStepSortDeletedAt	RecipeStepSort	= "deletedAt"
)

var RecipeStepSortDBValue = map[graphql_models.RecipeStepSort]RecipeStepSort{
	graphql_models.RecipeStepSortID:	RecipeStepSortID,
	graphql_models.RecipeStepSortNum:	RecipeStepSortNum,
	graphql_models.RecipeStepSortNote:	RecipeStepSortNote,
	graphql_models.RecipeStepSortCreatedAt:	RecipeStepSortCreatedAt,
	graphql_models.RecipeStepSortUpdatedAt:	RecipeStepSortUpdatedAt,
	graphql_models.RecipeStepSortDeletedAt:	RecipeStepSortDeletedAt,
}
var RecipeStepSortAPIValue = map[RecipeStepSort]graphql_models.RecipeStepSort{
	RecipeStepSortID:		graphql_models.RecipeStepSortID,
	RecipeStepSortNum:		graphql_models.RecipeStepSortNum,
	RecipeStepSortNote:		graphql_models.RecipeStepSortNote,
	RecipeStepSortCreatedAt:	graphql_models.RecipeStepSortCreatedAt,
	RecipeStepSortUpdatedAt:	graphql_models.RecipeStepSortUpdatedAt,
	RecipeStepSortDeletedAt:	graphql_models.RecipeStepSortDeletedAt,
}

func NullDotStringToPointerRecipeStepSort(v null.String) *graphql_models.RecipeStepSort {
	s := StringToRecipeStepSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeStepSort(v null.String) graphql_models.RecipeStepSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeStepSort(v.String)
}

func StringToRecipeStepSort(v string) graphql_models.RecipeStepSort {
	s := RecipeStepSortAPIValue[RecipeStepSort(v)]
	return s
}

func StringToPointerRecipeStepSort(v string) *graphql_models.RecipeStepSort {
	s := StringToRecipeStepSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeStepSortToString(v *graphql_models.RecipeStepSort) string {
	if v == nil {
		return ""
	}
	return RecipeStepSortToString(*v)
}

func PointerRecipeStepSortToNullDotString(v *graphql_models.RecipeStepSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeStepSortToNullDotString(*v)
}

func RecipeStepSortToNullDotString(v graphql_models.RecipeStepSort) null.String {
	s := RecipeStepSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeStepSortToString(v graphql_models.RecipeStepSort) string {
	s := RecipeStepSortDBValue[v]
	return string(s)
}

type AuthGroupPermissionSort string

const (
	AuthGroupPermissionSortID AuthGroupPermissionSort = "id"
)

var AuthGroupPermissionSortDBValue = map[graphql_models.AuthGroupPermissionSort]AuthGroupPermissionSort{
	graphql_models.AuthGroupPermissionSortID: AuthGroupPermissionSortID,
}
var AuthGroupPermissionSortAPIValue = map[AuthGroupPermissionSort]graphql_models.AuthGroupPermissionSort{
	AuthGroupPermissionSortID: graphql_models.AuthGroupPermissionSortID,
}

func NullDotStringToPointerAuthGroupPermissionSort(v null.String) *graphql_models.AuthGroupPermissionSort {
	s := StringToAuthGroupPermissionSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToAuthGroupPermissionSort(v null.String) graphql_models.AuthGroupPermissionSort {
	if !v.Valid {
		return ""
	}
	return StringToAuthGroupPermissionSort(v.String)
}

func StringToAuthGroupPermissionSort(v string) graphql_models.AuthGroupPermissionSort {
	s := AuthGroupPermissionSortAPIValue[AuthGroupPermissionSort(v)]
	return s
}

func StringToPointerAuthGroupPermissionSort(v string) *graphql_models.AuthGroupPermissionSort {
	s := StringToAuthGroupPermissionSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerAuthGroupPermissionSortToString(v *graphql_models.AuthGroupPermissionSort) string {
	if v == nil {
		return ""
	}
	return AuthGroupPermissionSortToString(*v)
}

func PointerAuthGroupPermissionSortToNullDotString(v *graphql_models.AuthGroupPermissionSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return AuthGroupPermissionSortToNullDotString(*v)
}

func AuthGroupPermissionSortToNullDotString(v graphql_models.AuthGroupPermissionSort) null.String {
	s := AuthGroupPermissionSortToString(v)
	return null.NewString(s, s != "")
}

func AuthGroupPermissionSortToString(v graphql_models.AuthGroupPermissionSort) string {
	s := AuthGroupPermissionSortDBValue[v]
	return string(s)
}

type RecipeBatchFragranceSort string

const (
	RecipeBatchFragranceSortID		RecipeBatchFragranceSort	= "id"
	RecipeBatchFragranceSortWeight		RecipeBatchFragranceSort	= "weight"
	RecipeBatchFragranceSortCost		RecipeBatchFragranceSort	= "cost"
	RecipeBatchFragranceSortCreatedAt	RecipeBatchFragranceSort	= "createdAt"
	RecipeBatchFragranceSortUpdatedAt	RecipeBatchFragranceSort	= "updatedAt"
	RecipeBatchFragranceSortDeletedAt	RecipeBatchFragranceSort	= "deletedAt"
)

var RecipeBatchFragranceSortDBValue = map[graphql_models.RecipeBatchFragranceSort]RecipeBatchFragranceSort{
	graphql_models.RecipeBatchFragranceSortID:		RecipeBatchFragranceSortID,
	graphql_models.RecipeBatchFragranceSortWeight:		RecipeBatchFragranceSortWeight,
	graphql_models.RecipeBatchFragranceSortCost:		RecipeBatchFragranceSortCost,
	graphql_models.RecipeBatchFragranceSortCreatedAt:	RecipeBatchFragranceSortCreatedAt,
	graphql_models.RecipeBatchFragranceSortUpdatedAt:	RecipeBatchFragranceSortUpdatedAt,
	graphql_models.RecipeBatchFragranceSortDeletedAt:	RecipeBatchFragranceSortDeletedAt,
}
var RecipeBatchFragranceSortAPIValue = map[RecipeBatchFragranceSort]graphql_models.RecipeBatchFragranceSort{
	RecipeBatchFragranceSortID:		graphql_models.RecipeBatchFragranceSortID,
	RecipeBatchFragranceSortWeight:		graphql_models.RecipeBatchFragranceSortWeight,
	RecipeBatchFragranceSortCost:		graphql_models.RecipeBatchFragranceSortCost,
	RecipeBatchFragranceSortCreatedAt:	graphql_models.RecipeBatchFragranceSortCreatedAt,
	RecipeBatchFragranceSortUpdatedAt:	graphql_models.RecipeBatchFragranceSortUpdatedAt,
	RecipeBatchFragranceSortDeletedAt:	graphql_models.RecipeBatchFragranceSortDeletedAt,
}

func NullDotStringToPointerRecipeBatchFragranceSort(v null.String) *graphql_models.RecipeBatchFragranceSort {
	s := StringToRecipeBatchFragranceSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToRecipeBatchFragranceSort(v null.String) graphql_models.RecipeBatchFragranceSort {
	if !v.Valid {
		return ""
	}
	return StringToRecipeBatchFragranceSort(v.String)
}

func StringToRecipeBatchFragranceSort(v string) graphql_models.RecipeBatchFragranceSort {
	s := RecipeBatchFragranceSortAPIValue[RecipeBatchFragranceSort(v)]
	return s
}

func StringToPointerRecipeBatchFragranceSort(v string) *graphql_models.RecipeBatchFragranceSort {
	s := StringToRecipeBatchFragranceSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerRecipeBatchFragranceSortToString(v *graphql_models.RecipeBatchFragranceSort) string {
	if v == nil {
		return ""
	}
	return RecipeBatchFragranceSortToString(*v)
}

func PointerRecipeBatchFragranceSortToNullDotString(v *graphql_models.RecipeBatchFragranceSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return RecipeBatchFragranceSortToNullDotString(*v)
}

func RecipeBatchFragranceSortToNullDotString(v graphql_models.RecipeBatchFragranceSort) null.String {
	s := RecipeBatchFragranceSortToString(v)
	return null.NewString(s, s != "")
}

func RecipeBatchFragranceSortToString(v graphql_models.RecipeBatchFragranceSort) string {
	s := RecipeBatchFragranceSortDBValue[v]
	return string(s)
}

func AdditiveWithUintID(id uint) *graphql_models.Additive {
	return &graphql_models.Additive{
		ID: AdditiveIDToGraphQL(id),
	}
}

func AdditiveWithIntID(id int) *graphql_models.Additive {
	return AdditiveWithUintID(uint(id))
}

func AdditiveWithNullDotUintID(id null.Uint) *graphql_models.Additive {
	return AdditiveWithUintID(id.Uint)
}

func AdditiveWithNullDotIntID(id null.Int) *graphql_models.Additive {
	return AdditiveWithUintID(uint(id.Int))
}

func AdditivesToGraphQL(am []*models.Additive) []*graphql_models.Additive {
	ar := make([]*graphql_models.Additive, len(am))
	for i, m := range am {
		ar[i] = AdditiveToGraphQL(m)
	}
	return ar
}

func AdditiveIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Additive)
}

func AdditiveToGraphQL(m *models.Additive) *graphql_models.Additive {
	if m == nil {
		return nil
	}

	r := &graphql_models.Additive{
		ID:		AdditiveIDToGraphQL(uint(m.ID)),
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if m.R != nil && m.R.RecipeAdditive != nil {
		r.RecipeAdditive = RecipeAdditiveToGraphQL(m.R.RecipeAdditive)
	}
	if m.R != nil && m.R.RecipeBatchAdditive != nil {
		r.RecipeBatchAdditive = RecipeBatchAdditiveToGraphQL(m.R.RecipeBatchAdditive)
	}
	if m.R != nil && m.R.AdditiveInventories != nil {
		r.AdditiveInventories = AdditiveInventoriesToGraphQL(m.R.AdditiveInventories)
	}

	return r
}

func AdditiveID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AdditiveIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AdditiveInventoryWithUintID(id uint) *graphql_models.AdditiveInventory {
	return &graphql_models.AdditiveInventory{
		ID: AdditiveInventoryIDToGraphQL(id),
	}
}

func AdditiveInventoryWithIntID(id int) *graphql_models.AdditiveInventory {
	return AdditiveInventoryWithUintID(uint(id))
}

func AdditiveInventoryWithNullDotUintID(id null.Uint) *graphql_models.AdditiveInventory {
	return AdditiveInventoryWithUintID(id.Uint)
}

func AdditiveInventoryWithNullDotIntID(id null.Int) *graphql_models.AdditiveInventory {
	return AdditiveInventoryWithUintID(uint(id.Int))
}

func AdditiveInventoriesToGraphQL(am []*models.AdditiveInventory) []*graphql_models.AdditiveInventory {
	ar := make([]*graphql_models.AdditiveInventory, len(am))
	for i, m := range am {
		ar[i] = AdditiveInventoryToGraphQL(m)
	}
	return ar
}

func AdditiveInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AdditiveInventory)
}

func AdditiveInventoryToGraphQL(m *models.AdditiveInventory) *graphql_models.AdditiveInventory {
	if m == nil {
		return nil
	}

	r := &graphql_models.AdditiveInventory{
		ID:		AdditiveInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate:	boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:	boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.AdditiveID) {
		if m.R != nil && m.R.Additive != nil {
			r.Additive = AdditiveToGraphQL(m.R.Additive)
		} else {
			r.Additive = AdditiveWithIntID(m.AdditiveID)
		}
	}
	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func AdditiveInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AdditiveInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AuthGroupWithUintID(id uint) *graphql_models.AuthGroup {
	return &graphql_models.AuthGroup{
		ID: AuthGroupIDToGraphQL(id),
	}
}

func AuthGroupWithIntID(id int) *graphql_models.AuthGroup {
	return AuthGroupWithUintID(uint(id))
}

func AuthGroupWithNullDotUintID(id null.Uint) *graphql_models.AuthGroup {
	return AuthGroupWithUintID(id.Uint)
}

func AuthGroupWithNullDotIntID(id null.Int) *graphql_models.AuthGroup {
	return AuthGroupWithUintID(uint(id.Int))
}

func AuthGroupsToGraphQL(am []*models.AuthGroup) []*graphql_models.AuthGroup {
	ar := make([]*graphql_models.AuthGroup, len(am))
	for i, m := range am {
		ar[i] = AuthGroupToGraphQL(m)
	}
	return ar
}

func AuthGroupIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AuthGroup)
}

func AuthGroupToGraphQL(m *models.AuthGroup) *graphql_models.AuthGroup {
	if m == nil {
		return nil
	}

	r := &graphql_models.AuthGroup{
		ID:	AuthGroupIDToGraphQL(uint(m.ID)),
		Name:	m.Name,
	}

	if m.R != nil && m.R.GroupAuthGroupPermissions != nil {
		r.GroupAuthGroupPermissions = AuthGroupPermissionsToGraphQL(m.R.GroupAuthGroupPermissions)
	}
	if m.R != nil && m.R.GroupAuthUserGroups != nil {
		r.GroupAuthUserGroups = AuthUserGroupsToGraphQL(m.R.GroupAuthUserGroups)
	}

	return r
}

func AuthGroupID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AuthGroupIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AuthGroupPermissionWithUintID(id uint) *graphql_models.AuthGroupPermission {
	return &graphql_models.AuthGroupPermission{
		ID: AuthGroupPermissionIDToGraphQL(id),
	}
}

func AuthGroupPermissionWithIntID(id int) *graphql_models.AuthGroupPermission {
	return AuthGroupPermissionWithUintID(uint(id))
}

func AuthGroupPermissionWithNullDotUintID(id null.Uint) *graphql_models.AuthGroupPermission {
	return AuthGroupPermissionWithUintID(id.Uint)
}

func AuthGroupPermissionWithNullDotIntID(id null.Int) *graphql_models.AuthGroupPermission {
	return AuthGroupPermissionWithUintID(uint(id.Int))
}

func AuthGroupPermissionsToGraphQL(am []*models.AuthGroupPermission) []*graphql_models.AuthGroupPermission {
	ar := make([]*graphql_models.AuthGroupPermission, len(am))
	for i, m := range am {
		ar[i] = AuthGroupPermissionToGraphQL(m)
	}
	return ar
}

func AuthGroupPermissionIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AuthGroupPermissions)
}

func AuthGroupPermissionToGraphQL(m *models.AuthGroupPermission) *graphql_models.AuthGroupPermission {
	if m == nil {
		return nil
	}

	r := &graphql_models.AuthGroupPermission{
		ID: AuthGroupPermissionIDToGraphQL(uint(m.ID)),
	}

	if boilergql.IntIsFilled(m.GroupID) {
		if m.R != nil && m.R.Group != nil {
			r.Group = AuthGroupToGraphQL(m.R.Group)
		} else {
			r.Group = AuthGroupWithIntID(m.GroupID)
		}
	}
	if boilergql.IntIsFilled(m.PermissionID) {
		if m.R != nil && m.R.Permission != nil {
			r.Permission = AuthPermissionToGraphQL(m.R.Permission)
		} else {
			r.Permission = AuthPermissionWithIntID(m.PermissionID)
		}
	}

	return r
}

func AuthGroupPermissionID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AuthGroupPermissionIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AuthPermissionWithUintID(id uint) *graphql_models.AuthPermission {
	return &graphql_models.AuthPermission{
		ID: AuthPermissionIDToGraphQL(id),
	}
}

func AuthPermissionWithIntID(id int) *graphql_models.AuthPermission {
	return AuthPermissionWithUintID(uint(id))
}

func AuthPermissionWithNullDotUintID(id null.Uint) *graphql_models.AuthPermission {
	return AuthPermissionWithUintID(id.Uint)
}

func AuthPermissionWithNullDotIntID(id null.Int) *graphql_models.AuthPermission {
	return AuthPermissionWithUintID(uint(id.Int))
}

func AuthPermissionsToGraphQL(am []*models.AuthPermission) []*graphql_models.AuthPermission {
	ar := make([]*graphql_models.AuthPermission, len(am))
	for i, m := range am {
		ar[i] = AuthPermissionToGraphQL(m)
	}
	return ar
}

func AuthPermissionIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AuthPermission)
}

func AuthPermissionToGraphQL(m *models.AuthPermission) *graphql_models.AuthPermission {
	if m == nil {
		return nil
	}

	r := &graphql_models.AuthPermission{
		ID:		AuthPermissionIDToGraphQL(uint(m.ID)),
		Name:		m.Name,
		ContentTypeID:	boilergql.IntToString(m.ContentTypeID),
		Codename:	m.Codename,
	}

	if m.R != nil && m.R.PermissionAuthGroupPermissions != nil {
		r.PermissionAuthGroupPermissions = AuthGroupPermissionsToGraphQL(m.R.PermissionAuthGroupPermissions)
	}
	if m.R != nil && m.R.PermissionAuthUserUserPermissions != nil {
		r.PermissionAuthUserUserPermissions = AuthUserUserPermissionsToGraphQL(m.R.PermissionAuthUserUserPermissions)
	}

	return r
}

func AuthPermissionID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AuthPermissionIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AuthUserWithUintID(id uint) *graphql_models.AuthUser {
	return &graphql_models.AuthUser{
		ID: AuthUserIDToGraphQL(id),
	}
}

func AuthUserWithIntID(id int) *graphql_models.AuthUser {
	return AuthUserWithUintID(uint(id))
}

func AuthUserWithNullDotUintID(id null.Uint) *graphql_models.AuthUser {
	return AuthUserWithUintID(id.Uint)
}

func AuthUserWithNullDotIntID(id null.Int) *graphql_models.AuthUser {
	return AuthUserWithUintID(uint(id.Int))
}

func AuthUsersToGraphQL(am []*models.AuthUser) []*graphql_models.AuthUser {
	ar := make([]*graphql_models.AuthUser, len(am))
	for i, m := range am {
		ar[i] = AuthUserToGraphQL(m)
	}
	return ar
}

func AuthUserIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AuthUser)
}

func AuthUserToGraphQL(m *models.AuthUser) *graphql_models.AuthUser {
	if m == nil {
		return nil
	}

	r := &graphql_models.AuthUser{
		ID:		AuthUserIDToGraphQL(uint(m.ID)),
		Password:	m.Password,
		LastLogin:	boilergql.NullDotTimeToPointerInt(m.LastLogin),
		IsSuperuser:	m.IsSuperuser,
		Username:	m.Username,
		FirstName:	m.FirstName,
		LastName:	m.LastName,
		Email:		m.Email,
		IsStaff:	m.IsStaff,
		IsActive:	m.IsActive,
		DateJoined:	boilergql.TimeDotTimeToInt(m.DateJoined),
	}

	if m.R != nil && m.R.UserAuthUserGroups != nil {
		r.UserAuthUserGroups = AuthUserGroupsToGraphQL(m.R.UserAuthUserGroups)
	}
	if m.R != nil && m.R.UserAuthUserUserPermissions != nil {
		r.UserAuthUserUserPermissions = AuthUserUserPermissionsToGraphQL(m.R.UserAuthUserUserPermissions)
	}

	return r
}

func AuthUserID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AuthUserIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AuthUserGroupWithUintID(id uint) *graphql_models.AuthUserGroup {
	return &graphql_models.AuthUserGroup{
		ID: AuthUserGroupIDToGraphQL(id),
	}
}

func AuthUserGroupWithIntID(id int) *graphql_models.AuthUserGroup {
	return AuthUserGroupWithUintID(uint(id))
}

func AuthUserGroupWithNullDotUintID(id null.Uint) *graphql_models.AuthUserGroup {
	return AuthUserGroupWithUintID(id.Uint)
}

func AuthUserGroupWithNullDotIntID(id null.Int) *graphql_models.AuthUserGroup {
	return AuthUserGroupWithUintID(uint(id.Int))
}

func AuthUserGroupsToGraphQL(am []*models.AuthUserGroup) []*graphql_models.AuthUserGroup {
	ar := make([]*graphql_models.AuthUserGroup, len(am))
	for i, m := range am {
		ar[i] = AuthUserGroupToGraphQL(m)
	}
	return ar
}

func AuthUserGroupIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AuthUserGroups)
}

func AuthUserGroupToGraphQL(m *models.AuthUserGroup) *graphql_models.AuthUserGroup {
	if m == nil {
		return nil
	}

	r := &graphql_models.AuthUserGroup{
		ID: AuthUserGroupIDToGraphQL(uint(m.ID)),
	}

	if boilergql.IntIsFilled(m.UserID) {
		if m.R != nil && m.R.User != nil {
			r.User = AuthUserToGraphQL(m.R.User)
		} else {
			r.User = AuthUserWithIntID(m.UserID)
		}
	}
	if boilergql.IntIsFilled(m.GroupID) {
		if m.R != nil && m.R.Group != nil {
			r.Group = AuthGroupToGraphQL(m.R.Group)
		} else {
			r.Group = AuthGroupWithIntID(m.GroupID)
		}
	}

	return r
}

func AuthUserGroupID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AuthUserGroupIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func AuthUserUserPermissionWithUintID(id uint) *graphql_models.AuthUserUserPermission {
	return &graphql_models.AuthUserUserPermission{
		ID: AuthUserUserPermissionIDToGraphQL(id),
	}
}

func AuthUserUserPermissionWithIntID(id int) *graphql_models.AuthUserUserPermission {
	return AuthUserUserPermissionWithUintID(uint(id))
}

func AuthUserUserPermissionWithNullDotUintID(id null.Uint) *graphql_models.AuthUserUserPermission {
	return AuthUserUserPermissionWithUintID(id.Uint)
}

func AuthUserUserPermissionWithNullDotIntID(id null.Int) *graphql_models.AuthUserUserPermission {
	return AuthUserUserPermissionWithUintID(uint(id.Int))
}

func AuthUserUserPermissionsToGraphQL(am []*models.AuthUserUserPermission) []*graphql_models.AuthUserUserPermission {
	ar := make([]*graphql_models.AuthUserUserPermission, len(am))
	for i, m := range am {
		ar[i] = AuthUserUserPermissionToGraphQL(m)
	}
	return ar
}

func AuthUserUserPermissionIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.AuthUserUserPermissions)
}

func AuthUserUserPermissionToGraphQL(m *models.AuthUserUserPermission) *graphql_models.AuthUserUserPermission {
	if m == nil {
		return nil
	}

	r := &graphql_models.AuthUserUserPermission{
		ID: AuthUserUserPermissionIDToGraphQL(uint(m.ID)),
	}

	if boilergql.IntIsFilled(m.UserID) {
		if m.R != nil && m.R.User != nil {
			r.User = AuthUserToGraphQL(m.R.User)
		} else {
			r.User = AuthUserWithIntID(m.UserID)
		}
	}
	if boilergql.IntIsFilled(m.PermissionID) {
		if m.R != nil && m.R.Permission != nil {
			r.Permission = AuthPermissionToGraphQL(m.R.Permission)
		} else {
			r.Permission = AuthPermissionWithIntID(m.PermissionID)
		}
	}

	return r
}

func AuthUserUserPermissionID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func AuthUserUserPermissionIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func FragranceWithUintID(id uint) *graphql_models.Fragrance {
	return &graphql_models.Fragrance{
		ID: FragranceIDToGraphQL(id),
	}
}

func FragranceWithIntID(id int) *graphql_models.Fragrance {
	return FragranceWithUintID(uint(id))
}

func FragranceWithNullDotUintID(id null.Uint) *graphql_models.Fragrance {
	return FragranceWithUintID(id.Uint)
}

func FragranceWithNullDotIntID(id null.Int) *graphql_models.Fragrance {
	return FragranceWithUintID(uint(id.Int))
}

func FragrancesToGraphQL(am []*models.Fragrance) []*graphql_models.Fragrance {
	ar := make([]*graphql_models.Fragrance, len(am))
	for i, m := range am {
		ar[i] = FragranceToGraphQL(m)
	}
	return ar
}

func FragranceIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Fragrance)
}

func FragranceToGraphQL(m *models.Fragrance) *graphql_models.Fragrance {
	if m == nil {
		return nil
	}

	r := &graphql_models.Fragrance{
		ID:		FragranceIDToGraphQL(uint(m.ID)),
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if m.R != nil && m.R.RecipeBatchFragrance != nil {
		r.RecipeBatchFragrance = RecipeBatchFragranceToGraphQL(m.R.RecipeBatchFragrance)
	}
	if m.R != nil && m.R.RecipeFragrance != nil {
		r.RecipeFragrance = RecipeFragranceToGraphQL(m.R.RecipeFragrance)
	}
	if m.R != nil && m.R.FragranceInventories != nil {
		r.FragranceInventories = FragranceInventoriesToGraphQL(m.R.FragranceInventories)
	}

	return r
}

func FragranceID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func FragranceIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func FragranceInventoryWithUintID(id uint) *graphql_models.FragranceInventory {
	return &graphql_models.FragranceInventory{
		ID: FragranceInventoryIDToGraphQL(id),
	}
}

func FragranceInventoryWithIntID(id int) *graphql_models.FragranceInventory {
	return FragranceInventoryWithUintID(uint(id))
}

func FragranceInventoryWithNullDotUintID(id null.Uint) *graphql_models.FragranceInventory {
	return FragranceInventoryWithUintID(id.Uint)
}

func FragranceInventoryWithNullDotIntID(id null.Int) *graphql_models.FragranceInventory {
	return FragranceInventoryWithUintID(uint(id.Int))
}

func FragranceInventoriesToGraphQL(am []*models.FragranceInventory) []*graphql_models.FragranceInventory {
	ar := make([]*graphql_models.FragranceInventory, len(am))
	for i, m := range am {
		ar[i] = FragranceInventoryToGraphQL(m)
	}
	return ar
}

func FragranceInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.FragranceInventory)
}

func FragranceInventoryToGraphQL(m *models.FragranceInventory) *graphql_models.FragranceInventory {
	if m == nil {
		return nil
	}

	r := &graphql_models.FragranceInventory{
		ID:		FragranceInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate:	boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:	boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.FragranceID) {
		if m.R != nil && m.R.Fragrance != nil {
			r.Fragrance = FragranceToGraphQL(m.R.Fragrance)
		} else {
			r.Fragrance = FragranceWithIntID(m.FragranceID)
		}
	}
	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func FragranceInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func FragranceInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LipidWithUintID(id uint) *graphql_models.Lipid {
	return &graphql_models.Lipid{
		ID: LipidIDToGraphQL(id),
	}
}

func LipidWithIntID(id int) *graphql_models.Lipid {
	return LipidWithUintID(uint(id))
}

func LipidWithNullDotUintID(id null.Uint) *graphql_models.Lipid {
	return LipidWithUintID(id.Uint)
}

func LipidWithNullDotIntID(id null.Int) *graphql_models.Lipid {
	return LipidWithUintID(uint(id.Int))
}

func LipidsToGraphQL(am []*models.Lipid) []*graphql_models.Lipid {
	ar := make([]*graphql_models.Lipid, len(am))
	for i, m := range am {
		ar[i] = LipidToGraphQL(m)
	}
	return ar
}

func LipidIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Lipid)
}

func LipidToGraphQL(m *models.Lipid) *graphql_models.Lipid {
	if m == nil {
		return nil
	}

	r := &graphql_models.Lipid{
		ID:		LipidIDToGraphQL(uint(m.ID)),
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
		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if m.R != nil && m.R.RecipeBatchLipid != nil {
		r.RecipeBatchLipid = RecipeBatchLipidToGraphQL(m.R.RecipeBatchLipid)
	}
	if m.R != nil && m.R.RecipeLipid != nil {
		r.RecipeLipid = RecipeLipidToGraphQL(m.R.RecipeLipid)
	}
	if m.R != nil && m.R.LipidInventories != nil {
		r.LipidInventories = LipidInventoriesToGraphQL(m.R.LipidInventories)
	}

	return r
}

func LipidID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LipidIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LipidInventoryWithUintID(id uint) *graphql_models.LipidInventory {
	return &graphql_models.LipidInventory{
		ID: LipidInventoryIDToGraphQL(id),
	}
}

func LipidInventoryWithIntID(id int) *graphql_models.LipidInventory {
	return LipidInventoryWithUintID(uint(id))
}

func LipidInventoryWithNullDotUintID(id null.Uint) *graphql_models.LipidInventory {
	return LipidInventoryWithUintID(id.Uint)
}

func LipidInventoryWithNullDotIntID(id null.Int) *graphql_models.LipidInventory {
	return LipidInventoryWithUintID(uint(id.Int))
}

func LipidInventoriesToGraphQL(am []*models.LipidInventory) []*graphql_models.LipidInventory {
	ar := make([]*graphql_models.LipidInventory, len(am))
	for i, m := range am {
		ar[i] = LipidInventoryToGraphQL(m)
	}
	return ar
}

func LipidInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.LipidInventory)
}

func LipidInventoryToGraphQL(m *models.LipidInventory) *graphql_models.LipidInventory {
	if m == nil {
		return nil
	}

	r := &graphql_models.LipidInventory{
		ID:		LipidInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate:	boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:	boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,
		Sap:		m.Sap,
		Naoh:		m.Naoh,
		Koh:		m.Koh,
		GramsPerLiter:	m.GramsPerLiter,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.LipidID) {
		if m.R != nil && m.R.Lipid != nil {
			r.Lipid = LipidToGraphQL(m.R.Lipid)
		} else {
			r.Lipid = LipidWithIntID(m.LipidID)
		}
	}
	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func LipidInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LipidInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LyeWithUintID(id uint) *graphql_models.Lye {
	return &graphql_models.Lye{
		ID: LyeIDToGraphQL(id),
	}
}

func LyeWithIntID(id int) *graphql_models.Lye {
	return LyeWithUintID(uint(id))
}

func LyeWithNullDotUintID(id null.Uint) *graphql_models.Lye {
	return LyeWithUintID(id.Uint)
}

func LyeWithNullDotIntID(id null.Int) *graphql_models.Lye {
	return LyeWithUintID(uint(id.Int))
}

func LyesToGraphQL(am []*models.Lye) []*graphql_models.Lye {
	ar := make([]*graphql_models.Lye, len(am))
	for i, m := range am {
		ar[i] = LyeToGraphQL(m)
	}
	return ar
}

func LyeIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Lye)
}

func LyeToGraphQL(m *models.Lye) *graphql_models.Lye {
	if m == nil {
		return nil
	}

	r := &graphql_models.Lye{
		ID:		LyeIDToGraphQL(uint(m.ID)),
		Kind:		m.Kind,
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if m.R != nil && m.R.RecipeBatchLye != nil {
		r.RecipeBatchLye = RecipeBatchLyeToGraphQL(m.R.RecipeBatchLye)
	}
	if m.R != nil && m.R.LyeInventories != nil {
		r.LyeInventories = LyeInventoriesToGraphQL(m.R.LyeInventories)
	}

	return r
}

func LyeID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LyeIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func LyeInventoryWithUintID(id uint) *graphql_models.LyeInventory {
	return &graphql_models.LyeInventory{
		ID: LyeInventoryIDToGraphQL(id),
	}
}

func LyeInventoryWithIntID(id int) *graphql_models.LyeInventory {
	return LyeInventoryWithUintID(uint(id))
}

func LyeInventoryWithNullDotUintID(id null.Uint) *graphql_models.LyeInventory {
	return LyeInventoryWithUintID(id.Uint)
}

func LyeInventoryWithNullDotIntID(id null.Int) *graphql_models.LyeInventory {
	return LyeInventoryWithUintID(uint(id.Int))
}

func LyeInventoriesToGraphQL(am []*models.LyeInventory) []*graphql_models.LyeInventory {
	ar := make([]*graphql_models.LyeInventory, len(am))
	for i, m := range am {
		ar[i] = LyeInventoryToGraphQL(m)
	}
	return ar
}

func LyeInventoryIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.LyeInventory)
}

func LyeInventoryToGraphQL(m *models.LyeInventory) *graphql_models.LyeInventory {
	if m == nil {
		return nil
	}

	r := &graphql_models.LyeInventory{
		ID:		LyeInventoryIDToGraphQL(uint(m.ID)),
		PurchaseDate:	boilergql.TimeDotTimeToInt(m.PurchaseDate),
		ExpiryDate:	boilergql.TimeDotTimeToInt(m.ExpiryDate),
		Cost:		m.Cost,
		Weight:		m.Weight,
		Concentration:	m.Concentration,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.LyeID) {
		if m.R != nil && m.R.Lye != nil {
			r.Lye = LyeToGraphQL(m.R.Lye)
		} else {
			r.Lye = LyeWithIntID(m.LyeID)
		}
	}
	if boilergql.IntIsFilled(m.SupplierID) {
		if m.R != nil && m.R.Supplier != nil {
			r.Supplier = SupplierToGraphQL(m.R.Supplier)
		} else {
			r.Supplier = SupplierWithIntID(m.SupplierID)
		}
	}

	return r
}

func LyeInventoryID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func LyeInventoryIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeWithUintID(id uint) *graphql_models.Recipe {
	return &graphql_models.Recipe{
		ID: RecipeIDToGraphQL(id),
	}
}

func RecipeWithIntID(id int) *graphql_models.Recipe {
	return RecipeWithUintID(uint(id))
}

func RecipeWithNullDotUintID(id null.Uint) *graphql_models.Recipe {
	return RecipeWithUintID(id.Uint)
}

func RecipeWithNullDotIntID(id null.Int) *graphql_models.Recipe {
	return RecipeWithUintID(uint(id.Int))
}

func RecipesToGraphQL(am []*models.Recipe) []*graphql_models.Recipe {
	ar := make([]*graphql_models.Recipe, len(am))
	for i, m := range am {
		ar[i] = RecipeToGraphQL(m)
	}
	return ar
}

func RecipeIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Recipe)
}

func RecipeToGraphQL(m *models.Recipe) *graphql_models.Recipe {
	if m == nil {
		return nil
	}

	r := &graphql_models.Recipe{
		ID:		RecipeIDToGraphQL(uint(m.ID)),
		Name:		m.Name,
		Note:		m.Note,
		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if m.R != nil && m.R.RecipeAdditives != nil {
		r.RecipeAdditives = RecipeAdditivesToGraphQL(m.R.RecipeAdditives)
	}
	if m.R != nil && m.R.RecipeBatches != nil {
		r.RecipeBatches = RecipeBatchesToGraphQL(m.R.RecipeBatches)
	}
	if m.R != nil && m.R.RecipeFragrances != nil {
		r.RecipeFragrances = RecipeFragrancesToGraphQL(m.R.RecipeFragrances)
	}
	if m.R != nil && m.R.RecipeLipids != nil {
		r.RecipeLipids = RecipeLipidsToGraphQL(m.R.RecipeLipids)
	}
	if m.R != nil && m.R.RecipeSteps != nil {
		r.RecipeSteps = RecipeStepsToGraphQL(m.R.RecipeSteps)
	}

	return r
}

func RecipeID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeAdditiveWithUintID(id uint) *graphql_models.RecipeAdditive {
	return &graphql_models.RecipeAdditive{
		ID: RecipeAdditiveIDToGraphQL(id),
	}
}

func RecipeAdditiveWithIntID(id int) *graphql_models.RecipeAdditive {
	return RecipeAdditiveWithUintID(uint(id))
}

func RecipeAdditiveWithNullDotUintID(id null.Uint) *graphql_models.RecipeAdditive {
	return RecipeAdditiveWithUintID(id.Uint)
}

func RecipeAdditiveWithNullDotIntID(id null.Int) *graphql_models.RecipeAdditive {
	return RecipeAdditiveWithUintID(uint(id.Int))
}

func RecipeAdditivesToGraphQL(am []*models.RecipeAdditive) []*graphql_models.RecipeAdditive {
	ar := make([]*graphql_models.RecipeAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeAdditiveToGraphQL(m)
	}
	return ar
}

func RecipeAdditiveIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeAdditive)
}

func RecipeAdditiveToGraphQL(m *models.RecipeAdditive) *graphql_models.RecipeAdditive {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeAdditive{
		ID:		RecipeAdditiveIDToGraphQL(uint(m.ID)),
		Percentage:	m.Percentage,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.AdditiveID) {
		if m.R != nil && m.R.Additive != nil {
			r.Additive = AdditiveToGraphQL(m.R.Additive)
		} else {
			r.Additive = AdditiveWithIntID(m.AdditiveID)
		}
	}
	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeAdditiveID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeAdditiveIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchWithUintID(id uint) *graphql_models.RecipeBatch {
	return &graphql_models.RecipeBatch{
		ID: RecipeBatchIDToGraphQL(id),
	}
}

func RecipeBatchWithIntID(id int) *graphql_models.RecipeBatch {
	return RecipeBatchWithUintID(uint(id))
}

func RecipeBatchWithNullDotUintID(id null.Uint) *graphql_models.RecipeBatch {
	return RecipeBatchWithUintID(id.Uint)
}

func RecipeBatchWithNullDotIntID(id null.Int) *graphql_models.RecipeBatch {
	return RecipeBatchWithUintID(uint(id.Int))
}

func RecipeBatchesToGraphQL(am []*models.RecipeBatch) []*graphql_models.RecipeBatch {
	ar := make([]*graphql_models.RecipeBatch, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchToGraphQL(m)
	}
	return ar
}

func RecipeBatchIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatch)
}

func RecipeBatchToGraphQL(m *models.RecipeBatch) *graphql_models.RecipeBatch {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeBatch{
		ID:			RecipeBatchIDToGraphQL(uint(m.ID)),
		Tag:			m.Tag,
		ProductionDate:		boilergql.TimeDotTimeToInt(m.ProductionDate),
		SellableDate:		boilergql.TimeDotTimeToInt(m.SellableDate),
		Note:			m.Note,
		LipidWeight:		m.LipidWeight,
		ProductionWeight:	m.ProductionWeight,
		CuredWeight:		m.CuredWeight,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}
	if m.R != nil && m.R.BatchRecipeBatchAdditives != nil {
		r.BatchRecipeBatchAdditives = RecipeBatchAdditivesToGraphQL(m.R.BatchRecipeBatchAdditives)
	}
	if m.R != nil && m.R.BatchRecipeBatchFragrances != nil {
		r.BatchRecipeBatchFragrances = RecipeBatchFragrancesToGraphQL(m.R.BatchRecipeBatchFragrances)
	}
	if m.R != nil && m.R.BatchRecipeBatchLipids != nil {
		r.BatchRecipeBatchLipids = RecipeBatchLipidsToGraphQL(m.R.BatchRecipeBatchLipids)
	}
	if m.R != nil && m.R.BatchRecipeBatchLyes != nil {
		r.BatchRecipeBatchLyes = RecipeBatchLyesToGraphQL(m.R.BatchRecipeBatchLyes)
	}
	if m.R != nil && m.R.BatchRecipeBatchNotes != nil {
		r.BatchRecipeBatchNotes = RecipeBatchNotesToGraphQL(m.R.BatchRecipeBatchNotes)
	}

	return r
}

func RecipeBatchID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchAdditiveWithUintID(id uint) *graphql_models.RecipeBatchAdditive {
	return &graphql_models.RecipeBatchAdditive{
		ID: RecipeBatchAdditiveIDToGraphQL(id),
	}
}

func RecipeBatchAdditiveWithIntID(id int) *graphql_models.RecipeBatchAdditive {
	return RecipeBatchAdditiveWithUintID(uint(id))
}

func RecipeBatchAdditiveWithNullDotUintID(id null.Uint) *graphql_models.RecipeBatchAdditive {
	return RecipeBatchAdditiveWithUintID(id.Uint)
}

func RecipeBatchAdditiveWithNullDotIntID(id null.Int) *graphql_models.RecipeBatchAdditive {
	return RecipeBatchAdditiveWithUintID(uint(id.Int))
}

func RecipeBatchAdditivesToGraphQL(am []*models.RecipeBatchAdditive) []*graphql_models.RecipeBatchAdditive {
	ar := make([]*graphql_models.RecipeBatchAdditive, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchAdditiveToGraphQL(m)
	}
	return ar
}

func RecipeBatchAdditiveIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchAdditive)
}

func RecipeBatchAdditiveToGraphQL(m *models.RecipeBatchAdditive) *graphql_models.RecipeBatchAdditive {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeBatchAdditive{
		ID:	RecipeBatchAdditiveIDToGraphQL(uint(m.ID)),
		Weight:	m.Weight,
		Cost:	m.Cost,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.AdditiveID) {
		if m.R != nil && m.R.Additive != nil {
			r.Additive = AdditiveToGraphQL(m.R.Additive)
		} else {
			r.Additive = AdditiveWithIntID(m.AdditiveID)
		}
	}
	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchAdditiveID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchAdditiveIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchFragranceWithUintID(id uint) *graphql_models.RecipeBatchFragrance {
	return &graphql_models.RecipeBatchFragrance{
		ID: RecipeBatchFragranceIDToGraphQL(id),
	}
}

func RecipeBatchFragranceWithIntID(id int) *graphql_models.RecipeBatchFragrance {
	return RecipeBatchFragranceWithUintID(uint(id))
}

func RecipeBatchFragranceWithNullDotUintID(id null.Uint) *graphql_models.RecipeBatchFragrance {
	return RecipeBatchFragranceWithUintID(id.Uint)
}

func RecipeBatchFragranceWithNullDotIntID(id null.Int) *graphql_models.RecipeBatchFragrance {
	return RecipeBatchFragranceWithUintID(uint(id.Int))
}

func RecipeBatchFragrancesToGraphQL(am []*models.RecipeBatchFragrance) []*graphql_models.RecipeBatchFragrance {
	ar := make([]*graphql_models.RecipeBatchFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchFragranceToGraphQL(m)
	}
	return ar
}

func RecipeBatchFragranceIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchFragrance)
}

func RecipeBatchFragranceToGraphQL(m *models.RecipeBatchFragrance) *graphql_models.RecipeBatchFragrance {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeBatchFragrance{
		ID:	RecipeBatchFragranceIDToGraphQL(uint(m.ID)),
		Weight:	m.Weight,
		Cost:	m.Cost,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.FragranceID) {
		if m.R != nil && m.R.Fragrance != nil {
			r.Fragrance = FragranceToGraphQL(m.R.Fragrance)
		} else {
			r.Fragrance = FragranceWithIntID(m.FragranceID)
		}
	}
	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchFragranceID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchFragranceIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchLipidWithUintID(id uint) *graphql_models.RecipeBatchLipid {
	return &graphql_models.RecipeBatchLipid{
		ID: RecipeBatchLipidIDToGraphQL(id),
	}
}

func RecipeBatchLipidWithIntID(id int) *graphql_models.RecipeBatchLipid {
	return RecipeBatchLipidWithUintID(uint(id))
}

func RecipeBatchLipidWithNullDotUintID(id null.Uint) *graphql_models.RecipeBatchLipid {
	return RecipeBatchLipidWithUintID(id.Uint)
}

func RecipeBatchLipidWithNullDotIntID(id null.Int) *graphql_models.RecipeBatchLipid {
	return RecipeBatchLipidWithUintID(uint(id.Int))
}

func RecipeBatchLipidsToGraphQL(am []*models.RecipeBatchLipid) []*graphql_models.RecipeBatchLipid {
	ar := make([]*graphql_models.RecipeBatchLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLipidToGraphQL(m)
	}
	return ar
}

func RecipeBatchLipidIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchLipid)
}

func RecipeBatchLipidToGraphQL(m *models.RecipeBatchLipid) *graphql_models.RecipeBatchLipid {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeBatchLipid{
		ID:	RecipeBatchLipidIDToGraphQL(uint(m.ID)),
		Weight:	m.Weight,
		Cost:	m.Cost,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.LipidID) {
		if m.R != nil && m.R.Lipid != nil {
			r.Lipid = LipidToGraphQL(m.R.Lipid)
		} else {
			r.Lipid = LipidWithIntID(m.LipidID)
		}
	}
	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchLipidID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchLipidIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchLyeWithUintID(id uint) *graphql_models.RecipeBatchLye {
	return &graphql_models.RecipeBatchLye{
		ID: RecipeBatchLyeIDToGraphQL(id),
	}
}

func RecipeBatchLyeWithIntID(id int) *graphql_models.RecipeBatchLye {
	return RecipeBatchLyeWithUintID(uint(id))
}

func RecipeBatchLyeWithNullDotUintID(id null.Uint) *graphql_models.RecipeBatchLye {
	return RecipeBatchLyeWithUintID(id.Uint)
}

func RecipeBatchLyeWithNullDotIntID(id null.Int) *graphql_models.RecipeBatchLye {
	return RecipeBatchLyeWithUintID(uint(id.Int))
}

func RecipeBatchLyesToGraphQL(am []*models.RecipeBatchLye) []*graphql_models.RecipeBatchLye {
	ar := make([]*graphql_models.RecipeBatchLye, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchLyeToGraphQL(m)
	}
	return ar
}

func RecipeBatchLyeIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchLye)
}

func RecipeBatchLyeToGraphQL(m *models.RecipeBatchLye) *graphql_models.RecipeBatchLye {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeBatchLye{
		ID:		RecipeBatchLyeIDToGraphQL(uint(m.ID)),
		Weight:		m.Weight,
		Discount:	m.Discount,
		Cost:		m.Cost,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.LyeID) {
		if m.R != nil && m.R.Lye != nil {
			r.Lye = LyeToGraphQL(m.R.Lye)
		} else {
			r.Lye = LyeWithIntID(m.LyeID)
		}
	}
	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchLyeID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchLyeIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeBatchNoteWithUintID(id uint) *graphql_models.RecipeBatchNote {
	return &graphql_models.RecipeBatchNote{
		ID: RecipeBatchNoteIDToGraphQL(id),
	}
}

func RecipeBatchNoteWithIntID(id int) *graphql_models.RecipeBatchNote {
	return RecipeBatchNoteWithUintID(uint(id))
}

func RecipeBatchNoteWithNullDotUintID(id null.Uint) *graphql_models.RecipeBatchNote {
	return RecipeBatchNoteWithUintID(id.Uint)
}

func RecipeBatchNoteWithNullDotIntID(id null.Int) *graphql_models.RecipeBatchNote {
	return RecipeBatchNoteWithUintID(uint(id.Int))
}

func RecipeBatchNotesToGraphQL(am []*models.RecipeBatchNote) []*graphql_models.RecipeBatchNote {
	ar := make([]*graphql_models.RecipeBatchNote, len(am))
	for i, m := range am {
		ar[i] = RecipeBatchNoteToGraphQL(m)
	}
	return ar
}

func RecipeBatchNoteIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeBatchNote)
}

func RecipeBatchNoteToGraphQL(m *models.RecipeBatchNote) *graphql_models.RecipeBatchNote {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeBatchNote{
		ID:	RecipeBatchNoteIDToGraphQL(uint(m.ID)),
		Note:	m.Note,
		Link:	m.Link,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.BatchID) {
		if m.R != nil && m.R.Batch != nil {
			r.Batch = RecipeBatchToGraphQL(m.R.Batch)
		} else {
			r.Batch = RecipeBatchWithIntID(m.BatchID)
		}
	}

	return r
}

func RecipeBatchNoteID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeBatchNoteIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeFragranceWithUintID(id uint) *graphql_models.RecipeFragrance {
	return &graphql_models.RecipeFragrance{
		ID: RecipeFragranceIDToGraphQL(id),
	}
}

func RecipeFragranceWithIntID(id int) *graphql_models.RecipeFragrance {
	return RecipeFragranceWithUintID(uint(id))
}

func RecipeFragranceWithNullDotUintID(id null.Uint) *graphql_models.RecipeFragrance {
	return RecipeFragranceWithUintID(id.Uint)
}

func RecipeFragranceWithNullDotIntID(id null.Int) *graphql_models.RecipeFragrance {
	return RecipeFragranceWithUintID(uint(id.Int))
}

func RecipeFragrancesToGraphQL(am []*models.RecipeFragrance) []*graphql_models.RecipeFragrance {
	ar := make([]*graphql_models.RecipeFragrance, len(am))
	for i, m := range am {
		ar[i] = RecipeFragranceToGraphQL(m)
	}
	return ar
}

func RecipeFragranceIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeFragrance)
}

func RecipeFragranceToGraphQL(m *models.RecipeFragrance) *graphql_models.RecipeFragrance {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeFragrance{
		ID:		RecipeFragranceIDToGraphQL(uint(m.ID)),
		Percentage:	m.Percentage,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.FragranceID) {
		if m.R != nil && m.R.Fragrance != nil {
			r.Fragrance = FragranceToGraphQL(m.R.Fragrance)
		} else {
			r.Fragrance = FragranceWithIntID(m.FragranceID)
		}
	}
	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeFragranceID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeFragranceIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeLipidWithUintID(id uint) *graphql_models.RecipeLipid {
	return &graphql_models.RecipeLipid{
		ID: RecipeLipidIDToGraphQL(id),
	}
}

func RecipeLipidWithIntID(id int) *graphql_models.RecipeLipid {
	return RecipeLipidWithUintID(uint(id))
}

func RecipeLipidWithNullDotUintID(id null.Uint) *graphql_models.RecipeLipid {
	return RecipeLipidWithUintID(id.Uint)
}

func RecipeLipidWithNullDotIntID(id null.Int) *graphql_models.RecipeLipid {
	return RecipeLipidWithUintID(uint(id.Int))
}

func RecipeLipidsToGraphQL(am []*models.RecipeLipid) []*graphql_models.RecipeLipid {
	ar := make([]*graphql_models.RecipeLipid, len(am))
	for i, m := range am {
		ar[i] = RecipeLipidToGraphQL(m)
	}
	return ar
}

func RecipeLipidIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeLipid)
}

func RecipeLipidToGraphQL(m *models.RecipeLipid) *graphql_models.RecipeLipid {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeLipid{
		ID:		RecipeLipidIDToGraphQL(uint(m.ID)),
		Percentage:	m.Percentage,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.LipidID) {
		if m.R != nil && m.R.Lipid != nil {
			r.Lipid = LipidToGraphQL(m.R.Lipid)
		} else {
			r.Lipid = LipidWithIntID(m.LipidID)
		}
	}
	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeLipidID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeLipidIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func RecipeStepWithUintID(id uint) *graphql_models.RecipeStep {
	return &graphql_models.RecipeStep{
		ID: RecipeStepIDToGraphQL(id),
	}
}

func RecipeStepWithIntID(id int) *graphql_models.RecipeStep {
	return RecipeStepWithUintID(uint(id))
}

func RecipeStepWithNullDotUintID(id null.Uint) *graphql_models.RecipeStep {
	return RecipeStepWithUintID(id.Uint)
}

func RecipeStepWithNullDotIntID(id null.Int) *graphql_models.RecipeStep {
	return RecipeStepWithUintID(uint(id.Int))
}

func RecipeStepsToGraphQL(am []*models.RecipeStep) []*graphql_models.RecipeStep {
	ar := make([]*graphql_models.RecipeStep, len(am))
	for i, m := range am {
		ar[i] = RecipeStepToGraphQL(m)
	}
	return ar
}

func RecipeStepIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.RecipeStep)
}

func RecipeStepToGraphQL(m *models.RecipeStep) *graphql_models.RecipeStep {
	if m == nil {
		return nil
	}

	r := &graphql_models.RecipeStep{
		ID:	RecipeStepIDToGraphQL(uint(m.ID)),
		Num:	m.Num,
		Note:	m.Note,

		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if boilergql.IntIsFilled(m.RecipeID) {
		if m.R != nil && m.R.Recipe != nil {
			r.Recipe = RecipeToGraphQL(m.R.Recipe)
		} else {
			r.Recipe = RecipeWithIntID(m.RecipeID)
		}
	}

	return r
}

func RecipeStepID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func RecipeStepIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}

func SupplierWithUintID(id uint) *graphql_models.Supplier {
	return &graphql_models.Supplier{
		ID: SupplierIDToGraphQL(id),
	}
}

func SupplierWithIntID(id int) *graphql_models.Supplier {
	return SupplierWithUintID(uint(id))
}

func SupplierWithNullDotUintID(id null.Uint) *graphql_models.Supplier {
	return SupplierWithUintID(id.Uint)
}

func SupplierWithNullDotIntID(id null.Int) *graphql_models.Supplier {
	return SupplierWithUintID(uint(id.Int))
}

func SuppliersToGraphQL(am []*models.Supplier) []*graphql_models.Supplier {
	ar := make([]*graphql_models.Supplier, len(am))
	for i, m := range am {
		ar[i] = SupplierToGraphQL(m)
	}
	return ar
}

func SupplierIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.Supplier)
}

func SupplierToGraphQL(m *models.Supplier) *graphql_models.Supplier {
	if m == nil {
		return nil
	}

	r := &graphql_models.Supplier{
		ID:		SupplierIDToGraphQL(uint(m.ID)),
		Name:		m.Name,
		Website:	m.Website,
		Note:		m.Note,
		CreatedAt:	boilergql.TimeDotTimeToInt(m.CreatedAt),
		UpdatedAt:	boilergql.TimeDotTimeToInt(m.UpdatedAt),
		DeletedAt:	boilergql.NullDotTimeToPointerInt(m.DeletedAt),
	}

	if m.R != nil && m.R.AdditiveInventories != nil {
		r.AdditiveInventories = AdditiveInventoriesToGraphQL(m.R.AdditiveInventories)
	}
	if m.R != nil && m.R.FragranceInventories != nil {
		r.FragranceInventories = FragranceInventoriesToGraphQL(m.R.FragranceInventories)
	}
	if m.R != nil && m.R.LipidInventories != nil {
		r.LipidInventories = LipidInventoriesToGraphQL(m.R.LipidInventories)
	}
	if m.R != nil && m.R.LyeInventories != nil {
		r.LyeInventories = LyeInventoriesToGraphQL(m.R.LyeInventories)
	}

	return r
}

func SupplierID(v string) int {
	return boilergql.IDToBoilerInt(v)
}

func SupplierIDs(a []string) []int {
	return boilergql.IDsToBoilerInt(a)
}
