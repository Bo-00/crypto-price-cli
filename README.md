# �� Crypto Price CLI

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/yourusername/crypto-price-cli.svg)](https://github.com/yourusername/crypto-price-cli/releases)
[![Downloads](https://img.shields.io/github/downloads/yourusername/crypto-price-cli/total.svg)](https://github.com/yourusername/crypto-price-cli/releases)

一个功能强大的加密货币价格查询命令行工具，支持实时价格查询、市场数据显示等功能。

![Demo](https://via.placeholder.com/800x400/1e1e1e/ffffff?text=CLI+Demo+Screenshot)

## ✨ 特性

- 🔍 **实时价格查询** - 支持查询任何加密货币的实时价格
- 📊 **市场数据** - 显示市值、交易量、24 小时变化等详细信息
- 🌍 **多币种支持** - 支持 USD、EUR、JPY、CNY 等多种法币
- 📈 **排行榜** - 查看按市值排序的加密货币排行榜
- 🎨 **美观界面** - 彩色输出和表格显示，用户体验友好
- ⚡ **快速响应** - 基于 CoinGecko API，数据准确可靠
- 🚀 **跨平台** - 支持 Windows、macOS、Linux

## 🛠️ 安装

### 方法 1：下载预编译二进制文件（推荐）

前往 [Releases](https://github.com/yourusername/crypto-price-cli/releases) 页面下载适合你系统的版本：

#### macOS

```bash
# Apple Silicon (M1/M2)
curl -L https://github.com/yourusername/crypto-price-cli/releases/latest/download/crypto-darwin-arm64 -o crypto
chmod +x crypto
sudo mv crypto /usr/local/bin/

# Intel
curl -L https://github.com/yourusername/crypto-price-cli/releases/latest/download/crypto-darwin-amd64 -o crypto
chmod +x crypto
sudo mv crypto /usr/local/bin/
```

#### Linux

```bash
curl -L https://github.com/yourusername/crypto-price-cli/releases/latest/download/crypto-linux-amd64 -o crypto
chmod +x crypto
sudo mv crypto /usr/local/bin/
```

#### Windows

下载 `crypto-windows-amd64.exe`，重命名为 `crypto.exe` 并添加到 PATH。

### 方法 2：通过 Go 安装

```bash
go install github.com/yourusername/crypto-price-cli@latest
```

### 方法 3：从源码构建

```bash
git clone https://github.com/yourusername/crypto-price-cli.git
cd crypto-price-cli
chmod +x install.sh
./install.sh
```

## 📖 使用方法

### 快速开始

```bash
# 查询比特币价格
crypto price bitcoin

# 查看前10大加密货币
crypto list

# 获取帮助
crypto --help
```

### 高级用法

```bash
# 查询以太坊的欧元价格
crypto price ethereum eur

# 查看前20名的人民币价格
crypto list --limit 20 --currency cny

# 使用简写
crypto price btc
crypto list -l 50 -c jpy
```

更多使用方法请查看 [USAGE.md](USAGE.md)

## 🌍 支持的货币

- **法币**: USD, EUR, GBP, JPY, CNY, KRW, INR, CAD, AUD, CHF
- **加密货币**: BTC, ETH（作为计价单位）

## 🔧 技术栈

- **语言**: Go 1.21+
- **CLI 框架**: [Cobra](https://github.com/spf13/cobra)
- **配置管理**: [Viper](https://github.com/spf13/viper)
- **颜色输出**: [Fatih Color](https://github.com/fatih/color)
- **表格显示**: [TableWriter](https://github.com/olekukonko/tablewriter)
- **数据源**: [CoinGecko API](https://www.coingecko.com/en/api)

## 📝 Go 语言学习要点

此项目展示了以下 Go 语言最佳实践：

1. **项目结构**

   - 清晰的包分离（cmd, api）
   - 合理的文件组织

2. **错误处理**

   - Go 风格的错误处理
   - 错误包装和传递

3. **接口设计**

   - HTTP 客户端封装
   - API 抽象层

4. **并发安全**

   - HTTP 客户端配置
   - 超时处理

5. **命令行工具**
   - Cobra 框架使用
   - 标志和参数处理

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本项目
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 🐛 问题反馈

如果遇到问题，请在 [Issues](https://github.com/yourusername/crypto-price-cli/issues) 中报告。

## 📄 许可证

MIT License - 查看 [LICENSE](LICENSE) 文件了解详情

## ⭐ Star History

如果这个项目对你有帮助，请给个 Star ⭐️

## 🙏 致谢

- [CoinGecko](https://www.coingecko.com/) - 提供免费的加密货币 API
- [Cobra](https://github.com/spf13/cobra) - 优秀的 CLI 框架
- [Go 语言学习笔记](https://github.com/qyuhen/book) - 雨痕老师的经典教程
