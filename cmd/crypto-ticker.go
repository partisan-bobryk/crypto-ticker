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
	coinIds := []string{"bitcoin", "chainlink", "polkadot", "ethereum"}
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
		"DOT $%4.2f | LINK $%4.2f | BTC $%6.2f | ETH $%4.2f",
		jsonPayload["polkadot"]["usd"],
		jsonPayload["chainlink"]["usd"],
		jsonPayload["bitcoin"]["usd"],
		jsonPayload["ethereum"]["usd"],
	)
	systray.SetTitle(fmt.Sprintf(" %s ", ticker))
}

func registerCoins() {
	// testCoin1 := coinDisplay{
	// 	CoinStats: coinStats{
	// 		Name: "chainlink",
	// 	},
	// }
	// testCoin2 := coinDisplay{
	// 	CoinStats: coinStats{
	// 		Name: "polkadot",
	// 	},
	// }
	// coinDisplayList := []coinDisplay{testCoin1, testCoin2}

	go func() {
		for {
			setDisplayValues()
			time.Sleep(1 * time.Minute)
		}
	}()

	// for _, coin := range coinDisplayList {
	// 	coin.Register()
	// }
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
