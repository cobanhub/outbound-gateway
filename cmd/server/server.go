package server

import (
	"github.com/cobanhub/madakaripura/internal/handler/controller"
)

func Start() {
	apiOpts := controller.APIOptions{
		Prefix:         "/madakaripura",
		DefaultTimeout: 30,
		EnableSwagger:  true,
		// Interactor:     interactor.NewInteractor(),
	}
	controller := controller.NewAPI(apiOpts)
	controller.Register()
}
