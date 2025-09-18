package sum

import(
	"testing"
)

func TestSum(t *testing.T) {
	if sum_num := sum(2, 3); sum_num==5 {
		t.Errorf("sum_num = 3; %d", sum_num)
		t.Fatalf("sum_num = 3; %d", sum_num)
		t.Skipf("sum_num = 3; %d", sum_num)
		t.Logf("sum_num = 3; %d", sum_num)
	}
}