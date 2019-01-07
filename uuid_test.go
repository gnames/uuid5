package uuid5

import (
	"testing"
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
