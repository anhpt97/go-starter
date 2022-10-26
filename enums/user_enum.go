package enums

type UserRole string

type role struct {
	Admin UserRole
	User  UserRole
}

type UserStatus string

type status struct {
	NotActivated UserStatus
	Active       UserStatus
	IsDisabled   UserStatus
}

var User = struct {
	Role   role
	Status status
}{
	Role: role{
		Admin: "ADMIN",
		User:  "USER",
	},
	Status: status{
		NotActivated: "NOT_ACTIVATED",
		Active:       "ACTIVE",
		IsDisabled:   "IS_DISABLED",
	},
}
