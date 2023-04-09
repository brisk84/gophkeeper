package usecase

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/brisk84/gophkeeper/domain"
	"github.com/brisk84/gophkeeper/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestSaveLogin(t *testing.T) {
	stor := new(storageMock)

	ctx := context.Background()
	mockId := 1
	mockUserId := 4
	mockTitle := "some service"
	mockLogin := "login"
	mockPass := "pass"

	m := make(map[string]string)
	m["login"] = mockLogin
	m["pass"] = mockPass
	data, err := json.Marshal(m)
	assert.NoError(t, err)
	stor.On("SaveData", ctx, mockUserId, mockTitle, data, "login").Return(mockId, nil).Once()

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal(err)
	}
	u, err := New(lg, stor)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		id, err := u.SaveLogin(ctx, mockUserId, "some service", "login", "pass")
		assert.NoError(t, err)
		assert.Equal(t, mockId, id)

		stor.AssertExpectations(t)
	})
}

func TestAuth(t *testing.T) {
	stor := new(storageMock)

	ctx := context.Background()
	mockUser := domain.User{Login: "user01", Password: "123456"}
	token := "b80f2699a2a444546aa7ef44b3aca41b42a097e43fbf16c60e731913129529c4"
	stor.On("GetByToken", ctx, token).Return(mockUser, nil)

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal(err)
	}
	u, err := New(lg, stor)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		user, err := u.Auth(ctx, token)
		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)

		stor.AssertExpectations(t)
	})
}

// func TestLogin(t *testing.T) {
// 	stor := new(storageMock)

// 	ctx := context.Background()
// 	user := domain.User{Login: "user01", Password: "123456"}
// 	mockToken := "b80f2699a2a444546aa7ef44b3aca41b42a097e43fbf16c60e731913129529c4"
// 	// stor.On("GetByLogin", ctx, "user01").Return("b80f2699a2a444546aa7ef44b3aca41b42a097e43fbf16c60e731913129529c4")
// 	stor.On("SaveUser", ctx, user).Return(mockToken, nil).Once()

// 	lg, err := logger.New(true)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	u, err := New(lg, stor)
// 	assert.NoError(t, err)

// 	t.Run("success", func(t *testing.T) {
// 		token, err := u.Register(ctx, user)
// 		assert.NoError(t, err)
// 		assert.Equal(t, mockToken, token)

// 		stor.AssertExpectations(t)

// 	})
// }

// func TestRegister(t *testing.T) {
// 	stor := new(storageMock)

// 	ctx := context.Background()
// 	user := domain.User{Login: "user01", Password: "123456"}
// 	mockToken := "b80f2699a2a444546aa7ef44b3aca41b42a097e43fbf16c60e731913129529c4"
// 	// stor.On("GetByLogin", ctx, "user01").Return("b80f2699a2a444546aa7ef44b3aca41b42a097e43fbf16c60e731913129529c4")
// 	stor.On("SaveUser", ctx, user).Return(mockToken, nil).Once()

// 	lg, err := logger.New(true)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	u, err := New(lg, stor)
// 	assert.NoError(t, err)

// 	t.Run("success", func(t *testing.T) {
// 		token, err := u.Register(ctx, user)
// 		assert.NoError(t, err)
// 		assert.Equal(t, mockToken, token)

// 		stor.AssertExpectations(t)

// 	})

// 	// mockTaskRepository := new(mocks.TaskRepository)
// 	// userObjectID := primitive.NewObjectID()
// 	// userID := userObjectID.Hex()

// 	// t.Run("success", func(t *testing.T) {

// 	// 	mockTask := domain.Task{
// 	// 		ID:     primitive.NewObjectID(),
// 	// 		Title:  "Test Title",
// 	// 		UserID: userObjectID,
// 	// 	}

// 	// 	mockListTask := make([]domain.Task, 0)
// 	// 	mockListTask = append(mockListTask, mockTask)

// 	// 	mockTaskRepository.On("FetchByUserID", mock.Anything, userID).Return(mockListTask, nil).Once()

// 	// 	u := usecase.NewTaskUsecase(mockTaskRepository, time.Second*2)

// 	// 	list, err := u.FetchByUserID(context.Background(), userID)

// 	// 	assert.NoError(t, err)
// 	// 	assert.NotNil(t, list)
// 	// 	assert.Len(t, list, len(mockListTask))

// 	// 	mockTaskRepository.AssertExpectations(t)
// 	// })

// 	// t.Run("error", func(t *testing.T) {
// 	// 	mockTaskRepository.On("FetchByUserID", mock.Anything, userID).Return(nil, errors.New("Unexpected")).Once()

// 	// 	u := usecase.NewTaskUsecase(mockTaskRepository, time.Second*2)

// 	// 	list, err := u.FetchByUserID(context.Background(), userID)

// 	// 	assert.Error(t, err)
// 	// 	assert.Nil(t, list)

// 	// 	mockTaskRepository.AssertExpectations(t)
// 	// })

// }
