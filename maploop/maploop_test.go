package maploop

import (
	"fmt"
	"testing"
)

func TestMapLoop(t *testing.T) {
	studentGrades := make(map[string]int)

	grades := []struct {
		name  string
		grade int
	}{
		{"John", 90},
		{"Jane", 89},
		{"Michael", 24},
	}

	for _, data := range grades {
		studentGrades[data.name] = data.grade
	}

	for name, grade := range studentGrades {
		fmt.Printf("Nama: %s, Nilai: %d\n", name, grade)
	}
}
