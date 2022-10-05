package server

import (
	"context"

	models "github.com/Abeldlp/manager-mail/Models"
	"github.com/Abeldlp/manager-mail/config"
	"github.com/Abeldlp/manager-mail/mailpb"
)

type Server struct {
	mailpb.UnimplementedMailServiceServer
}

func (*Server) GetAllMails(ctx context.Context, req *mailpb.GetAllMailRequest) (*mailpb.GetAllMailResponse, error) {
	var mailRes []models.Mail
	var mails []*mailpb.Mail
	config.DB.Order("created_at desc").Find(&mailRes)

	for _, item := range mailRes {
		mails = append(mails, &mailpb.Mail{
			Subject:  item.Subject,
			Sender:   item.Sender,
			Receiver: item.Receiver,
			Content:  item.Content,
		})
	}

	res := &mailpb.GetAllMailResponse{
		Mail: mails,
	}

	return res, nil
}

func (*Server) CreateMail(ctx context.Context, req *mailpb.CreateMailRequest) (*mailpb.CreateMailResponse, error) {
	config.DB.Create(&models.Mail{
		Subject:  req.Mail.Subject,
		Sender:   req.Mail.Sender,
		Receiver: req.Mail.Receiver,
		Content:  req.Mail.Content,
	})

	returnResponse := &mailpb.CreateMailResponse{
		Mail: &mailpb.Mail{
			Subject:  req.Mail.Subject,
			Sender:   req.Mail.Sender,
			Receiver: req.Mail.Receiver,
			Content:  req.Mail.Content,
		},
	}

	return returnResponse, nil
}
