package test

import (
	"context"
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/log"
	"github.com/c1emon/lemontree/model"
	"github.com/c1emon/lemontree/persister"
	"testing"
)

func Test_DbCreate(t *testing.T) {
	config.SetConfig(8080, "postgres", "host=10.0.0.70 port=5432 user=postgres dbname=lemon_tree password=123456 sslmode=disable")

	defer func() {
		if err := persister.DisConnect(); err != nil {
			log.GetLogger().Warnf("unable close db: %s", err)
		}
	}()

	r := persister.NewDefaultOrganizationRepository()
	r.CreateOne(context.Background(), &model.Organization{Name: "ta2aest"})
	r.CreateOne(context.Background(), &model.Organization{Name: "taaest"})

}
