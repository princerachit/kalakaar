package configuration

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestDefaultStep_Execute(t *testing.T) {
	type fields struct {
		id      string
		name    string
		command string
	}
	tests := []struct {
		name   string
		fields fields
		result Result
		errors []error
	}{
		{
			name: "Success",
			fields: fields{
				id:      "test",
				name:    "test",
				command: "echo testoutput",
			},
			result: &DefaultResult{result: []string{"testoutput"}},
			errors: nil,
		},
		{
			name: "Failure",
			fields: fields{
				id:      "test",
				name:    "test",
				command: "nonexistentbinary testoutput",
			},
			result: nil,
			errors: []error{errors.New("exit status 127")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DefaultStep{
				id:      tt.fields.id,
				name:    tt.fields.name,
				command: tt.fields.command,
			}
			result, errs := s.Execute()
			if tt.result != nil {
				res := strings.TrimSuffix(result.Get().(string), "\n")
				if !reflect.DeepEqual(res, tt.result.Get()) {
					t.Errorf("DefaultStep.Execute() got = %v, want = %v", res, tt.result.Get())
				}
			} else {
				// // TODO: Fix this test
				got := errs[0].Error()
				want := tt.errors[0].Error()
				if got != want {
					t.Errorf("DefaultStep.Execute() got = %v, want = %v", got, want)
				}
			}

		})
	}
}
