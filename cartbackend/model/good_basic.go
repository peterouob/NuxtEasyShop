package model

type Good struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Img         string `json:"img"`
	Description string `json:"description"`
}

func (g *Good) TableName() string {
	return "good"
}
