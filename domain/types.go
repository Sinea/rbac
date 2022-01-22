package domain

import "fmt"

const permissionSeparator = ":"

// User is a 'pointer' to a user. It's knowledge of the actual user is just it's ID
type User struct {
	ID    string
	Roles []Role
}

// Role encapsulates a list of permissions
type Role struct {
	Name        string
	Permissions []Permission
}

// Action that can be taken on a resource
type Action string

// Resource on which actions can be taken
type Resource string

// Permission represents an action that can be taken on a resource
type Permission struct {
	resource Resource
	action   Action
}

// Resource returns the resource of thie permission
func (p Permission) Resource() Resource {
	return p.resource
}

// Action returns the action of the permission
func (p Permission) Action() Action {
	return p.action
}

func (p Permission) String() string {
	return fmt.Sprintf("%s%s%s", p.resource, permissionSeparator, p.action)
}

// NewPermission creates a new permission
func NewPermission(resource Resource, action Action) Permission {
	return Permission{resource: resource, action: action}
}
