//go:generate gorunpkg github.com/99designs/gqlgen

package graphql

import (
	context "context"

	"github.com/koneko096/blacktube-graphql/manager"
	models "github.com/koneko096/blacktube-graphql/models"
)

type Resolver struct {
	UserManager  *manager.UserQueryManager
	VideoManager *manager.VideoQueryManager
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (models.User, error) {
	return r.UserManager.NewUser(input)
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, mutation models.UserMutation) (models.User, error) {
	u := models.User{
		ID:       id,
		Name:     *mutation.Name,
		Password: *mutation.Password,
	}
	return r.UserManager.UpdateUser(u)
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (models.User, error) {
	return r.UserManager.DeleteUser(id)
}
func (r *mutationResolver) CreateVideo(ctx context.Context, input models.NewVideo) (models.VideoNested, error) {
	return r.VideoManager.NewVideo(input)
}
func (r *mutationResolver) DeleteVideo(ctx context.Context, id int) (models.VideoNested, error) {
	return r.VideoManager.DeleteVideo(id)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	return r.UserManager.AllUsers()
}
func (r *queryResolver) Videos(ctx context.Context) ([]models.VideoNested, error) {
	return r.VideoManager.AllVideos()
}
