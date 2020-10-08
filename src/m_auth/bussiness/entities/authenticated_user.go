package authbussinessentities

// AuthenticatedUser ...
type AuthenticatedUser struct {
	ID           int64
	Name         string
	Email        string
	PermissionID int64
}
