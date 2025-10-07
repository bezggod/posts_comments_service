package user_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_commets_service/internal/domain/models"
	"testing"
)

func TestUserUseCase_Create(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		req CreateUserReq
	}

	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: CreateUserReq{
					Name: "TEST",
				},
			},
			want: &models.User{
				Name: "TEST",
			},
			before: func(m mockService, args args) {

				m.userRepo.EXPECT().Create(mock.Anything, args.req.Name).Return(&models.User{
					Name: "TEST",
				}, nil)
			},
		},
		{
			name: "error on Create",
			args: args{
				req: CreateUserReq{
					Name: "TEST",
				},
			},
			wantErr: true,
			before: func(m mockService, args args) {

				m.userRepo.EXPECT().Create(mock.Anything, args.req.Name).Return(nil, errTest)
			},
		},
		{
			name: "empty name",
			args: args{
				req: CreateUserReq{
					Name: "",
				},
			},
			wantErr: true,
			before:  func(m mockService, args args) {},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			usecase, mocks := makeService(t)

			if tc.before != nil {
				tc.before(mocks, tc.args)
			}
			user, err := usecase.Create(context.Background(), tc.args.req.Name)
			if tc.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tc.want, user)
		})
	}
}
