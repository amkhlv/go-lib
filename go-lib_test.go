/**
 * Author: Andrei Mikhailov
 * File: go-lib_test.go
 */

package amkhlv

import (
	"regexp"
	"testing"
)

func TestUnwrapResult(t *testing.T) {

	r := UnwrapResult(regexp.Compile("^a"))

	if r.MatchString("aaa") {
	} else {
		t.Errorf("did not work...")
	}
}
