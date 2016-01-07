package main

import (
	"github.com/f110/metro-delay/env"
	"log"
)

func main() {
	conf, err := NewConf("./conf.json")
	if err != nil {
		log.Print("could not load config file")
		log.Fatal(err)
	}

	watcher, err := NewMetroWatcher(conf)
	if err != nil {
		log.Print("could not create watcher instance")
		log.Fatal(err)
	}

	notifier, err := NewSlackNotifier(conf)
	if err != nil {
		log.Print("could not create notifier instance")
		log.Fatal(err)
	}

	if env.DEBUG {
		log.Print(notifier.Notify(RailwayFukutoshin, "start program"))
	}
	log.Fatal(watcher.Start(notifier))
}
