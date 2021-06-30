package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("CoinTicker")
	mQuit := systray.AddMenuItem("Exit", "Close Crypto Ticker")
	systray.AddSeparator()
	registerCoins()

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				fmt.Println("CoinTicker Closing")
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {}

func setDisplayValues() {
	coinIds := []string{"bitcoin", "stellar", "matic-network", "ethereum", "enjincoin"}
	queryIds := strings.Join(coinIds, "%2C")

	// Make a network call for the list of coins
	res, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd&include_24hr_change=true", queryIds))
	jsonPayload := make(map[string]map[string]interface{})

	if err != nil {
		fmt.Println("Failed to fetch CoinGecko Data")
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&jsonPayload)

	var ticker string

	ticker = fmt.Sprintf(
		"ENJ $%4.2f | MATIC $%4.2f | XLM $%4.2f | BTC $%6.2f | ETH $%4.2f",
		jsonPayload["enjincoin"]["usd"],
		jsonPayload["matic-network"]["usd"],
		jsonPayload["stellar"]["usd"],
		jsonPayload["bitcoin"]["usd"],
		jsonPayload["ethereum"]["usd"],
	)
	systray.SetTitle(fmt.Sprintf(" %s ", ticker))
}

func registerCoins() {
	go func() {
		for {
			setDisplayValues()
			time.Sleep(1 * time.Minute)
		}
	}()
}

type coinStats struct {
	Name         string
	Change24h    *string
	CurrentValue *string
}

type coinDisplay struct {
	CoinStats coinStats
	MenuItem  *systray.MenuItem
}

func (cd *coinDisplay) Register() {
	cd.MenuItem = systray.AddMenuItem(cd.CoinStats.Name, cd.CoinStats.Name)
}
