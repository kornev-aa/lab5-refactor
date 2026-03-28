package gui

import (
    "fmt"
    "strconv"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/kornev-aa/lab5-cache/internal/pkg/weather"
    "github.com/kornev-aa/lab5-cache/pkg/logger"
    "github.com/kornev-aa/lab5-cache/pkg/storage"
)

type GUIApp struct {
    log     *logger.Logger
    storage storage.LocationStorage
    weather *weather.WeatherService
}

func NewGUIApp(log *logger.Logger, storage storage.LocationStorage) *GUIApp {
    return &GUIApp{
        log:     log,
        storage: storage,
        weather: weather.NewWeatherService(),
    }
}

func (g *GUIApp) Run() error {
    g.log.Info("Запуск GUI приложения")

    myApp := app.New()
    myWindow := myApp.NewWindow("Погода")
    myWindow.Resize(fyne.NewSize(400, 300))

    lat, _ := g.storage.GetLatitude()
    lon, _ := g.storage.GetLongitude()

    latEntry := widget.NewEntry()
    latEntry.SetText(fmt.Sprintf("%.4f", lat))
    lonEntry := widget.NewEntry()
    lonEntry.SetText(fmt.Sprintf("%.4f", lon))

    resultLabel := widget.NewLabel("Нажмите 'Получить погоду'")
    resultLabel.Wrapping = fyne.TextWrapWord

    getWeatherBtn := widget.NewButton("Получить погоду", func() {
        g.log.Debug("Кнопка 'Получить погоду' нажата")

        latVal, err1 := strconv.ParseFloat(latEntry.Text, 64)
        lonVal, err2 := strconv.ParseFloat(lonEntry.Text, 64)

        if err1 != nil || err2 != nil {
            resultLabel.SetText("Ошибка: введите корректные координаты")
            g.log.Error("Неверные координаты", fmt.Errorf("lat=%s, lon=%s", latEntry.Text, lonEntry.Text))
            return
        }

        g.log.Debug(fmt.Sprintf("Запрос погоды для координат: %.4f, %.4f", latVal, lonVal))

        weatherData, err := g.weather.GetWeather(latVal, lonVal)
        if err != nil {
            resultLabel.SetText(fmt.Sprintf("Ошибка: %s", err.Error()))
            g.log.Error("Ошибка получения погоды", err)
            return
        }

        resultLabel.SetText(fmt.Sprintf("Температура: %.2f°C", weatherData.Current.Temperature))
        g.log.Info(fmt.Sprintf("Получена погода: %.2f°C", weatherData.Current.Temperature))
    })

    saveLocationBtn := widget.NewButton("Сохранить координаты", func() {
        g.log.Debug("Кнопка 'Сохранить координаты' нажата")

        latVal, err1 := strconv.ParseFloat(latEntry.Text, 64)
        lonVal, err2 := strconv.ParseFloat(lonEntry.Text, 64)

        if err1 != nil || err2 != nil {
            resultLabel.SetText("Ошибка: введите корректные координаты")
            g.log.Error("Неверные координаты", fmt.Errorf("lat=%s, lon=%s", latEntry.Text, lonEntry.Text))
            return
        }

        if err := g.storage.SaveLocation(latVal, lonVal); err != nil {
            resultLabel.SetText(fmt.Sprintf("Ошибка сохранения: %s", err.Error()))
            g.log.Error("Ошибка сохранения координат", err)
            return
        }

        resultLabel.SetText("Координаты сохранены!")
        g.log.Info(fmt.Sprintf("Координаты сохранены: %.4f, %.4f", latVal, lonVal))
    })

    content := container.NewVBox(
        widget.NewLabel("Широта:"),
        latEntry,
        widget.NewLabel("Долгота:"),
        lonEntry,
        container.NewHBox(getWeatherBtn, saveLocationBtn),
        widget.NewSeparator(),
        resultLabel,
    )

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
    return nil
}
