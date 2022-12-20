package persister

import (
	"context"
	"fmt"
	"github.com/c1emon/lemontree/ent"
)

type EntOAuthClientRepository struct {
	client *ent.Client
}

func (r *EntOAuthClientRepository) SetClient(c any) {
	r.client = c.(*ent.Client)
}

func (r *EntOAuthClientRepository) Create(ctx context.Context, client ent.OAuthClient) (*ent.OAuthClient, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	c, err := tx.OAuthClient.Create().SetName(client.Name).Save(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("failed creating the group: %w", err))
	}
	err = tx.Organization.Update().Where().AddOauthClients(c).Exec(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("failed creating the group: %w", err))
	}

	return c, tx.Commit()
}

func (r *EntOAuthClientRepository) Delete(ctx context.Context, id string, client ent.OAuthClient) error {
	//TODO implement me
	panic("implement me")
}

func (r *EntOAuthClientRepository) Update(ctx context.Context, id string, client ent.OAuthClient) (*ent.OAuthClient, error) {
	//TODO implement me
	panic("implement me")
}

func (r *EntOAuthClientRepository) GetOneById(ctx context.Context, id string) (*ent.OAuthClient, error) {
	//TODO implement me
	panic("implement me")
}
