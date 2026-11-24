package main

import "testing"

func TestDefaultHostname(t *testing.T) {
	tests := []struct {
		targetHost string
		want       string
		wantErr    bool
	}{
		{targetHost: "pdp1153.example.com", want: "pdp1153"},
		{targetHost: "my-host.example.com", want: "my-host"},
		{targetHost: "printer", want: "printer"},
		{targetHost: "192.0.2.7", wantErr: true},
		{targetHost: "2001:db8::7", wantErr: true},
		{targetHost: "", wantErr: true},
		{targetHost: ".example.com", wantErr: true},
	}
	for _, tt := range tests {
		got, err := defaultHostname(tt.targetHost)
		if (err != nil) != tt.wantErr {
			t.Errorf("defaultHostname(%q) error = %v, wantErr %v", tt.targetHost, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("defaultHostname(%q) = %q, want %q", tt.targetHost, got, tt.want)
		}
	}
}
