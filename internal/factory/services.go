package factory

import (
	"sync"

	"github.com/c1emon/lemontree/internal/org"
	"github.com/c1emon/lemontree/internal/user"
)

var onceUserService = sync.Once{}
var UserService *user.UserService

func GetUserService() *user.UserService {
	onceUserService.Do(func() {
		UserService = user.NewUserService()
	})

	return UserService
}

var onceUserIdentityService = sync.Once{}
var userIdentityService *user.UserIdentityService

func GetUserIdentityService() *user.UserIdentityService {
	onceUserIdentityService.Do(func() {
		userIdentityService = user.NewUserIdentityService()
	})

	return userIdentityService
}

var onceOrganizationService = sync.Once{}
var orgService *org.OrganizationService

func GetOrganizationService() *org.OrganizationService {
	onceOrganizationService.Do(func() {
		orgService = org.NewOrganizationService()
	})

	return orgService
}
