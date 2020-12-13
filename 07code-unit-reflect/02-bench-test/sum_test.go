package benchUnit

import "testing"

func BenchmarkGetSum(b *testing.B) {
	b.Log("BenchmarkGetSum start:")

	b.ReportAllocs()	// 报告 内存开销

	for i := 0; i < b.N; i++ {
		GetSum(10)
	}

}

func BenchmarkGetSumRecursive(b *testing.B) {
	b.Log("BenchmarkGetSumRecursive start:")

	b.ReportAllocs()	// 报告 内存开销

	for i := 0; i < b.N; i++ {
		GetSumRecursive(10)
	}
}

func BenchmarkGetSum2(b *testing.B) {
	b.Log("BenchmarkGetSum start:")

	b.ReportAllocs()	// 报告 内存开销

	for i := 0; i < b.N; i++ {
		GetSum(100)
	}

}

func BenchmarkGetSumRecursive2(b *testing.B) {
	b.Log("BenchmarkGetSumRecursive start:")

	b.ReportAllocs()	// 报告 内存开销

	for i := 0; i < b.N; i++ {
		GetSumRecursive(100)
	}
}