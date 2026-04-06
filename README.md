# 诗吟 shiyin

终端诗词阅读器，在命令行中静心读诗。

![demo](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go) ![license](https://img.shields.io/badge/license-MIT-green)

## 功能

- 无边框，诗词居中显示，界面纯净
- 收录唐诗三百首、宋词三百首
- 支持翻页、随机跳转
- 收藏夹，按 `f` 收藏/取消，跨会话持久化
- 启动时选择诗词集合，或通过 CLI 参数直接进入
- 支持乱序模式

## 预览

![example](example.jpg)

## 安装

### 下载预编译二进制

前往 [Releases](https://github.com/Wang-mis/shiyin/releases) 下载对应平台的文件：

| 平台 | 文件 |
|------|------|
| Windows (x64) | `shiyin-windows-amd64.exe` |
| Linux (x64) | `shiyin-linux-amd64` |
| Linux (ARM64) | `shiyin-linux-arm64` |
| macOS (Intel) | `shiyin-darwin-amd64` |
| macOS (Apple Silicon) | `shiyin-darwin-arm64` |

Linux / macOS 下载后需添加执行权限：

```bash
chmod +x shiyin-*
```

### 从源码构建

需要 Go 1.21 或更高版本。

```bash
git clone https://github.com/Wang-mis/shiyin.git
cd shiyin
go build -o shiyin ./cmd/
```

## 用法

```bash
shiyin              # 启动集合选择器
shiyin tang         # 直接打开唐诗三百首
shiyin ci           # 直接打开宋词三百首
shiyin all          # 全部诗词
shiyin fav          # 直接打开收藏夹
shiyin all -s       # 全部诗词，乱序
```

## 快捷键

| 按键 | 功能 |
|------|------|
| `→` / `l` | 下一首 |
| `←` / `p` | 上一首 |
| `r` | 随机跳转 |
| `f` | 收藏 / 取消收藏当前诗词 |
| `o` | 在浏览器中打开当前诗词的古文岛详情页 |
| `h` | 切换底部帮助栏 |
| `Esc` | 返回集合选择 |
| `q` / `Ctrl+C` | 退出 |

收藏数据存储在：
- Windows：`%APPDATA%\shiyin\favorites.json`
- Linux / macOS：`~/.config/shiyin/favorites.json`

## 数据来源

诗词数据来自 [chinese-poetry/huajianji](https://github.com/chinese-poetry/huajianji)。

## License

MIT
