package runners

import "crypto/rand"

func randomString() string {
	data := make([]byte, 64)
	_, err := rand.Read(data)

	// Should nevere happen
	if err != nil {
		panic(err)
	}

	return string(data)
}
