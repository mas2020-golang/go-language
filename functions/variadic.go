/*
Example that show how to use the variadic functions.
*/
package main

import "strings"

func variadic(elems ...string) string {
	//var res string
	//for _, s := range elems {
	//	res += fmt.Sprintf("%s-", s)
	//}
	//return res[:len(res)-1]
	return strings.Join(elems, "-")
}
