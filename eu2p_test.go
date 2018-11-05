package eu2p

import (
	"testing"
)

func Test_GetEmojiPng(t *testing.T) {
	img := GetEmojiImg("1f487", 32, 32)
	if img != nil {
		t.Log("success")
	}
}
