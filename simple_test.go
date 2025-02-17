package main

import "testing"

func TestAdd(t *testing.T) {
    got := add(2, 3)
    want := 5

    if got != want {
        t.Errorf("add(2, 3) = %d; want %d", got, want)
    }
}
