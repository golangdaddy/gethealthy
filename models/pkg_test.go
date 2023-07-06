package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	assert := assert.New(t)

	user := DemoUser()

	thread := user.NewThread("DOOM BREAK: WATCH THIS KITTEN VIDEO")

	reply := thread.Reply(user, "hey watch this video")

	assert.NotNil(reply)
	assert.Len(thread.ID, 36)
	assert.Len(reply.ID, 73)
}
