package main

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"

	tgClioent "Bot/clients/telegram"
	"Bot/consumer/event-consumer"
	"Bot/events/telegram"
	"Bot/storage/files"

	cnf "github.com/PulsarG/ConfigManager"
)

// 5708011095:AAHJiuyPCem8MSmZqbKpJCFzR11xT3lEwIk

const (
	tgBotHost = "api.telegram.org"

	storagePage = "storage"
	bathSize    = 100
)

var isStart = false
var ch1 = make(chan string)
var ctx = context.Background()

func main() {
	startAppWindow()
}

func startBot(ctx context.Context, token string) {
	select {
	case <-ctx.Done():
		log.Print("service hand stopped")
		return
	default:
		eventsProcessor := telegram.New(tgClioent.New(tgBotHost, token), files.New(storagePage))

		log.Print("service started")

		consumer := event_consumer.New(eventsProcessor, eventsProcessor, bathSize)

		if err := consumer.Start(); err != nil {
			log.Fatal("service stopped", err)
		}
		isStart = true
	}
}

func controlBot(ctx context.Context, cancel context.CancelFunc, token string) {
	go startBot(ctx, token)
}

func startAppWindow() {
	App := app.New()
	mainWindow := App.NewWindow("Bot")
	mainWindow.Resize(fyne.NewSize(250, 150))

	ctx, cancel := context.WithCancel(ctx)

	inputKey := widget.NewPasswordEntry()
	inputKey.PlaceHolder = "Token"

	startButton := widget.NewButton("Start", func() { controlBot(ctx, cancel, cnf.GetFromIni("TOKENS", "testBotPulsar")) })

	cont := container.NewVBox(inputKey, startButton)

	mainWindow.SetContent(cont)
	mainWindow.Show()
	App.Run()
}
