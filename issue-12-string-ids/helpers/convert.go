// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package helpers

import (
	null "github.com/volatiletech/null/v8"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/models"
	"github.com/web-ridge/utils-go/boilergql"
)

func CommentWithStringID(id string) *graphql_models.Comment {
	return &graphql_models.Comment{
		ID: id,
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

func CommentToGraphQL(m *models.Comment) *graphql_models.Comment {
	if m == nil {
		return nil
	}

	r := &graphql_models.Comment{
		ID:      m.ID,
		Content: m.Content,
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

func CommentLikeWithStringID(id string) *graphql_models.CommentLike {
	return &graphql_models.CommentLike{
		ID: id,
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

func CommentLikeToGraphQL(m *models.CommentLike) *graphql_models.CommentLike {
	if m == nil {
		return nil
	}

	r := &graphql_models.CommentLike{
		ID: m.ID,

		LikeType:  m.LikeType,
		CreatedAt: boilergql.NullDotTimeToPointerInt(m.CreatedAt),
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

func FriendshipWithStringID(id string) *graphql_models.Friendship {
	return &graphql_models.Friendship{
		ID: id,
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

func FriendshipToGraphQL(m *models.Friendship) *graphql_models.Friendship {
	if m == nil {
		return nil
	}

	r := &graphql_models.Friendship{
		ID:        m.ID,
		CreatedAt: boilergql.NullDotTimeToPointerInt(m.CreatedAt),
	}

	if m.R != nil && m.R.Users != nil {
		r.Users = UsersToGraphQL(m.R.Users)
	}

	return r
}

func ImageWithStringID(id string) *graphql_models.Image {
	return &graphql_models.Image{
		ID: id,
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

func ImageToGraphQL(m *models.Image) *graphql_models.Image {
	if m == nil {
		return nil
	}

	r := &graphql_models.Image{
		ID: m.ID,

		Views:       boilergql.NullDotIntToPointerInt(m.Views),
		OriginalURL: boilergql.NullDotStringToPointerString(m.OriginalURL),
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

func ImageVariationWithStringID(id string) *graphql_models.ImageVariation {
	return &graphql_models.ImageVariation{
		ID: id,
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

func LikeWithStringID(id string) *graphql_models.Like {
	return &graphql_models.Like{
		ID: id,
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

func LikeToGraphQL(m *models.Like) *graphql_models.Like {
	if m == nil {
		return nil
	}

	r := &graphql_models.Like{
		ID: m.ID,

		LikeType:  m.LikeType,
		CreatedAt: boilergql.NullDotTimeToPointerInt(m.CreatedAt),
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

func PostWithStringID(id string) *graphql_models.Post {
	return &graphql_models.Post{
		ID: id,
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

func PostToGraphQL(m *models.Post) *graphql_models.Post {
	if m == nil {
		return nil
	}

	r := &graphql_models.Post{
		ID:      m.ID,
		Content: m.Content,
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

func UserWithStringID(id string) *graphql_models.User {
	return &graphql_models.User{
		ID: id,
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

func UserToGraphQL(m *models.User) *graphql_models.User {
	if m == nil {
		return nil
	}

	r := &graphql_models.User{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Email:     m.Email,
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
