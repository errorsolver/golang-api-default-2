package entity

type Person struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname"`
	Age       int8   `json:"age" form:"age" binding:"gte=1,lte=100"`
	Email     string `json:"email" form:"email" binding:"email"`
}

type Video struct {
	Title  string `json:"title" form:"title" binding:"min=2,max=10" validate:"is-cool"`
	Desc   string `json:"desc" form:"desc" binding:"max=20"`
	URL    string `json:"url" form:"url" binding:"required,url"`
	Author Person `json:"author" form:"author" binding:"required"`
}
