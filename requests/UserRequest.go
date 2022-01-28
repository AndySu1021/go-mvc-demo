package requests

type CreateUserParam struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required"`
}

type GetUserListParam struct {
	Username string `form:"username"`
	Page int `form:"page" binding:"required,gte=1"`
	PageSize int `form:"page_size" binding:"required,gte=10"`
}

type UpdateUserParam struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required"`
	Status   bool   `json:"status" binding:"required"`
}
