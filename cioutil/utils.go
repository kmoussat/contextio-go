package cioutil

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"net/url"
	"strconv"
	"time"
)

// Cio struct contains the api key and secret, and other configuration settings,
// along with an optional logger, and provides access to methods used by all
// ContextIO structs/object.
type Cio struct {
	apiKey         string
	apiSecret      string
	Log            Logger
	Host           string
	RequestTimeout time.Duration

	// PreRequestHook is a function that will be executed before the request is made,
	// its arguments are the User ID, Account Label, the Method (GET/POST/etc),
	// the URL, and the redacted body values
	PreRequestHook func(string, string, string, string, url.Values)

	// PostRequestShouldRetryHook is a function that will be executed after each request is made,
	// and will be called at least once.
	// Its arguments are the request Attempt # (starts at 1), User ID, Account Label,
	// the Method (GET/POST/etc), the URL, response Status Code, response Payload,
	// and any error received while attempting this request.
	// The returned boolean is whether this request should be
	// retried or not, which if false then this is the last call of this function,
	// but if true means this function will be called again.
	PostRequestShouldRetryHook func(int, string, string, string, string, int, string, error) bool

	// ResponseBodyCloseErrorHook is a function that will execute if there is an
	// error closing the response body
	ResponseBodyCloseErrorHook func(error)
}

// NewCio returns a CIO struct for embedding in a concrete type.
func NewCio(key string, secret string, logger Logger, host string, requestTimeout time.Duration) Cio {
	return Cio{
		apiKey:         key,
		apiSecret:      secret,
		Log:            logger,
		Host:           host,
		RequestTimeout: requestTimeout,
	}
}

// ValidateCallback returns true if this Webhook Callback or User Account Status Callback authenticates
func (cio Cio) ValidateCallback(token string, signature string, timestamp int) bool {
	// Hash timestamp and token with secret, compare to signature
	message := strconv.Itoa(timestamp) + token
	hash := hashHmac(sha256.New, message, cio.apiSecret)
	return len(hash) > 0 && signature == hash
}

// hashHmac returns the hash of a message hashed with the provided hash function, using the provided secret
func hashHmac(hashAlgorithm func() hash.Hash, message string, secret string) string {
	h := hmac.New(hashAlgorithm, []byte(secret))
	if _, err := h.Write([]byte(message)); err != nil {
		panic("hash.Hash unable to write message bytes, with error: " + err.Error())
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Logger interface which *log.Logger uses.
// Allows injection of user specified loggers, such as log.Logger or logrus.
type Logger interface {
	Printf(format string, v ...interface{})
}

// TestLogger is a *bytes.Buffer that implements the logging interface
type TestLogger struct {
	*bytes.Buffer
}

// Printf prints the arguments to the buffer, using fmt.Sprintf
func (l *TestLogger) Printf(format string, v ...interface{}) {
	_, err := l.Write([]byte(fmt.Sprintf(format, v...)))
	if err != nil {
		panic("Error writing to test logger: " + err.Error())
	}
}
