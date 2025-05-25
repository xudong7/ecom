package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	if hashed == password {
		t.Fatal("hashed password should not be equal to plain password")
	}
}

func TestComparePasswords(t *testing.T) {
	password := "password"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	if !ComparePasswords(hashed, []byte(password)) {
		t.Fatal("expected passwords to match")
	}

	if ComparePasswords(hashed, []byte("wrongpassword")) {
		t.Fatal("expected passwords to not match")
	}
}
