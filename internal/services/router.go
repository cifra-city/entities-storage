package service

import (
	"context"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/entities-storage/internal/config"
	"github.com/cifra-city/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	auth := service.TokenManager.AuthMiddleware(service.Config.JWT.AccessToken.SecretKey)
	adminGrant := service.TokenManager.RoleGrant(service.Config.JWT.AccessToken.SecretKey, tokens.AdminRole, tokens.ModeratorRole)

	r.Route("/entities-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(auth)
				r.Route("/create", func(r chi.Router) {
					r.Post("/places", nil)
					r.Post("/distributor", nil)
				})
				r.Route("/update", func(r chi.Router) {
					r.Route("/places/{id}", func(r chi.Router) {
						r.Put("/name", nil)
						r.Put("/description", nil)
						r.Put("/location", nil)
						r.Put("/type", nil)
						r.Post("/photo", nil)
						r.Route("/schedule", func(r chi.Router) {
							r.Delete("/{day_week}", nil)
							r.Put("/{day_week}", nil)
						})
					})
					r.Route("/distributor", func(r chi.Router) {
						r.Put("/name", nil)
					})
				})
				r.Route("/staff", func(r chi.Router) {
					r.Route("/places", func(r chi.Router) {
						r.Post("/add", nil)
						r.Patch("/update", nil)
						r.Delete("/remove", nil)
					})
					r.Route("/distributor", func(r chi.Router) {
						r.Post("/add", nil)
						r.Patch("/update", nil)
						r.Delete("/remove", nil)
					})
				})
			})
			r.Route("/public", func(r chi.Router) {
				r.Route("/place", func(r chi.Router) {
					r.Get("/{id}", nil)
				})
			})
			r.Route("/public", func(r chi.Router) {
				r.Route("/distributor", func(r chi.Router) {
					r.Get("/{id}", nil)
				})
				r.Route("/places", func(r chi.Router) {

					r.Get("/{id}", nil)
					r.Get("/{id}/{photos}", nil)

					r.Get("/{city_id}", nil)

					r.Get("/{street_id}/{house_number}", nil)
					r.Get("/{street_id}/{house_number}/{photos}", nil)

					r.Get("/{street_id}/{type}", nil)
				})
			})

			r.Route("admin", func(r chi.Router) {
				r.Use(adminGrant)
				r.Delete("/places/{id}", nil)
				r.Delete("/distributor/{id}", nil)
				r.Put("/type/{id}", nil)
			})
		})
	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)
	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
