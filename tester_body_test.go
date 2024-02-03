package httpcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTester_WithString(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()

	checker.Test(mockT, "GET", "/mirrorbody").
		WithString("Test123").
		Check().
		HasString("Test123")

	assert.False(t, mockT.Failed())
}

func TestTester_WithBody(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()

	checker.Test(mockT, "GET", "/mirrorbody").
		WithBody([]byte("Test123")).
		Check().
		HasBody([]byte("Test123"))

	assert.False(t, mockT.Failed())
}

func TestTester_HasBody(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()
	checker.Test(mockT, "GET", "/byte").
		Check().
		HasBody([]byte("hello world"))
	assert.False(t, mockT.Failed())
}

func TestTester_MustHasBody(t *testing.T) {
	t.Skip("skip this test because it expects failure.")
	checker := newTestChecker()
	checker.Test(t, "GET", "/byte").
		Check().
		MustHasBody([]byte("hello")).
		Cb(func(response *http.Response) {
			t.Fatal("it is expected that this assertion will not be executed.")
		})
}

func TestTester_HasString(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()
	checker.Test(mockT, "GET", "/byte").
		Check().
		HasString("hello world")
	assert.False(t, mockT.Failed())
}

func TestTester_MustHasString(t *testing.T) {
	t.Skip("skip this test because it expects failure.")
	checker := newTestChecker()
	checker.Test(t, "GET", "/byte").
		Check().
		MustHasString("hello").
		Cb(func(response *http.Response) {
			t.Fatal("it is expected that this assertion will not be executed.")
		})
}

func TestTester_ContainsBody(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()
	checker.Test(mockT, "GET", "/byte").
		Check().
		ContainsBody([]byte("llo")).
		ContainsBody([]byte("hell"))
	assert.False(t, mockT.Failed())
}

func TestTester_MustContainsBody(t *testing.T) {
	t.Skip("skip this test because it expects failure.")
	checker := newTestChecker()
	checker.Test(t, "GET", "/byte").
		Check().
		MustContainsBody([]byte("aloha")).
		Cb(func(response *http.Response) {
			t.Fatal("it is expected that this assertion will not be executed.")
		})
}

func TestTester_NotContainsBody(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()
	checker.Test(mockT, "GET", "/byte").
		Check().
		NotContainsBody([]byte("aloha"))
	assert.False(t, mockT.Failed())
}

func TestTester_MustNotContainsBody(t *testing.T) {
	t.Skip("skip this test because it expects failure.")
	checker := newTestChecker()
	checker.Test(t, "GET", "/byte").
		Check().
		MustNotContainsBody([]byte("hello")).
		Cb(func(response *http.Response) {
			t.Fatal("it is expected that this assertion will not be executed.")
		})
}

func TestTester_ContainsString(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()
	checker.Test(mockT, "GET", "/byte").
		Check().
		ContainsString("llo").
		ContainsString("hell")
	assert.False(t, mockT.Failed())
}

func TestTester_MustContainsString(t *testing.T) {
	t.Skip("skip this test because it expects failure.")
	checker := newTestChecker()
	checker.Test(t, "GET", "/byte").
		Check().
		MustContainsString("aloha").
		Cb(func(response *http.Response) {
			t.Fatal("it is expected that this assertion will not be executed.")
		})
}

func TestTester_NotContainsString(t *testing.T) {
	mockT := new(testing.T)
	checker := newTestChecker()
	checker.Test(mockT, "GET", "/byte").
		Check().
		NotContainsString("aloha")
	assert.False(t, mockT.Failed())
}

func TestTester_MustNotContainsString(t *testing.T) {
	t.Skip("skip this test because it expects failure.")
	checker := newTestChecker()
	checker.Test(t, "GET", "/byte").
		Check().
		MustNotContainsString("hello").
		Cb(func(response *http.Response) {
			t.Fatal("it is expected that this assertion will not be executed.")
		})
}

func TestTester_MatchesJSONQuery(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Name": "Some", "Age": 30}`))
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	checker := newTestChecker()
	checker.Test(t, "GET", "/json").
		Check().
		MatchesJSONQuery(`.Name`)
}

func TestTester_NotMatchesJSONQuery(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Name": "Some", "Age": 30}`))
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	checker := newTestChecker()
	checker.Test(t, "GET", "/json").
		Check().
		NotMatchesJSONQuery(`.XXX`)
}
