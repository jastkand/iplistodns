package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestTldPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{name: "subdomain", input: "a.b.example.com", output: "example.com"},
		{name: "trim and lowercase", input: "  API.MAX.RU ", output: "max.ru"},
		{name: "trailing dot", input: "yandex.net.", output: "yandex.net"},
		{name: "single label", input: "localhost", output: ""},
		{name: "empty", input: "", output: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tldPlusOne(tt.input)
			if got != tt.output {
				t.Fatalf("tldPlusOne(%q) = %q, want %q", tt.input, got, tt.output)
			}
		})
	}
}

func TestTrimPort(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{name: "has port", input: "example.com:443", output: "example.com"},
		{name: "no port", input: "example.com", output: "example.com"},
		{name: "empty", input: "", output: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trimPort(tt.input)
			if got != tt.output {
				t.Fatalf("trimPort(%q) = %q, want %q", tt.input, got, tt.output)
			}
		})
	}
}

func TestExtractDNSFromFixture(t *testing.T) {
	data, err := os.ReadFile("tests/ip-list-test.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	items := map[string]entry{}
	if err := json.Unmarshal(data, &items); err != nil {
		t.Fatalf("failed to parse fixture: %v", err)
	}

	got := extractDNS(items)
	want := []string{
		"ru",
		"рф",
		"moscow",
		"москва",
		"рус",
		"by",
		"su",
		"mycdn.me",
		"vk.com",
		"vkuser.net",
		"yandex.com",
		"yandex.net",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("extractDNS() mismatch\n got: %#v\nwant: %#v", got, want)
	}
}
