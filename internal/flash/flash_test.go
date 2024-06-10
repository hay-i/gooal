package flash

import (
	"testing"

	"github.com/hay-i/gooal/pkg/test"
)

func TestAddFlash(t *testing.T) {
	t.Run("should add a flash to the context", func(t *testing.T) {
		ctx := test.SetupEchoContext()

		Add(ctx, "error message", Error)

		flashes := Get(ctx)[Error]

		if len(flashes) != 1 {
			t.Errorf("got %d flashes, want 1", len(flashes))
		}
	})
}

func TestGetFlashes(t *testing.T) {
	t.Run("should return a map of flashes", func(t *testing.T) {
		ctx := test.SetupEchoContext()

		Add(ctx, "error message", Error)
		Add(ctx, "warning message", Warning)
		Add(ctx, "info message", Info)
		Add(ctx, "success message", Success)

		flashes := Get(ctx)

		for _, flashType := range []Type{Error, Warning, Info, Success} {
			if len(flashes[flashType]) != 1 {
				t.Errorf("got %d flashes, want 0", len(flashes[flashType]))
			}
		}
	})
}
