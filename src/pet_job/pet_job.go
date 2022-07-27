package pet_job

type PetJob int

const (
	None PetJob = iota
	Received
	Brushed
	Washed
	Dried
	Complete
)

var PetJobs = []PetJob{Received, Brushed, Washed, Dried, Complete}

func (job PetJob) String() string {
	return []string{"None", "Received", "Brushed", "Washed", "Dried", "Complete"}[job]
}

func Parse(job string) PetJob {
	for _, j := range PetJobs {
		if job == j.String() {
			return j
		}
	}
	return None
}
