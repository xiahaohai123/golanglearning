package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	//check := func(url string) {
	//	results[url] = wc(url)
	//}

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// 注意 resultChannel，它在循环中持续输出结果
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
