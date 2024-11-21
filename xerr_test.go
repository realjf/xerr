// #############################################################################
// # File: xerr_test.go                                                        #
// # Project: xerr                                                             #
// # Created Date: 2024/11/21 18:31:01                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/11/21 19:00:46                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package xerr_test

import (
	"fmt"
	"testing"

	"github.com/realjf/xerr"
)

func TestNew(t *testing.T) {
	e := xerr.New()
	fmt.Printf("%s\n", e)

	e2 := xerr.WrapError(e)
	fmt.Printf("%s\n", e2)

	e3 := xerr.NewCode(2)
	fmt.Printf("%s\n", e3)

	e4 := xerr.NewError("hello")
	fmt.Printf("%s\n", e4)
}
