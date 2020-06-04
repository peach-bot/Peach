package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpcodes(t *testing.T) {
	a := assert.New(t)
	// Test int generation
	a.Equal(opcodeDispatch, opcode(0))
	a.Equal(opcodeVoiceStateUpdate, opcode(4))
	a.Equal(opcodeResume, opcode(6))
	a.Equal(opcodeHeartbeatACK, opcode(11))
	// Test string method
	a.Equal(opcodeDispatch.String(), "opcodeDispatch")
}

func TestCloseCodes(t *testing.T) {
	a := assert.New(t)
	// Test int generation
	a.Equal(closecodeUnknownError, closecode(4000))
	a.Equal(closecodeAlreadyAuthenticated, closecode(4005))
	a.Equal(closecodeInvalidSquence, closecode(4007))
	a.Equal(closecodeDisallowedIntents, closecode(4014))
	// Test string method
	a.Equal(closecodeUnknownError.String(), "closecodeUnknownError")
}
