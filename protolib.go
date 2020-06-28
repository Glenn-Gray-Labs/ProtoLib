package ProtoLib

import (
	"github.com/Glenn-Gray-Labs/imgui"
)

func Button(text string) bool {
	return imgui.Button(text)
}
