package main

import (
    "os"
    "time"
    "github.com/kornev-aa/lab5-refactor/internal/adapters/weather"
    "github.com/kornev-aa/lab5-refactor/internal/pkg/app/cli"
    "github.com/kornev-aa/lab5-refactor/pkg/cache"
    "github.com/kornev-aa/lab5-refactor/pkg/logger"
)

func main() {
    log := logger.New()

    // Создаём кэш в памяти
    memCache := cache.NewMemoryCache()
    cacheTTL := 5 * time.Minute

    // Создаём WeatherInfo с зависимостями
    wi := weather.New(log, memCache, cacheTTL)

    // Создаём приложение с зависимостями
    app := cli.New(log, wi)

    if err := app.Run(); err != nil {
        log.Error("Приложение завершилось с ошибкой", err)
        os.Exit(1)
    }

    log.Info("Приложение завершилось успешно")
}
