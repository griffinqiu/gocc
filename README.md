# OpenCC for Go

[![Go](https://github.com/longbridge/opencc/workflows/Go/badge.svg)](https://github.com/longbridge/opencc/actions?query=workflow%3AGo)

This is a pure Go version of OpenCC([Open Chinese Convert 開放中文轉換](https://github.com/BYVoid/OpenCC/)) which is a project for conversion between Traditional and Simplified Chinese developed by [BYVoid](https://www.byvoid.com/). Avoid C library dependency and use Go Embed feature for embed dict into to binary for deploy easily.

此项目是基于 Native Go 方式实现 OpenCC，利用 OpenCC 官方的词典，减少 C 库对环境的依赖，同时，基于 Go Embed 特性，可以让我们编译的时候，直接将字典打包到 Go 的二进制里面，以获得较好的开发部署体验。

## Installation

```sh
go get github.com/longbridge/opencc
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/longbridge/opencc"
)

func main() {
    s2t, err := opencc.New("s2t")
    if err != nil {
        log.Fatal(err)
    }


    in := `自然语言处理是人工智能领域中的一个重要方向。`
    out, err := s2t.Convert(in)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n%s\n", in, out)
    //自然语言处理是人工智能领域中的一个重要方向。
    //自然語言處理是人工智能領域中的一個重要方向。
}
```

## 預設配置文件

- `s2t.json` Simplified Chinese to Traditional Chinese 簡體到繁體
- `t2s.json` Traditional Chinese to Simplified Chinese 繁體到簡體
- `s2tw.json` Simplified Chinese to Traditional Chinese (Taiwan Standard) 簡體到臺灣正體
- `tw2s.json` Traditional Chinese (Taiwan Standard) to Simplified Chinese 臺灣正體到簡體
- `s2hk.json` Simplified Chinese to Traditional Chinese (Hong Kong variant) 簡體到香港繁體
- `hk2s.json` Traditional Chinese (Hong Kong variant) to Simplified Chinese 香港繁體到簡體
- `s2twp.json` Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom 簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙
- `tw2sp.json` Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom 繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙
- `t2tw.json` Traditional Chinese (OpenCC Standard) to Taiwan Standard 繁體（OpenCC 標準）到臺灣正體
- `hk2t.json` Traditional Chinese (Hong Kong variant) to Traditional Chinese 香港繁體到繁體（OpenCC 標準）
- `t2hk.json` Traditional Chinese (OpenCC Standard) to Hong Kong variant 繁體（OpenCC 標準）到香港繁體
- `t2jp.json` Traditional Chinese Characters (Kyūjitai) to New Japanese Kanji (Shinjitai) 繁體（OpenCC 標準，舊字體）到日文新字體
- `jp2t.json` New Japanese Kanji (Shinjitai) to Traditional Chinese Characters (Kyūjitai) 日文新字體到繁體（OpenCC 標準，舊字體）
- `tw2t.json` Traditional Chinese (Taiwan standard) to Traditional Chinese 臺灣正體到繁體（OpenCC 標準）
- `s2hk-finance.json` 针对香港市场金融数据，做了特殊补充。

## Development Guides

- dictionary - 请勿修改！这个用来同步 OpenCC 官方的字典，请勿改动，这个文件夹应该是靠命令来生成的。
- addition-dictionary - 用来存放此项目提前修复的字典，执行 `make update:data` 的时候，会把这里的内容补充到 dictionary 里面。

采用 `make update:data` 命令可以从 OpenCC 官方仓库更新词典。

```bash
$ make update:data
```

## Benchmarks

See [benchmark_test.go](https://github.com/longbridge/opencc/tree/master/tests/benchmark_test.go)

### Short text (100 chars)

| Mode         | Number of Chars | Duration / op |
| ------------ | --------------- | ------------- |
| s2t          | 100             | 0.04 ms       |
| t2s          | 100             | 0.04 ms       |
| s2hk-finance | 100             | 0.07 ms       |
| s2tw         | 100             | 0.063 ms      |

### Long Text (14K)

Use 14K text to run benchmark.

| Mode         | Number of Chars | Duration / op |
| ------------ | --------------- | ------------- |
| s2t          | 14K             | 2.5 ms        |
| t2s          | 14K             | 2.8 ms        |
| s2hk-finance | 14K             | 5 ms          |
| s2tw         | 14K             | 3.8 ms        |

## License

Apache License
