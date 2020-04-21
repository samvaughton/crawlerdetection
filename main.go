package main

import (
	"regexp"
)

func IsCrawler(userAgent string) bool {
	_, matched := GetCrawler(userAgent)

	return matched
}

func GetCrawler(userAgent string) (*CrawlerPattern, bool) {
	patterns := getCrawlerPatterns()

	for _, pattern := range patterns {
		if isMatch(pattern, userAgent) {
			return &pattern, true
		}
	}

	return nil, false
}

func isMatch(pattern CrawlerPattern, userAgent string) bool {
	matched, err := regexp.MatchString(pattern.Pattern, userAgent)

	if err != nil {
		panic(err)
	}

	return matched
}
