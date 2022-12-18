package doauth

import (
	"context"
	"github.com/ory/fosite"
	"time"
)

type RedisStorage struct {
}

func (r RedisStorage) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisStorage) ClientAssertionJWTValid(ctx context.Context, jti string) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisStorage) SetClientAssertionJWT(ctx context.Context, jti string, exp time.Time) error {
	//TODO implement me
	panic("implement me")
}
