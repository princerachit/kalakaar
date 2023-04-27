package configuration

import (
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
			name: "TestDefaultStep_Execute",
			fields: fields{
				id:      "test",
				name:    "test",
				command: "echo testoutput",
			},
			result: &DefaultResult{result: []string{"testoutput"}},
			errors: nil,
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
			res := strings.TrimSuffix(result.Get().(string), "\n")

			if !reflect.DeepEqual(res, tt.result.Get()) {
				t.Errorf("DefaultStep.Execute() got = %v, want = %v", res, tt.result.Get())
			}
			// TODO: Fix this test
			if !reflect.DeepEqual(errs, tt.errors) {
				t.Errorf("DefaultStep.Execute() got1 = %v, want = %v", errs, tt.errors)
			}
		})
	}
}
