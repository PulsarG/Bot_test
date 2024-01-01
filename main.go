package main

import (
	//	"flag"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	tgClioent "Bot/clients/telegram"
	"Bot/consumer/event-consumer"
	"Bot/events/telegram"
	"Bot/storage/files"
)

// 5708011095:AAHJiuyPCem8MSmZqbKpJCFzR11xT3lEwIk

const (
	tgBotHost = "api.telegram.org"

	storagePage = "storage"
	bathSize    = 100
)

var isStart = false

//var infoPanel = widget.NewLabel("Bot Stropped")

func main() {
	startAppWindow()
}

func mustToken(token string) string {
	/* token := flag.String("tg-bot-token", "", "token for bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("bad token")
	} */

	return token
}

func startBot(token string) {
	eventsProcessor := telegram.New(tgClioent.New(tgBotHost, mustToken(token)), files.New(storagePage))

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, bathSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}

	isStart = true
}

func startAppWindow() {
	App := app.New()
	mainWindow := App.NewWindow("Bot")
	mainWindow.Resize(fyne.NewSize(250, 150))

	inputKey := widget.NewPasswordEntry()
	inputKey.PlaceHolder = "Token"
	startButton := widget.NewButton("Start", func() { startBot(inputKey.Text) })
	//infoPanel := widget.NewLabel("Bot Stropped")

	cont := container.NewVBox(inputKey, startButton)

	mainWindow.SetContent(cont)
	mainWindow.Show()
	App.Run()
}
