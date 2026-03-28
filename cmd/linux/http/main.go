package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    httphandlers "github.com/kornev-aa/lab5-refactor/internal/pkg/http"  // ← переименовали
    "github.com/kornev-aa/lab5-refactor/pkg/config"
    "github.com/kornev-aa/lab5-refactor/pkg/logger"
    "github.com/kornev-aa/lab5-refactor/pkg/storage"
)

func main() {
    cfg, err := config.Load("./config.json")
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    logg := logger.New()

    var store storage.LocationStorage
    switch cfg.StorageType {
    case "file":
        store = storage.NewFileStorage(cfg.FilePath)
        logg.Info("Using file storage")
    default:
        logg.Error("Unknown storage type", nil)
        return
    }

    handlers := httphandlers.NewHandlers(logg, store)  // ← используем новое имя

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/weather", handlers.GetWeather)
    r.Post("/location", handlers.SaveLocation)
    r.Get("/location", handlers.GetLocation)

    logg.Info(fmt.Sprintf("HTTP server starting on :8080"))
    http.ListenAndServe(":8080", r)
}
