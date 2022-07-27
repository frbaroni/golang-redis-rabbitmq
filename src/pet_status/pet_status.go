package pet_status

type PetStatus int

const (
	Received PetStatus = iota
	Brushed
	Washed
	Dried
	Complete
)

func (status PetStatus) String() string {
	return []string{"Received", "Brushed", "Washed", "Dried", "Complete"}[status]
}
