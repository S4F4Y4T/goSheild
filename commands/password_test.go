package commands

import (
	"testing"
)

func BenchmarkGenerateRandomPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := generateRandomPassword(16)
		if err != nil {
			b.Fatal(err)
		}
	}
}
