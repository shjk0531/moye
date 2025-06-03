package constants

type UserPermissionName string

const (
	UserPermissionLoungeCreate UserPermissionName = "lounge:create"
	UserPermissionLoungeReadAll UserPermissionName = "lounge:read:all"
	UserPermissionLoungeUpdateOwn UserPermissionName = "lounge:update:own"
	UserPermissionLoungeDeleteOwn UserPermissionName = "lounge:delete:own"
	UserPermissionLoungeUpdateAll UserPermissionName = "lounge:update:all"
	UserPermissionLoungeDeleteAll UserPermissionName = "lounge:delete:all"
)
