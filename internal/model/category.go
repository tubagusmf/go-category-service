package model

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"article_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryUsecase interface {
	FindCategoryById(id string) (Category, error)
}

type CategoryRepository interface {
	FindCategoryById(id string) (Category, error)
}
