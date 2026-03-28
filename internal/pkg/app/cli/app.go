package cli

import (
    "fmt"
    "github.com/kornev-aa/lab5-refactor/internal/domain/models"
)

type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string, err error)
}

type WeatherInfo interface {
    GetTemperature(lat, lon float64) models.TempInfo
}

type cliApp struct {
    l  Logger
    wi WeatherInfo
}

func New(l Logger, wi WeatherInfo) *cliApp {
    return &cliApp{
        l:  l,
        wi: wi,
    }
}

func (c *cliApp) Run() error {
    c.l.Info("Запуск приложения")

    // Координаты Гродно (можно потом загружать из storage)
    lat := 53.6688
    lon := 23.8223

    c.l.Info(fmt.Sprintf("Координаты: широта=%.4f, долгота=%.4f", lat, lon))

    temp := c.wi.GetTemperature(lat, lon)

    fmt.Printf("Температура воздуха - %.2f градусов цельсия\n", temp.Temp)
    return nil
}
