# Telegram Bot em Go

Um bot de Telegram feito em Go que responde a comandos úteis como lembretes, previsão do tempo, cotação de ações e criptomoedas.

![Go Version](https://img.shields.io/badge/go-1.24-blue)
![Last Commit](https://img.shields.io/github/last-commit/douglashauschild/telegram-bot-go)

## 📦 Funcionalidades
- ```/start``` → Mensagem de boas-vindas e ajuda.
- ```/reminder <texto>``` → Cria um lembrete simples.
- ```/weather <cidade>``` → Informa a previsão do tempo (mock).
- ```/stock <código>``` → Mostra a cotação de uma ação (mock).
- ```/crypto <moeda>``` → Mostra a cotação de uma criptomoeda (mock).

## 🚀 Como executar o projeto

### 1. Pré-requisitos
- Go 1.18 ou superior.
- Conta no Telegram.
  
### 2. Clonar o repositório

```bash
git clone https://github.com/douglashauschild/telegram-go-bot.git
cd telegram-go-bot
```

### 3. Instalar dependências
```bash
go mod tidy
```

### 4. Configurar o arquivo ```.env```
```bash
TELEGRAM_TOKEN=123456789:ABCDEFseuTokenAqui
```

### 5. Obter o token do Telegram
1. Abra o Telegram e procure por: ```@BotFather```.
2. Digite: ```/start```.
3. Em seguida, envie: ```/newbot```.
4. Escolha um nome e um username para o seu bot (o username precisa terminar com bot, por exemplo: ```MeuTesteBot```).
5. O BotFather vai te enviar um **token** no formato:
```bash
123456789:ABCDEFseuTokenAqui
```
6. Copie esse token e coloque no arquivo ```.env``` como mostrado acima.

### 6. Executar o projeto
```bash
go run main.go
```

# 📁 Estrutura do Projeto
```bash
telegram-go-bot/
├── .env
├── go.mod
├── go.sum
├── main.go
└── README.md
```

# 📝 Notas
- Este projeto utiliza respostas **mockadas** (fixas) para simplificar a lógica.
- Para integração com APIs reais (clima, finanças), podemos evoluir facilmente.
- Variáveis de ambiente são gerenciadas com o pacote ```godotenv```.

# 🤝 Contribuindo
1. Faça um fork.
2. Crie sua feature branch: ```git checkout -b minha-feature```.
3. Commit: ```git commit -m 'Minha feature'```.
4. Push: ```git push origin minha-feature```.
5. Abra um Pull Request.
