package models

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// struct for authentification
type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// struct for requests that dont require password
type UserResponse struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUserResponse(users []User) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return userResponses
}
