package vcsclient

import (
	"net/http"
	"testing"

	"sourcegraph.com/sourcegraph/go-vcs/vcs"
)

func TestRepository_MergeBase(t *testing.T) {
	setup()
	defer teardown()

	repoID := "a.b/c"
	repo_, _ := vcsclient.Repository(repoID)
	repo := repo_.(*repository)

	want := vcs.CommitID("abcd")

	var called bool
	mux.HandleFunc(urlPath(t, RouteRepoMergeBase, repo, map[string]string{"RepoID": repoID, "CommitIDA": "a", "CommitIDB": "b"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		http.Redirect(w, r, urlPath(t, RouteRepoCommit, repo, map[string]string{"CommitID": "abcd"}), http.StatusFound)
	})

	commitID, err := repo.MergeBase("a", "b")
	if err != nil {
		t.Errorf("Repository.MergeBase returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if commitID != want {
		t.Errorf("Repository.MergeBase returned %+v, want %+v", commitID, want)
	}
}

func TestRepository_CrossRepoMergeBase(t *testing.T) {
	setup()
	defer teardown()

	repoID := "a.b/c"
	repo_, _ := vcsclient.Repository(repoID)
	repo := repo_.(*repository)

	want := vcs.CommitID("abcd")

	var called bool
	mux.HandleFunc(urlPath(t, RouteRepoCrossRepoMergeBase, repo, map[string]string{"RepoID": repoID, "CommitIDA": "a", "BRepoID": "x.com/y", "CommitIDB": "b"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		http.Redirect(w, r, urlPath(t, RouteRepoCommit, repo, map[string]string{"CommitID": "abcd"}), http.StatusFound)
	})

	bRepoID := "x.com/y"
	bRepo, _ := vcsclient.Repository(bRepoID)

	commitID, err := repo.CrossRepoMergeBase("a", bRepo, "b")
	if err != nil {
		t.Errorf("Repository.CrossRepoMergeBase returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if commitID != want {
		t.Errorf("Repository.CrossRepoMergeBase returned %+v, want %+v", commitID, want)
	}
}
