package main

import (
	"fmt"
	"regexp"
	"strings"
)

var tableRegex = regexp.MustCompile(
	"^*([^\\n ].*\\|.*)\\n" + // Header
		" {0,3}(?:\\| *)?(:?-+:? *(?:\\| *:?-+:? *)*)(?:\\| *)?" + // Align
		`((\|?.+\|?\n)|(\|?.+\|?)){1,}`, // Cells
)

func findTable(str string) string {
	var match = tableRegex.FindAllString(str, -1)
	for i, v := range match {
		fmt.Println(i, v)
	}
	return str
}
func findTableAndReplace(str string) string {
	var match = tableRegex.FindAllString(str, -1)
	for i, v := range match {
		fmt.Println(i, v)
		tableNormalized := normalizeTable(v)
		fmt.Println(tableNormalized)
		str = strings.Replace(str, v, parseMarkdownTableToHtml(tableNormalized), 1)
	}
	return str
}
func normalizeTable(str string) string {
	str = strings.Trim(str, "\n ")
	var lines = strings.Split(str, "\n")
	var newLines = []string{}
	for _, line := range lines {
		line = strings.Trim(line, " ")
		if line[0:1] == "|" {
			line = line[1:]
		}
		if line[len(line)-1:] == "|" {
			line = line[0 : len(line)-1]
		}
		newLines = append(newLines, line)
	}
	return strings.Join(newLines, "\n")
}

func parseMarkdownTableToHtml(markdownTable string) string {
	// Alignments
	var r1 = regexp.MustCompile(`^\:(\-|\w){1,}$`)
	var r2 = regexp.MustCompile(`^\:(\-|\w){1,}\:$`)
	var r3 = regexp.MustCompile(`^\:(\-|\w){1,}$`)
	lines := strings.Split(markdownTable, "\n")

	// Handle table align (| :---|:----:|---: |)
	alignments := []string{}
	items := strings.Split(lines[1], "|")
	for _, item := range items {
		if item == "" {
			continue
		}
		trimedItem := strings.Trim(item, " ")
		if r1.MatchString(trimedItem) {
			alignments = append(alignments, "left")
		} else if r2.MatchString(trimedItem) {
			alignments = append(alignments, "center")
		} else if r3.MatchString(trimedItem) {
			alignments = append(alignments, "right")
		} else {
			alignments = append(alignments, "left")
		}
	}
	// Header
	header := strings.Split(lines[0], "|")
	headerHtml := `<thead><tr>`
	for i := 0; i < len(header); i++ {
		headerHtml += `<th`
		if len(alignments) == 0 {
			headerHtml += ` style="text-align: ` + alignments[i] + `"`
		}
		headerHtml += `>` + header[i] + `</th>`
	}
	headerHtml += `</tr></thead>`

	// Body
	bodyHtml := `<tbody>`
	for i := 0; i < len(lines); i++ {
		if i < 2 {
			continue // Headers and alignments
		}
		row := strings.Split(lines[i], "|")
		bodyHtml += `<tr>`
		for j := 0; j < len(row); j++ {
			bodyHtml += `<td`
			if len(alignments) == 0 {
				bodyHtml += ` style="text-align: ` + alignments[j] + `"`
			}
			bodyHtml += `>` + row[j] + `</td>`
		}
		bodyHtml += `</tr>`
	}
	bodyHtml += `</tbody>`
	// Table
	return `<table>` + headerHtml + bodyHtml + `</table>`

}
