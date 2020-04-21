package main

import (
	"fmt"
	"testing"
)

func TestMatchesWork(t *testing.T) {

	t.Run("test all patterns match their instances", func(t *testing.T) {
		patterns := getCrawlerPatterns()

		for _, pattern := range patterns {
			for _, instance := range pattern.Instances {

				if isMatch(pattern, instance) == false {
					fmt.Println("Pattern " + pattern.Pattern + " failed to match: " + instance)
					t.Fail()
				}

			}
		}
	})

	t.Run("test api usage of IsCrawler", func(t *testing.T) {

		if IsCrawler("Googlebot/") == false {
			t.Fail()
		}

		if IsCrawler("Googlebot/2.1") == false {
			t.Fail()
		}

	})

}
