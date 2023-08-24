package utils

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmailCode(srcEmail, toEmail string, emailKey string, code string) error {
	e := email.NewEmail()            //实例化
	e.From = "渊虹y<" + srcEmail + ">" //谁发的
	e.To = []string{toEmail}         //给谁发guojunlh@qq.com
	e.Bcc = []string{toEmail}        //抄送给谁
	e.Subject = "邮箱验证码"              //标题
	//e.Text = []byte("邮件内容")              //普通文本
	e.HTML = []byte(toEmail + ":您的邮箱验证码为‘" + code + "’,有效时间五分钟，请于五分钟之内填写") //html文本
	//服务器域名和端口号			授权 							邮箱中的授权码
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", srcEmail, emailKey, "smtp.163.com"))
	if err != nil {
		fmt.Printf("err %v ", err)
		return err
	}
	//fmt.Println("发送成功")
	fmt.Printf("code: %v\n", code)
	return nil

}
func SendAccountEmail(srcEmail, toEmail string, emailKey string, Account string) error {
	e := email.NewEmail()            //实例化
	e.From = "渊虹y<" + srcEmail + ">" //谁发的
	e.To = []string{toEmail}         //给谁发guojunlh@qq.com
	e.Bcc = []string{toEmail}        //抄送给谁
	e.Subject = "账号注册"               //标题
	//e.Text = []byte("邮件内容")              //普通文本
	e.HTML = []byte(toEmail + ":您'富硒农产品系统'注册的账号为<a>" + Account + "</a>,密码为您注册时填写的密码,欢迎使用本系统！") //hrml文本
	//服务器域名和端口号			授权 							邮箱中的授权码
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", srcEmail, emailKey, "smtp.163.com"))
	if err != nil {
		return err
	}
	//fmt.Println("发送成功")
	return nil
}

// 自己通知自己的就好了
func IssuesEmail(srcEmail string, emailKey string, title string, content string) error {
	e := email.NewEmail()            //实例化
	e.From = "渊虹y<" + srcEmail + ">" //谁发的
	e.To = []string{srcEmail}        //给谁发
	e.Bcc = []string{srcEmail}       //抄送给谁
	e.Subject = "Issues邮件"           //标题
	//e.Text = []byte("邮件内容")              //普通文本
	e.HTML = []byte("主题:<p>" + title + "</p>详细描述：<p>" + content + "<p>") //hrml文本
	//服务器域名和端口号			授权 							邮箱中的授权码
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", srcEmail, emailKey, "smtp.163.com"))
	if err != nil {
		return err
	}
	return nil
}
