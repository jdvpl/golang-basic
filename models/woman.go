package models

type Woman struct {
	Man
}

func (w *Woman) Breathe()               { w.Breathing = true }
func (w *Woman) Thinking()              { w.Think = true }
func (w *Woman) Work()                  { w.Working = true }
func (m *Woman) Sex() string            { return "Woman" }
func (m *Woman) AddingName(name string) { m.Name = name }
func (m *Woman) BeingAlive() bool       { return true }
