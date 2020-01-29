package uuid5

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Homo sapiens", "16f235a0-e4a3-529c-9b83-bd15fe722110"},
		{"Betula verucosa", "4c19ac07-ec67-5cff-97bf-7d9ecbe12e34"},
		// count spaces before and after
		{" Betula verucosa \n ", "d764732b-ed00-5e31-a2b9-de044ad6b4c5"},
		// double spaces are preserved
		{"Betula  verucosa", "2883a5c0-d425-518f-8b21-8bd434840c93"},
	}
	for _, c := range cases {
		got := UUID5(c.in).String()
		if got != c.want {
			t.Errorf("uuid5.String of '%s' == '%s', want '%s'", c.in, got, c.want)
		}
	}
}

func TestStrings(t *testing.T) {
	in := []string{"Homo sapiens", "Betula verucosa",
		" Betula verucosa \n", "Betula  verucosa"}
	want := []string{"16f235a0-e4a3-529c-9b83-bd15fe722110",
		"4c19ac07-ec67-5cff-97bf-7d9ecbe12e34",
		"8e695ff8-dbd8-5ab7-b69d-22c4ebfd2ef9",
		"2883a5c0-d425-518f-8b21-8bd434840c93"}
	uuids := UUID5s(in)
	got := Strings(uuids)
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("uuid.Strings: %s == %s, want %s for element %d\n", in[i], got[i], want[i], i)
		}
	}
}

// Benchmarks. To run all of them use
// go test ./... -bench=. -benchmem -count=10 -run=XXX > bench.txt && benchstat bench.txt

// BenchmarkUUIDString checks the speed of generating a string version of UUID
func BenchmarkUUIDString(b *testing.B) {
	traceFile := "uuid-string.trace"
	runBenchmark("UUIDString", b, traceFile, true)
}

// BenchmarkUUIDRaw checks the speed of generating a string version of UUID
func BenchmarkUUIDRaw(b *testing.B) {
	traceFile := "uuid-raw.trace"
	runBenchmark("UUIDRaw", b, traceFile, false)
}

// BenchmarkUUID checks the speed of generating a string version of UUID
func BenchmarkUUID(b *testing.B) {
	gnUUID := uuid.NewV5(uuid.NamespaceDNS, "globalnames.org")
	n := "UUID"
	var res uuid.UUID
	b.Run(n, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res = uuid.NewV5(gnUUID, "Influenza A virus (A/blue-winged teal/Missouri/11OS2563/2011(H12N4))")
			_ = res.String
		}
		_ = fmt.Sprintf("%d", len(res))
	})
}

func runBenchmark(n string, b *testing.B, traceFile string,
	withString bool) {
	f, err := os.Create(traceFile)
	if err != nil {
		panic(err)
	}
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer b.StopTimer()
	defer trace.Stop()

	var res uuid.UUID
	b.Run(n, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res = UUID5("Influenza A virus (A/blue-winged teal/Missouri/11OS2563/2011(H12N4))")
			if withString {
				_ = res.String
			}
		}

		_ = fmt.Sprintf("%d", len(res))
	})
}
