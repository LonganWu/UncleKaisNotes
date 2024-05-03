package main

// Mailer 邮箱
type Mailer interface {
	SendMail(subject, destination, body string) error
}
type mailClient struct {
	mlr Mailer
}

func New(mlr Mailer) *mailClient {
	return &mailClient{
		mlr: mlr,
	}
}

// ComposeAndSend 将传入的电子邮件内容（body）与签名（signTxt）编排合并后传给Mailer接口实现者的SendMail方法，由其将邮件发送出去。
func (c *mailClient) ComposeAndSend(subject string,
	destinations []string, body string) (string, error) {
	signTxt := sign.Get()
	newBody := body + "\n" + signTxt

	for _, dest := range destinations {
		err := c.mlr.SendMail(subject, dest, newBody)
		if err != nil {
			return "", err
		}
	}
	return newBody, nil
}

// Sign 获取签名
type Sign struct {
}

func (s Sign) Get() string {
	return "sign string"
}

var getSign = Get

func Get(string) string {
	return "sign"
}

var sign Sign
