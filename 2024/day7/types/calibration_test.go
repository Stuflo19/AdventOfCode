package types

import "testing"

func TestNewCalibration(t *testing.T) {
	calibration := NewCalibration(7290, []int{6, 8, 6, 15})

	calibration.FindTimesOrAddForExpected2()

	if calibration.Total != 7290 {
		t.Errorf("Expected total to be 7290, but got %v", calibration.Total)
	}
}
