# lab5 - Приложение погоды с CLI, HTTP и GUI

## Запуск

- CLI: `make run-cli`
- HTTP: `make run-http` (сервер на :8080)
- GUI: `make run-gui`

## API

- `GET /weather` — текущая погода
- `GET /location` — текущие координаты
- `POST /location?lat=...&lon=...` — сохранить координаты
