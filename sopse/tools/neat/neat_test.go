package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("text")
	assert.Equal(t, "mC2ePrmW9VnmM_TRlN7zdh2Qn1o7ZH0ahR_q1nwyydE", hash)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123!\n")
	assert.Equal(t, "name-123", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1000, 0).Local()

	// success
	tobj := Time(1000)
	assert.Equal(t, want, tobj)
}
