# 🚀 发布指南

本文档介绍如何发布和分享你的 Crypto Price CLI 工具。

## 📦 发布平台

### 1. GitHub Releases（推荐首选）

#### 优势

- ✅ 免费且易用
- ✅ 自动构建多平台二进制文件
- ✅ 版本管理和发布说明
- ✅ 下载统计

#### 发布步骤

1. **创建 GitHub 仓库**

   ```bash
   # 在 https://github.com/new 创建新仓库
   # 仓库名建议: crypto-price-cli
   ```

2. **推送代码**

   ```bash
   git init
   git add .
   git commit -m "Initial commit: Crypto Price CLI v1.0.0"
   git remote add origin https://github.com/yourusername/crypto-price-cli.git
   git branch -M main
   git push -u origin main
   ```

3. **创建版本标签**

   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

4. **自动构建发布**
   - GitHub Actions 会自动构建多平台二进制文件
   - 自动创建 Release 页面
   - 包含详细的安装说明

### 2. Go 模块发布

#### 优势

- ✅ Go 开发者可以直接 `go install`
- ✅ 版本管理和依赖解析
- ✅ 集成到 Go 生态系统

#### 发布步骤

1. **确保模块路径正确**

   ```go
   // go.mod
   module github.com/yourusername/crypto-price-cli
   ```

2. **推送到 GitHub**

   ```bash
   git push origin main
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **用户安装**
   ```bash
   go install github.com/yourusername/crypto-price-cli@latest
   ```

### 3. Homebrew 发布（macOS）

#### 优势

- ✅ macOS 用户友好
- ✅ 自动依赖管理
- ✅ 简单的安装命令

#### 发布步骤

1. **创建 Formula**

   ```ruby
   # crypto-price-cli.rb
   class CryptoPriceCli < Formula
     desc "A powerful CLI tool for querying cryptocurrency prices"
     homepage "https://github.com/yourusername/crypto-price-cli"
     url "https://github.com/yourusername/crypto-price-cli/releases/download/v1.0.0/crypto-darwin-amd64"
     sha256 "your-sha256-hash"
     version "1.0.0"

     def install
       bin.install "crypto-darwin-amd64" => "crypto"
     end

     test do
       system "#{bin}/crypto", "--version"
     end
   end
   ```

2. **提交到 Homebrew**
   - Fork [homebrew-core](https://github.com/Homebrew/homebrew-core)
   - 添加 Formula 文件
   - 提交 Pull Request

### 4. 其他包管理器

#### Arch Linux (AUR)

```bash
# 创建 PKGBUILD 文件
# 提交到 AUR
```

#### Snap (Ubuntu)

```yaml
# snapcraft.yaml
name: crypto-price-cli
version: "1.0.0"
summary: Cryptocurrency price CLI tool
description: |
  A powerful CLI tool for querying cryptocurrency prices
```

#### Chocolatey (Windows)

```xml
<!-- crypto-price-cli.nuspec -->
<package>
  <metadata>
    <id>crypto-price-cli</id>
    <version>1.0.0</version>
    <title>Crypto Price CLI</title>
    <description>Cryptocurrency price query tool</description>
  </metadata>
</package>
```

## 🌟 推广策略

### 1. 社区分享

#### Reddit

- r/golang
- r/commandline
- r/cryptocurrency
- r/programming

#### 论坛

- Go 语言中文网
- V2EX
- 掘金
- 知乎

#### 社交媒体

- Twitter/X
- LinkedIn
- 微博

### 2. 技术博客

#### 文章主题

- "用 Go 构建加密货币价格查询 CLI 工具"
- "Go 语言 CLI 开发最佳实践"
- "从零开始发布一个 Go CLI 工具"

#### 发布平台

- 掘金
- CSDN
- 博客园
- Medium

### 3. 产品展示

#### Demo 视频

- 录制使用演示
- 上传到 YouTube/B 站
- 添加到 README

#### 截图

- 制作美观的 CLI 截图
- 展示主要功能
- 用于文档和推广

## 📊 数据跟踪

### 1. GitHub 指标

- Stars 数量
- Fork 数量
- Issues 和 PR
- 下载次数

### 2. 使用统计

```go
// 可选：添加匿名使用统计
func trackUsage() {
    // 发送匿名使用数据
}
```

### 3. 用户反馈

- GitHub Issues
- 用户调研
- 功能请求

## 🔄 持续改进

### 1. 版本发布流程

```bash
# 1. 开发新功能
git checkout -b feature/new-feature

# 2. 测试和代码审查
go test ./...
golangci-lint run

# 3. 合并到主分支
git checkout main
git merge feature/new-feature

# 4. 创建新版本
git tag v1.1.0
git push origin v1.1.0

# 5. 更新文档
```

### 2. 功能路线图

- [ ] 添加更多交易所数据源
- [ ] 支持价格提醒
- [ ] 添加图表显示
- [ ] 支持投资组合跟踪

### 3. 社区建设

- 响应用户问题
- 接受贡献
- 维护文档
- 发布更新

## 💡 推广文案模板

### GitHub README 徽章

```markdown
[![GitHub stars](https://img.shields.io/github/stars/yourusername/crypto-price-cli.svg)](https://github.com/yourusername/crypto-price-cli/stargazers)
[![GitHub downloads](https://img.shields.io/github/downloads/yourusername/crypto-price-cli/total.svg)](https://github.com/yourusername/crypto-price-cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/crypto-price-cli)](https://goreportcard.com/report/github.com/yourusername/crypto-price-cli)
```

### 社交媒体文案

```
🚀 刚刚发布了一个用 Go 语言开发的加密货币价格查询 CLI 工具！

✨ 特性：
- 实时价格查询
- 多货币支持
- 美观的界面
- 跨平台支持

GitHub: https://github.com/yourusername/crypto-price-cli

#golang #cli #cryptocurrency #opensource
```

### 博客文章大纲

1. 项目背景和动机
2. 技术选型和架构设计
3. 开发过程和挑战
4. Go 语言最佳实践
5. 发布和推广经验
6. 未来计划

通过以上多种方式，你的 CLI 工具可以获得更多的关注和使用！
