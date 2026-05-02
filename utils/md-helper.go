package utils

import (
	"bytes"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

// MdToHtml 将 Markdown 转换为 HTML
func MdToHtml(md string) string {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		return ""
	}
	return buf.String()
}

// MdToText 将 Markdown 转换为纯文本 (过滤掉所有 HTML 标签)
func MdToText(md string) string {
	// 1. 先转成 HTML
	html := MdToHtml(md)

	// 2. 使用 bluemonday 过滤掉所有 HTML 标签，只剩下文本内容
	// StrictPolicy 会剔除所有标签，相当于你的正则过滤
	p := bluemonday.StrictPolicy()
	text := p.Sanitize(html)

	return strings.TrimSpace(text)
}

// MdGetSummary 提取 Markdown 摘要
// length 为可选参数，如果不传则默认 100
func MdGetSummary(md string, length ...int) string {
	// 1. 设置默认值
	finalLength := 100
	if len(length) > 0 {
		finalLength = length[0]
	}

	// 2. 获取纯文本并转为 rune 保证中文安全
	text := MdToText(md)
	runeText := []rune(text)

	if len(runeText) <= finalLength {
		return string(runeText)
	}

	// 3. 截取
	return string(runeText[:finalLength]) + "..."
}
