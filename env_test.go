package env

import (
	"testing"

	"github.com/code-gorilla-au/env/internal/testutils"
)

func TestGetAsString(t *testing.T) {
	t.Setenv("TEST_ENV", "test")

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return string",
			args: args{
				key: "TEST_ENV",
			},
			want: "test",
		},
		{
			name: "should return empty string",
			args: args{
				key: "FAKE_ENV",
			},
			want: "",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := GetAsString(tt.args.key)
			testutils.AssertEqual(t, tt.want, got)
		})
	}
}

func TestGetAsInt(t *testing.T) {
	t.Setenv("TEST_ENV", "1")

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return int",
			args: args{
				name: "TEST_ENV",
			},
			want: 1,
		},
		{
			name: "should return 0",
			args: args{
				name: "FAKE_ENV",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAsInt(tt.args.name)
			testutils.AssertEqual(t, tt.want, got)
		})
	}
}

func TestGetAsBool(t *testing.T) {
	t.Setenv("TEST_ENV", "true")

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{
				name: "TEST_ENV",
			},
			want: true,
		},
		{
			name: "should return false if no env var",
			args: args{
				name: "FAKE_ENV",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAsBool(tt.args.name)
			testutils.AssertEqual(t, tt.want, got)
		})
	}
}

func TestGetAsSlice(t *testing.T) {
	t.Setenv("TEST_ENV", "hello,world")

	type args struct {
		name string
		sep  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should return slice",
			args: args{
				name: "TEST_ENV",
			},
			want: []string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAsSlice(tt.args.name, tt.args.sep)
			testutils.AssertEqual(t, tt.want, got)

		})
	}
}

func TestPanicSlice(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetAsSlice did not panic")
		}
	}()

	WithStrictMode()

	result := GetAsSlice("TEST_ENV", ",")
	testutils.AssertEqual(t, []string{"hello", "world"}, result)
}

func TestPanicInt(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetAsInt did not panic")
		}
	}()

	WithStrictMode()

	GetAsInt("TEST_ENV")
}

func TestNonStrictInt(t *testing.T) {
	t.Setenv("TEST_ENV", "hello")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetAsInt did not panic")
		}
	}()

	WithStrictMode()

	GetAsInt("TEST_ENV")
}

func TestPanicBool(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetAsBool did not panic")
		}
	}()

	WithStrictMode()

	GetAsBool("FLASH")
}

func TestGetAsStringWithDefault(t *testing.T) {

	t.Setenv("TEST_ENV", "test")

	result := GetAsStringWithDefault("TEST_ENV", "default")
	testutils.AssertEqual(t, "test", result)
}

func TestGetAsStringWithDefaultShouldReturnDefault(t *testing.T) {

	result := GetAsStringWithDefault("TEST_ENV", "default")
	testutils.AssertEqual(t, "default", result)
}

func TestGetAsIntWithDefault(t *testing.T) {
	t.Setenv("TEST_ENV", "99")

	result := GetAsIntWithDefault("TEST_ENV", 42)
	testutils.AssertEqual(t, 99, result)
}

func TestGetAsIntWithDefaultShouldReturnDefault(t *testing.T) {

	result := GetAsIntWithDefault("TEST_ENV", 42)
	testutils.AssertEqual(t, 42, result)
}

func TestGetAsBoolWithDefault(t *testing.T) {
	t.Setenv("TEST_ENV", "true")

	result := GetAsBoolWithDefault("TEST_ENV", false)
	testutils.AssertEqual(t, true, result)
}

func TestGetAsBoolWithDefaultShouldReturnDefault(t *testing.T) {

	result := GetAsBoolWithDefault("TEST_ENV", true)
	testutils.AssertEqual(t, true, result)
}

func TestNonStrictNonBool(t *testing.T) {
	t.Setenv("TEST_ENV", "hello")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetAsInt did not panic")
		}
	}()

	WithStrictMode()

	GetAsBool("TEST_ENV")
}

func TestGetAsSliceWithDefault(t *testing.T) {
	t.Setenv("TEST_ENV", "hello,world")

	result := GetAsSliceWithDefault("TEST_ENV", ",", []string{"default"})
	testutils.AssertEqual(t, []string{"hello", "world"}, result)
}

func TestGetAsSliceWithDefaultShouldReturnDefault(t *testing.T) {

	result := GetAsSliceWithDefault("TEST_ENV", ",", []string{"default"})
	testutils.AssertEqual(t, []string{"default"}, result)
}
