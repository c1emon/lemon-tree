package persister

import (
	"context"
	"github.com/c1emon/lemontree/ent"
)

type EntUserRepository struct {
	client *ent.Client
}

func (r *EntUserRepository) SetClient(c any) {
	r.client = c.(*ent.Client)
}

func (r *EntUserRepository) Create(ctx context.Context, user ent.User) (*ent.User, error) {
	return r.client.User.Create().
		SetName(user.Name).
		SetUsername(user.Username).
		Save(ctx)
}

func (r *EntUserRepository) Delete(ctx context.Context, id string, user ent.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *EntUserRepository) Update(ctx context.Context, id string, user ent.User) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *EntUserRepository) GetOneById(ctx context.Context, id string) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}
