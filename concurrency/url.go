package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(fn WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultCh := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultCh <- result{u, fn(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultCh
		results[r.string] = r.bool
	}
	return results
}

func MockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}
