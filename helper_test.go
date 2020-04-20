package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"testing"

	"gopkg.in/square/go-jose.v2"
)

func TestIssueJWT(t *testing.T) {
	rsaPrivateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	testKey := &jose.JSONWebKey{
		Key:   rsaPrivateKey,
		KeyID: "rsa-test-key",
	}
	rsaPublicKey := &rsaPrivateKey.PublicKey
	h := New("line")
	jwt, err := h.issueJWT(testKey)
	if err != nil {
		t.Error("error on generate JWT", err)
	}
	parsed, err := jose.ParseSigned(jwt)
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed.Signatures) != 1 {
		t.Fatal("Too many or too few signatures.")
	}
	payload, err := parsed.Verify(rsaPublicKey)
	if err != nil {
		t.Fatalf("signature did not validate: %v", err)
	}

	got := make(map[string]interface{})
	err = json.Unmarshal(payload, &got)
	if err != nil {
		t.Fatal(err)
	}
	if got["iss"].(string) != "line" {
		t.Error("invalid issuer")
	}
	if got["sub"].(string) != "line" {
		t.Error("invalid subject")
	}
	if got["aud"].(string) != AudienceLINE {
		t.Error("invalid audience")
	}
	f, ok := got["token_exp"].(float64)
	if !ok {
		t.Error("invalid token expire")
	}
	tokenExp := int(f)
	if tokenExp != h.TokenExpire {
		t.Errorf("expect %d, got %d", h.TokenExpire, tokenExp)
	}
}
