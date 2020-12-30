package helpers

import (
	"context"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/social-network/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var TablePreloadMap = map[string]map[string]boilergql.ColumnSetting{
	models.TableNames.Comment: {
		"commentLikes": {
			Name:			models.CommentRels.CommentLikes,
			RelationshipModelName:	models.TableNames.CommentLike,
			IDAvailable:		false,
		},
		"post": {
			Name:			models.CommentRels.Post,
			RelationshipModelName:	models.TableNames.Post,
			IDAvailable:		true,
		},
		"user": {
			Name:			models.CommentRels.User,
			RelationshipModelName:	models.TableNames.User,
			IDAvailable:		true,
		},
	},
	models.TableNames.CommentLike: {
		"comment": {
			Name:			models.CommentLikeRels.Comment,
			RelationshipModelName:	models.TableNames.Comment,
			IDAvailable:		true,
		},
		"user": {
			Name:			models.CommentLikeRels.User,
			RelationshipModelName:	models.TableNames.User,
			IDAvailable:		true,
		},
	},
	models.TableNames.Friendship: {
		"users": {
			Name:			models.FriendshipRels.Users,
			RelationshipModelName:	models.TableNames.User,
			IDAvailable:		false,
		},
	},
	models.TableNames.Image: {
		"imageVariations": {
			Name:			models.ImageRels.ImageVariations,
			RelationshipModelName:	models.TableNames.ImageVariation,
			IDAvailable:		false,
		},
		"post": {
			Name:			models.ImageRels.Post,
			RelationshipModelName:	models.TableNames.Post,
			IDAvailable:		true,
		},
	},
	models.TableNames.ImageVariation: {
		"image": {
			Name:			models.ImageVariationRels.Image,
			RelationshipModelName:	models.TableNames.Image,
			IDAvailable:		true,
		},
	},
	models.TableNames.Like: {
		"post": {
			Name:			models.LikeRels.Post,
			RelationshipModelName:	models.TableNames.Post,
			IDAvailable:		true,
		},
		"user": {
			Name:			models.LikeRels.User,
			RelationshipModelName:	models.TableNames.User,
			IDAvailable:		true,
		},
	},
	models.TableNames.Post: {
		"comments": {
			Name:			models.PostRels.Comments,
			RelationshipModelName:	models.TableNames.Comment,
			IDAvailable:		false,
		},
		"images": {
			Name:			models.PostRels.Images,
			RelationshipModelName:	models.TableNames.Image,
			IDAvailable:		false,
		},
		"likes": {
			Name:			models.PostRels.Likes,
			RelationshipModelName:	models.TableNames.Like,
			IDAvailable:		false,
		},
		"user": {
			Name:			models.PostRels.User,
			RelationshipModelName:	models.TableNames.User,
			IDAvailable:		true,
		},
	},
	models.TableNames.User: {
		"commentLikes": {
			Name:			models.UserRels.CommentLikes,
			RelationshipModelName:	models.TableNames.CommentLike,
			IDAvailable:		false,
		},
		"comments": {
			Name:			models.UserRels.Comments,
			RelationshipModelName:	models.TableNames.Comment,
			IDAvailable:		false,
		},
		"friendships": {
			Name:			models.UserRels.Friendships,
			RelationshipModelName:	models.TableNames.Friendship,
			IDAvailable:		false,
		},
		"likes": {
			Name:			models.UserRels.Likes,
			RelationshipModelName:	models.TableNames.Like,
			IDAvailable:		false,
		},
		"posts": {
			Name:			models.UserRels.Posts,
			RelationshipModelName:	models.TableNames.Post,
			IDAvailable:		false,
		},
	},
}

func GetCommentPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Comment, "")
}
func GetCommentNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Comment, DefaultLevels.EdgesNode)
}
func GetCommentPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Comment, level)
}

func GetCommentLikePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.CommentLike, "")
}
func GetCommentLikeNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.CommentLike, DefaultLevels.EdgesNode)
}
func GetCommentLikePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.CommentLike, level)
}

func GetFriendshipPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Friendship, "")
}
func GetFriendshipNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Friendship, DefaultLevels.EdgesNode)
}
func GetFriendshipPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Friendship, level)
}

func GetImagePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Image, "")
}
func GetImageNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Image, DefaultLevels.EdgesNode)
}
func GetImagePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Image, level)
}

func GetImageVariationPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.ImageVariation, "")
}
func GetImageVariationNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.ImageVariation, DefaultLevels.EdgesNode)
}
func GetImageVariationPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.ImageVariation, level)
}

func GetLikePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Like, "")
}
func GetLikeNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Like, DefaultLevels.EdgesNode)
}
func GetLikePreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Like, level)
}

func GetPostPreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Post, "")
}
func GetPostNodePreloadMods(ctx context.Context) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Post, DefaultLevels.EdgesNode)
}
func GetPostPreloadModsWithLevel(ctx context.Context, level string) (queryMods []qm.QueryMod) {
	return boilergql.GetPreloadModsWithLevel(ctx, TablePreloadMap, models.TableNames.Post, level)
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

var CommentLikePayloadPreloadLevels = struct {
	CommentLike string
}{
	CommentLike: "commentLike",
}

var CommentPayloadPreloadLevels = struct {
	Comment string
}{
	Comment: "comment",
}

var FriendshipPayloadPreloadLevels = struct {
	Friendship string
}{
	Friendship: "friendship",
}

var ImagePayloadPreloadLevels = struct {
	Image string
}{
	Image: "image",
}

var ImageVariationPayloadPreloadLevels = struct {
	ImageVariation string
}{
	ImageVariation: "imageVariation",
}

var LikePayloadPreloadLevels = struct {
	Like string
}{
	Like: "like",
}

var PostPayloadPreloadLevels = struct {
	Post string
}{
	Post: "post",
}

var UserPayloadPreloadLevels = struct {
	User string
}{
	User: "user",
}

var DefaultLevels = struct {
	EdgesNode string
}{
	EdgesNode: "edges.node",
}
