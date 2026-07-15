package db

func exists(u string) bool {
}

func add(u string, pH []byte, s bool, lH bool) bool {
	if exists(u) {
		return false
	}
	
}
