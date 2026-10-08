# Telegram Bot - Go

Un chatbot de Telegram completo desarrollado en Go con arquitectura limpia y escalable.

## 📋 Características

- ✅ Arquitectura modular y escalable
- ✅ Base de datos PostgreSQL con GORM
- ✅ Caché con Redis
- ✅ Sistema de middleware (logging, autenticación, rate limiting)
- ✅ Manejo de comandos y mensajes
- ✅ Soporte para fotos y documentos
- ✅ Logging estructurado con Zap
- ✅ Configuración con variables de entorno
- ✅ Soporte para webhooks y long polling

## 🏗️ Arquitectura

```
telegram-bot/
├── config/          # Configuración del bot
│   └── config.go
├── models/          # Modelos de datos
│   ├── user.go
│   └── message.go
├── database/        # Conexión a base de datos
│   ├── database.go
│   └── redis.go
├── handlers/        # Manejadores de comandos y mensajes
│   ├── command_handler.go
│   └── message_handler.go
├── bot/            # Lógica principal del bot
│   └── bot.go
├── middleware/     # Middleware para autenticación, logging, etc.
│   ├── logging.go
│   ├── auth.go
│   └── rate_limit.go
├── utils/          # Utilidades y helpers
│   ├── logger.go
│   ├── helpers.go
│   └── validator.go
├── main.go         # Punto de entrada
├── go.mod          # Dependencias
└── .env.example    # Ejemplo de configuración
```

## 🚀 Instalación

1. Clonar el repositorio:
```bash
git clone https://github.com/nocturne-cybersecurity/telebot
cd telegram-bot
```

2. Instalar dependencias:
```bash
go mod download
```

3. Configurar variables de entorno:
```bash
cp .env.example .env
# Editar .env con tus configuraciones
```

4. Crear base de datos PostgreSQL:
```sql
CREATE DATABASE botdb;
```

5. Ejecutar el bot:
```bash
go run main.go
```

## ⚙️ Configuración

Variables de entorno requeridas:

- `TELEGRAM_BOT_TOKEN`: Token del bot de Telegram (obtenido de @BotFather)
- `DATABASE_URL`: URL de conexión a PostgreSQL
- `REDIS_URL`: URL de conexión a Redis
- `ENVIRONMENT`: Entorno (development/production)
- `LOG_LEVEL`: Nivel de logging (debug/info/warn/error)

## 📱 Comandos Disponibles

- `/start` - Iniciar el bot y registrar usuario
- `/help` - Mostrar ayuda y comandos disponibles
- `/about` - Información sobre el bot
- `/settings` - Configuración del usuario
- `/stats` - Estadísticas de uso (solo administradores)

## 🔧 Tecnologías

- **Go 1.21+** - Lenguaje de programación
- **Telebot v3** - Librería para Telegram Bot API
- **GORM** - ORM para PostgreSQL
- **Redis** - Caché y rate limiting
- **Zap** - Logging estructurado
- **Viper** - Gestión de configuración

## 📊 Base de Datos

### Tablas

- **users**: Información de usuarios
- **user_states**: Estados de conversación
- **messages**: Historial de mensajes

## 🛠️ Desarrollo

### Agregar nuevo comando

1. Crear handler en `handlers/command_handler.go`
2. Registrar el comando en `bot/bot.go`
3. Agregar documentación en README

### Agregar nuevo middleware

1. Crear archivo en `middleware/`
2. Implementar la interfaz de middleware
3. Aplicar en los handlers correspondientes

## 📝 Licencia

MIT License

## 👤 Autor

Tu nombre aquí
