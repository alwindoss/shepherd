package server

import (
	"database/sql"
	"io/fs"
	"log/slog"
	"net/http"
	"shepherd/internal/handler"
	"shepherd/internal/repository"
	"shepherd/internal/service"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func Start(cfg *Config) error {
	secretKey, err := handler.GetJWTSecret()
	if err != nil {
		return err
	}
	hdlr := &handler.SessionHandler{
		JWTSecret: secretKey,
	}
	pgHdlr := &handler.PageHandler{
		JWTSecret: secretKey,
	}
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	publicFS, _ := fs.Sub(cfg.FS, "public")
	r.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.FS(publicFS))))

	// we will implement these handlers in the next sections
	r.Get("/", pgHdlr.HomePageHandler)
	r.Get("/about", pgHdlr.AboutPageHandler)
	r.Get("/dashboard", pgHdlr.DashboardPageHandler)
	r.Get("/pages/login", pgHdlr.LoginPageHandler)
	r.Post("/pages/login", pgHdlr.LoginFormHandler)

	// r.Post("/sessions/login", hdlr.LoginHandler)
	r.Post("/sessions/refresh", hdlr.RefreshHandler)
	r.Post("/sessions/logout", hdlr.LogoutHandler)

	db, err := sql.Open("postgres", "postgres://shepherd_user:shepherd_password@localhost:5432/shepherd_db?sslmode=disable")
	if err != nil {
		return err
	}

	memRepo := repository.NewMembersRepository(db)
	memSvc := service.NewMemberService(memRepo)
	memHdlr := handler.NewMemberHandler(memSvc)
	configureMemberRoutes(r, memHdlr)

	evntHdlr := handler.NewEventHandler(nil)
	configureEventRoutes(r, evntHdlr)

	donHdlr := handler.NewDonationHandler(nil)
	configureDonationsRoutes(r, donHdlr)

	attr := slog.String("Host:Port", cfg.Addr)
	slog.Info("Shepherd server started on", attr)
	// start the server on port 8000
	return http.ListenAndServe(cfg.Addr, r)
}
