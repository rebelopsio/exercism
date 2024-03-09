// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Address map[string]string
	Name    string
	Age     int
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	nr := &Resident{}
	nr.Address = address
	nr.Name = name
	nr.Age = age
	return nr
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r == nil {
		return false
	}
	if r.Address == nil {
		return false
	}
	if r.Address["street"] == "" {
		return false
	}
	if r.Name == "" {
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Address = nil
	r.Name = ""
	r.Age = 0
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	c := 0
	for _, v := range residents {
		if v.HasRequiredInfo() {
			c += 1
		}
	}
	return c
}
