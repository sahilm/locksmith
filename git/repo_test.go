package git_test

import (
	"io/ioutil"
	"testing"

	"time"

	"os"

	"github.com/kylelemons/godebug/pretty"
	"github.com/sahilm/locksmith/git"
)

func TestRepo(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)
	repo := git.Repo{Dir: tmpdir, Timeout: 1 * time.Second}
	err = repo.Init()
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"foo", "bar"} {
		err = ioutil.WriteFile(tmpdir+"/"+f, nil, 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
	err = repo.Add(".")
	if err != nil {
		t.Fatal(err)
	}
	err = repo.Commit("add foo & bar")
	if err != nil {
		t.Fatal(err)
	}
	files, err := repo.ListFiles()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"bar", "foo"}
	if diff := pretty.Compare(files, want); diff != "" {
		t.Errorf("%v", diff)
	}
}
