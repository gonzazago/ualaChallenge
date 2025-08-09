package post

import (
	"github.com/go-playground/assert/v2"
	"strings"
	"testing"
)

func TestCreatePostRequest_Validate(t *testing.T) {

	type args struct {
		createRequest CreatePostRequest
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
			name: "validate success",
			args: args{
				createRequest: CreatePostRequest{
					UserID: "user-id",
					Text:   "text",
				},
			},
			want: want{
				valid: true,
				cause: "",
			},
		},
		{
			name: "validate failure - user missing",
			args: args{
				createRequest: CreatePostRequest{
					UserID: "",
					Text:   "text",
				},
			},
			want: want{
				valid: false,
				cause: "user id is required",
			},
		},
		{
			name: "validate failure - text missing",
			args: args{
				createRequest: CreatePostRequest{
					UserID: "user-id",
					Text:   "",
				},
			},
			want: want{
				valid: false,
				cause: "text is required",
			},
		},
		{
			name: "validate failure - text to long",
			args: args{
				createRequest: CreatePostRequest{
					UserID: "user-id",
					Text:   strings.Repeat("a", 281),
				},
			},
			want: want{
				valid: false,
				cause: "text is too long, not exceed 280 characters",
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
