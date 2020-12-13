package benchUnit

import (
	"testing"
)

var (
	filename = "../Files/db.json"
	p = Person{"sam", 30, 2000, 0, []string{"吃东西", "总结", "smokes"}}
)

func BenchmarkEncodePerson2JsonFile(b *testing.B) {
	b.Log("BenchmarkEncodePerson2JsonFile bench start")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EncodePerson2JsonFile(filename, &p)
	}

}

func BenchmarkDecodePerson2JsonFile(b *testing.B) {
	b.Log("BenchmarkDecodePerson2JsonFile bench start")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeJsonFile2Person(filename)
	}

}
