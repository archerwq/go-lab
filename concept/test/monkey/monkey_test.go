/*
It's not possible through regular language constructs, but we can always bend computers to our will!
Monkey implements monkeypatching by rewriting the running executable at runtime and inserting a jump
to the function you want called instead. This is as unsafe as it sounds and I don't recommend anyone
do it outside of a testing environment.

Monkey sometimes fails to patch a function if inlining is enabled.
Try running your tests with inlining disabled, for example: go test -gcflags=-l

https://github.com/bouk/monkey
https://bou.ke/blog/monkey-patching-in-go/
https://www.jianshu.com/p/2f675d5e334e
*/
package monkey

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"bou.ke/monkey"
)

// go test -timeout 30s github.com/archerwq/go-lab/concept/test/monkey -run ^TestPatch$ -v -gcflags=-l
func TestPatchFunc(t *testing.T) {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}

func TestPatchInstanceMethod(t *testing.T) {
	getNotAllowedErr := fmt.Errorf("GET not allowed")
	// Note that patching the method for just one instance is currently not possible,
	// PatchInstanceMethod will patch it for all instances. Don't bother trying
	// monkey.Patch(instance.Method, replacement), it won't work.
	// monkey.UnpatchInstanceMethod(<type>, <name>) will undo PatchInstanceMethod.
	monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient),
		"Get", func(c *http.Client, url string) (*http.Response, error) {
			return nil, getNotAllowedErr
		})
	_, err := http.Get("http://google.com")
	if err != getNotAllowedErr {
		t.Errorf("got error: %v, want error: %v", err, getNotAllowedErr)
	}
	monkey.UnpatchAll()
}

func TestPatchGuard(t *testing.T) {
	var guard *monkey.PatchGuard
	notAllowedErr := fmt.Errorf("only https requests allowed")
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient),
		"Get", func(c *http.Client, url string) (*http.Response, error) {
			guard.Unpatch()
			defer guard.Restore()

			if !strings.HasPrefix(url, "https://") {
				return nil, notAllowedErr
			}

			return c.Get(url)
		})

	_, err := http.Get("http://google.com")
	if err != notAllowedErr {
		t.Errorf("got error: %v, want error: %v", err, notAllowedErr)
	}
	resp, err := http.Get("https://google.com")
	if err != nil || resp.StatusCode != 200 {
		t.Error("should be 200 OK")
	}
	monkey.UnpatchAll()
}
