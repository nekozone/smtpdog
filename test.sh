#!/bin/bash
swaks --to username+dd@m.dogcraft.top  --from username@example.com --header-X-Mailer SMTP  --ehlo example.com --header "Subject:测试" --body '已收到\n首先进入“Mail”，点击“其他”，新建一个邮件帐户，输入您邮箱的完整地址，点击“存储”。\nSMTP瑞士军刀Swaks是由John Jetmore编写和维护的一种功能强大，灵活，可脚本化，面向 事务的SMTP测试工具。可向任意目标发送任意内容的邮件\n' --h-From: 'example<username@example.com>' --attach go.sum
