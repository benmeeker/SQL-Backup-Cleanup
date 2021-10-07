package main

func main() {
	var filepath []string
	filepath = append(filepath, `\\ATAT\SQL Backups\SQL Backups\EVEREST_SYSTEM`)
	filepath = append(filepath, `\\ATAT\SQL Backups\SQL Backups\EVEREST_TNS`)
	filepath = append(filepath, `\\ATAT\SQL Backups\SQL Backups\EVEREST_EENT`)
	filepath = append(filepath, `\\ATAT\SQL Backups\SQL Backups\EVEREST_WHITEFIR`)
	filepath = append(filepath, `\\ATAT\SQL Backups\SQL Backups\Namifiers`)
	filepath = append(filepath, `\\ATAT\SQL Backups\SQL Backups\StarShip`)

	for _, path := range filepath {
		delete(search(path))
	}
}
