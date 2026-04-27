package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/idna"
)

type entry struct {
	Domains []string `json:"domains"`
}

var alwaysTLDs = []string{"ru", "рф", "moscow", "москва", "рус", "by", "su"}

func tldPlusOne(domain string) string {
	d := strings.TrimSpace(strings.ToLower(domain))
	d = strings.TrimSuffix(d, ".")
	if d == "" {
		return ""
	}

	parts := strings.Split(d, ".")
	if len(parts) < 2 {
		return ""
	}

	return parts[len(parts)-2] + "." + parts[len(parts)-1]
}

func main() {
	input := flag.String("input", "ip-list.json", "Path to source JSON file")
	flag.Parse()
	if flag.NArg() > 0 {
		*input = flag.Arg(0)
	}

	data, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %q: %v\n", *input, err)
		os.Exit(1)
	}

	items := map[string]entry{}
	if err := json.Unmarshal(data, &items); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse %q: %v\n", *input, err)
		os.Exit(1)
	}

	for _, d := range extractDNS(items) {
		fmt.Println(d)
	}
}

func extractDNS(items map[string]entry) []string {
	always := alwaysTLDs

	alwaysSet := map[string]struct{}{}
	for _, tld := range always {
		punycode, _ := idna.ToASCII(tld)
		alwaysSet[tld] = struct{}{}
		alwaysSet[punycode] = struct{}{}
	}

	collected := map[string]struct{}{}
	for _, item := range items {
		for _, domain := range item.Domains {
			domain = trimPort(domain)
			t1 := tldPlusOne(domain)
			if t1 == "" {
				continue
			}

			parts := strings.Split(t1, ".")
			tld := parts[len(parts)-1]
			if _, ok := alwaysSet[tld]; ok {
				continue
			}

			collected[t1] = struct{}{}
		}
	}

	result := make([]string, 0, len(collected))
	for d := range collected {
		result = append(result, d)
	}
	sort.Strings(result)

	out := make([]string, 0, len(always)+len(result))
	out = append(out, always...)
	out = append(out, result...)
	return out
}

func trimPort(domain string) string {
	if idx := strings.LastIndex(domain, ":"); idx != -1 {
		return domain[:idx]
	}
	return domain
}
