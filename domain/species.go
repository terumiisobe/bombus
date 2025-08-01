package domain

type Species struct {
	id             int
	scientificName string
	commonName     string
}

func NewSpecies(id int, scientificName, commonName string) Species {
	return Species{id, scientificName, commonName}
}

func (s Species) GetId() int {
	return s.id
}

func (s Species) GetScientificName() string {
	return s.scientificName
}

func (s Species) GetCommonName() string {
	return s.commonName
}

func (s Species) SetId(id int) {
	s.id = id
}

func (s Species) SetScientificName(scientificName string) {
	s.scientificName = scientificName
}

func (s Species) SetCommonName(commonName string) {
	s.commonName = commonName
}
