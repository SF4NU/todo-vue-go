package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) { //la funzione in pratica si occupa di criptare la password ricevuta in fase di registrazione, utilizza il pacchetto bcrypt per la crittografia e
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //bcrypt.DefaultCost - il costo per la generazione del hash è quello base di bcrypt quindi in teoria se aumentasse ci metterebbe di più a generare la password criptata anche se con un costo più alto c'è più protezione, per semplicità in questo caso lascio predefinito
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error { //la funzione si occupa di verificare se la password inserita (non criptata) combacia con quella presente nel database
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
