package app

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	v1 "social-media-app/internal/handler/http/v1"
	"social-media-app/internal/postgres"
	"social-media-app/internal/usecase"
	"syscall"
	"time"
)

func Run() {
	log := logrus.New()

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := sqlx.Open("pgx", cfg.Postgres.Conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConn)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConn)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Second)
	db.SetConnMaxIdleTime(cfg.Postgres.IdleConnMaxLifetime * time.Second)

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := postgres.NewUserRepository(db)

	createUserUC := usecase.NewCreateUserUseCase(userRepo)

	userHandlerV1 := v1.NewUserHandler(createUserUC)

	f := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          v1.HandleError,
	})

	f.Use(func(ctx *fiber.Ctx) error {
		err = ctx.Next()
		if err != nil {
			err = ctx.App().ErrorHandler(ctx, err)
			if err != nil {
				return err
			}
		}

		log.WithField(
			"status", ctx.Response().StatusCode(),
		).WithField(
			"method", ctx.Method(),
		).WithField(
			"path", ctx.Path(),
		).Info("Request")
		return nil
	})

	routerV1 := f.Group("v1")

	usersV1 := routerV1.Group("users")

	userHandlerV1.Register(usersV1)

	go func() {
		err = f.Listen(net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port))
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	log.Info("application has started")

	<-exit

	err = f.Shutdown()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("application has been shut down")
}
