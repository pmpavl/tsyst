package passage

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/pmpavl/tsyst/app/api/passage/request"
	"github.com/pmpavl/tsyst/app/api/passage/response"
	"github.com/pmpavl/tsyst/models"
	"github.com/pmpavl/tsyst/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Create(c *gin.Context) {
	req, err := request.GetCreateRequest(c)
	if err != nil {
		s.errorResponse(c, http.StatusBadRequest, errors.Wrap(err, "get create request"))

		return //! 400
	}

	user, err := s.dbUsers.ReadByAccessToken(c.Request.Context(), req.AccessToken)
	if errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusForbidden, ErrAccessTokenNotExist)

		return //! 403
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db users read by access token"))

		return //! 500
	}

	test, err := s.dbTests.ReadByPath(c.Request.Context(), req.Path)
	if err == mongo.ErrNoDocuments {
		s.errorResponse(c, http.StatusBadRequest, ErrTestPathNotExist)

		return //! 400
	} else if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db tests read by path"))

		return //! 500
	}

	lastPassage, err := s.dbPassages.SearchLastUserPassage(c.Request.Context(), user.ID, test.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db passages search last user passage"))

		return //! 500
	}

	if lastPassage != nil && test.Repeat != nil {
		if err := s.opportunityCreatePassage(c.Request.Context(), lastPassage.End, test.Repeat); err != nil {
			s.errorResponse(c, http.StatusForbidden, err)

			return //! 403
		}
	}

	tasks, err := s.searchPassageTasks(c.Request.Context(), test.Tasks)
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, err)

		return //! 500
	}

	id, err := s.dbPassages.Create(c.Request.Context(), models.NewPassage(user.ID, test, tasks))
	if err != nil {
		s.errorResponse(c, http.StatusInternalServerError, errors.Wrap(err, "db passages create"))

		return //! 500
	}

	s.okResponse(c, response.CreateResponse{ID: id}) //! 200
}

// Возможно ли создание прохождения.
func (s *Service) opportunityCreatePassage(
	ctx context.Context,
	lastPassageEnd time.Time,
	testRepeat *models.TestRepeat,
) error {
	// Есть незаконченный тест
	if lastPassageEnd.After(time.Now()) {
		return ErrAlreadyStart
	}

	// Тест нельзя перепроходить
	if testRepeat.Type == constants.Disposable {
		return ErrDisposable
	}

	// Не прошло нужное время до повторного прохождения
	if testRepeat.Type == constants.Repeatable &&
		lastPassageEnd.Add(testRepeat.TimeToRepeat.Time()).After(time.Now()) {
		return ErrTimeToRepeat
	}

	return nil
}

// Составление списка задач для прохождения.
func (s *Service) searchPassageTasks(
	ctx context.Context,
	testTasks []*models.TestTask,
) ([]*models.PassageTask, error) {
	var (
		tasks   = make([]*models.PassageTask, 0, len(testTasks))
		usedIDs = make([]primitive.ObjectID, 0, len(testTasks))
	)

	for _, testTask := range testTasks {
		var (
			task *models.Task
			err  error
		)

		if testTask.ID != nil { // Конкретная задача
			task, err = s.dbTasks.Read(ctx, *testTask.ID)
		} else { // Случайная задача
			task, err = s.dbTasks.Search(ctx, testTask.Complexity, testTask.Themes, usedIDs)
			if task != nil {
				usedIDs = append(usedIDs, task.ID)
			}
		}

		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrIncorrectlyCompiledTest
		} else if err != nil {
			return nil, errors.Wrap(err, "db tasks search")
		}

		tasks = append(tasks, models.NewPassageTask(task, testTask.Points))
	}

	return tasks, nil
}
