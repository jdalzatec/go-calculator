package evaluator_test

import (
	"testing"

	"github.com/jdalzatec/calculator/pkg/evaluator"
	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	type args struct {
		lhs      int64
		rhs      int64
		operator string
	}
	tests := []struct {
		name    string
		args    args
		want    *int64
		wantErr bool
	}{
		{
			name: "adds two numbers",
			args: args{
				lhs:      1,
				rhs:      2,
				operator: "+",
			},
			want:    new(int64(3)),
			wantErr: false,
		},
		{
			name: "subtracts two numbers",
			args: args{
				lhs:      1,
				rhs:      2,
				operator: "-",
			},
			want:    new(int64(-1)),
			wantErr: false,
		},
		{
			name: "multiply two numbers",
			args: args{
				lhs:      8,
				rhs:      7,
				operator: "*",
			},
			want:    new(int64(56)),
			wantErr: false,
		},
		{
			name: "division with rhs zero",
			args: args{
				lhs:      1,
				rhs:      0,
				operator: "/",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "divide two numbers",
			args: args{
				lhs:      1,
				rhs:      2,
				operator: "/",
			},
			want:    new(int64(0)),
			wantErr: false,
		},
		{
			name: "unknown operator",
			args: args{
				lhs:      1,
				rhs:      2,
				operator: "^",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evaluator.Evaluate(tt.args.lhs, tt.args.rhs, tt.args.operator)
			if tt.wantErr {
				assert.Nil(t, got)
				assert.Error(t, err)
			} else {
				assert.NoErrorf(t, err, "evaluator.Evaluate() error = %v", err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
