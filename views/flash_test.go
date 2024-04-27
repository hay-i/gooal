package views

import (
	"testing"

	"github.com/hay-i/chronologger/test"
)

func TestAddFlash(t *testing.T) {
	t.Run("should add a flash to the context", func(t *testing.T) {
		ctx := test.SetupEchoContext()

		AddFlash(ctx, "error message", FlashError)

		flashes := GetFlashes(ctx)[FlashError]

		if len(flashes) != 1 {
			t.Errorf("got %d flashes, want 1", len(flashes))
		}
	})
}

func TestGetFlashes(t *testing.T) {
	t.Run("should return a map of flashes", func(t *testing.T) {
		ctx := test.SetupEchoContext()

		AddFlash(ctx, "error message", FlashError)
		AddFlash(ctx, "warning message", FlashWarning)
		AddFlash(ctx, "info message", FlashInfo)
		AddFlash(ctx, "success message", FlashSuccess)

		flashes := GetFlashes(ctx)

		for _, flashType := range []FlashType{FlashError, FlashWarning, FlashInfo, FlashSuccess} {
			if len(flashes[flashType]) != 1 {
				t.Errorf("got %d flashes, want 0", len(flashes[flashType]))
			}
		}
	})
}
