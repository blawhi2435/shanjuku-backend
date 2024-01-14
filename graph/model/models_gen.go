// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateGroupInput struct {
	Name string `json:"name"`
}

type CreateGroupPayload struct {
	Group *Group `json:"group"`
}

type DeleteGroupInput struct {
	ID string `json:"id"`
}

type DeleteGroupPayload struct {
	Success bool `json:"success"`
}

type EditGroupInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EditGroupPayload struct {
	Group *Group `json:"group,omitempty"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type InviteUserInput struct {
	GroupID string   `json:"groupId"`
	UserIds []string `json:"userIds"`
}

type InviteUserPayload struct {
	Users []*User `json:"users"`
}

type LoginInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginPayload struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type LogoutInput struct {
	Account string `json:"account"`
}

type LogoutPayload struct {
	Success bool `json:"success"`
}

type Mutation struct {
}

type Query struct {
}

type RegisterInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type RemoveUserInput struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

type RemoveUserPayload struct {
	Success bool `json:"success"`
}

type User struct {
	ID      string `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
}
