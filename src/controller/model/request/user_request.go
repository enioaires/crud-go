package request

type UserRequest struct {
	Email     string     `json:"email" binding:"required,email"`
	Password  string     `json:"password" binding:"required,min=6,max=50,containsany=@#$%^&+=_!?~"`
	Name      string     `json:"name" binding:"required,min=3,max=50"`
	Birthday  string     `json:"birthday" binding:"required"`
	IsActive  bool       `json:"is_active" binding:"required"`
	IsOver18  bool       `json:"is_over_18" binding:"required,eq=true"`
	Documents []Document `json:"documents" binding:"required,min=1,max=2"`
}

type Document struct {
	Type   string `json:"type" binding:"required"`
	Number string `json:"number" binding:"required,min=3,max=50"`
}
