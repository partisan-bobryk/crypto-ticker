# crypto-ticker
Display price of certain crypto coins. Source of price from coingecko.

## Installation

__MacOS__
1. Pull down repo
2. Use the cli to run the program by executing `./bin/coin-watch`. A process will start and a system tray should contain a list of cryto with their price.

__Linux__
1. Pull down repo
2. Make sure you have libgtk-3 installed. Install by running `sudo apt-get install gcc libgtk-3-dev libappindicator3-dev`.
3. Using latest go version run `go build -o ./bin/coin-watch ./cmd/crypto-ticker.go` in the root directory. 
4. Run the compiled executable. `./bin/coin-watch`
