# Go Crawler Detection

Tiny Go library to detect if a useragent is a bot, spider, crawler etc. Uses https://github.com/monperrus/crawler-user-agents for the data.

## Usage

`import github.com/samvaughton/go-crawler-detection`

### API

 - `crawlerdetection.GetCrawler(string) CrawlerPattern`
 - `crawlerdetection.IsCrawler(string) bool`
 
## Examples

`crawlerdetection.IsCrawler("Googlebot/2.1")` ***` = true`***


