package views

import (
	"testing"

	"github.com/hay-i/chronologger/test"
)

func TestGetFlashes(t *testing.T) {
	t.Run("should return a map of flashes", func(t *testing.T) {
		ctx := test.SetupEchoContext()

		flashes := GetFlashes(ctx)

		for _, flashType := range []FlashType{FlashError, FlashWarning, FlashInfo, FlashSuccess} {
			if len(flashes[flashType]) != 1 {
				t.Errorf("got %d flashes, want 0", len(flashes[flashType]))
			}
		}
	})
}
