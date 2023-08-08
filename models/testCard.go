package models

type TestCard struct {
	Path        string    `json:"path"`        // Путь
	Name        string    `json:"name"`        // Название
	Description string    `json:"description"` // Описание
	Tags        *TestTags `json:"tags"`        // Теги теста
}

func NewTestCard(path, name, description string, tags *TestTags) *TestCard {
	return &TestCard{
		Path:        path,
		Name:        name,
		Description: description,
		Tags:        tags,
	}
}
