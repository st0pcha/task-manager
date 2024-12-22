package types

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserWithTasksResponse struct {
	User  UserResponse   `json:"user"`
	Tasks []TaskResponse `json:"tasks"`
}
