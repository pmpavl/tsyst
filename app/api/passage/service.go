package passage

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/constants"
	"github.com/pmpavl/tsyst/pkg/log"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrAccessTokenNotExist     = errors.New("access token not exist")    // Access токен не существует
	ErrTestPathNotExist        = errors.New("test path not exist")       // Пути теста не существует
	ErrAlreadyStart            = errors.New("already start")             // Есть незаконченный тест
	ErrDisposable              = errors.New("disposable")                // Тест нельзя перепроходить
	ErrTimeToRepeat            = errors.New("time to repeat")            // Не прошло время до повторного прохождения
	ErrIncorrectlyCompiledTest = errors.New("incorrectly compiled test") // Неправильно составлен тест
	ErrPassageIDNotExist       = errors.New("passage id not exist")      // Прохождение не существует
	ErrIncorrectPassageUser    = errors.New("incorrect passage user")    // Неверный пользователь
	ErrAlreadyEnd              = errors.New("already end")               // Прохождение завершено
	ErrNotActualTaskNum        = errors.New("not actual task num")       // Не актуальный номер задачи
)

type DBUsers interface {
	ReadByAccessToken(ctx context.Context, accessToken string) (*models.User, error)
}

type DBTests interface {
	ReadByPath(ctx context.Context, path string) (*models.Test, error)
}

type DBTasks interface {
	Read(ctx context.Context, id primitive.ObjectID) (*models.Task, error)
	Search(
		ctx context.Context, complexity constants.ComplexityType, themes []string, usedIDs []primitive.ObjectID,
	) (*models.Task, error)
}

type DBPassages interface {
	Create(ctx context.Context, passage *models.Passage) (primitive.ObjectID, error)
	Read(ctx context.Context, id primitive.ObjectID) (*models.Passage, error)
	SearchLastUserPassage(ctx context.Context, userID, testID primitive.ObjectID) (*models.Passage, error)
	Update(ctx context.Context, passage *models.Passage) error
}

type Service struct {
	log *zerolog.Logger

	dbUsers    DBUsers
	dbTests    DBTests
	dbTasks    DBTasks
	dbPassages DBPassages
}

func New(dbUsers DBUsers, dbTests DBTests, dbTasks DBTasks, dbPassages DBPassages) *Service {
	return &Service{
		log:        log.For("service-passage"),
		dbUsers:    dbUsers,
		dbTests:    dbTests,
		dbTasks:    dbTasks,
		dbPassages: dbPassages,
	}
}

func (s *Service) okResponse(c *gin.Context, response any) { s.response(c, http.StatusOK, response) }

type ErrorDefault struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (s *Service) errorResponse(c *gin.Context, code int, err error) {
	s.response(c, code, ErrorDefault{
		Code:    http.StatusText(code),
		Message: err.Error(),
	})
}

func (s *Service) response(c *gin.Context, code int, response any) { c.JSON(code, response) }
