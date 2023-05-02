/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
    and usage of using your command. For example:

    Cobra is a CLI library for Go that empowers applications.
    This application is a tool to generate the needed files
    to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started\n", appVersion)
		rand.Seed(time.Now().UnixNano())

		settings := telebot.Settings{
			Token:  os.Getenv("TELEGRAM_TOKEN"),
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}

		kbot, err := telebot.NewBot(settings)

		if err != nil {
			log.Fatal(err)
			return
		}

		kbot.Handle("/start", func(m telebot.Context) error {
			err = m.Send(fmt.Sprintf("Hello, I'm Kbot %s!", appVersion))
			return err
		})

		kbot.Handle("/help", func(m telebot.Context) error {
			err = m.Send(fmt.Sprintf("I'm dead simple bot, I can echo back everything you type, " +
				"also I understand /start /help /randpic commands.\n" +
				"On /randpic I will send you 1 random image of 1000, with size 200x300"))
			return err
		})

		kbot.Handle("/randpic", func(m telebot.Context) error {
			url := "https://picsum.photos/id/" + strconv.Itoa(rand.Intn(1000)) + "/200/300"
			log.Println("got /randpic request, serving from", url)
			photo := &telebot.Photo{File: telebot.FromURL(url)}
			err = m.Send(photo)
			return err
		})

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Println(m.Message().Payload, m.Text())
			err = m.Send(m.Text())
			return err
		})

		kbot.Start()

	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
