package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente padrão")
	}

	botToken := os.Getenv("TELEGRAM_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_TOKEN não definido")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Bot autorizado em conta %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			msg.Text = handleStart()
		case strings.HasPrefix(update.Message.Text, "/reminder"):
			msg.Text = handleReminder(update.Message.Text)
		case strings.HasPrefix(update.Message.Text, "/weather"):
			msg.Text = handleWeather(update.Message.Text)
		case strings.HasPrefix(update.Message.Text, "/stock"):
			msg.Text = handleStock(update.Message.Text)
		case strings.HasPrefix(update.Message.Text, "/crypto"):
			msg.Text = handleCrypto(update.Message.Text)
		default:
			msg.Text = "Comando não reconhecido. Use /reminder, /weather, /stock ou /crypto"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}

func handleStart() string {
	return `👋 Olá! Eu sou seu bot pessoal.

	Você pode usar os seguintes comandos:
	/reminder <texto> - Criar um lembrete.
	/weather <cidade> - Ver a previsão do tempo.
	/stock <código> - Consultar uma ação.
	/crypto <moeda> - Consultar uma criptomoeda.

	Digite um comando para começar!`
}

func handleReminder(text string) string {
	args := strings.TrimSpace(strings.TrimPrefix(text, "/reminder"))
	if args == "" {
		return "Por favor, forneça um lembrete. Ex: /reminder Comprar leite"
	}
	// Aqui poderia salvar o lembrete num banco de dados ou agendar tarefa.
	return "Lembrete criado: " + args
}

func handleWeather(text string) string {
	args := strings.TrimSpace(strings.TrimPrefix(text, "/weather"))
	city := "sua cidade"
	if args != "" {
		city = args
	}
	// Aqui poderia chamar uma API real de clima (OpenWeatherMap por exemplo)
	return "Previsão para " + city + ": ☀️ 25°C, céu limpo."
}

func handleStock(text string) string {
	args := strings.TrimSpace(strings.TrimPrefix(text, "/stock"))
	stock := "PETR4"
	if args != "" {
		stock = strings.ToUpper(args)
	}
	// Aqui poderia chamar uma API real de bolsa
	return "Cotação de " + stock + ": R$ 28,50 📈"
}

func handleCrypto(text string) string {
	args := strings.TrimSpace(strings.TrimPrefix(text, "/crypto"))
	crypto := "BTC"
	if args != "" {
		crypto = strings.ToUpper(args)
	}
	// Aqui poderia chamar uma API real de cripto
	return "Cotação de " + crypto + ": US$ 65.000 🚀"
}
