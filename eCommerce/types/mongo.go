package types

// -> schema

type User struct {
	User   string   `json:"user" bson:"user"`
	Bucket []string `json:"bucket" bson:"bucket"`
}

type Content struct {
	Name      string `json:"name" bson:"name"`
	Price     int64  `json:"price" bson:"price"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
}

type History struct {
	User        string   `json:"user" bson:"user"`
	ContentList []string `json:"contentList" bson:"contentList"`
	CreatedAt   int64    `json:"createdAt" bson:"createdAt"`
}

// -> request

type UserRequest struct {
	User string `form:"user" binding:"required"`
}

type ContentRequest struct {
	Name string `form:"content"`
}

type CreateUserRequest struct {
	User string `json:"user" binding:"required"`
}

type CreateContentRequest struct {
	Name  string `json:"content" binding:"required"`
	Price int64  `json:"price" binding:"required"`
}

type BuyRequest struct {
	User string `json:"user" binding:"required"`
}

type CreateBucketRequest struct {
	User    string `json:"user" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type GeneralResponse struct {
	ResultCode  int         `json:"resultCode"`
	Description string      `json:"description"`
	ErrCode     int         `json:"errCode"`
	Result      interface{} `json:"result"`
}
