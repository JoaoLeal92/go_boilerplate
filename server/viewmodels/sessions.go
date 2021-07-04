package viewmodels

// CreateSessionViewModel contains the data required to create a new session for a user
type CreateSessionViewModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
