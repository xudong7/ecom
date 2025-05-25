package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Fatalf("failed to create JWT: %v", err)
	}

	if token == "" {
		t.Fatal("expected non-empty token")
	}
}
