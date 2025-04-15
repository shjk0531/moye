package constants

type UserPermissionName string

const (
	UserPermissionStudyCreate UserPermissionName = "study:create"
	UserPermissionStudyReadAll UserPermissionName = "study:read:all"
	UserPermissionStudyUpdateOwn UserPermissionName = "study:update:own"
	UserPermissionStudyDeleteOwn UserPermissionName = "study:delete:own"
	UserPermissionStudyUpdateAll UserPermissionName = "study:update:all"
	UserPermissionStudyDeleteAll UserPermissionName = "study:delete:all"
)
