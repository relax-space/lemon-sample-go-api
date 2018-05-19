package kit

type Job struct {
	OnTime  int `json:"onTime"`
	OffTime int `json:"offTime"`
}

func New() (job *Job) {
	job = &Job{
		OnTime:  8,
		OffTime: 17,
	}
	return
}

type Employee struct {
	Company  string `json:"company"`
	BossMood string `json:"bossMood"`
}

func (Employee) ComputerWorkTime(options ...func(*Job)) (workTime int) {
	job := New()
	for _, v := range options {
		if options == nil {
			continue
		}
		v(job)
	}
	workTime = job.OffTime - job.OnTime - 1
	return
}
