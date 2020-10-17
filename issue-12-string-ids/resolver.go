// Generated with https://github.com/web-ridge/gqlgen-sqlboiler.
package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/auth"
	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/graphql_models"
	. "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/helpers"
	dm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/models"
	"github.com/web-ridge/utils-go/boilergql"
)

type Resolver struct {
	db *sql.DB
}

const inputKey = "input"

const publicCommentCreateError = "could not create comment"

func (r *mutationResolver) CreateComment(ctx context.Context, input fm.CommentCreateInput) (*fm.CommentPayload, error) {

	m := CommentCreateInputToBoiler(&input)

	m.UserID = auth.UserIDFromContext(ctx)

	whiteList := CommentCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),

		dm.CommentColumns.UserID,
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicCommentCreateError)
		return nil, errors.New(publicCommentCreateError)
	}

	// resolve requested fields after creating
	mods := GetCommentPreloadModsWithLevel(ctx, CommentPayloadPreloadLevels.Comment)
	mods = append(mods, dm.CommentWhere.ID.EQ(m.ID))

	mods = append(mods, dm.CommentWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	pM, err := dm.Comments(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentCreateError)
		return nil, errors.New(publicCommentCreateError)
	}
	return &fm.CommentPayload{
		Comment: CommentToGraphQL(pM),
	}, nil
}

//const publicCommentBatchCreateError = "could not create comments"

func (r *mutationResolver) CreateComments(ctx context.Context, input fm.CommentsCreateInput) (*fm.CommentsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicCommentUpdateError = "could not update comment"

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input fm.CommentUpdateInput) (*fm.CommentPayload, error) {
	m := CommentUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.Comments(
		dm.CommentWhere.ID.EQ(dbID),

		dm.CommentWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicCommentUpdateError)
		return nil, errors.New(publicCommentUpdateError)
	}

	// resolve requested fields after updating
	mods := GetCommentPreloadModsWithLevel(ctx, CommentPayloadPreloadLevels.Comment)
	mods = append(mods, dm.CommentWhere.ID.EQ(dbID))

	mods = append(mods, dm.CommentWhere.UserID.EQ(auth.UserIDFromContext(ctx)))
	pM, err := dm.Comments(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentUpdateError)
		return nil, errors.New(publicCommentUpdateError)
	}
	return &fm.CommentPayload{
		Comment: CommentToGraphQL(pM),
	}, nil
}

const publicCommentBatchUpdateError = "could not update comments"

func (r *mutationResolver) UpdateComments(ctx context.Context, filter *fm.CommentFilter, input fm.CommentUpdateInput) (*fm.CommentsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.CommentWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, CommentFilterToMods(filter)...)

	m := CommentUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Comments(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicCommentBatchUpdateError)
		return nil, errors.New(publicCommentBatchUpdateError)
	}

	return &fm.CommentsUpdatePayload{
		Ok: true,
	}, nil
}

const publicCommentDeleteError = "could not delete comment"

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*fm.CommentDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.CommentWhere.ID.EQ(dbID),

		dm.CommentWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	}
	if _, err := dm.Comments(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicCommentDeleteError)
		return nil, errors.New(publicCommentDeleteError)
	}

	return &fm.CommentDeletePayload{
		ID: id,
	}, nil
}

const publicCommentBatchDeleteError = "could not delete comments"

func (r *mutationResolver) DeleteComments(ctx context.Context, filter *fm.CommentFilter) (*fm.CommentsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.CommentWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, CommentFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.CommentColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Comment))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.Comments(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicCommentBatchDeleteError)
		return nil, errors.New(publicCommentBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.Comments(dm.CommentWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicCommentBatchDeleteError)
		return nil, errors.New(publicCommentBatchDeleteError)
	}
	return &fm.CommentsDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicCommentLikeCreateError = "could not create commentLike"

func (r *mutationResolver) CreateCommentLike(ctx context.Context, input fm.CommentLikeCreateInput) (*fm.CommentLikePayload, error) {

	m := CommentLikeCreateInputToBoiler(&input)

	m.UserID = auth.UserIDFromContext(ctx)

	whiteList := CommentLikeCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),

		dm.CommentLikeColumns.UserID,
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicCommentLikeCreateError)
		return nil, errors.New(publicCommentLikeCreateError)
	}

	// resolve requested fields after creating
	mods := GetCommentLikePreloadModsWithLevel(ctx, CommentLikePayloadPreloadLevels.CommentLike)
	mods = append(mods, dm.CommentLikeWhere.ID.EQ(m.ID))

	mods = append(mods, dm.CommentLikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	pM, err := dm.CommentLikes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentLikeCreateError)
		return nil, errors.New(publicCommentLikeCreateError)
	}
	return &fm.CommentLikePayload{
		CommentLike: CommentLikeToGraphQL(pM),
	}, nil
}

//const publicCommentLikeBatchCreateError = "could not create commentLikes"

func (r *mutationResolver) CreateCommentLikes(ctx context.Context, input fm.CommentLikesCreateInput) (*fm.CommentLikesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicCommentLikeUpdateError = "could not update commentLike"

func (r *mutationResolver) UpdateCommentLike(ctx context.Context, id string, input fm.CommentLikeUpdateInput) (*fm.CommentLikePayload, error) {
	m := CommentLikeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.CommentLikes(
		dm.CommentLikeWhere.ID.EQ(dbID),

		dm.CommentLikeWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicCommentLikeUpdateError)
		return nil, errors.New(publicCommentLikeUpdateError)
	}

	// resolve requested fields after updating
	mods := GetCommentLikePreloadModsWithLevel(ctx, CommentLikePayloadPreloadLevels.CommentLike)
	mods = append(mods, dm.CommentLikeWhere.ID.EQ(dbID))

	mods = append(mods, dm.CommentLikeWhere.UserID.EQ(auth.UserIDFromContext(ctx)))
	pM, err := dm.CommentLikes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentLikeUpdateError)
		return nil, errors.New(publicCommentLikeUpdateError)
	}
	return &fm.CommentLikePayload{
		CommentLike: CommentLikeToGraphQL(pM),
	}, nil
}

const publicCommentLikeBatchUpdateError = "could not update commentLikes"

func (r *mutationResolver) UpdateCommentLikes(ctx context.Context, filter *fm.CommentLikeFilter, input fm.CommentLikeUpdateInput) (*fm.CommentLikesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.CommentLikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, CommentLikeFilterToMods(filter)...)

	m := CommentLikeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.CommentLikes(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicCommentLikeBatchUpdateError)
		return nil, errors.New(publicCommentLikeBatchUpdateError)
	}

	return &fm.CommentLikesUpdatePayload{
		Ok: true,
	}, nil
}

const publicCommentLikeDeleteError = "could not delete commentLike"

func (r *mutationResolver) DeleteCommentLike(ctx context.Context, id string) (*fm.CommentLikeDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.CommentLikeWhere.ID.EQ(dbID),

		dm.CommentLikeWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	}
	if _, err := dm.CommentLikes(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicCommentLikeDeleteError)
		return nil, errors.New(publicCommentLikeDeleteError)
	}

	return &fm.CommentLikeDeletePayload{
		ID: id,
	}, nil
}

const publicCommentLikeBatchDeleteError = "could not delete commentLikes"

func (r *mutationResolver) DeleteCommentLikes(ctx context.Context, filter *fm.CommentLikeFilter) (*fm.CommentLikesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.CommentLikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, CommentLikeFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.CommentLikeColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.CommentLike))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.CommentLikes(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicCommentLikeBatchDeleteError)
		return nil, errors.New(publicCommentLikeBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.CommentLikes(dm.CommentLikeWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicCommentLikeBatchDeleteError)
		return nil, errors.New(publicCommentLikeBatchDeleteError)
	}
	return &fm.CommentLikesDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicFriendshipCreateError = "could not create friendship"

func (r *mutationResolver) CreateFriendship(ctx context.Context, input fm.FriendshipCreateInput) (*fm.FriendshipPayload, error) {

	m := FriendshipCreateInputToBoiler(&input)

	whiteList := FriendshipCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicFriendshipCreateError)
		return nil, errors.New(publicFriendshipCreateError)
	}

	// resolve requested fields after creating
	mods := GetFriendshipPreloadModsWithLevel(ctx, FriendshipPayloadPreloadLevels.Friendship)
	mods = append(mods, dm.FriendshipWhere.ID.EQ(m.ID))

	pM, err := dm.Friendships(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFriendshipCreateError)
		return nil, errors.New(publicFriendshipCreateError)
	}
	return &fm.FriendshipPayload{
		Friendship: FriendshipToGraphQL(pM),
	}, nil
}

//const publicFriendshipBatchCreateError = "could not create friendships"

func (r *mutationResolver) CreateFriendships(ctx context.Context, input fm.FriendshipsCreateInput) (*fm.FriendshipsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicFriendshipUpdateError = "could not update friendship"

func (r *mutationResolver) UpdateFriendship(ctx context.Context, id string, input fm.FriendshipUpdateInput) (*fm.FriendshipPayload, error) {
	m := FriendshipUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.Friendships(
		dm.FriendshipWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFriendshipUpdateError)
		return nil, errors.New(publicFriendshipUpdateError)
	}

	// resolve requested fields after updating
	mods := GetFriendshipPreloadModsWithLevel(ctx, FriendshipPayloadPreloadLevels.Friendship)
	mods = append(mods, dm.FriendshipWhere.ID.EQ(dbID))

	pM, err := dm.Friendships(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFriendshipUpdateError)
		return nil, errors.New(publicFriendshipUpdateError)
	}
	return &fm.FriendshipPayload{
		Friendship: FriendshipToGraphQL(pM),
	}, nil
}

const publicFriendshipBatchUpdateError = "could not update friendships"

func (r *mutationResolver) UpdateFriendships(ctx context.Context, filter *fm.FriendshipFilter, input fm.FriendshipUpdateInput) (*fm.FriendshipsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, FriendshipFilterToMods(filter)...)

	m := FriendshipUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Friendships(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicFriendshipBatchUpdateError)
		return nil, errors.New(publicFriendshipBatchUpdateError)
	}

	return &fm.FriendshipsUpdatePayload{
		Ok: true,
	}, nil
}

const publicFriendshipDeleteError = "could not delete friendship"

func (r *mutationResolver) DeleteFriendship(ctx context.Context, id string) (*fm.FriendshipDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.FriendshipWhere.ID.EQ(dbID),
	}
	if _, err := dm.Friendships(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicFriendshipDeleteError)
		return nil, errors.New(publicFriendshipDeleteError)
	}

	return &fm.FriendshipDeletePayload{
		ID: id,
	}, nil
}

const publicFriendshipBatchDeleteError = "could not delete friendships"

func (r *mutationResolver) DeleteFriendships(ctx context.Context, filter *fm.FriendshipFilter) (*fm.FriendshipsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, FriendshipFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.FriendshipColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Friendship))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.Friendships(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicFriendshipBatchDeleteError)
		return nil, errors.New(publicFriendshipBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.Friendships(dm.FriendshipWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicFriendshipBatchDeleteError)
		return nil, errors.New(publicFriendshipBatchDeleteError)
	}
	return &fm.FriendshipsDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicImageCreateError = "could not create image"

func (r *mutationResolver) CreateImage(ctx context.Context, input fm.ImageCreateInput) (*fm.ImagePayload, error) {

	m := ImageCreateInputToBoiler(&input)

	whiteList := ImageCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicImageCreateError)
		return nil, errors.New(publicImageCreateError)
	}

	// resolve requested fields after creating
	mods := GetImagePreloadModsWithLevel(ctx, ImagePayloadPreloadLevels.Image)
	mods = append(mods, dm.ImageWhere.ID.EQ(m.ID))

	pM, err := dm.Images(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageCreateError)
		return nil, errors.New(publicImageCreateError)
	}
	return &fm.ImagePayload{
		Image: ImageToGraphQL(pM),
	}, nil
}

//const publicImageBatchCreateError = "could not create images"

func (r *mutationResolver) CreateImages(ctx context.Context, input fm.ImagesCreateInput) (*fm.ImagesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicImageUpdateError = "could not update image"

func (r *mutationResolver) UpdateImage(ctx context.Context, id string, input fm.ImageUpdateInput) (*fm.ImagePayload, error) {
	m := ImageUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.Images(
		dm.ImageWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicImageUpdateError)
		return nil, errors.New(publicImageUpdateError)
	}

	// resolve requested fields after updating
	mods := GetImagePreloadModsWithLevel(ctx, ImagePayloadPreloadLevels.Image)
	mods = append(mods, dm.ImageWhere.ID.EQ(dbID))

	pM, err := dm.Images(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageUpdateError)
		return nil, errors.New(publicImageUpdateError)
	}
	return &fm.ImagePayload{
		Image: ImageToGraphQL(pM),
	}, nil
}

const publicImageBatchUpdateError = "could not update images"

func (r *mutationResolver) UpdateImages(ctx context.Context, filter *fm.ImageFilter, input fm.ImageUpdateInput) (*fm.ImagesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, ImageFilterToMods(filter)...)

	m := ImageUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Images(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicImageBatchUpdateError)
		return nil, errors.New(publicImageBatchUpdateError)
	}

	return &fm.ImagesUpdatePayload{
		Ok: true,
	}, nil
}

const publicImageDeleteError = "could not delete image"

func (r *mutationResolver) DeleteImage(ctx context.Context, id string) (*fm.ImageDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.ImageWhere.ID.EQ(dbID),
	}
	if _, err := dm.Images(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicImageDeleteError)
		return nil, errors.New(publicImageDeleteError)
	}

	return &fm.ImageDeletePayload{
		ID: id,
	}, nil
}

const publicImageBatchDeleteError = "could not delete images"

func (r *mutationResolver) DeleteImages(ctx context.Context, filter *fm.ImageFilter) (*fm.ImagesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, ImageFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.ImageColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Image))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.Images(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicImageBatchDeleteError)
		return nil, errors.New(publicImageBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.Images(dm.ImageWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicImageBatchDeleteError)
		return nil, errors.New(publicImageBatchDeleteError)
	}
	return &fm.ImagesDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicImageVariationCreateError = "could not create imageVariation"

func (r *mutationResolver) CreateImageVariation(ctx context.Context, input fm.ImageVariationCreateInput) (*fm.ImageVariationPayload, error) {

	m := ImageVariationCreateInputToBoiler(&input)

	whiteList := ImageVariationCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicImageVariationCreateError)
		return nil, errors.New(publicImageVariationCreateError)
	}

	// resolve requested fields after creating
	mods := GetImageVariationPreloadModsWithLevel(ctx, ImageVariationPayloadPreloadLevels.ImageVariation)
	mods = append(mods, dm.ImageVariationWhere.ID.EQ(m.ID))

	pM, err := dm.ImageVariations(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageVariationCreateError)
		return nil, errors.New(publicImageVariationCreateError)
	}
	return &fm.ImageVariationPayload{
		ImageVariation: ImageVariationToGraphQL(pM),
	}, nil
}

//const publicImageVariationBatchCreateError = "could not create imageVariations"

func (r *mutationResolver) CreateImageVariations(ctx context.Context, input fm.ImageVariationsCreateInput) (*fm.ImageVariationsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicImageVariationUpdateError = "could not update imageVariation"

func (r *mutationResolver) UpdateImageVariation(ctx context.Context, id string, input fm.ImageVariationUpdateInput) (*fm.ImageVariationPayload, error) {
	m := ImageVariationUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.ImageVariations(
		dm.ImageVariationWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicImageVariationUpdateError)
		return nil, errors.New(publicImageVariationUpdateError)
	}

	// resolve requested fields after updating
	mods := GetImageVariationPreloadModsWithLevel(ctx, ImageVariationPayloadPreloadLevels.ImageVariation)
	mods = append(mods, dm.ImageVariationWhere.ID.EQ(dbID))

	pM, err := dm.ImageVariations(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageVariationUpdateError)
		return nil, errors.New(publicImageVariationUpdateError)
	}
	return &fm.ImageVariationPayload{
		ImageVariation: ImageVariationToGraphQL(pM),
	}, nil
}

const publicImageVariationBatchUpdateError = "could not update imageVariations"

func (r *mutationResolver) UpdateImageVariations(ctx context.Context, filter *fm.ImageVariationFilter, input fm.ImageVariationUpdateInput) (*fm.ImageVariationsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, ImageVariationFilterToMods(filter)...)

	m := ImageVariationUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.ImageVariations(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicImageVariationBatchUpdateError)
		return nil, errors.New(publicImageVariationBatchUpdateError)
	}

	return &fm.ImageVariationsUpdatePayload{
		Ok: true,
	}, nil
}

const publicImageVariationDeleteError = "could not delete imageVariation"

func (r *mutationResolver) DeleteImageVariation(ctx context.Context, id string) (*fm.ImageVariationDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.ImageVariationWhere.ID.EQ(dbID),
	}
	if _, err := dm.ImageVariations(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicImageVariationDeleteError)
		return nil, errors.New(publicImageVariationDeleteError)
	}

	return &fm.ImageVariationDeletePayload{
		ID: id,
	}, nil
}

const publicImageVariationBatchDeleteError = "could not delete imageVariations"

func (r *mutationResolver) DeleteImageVariations(ctx context.Context, filter *fm.ImageVariationFilter) (*fm.ImageVariationsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, ImageVariationFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.ImageVariationColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.ImageVariation))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.ImageVariations(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicImageVariationBatchDeleteError)
		return nil, errors.New(publicImageVariationBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.ImageVariations(dm.ImageVariationWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicImageVariationBatchDeleteError)
		return nil, errors.New(publicImageVariationBatchDeleteError)
	}
	return &fm.ImageVariationsDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicLikeCreateError = "could not create like"

func (r *mutationResolver) CreateLike(ctx context.Context, input fm.LikeCreateInput) (*fm.LikePayload, error) {

	m := LikeCreateInputToBoiler(&input)

	m.UserID = auth.UserIDFromContext(ctx)

	whiteList := LikeCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),

		dm.LikeColumns.UserID,
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicLikeCreateError)
		return nil, errors.New(publicLikeCreateError)
	}

	// resolve requested fields after creating
	mods := GetLikePreloadModsWithLevel(ctx, LikePayloadPreloadLevels.Like)
	mods = append(mods, dm.LikeWhere.ID.EQ(m.ID))

	mods = append(mods, dm.LikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	pM, err := dm.Likes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLikeCreateError)
		return nil, errors.New(publicLikeCreateError)
	}
	return &fm.LikePayload{
		Like: LikeToGraphQL(pM),
	}, nil
}

//const publicLikeBatchCreateError = "could not create likes"

func (r *mutationResolver) CreateLikes(ctx context.Context, input fm.LikesCreateInput) (*fm.LikesPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicLikeUpdateError = "could not update like"

func (r *mutationResolver) UpdateLike(ctx context.Context, id string, input fm.LikeUpdateInput) (*fm.LikePayload, error) {
	m := LikeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.Likes(
		dm.LikeWhere.ID.EQ(dbID),

		dm.LikeWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLikeUpdateError)
		return nil, errors.New(publicLikeUpdateError)
	}

	// resolve requested fields after updating
	mods := GetLikePreloadModsWithLevel(ctx, LikePayloadPreloadLevels.Like)
	mods = append(mods, dm.LikeWhere.ID.EQ(dbID))

	mods = append(mods, dm.LikeWhere.UserID.EQ(auth.UserIDFromContext(ctx)))
	pM, err := dm.Likes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLikeUpdateError)
		return nil, errors.New(publicLikeUpdateError)
	}
	return &fm.LikePayload{
		Like: LikeToGraphQL(pM),
	}, nil
}

const publicLikeBatchUpdateError = "could not update likes"

func (r *mutationResolver) UpdateLikes(ctx context.Context, filter *fm.LikeFilter, input fm.LikeUpdateInput) (*fm.LikesUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.LikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, LikeFilterToMods(filter)...)

	m := LikeUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Likes(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicLikeBatchUpdateError)
		return nil, errors.New(publicLikeBatchUpdateError)
	}

	return &fm.LikesUpdatePayload{
		Ok: true,
	}, nil
}

const publicLikeDeleteError = "could not delete like"

func (r *mutationResolver) DeleteLike(ctx context.Context, id string) (*fm.LikeDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.LikeWhere.ID.EQ(dbID),

		dm.LikeWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	}
	if _, err := dm.Likes(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLikeDeleteError)
		return nil, errors.New(publicLikeDeleteError)
	}

	return &fm.LikeDeletePayload{
		ID: id,
	}, nil
}

const publicLikeBatchDeleteError = "could not delete likes"

func (r *mutationResolver) DeleteLikes(ctx context.Context, filter *fm.LikeFilter) (*fm.LikesDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.LikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, LikeFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.LikeColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Like))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.Likes(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicLikeBatchDeleteError)
		return nil, errors.New(publicLikeBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.Likes(dm.LikeWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicLikeBatchDeleteError)
		return nil, errors.New(publicLikeBatchDeleteError)
	}
	return &fm.LikesDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicPostCreateError = "could not create post"

func (r *mutationResolver) CreatePost(ctx context.Context, input fm.PostCreateInput) (*fm.PostPayload, error) {

	m := PostCreateInputToBoiler(&input)

	m.UserID = auth.UserIDFromContext(ctx)

	whiteList := PostCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),

		dm.PostColumns.UserID,
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicPostCreateError)
		return nil, errors.New(publicPostCreateError)
	}

	// resolve requested fields after creating
	mods := GetPostPreloadModsWithLevel(ctx, PostPayloadPreloadLevels.Post)
	mods = append(mods, dm.PostWhere.ID.EQ(m.ID))

	mods = append(mods, dm.PostWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	pM, err := dm.Posts(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicPostCreateError)
		return nil, errors.New(publicPostCreateError)
	}
	return &fm.PostPayload{
		Post: PostToGraphQL(pM),
	}, nil
}

//const publicPostBatchCreateError = "could not create posts"

func (r *mutationResolver) CreatePosts(ctx context.Context, input fm.PostsCreateInput) (*fm.PostsPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicPostUpdateError = "could not update post"

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input fm.PostUpdateInput) (*fm.PostPayload, error) {
	m := PostUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.Posts(
		dm.PostWhere.ID.EQ(dbID),

		dm.PostWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicPostUpdateError)
		return nil, errors.New(publicPostUpdateError)
	}

	// resolve requested fields after updating
	mods := GetPostPreloadModsWithLevel(ctx, PostPayloadPreloadLevels.Post)
	mods = append(mods, dm.PostWhere.ID.EQ(dbID))

	mods = append(mods, dm.PostWhere.UserID.EQ(auth.UserIDFromContext(ctx)))
	pM, err := dm.Posts(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicPostUpdateError)
		return nil, errors.New(publicPostUpdateError)
	}
	return &fm.PostPayload{
		Post: PostToGraphQL(pM),
	}, nil
}

const publicPostBatchUpdateError = "could not update posts"

func (r *mutationResolver) UpdatePosts(ctx context.Context, filter *fm.PostFilter, input fm.PostUpdateInput) (*fm.PostsUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.PostWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, PostFilterToMods(filter)...)

	m := PostUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Posts(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicPostBatchUpdateError)
		return nil, errors.New(publicPostBatchUpdateError)
	}

	return &fm.PostsUpdatePayload{
		Ok: true,
	}, nil
}

const publicPostDeleteError = "could not delete post"

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*fm.PostDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.PostWhere.ID.EQ(dbID),

		dm.PostWhere.UserID.EQ(auth.UserIDFromContext(ctx)),
	}
	if _, err := dm.Posts(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicPostDeleteError)
		return nil, errors.New(publicPostDeleteError)
	}

	return &fm.PostDeletePayload{
		ID: id,
	}, nil
}

const publicPostBatchDeleteError = "could not delete posts"

func (r *mutationResolver) DeletePosts(ctx context.Context, filter *fm.PostFilter) (*fm.PostsDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, dm.PostWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))
	mods = append(mods, PostFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.PostColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.Post))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.Posts(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicPostBatchDeleteError)
		return nil, errors.New(publicPostBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.Posts(dm.PostWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicPostBatchDeleteError)
		return nil, errors.New(publicPostBatchDeleteError)
	}
	return &fm.PostsDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicUserCreateError = "could not create user"

func (r *mutationResolver) CreateUser(ctx context.Context, input fm.UserCreateInput) (*fm.UserPayload, error) {

	m := UserCreateInputToBoiler(&input)

	whiteList := UserCreateInputToBoilerWhitelist(
		boilergql.GetInputFromContext(ctx, inputKey),
	)
	if err := m.Insert(ctx, r.db, whiteList); err != nil {
		log.Error().Err(err).Msg(publicUserCreateError)
		return nil, errors.New(publicUserCreateError)
	}

	// resolve requested fields after creating
	mods := GetUserPreloadModsWithLevel(ctx, UserPayloadPreloadLevels.User)
	mods = append(mods, dm.UserWhere.ID.EQ(m.ID))

	pM, err := dm.Users(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicUserCreateError)
		return nil, errors.New(publicUserCreateError)
	}
	return &fm.UserPayload{
		User: UserToGraphQL(pM),
	}, nil
}

//const publicUserBatchCreateError = "could not create users"

func (r *mutationResolver) CreateUsers(ctx context.Context, input fm.UsersCreateInput) (*fm.UsersPayload, error) {
	// TODO: Implement batch create
	return nil, nil
}

const publicUserUpdateError = "could not update user"

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input fm.UserUpdateInput) (*fm.UserPayload, error) {
	m := UserUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)

	dbID := id

	if _, err := dm.Users(
		dm.UserWhere.ID.EQ(dbID),
	).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicUserUpdateError)
		return nil, errors.New(publicUserUpdateError)
	}

	// resolve requested fields after updating
	mods := GetUserPreloadModsWithLevel(ctx, UserPayloadPreloadLevels.User)
	mods = append(mods, dm.UserWhere.ID.EQ(dbID))

	pM, err := dm.Users(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicUserUpdateError)
		return nil, errors.New(publicUserUpdateError)
	}
	return &fm.UserPayload{
		User: UserToGraphQL(pM),
	}, nil
}

const publicUserBatchUpdateError = "could not update users"

func (r *mutationResolver) UpdateUsers(ctx context.Context, filter *fm.UserFilter, input fm.UserUpdateInput) (*fm.UsersUpdatePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, UserFilterToMods(filter)...)

	m := UserUpdateInputToModelM(boilergql.GetInputFromContext(ctx, inputKey), input)
	if _, err := dm.Users(mods...).UpdateAll(ctx, r.db, m); err != nil {
		log.Error().Err(err).Msg(publicUserBatchUpdateError)
		return nil, errors.New(publicUserBatchUpdateError)
	}

	return &fm.UsersUpdatePayload{
		Ok: true,
	}, nil
}

const publicUserDeleteError = "could not delete user"

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*fm.UserDeletePayload, error) {
	dbID := id

	mods := []qm.QueryMod{
		dm.UserWhere.ID.EQ(dbID),
	}
	if _, err := dm.Users(mods...).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicUserDeleteError)
		return nil, errors.New(publicUserDeleteError)
	}

	return &fm.UserDeletePayload{
		ID: id,
	}, nil
}

const publicUserBatchDeleteError = "could not delete users"

func (r *mutationResolver) DeleteUsers(ctx context.Context, filter *fm.UserFilter) (*fm.UsersDeletePayload, error) {
	var mods []qm.QueryMod

	mods = append(mods, UserFilterToMods(filter)...)
	mods = append(mods, qm.Select(dm.UserColumns.ID))
	mods = append(mods, qm.From(dm.TableNames.User))
	var IDsToRemove []boilergql.RemovedStringID
	if err := dm.Users(mods...).Bind(ctx, r.db, &IDsToRemove); err != nil {
		log.Error().Err(err).Msg(publicUserBatchDeleteError)
		return nil, errors.New(publicUserBatchDeleteError)
	}

	boilerIDs := boilergql.RemovedIDsToBoilerString(IDsToRemove)
	if _, err := dm.Users(dm.UserWhere.ID.IN(boilerIDs)).DeleteAll(ctx, r.db); err != nil {
		log.Error().Err(err).Msg(publicUserBatchDeleteError)
		return nil, errors.New(publicUserBatchDeleteError)
	}
	return &fm.UsersDeletePayload{
		Ids: boilerIDs,
	}, nil
}

const publicCommentSingleError = "could not get comment"

func (r *queryResolver) Comment(ctx context.Context, id string) (*fm.Comment, error) {
	dbID := id

	mods := GetCommentPreloadMods(ctx)
	mods = append(mods, dm.CommentWhere.ID.EQ(dbID))

	mods = append(mods, dm.CommentWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	m, err := dm.Comments(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentSingleError)
		return nil, errors.New(publicCommentSingleError)
	}
	return CommentToGraphQL(m), nil
}

const publicCommentListError = "could not list comments"

func (r *queryResolver) Comments(ctx context.Context, filter *fm.CommentFilter) ([]*fm.Comment, error) {
	mods := GetCommentPreloadMods(ctx)

	mods = append(mods, dm.CommentWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	mods = append(mods, CommentFilterToMods(filter)...)
	a, err := dm.Comments(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentListError)
		return nil, errors.New(publicCommentListError)
	}
	return CommentsToGraphQL(a), nil
}

const publicCommentLikeSingleError = "could not get commentLike"

func (r *queryResolver) CommentLike(ctx context.Context, id string) (*fm.CommentLike, error) {
	dbID := id

	mods := GetCommentLikePreloadMods(ctx)
	mods = append(mods, dm.CommentLikeWhere.ID.EQ(dbID))

	mods = append(mods, dm.CommentLikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	m, err := dm.CommentLikes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentLikeSingleError)
		return nil, errors.New(publicCommentLikeSingleError)
	}
	return CommentLikeToGraphQL(m), nil
}

const publicCommentLikeListError = "could not list commentLikes"

func (r *queryResolver) CommentLikes(ctx context.Context, filter *fm.CommentLikeFilter) ([]*fm.CommentLike, error) {
	mods := GetCommentLikePreloadMods(ctx)

	mods = append(mods, dm.CommentLikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	mods = append(mods, CommentLikeFilterToMods(filter)...)
	a, err := dm.CommentLikes(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicCommentLikeListError)
		return nil, errors.New(publicCommentLikeListError)
	}
	return CommentLikesToGraphQL(a), nil
}

const publicFriendshipSingleError = "could not get friendship"

func (r *queryResolver) Friendship(ctx context.Context, id string) (*fm.Friendship, error) {
	dbID := id

	mods := GetFriendshipPreloadMods(ctx)
	mods = append(mods, dm.FriendshipWhere.ID.EQ(dbID))

	m, err := dm.Friendships(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFriendshipSingleError)
		return nil, errors.New(publicFriendshipSingleError)
	}
	return FriendshipToGraphQL(m), nil
}

const publicFriendshipListError = "could not list friendships"

func (r *queryResolver) Friendships(ctx context.Context, filter *fm.FriendshipFilter) ([]*fm.Friendship, error) {
	mods := GetFriendshipPreloadMods(ctx)

	mods = append(mods, FriendshipFilterToMods(filter)...)
	a, err := dm.Friendships(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicFriendshipListError)
		return nil, errors.New(publicFriendshipListError)
	}
	return FriendshipsToGraphQL(a), nil
}

const publicImageSingleError = "could not get image"

func (r *queryResolver) Image(ctx context.Context, id string) (*fm.Image, error) {
	dbID := id

	mods := GetImagePreloadMods(ctx)
	mods = append(mods, dm.ImageWhere.ID.EQ(dbID))

	m, err := dm.Images(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageSingleError)
		return nil, errors.New(publicImageSingleError)
	}
	return ImageToGraphQL(m), nil
}

const publicImageListError = "could not list images"

func (r *queryResolver) Images(ctx context.Context, filter *fm.ImageFilter) ([]*fm.Image, error) {
	mods := GetImagePreloadMods(ctx)

	mods = append(mods, ImageFilterToMods(filter)...)
	a, err := dm.Images(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageListError)
		return nil, errors.New(publicImageListError)
	}
	return ImagesToGraphQL(a), nil
}

const publicImageVariationSingleError = "could not get imageVariation"

func (r *queryResolver) ImageVariation(ctx context.Context, id string) (*fm.ImageVariation, error) {
	dbID := id

	mods := GetImageVariationPreloadMods(ctx)
	mods = append(mods, dm.ImageVariationWhere.ID.EQ(dbID))

	m, err := dm.ImageVariations(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageVariationSingleError)
		return nil, errors.New(publicImageVariationSingleError)
	}
	return ImageVariationToGraphQL(m), nil
}

const publicImageVariationListError = "could not list imageVariations"

func (r *queryResolver) ImageVariations(ctx context.Context, filter *fm.ImageVariationFilter) ([]*fm.ImageVariation, error) {
	mods := GetImageVariationPreloadMods(ctx)

	mods = append(mods, ImageVariationFilterToMods(filter)...)
	a, err := dm.ImageVariations(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicImageVariationListError)
		return nil, errors.New(publicImageVariationListError)
	}
	return ImageVariationsToGraphQL(a), nil
}

const publicLikeSingleError = "could not get like"

func (r *queryResolver) Like(ctx context.Context, id string) (*fm.Like, error) {
	dbID := id

	mods := GetLikePreloadMods(ctx)
	mods = append(mods, dm.LikeWhere.ID.EQ(dbID))

	mods = append(mods, dm.LikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	m, err := dm.Likes(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLikeSingleError)
		return nil, errors.New(publicLikeSingleError)
	}
	return LikeToGraphQL(m), nil
}

const publicLikeListError = "could not list likes"

func (r *queryResolver) Likes(ctx context.Context, filter *fm.LikeFilter) ([]*fm.Like, error) {
	mods := GetLikePreloadMods(ctx)

	mods = append(mods, dm.LikeWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	mods = append(mods, LikeFilterToMods(filter)...)
	a, err := dm.Likes(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicLikeListError)
		return nil, errors.New(publicLikeListError)
	}
	return LikesToGraphQL(a), nil
}

const publicPostSingleError = "could not get post"

func (r *queryResolver) Post(ctx context.Context, id string) (*fm.Post, error) {
	dbID := id

	mods := GetPostPreloadMods(ctx)
	mods = append(mods, dm.PostWhere.ID.EQ(dbID))

	mods = append(mods, dm.PostWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	m, err := dm.Posts(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicPostSingleError)
		return nil, errors.New(publicPostSingleError)
	}
	return PostToGraphQL(m), nil
}

const publicPostListError = "could not list posts"

func (r *queryResolver) Posts(ctx context.Context, filter *fm.PostFilter) ([]*fm.Post, error) {
	mods := GetPostPreloadMods(ctx)

	mods = append(mods, dm.PostWhere.UserID.EQ(
		auth.UserIDFromContext(ctx),
	))

	mods = append(mods, PostFilterToMods(filter)...)
	a, err := dm.Posts(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicPostListError)
		return nil, errors.New(publicPostListError)
	}
	return PostsToGraphQL(a), nil
}

const publicUserSingleError = "could not get user"

func (r *queryResolver) User(ctx context.Context, id string) (*fm.User, error) {
	dbID := id

	mods := GetUserPreloadMods(ctx)
	mods = append(mods, dm.UserWhere.ID.EQ(dbID))

	m, err := dm.Users(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicUserSingleError)
		return nil, errors.New(publicUserSingleError)
	}
	return UserToGraphQL(m), nil
}

const publicUserListError = "could not list users"

func (r *queryResolver) Users(ctx context.Context, filter *fm.UserFilter) ([]*fm.User, error) {
	mods := GetUserPreloadMods(ctx)

	mods = append(mods, UserFilterToMods(filter)...)
	a, err := dm.Users(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicUserListError)
		return nil, errors.New(publicUserListError)
	}
	return UsersToGraphQL(a), nil
}

func (r *Resolver) Mutation() fm.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() fm.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
