package llm

import (
	"reflect"
	"testing"
)

func TestCallLocalLLM(t *testing.T) {
	type args struct {
		sessionID *string
		message   *string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallLocalLLM(tt.args.sessionID, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallLocalLLM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendAPIRequest(t *testing.T) {
	type args struct {
		sessionID *string
		message   *string
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendAPIRequest(tt.args.sessionID, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendAPIRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendAPIRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
