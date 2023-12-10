package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	userDomain "github.com/radish-miyazaki/code-kakitai/domain/user"
)

func TestUserRepository_FindByID(t *testing.T) {
	address, _ := userDomain.NewAddress("宮崎県", "宮崎市", "青島1-1-1")
	user, _ := userDomain.Reconstruct(
		"01HCNYK0PKYZWB0ZT1KR0EPWGP",
		"example@test.com",
		"08000000000",
		"山田",
		"太郎",
		*address,
	)

	tests := []struct {
		name string
		want *userDomain.User
	}{
		{
			name: "Valid",
			want: user,
		},
	}
	userRepo := NewUserRepository()
	ctx := context.Background()
	resetTestData(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userRepo.FindByID(ctx, "01HCNYK0PKYZWB0ZT1KR0EPWGP")
			if err != nil {
				t.Error(err)
			}
			if got == nil {
				t.Error("got is nil")
			}

			fmt.Println(got)
			fmt.Println(tt.want)

			if diff := cmp.Diff(got.ID(), tt.want.ID()); diff != "" {
				t.Errorf("FindByID() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
