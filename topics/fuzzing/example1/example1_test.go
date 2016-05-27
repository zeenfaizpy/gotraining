// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to validate the api endpoints.
package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/gotraining/topics/fuzzing/example1"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	api.Routes()
}

// TestProcess tests the Process endpoint with proper data.
func TestProcess(t *testing.T) {
	data := []struct {
		url    string
		status int
		val    []byte
		resp   string
	}{
		{"/process", http.StatusOK, []byte("Bill:46,Lisa:42,John:35,Eduardo:46"), `[{"Name":"Bill","Age":46},{"Name":"Lisa","Age":42},{"Name":"John","Age":35},{"Name":"Eduardo","Age":46}]`},
	}

	t.Log("Given the need to test the Process endpoint.")
	{
		for i, d := range data {
			t.Logf("\tTest %d:\tWhen checking %q for status code %d with data %s", i, d.url, d.status, d.val)
			{
				r, _ := http.NewRequest("POST", d.url, bytes.NewBuffer(d.val))
				w := httptest.NewRecorder()
				http.DefaultServeMux.ServeHTTP(w, r)

				if w.Code != d.status {
					t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d].", failed, d.status, w.Code)
				}
				t.Logf("\t%s\tShould receive a status code of %d for the response.", succeed, d.status)

				recv := w.Body.String()

				if d.resp != recv[:len(recv)-1] {
					t.Log("GOT:", recv)
					t.Log("EXP:", d.resp)
					t.Fatalf("\t%s\tShould get the expected result.", failed)
				}
				t.Logf("\t%s\tShould get the expected result.", succeed)
			}
		}
	}
}
