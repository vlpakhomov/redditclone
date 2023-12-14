package repository

import (
	"context"
	"fmt"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type postRepositoryMongoDB struct {
	client *mongo.Client
	posts  *mongo.Collection
}

var _ interfaces.IPostRepository = (*postRepositoryMongoDB)(nil)

func NewPostRepositoryMongoDB(client *mongo.Client, col *mongo.Collection) *postRepositoryMongoDB {
	return &postRepositoryMongoDB{
		client: client,
		posts:  col,
	}
}

func (r *postRepositoryMongoDB) Add(ctx context.Context, post entity.PostExtend) error {
	_, errInsertOne := r.posts.InsertOne(ctx, post)
	if errInsertOne != nil {
		return fmt.Errorf("[postRepositoryMongoDB.Add]: %w", errInsertOne)
	}

	return nil
}

func (r *postRepositoryMongoDB) Get(ctx context.Context, postID string) (entity.PostExtend, error) {
	filter := bson.M{"id": postID}
	res := r.posts.FindOne(ctx, filter)
	if res.Err() == mongo.ErrNoDocuments {
		return entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.Get]: %w", interfaces.ErrPostNotExists)
	}
	if res.Err() != nil {
		return entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.Get]: %w", res.Err())
	}

	post := entity.PostExtend{}
	errDecode := res.Decode(&post)
	if errDecode != nil {
		return entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.Get]: %w", errDecode)
	}

	return post, nil
}

func (r *postRepositoryMongoDB) GetWhereCategory(ctx context.Context, category string) ([]entity.PostExtend, error) {
	filter := bson.M{"post.category": category}
	cur, errFind := r.posts.Find(ctx, filter)
	if errFind != nil {
		return []entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetWhereCategory]: %w", errFind)
	}

	posts := []entity.PostExtend{}
	errAll := cur.All(ctx, &posts)
	if errAll != nil {
		return []entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetWhereCategory]: %w", errAll)
	}

	return posts, nil
}

func (r *postRepositoryMongoDB) GetWhereUsername(ctx context.Context, username string) ([]entity.PostExtend, error) {
	filter := bson.M{"post.author.username": username}
	cur, errFind := r.posts.Find(ctx, filter)
	if errFind != nil {
		return []entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetWhereUsername]: %w", errFind)
	}

	posts := []entity.PostExtend{}
	errAll := cur.All(ctx, &posts)
	if errAll != nil {
		return []entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetWhereUsername]: %w", errAll)
	}

	return posts, nil
}

func (r *postRepositoryMongoDB) GetAll(ctx context.Context) ([]entity.PostExtend, error) {
	filter := bson.M{}
	cur, errFind := r.posts.Find(ctx, filter)
	if errFind != nil {
		return []entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetAll]: %w", errFind)
	}

	posts := []entity.PostExtend{}
	errAll := cur.All(ctx, &posts)
	if errAll != nil {
		return []entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetAll]: %w", errAll)
	}

	return posts, nil
}

func (r *postRepositoryMongoDB) Update(ctx context.Context, postID string, newPost entity.PostExtend) error {
	filter := bson.M{"id": postID}

	res, errUpdateOne := r.posts.UpdateOne(ctx, filter, bson.M{"$set": newPost})
	if errUpdateOne != nil {
		return fmt.Errorf("[postRepositoryMongoDB.Update]: %w", errUpdateOne)
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("[postRepositoryMongoDB.Update]: %w", interfaces.ErrPostNotExists)
	}

	return nil
}

func (r *postRepositoryMongoDB) Delete(ctx context.Context, postID string) error {
	filter := bson.M{"id": postID}
	res, errDelete := r.posts.DeleteOne(ctx, filter)
	if errDelete != nil {
		return fmt.Errorf("[postRepositoryMongoDB.Delete]: %w", errDelete)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("[postRepositoryMongoDB.Update]: %w", interfaces.ErrPostNotExists)
	}

	return nil
}

func (r *postRepositoryMongoDB) AddComment(ctx context.Context, postID string, comment entity.CommentExtend) (entity.PostExtend, error) {
	newPost, errGetPost := r.Get(ctx, postID)
	if errGetPost != nil {
		return entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.AddComment]: %w", errGetPost)
	}

	newPost.Comments = append(newPost.Comments, comment)
	errUpdatePost := r.Update(ctx, postID, newPost)
	if errUpdatePost != nil {
		return entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.AddComment]: %w", errUpdatePost)
	}
	return newPost, nil
}

func (r *postRepositoryMongoDB) GetComment(ctx context.Context, postID string, commentID string) (entity.CommentExtend, error) {
	post, errGetPost := r.Get(ctx, postID)
	if errGetPost != nil {
		return entity.CommentExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetComment]->%w", errGetPost)
	}

	for i := range post.Comments {
		if post.Comments[i].ID == commentID {
			return post.Comments[i], nil
		}
	}

	return entity.CommentExtend{}, fmt.Errorf("[postRepositoryMongoDB.GetComment]: %w", interfaces.ErrCommentNotExists)
}

func (r *postRepositoryMongoDB) DeleteComment(ctx context.Context, postID string, commentID string) (entity.PostExtend, error) {
	post, errGetPost := r.Get(ctx, postID)
	if errGetPost != nil {
		return entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.DeleteComment]: %w", errGetPost)
	}

	for i := range post.Comments {
		if post.Comments[i].ID == commentID {
			copy(post.Comments[i:], post.Comments[i+1:])
			post.Comments[len(post.Comments)-1] = entity.CommentExtend{}
			post.Comments = post.Comments[:len(post.Comments)-1]
			r.Update(ctx, postID, post)
			return post, nil
		}
	}

	return entity.PostExtend{}, fmt.Errorf("[postRepositoryMongoDB.DeleteComment]: %w", interfaces.ErrCommentNotExists)
}
