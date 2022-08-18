package uid

type Person struct {
	Uid   string `json:"UID" binding:"required"`
	Name  string `json:"Name" binding:"required"`
	Email string `json:"Email" binding:"required"`
}

type Controller interface {
	ShowAll() []Person
	Add(person Person)
}
type C2 interface {
	GetByid(id string) string
}

type People struct {
	Person []Person
}

func New() *People {
	return &People{}
}

func (p *People) Add(person Person) {
	p.Person = append(p.Person, person)
	// println(binary.Size(p))
}

func (p *People) ShowAll() []Person {
	return p.Person
}

func (p *People) GetByid(id string) string {
	for _, b := range p.Person {
		if b.Uid == id {
			print(b.Email)
			return b.Email
		}
	}
	return "bruh"
}
