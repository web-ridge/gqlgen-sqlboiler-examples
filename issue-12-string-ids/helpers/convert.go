package helpers

import (
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/null/v8"
)

type PostSort string

const (
	PostSortID	PostSort	= "id"
	PostSortContent	PostSort	= "content"
)

var PostSortDBValue = map[graphql_models.PostSort]PostSort{
	graphql_models.PostSortID:	PostSortID,
	graphql_models.PostSortContent:	PostSortContent,
}
var PostSortAPIValue = map[PostSort]graphql_models.PostSort{
	PostSortID:		graphql_models.PostSortID,
	PostSortContent:	graphql_models.PostSortContent,
}

func NullDotStringToPointerPostSort(v null.String) *graphql_models.PostSort {
	s := StringToPostSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToPostSort(v null.String) graphql_models.PostSort {
	if !v.Valid {
		return ""
	}
	return StringToPostSort(v.String)
}

func StringToPostSort(v string) graphql_models.PostSort {
	s := PostSortAPIValue[PostSort(v)]
	return s
}

func StringToPointerPostSort(v string) *graphql_models.PostSort {
	s := StringToPostSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerPostSortToString(v *graphql_models.PostSort) string {
	if v == nil {
		return ""
	}
	return PostSortToString(*v)
}

func PointerPostSortToNullDotString(v *graphql_models.PostSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return PostSortToNullDotString(*v)
}

func PostSortToNullDotString(v graphql_models.PostSort) null.String {
	s := PostSortToString(v)
	return null.NewString(s, s != "")
}

func PostSortToString(v graphql_models.PostSort) string {
	s := PostSortDBValue[v]
	return string(s)
}

type UserSort string

const (
	UserSortID		UserSort	= "id"
	UserSortFirstName	UserSort	= "firstName"
	UserSortLastName	UserSort	= "lastName"
	UserSortEmail		UserSort	= "email"
	UserSortPassword	UserSort	= "password"
)

var UserSortDBValue = map[graphql_models.UserSort]UserSort{
	graphql_models.UserSortID:		UserSortID,
	graphql_models.UserSortFirstName:	UserSortFirstName,
	graphql_models.UserSortLastName:	UserSortLastName,
	graphql_models.UserSortEmail:		UserSortEmail,
	graphql_models.UserSortPassword:	UserSortPassword,
}
var UserSortAPIValue = map[UserSort]graphql_models.UserSort{
	UserSortID:		graphql_models.UserSortID,
	UserSortFirstName:	graphql_models.UserSortFirstName,
	UserSortLastName:	graphql_models.UserSortLastName,
	UserSortEmail:		graphql_models.UserSortEmail,
	UserSortPassword:	graphql_models.UserSortPassword,
}

func NullDotStringToPointerUserSort(v null.String) *graphql_models.UserSort {
	s := StringToUserSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToUserSort(v null.String) graphql_models.UserSort {
	if !v.Valid {
		return ""
	}
	return StringToUserSort(v.String)
}

func StringToUserSort(v string) graphql_models.UserSort {
	s := UserSortAPIValue[UserSort(v)]
	return s
}

func StringToPointerUserSort(v string) *graphql_models.UserSort {
	s := StringToUserSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerUserSortToString(v *graphql_models.UserSort) string {
	if v == nil {
		return ""
	}
	return UserSortToString(*v)
}

func PointerUserSortToNullDotString(v *graphql_models.UserSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return UserSortToNullDotString(*v)
}

func UserSortToNullDotString(v graphql_models.UserSort) null.String {
	s := UserSortToString(v)
	return null.NewString(s, s != "")
}

func UserSortToString(v graphql_models.UserSort) string {
	s := UserSortDBValue[v]
	return string(s)
}

type ImageVariationSort string

const (
	ImageVariationSortID ImageVariationSort = "id"
)

var ImageVariationSortDBValue = map[graphql_models.ImageVariationSort]ImageVariationSort{
	graphql_models.ImageVariationSortID: ImageVariationSortID,
}
var ImageVariationSortAPIValue = map[ImageVariationSort]graphql_models.ImageVariationSort{
	ImageVariationSortID: graphql_models.ImageVariationSortID,
}

func NullDotStringToPointerImageVariationSort(v null.String) *graphql_models.ImageVariationSort {
	s := StringToImageVariationSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToImageVariationSort(v null.String) graphql_models.ImageVariationSort {
	if !v.Valid {
		return ""
	}
	return StringToImageVariationSort(v.String)
}

func StringToImageVariationSort(v string) graphql_models.ImageVariationSort {
	s := ImageVariationSortAPIValue[ImageVariationSort(v)]
	return s
}

func StringToPointerImageVariationSort(v string) *graphql_models.ImageVariationSort {
	s := StringToImageVariationSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerImageVariationSortToString(v *graphql_models.ImageVariationSort) string {
	if v == nil {
		return ""
	}
	return ImageVariationSortToString(*v)
}

func PointerImageVariationSortToNullDotString(v *graphql_models.ImageVariationSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return ImageVariationSortToNullDotString(*v)
}

func ImageVariationSortToNullDotString(v graphql_models.ImageVariationSort) null.String {
	s := ImageVariationSortToString(v)
	return null.NewString(s, s != "")
}

func ImageVariationSortToString(v graphql_models.ImageVariationSort) string {
	s := ImageVariationSortDBValue[v]
	return string(s)
}

type CommentLikeSort string

const (
	CommentLikeSortID		CommentLikeSort	= "id"
	CommentLikeSortLikeType		CommentLikeSort	= "likeType"
	CommentLikeSortCreatedAt	CommentLikeSort	= "createdAt"
)

var CommentLikeSortDBValue = map[graphql_models.CommentLikeSort]CommentLikeSort{
	graphql_models.CommentLikeSortID:		CommentLikeSortID,
	graphql_models.CommentLikeSortLikeType:		CommentLikeSortLikeType,
	graphql_models.CommentLikeSortCreatedAt:	CommentLikeSortCreatedAt,
}
var CommentLikeSortAPIValue = map[CommentLikeSort]graphql_models.CommentLikeSort{
	CommentLikeSortID:		graphql_models.CommentLikeSortID,
	CommentLikeSortLikeType:	graphql_models.CommentLikeSortLikeType,
	CommentLikeSortCreatedAt:	graphql_models.CommentLikeSortCreatedAt,
}

func NullDotStringToPointerCommentLikeSort(v null.String) *graphql_models.CommentLikeSort {
	s := StringToCommentLikeSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToCommentLikeSort(v null.String) graphql_models.CommentLikeSort {
	if !v.Valid {
		return ""
	}
	return StringToCommentLikeSort(v.String)
}

func StringToCommentLikeSort(v string) graphql_models.CommentLikeSort {
	s := CommentLikeSortAPIValue[CommentLikeSort(v)]
	return s
}

func StringToPointerCommentLikeSort(v string) *graphql_models.CommentLikeSort {
	s := StringToCommentLikeSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerCommentLikeSortToString(v *graphql_models.CommentLikeSort) string {
	if v == nil {
		return ""
	}
	return CommentLikeSortToString(*v)
}

func PointerCommentLikeSortToNullDotString(v *graphql_models.CommentLikeSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return CommentLikeSortToNullDotString(*v)
}

func CommentLikeSortToNullDotString(v graphql_models.CommentLikeSort) null.String {
	s := CommentLikeSortToString(v)
	return null.NewString(s, s != "")
}

func CommentLikeSortToString(v graphql_models.CommentLikeSort) string {
	s := CommentLikeSortDBValue[v]
	return string(s)
}

type LikeSort string

const (
	LikeSortID		LikeSort	= "id"
	LikeSortLikeType	LikeSort	= "likeType"
	LikeSortCreatedAt	LikeSort	= "createdAt"
)

var LikeSortDBValue = map[graphql_models.LikeSort]LikeSort{
	graphql_models.LikeSortID:		LikeSortID,
	graphql_models.LikeSortLikeType:	LikeSortLikeType,
	graphql_models.LikeSortCreatedAt:	LikeSortCreatedAt,
}
var LikeSortAPIValue = map[LikeSort]graphql_models.LikeSort{
	LikeSortID:		graphql_models.LikeSortID,
	LikeSortLikeType:	graphql_models.LikeSortLikeType,
	LikeSortCreatedAt:	graphql_models.LikeSortCreatedAt,
}

func NullDotStringToPointerLikeSort(v null.String) *graphql_models.LikeSort {
	s := StringToLikeSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToLikeSort(v null.String) graphql_models.LikeSort {
	if !v.Valid {
		return ""
	}
	return StringToLikeSort(v.String)
}

func StringToLikeSort(v string) graphql_models.LikeSort {
	s := LikeSortAPIValue[LikeSort(v)]
	return s
}

func StringToPointerLikeSort(v string) *graphql_models.LikeSort {
	s := StringToLikeSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerLikeSortToString(v *graphql_models.LikeSort) string {
	if v == nil {
		return ""
	}
	return LikeSortToString(*v)
}

func PointerLikeSortToNullDotString(v *graphql_models.LikeSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return LikeSortToNullDotString(*v)
}

func LikeSortToNullDotString(v graphql_models.LikeSort) null.String {
	s := LikeSortToString(v)
	return null.NewString(s, s != "")
}

func LikeSortToString(v graphql_models.LikeSort) string {
	s := LikeSortDBValue[v]
	return string(s)
}

type ImageSort string

const (
	ImageSortID		ImageSort	= "id"
	ImageSortViews		ImageSort	= "views"
	ImageSortOriginalURL	ImageSort	= "originalUrl"
)

var ImageSortDBValue = map[graphql_models.ImageSort]ImageSort{
	graphql_models.ImageSortID:		ImageSortID,
	graphql_models.ImageSortViews:		ImageSortViews,
	graphql_models.ImageSortOriginalURL:	ImageSortOriginalURL,
}
var ImageSortAPIValue = map[ImageSort]graphql_models.ImageSort{
	ImageSortID:		graphql_models.ImageSortID,
	ImageSortViews:		graphql_models.ImageSortViews,
	ImageSortOriginalURL:	graphql_models.ImageSortOriginalURL,
}

func NullDotStringToPointerImageSort(v null.String) *graphql_models.ImageSort {
	s := StringToImageSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToImageSort(v null.String) graphql_models.ImageSort {
	if !v.Valid {
		return ""
	}
	return StringToImageSort(v.String)
}

func StringToImageSort(v string) graphql_models.ImageSort {
	s := ImageSortAPIValue[ImageSort(v)]
	return s
}

func StringToPointerImageSort(v string) *graphql_models.ImageSort {
	s := StringToImageSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerImageSortToString(v *graphql_models.ImageSort) string {
	if v == nil {
		return ""
	}
	return ImageSortToString(*v)
}

func PointerImageSortToNullDotString(v *graphql_models.ImageSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return ImageSortToNullDotString(*v)
}

func ImageSortToNullDotString(v graphql_models.ImageSort) null.String {
	s := ImageSortToString(v)
	return null.NewString(s, s != "")
}

func ImageSortToString(v graphql_models.ImageSort) string {
	s := ImageSortDBValue[v]
	return string(s)
}

type CommentSort string

const (
	CommentSortID		CommentSort	= "id"
	CommentSortContent	CommentSort	= "content"
)

var CommentSortDBValue = map[graphql_models.CommentSort]CommentSort{
	graphql_models.CommentSortID:		CommentSortID,
	graphql_models.CommentSortContent:	CommentSortContent,
}
var CommentSortAPIValue = map[CommentSort]graphql_models.CommentSort{
	CommentSortID:		graphql_models.CommentSortID,
	CommentSortContent:	graphql_models.CommentSortContent,
}

func NullDotStringToPointerCommentSort(v null.String) *graphql_models.CommentSort {
	s := StringToCommentSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToCommentSort(v null.String) graphql_models.CommentSort {
	if !v.Valid {
		return ""
	}
	return StringToCommentSort(v.String)
}

func StringToCommentSort(v string) graphql_models.CommentSort {
	s := CommentSortAPIValue[CommentSort(v)]
	return s
}

func StringToPointerCommentSort(v string) *graphql_models.CommentSort {
	s := StringToCommentSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerCommentSortToString(v *graphql_models.CommentSort) string {
	if v == nil {
		return ""
	}
	return CommentSortToString(*v)
}

func PointerCommentSortToNullDotString(v *graphql_models.CommentSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return CommentSortToNullDotString(*v)
}

func CommentSortToNullDotString(v graphql_models.CommentSort) null.String {
	s := CommentSortToString(v)
	return null.NewString(s, s != "")
}

func CommentSortToString(v graphql_models.CommentSort) string {
	s := CommentSortDBValue[v]
	return string(s)
}

type FriendshipSort string

const (
	FriendshipSortID	FriendshipSort	= "id"
	FriendshipSortCreatedAt	FriendshipSort	= "createdAt"
)

var FriendshipSortDBValue = map[graphql_models.FriendshipSort]FriendshipSort{
	graphql_models.FriendshipSortID:	FriendshipSortID,
	graphql_models.FriendshipSortCreatedAt:	FriendshipSortCreatedAt,
}
var FriendshipSortAPIValue = map[FriendshipSort]graphql_models.FriendshipSort{
	FriendshipSortID:		graphql_models.FriendshipSortID,
	FriendshipSortCreatedAt:	graphql_models.FriendshipSortCreatedAt,
}

func NullDotStringToPointerFriendshipSort(v null.String) *graphql_models.FriendshipSort {
	s := StringToFriendshipSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToFriendshipSort(v null.String) graphql_models.FriendshipSort {
	if !v.Valid {
		return ""
	}
	return StringToFriendshipSort(v.String)
}

func StringToFriendshipSort(v string) graphql_models.FriendshipSort {
	s := FriendshipSortAPIValue[FriendshipSort(v)]
	return s
}

func StringToPointerFriendshipSort(v string) *graphql_models.FriendshipSort {
	s := StringToFriendshipSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerFriendshipSortToString(v *graphql_models.FriendshipSort) string {
	if v == nil {
		return ""
	}
	return FriendshipSortToString(*v)
}

func PointerFriendshipSortToNullDotString(v *graphql_models.FriendshipSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return FriendshipSortToNullDotString(*v)
}

func FriendshipSortToNullDotString(v graphql_models.FriendshipSort) null.String {
	s := FriendshipSortToString(v)
	return null.NewString(s, s != "")
}

func FriendshipSortToString(v graphql_models.FriendshipSort) string {
	s := FriendshipSortDBValue[v]
	return string(s)
}

func CommentWithStringID(id string) *graphql_models.Comment {
	return &graphql_models.Comment{
		ID: CommentIDToGraphQL(id),
	}
}

func CommentWithNullDotStringID(id null.String) *graphql_models.Comment {
	return CommentWithStringID(id.String)
}

func CommentsToGraphQL(am []*models.Comment) []*graphql_models.Comment {
	ar := make([]*graphql_models.Comment, len(am))
	for i, m := range am {
		ar[i] = CommentToGraphQL(m)
	}
	return ar
}

func CommentIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.Comment)
}

func CommentToGraphQL(m *models.Comment) *graphql_models.Comment {
	if m == nil {
		return nil
	}

	r := &graphql_models.Comment{
		ID:		m.ID,
		Content:	m.Content,
	}

	if boilergql.NullDotStringIsFilled(m.PostID) {
		if m.R != nil && m.R.Post != nil {
			r.Post = PostToGraphQL(m.R.Post)
		} else {
			r.Post = PostWithNullDotStringID(m.PostID)
		}
	}
	if boilergql.StringIsFilled(m.UserID) {
		if m.R != nil && m.R.User != nil {
			r.User = UserToGraphQL(m.R.User)
		} else {
			r.User = UserWithStringID(m.UserID)
		}
	}
	if m.R != nil && m.R.CommentLikes != nil {
		r.CommentLikes = CommentLikesToGraphQL(m.R.CommentLikes)
	}

	return r
}

func CommentID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func CommentIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}

func CommentLikeWithStringID(id string) *graphql_models.CommentLike {
	return &graphql_models.CommentLike{
		ID: CommentLikeIDToGraphQL(id),
	}
}

func CommentLikeWithNullDotStringID(id null.String) *graphql_models.CommentLike {
	return CommentLikeWithStringID(id.String)
}

func CommentLikesToGraphQL(am []*models.CommentLike) []*graphql_models.CommentLike {
	ar := make([]*graphql_models.CommentLike, len(am))
	for i, m := range am {
		ar[i] = CommentLikeToGraphQL(m)
	}
	return ar
}

func CommentLikeIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.CommentLike)
}

func CommentLikeToGraphQL(m *models.CommentLike) *graphql_models.CommentLike {
	if m == nil {
		return nil
	}

	r := &graphql_models.CommentLike{
		ID:	m.ID,

		LikeType:	m.LikeType,
		CreatedAt:	boilergql.NullDotTimeToPointerInt(m.CreatedAt),
	}

	if boilergql.StringIsFilled(m.CommentID) {
		if m.R != nil && m.R.Comment != nil {
			r.Comment = CommentToGraphQL(m.R.Comment)
		} else {
			r.Comment = CommentWithStringID(m.CommentID)
		}
	}
	if boilergql.StringIsFilled(m.UserID) {
		if m.R != nil && m.R.User != nil {
			r.User = UserToGraphQL(m.R.User)
		} else {
			r.User = UserWithStringID(m.UserID)
		}
	}

	return r
}

func CommentLikeID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func CommentLikeIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}

func FriendshipWithStringID(id string) *graphql_models.Friendship {
	return &graphql_models.Friendship{
		ID: FriendshipIDToGraphQL(id),
	}
}

func FriendshipWithNullDotStringID(id null.String) *graphql_models.Friendship {
	return FriendshipWithStringID(id.String)
}

func FriendshipsToGraphQL(am []*models.Friendship) []*graphql_models.Friendship {
	ar := make([]*graphql_models.Friendship, len(am))
	for i, m := range am {
		ar[i] = FriendshipToGraphQL(m)
	}
	return ar
}

func FriendshipIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.Friendship)
}

func FriendshipToGraphQL(m *models.Friendship) *graphql_models.Friendship {
	if m == nil {
		return nil
	}

	r := &graphql_models.Friendship{
		ID:		m.ID,
		CreatedAt:	boilergql.NullDotTimeToPointerInt(m.CreatedAt),
	}

	if m.R != nil && m.R.Users != nil {
		r.Users = UsersToGraphQL(m.R.Users)
	}

	return r
}

func FriendshipID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func FriendshipIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}

func ImageWithStringID(id string) *graphql_models.Image {
	return &graphql_models.Image{
		ID: ImageIDToGraphQL(id),
	}
}

func ImageWithNullDotStringID(id null.String) *graphql_models.Image {
	return ImageWithStringID(id.String)
}

func ImagesToGraphQL(am []*models.Image) []*graphql_models.Image {
	ar := make([]*graphql_models.Image, len(am))
	for i, m := range am {
		ar[i] = ImageToGraphQL(m)
	}
	return ar
}

func ImageIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.Image)
}

func ImageToGraphQL(m *models.Image) *graphql_models.Image {
	if m == nil {
		return nil
	}

	r := &graphql_models.Image{
		ID:	m.ID,

		Views:		boilergql.NullDotIntToPointerInt(m.Views),
		OriginalURL:	boilergql.NullDotStringToPointerString(m.OriginalURL),
	}

	if boilergql.StringIsFilled(m.PostID) {
		if m.R != nil && m.R.Post != nil {
			r.Post = PostToGraphQL(m.R.Post)
		} else {
			r.Post = PostWithStringID(m.PostID)
		}
	}
	if m.R != nil && m.R.ImageVariations != nil {
		r.ImageVariations = ImageVariationsToGraphQL(m.R.ImageVariations)
	}

	return r
}

func ImageID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func ImageIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}

func ImageVariationWithStringID(id string) *graphql_models.ImageVariation {
	return &graphql_models.ImageVariation{
		ID: ImageVariationIDToGraphQL(id),
	}
}

func ImageVariationWithNullDotStringID(id null.String) *graphql_models.ImageVariation {
	return ImageVariationWithStringID(id.String)
}

func ImageVariationsToGraphQL(am []*models.ImageVariation) []*graphql_models.ImageVariation {
	ar := make([]*graphql_models.ImageVariation, len(am))
	for i, m := range am {
		ar[i] = ImageVariationToGraphQL(m)
	}
	return ar
}

func ImageVariationIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.ImageVariation)
}

func ImageVariationToGraphQL(m *models.ImageVariation) *graphql_models.ImageVariation {
	if m == nil {
		return nil
	}

	r := &graphql_models.ImageVariation{
		ID: m.ID,
	}

	if boilergql.StringIsFilled(m.ImageID) {
		if m.R != nil && m.R.Image != nil {
			r.Image = ImageToGraphQL(m.R.Image)
		} else {
			r.Image = ImageWithStringID(m.ImageID)
		}
	}

	return r
}

func ImageVariationID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func ImageVariationIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}

func LikeWithStringID(id string) *graphql_models.Like {
	return &graphql_models.Like{
		ID: LikeIDToGraphQL(id),
	}
}

func LikeWithNullDotStringID(id null.String) *graphql_models.Like {
	return LikeWithStringID(id.String)
}

func LikesToGraphQL(am []*models.Like) []*graphql_models.Like {
	ar := make([]*graphql_models.Like, len(am))
	for i, m := range am {
		ar[i] = LikeToGraphQL(m)
	}
	return ar
}

func LikeIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.Like)
}

func LikeToGraphQL(m *models.Like) *graphql_models.Like {
	if m == nil {
		return nil
	}

	r := &graphql_models.Like{
		ID:	m.ID,

		LikeType:	m.LikeType,
		CreatedAt:	boilergql.NullDotTimeToPointerInt(m.CreatedAt),
	}

	if boilergql.StringIsFilled(m.PostID) {
		if m.R != nil && m.R.Post != nil {
			r.Post = PostToGraphQL(m.R.Post)
		} else {
			r.Post = PostWithStringID(m.PostID)
		}
	}
	if boilergql.StringIsFilled(m.UserID) {
		if m.R != nil && m.R.User != nil {
			r.User = UserToGraphQL(m.R.User)
		} else {
			r.User = UserWithStringID(m.UserID)
		}
	}

	return r
}

func LikeID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func LikeIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}

func PostWithStringID(id string) *graphql_models.Post {
	return &graphql_models.Post{
		ID: PostIDToGraphQL(id),
	}
}

func PostWithNullDotStringID(id null.String) *graphql_models.Post {
	return PostWithStringID(id.String)
}

func PostsToGraphQL(am []*models.Post) []*graphql_models.Post {
	ar := make([]*graphql_models.Post, len(am))
	for i, m := range am {
		ar[i] = PostToGraphQL(m)
	}
	return ar
}

func PostIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.Post)
}

func PostToGraphQL(m *models.Post) *graphql_models.Post {
	if m == nil {
		return nil
	}

	r := &graphql_models.Post{
		ID:		m.ID,
		Content:	m.Content,
	}

	if boilergql.StringIsFilled(m.UserID) {
		if m.R != nil && m.R.User != nil {
			r.User = UserToGraphQL(m.R.User)
		} else {
			r.User = UserWithStringID(m.UserID)
		}
	}
	if m.R != nil && m.R.Comments != nil {
		r.Comments = CommentsToGraphQL(m.R.Comments)
	}
	if m.R != nil && m.R.Images != nil {
		r.Images = ImagesToGraphQL(m.R.Images)
	}
	if m.R != nil && m.R.Likes != nil {
		r.Likes = LikesToGraphQL(m.R.Likes)
	}

	return r
}

func PostID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func PostIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}

func UserWithStringID(id string) *graphql_models.User {
	return &graphql_models.User{
		ID: UserIDToGraphQL(id),
	}
}

func UserWithNullDotStringID(id null.String) *graphql_models.User {
	return UserWithStringID(id.String)
}

func UsersToGraphQL(am []*models.User) []*graphql_models.User {
	ar := make([]*graphql_models.User, len(am))
	for i, m := range am {
		ar[i] = UserToGraphQL(m)
	}
	return ar
}

func UserIDToGraphQL(v string) string {
	return boilergql.StringIDToGraphQL(v, models.TableNames.User)
}

func UserToGraphQL(m *models.User) *graphql_models.User {
	if m == nil {
		return nil
	}

	r := &graphql_models.User{
		ID:		m.ID,
		FirstName:	m.FirstName,
		LastName:	m.LastName,
		Email:		m.Email,
		Password:	boilergql.ByteSliceToString(m.Password),
	}

	if m.R != nil && m.R.Comments != nil {
		r.Comments = CommentsToGraphQL(m.R.Comments)
	}
	if m.R != nil && m.R.CommentLikes != nil {
		r.CommentLikes = CommentLikesToGraphQL(m.R.CommentLikes)
	}
	if m.R != nil && m.R.Likes != nil {
		r.Likes = LikesToGraphQL(m.R.Likes)
	}
	if m.R != nil && m.R.Posts != nil {
		r.Posts = PostsToGraphQL(m.R.Posts)
	}
	if m.R != nil && m.R.Friendships != nil {
		r.Friendships = FriendshipsToGraphQL(m.R.Friendships)
	}

	return r
}

func UserID(v string) string {
	return boilergql.StringIDToBoilerString(v)
}

func UserIDs(a []string) []string {
	return boilergql.StringIDsToBoilerString(a)
}
