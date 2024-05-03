package main

import (
	"fmt"
	"testing"
)

// 发送成功的假邮箱
type fakeOkMailer struct{}

func (m *fakeOkMailer) SendMail(subject string, dest string, body string) error {
	return nil
}

// 发送失败的假邮箱
type fakeFailMailer struct{}

func (m *fakeFailMailer) SendMail(subject string, dest string, body string) error {
	return fmt.Errorf("can not reach the mail server of dest [%s]", dest)
}

// 测试 邮箱发送成功
func TestComposeAndSendOk(t *testing.T) {
	m := &fakeOkMailer{}
	mc := New(m)
	_, err := mc.ComposeAndSend("hello, fake test", []string{"xxx@example.com"}, "the test body")
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}
}

// 测试 邮箱发送失败
func TestComposeAndSendFail(t *testing.T) {
	m := &fakeFailMailer{}
	mc := New(m)
	_, err := mc.ComposeAndSend("hello, fake test", []string{"xxx@example.com"}, "the test body")
	if err == nil {
		t.Errorf("want non-nil, got nil")
	}
}
