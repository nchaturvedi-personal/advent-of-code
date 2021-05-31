package main

import (
	"log"
	"testing"
)

func Test_passwordCheck(t *testing.T) {
	//Given
	passwordPolicy := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

	//When
	t.Run("when given the old password policy", func(t *testing.T) {
		validPasswords, err := checkPasswords(passwordPolicy)
		if err != nil {
			log.Fatal(err)
		}

		if validPasswords != 2 {
			t.Fatal("Password Check Failed!")
		}
	})
	//When
	t.Run("when given the new password policy", func(t *testing.T) {
		validPasswordsNewPolicy, err := checkPasswordsNewPolicy(passwordPolicy)
		if err != nil {
			log.Fatal(err)
		}
		if validPasswordsNewPolicy != 1 {
			t.Fatal("Password Check with new policy failed!")
		}
	})

}
