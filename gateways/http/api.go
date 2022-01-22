package http

import (
	"fmt"
	"usecases"
)

type App struct {
	permissionsUseCase *usecases.PermissionsUseCase
}

func (a *App) Start() {
	fmt.Println("Starting http app")
}

func (a *App) Stop() {
	fmt.Println("Stopping http app")
}

func NewApp(permissionsUseCase *usecases.PermissionsUseCase) *App {
	return &App{permissionsUseCase: permissionsUseCase}
}
