package main

import (
	"assets/config"
	"assets/internal/assets"
	"assets/internal/auth"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.New(ctx, cfg.PgDBDSN)
	if err != nil {
		log.Fatal(err)
	}

	authRepository := auth.NewRepository(pool)
	authService := auth.NewService(authRepository, cfg.SessionDurationHours)
	authAPI := auth.NewAPI(authService)

	assetsRepository := assets.NewRepository(pool)
	assetsService := assets.NewService(assetsRepository)
	assetsAPI := assets.NewAPI(assetsService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth", authAPI.Auth)
	mux.HandleFunc("/api/upload-asset/{asset_name}",
		auth.Middleware(assetsAPI.Upload, authService),
	)
	mux.HandleFunc("/api/asset/{asset_name}",
		auth.Middleware(assetsAPI.GetAsset, authService),
	)
	mux.HandleFunc("/api/delete-asset/{asset_name}",
		auth.Middleware(assetsAPI.DeleteAsset, authService),
	)
	mux.HandleFunc("/api/assets",
		assetsAPI.GetList,
	)

	log.Printf("HTTP server is started: %s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.AppPort), mux))
	//log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", cfg.AppPort), "./certs/server.crt", "./certs/server.key", mux))
}
