package main

import (
    "fmt"
    "os"
    "github.com/kornev-aa/lab5-refactor/internal/pkg/gui"
    "github.com/kornev-aa/lab5-refactor/pkg/config"
    "github.com/kornev-aa/lab5-refactor/pkg/logger"
    "github.com/kornev-aa/lab5-refactor/pkg/storage"
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

    app := gui.NewGUIApp(log, store)
    if err := app.Run(); err != nil {
        log.Error("Ошибка GUI", err)
        os.Exit(1)
    }
}
