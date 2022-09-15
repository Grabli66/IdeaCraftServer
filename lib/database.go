package database

import (
	"log"
	"time"
)

type GetIdeasFilter struct {
	StartDate  time.Time
	Count      int
	IsFavorite bool
	UserId     *int
}

type IDatabase interface {
	// Инициализирует
	Init()
	// Возвращает идею по идентификатору
	GetIdeaById(id int) IdeaDB
	// Возвращает идеи
	GetIdeas(filter GetIdeasFilter) []IdeaDB
}

// Идея базы данных
type IdeaDB struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Desctiption string `json:"desctiption"`
}

type MockDatabase struct {
	ideas map[int]IdeaDB
}

var database IDatabase = &MockDatabase{}

// Инициализирует
func (m *MockDatabase) Init() {
	m.ideas = map[int]IdeaDB{
		0: {
			Id:          0,
			Title:       "Идея 1",
			Desctiption: "Описание 1",
		},
		1: {
			Id:          1,
			Title:       "Идея 2",
			Desctiption: "Описание 2",
		},
		2: {
			Id:          2,
			Title:       "Идея 3",
			Desctiption: "Описание 3",
		},
	}
}

// Возвращает идею по идентификатору
func (m *MockDatabase) GetIdeaById(id int) IdeaDB {
	return m.ideas[id]
}

// Возвращает идеи
func (m *MockDatabase) GetIdeas(filter GetIdeasFilter) []IdeaDB {
	log.Printf("%v \n", filter)
	values := make([]IdeaDB, 0, len(m.ideas))
	for _, v := range m.ideas {
		values = append(values, v)
	}

	return values
}

// Возвращает базу
func GetDatabase() IDatabase {
	return database
}
