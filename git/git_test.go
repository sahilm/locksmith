package git_test

// Silence tests till I create a mechanism to create a temp testing locks dir

//func TestListFiles(t *testing.T) {
//	t.Run("it returns the list of files in the repo", func(t *testing.T) {
//		got, err := git.ListFiles("testdata/repo/.git", 100*time.Millisecond)
//		if err != nil {
//			t.Fatal(err)
//		}
//		want := []string{"bar", "foo"}
//		if diff := pretty.Compare(got, want); diff != "" {
//			t.Errorf("%v", diff)
//		}
//	})
//	t.Run("it fails with a nice error message", func(t *testing.T) {
//		_, err := git.ListFiles("testdata/repo", 100*time.Millisecond)
//		if err == nil {
//			t.Fatal("got no error, expected error")
//		}
//		got := err.Error()
//		want := "failed to list files fatal: Not a git repository: 'testdata/repo': exit status 128"
//		if got != want {
//			t.Errorf("got: %v, want: %v", got, want)
//		}
//	})
//	t.Run("it times out", func(t *testing.T) {
//		_, err := git.ListFiles("testdata/repo", -1*time.Millisecond)
//		if err == nil {
//			t.Fatal("got no error, expected error")
//		}
//		got := err.Error()
//		want := "failed to list files : git --git-dir testdata/repo ls-files timed out after -1ms"
//		if got != want {
//			t.Errorf("got: %v, want: %v", got, want)
//		}
//	})
//}
