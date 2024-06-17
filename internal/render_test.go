package internal

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/wijohnst/templ-htmx/internal/test_utils"
	"github.com/wijohnst/templ-htmx/views"
)

const sut = "Render"

type Component interface {
	Render(ctx context.Context, writer io.Writer) error
}

type MockComponent struct {
	RenderFunc func(ctx context.Context, writer io.Writer) error
}

func (m *MockComponent) Render(ctx context.Context, writer io.Writer) error {
	return m.RenderFunc(ctx, writer)
}

func TestRender(t *testing.T) {
	// Test 0 - No Error
	desc := "Returns correct markup"
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	template := views.Index()
	ctx.Request = &http.Request{}

	Render(ctx, template)

	expected := "<div>Hello, World </div>"
	actual := w.Body.String()

	test_utils.Assert(sut, desc, 0, t, expected, actual, nil)

	// Test 1 - Template throws
	desc = "Returns an error"
	mockTemplate := &MockComponent{
		RenderFunc: func(ctx context.Context, writer io.Writer) error {
			return errors.New("")
		},
	}

	err := Render(ctx, mockTemplate)

	test_utils.TestShouldThrow(sut, desc, 1, t, err)

}
