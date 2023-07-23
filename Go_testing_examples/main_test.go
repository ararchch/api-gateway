// test by running go test -v (to see all test cases)
package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
        t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestSum2(t *testing.T) {
    total := Sum2(5,5)
    if total != 10 {
        t.Errorf("Sum was incorrect, got %d, want %d.", total, 10)
    }

}