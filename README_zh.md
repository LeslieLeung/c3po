# C3PO - 由ChatGPT赋能

[English](README.md) | 简体中文

## 介绍

C3PO 使用 ChatGPT 作为后端生成 i18n 文件的翻译。

![](example/example.gif)

> 为您介绍C3PO，您在国际化和翻译方面的终极伙伴。 C3PO是一款强大的i18n翻译工具，消除语言障碍并帮助您与全世界的人们联系，让您的生活更轻松。C3PO基于最先进的技术，具有精准和及时的翻译功能。无论您是希望拓展全球业务的企业，还是仅仅对探索外语感兴趣的个人，C3PO都可以轻松应对。有了它用户友好的界面和高级的翻译能力，C3PO是您所有 i18n 翻译需求的首选解决方案。与C3PO一起跟语言障碍说再见，与机会无限的世界打招呼。
>
> -- 由ChatGPT生成，部分翻译有重新校对以保证流畅

## 安装

### go install

```bash
go install github.com/leslieleung/c3po
```

## 用法

在使用 C3PO 之前，您需要拥有 OpenAI API 密钥。然后，您需要通过 `export OPENAI_API_KEY="sk-xxxxxx"` 或 `echo "sk-xxxxxx">〜/.c3pocfg` 设置它。

### 翻译文本

```bash
c3po translate -t "text to translate" -l "en,de"
```

`-l "en,de"` 是ISO-639-1语言代码，用逗号分隔。

### 翻译文件

```bash
c3po translateFile -f "path/to/file"
```

当前仅支持 csv 格式。您的 csv 文件应如下所示：

```csv
key,zh,en,de,fr
hello,你好,,,
```

第一行应为标题，第一列应为翻译的key，第二列应为源语言。
然后是目标语言。

### 调试

使用带有 `-v` 选项的任何命令都将将日志级别设置为 `DEBUG` ，应为您提供足够的信息以进行调试。

## 路线图

- [ ] 支持GNU mo / po文件
- [ ] 稳定性改进

## 鸣谢

这个项目部分灵感来自于 [chatgpt-i18n](https://github.com/ObservedObservergpt-i18n)。它提供了一个 Web 界面来翻译 i18n 文件。
但是，仅支持json格式。我打算支持更多格式，例如GNU mo/po文件。