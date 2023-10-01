package test

import (
	"context"
	"testing"

	"github.com/c1emon/lemontree/internal/oidcx/client"
)

func Test_ClientCreate(t *testing.T) {
	start()
	defer stop()

	svc := client.NewOidcClientService()
	svc.Create(context.Background())

}
