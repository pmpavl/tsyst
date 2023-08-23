package models

type TestCard struct {
	Path        string       `json:"path"`                  // Путь
	Name        string       `json:"name"`                  // Название
	Description string       `json:"description"`           // Описание
	Tags        *TestTags    `json:"tags"`                  // Теги теста
	LastPassage *TestPassage `json:"lastPassage,omitempty"` // Последнее прохождение
}

func NewTestCard(path, name, description string, tags *TestTags) *TestCard {
	return &TestCard{
		Path:        path,
		Name:        name,
		Description: description,
		Tags:        tags,
	}
}

func (t *TestCard) AddLastPassage(lastPassage *TestPassage) { t.LastPassage = lastPassage }
