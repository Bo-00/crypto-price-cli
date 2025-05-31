# Crypto Price CLI 使用指南

## 基本命令

### 1. 查询加密货币价格

```bash
# 查询比特币价格（默认USD）
crypto price bitcoin
crypto price btc

# 查询以太坊价格
crypto price ethereum
crypto price eth

# 指定货币单位
crypto price bitcoin eur      # 欧元
crypto price eth jpy          # 日元
crypto price btc cny          # 人民币

# 使用标志指定货币
crypto price bitcoin --currency eur
crypto price btc -c cny
```

### 2. 查看加密货币排行榜

```bash
# 查看前10名
crypto list

# 查看前20名
crypto list --limit 20
crypto list -l 20

# 指定显示货币
crypto list --currency eur
crypto list -c jpy

# 组合使用
crypto list -l 50 -c eur
```

## 支持的加密货币（内置快速查询）

以下加密货币可以直接使用符号或全名查询，无需网络搜索：

| 名称             | 符号  | 支持的查询方式            |
| ---------------- | ----- | ------------------------- |
| Bitcoin          | BTC   | `bitcoin`, `btc`          |
| Ethereum         | ETH   | `ethereum`, `eth`         |
| Binance Coin     | BNB   | `binancecoin`, `bnb`      |
| Cardano          | ADA   | `cardano`, `ada`          |
| Solana           | SOL   | `solana`, `sol`           |
| XRP              | XRP   | `xrp`, `ripple`           |
| Polkadot         | DOT   | `polkadot`, `dot`         |
| Dogecoin         | DOGE  | `dogecoin`, `doge`        |
| Avalanche        | AVAX  | `avalanche`, `avax`       |
| Shiba Inu        | SHIB  | `shiba`, `shib`           |
| Polygon          | MATIC | `polygon`, `matic`        |
| Chainlink        | LINK  | `chainlink`, `link`       |
| Litecoin         | LTC   | `litecoin`, `ltc`         |
| Uniswap          | UNI   | `uniswap`, `uni`          |
| Cosmos           | ATOM  | `cosmos`, `atom`          |
| Algorand         | ALGO  | `algorand`, `algo`        |
| TRON             | TRX   | `tron`, `trx`             |
| Stellar          | XLM   | `stellar`, `xlm`          |
| Monero           | XMR   | `monero`, `xmr`           |
| Ethereum Classic | ETC   | `ethereum-classic`, `etc` |
| VeChain          | VET   | `vechain`, `vet`          |
| Filecoin         | FIL   | `filecoin`, `fil`         |
| Tether           | USDT  | `tether`, `usdt`          |
| USD Coin         | USDC  | `usd-coin`, `usdc`        |
| Binance USD      | BUSD  | `binance-usd`, `busd`     |

## 支持的货币单位

### 法定货币

- `usd` - 美元 ($)
- `eur` - 欧元 (€)
- `gbp` - 英镑 (£)
- `jpy` - 日元 (¥)
- `cny` - 人民币 (¥)
- `krw` - 韩元 (₩)
- `inr` - 印度卢比 (₹)
- `cad` - 加拿大元 (C$)
- `aud` - 澳元 (A$)
- `chf` - 瑞士法郎 (CHF)

### 加密货币计价

- `btc` - 比特币 (₿)
- `eth` - 以太坊 (Ξ)

## 示例使用场景

### 日常查询

```bash
# 快速查看比特币价格
crypto price btc

# 查看以太坊的人民币价格
crypto price eth cny

# 查看前10大加密货币
crypto list
```

### 投资分析

```bash
# 查看详细的比特币信息
crypto price bitcoin

# 对比不同货币单位的价格
crypto price eth usd
crypto price eth eur
crypto price eth btc

# 查看更多排行榜信息
crypto list -l 50
```

### 国际用户

```bash
# 欧洲用户
crypto price bitcoin eur
crypto list -c eur

# 日本用户
crypto price bitcoin jpy
crypto list -c jpy

# 中国用户
crypto price bitcoin cny
crypto list -c cny
```

## 网络配置

如果遇到网络连接问题，可以设置代理：

```bash
# 设置HTTP代理
export https_proxy=http://127.0.0.1:7890
export http_proxy=http://127.0.0.1:7890

# 设置SOCKS5代理
export all_proxy=socks5://127.0.0.1:7890

# 然后使用crypto命令
crypto price bitcoin
```

## 卸载

如果需要卸载 CLI 工具：

```bash
sudo rm /usr/local/bin/crypto
```

## 获取帮助

```bash
# 查看总体帮助
crypto --help

# 查看特定命令帮助
crypto price --help
crypto list --help
```
