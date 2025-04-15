package constants

type UserRoleName string

const (
	UserRoleAdmin UserRoleName = "admin"
	UserRoleUser  UserRoleName = "user"
	UserRoleGuest UserRoleName = "guest"
	UserRoleBlock UserRoleName = "block"
	UserRoleEditor UserRoleName = "editor"
	UserRoleWriter UserRoleName = "writer"
	UserRoleReader UserRoleName = "reader"
)
