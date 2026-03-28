package http

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/kornev-aa/lab5-cache/internal/pkg/weather"
    "github.com/kornev-aa/lab5-cache/pkg/logger"
    "github.com/kornev-aa/lab5-cache/pkg/storage"
)

type Handlers struct {
    log     *logger.Logger
    storage storage.LocationStorage
    weather *weather.WeatherService
}

func NewHandlers(log *logger.Logger, storage storage.LocationStorage) *Handlers {
    return &Handlers{
        log:     log,
        storage: storage,
        weather: weather.NewWeatherService(),
    }
}

func (h *Handlers) GetWeather(w http.ResponseWriter, r *http.Request) {
    h.log.Debug("GET /weather")

    lat, err := h.storage.GetLatitude()
    if err != nil {
        h.log.Error("Failed to get latitude", err)
        http.Error(w, "Failed to get latitude", http.StatusInternalServerError)
        return
    }

    lon, err := h.storage.GetLongitude()
    if err != nil {
        h.log.Error("Failed to get longitude", err)
        http.Error(w, "Failed to get longitude", http.StatusInternalServerError)
        return
    }

    h.log.Debug("Getting weather for coordinates: " + formatCoords(lat, lon))

    result, err := h.weather.GetWeather(lat, lon)
    if err != nil {
        h.log.Error("Failed to get weather", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func (h *Handlers) SaveLocation(w http.ResponseWriter, r *http.Request) {
    h.log.Debug("POST /location")

    latStr := r.URL.Query().Get("lat")
    lonStr := r.URL.Query().Get("lon")

    lat, err := strconv.ParseFloat(latStr, 64)
    if err != nil {
        h.log.Error("Invalid latitude: "+latStr, err)
        http.Error(w, "Invalid latitude", http.StatusBadRequest)
        return
    }

    lon, err := strconv.ParseFloat(lonStr, 64)
    if err != nil {
        h.log.Error("Invalid longitude: "+lonStr, err)
        http.Error(w, "Invalid longitude", http.StatusBadRequest)
        return
    }

    if err := h.storage.SaveLocation(lat, lon); err != nil {
        h.log.Error("Failed to save location", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    h.log.Info("Location saved: " + formatCoords(lat, lon))
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func (h *Handlers) GetLocation(w http.ResponseWriter, r *http.Request) {
    h.log.Debug("GET /location")

    lat, err := h.storage.GetLatitude()
    if err != nil {
        h.log.Error("Failed to get latitude", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    lon, err := h.storage.GetLongitude()
    if err != nil {
        h.log.Error("Failed to get longitude", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    result := map[string]float64{"latitude": lat, "longitude": lon}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func formatCoords(lat, lon float64) string {
    return "lat=" + strconv.FormatFloat(lat, 'f', 4, 64) + ", lon=" + strconv.FormatFloat(lon, 'f', 4, 64)
}
