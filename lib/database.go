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

	// Добавляет комментарий
	AddComment(comment CommentDB) CommentDB

	// Возвращает комментарии для идеи
	GetComments(ideaId int) []CommentDB
}

// Идея базы данных
type IdeaDB struct {
	// Идентификатор идеи
	Id int `json:"id"`
	// Заголовок
	Title string `json:"title"`
	// Описание идеи - основной текст
	Desctiption string `json:"desctiption"`
	// Дата создания идеи
	DateTime time.Time `json:"datetime"`
	// Идентификатор пользователя, который создал идею
	UserId int `json:"userid"`
}

// Комментарий к идеи
type CommentDB struct {
	// Идентификатор комментария
	Id int `json:"id"`
	// Текст комментария
	Text string `json:"text"`
	// Идентификатор пользователя, который сделал комментарий
	UserId int `json:"userid"`
}

// Контейнер для данных идеи
type ideaDataContainer struct {
	// Идея
	idea IdeaDB
	// Комментарии к идеи
	comments map[int]CommentDB
}

type MockDatabase struct {
	ideasLock   sync.RWMutex
	ideas       map[int]ideaDataContainer
	commentLock sync.RWMutex
	// Комментарии
	comments map[int]CommentDB
}

var database IDatabase = &MockDatabase{}

// Возвращает базу
func GetDatabase() IDatabase {
	return database
}

// Инициализирует
func (m *MockDatabase) Init() {
	m.ideas = map[int]ideaDataContainer{
		0: {
			idea: IdeaDB{
				Id:          0,
				Title:       "Идея 1",
				Desctiption: "Описание 1",
			},
		},

		1: {
			idea: IdeaDB{
				Id:          1,
				Title:       "Идея 2",
				Desctiption: "Описание 2",
			},
		},

		2: {
			idea: IdeaDB{
				Id:          2,
				Title:       "Идея 3",
				Desctiption: "Описание 3",
			},
		},
	}
}

// Возвращает идею по идентификатору
func (m *MockDatabase) GetIdeaById(id int) IdeaDB {
	m.ideasLock.RLock()
	defer m.ideasLock.RUnlock()

	return m.ideas[id].idea
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
			if v.idea.UserId == *filter.UserId {
				values = append(values, v.idea)
			}
		} else if filter.UserId != nil {
			values = append(values, v.idea)
		} else if v.idea.DateTime.Before(filter.StartDate) {
			values = append(values, v.idea)
		}
	}

	return values
}

// Добавляет новую идею
func (m *MockDatabase) AddIdea(idea IdeaDB) {
	m.ideasLock.Lock()
	defer m.ideasLock.Unlock()

	m.ideas[idea.Id] = ideaDataContainer{
		idea: idea,
	}
}

// Редактирует идею
func (m *MockDatabase) EditIdea(idea IdeaDB) {
	m.ideasLock.Lock()
	defer m.ideasLock.Unlock()

	m.ideas[idea.Id] = ideaDataContainer{
		idea: idea,
	}
}

// Удаляет идею
func (m *MockDatabase) DeleteIdea(id int) {
	m.ideasLock.Lock()
	defer m.ideasLock.Unlock()

	delete(m.ideas, id)
}

// Добавляет комментарий
func (m *MockDatabase) AddComment(comment CommentDB) CommentDB {
	return comment
}

// Возвращает комментарии
func (m *MockDatabase) GetComments(ideaId int) []CommentDB {
	return nil
}
