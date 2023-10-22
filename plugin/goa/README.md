httpcheck Goa plugin
===

This plugin provides a mounter for the [Goa](http://github.com/goadesign/goa) HTTP endpoints, enabling seamless integration with [httpcheck](https://github.com/ikawaha/httpcheck). The mounter satisfies the http.Handler interface, allowing it to be directly set and utilized within httpcheck.

## Usage Example

The [following example](https://github.com/ikawaha/httpcheck/blob/5e434b64049b13f5558e83d26abdeaa28b281cd7/plugin/goa/_test/calc_test.go#L18-L35) demonstrates how to use the goa plugin to test a Goa HTTP endpoint. In this example, we're testing an endpoint that multiplies two numbers.

```go
package calc

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/ikawaha/httpcheck"
	"github.com/ikawaha/httpcheck/plugin/goa"
	gen "github.com/ikawaha/httpcheck/plugin/goa/calc/gen/calc"
	"github.com/ikawaha/httpcheck/plugin/goa/calc/gen/http/calc/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMounter(t *testing.T) {
	m := goa.NewMounter()
	m.Mount(goa.EndpointModular{
		Builder:  server.NewMultiplyHandler,
		Mounter:  server.MountMultiplyHandler,
		Endpoint: gen.NewMultiplyEndpoint(NewCalc()),
	})
	httpcheck.New(m).Test(t, "GET", "/multiply/3/5").
		Check().
		MustHasStatus(http.StatusOK).
		Cb(func(r *http.Response) {
			b, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			i, err := strconv.Atoi(strings.TrimSpace(string(b)))
			assert.NoError(t, err, string(b))
			assert.Equal(t, 3*5, i)
		})
}
```