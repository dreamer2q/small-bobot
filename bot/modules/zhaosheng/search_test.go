package zhaosheng

import "testing"

func TestGetQueryResult(t *testing.T) {
	res, err := GetQueryResult("123", "321")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("result: %v", res)
}
