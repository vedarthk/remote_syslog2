package loggly

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestCerts(t *testing.T) {
	h := md5.New()
	h.Write(certs())
	expected := "adfb41fd477c07cf086ab1cf8cf72025"
	actual := fmt.Sprintf("%x", h.Sum(nil))
	if actual != expected {
		t.Errorf("Mismatched hash for loggly certs, expected %s got %s", expected, actual)
	}
}

func TestRootCA(t *testing.T) {
	pool := RootCA()
	if len(pool.Subjects()) != 3 {
		t.Errorf("error loading RootCA")
	}
}
