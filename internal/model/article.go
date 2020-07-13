package model

type Article struct {
	*Model
	Title	string	`json:"title"`
	Desc	string	`json:"desc"`
	Content	string	`json:"content"`
	ConverImageUrl	string	`json:"conver_image_url"`
	State	string	`json:"state"`
}

func (a Article) TableName() string {
	return "blog_acticle"
}