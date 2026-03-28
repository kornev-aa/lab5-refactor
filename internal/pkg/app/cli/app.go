package cli

import (
    "fmt"
    "time"
    "github.com/kornev-aa/lab5-cache/internal/pkg/weather"
    "github.com/kornev-aa/lab5-cache/pkg/cache"
    "github.com/kornev-aa/lab5-cache/pkg/storage"
)

type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string, err error)
}

type cliApp struct {
    log     Logger
    storage storage.LocationStorage
    weather *weather.WeatherService
}

func New(log Logger, storage storage.LocationStorage, cache cache.Cache, cacheTTL time.Duration) *cliApp {
    return &cliApp{
        log:     log,
        storage: storage,
        weather: weather.NewWeatherService(cache, cacheTTL),
    }
}

func (c *cliApp) Run() error {
    c.log.Info("Запуск приложения")

    lat, err := c.storage.GetLatitude()
    if err != nil {
        c.log.Debug("Не удалось загрузить широту, используем значение по умолчанию")
        lat = 53.6688
    }

    lon, err := c.storage.GetLongitude()
    if err != nil {
        c.log.Debug("Не удалось загрузить долготу, используем значение по умолчанию")
        lon = 23.8223
    }

    c.log.Info(fmt.Sprintf("Координаты: широта=%.4f, долгота=%.4f", lat, lon))

    result, err := c.weather.GetWeather(lat, lon)
    if err != nil {
        c.log.Error("Ошибка получения погоды", err)
        return err
    }

    fmt.Printf("Температура воздуха - %.2f градусов цельсия\n", result.Current.Temperature)
    return nil
}

func (c *cliApp) SaveLocation(lat, lon float64) error {
    c.log.Info(fmt.Sprintf("Сохранение координат: %.4f, %.4f", lat, lon))
    return c.storage.SaveLocation(lat, lon)
}
