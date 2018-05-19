package kit

import (
	"fmt"
	"testing"
)

func Test_compute(t *testing.T) {
	value := compute(1, 2, add)
	fmt.Printf("add:%v\n", value)
	value = compute(8, 9, mul)
	fmt.Printf("mul:%v\n", value)
}

func Test_ComputerWorkTime_Eland(t *testing.T) {
	employee := Employee{
		Company:  "eland",
		BossMood: "good",
	}
	workTime := employee.ComputerWorkTime(func(job *Job) {
		if employee.Company == "eland" && employee.BossMood == "good" {
			job.OffTime = 13
		}
	})
	fmt.Println(workTime)
}
func Test_ComputerWorkTime_xiaomi(t *testing.T) {
	employee := Employee{
		Company:  "xiaomi",
		BossMood: "good",
	}
	workTime := employee.ComputerWorkTime(func(job *Job) {
		if employee.Company == "eland" && employee.BossMood == "good" {
			job.OffTime = 13
		}
	})
	fmt.Println(workTime)
}
