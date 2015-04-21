package vcsclient

import (
	"net/http"
	"reflect"
	"testing"

	"sourcegraph.com/sourcegraph/go-vcs/vcs"
)

func TestRepository_Diff(t *testing.T) {
	setup()
	defer teardown()

	repoID := "a.b/c"
	repo_, _ := vcsclient.Repository(repoID)
	repo := repo_.(*repository)

	want := &vcs.Diff{Raw: "diff"}

	var called bool
	mux.HandleFunc(urlPath(t, RouteRepoDiff, repo, map[string]string{"RepoID": repoID, "Base": "b", "Head": "h"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	diff, err := repo.Diff("b", "h", nil)
	if err != nil {
		t.Errorf("Repository.Diff returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(diff, want) {
		t.Errorf("Repository.Diff returned %+v, want %+v", diff, want)
	}
}

func TestRepository_CrossRepoDiff(t *testing.T) {
	setup()
	defer teardown()

	repoID := "a.b/c"
	repo_, _ := vcsclient.Repository(repoID)
	repo := repo_.(*repository)

	want := &vcs.Diff{Raw: "diff"}

	var called bool
	mux.HandleFunc(urlPath(t, RouteRepoCrossRepoDiff, repo, map[string]string{"RepoID": repoID, "Base": "b", "HeadRepoID": "x.com/y", "Head": "h"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	headRepoID := "x.com/y"
	headRepo, _ := vcsclient.Repository(headRepoID)

	diff, err := repo.CrossRepoDiff("b", headRepo, "h", nil)
	if err != nil {
		t.Errorf("Repository.CrossRepoDiff returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(diff, want) {
		t.Errorf("Repository.CrossRepoDiff returned %+v, want %+v", diff, want)
	}
}
