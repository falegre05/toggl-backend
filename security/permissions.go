package security

import (
	"errors"

	"github.com/graphql-go/graphql"
)

type permissionLevel int

const (
	// PermissionsPublic level basic level, without authentication
	plPublic permissionLevel = iota
	// PermissionsUser authentication as user
	plUser
)

var (
	// PermissionsPublic level basic level, without authentication
	PermissionsPublic = plPublic
	// PermissionsUser authentication as string, device and user
	PermissionsUser = plUser
)

// Check verify that minimum permission is met before executing the resolver
func Check(permission permissionLevel, successFn graphql.FieldResolveFn) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (res interface{}, err error) {
		var maxPermissionLevel = plPublic
		if GetUserID(params.Context) != nil {
			maxPermissionLevel = plUser
		}

		if permission > maxPermissionLevel {
			return nil, errors.New("unauthorized")
		}

		res, err = successFn(params)
		return
	}
}
