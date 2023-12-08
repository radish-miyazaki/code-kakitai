package product

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/radish-miyazaki/go-pkg/ulid"
)

func Test_NewProduct(t *testing.T) {
	ownerID := ulid.NewULID()

	type args struct {
		ownerID     string
		name        string
		description string
		price       int64
		stock       int
	}
	tests := []struct {
		name    string
		args    args
		want    *Product
		wantErr bool
	}{
		{
			name: "Valid",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			want: &Product{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			wantErr: false,
		},
		{
			name: "Invalid: owner id is invalid",
			args: args{
				ownerID:     "invalid",
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid: name is invalid",
			args: args{
				ownerID:     ownerID,
				name:        "",
				description: "test",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid: description is invalid",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid: price is invalid",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       0,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid: stock is invalid",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       100,
				stock:       -1,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduct(tt.args.ownerID, tt.args.name, tt.args.description, tt.args.price, tt.args.stock)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(Product{}), cmpopts.IgnoreFields(Product{}, "id"))
			if diff != "" {
				t.Errorf("diff: %s", diff)
			}
		})
	}
}
