package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	testCases := map[string]struct {
		path       string
		ua         string
		goldenFile string
	}{
		"job slug fb": {
			path:       "/job/blockchain-developer-44",
			ua:         "facebookexternalhit",
			goldenFile: "goldenFiles/job_slug_fb.html",
		},
		"job slug -fb": {
			path:       "/job/blockchain-developer-44",
			ua:         "",
			goldenFile: "goldenFiles/job_slug_-fb.html",
		},
		"job -slug fb": {
			path:       "/job/x",
			ua:         "facebookexternalhit",
			goldenFile: "goldenFiles/job_-slug_fb.html",
		},
		"job -slug -fb": {
			path:       "/job/x",
			ua:         "",
			goldenFile: "goldenFiles/job_-slug_-fb.html",
		},
		"profile slug fb": {
			path:       "/profile/johan-deecke-292",
			ua:         "facebookexternalhit",
			goldenFile: "goldenFiles/profile_slug_fb.html",
		},
		"profile slug -fb": {
			path:       "/profile/johan-deecke-292",
			ua:         "",
			goldenFile: "goldenFiles/profile_slug_-fb.html",
		},
		"profile -slug fb": {
			path:       "/profile/x",
			ua:         "facebookexternalhit",
			goldenFile: "goldenFiles/profile_-slug_fb.html",
		},
		"profile -slug -fb": {
			path:       "/profile/x",
			ua:         "",
			goldenFile: "goldenFiles/profile_-slug_-fb.html",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			router := setupRouter()
			w := httptest.NewRecorder()
			w.Header().Set("User-Agent", tc.ua)
			w.WriteHeader(200)
			req, _ := http.NewRequest("GET", tc.path, nil)
			router.ServeHTTP(w, req)

			got := w.Body.String()

			f, err := os.Open(tc.goldenFile)
			if err != nil {
				t.Fatalf("os.Open() err = %s; want nil", err)
			}
			want, err := ioutil.ReadAll(f)
			f.Close()

			assert.Equal(t, 200, w.Code)
			assert.Equal(t, want, got)
		})
	}
}
