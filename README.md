# Telegram Bot Go

A simple Telegram bot built with Go (Golang), using the official [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api).  
This project demonstrates how to build and run a basic bot that interacts with users through messages and commands.

![Go Version](https://img.shields.io/badge/go-1.24-blue)
![Last Commit](https://img.shields.io/github/last-commit/douglashauschild/telegram-bot-go)

## 🚀 Getting Started

### 📦 Requirements

- Go 1.18+
- Telegram account and bot token (get one at [BotFather](https://t.me/botfather))

### 🛠️ Installation

```bash
git clone https://github.com/douglashauschild/telegram-bot-go.git
cd telegram-bot-go
go mod tidy
```
### ▶️ Running the Bot
Update the ```.env``` file with your Telegram bot token:
```go
TELEGRAM_TOKEN=123456789:ABCDEFseuTokenAqui
```
Then run the bot:
```bach
go run main.go
```
## 💡 Features
- Listens to incoming messages
- Responds to specific keywords and commands
- Simple and extensible architecture

## 📁 Project Structure
```go
telegram-bot-go/
  ├── main.go         # Bot entry point
  ├── go.mod          # Go module definition
```

## 📚 References
- [Go Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api)
- [Telegram Bot API Documentation](https://core.telegram.org/bots/api)

## 🧪 Next Steps
- Add support for more commands
- Improve error handling and logging
- Add modular command handlers

## 👨🏻‍💻 Author
Douglas Hauschild  
[LinkedIn](https://www.linkedin.com/in/douglas-hauschild-66449122b/) | [GitHub](https://github.com/douglashauschild)

> Feel free to fork and customize this project to build your own Telegram bot!
