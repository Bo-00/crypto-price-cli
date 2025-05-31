package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// 常见加密货币映射表，避免频繁API调用
var commonCoins = map[string]string{
	"bitcoin":          "bitcoin",
	"btc":              "bitcoin",
	"ethereum":         "ethereum",
	"eth":              "ethereum",
	"binancecoin":      "binancecoin",
	"bnb":              "binancecoin",
	"cardano":          "cardano",
	"ada":              "cardano",
	"solana":           "solana",
	"sol":              "solana",
	"xrp":              "ripple",
	"ripple":           "ripple",
	"polkadot":         "polkadot",
	"dot":              "polkadot",
	"dogecoin":         "dogecoin",
	"doge":             "dogecoin",
	"avalanche":        "avalanche-2",
	"avax":             "avalanche-2",
	"shiba":            "shiba-inu",
	"shib":             "shiba-inu",
	"polygon":          "matic-network",
	"matic":            "matic-network",
	"chainlink":        "chainlink",
	"link":             "chainlink",
	"litecoin":         "litecoin",
	"ltc":              "litecoin",
	"uniswap":          "uniswap",
	"uni":              "uniswap",
	"cosmos":           "cosmos",
	"atom":             "cosmos",
	"algorand":         "algorand",
	"algo":             "algorand",
	"tron":             "tron",
	"trx":              "tron",
	"stellar":          "stellar",
	"xlm":              "stellar",
	"monero":           "monero",
	"xmr":              "monero",
	"ethereum-classic": "ethereum-classic",
	"etc":              "ethereum-classic",
	"vechain":          "vechain",
	"vet":              "vechain",
	"filecoin":         "filecoin",
	"fil":              "filecoin",
	"tether":           "tether",
	"usdt":             "tether",
	"usd-coin":         "usd-coin",
	"usdc":             "usd-coin",
	"binance-usd":      "binance-usd",
	"busd":             "binance-usd",
}

// CoinGeckoAPI represents the CoinGecko API client
type CoinGeckoAPI struct {
	BaseURL string
	Client  *http.Client
}

// NewCoinGeckoAPI creates a new CoinGecko API client
func NewCoinGeckoAPI() *CoinGeckoAPI {
	return &CoinGeckoAPI{
		BaseURL: "https://api.coingecko.com/api/v3",
		Client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// PriceData represents cryptocurrency price data
type PriceData struct {
	ID                 string  `json:"id"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	CurrentPrice       float64 `json:"current_price"`
	MarketCap          float64 `json:"market_cap"`
	MarketCapRank      int     `json:"market_cap_rank"`
	TotalVolume        float64 `json:"total_volume"`
	PriceChange24h     float64 `json:"price_change_24h"`
	PriceChangePerc24h float64 `json:"price_change_percentage_24h"`
	LastUpdated        string  `json:"last_updated"`
}

// CoinList represents a simplified coin list item
type CoinList struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// GetPrice fetches the current price for a specific cryptocurrency
func (api *CoinGeckoAPI) GetPrice(coinID string, currency string) (*PriceData, error) {
	if currency == "" {
		currency = "usd"
	}

	url := fmt.Sprintf("%s/coins/markets?ids=%s&vs_currency=%s&order=market_cap_desc&per_page=1&page=1&sparkline=false",
		api.BaseURL, coinID, currency)

	resp, err := api.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var prices []PriceData
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(prices) == 0 {
		return nil, fmt.Errorf("cryptocurrency not found: %s", coinID)
	}

	return &prices[0], nil
}

// GetTopCoins fetches top cryptocurrencies by market cap
func (api *CoinGeckoAPI) GetTopCoins(limit int, currency string) ([]PriceData, error) {
	if currency == "" {
		currency = "usd"
	}
	if limit <= 0 || limit > 250 {
		limit = 10
	}

	url := fmt.Sprintf("%s/coins/markets?vs_currency=%s&order=market_cap_desc&per_page=%d&page=1&sparkline=false",
		api.BaseURL, currency, limit)

	resp, err := api.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var prices []PriceData
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return prices, nil
}

// SearchCoin searches for a cryptocurrency by symbol or name
// 优先使用内置映射表，减少网络请求
func (api *CoinGeckoAPI) SearchCoin(query string) (string, error) {
	queryLower := strings.ToLower(query)

	// 首先检查常见币种映射表
	if coinID, exists := commonCoins[queryLower]; exists {
		return coinID, nil
	}

	// 如果不在常见币种中，再通过API查询
	return api.searchCoinFromAPI(query)
}

// searchCoinFromAPI 通过API搜索加密货币
func (api *CoinGeckoAPI) searchCoinFromAPI(query string) (string, error) {
	// 使用搜索API而不是获取完整列表
	url := fmt.Sprintf("%s/search?query=%s", api.BaseURL, query)

	resp, err := api.Client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to search cryptocurrency: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var searchResult struct {
		Coins []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"coins"`
	}

	if err := json.Unmarshal(body, &searchResult); err != nil {
		return "", fmt.Errorf("failed to unmarshal search response: %w", err)
	}

	if len(searchResult.Coins) == 0 {
		return "", fmt.Errorf("cryptocurrency not found: %s", query)
	}

	queryLower := strings.ToLower(query)

	// 寻找精确匹配
	for _, coin := range searchResult.Coins {
		if strings.ToLower(coin.Symbol) == queryLower || strings.ToLower(coin.Name) == queryLower {
			return coin.ID, nil
		}
	}

	// 如果没有精确匹配，返回第一个结果
	return searchResult.Coins[0].ID, nil
}
