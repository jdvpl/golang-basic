package models

type Man struct {
	Name      string
	Height    float64
	Weight    float64
	Breathing bool
	Think     bool
	Eating    bool
	Working   bool
	Hello     string
}

func (m *Man) Breathe()               { m.Breathing = true }
func (m *Man) Thinking()              { m.Think = true }
func (m *Man) Eat()                   { m.Working = true }
func (m *Man) Sex() string            { return "Man" }
func (m *Man) AddingName(name string) { m.Name = name }
func (m *Man) SayingHello() string    { return "Hello I am " + m.Name + "I am a " + m.Sex() }
func (m *Man) BeingAlive() bool       { return true }
