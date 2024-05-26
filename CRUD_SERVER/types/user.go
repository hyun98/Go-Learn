package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// Response

type GetUserResponse struct {
	ApiResponse *ApiResponse `json:"ApiResponse"`
	User        []*User      `json:"users"`
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserResponse struct {
	*ApiResponse
}

// Request

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type UpdateUserRequest struct {
	Name       string `json:"name" binding:"required"`
	UpdatedAge int64  `json:"updatedAge" binding:"required"`
}

func (c *UpdateUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.UpdatedAge,
	}
}

type DeleteUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *DeleteUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  0,
	}
}
