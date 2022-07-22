package stm

import (
	"fmt"
	"net/http"
	"time"
)

// PingSearchEngines requests some ping server from it calls Sitemap.PingSearchEngines.
func PingSearchEngines(opts *Options, urls ...string) {
	urls = append(urls, []string{
		"http://www.google.com/webmasters/tools/ping?sitemap=%s",
		"http://www.bing.com/webmaster/ping.aspx?siteMap=%s",
	}...)
	sitemapURL := opts.IndexLocation().URL()

	client := http.Client{Timeout: time.Duration(5 * time.Second)}

	for _, url := range urls {
		go func(baseurl string) {
			url := fmt.Sprintf(baseurl, sitemapURL)

			resp, err := client.Get(url)
			if err != nil {
				fmt.Printf("[E] Ping failed: %s (URL:%s)\n",
					err, url)
				return
			}
			defer resp.Body.Close()
		}(url)
	}

}
