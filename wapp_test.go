package wapp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	h := &Wapp{}

	serv := httptest.NewServer(h)
	defer serv.Close()

	t.Run("foo n point test", func(t *testing.T) {

		const expected = "borkab a legmenobb"
		requestURL := serv.URL + "/foo"
		requestBody := strings.NewReader(expected)
		req, err := http.NewRequest(http.MethodPost, requestURL, requestBody)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		fmt.Println(resp) //response objektum
		if err != nil {
			t.Fatal(err)
		}

		if http.StatusCreated != resp.StatusCode {
			t.Fatal("created code doesn't as expected. it should be: 201")
		} //ez az elvarasunk a szervertol

		req, err = http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err = http.DefaultClient.Do(req) //assimilar: httptest.Server.Client()
		if err != nil {
			t.Fatal(err)
		}
		if http.StatusOK != resp.StatusCode {
			t.Fatal("created code doesn't as expected. it should be: OK")
		} //ez az elvarasunk a szervertol
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if strings.TrimSpace(string(bs)) != expected {
			t.Fatal("we are sad")
		}

	})
}
