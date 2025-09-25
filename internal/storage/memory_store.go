package storage

import "errors"

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(c *Contact) error {
	c.ID = ms.nextID
	ms.contacts[c.ID] = c
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	out := make([]*Contact, 0, len(ms.contacts))
	for _, c := range ms.contacts {
		out = append(out, c)
	}
	return out, nil
}

func (ms *MemoryStore) GetByID(id int) (*Contact, error) {
	c, ok := ms.contacts[id]
	if !ok {
		return nil, errors.New("contact introuvable")
	}
	return c, nil
}

func (ms *MemoryStore) Update(id int, name, email string) error {
	c, err := ms.GetByID(id)
	if err != nil {
		return err
	}
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}
	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return errors.New("contact introuvable")
	}
	delete(ms.contacts, id)
	return nil
}
