package main

import (
	"flag"
	"fmt"
	"log"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/JimMir45/KAR/server/internal/config"
	"github.com/JimMir45/KAR/server/internal/domain"
	"github.com/JimMir45/KAR/server/internal/handler"
)

func main() {
	configPath := flag.String("config", "configs/config.yaml", "path to config file")
	flag.Parse()

	// Load config
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Init logger
	var logger *zap.Logger
	if cfg.Server.Mode == "release" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer func() { _ = logger.Sync() }()

	// Connect database
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("failed to get sql.DB", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)

	// Auto migrate (development only)
	if cfg.Server.Mode == "debug" {
		if err := autoMigrate(db); err != nil {
			logger.Fatal("failed to auto migrate", zap.Error(err))
		}
		logger.Info("database auto migration completed")
	}

	// Setup router
	r := handler.SetupRouter(cfg, logger)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("server starting", zap.String("addr", addr))
	if err := r.Run(addr); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.MemberProfile{},
		&domain.MemberLevel{},
		&domain.BalanceLog{},
		&domain.Track{},
		&domain.Sponsor{},
		&domain.Activity{},
		&domain.Registration{},
		&domain.LapRecord{},
		&domain.Order{},
		&domain.OrderItem{},
	)
}
