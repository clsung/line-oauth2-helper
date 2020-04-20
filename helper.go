package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/square/go-jose.v2"
)

// default values
const (
	AudienceLINE       = "https://api.line.me/"
	DefaultExpiry      = 30 * time.Minute
	DefaultTokenExpire = 2592000 // 30 Day
)

// Helper helps to generate JWT for line
type Helper struct {
	lineChannelID string
	//Expiry is the JWT expiry duration from now.
	Expiry time.Duration
	// TokenExp denotes the valid expiration time for the channel access token in seconds.
	TokenExpire int
}

// New helper
func New(channelID string) *Helper {
	return &Helper{
		lineChannelID: channelID,
		Expiry:        DefaultExpiry,
		TokenExpire:   DefaultTokenExpire,
	}
}

// WithExpiry set JWT valid duration
func (h *Helper) WithExpiry(expiry time.Duration) *Helper {
	h.Expiry = expiry
	return h
}

// WithTokenExpire set token expire in seconds
func (h *Helper) WithTokenExpire(tokenExp int) *Helper {
	h.TokenExpire = tokenExp
	return h
}

// GetLineJWTFromFile reads from file and return JWT or error
func (h *Helper) GetLineJWTFromFile(filePath string) (string, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("read file %s error: %w", filePath, err)
	}
	return h.GetLineJWT(r)
}

// GetLineJWT reads from io.Reader and return JWT or error
func (h *Helper) GetLineJWT(r io.Reader) (string, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}
	key, err := h.convertLineIssuedPrivateKey(buf)
	if err != nil {
		return "", fmt.Errorf("convertLineIssuedPrivateKey: %w", err)
	}
	return h.issueJWT(key)
}

func (h *Helper) convertLineIssuedPrivateKey(buf []byte) (*jose.JSONWebKey, error) {
	var err error
	m := make(map[string]map[string]string)
	if err = json.Unmarshal(buf, &m); err != nil {
		return nil, fmt.Errorf("json format: %w", err)
	}
	if _, ok := m["privateKey"]; ok {
		if buf, err = json.Marshal(m["privateKey"]); err != nil {
			return nil, err
		}
	}
	key := &jose.JSONWebKey{}
	if err = key.UnmarshalJSON(buf); err != nil {
		return nil, fmt.Errorf("invalid key json format: %w", err)
	}
	return key, nil
}

func (h *Helper) issueJWT(key *jose.JSONWebKey) (string, error) {
	opts := &jose.SignerOptions{}
	opts.WithType("JWT")
	payload := make(map[string]interface{})
	payload["iss"] = h.lineChannelID
	payload["sub"] = h.lineChannelID
	payload["aud"] = AudienceLINE
	payload["exp"] = time.Now().Add(h.Expiry).Unix()
	payload["token_exp"] = h.TokenExpire
	buf, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("payload: %w", err)
	}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, opts)
	if err != nil {
		return "", fmt.Errorf("NewSigner: %w", err)
	}
	obj, err := signer.Sign(buf)
	if err != nil {
		return "", fmt.Errorf("sign: %w", err)
	}
	msg, err := obj.CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("compact: %w", err)
	}
	return msg, nil
}
