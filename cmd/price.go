package cmd

import (
	"fmt"
	"strings"

	"github.com/yourusername/crypto-price-cli/api"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var currency string

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price [cryptocurrency] [currency]",
	Short: "Get the current price of a cryptocurrency",
	Long: `Get the current price of a cryptocurrency in a specified currency.
The currency parameter is optional and defaults to USD.

Examples:
  crypto price bitcoin
  crypto price btc
  crypto price ethereum usd
  crypto price eth eur
  crypto price bitcoin jpy`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		coinQuery := args[0]
		targetCurrency := "usd"

		if len(args) > 1 {
			targetCurrency = strings.ToLower(args[1])
		}

		if currency != "" {
			targetCurrency = strings.ToLower(currency)
		}

		client := api.NewCoinGeckoAPI()

		// First, search for the coin ID
		coinID, err := client.SearchCoin(coinQuery)
		if err != nil {
			color.Red("Error: %v", err)
			return
		}

		// Get price data
		priceData, err := client.GetPrice(coinID, targetCurrency)
		if err != nil {
			color.Red("Error fetching price: %v", err)
			return
		}

		displayPriceInfo(priceData, targetCurrency)
	},
}

func displayPriceInfo(data *api.PriceData, currency string) {
	// Title
	color.New(color.FgCyan, color.Bold).Printf("\n%s (%s)\n", data.Name, strings.ToUpper(data.Symbol))
	color.New(color.FgBlue).Println(strings.Repeat("=", len(data.Name)+len(data.Symbol)+3))

	// Current Price
	currencySymbol := getCurrencySymbol(currency)
	priceColor := color.New(color.FgGreen, color.Bold)
	priceColor.Printf("💰 Current Price: %s%s\n", currencySymbol, formatNumber(data.CurrentPrice))

	// 24h Change
	changeColor := color.New(color.FgRed)
	changeIcon := "📉"
	if data.PriceChangePerc24h > 0 {
		changeColor = color.New(color.FgGreen)
		changeIcon = "📈"
	}

	changeColor.Printf("%s 24h Change: %s%s (%.2f%%)\n",
		changeIcon,
		currencySymbol,
		formatNumber(data.PriceChange24h),
		data.PriceChangePerc24h)

	// Market Data
	fmt.Printf("📊 Market Cap: %s%s\n", currencySymbol, formatLargeNumber(data.MarketCap))
	fmt.Printf("📊 Market Cap Rank: #%d\n", data.MarketCapRank)
	fmt.Printf("📊 24h Volume: %s%s\n", currencySymbol, formatLargeNumber(data.TotalVolume))

	// Last Updated
	color.New(color.FgWhite).Printf("\n🕒 Last Updated: %s\n\n", data.LastUpdated)
}

func getCurrencySymbol(currency string) string {
	symbols := map[string]string{
		"usd": "$",
		"eur": "€",
		"gbp": "£",
		"jpy": "¥",
		"cny": "¥",
		"krw": "₩",
		"inr": "₹",
		"cad": "C$",
		"aud": "A$",
		"chf": "CHF ",
		"btc": "₿",
		"eth": "Ξ",
	}

	if symbol, exists := symbols[strings.ToLower(currency)]; exists {
		return symbol
	}
	return strings.ToUpper(currency) + " "
}

func formatNumber(num float64) string {
	if num >= 1 {
		return fmt.Sprintf("%.2f", num)
	} else if num >= 0.01 {
		return fmt.Sprintf("%.4f", num)
	} else {
		return fmt.Sprintf("%.8f", num)
	}
}

func formatLargeNumber(num float64) string {
	if num >= 1e12 {
		return fmt.Sprintf("%.2fT", num/1e12)
	} else if num >= 1e9 {
		return fmt.Sprintf("%.2fB", num/1e9)
	} else if num >= 1e6 {
		return fmt.Sprintf("%.2fM", num/1e6)
	} else if num >= 1e3 {
		return fmt.Sprintf("%.2fK", num/1e3)
	}
	return formatNumber(num)
}

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.Flags().StringVarP(&currency, "currency", "c", "", "Target currency (default: usd)")
}
