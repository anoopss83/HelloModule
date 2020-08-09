package hello

import(
	"testing"
)

func TestHello(t *testing.T){
	want := "Hello, world."
	got := Hello()
	if got != want {
		t.Errorf("Hello() got = %s  want = %s",got,want)
	}
}

func TestProverb(t *testing.T){
	want := "Concurrency is not parallelism."
	got := Proverb()
	if got != want {
		t.Errorf("Proverb() got = %s want = %s",got,want)
	}
}