package neutrino

type constants struct {
}

var Constants = &constants{}

func (c *constants) DatabaseName() string {
	return "neutrino"
}

func (c *constants) UsersCollection() string {
	return "users"
}

func (c *constants) ApplicationsCollection() string {
	return "applications"
}