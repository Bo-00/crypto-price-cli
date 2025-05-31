package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/yourusername/crypto-price-cli/api"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	limit        int
	listCurrency string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List top cryptocurrencies by market cap",
	Long: `List the top cryptocurrencies by market capitalization.
You can specify the number of coins to display and the currency.

Examples:
  crypto list
  crypto list --limit 20
  crypto list --currency eur
  crypto list -l 50 -c jpy`,
	Run: func(cmd *cobra.Command, args []string) {
		targetCurrency := "usd"
		if listCurrency != "" {
			targetCurrency = strings.ToLower(listCurrency)
		}

		client := api.NewCoinGeckoAPI()

		fmt.Printf("🔄 Fetching top %d cryptocurrencies...\n\n", limit)

		coins, err := client.GetTopCoins(limit, targetCurrency)
		if err != nil {
			color.Red("Error fetching cryptocurrency list: %v", err)
			return
		}

		displayCoinList(coins, targetCurrency)
	},
}

func displayCoinList(coins []api.PriceData, currency string) {
	currencySymbol := getCurrencySymbol(currency)

	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Symbol", "Price", "24h Change", "Market Cap", "Volume"})

	// Configure table appearance
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	// Add data rows
	for _, coin := range coins {
		// Format price
		price := fmt.Sprintf("%s%s", currencySymbol, formatNumber(coin.CurrentPrice))

		// Format 24h change with color
		changeStr := fmt.Sprintf("%.2f%%", coin.PriceChangePerc24h)
		if coin.PriceChangePerc24h > 0 {
			changeStr = "+" + changeStr
		}

		// Format market cap and volume
		marketCap := currencySymbol + formatLargeNumber(coin.MarketCap)
		volume := currencySymbol + formatLargeNumber(coin.TotalVolume)

		table.Append([]string{
			fmt.Sprintf("#%d", coin.MarketCapRank),
			coin.Name,
			strings.ToUpper(coin.Symbol),
			price,
			changeStr,
			marketCap,
			volume,
		})
	}

	// Display title
	color.New(color.FgCyan, color.Bold).Printf("📈 Top %d Cryptocurrencies (by Market Cap)\n", len(coins))
	color.New(color.FgBlue).Println(strings.Repeat("=", 50))

	table.Render()

	fmt.Printf("\n💡 Use 'crypto price <coin>' for detailed information\n")
	fmt.Printf("💡 Currency: %s\n\n", strings.ToUpper(currency))
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Number of cryptocurrencies to display (max 250)")
	listCmd.Flags().StringVarP(&listCurrency, "currency", "c", "usd", "Target currency for prices")
}
