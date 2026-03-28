package main

import (
    "fmt"
    "os"
    "time"
    "github.com/kornev-aa/lab5-cache/internal/pkg/app/cli"
    "github.com/kornev-aa/lab5-cache/pkg/cache"
    "github.com/kornev-aa/lab5-cache/pkg/config"
    "github.com/kornev-aa/lab5-cache/pkg/logger"
    "github.com/kornev-aa/lab5-cache/pkg/storage"
)

func main() {
    cfg, err := config.Load("./config.json")
    if err != nil {
        fmt.Printf("Ошибка загрузки конфига: %s\n", err.Error())
        os.Exit(1)
    }

    log := logger.New()

    var store storage.LocationStorage
    switch cfg.StorageType {
    case "file":
        store = storage.NewFileStorage(cfg.FilePath)
        log.Info("Используется файловое хранилище")
    default:
        log.Error("Неизвестный тип хранилища", nil)
        os.Exit(1)
    }

    // Выбираем тип кэша
    var cacheInstance cache.Cache
    switch cfg.CacheType {
    case "memory":
        cacheInstance = cache.NewMemoryCache()
        log.Info("Используется кэш в памяти")
    case "redis":
        cacheInstance = cache.NewRedisCache(cfg.RedisAddr)
        log.Info("Используется Redis кэш на " + cfg.RedisAddr)
    default:
        log.Error("Неизвестный тип кэша", nil)
        os.Exit(1)
    }

    cacheTTL := 5 * time.Minute

    app := cli.New(log, store, cacheInstance, cacheTTL)

    if len(os.Args) > 2 && os.Args[1] == "save" {
        var lat, lon float64
        fmt.Sscanf(os.Args[2], "%f", &lat)
        fmt.Sscanf(os.Args[3], "%f", &lon)
        if err := app.SaveLocation(lat, lon); err != nil {
            log.Error("Ошибка сохранения", err)
            os.Exit(1)
        }
        log.Info("Координаты сохранены")
        return
    }

    if err := app.Run(); err != nil {
        log.Error("Приложение завершилось с ошибкой", err)
        os.Exit(1)
    }

    log.Info("Приложение завершилось успешно")
}
