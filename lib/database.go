package database

import (
	"log"
	"sync"
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

	// Добавляет новую идею
	AddIdea(IdeaDB)

	// Редактирует идею
	EditIdea(IdeaDB)

	// Удаляет идею
	DeleteIdea(id int)

	// Возвращает идеи
	GetIdeas(filter GetIdeasFilter) []IdeaDB
}

// Идея базы данных
type IdeaDB struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Desctiption string    `json:"desctiption"`
	DateTime    time.Time `json:"datetime"`
	UserId      int       `json:"userid"`
}

type MockDatabase struct {
	sync.RWMutex
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
	m.RLock()
	defer m.RUnlock()

	return m.ideas[id]
}

// Возвращает идеи
func (m *MockDatabase) GetIdeas(filter GetIdeasFilter) []IdeaDB {
	log.Printf("%v \n", filter)
	values := make([]IdeaDB, 0, len(m.ideas))
	for i, v := range m.ideas {
		if i >= filter.Count {
			break
		}

		if filter.IsFavorite && filter.UserId != nil {
			if v.UserId == *filter.UserId {
				values = append(values, v)
			}
		} else if filter.UserId != nil {
			values = append(values, v)
		} else if v.DateTime.Before(filter.StartDate) {
			values = append(values, v)
		}
	}

	return values
}

// Добавляет новую идею
func (m *MockDatabase) AddIdea(idea IdeaDB) {
	m.Lock()
	defer m.Unlock()

	m.ideas[idea.Id] = idea
}

// Редактирует идею
func (m *MockDatabase) EditIdea(idea IdeaDB) {
	m.Lock()
	defer m.Unlock()

	m.ideas[idea.Id] = idea
}

// Удаляет идею
func (m *MockDatabase) DeleteIdea(id int) {
	m.Lock()
	defer m.Unlock()

	delete(m.ideas, id)
}

// Возвращает базу
func GetDatabase() IDatabase {
	return database
}
