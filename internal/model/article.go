package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Context       string `json:"context"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article_tag"
}