package test

import (
	"context"
	"fmt"
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/dao"
	"github.com/c1emon/lemontree/ent/department"
	"github.com/c1emon/lemontree/ent/organization"
	"github.com/c1emon/lemontree/log"
	"testing"
)

func Test_DbCreate(t *testing.T) {
	config.SetConfig(5432, "postgres", "host=10.0.0.70 port=5432 user=postgres dbname=lemon_tree password=123456 sslmode=disable")
	client := dao.GetEntClient()

	//client.Organization.Create().SetName("test1").SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())
	//client.Organization.Create().SetName("test2").SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())
	//client.Organization.Create().SetName("test3").SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())

	//client.Department.Create().SetName("o1d1").SetOrganizationID(1).SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())
	//client.Department.Create().SetName("o1d2").SetOrganizationID(1).SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())
	//
	//client.Department.Create().SetName("o2d1").SetOrganizationID(2).SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())
	//client.Department.Create().SetName("o2d2").SetOrganizationID(2).SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())
	//
	//client.Department.Create().SetName("o3d1").SetOrganizationID(3).SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())
	//client.Department.Create().SetName("o3d2").SetOrganizationID(3).SetUpdateTime(time.Now()).SetCreateTime(time.Now()).SaveX(context.Background())

	var oid = "asass"
	d, _ := client.Organization.Query().Where(organization.Oid(oid)).QueryDepartments().Where(department.ID(1)).Only(context.Background())
	fmt.Printf("d:\n%s", d)

	if err := client.Close(); err != nil {
		log.GetLogger().Warnf("unable close db client: %s", err)
	}
}
