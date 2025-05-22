package controller

import (
	"github.com/cobanhub/lib/router"
	"github.com/cobanhub/madakaripura/internal/handler/middleware"
	"github.com/cobanhub/madakaripura/internal/services/interactor"
)

type (
	API struct {
		prefix         string
		defaultTimeout int
		enableSwagger  bool
		mw             middleware.MiddlewareInterface
		interactor     interactor.Interactor
		rateLimit      map[string]router.PathLimit
	}

	APIOptions struct {
		Prefix         string
		EnableSwagger  bool
		DefaultTimeout int
		Middleware     middleware.MiddlewareInterface
		Interactor     interactor.Interactor
		RateLimit      map[string]router.PathLimit
		// Add more options as needed
	}
)

func NewAPI(opts APIOptions) *API {
	return &API{
		prefix:         opts.Prefix,
		defaultTimeout: opts.DefaultTimeout,
		enableSwagger:  opts.EnableSwagger,
		mw:             opts.Middleware,
		interactor:     opts.Interactor,
	}
}

func (api *API) Register() *router.Router {
	routes := router.New(router.Opts{
		AppName:      "madakaripura",
		Prefix:       api.prefix,
		WriteTimeout: 30,
		ReadTimeout:  30,
		RateLimit:    api.rateLimit,
	})

	routes.Group(api.prefix, func(v1 *router.Router) {
		v1.Group("/v1", func(route *router.Router) {
			// r.GET("/health", api.mw.HealthCheckHandler)
			// r.GET("/metrics", api.mw.MetricsHandler)
			// r.GET("/version", api.mw.VersionHandler)
			// r.GET("/swagger.yaml", api.mw.SwaggerHandler)
			// r.GET("/swagger-ui/", api.mw.SwaggerUIHandler)
			route.GET("/outbound/{integration}", api.HandleOutbound)
		})
	})

	return routes
}
