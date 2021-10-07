package main

func main() {
	var filepath []string
	// TODO: Insert as many filepath's as necessary -- this program is for local backup folders or network drives.
	filepath = append(filepath, `filepath here`)

	for _, path := range filepath {
		delete(search(path))
	}
}
