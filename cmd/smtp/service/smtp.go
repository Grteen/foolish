package service

import (
	"be/cmd/smtp/dal/rdb"
	"be/grpc/msmtpdemo"
	"be/pkg/errno"
	"be/pkg/kafka"
	"be/pkg/uuid"
	"context"
	"net/smtp"

	"github.com/redis/go-redis/v9"
)

// 邮箱服务器配置信息
type configInof struct {
	smtpAddr string
	smtpPort string
	secret   string
}

// 邮件内容信息
type emailContent struct {
	fromAddr    string
	contentType string
	theme       string
	message     string
	toAddr      []string
}

func sendEmail(c *configInof, e *emailContent) error {
	// 拼接smtp服务器地址
	smtpAddr := c.smtpAddr + ":" + c.smtpPort
	// 认证信息
	auth := smtp.PlainAuth("", e.fromAddr, c.secret, c.smtpAddr)
	// 配置邮件内容类型
	if e.contentType == "html" {
		e.contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		e.contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	// 当有多个收件人
	for _, to := range e.toAddr {
		msg := []byte("To: " + to + "\r\n" +
			"From: " + e.fromAddr + "\r\n" +
			"Subject: " + e.theme + "\r\n" +
			e.contentType + "\r\n\r\n" +
			"<html><h1>" + e.message + "</h1></html>")
		err := smtp.SendMail(smtpAddr, auth, e.fromAddr, []string{to}, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

type SmtpService struct {
	ctx context.Context
}

func NewSmtpService(ctx context.Context) *SmtpService {
	return &SmtpService{ctx: ctx}
}

// 发送邮件
func (s *SmtpService) SendSmtp(req *msmtpdemo.SendSmtpRequest) error {
	config := configInof{
		smtpAddr: "smtp.qq.com",
		secret:   "vpngxnpgjzehdghh",
		smtpPort: "25",
	}

	verify := uuid.GetVerify()
	rdb.SetRegisterVerify(s.ctx, req.Email, verify)
	content := emailContent{
		fromAddr:    "2106633192@qq.com",
		toAddr:      []string{req.Email},
		contentType: "html",
		theme:       "校园论坛注册",
		message:     verify,
	}

	err := sendEmail(&config, &content)
	if err != nil {
		kafka.ErrorLog(err.Error())
		return errno.ServiceFault
	} else {
		kafka.AccessLog(req.Email + " try to register")
	}

	return nil
}

// 查询redis内的验证码
func (s *SmtpService) QueryVerify(req *msmtpdemo.QueryVerifyRequest) (string, error) {
	verify, err := rdb.QueryRegisterVerify(s.ctx, req.Email)
	if err != nil {
		kafka.ErrorLog(err.Error())
		if err != redis.Nil {
			return "", errno.ServiceFault
		} else {
			return "", errno.WrongVerifyErr
		}
	}

	return verify, nil
}
