# 诗吟 shiyin

终端诗词阅读器，在命令行中静心读诗。

![demo](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go) ![license](https://img.shields.io/badge/license-MIT-green)

## 功能

- 无边框，诗词居中显示，界面纯净
- 收录唐诗三百首、宋词三百首
- 支持翻页、随机跳转
- 启动时选择诗词集合，或通过 CLI 参数直接进入
- 支持乱序模式

## 预览

```
              静  夜  思
           ────────────
               李  白
            唐代 · 五言绝句

        床前明月光，疑是地上霜。
        举头望明月，低头思故乡。
```

## 安装

**从源码构建：**

```bash
git clone https://github.com/Wang-mis/shiyin.git
cd shiyin
go build -o shiyin ./cmd/
```

需要 Go 1.21 或更高版本。

## 用法

```bash
shiyin              # 启动集合选择器
shiyin tang         # 直接打开唐诗三百首
shiyin ci           # 直接打开宋词三百首
shiyin all          # 全部诗词
shiyin all -s       # 全部诗词，乱序
```

## 快捷键

| 按键 | 功能 |
|------|------|
| `→` / `l` | 下一首 |
| `←` / `j` | 上一首 |
| `r` | 随机跳转 |
| `h` | 切换底部帮助栏 |
| `Esc` | 返回集合选择 |
| `q` / `Ctrl+C` | 退出 |

## 数据来源

诗词数据来自 [chinese-poetry/huajianji](https://github.com/chinese-poetry/huajianji)。

## License

MIT
