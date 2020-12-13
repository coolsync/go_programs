package mathUnit

import "testing"

func TestGetSum(t *testing.T) {
	if GetSum(10) != 55 {
		// t.Errorf("err info, 期望：%d, 实际：%d\n", 55, GetSum(10))
		t.Logf("err info, 期望：%d, 实际：%d\n", 55, GetSum(10))
		// FailNow marks the function as having failed and stops its execution
		// by calling runtime.Goexit
		t.FailNow()
	}
	t.Log("GetSum testing ok")
}

func TestGetSumRecursive(t *testing.T) {
	sum := GetSumRecursive(10)

	if sum != 55 {
		t.Errorf("err info, 期望：%d, 实际：%d\n", 55, GetSum(10))
	}
	t.Log("GetSumRecursive testing ok")
}