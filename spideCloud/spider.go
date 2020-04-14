package spideCloud

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/caojiehz/httpUtil"
	"strings"
)

func Spider() (cdns []CdnRegion) {
	data, err := httpUtil.Get(httpUtil.GetParasTuple{
		URL:      "https://www.feitsui.com/zh-hans/blog/page/11",
		RetryNum: 1,
	})

	if err != nil {
		fmt.Println("spide error:", err)
		return
	}

	body := strings.NewReader(string(data))
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("html error:", err)
		return
	}

	doc.Find("tr").Each(func(row int, s *goquery.Selection) {
		if row == 0 {
			return
		}

		var cdn CdnRegion
		var set bool
		s.Children().Each(func(col int, selection *goquery.Selection) {
			set = true
			if col == 0 {
				cdn.name = strings.TrimSpace(selection.Text())
			}
			if col == 1 {
				cdn.region = strings.TrimSpace(selection.Text())
			}
			if col == 2 {
				cdn.domain = selection.Text()
				vec := strings.Split(cdn.domain, "\n")
				if len(vec) == 4 {
					cdn.domain = strings.TrimSpace(vec[1])
					set = true
				}
			}
		})
		if set {
			cdns = append(cdns, cdn)
		}
	})
	return
}
