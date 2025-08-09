package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRequest_Validate(t *testing.T) {

	type args struct {
		createRequest CreateRequest
	}

	type want struct {
		valid bool
		cause string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Valid CreateRequest",
			args: args{
				createRequest: CreateRequest{
					Username: "test",
					Email:    "test@email.com",
				},
			},
			want: want{
				valid: true,
				cause: "",
			},
		},
		{
			name: "Invalid CreateRequest - missing username",
			args: args{
				createRequest: CreateRequest{
					Username: "",
					Email:    "test@email.com",
				},
			},
			want: want{
				valid: false,
				cause: "User name is required",
			},
		},
		{
			name: "Invalid CreateRequest - missing email",
			args: args{
				createRequest: CreateRequest{
					Username: "user",
					Email:    "",
				},
			},
			want: want{
				valid: false,
				cause: "Email is required",
			},
		},
		{
			name: "Invalid CreateRequest - invalid email",
			args: args{
				createRequest: CreateRequest{
					Username: "user",
					Email:    "fakemail",
				},
			},
			want: want{
				valid: false,
				cause: "Email is invalid",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, cause := tt.args.createRequest.Validate()
			assert.Equal(t, tt.want.valid, valid)
			assert.Equal(t, tt.want.cause, cause)
		})
	}
}
