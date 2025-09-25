package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
)

type jsonFileFormat struct {
	NextID   int        `json:"next_id"`
	Contacts []*Contact `json:"contacts"`
}

type JSONFileStore struct {
	path     string
	contacts map[int]*Contact
	nextID   int
}

func NewJSONFileStore(path string) (*JSONFileStore, error) {
	s := &JSONFileStore{
		path:     path,
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
	if err := s.load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return s, nil
}

func (s *JSONFileStore) Add(c *Contact) error {
	c.ID = s.nextID
	s.contacts[c.ID] = c
	s.nextID++
	return s.save()
}

func (s *JSONFileStore) GetAll() ([]*Contact, error) {
	out := make([]*Contact, 0, len(s.contacts))
	for _, c := range s.contacts {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func (s *JSONFileStore) GetByID(id int) (*Contact, error) {
	c, ok := s.contacts[id]
	if !ok {
		return nil, errors.New("contact introuvable")
	}
	return c, nil
}

func (s *JSONFileStore) Update(id int, name, email string) error {
	c, ok := s.contacts[id]
	if !ok {
		return errors.New("contact introuvable")
	}
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}
	return s.save()
}

func (s *JSONFileStore) Delete(id int) error {
	if _, ok := s.contacts[id]; !ok {
		return errors.New("contact introuvable")
	}
	delete(s.contacts, id)
	return s.save()
}

func (s *JSONFileStore) load() error {
	b, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}
	var payload jsonFileFormat
	if err := json.Unmarshal(b, &payload); err != nil {
		return err
	}
	s.contacts = make(map[int]*Contact, len(payload.Contacts))
	for _, c := range payload.Contacts {
		s.contacts[c.ID] = c
	}
	if payload.NextID > 0 {
		s.nextID = payload.NextID
	} else {
		maxID := 0
		for id := range s.contacts {
			if id > maxID {
				maxID = id
			}
		}
		s.nextID = maxID + 1
	}
	return nil
}

func (s *JSONFileStore) save() error {
	list := make([]*Contact, 0, len(s.contacts))
	for _, c := range s.contacts {
		list = append(list, c)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })

	payload := jsonFileFormat{
		NextID:   s.nextID,
		Contacts: list,
	}
	enc, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, enc, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}
