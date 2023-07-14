package router

import (
	"fmt"

	_ "github.com/trunov/erply-assignement-task/user-service/docs"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/trunov/erply-assignement-task/user-service/internal/config"
	"github.com/trunov/erply-assignement-task/user-service/internal/middleware"
	"github.com/trunov/erply-assignement-task/user-service/internal/transport/handler"
	use_case "github.com/trunov/erply-assignement-task/user-service/internal/use-case"
)

// @title Erply service API
// @version 1.0
// @description Service provides two endpoints to manage Erply customers. It connects to a local PostgreSQL database and an external Erply service for user authentication and reading/writing customer data.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func NewRouter(h *handler.Handler, erply use_case.Erply, cfg config.Config) chi.Router {

	r := chi.NewRouter()

	r.Route("/customer", func(r chi.Router) {
		r.Use(middleware.VerifyErplyUser(erply, cfg.ClientCode, cfg.Username, cfg.Password))
		r.Use(middleware.TokenAuthorization(cfg.Auth))

		r.Get("/{id}", h.GetCustomer)
		r.Post("/", h.AddCustomer)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", cfg.Port)),
	))

	return r
}
