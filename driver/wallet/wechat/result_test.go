package wechat

import (
	"testing"
)

func TestResult(t *testing.T) {
	cases := []struct {
		result *Result
		expect *PayResponse
		pass   bool
	}{
		{
			&Result{
				Body: []byte(`{"code_url":"https://xxx.com"}`),
			},
			&PayResponse{CodeUrl: "https://xxx.com"},
			true,
		},
		{
			&Result{
				Err: &Error{},
			},
			&PayResponse{CodeUrl: "https://xxx.com"},
			false,
		},
		{
			&Result{
				Body: []byte(``),
			},
			&PayResponse{CodeUrl: "https://xxx.com"},
			true,
		},
		{
			&Result{},
			&PayResponse{CodeUrl: "https://xxx.com"},
			true,
		},
		{
			&Result{
				Body: []byte(`{"`),
			},
			&PayResponse{CodeUrl: "https://xxx.com"},
			false,
		},
	}

	for _, c := range cases {
		dest := &PayResponse{}
		err := c.result.Scan(dest)
		pass := err == nil
		if pass != c.pass {
			t.Fatalf("expect %v, got %v, err %v", c.pass, pass, err)
		}
	}
}

func TestError(t *testing.T) {
	cases := []struct {
		err    *Error
		expect string
	}{
		{
			&Error{400, "code", "message"},
			`{"status":400,"code":"code","message":"message"}`,
		},
		{
			nil,
			"{}",
		},
	}

	for _, c := range cases {
		actual := c.err.Error()
		if actual != c.expect {
			t.Fatalf("expect %s, got %s", c.expect, actual)
		}
	}
}
