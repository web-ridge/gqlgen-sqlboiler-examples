package helpers

import (
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CommentCreateInputsToBoiler(am []*graphql_models.CommentCreateInput) []*models.Comment {
	ar := make([]*models.Comment, len(am))
	for i, m := range am {
		ar[i] = CommentCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentCreateInputToBoiler(
	m *graphql_models.CommentCreateInput,
) *models.Comment {
	if m == nil {
		return nil
	}

	r := &models.Comment{
		Content: m.Content,
	}
	return r
}

func CommentCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentCreateInput,
) models.M {
	model := CommentCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "content":
			modelM[models.CommentColumns.Content] = model.Content
		case "postId":
			modelM[models.CommentColumns.PostID] = model.PostID
		}
	}
	return modelM
}

func CommentCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.Content)
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.PostID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func CommentLikeCreateInputsToBoiler(am []*graphql_models.CommentLikeCreateInput) []*models.CommentLike {
	ar := make([]*models.CommentLike, len(am))
	for i, m := range am {
		ar[i] = CommentLikeCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentLikeCreateInputToBoiler(
	m *graphql_models.CommentLikeCreateInput,
) *models.CommentLike {
	if m == nil {
		return nil
	}

	r := &models.CommentLike{

		LikeType:	m.LikeType,
		CreatedAt:	boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func CommentLikeCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentLikeCreateInput,
) models.M {
	model := CommentLikeCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "commentId":
			modelM[models.CommentLikeColumns.CommentID] = model.CommentID
		case "likeType":
			modelM[models.CommentLikeColumns.LikeType] = model.LikeType
		case "createdAt":
			modelM[models.CommentLikeColumns.CreatedAt] = model.CreatedAt
		}
	}
	return modelM
}

func CommentLikeCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "commentId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CommentID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func CommentLikeUpdateInputsToBoiler(am []*graphql_models.CommentLikeUpdateInput) []*models.CommentLike {
	ar := make([]*models.CommentLike, len(am))
	for i, m := range am {
		ar[i] = CommentLikeUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentLikeUpdateInputToBoiler(
	m *graphql_models.CommentLikeUpdateInput,
) *models.CommentLike {
	if m == nil {
		return nil
	}

	r := &models.CommentLike{

		LikeType:	boilergql.PointerStringToString(m.LikeType),
		CreatedAt:	boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func CommentLikeUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentLikeUpdateInput,
) models.M {
	model := CommentLikeUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "commentId":
			modelM[models.CommentLikeColumns.CommentID] = model.CommentID
		case "likeType":
			modelM[models.CommentLikeColumns.LikeType] = model.LikeType
		case "createdAt":
			modelM[models.CommentLikeColumns.CreatedAt] = model.CreatedAt
		}
	}
	return modelM
}

func CommentLikeUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "commentId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CommentID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func CommentUpdateInputsToBoiler(am []*graphql_models.CommentUpdateInput) []*models.Comment {
	ar := make([]*models.Comment, len(am))
	for i, m := range am {
		ar[i] = CommentUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentUpdateInputToBoiler(
	m *graphql_models.CommentUpdateInput,
) *models.Comment {
	if m == nil {
		return nil
	}

	r := &models.Comment{
		Content: boilergql.PointerStringToString(m.Content),
	}
	return r
}

func CommentUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentUpdateInput,
) models.M {
	model := CommentUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "content":
			modelM[models.CommentColumns.Content] = model.Content
		case "postId":
			modelM[models.CommentColumns.PostID] = model.PostID
		}
	}
	return modelM
}

func CommentUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.Content)
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.PostID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FriendshipCreateInputsToBoiler(am []*graphql_models.FriendshipCreateInput) []*models.Friendship {
	ar := make([]*models.Friendship, len(am))
	for i, m := range am {
		ar[i] = FriendshipCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func FriendshipCreateInputToBoiler(
	m *graphql_models.FriendshipCreateInput,
) *models.Friendship {
	if m == nil {
		return nil
	}

	r := &models.Friendship{
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func FriendshipCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.FriendshipCreateInput,
) models.M {
	model := FriendshipCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "createdAt":
			modelM[models.FriendshipColumns.CreatedAt] = model.CreatedAt
		}
	}
	return modelM
}

func FriendshipCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FriendshipColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FriendshipUpdateInputsToBoiler(am []*graphql_models.FriendshipUpdateInput) []*models.Friendship {
	ar := make([]*models.Friendship, len(am))
	for i, m := range am {
		ar[i] = FriendshipUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func FriendshipUpdateInputToBoiler(
	m *graphql_models.FriendshipUpdateInput,
) *models.Friendship {
	if m == nil {
		return nil
	}

	r := &models.Friendship{
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func FriendshipUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.FriendshipUpdateInput,
) models.M {
	model := FriendshipUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "createdAt":
			modelM[models.FriendshipColumns.CreatedAt] = model.CreatedAt
		}
	}
	return modelM
}

func FriendshipUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FriendshipColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageCreateInputsToBoiler(am []*graphql_models.ImageCreateInput) []*models.Image {
	ar := make([]*models.Image, len(am))
	for i, m := range am {
		ar[i] = ImageCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageCreateInputToBoiler(
	m *graphql_models.ImageCreateInput,
) *models.Image {
	if m == nil {
		return nil
	}

	r := &models.Image{

		Views:		boilergql.PointerIntToNullDotInt(m.Views),
		OriginalURL:	boilergql.PointerStringToNullDotString(m.OriginalURL),
	}
	return r
}

func ImageCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageCreateInput,
) models.M {
	model := ImageCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "postId":
			modelM[models.ImageColumns.PostID] = model.PostID
		case "views":
			modelM[models.ImageColumns.Views] = model.Views
		case "originalUrl":
			modelM[models.ImageColumns.OriginalURL] = model.OriginalURL
		}
	}
	return modelM
}

func ImageCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.PostID)
		case "views":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.Views)
		case "originalUrl":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.OriginalURL)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageUpdateInputsToBoiler(am []*graphql_models.ImageUpdateInput) []*models.Image {
	ar := make([]*models.Image, len(am))
	for i, m := range am {
		ar[i] = ImageUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageUpdateInputToBoiler(
	m *graphql_models.ImageUpdateInput,
) *models.Image {
	if m == nil {
		return nil
	}

	r := &models.Image{

		Views:		boilergql.PointerIntToNullDotInt(m.Views),
		OriginalURL:	boilergql.PointerStringToNullDotString(m.OriginalURL),
	}
	return r
}

func ImageUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageUpdateInput,
) models.M {
	model := ImageUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "postId":
			modelM[models.ImageColumns.PostID] = model.PostID
		case "views":
			modelM[models.ImageColumns.Views] = model.Views
		case "originalUrl":
			modelM[models.ImageColumns.OriginalURL] = model.OriginalURL
		}
	}
	return modelM
}

func ImageUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.PostID)
		case "views":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.Views)
		case "originalUrl":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.OriginalURL)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageVariationCreateInputsToBoiler(am []*graphql_models.ImageVariationCreateInput) []*models.ImageVariation {
	ar := make([]*models.ImageVariation, len(am))
	for i, m := range am {
		ar[i] = ImageVariationCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageVariationCreateInputToBoiler(
	m *graphql_models.ImageVariationCreateInput,
) *models.ImageVariation {
	if m == nil {
		return nil
	}

	r := &models.ImageVariation{}
	return r
}

func ImageVariationCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageVariationCreateInput,
) models.M {
	model := ImageVariationCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "imageId":
			modelM[models.ImageVariationColumns.ImageID] = model.ImageID
		}
	}
	return modelM
}

func ImageVariationCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "imageId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageVariationColumns.ImageID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageVariationUpdateInputsToBoiler(am []*graphql_models.ImageVariationUpdateInput) []*models.ImageVariation {
	ar := make([]*models.ImageVariation, len(am))
	for i, m := range am {
		ar[i] = ImageVariationUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageVariationUpdateInputToBoiler(
	m *graphql_models.ImageVariationUpdateInput,
) *models.ImageVariation {
	if m == nil {
		return nil
	}

	r := &models.ImageVariation{}
	return r
}

func ImageVariationUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageVariationUpdateInput,
) models.M {
	model := ImageVariationUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "imageId":
			modelM[models.ImageVariationColumns.ImageID] = model.ImageID
		}
	}
	return modelM
}

func ImageVariationUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "imageId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageVariationColumns.ImageID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LikeCreateInputsToBoiler(am []*graphql_models.LikeCreateInput) []*models.Like {
	ar := make([]*models.Like, len(am))
	for i, m := range am {
		ar[i] = LikeCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func LikeCreateInputToBoiler(
	m *graphql_models.LikeCreateInput,
) *models.Like {
	if m == nil {
		return nil
	}

	r := &models.Like{

		LikeType:	m.LikeType,
		CreatedAt:	boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func LikeCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.LikeCreateInput,
) models.M {
	model := LikeCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "postId":
			modelM[models.LikeColumns.PostID] = model.PostID
		case "likeType":
			modelM[models.LikeColumns.LikeType] = model.LikeType
		case "createdAt":
			modelM[models.LikeColumns.CreatedAt] = model.CreatedAt
		}
	}
	return modelM
}

func LikeCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.PostID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LikeUpdateInputsToBoiler(am []*graphql_models.LikeUpdateInput) []*models.Like {
	ar := make([]*models.Like, len(am))
	for i, m := range am {
		ar[i] = LikeUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func LikeUpdateInputToBoiler(
	m *graphql_models.LikeUpdateInput,
) *models.Like {
	if m == nil {
		return nil
	}

	r := &models.Like{

		LikeType:	boilergql.PointerStringToString(m.LikeType),
		CreatedAt:	boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func LikeUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.LikeUpdateInput,
) models.M {
	model := LikeUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "postId":
			modelM[models.LikeColumns.PostID] = model.PostID
		case "likeType":
			modelM[models.LikeColumns.LikeType] = model.LikeType
		case "createdAt":
			modelM[models.LikeColumns.CreatedAt] = model.CreatedAt
		}
	}
	return modelM
}

func LikeUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.PostID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func PostCreateInputsToBoiler(am []*graphql_models.PostCreateInput) []*models.Post {
	ar := make([]*models.Post, len(am))
	for i, m := range am {
		ar[i] = PostCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func PostCreateInputToBoiler(
	m *graphql_models.PostCreateInput,
) *models.Post {
	if m == nil {
		return nil
	}

	r := &models.Post{
		Content: m.Content,
	}
	return r
}

func PostCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.PostCreateInput,
) models.M {
	model := PostCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "content":
			modelM[models.PostColumns.Content] = model.Content
		}
	}
	return modelM
}

func PostCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.PostColumns.Content)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func PostUpdateInputsToBoiler(am []*graphql_models.PostUpdateInput) []*models.Post {
	ar := make([]*models.Post, len(am))
	for i, m := range am {
		ar[i] = PostUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func PostUpdateInputToBoiler(
	m *graphql_models.PostUpdateInput,
) *models.Post {
	if m == nil {
		return nil
	}

	r := &models.Post{
		Content: boilergql.PointerStringToString(m.Content),
	}
	return r
}

func PostUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.PostUpdateInput,
) models.M {
	model := PostUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "content":
			modelM[models.PostColumns.Content] = model.Content
		}
	}
	return modelM
}

func PostUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.PostColumns.Content)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func UserCreateInputsToBoiler(am []*graphql_models.UserCreateInput) []*models.User {
	ar := make([]*models.User, len(am))
	for i, m := range am {
		ar[i] = UserCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func UserCreateInputToBoiler(
	m *graphql_models.UserCreateInput,
) *models.User {
	if m == nil {
		return nil
	}

	r := &models.User{
		FirstName:	m.FirstName,
		LastName:	m.LastName,
		Email:		m.Email,
		Password:	boilergql.StringToByteSlice(m.Password),
	}
	return r
}

func UserCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.UserCreateInput,
) models.M {
	model := UserCreateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "firstName":
			modelM[models.UserColumns.FirstName] = model.FirstName
		case "lastName":
			modelM[models.UserColumns.LastName] = model.LastName
		case "email":
			modelM[models.UserColumns.Email] = model.Email
		case "password":
			modelM[models.UserColumns.Password] = model.Password
		}
	}
	return modelM
}

func UserCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "firstName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.FirstName)
		case "lastName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.LastName)
		case "email":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.Email)
		case "password":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.Password)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func UserUpdateInputsToBoiler(am []*graphql_models.UserUpdateInput) []*models.User {
	ar := make([]*models.User, len(am))
	for i, m := range am {
		ar[i] = UserUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func UserUpdateInputToBoiler(
	m *graphql_models.UserUpdateInput,
) *models.User {
	if m == nil {
		return nil
	}

	r := &models.User{
		FirstName:	boilergql.PointerStringToString(m.FirstName),
		LastName:	boilergql.PointerStringToString(m.LastName),
		Email:		boilergql.PointerStringToString(m.Email),
		Password:	boilergql.PointerStringToByteSlice(m.Password),
	}
	return r
}

func UserUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.UserUpdateInput,
) models.M {
	model := UserUpdateInputToBoiler(&m)
	modelM := models.M{}
	for key := range input {
		switch key {
		case "firstName":
			modelM[models.UserColumns.FirstName] = model.FirstName
		case "lastName":
			modelM[models.UserColumns.LastName] = model.LastName
		case "email":
			modelM[models.UserColumns.Email] = model.Email
		case "password":
			modelM[models.UserColumns.Password] = model.Password
		}
	}
	return modelM
}

func UserUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	var columnsWhichAreSet []string
	for key := range input {
		switch key {
		case "firstName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.FirstName)
		case "lastName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.LastName)
		case "email":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.Email)
		case "password":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.Password)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}
