package doauth

import (
	"github.com/c1emon/lemontree/ent"
	"github.com/ory/fosite"
	"time"
)

type OAuthClient ent.OAuthClient

func (c *OAuthClient) GetID() string {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetHashedSecret() []byte {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetRedirectURIs() []string {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetGrantTypes() fosite.Arguments {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetResponseTypes() fosite.Arguments {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetScopes() fosite.Arguments {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) IsPublic() bool {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetAudience() fosite.Arguments {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetResponseModes() []fosite.ResponseModeType {
	//TODO implement me
	panic("implement me")
}

func (c *OAuthClient) GetEffectiveLifespan(gt fosite.GrantType, tt fosite.TokenType, fallback time.Duration) time.Duration {
	//TODO implement me
	panic("implement me")
}
