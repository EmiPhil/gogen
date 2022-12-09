package golang

import (
	"github.com/EmiPhil/gogen/code"
	"reflect"
	"testing"
)

func TestFunc_SetLeftBracketPrefix(t *testing.T) {
	type fields struct {
		Name    string
		Params  []string
		Returns []string
		Block   *code.Block
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := &Func{
				Name:    tt.fields.Name,
				Params:  tt.fields.Params,
				Returns: tt.fields.Returns,
				Block:   tt.fields.Block,
			}
			fn.SetLeftBracketPrefix()
		})
	}
}

func TestParams_String(t *testing.T) {
	tests := []struct {
		name string
		p    Params
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturns_String(t *testing.T) {
	tests := []struct {
		name string
		r    Returns
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScript_WriteFunc(t *testing.T) {
	type args struct {
		descriptor string
		name       string
		params     []string
		returns    []string
	}
	tests := []struct {
		name   string
		script *Script
		args   args
		want   string
	}{
		{
			name:   "basic",
			script: New("test"),
			args: args{
				descriptor: "helloWorld prints hello world!",
				name:       "helloWorld",
				params:     nil,
				returns:    nil,
			},
			want: `package test

// helloWorld prints hello world!
func helloWorld() {
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.script
			s.WriteFunc(tt.args.descriptor, tt.args.name, tt.args.params, tt.args.returns)
			if got := s.Render(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
