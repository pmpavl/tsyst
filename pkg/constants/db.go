package constants

import "github.com/pkg/errors"

var ErrResultNotObjectID = errors.New("result not object id")

type DatabaseName string

const (
	DatabaseCore     DatabaseName = "core"     // Core база с данными о пользователях
	DatabaseSource   DatabaseName = "source"   // База с исходными материалами
	DatabasePassages DatabaseName = "passages" // База с прохождениями тестов
)

func (n DatabaseName) String() string { return string(n) }

type CollectionName string

const (
	CollectionUsers    CollectionName = "users"    // Коллекция пользователей
	CollectionTests    CollectionName = "tests"    // Коллекция тестов
	CollectionTasks    CollectionName = "tasks"    // Коллекция задач
	CollectionPassages CollectionName = "passages" // Коллекция прохождений тестов
)

func (n CollectionName) String() string { return string(n) }
