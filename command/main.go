package main

import (
	"EduCRM/config"
	"EduCRM/cronjobs"
	"EduCRM/package/handler"
	"EduCRM/package/repository"
	"EduCRM/package/repository/psql"
	"EduCRM/package/repository/psql/migration"
	"EduCRM/package/service"
	"EduCRM/package/store"
	"EduCRM/server"
	"EduCRM/tools/logger"
	"context"
	"os"
	"os/signal"
	"syscall"
)

// @title EduCRM
// @version 1.0
// @description API Server for EduCRM Application
// @termsOfService gitlab.com/edu-crm
// @host gitlab.com/edu-crm
// @BasePath
// @contact.name   Bakhodir Yashin Mansur
// @contact.email  phapp0224mb@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	loggers := logger.GetLogger()
	cfg := config.Config()
	// Migration Up
	err := migration.MigratePsql(cfg.Postgres, loggers, true)
	if err != nil {
		loggers.Error("error while migrate up", err)
	}
	db, err := psql.NewPostgresDB(&cfg.Postgres, loggers)
	if err != nil {
		loggers.Fatalf("failed to initialize db: %s", err.Error())
		panic(err)
	}
	minio, err := store.MinioConnection(&cfg.Minio, loggers)
	if err != nil {
		loggers.Fatal("error while connect to minio server", err)
		panic(err)
	}
	repos := repository.NewRepository(db, loggers)
	newStore := store.NewStore(minio, &cfg.Minio, loggers)
	newService := service.NewService(repos, newStore, loggers)
	handlers := handler.NewHandler(newService, loggers)
	go cronjobs.NewCronJobs(db, minio, loggers)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Server.ServerPort, handlers.InitRoutes()); err != nil {
			loggers.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		loggers.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		loggers.Errorf("error occurred on db connection close: %s", err.Error())
	}
	// Start server (with or without graceful shutdown).
	// port := fmt.Sprintf(":%d", cfg.ServerPort)
	// err = app.Run(port)
	// if err != nil {
	// 	loggers.Fatal("error while running server", err)
	// }
}
