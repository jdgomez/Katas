package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChan := make(chan result)

	for _, url := range urls {
		go func() {
			resultsChan <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		res := <-resultsChan
		results[res.string] = res.bool
	}

	return results
}
