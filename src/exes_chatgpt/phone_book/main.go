package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Contact struct {
	Name   string
	Number string
}

type PhoneBook struct {
	Contacts []Contact
}

func checkNumber(number string) error {
	if len(number) == 0 {
		return errors.New("number is empty")
	}
	if len(number) != 11 {
		return errors.New("the number must consist of 11 digits")
	}
	if _, err := strconv.Atoi(number); err != nil {
		return errors.New("number must contain only digits")
	}
	return nil
}

func (p *PhoneBook) AddContact(name string, number string) error {
	if err := checkNumber(number); err != nil {
		return err
	}

	for _, c := range p.Contacts {
		if c.Name == name {
			return fmt.Errorf("контакт с именем %s уже существует", name)
		}
	}

	ct := Contact{Name: name, Number: number}
	p.Contacts = append(p.Contacts, ct)
	return nil
}

func (p *PhoneBook) FindByName(name string) (Contact, error) {
	for _, ct := range p.Contacts {
		if ct.Name == name {
			return ct, nil
		}
	}
	return Contact{}, fmt.Errorf("contact: %s not found", name)
}

func (p *PhoneBook) RemoveByName(name string) error {
	for i, ct := range p.Contacts {
		if ct.Name == name {
			p.Contacts = append(p.Contacts[:i], p.Contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact %q not found", name)
}

func main() {
	pb := PhoneBook{}
	if err := pb.AddContact("Петр", "79127535534"); err != nil {
		fmt.Println(err)
	}

	if contactPetr, err := pb.FindByName("Петр"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Наден контакт: %s: %s\n", contactPetr.Name, contactPetr.Number)
	}

	if err := pb.RemoveByName("Петр"); err != nil {
		fmt.Println(err)
	}
}
