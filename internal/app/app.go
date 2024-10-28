package app

import (
	"context"
	"go-api-example/internal/port"
	"go-api-example/internal/storage"
	"go-api-example/internal/types"
	"net/http"
	"time"
)

type App struct {
	storage storage.Storage
}

func NewApp(
	storage storage.Storage) *App {
	return &App{
		storage: storage,
	}
}

func (a *App) HealthzCheck(_ context.Context, _ port.HealthzCheckRequestObject) (port.HealthzCheckResponseObject, error) {
	return port.HealthzCheck200JSONResponse{Status: "OK", Timestamp: time.Now()}, nil
}

func (a *App) GetUserList(_ context.Context, req port.GetUserListRequestObject) (port.GetUserListResponseObject, error) {
	var (
		users              []*types.User
		limit, skip, total int
		err                error
	)
	limit = req.Params.Limit
	if limit == 0 {
		limit = 20
	}
	skip = req.Params.Skip

	if users, total, err = a.storage.GetAll(skip, limit); err != nil {
		return port.GetUserList502JSONResponse{Message: err.Error(), Code: http.StatusBadGateway}, nil
	}

	res := port.UserListResponse{
		Items:      make([]port.User, len(users)),
		TotalItems: total,
		Limit:      limit,
		Skip:       skip,
	}
	for i := range users {
		res.Items[i] = port.User{
			ID:   users[i].ID,
			Name: users[i].Name,
		}
	}

	return port.GetUserList200JSONResponse(res), nil
}

func (a *App) GetUser(_ context.Context, req port.GetUserRequestObject) (port.GetUserResponseObject, error) {
	var (
		user *types.User
		err  error
	)

	if user, err = a.storage.Get(req.ID); err != nil {
		return port.GetUser502JSONResponse{Message: err.Error(), Code: http.StatusBadGateway}, nil
	}

	res := port.User{
		ID:   user.ID,
		Name: user.Name,
	}

	return port.GetUser200JSONResponse(res), nil
}

func (a *App) CreateUser(_ context.Context, req port.CreateUserRequestObject) (port.CreateUserResponseObject, error) {
	var (
		user *types.User
		err  error
	)

	if user, err = a.storage.Create(req.Body.Name); err != nil {
		return port.CreateUser502JSONResponse{Message: err.Error(), Code: http.StatusBadGateway}, nil
	}

	res := port.User{
		ID:   user.ID,
		Name: user.Name,
	}
	return port.CreateUser201JSONResponse(res), nil
}
