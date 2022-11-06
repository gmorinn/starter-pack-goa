package api

import (
	"context"
	"fmt"
	"log"
	"os"
	db "starter-pack-goa/internal"
	"starter-pack-goa/utils"
	"strconv"

	"github.com/mailjet/mailjet-apiv3-go"
)

func (server *Server) sendEmail(ctx context.Context, id, email string) error {
	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {
		templateAPI, err := strconv.ParseInt(os.Getenv("API_MAILJET_SEND_CODE"), 10, 64)
		if err != nil {
			return fmt.Errorf("ERROR_PARSE_INT %s", err.Error())
		}
		code := utils.RandStringRunes(5)
		arg := db.UpdateUserConfirmCodeParams{
			Email:               email,
			PasswordConfirmCode: utils.NullS(code),
		}
		if err := q.UpdateUserConfirmCode(ctx, arg); err != nil {
			return fmt.Errorf("UPDATE_USER_CONFIRM_CODE %s", err.Error())
		}
		m := mailjet.NewMailjetClient(os.Getenv("API_MAILJET_KEY"), os.Getenv("API_MAILJET_SECRET"))
		messagesInfo := []mailjet.InfoMessagesV31{
			mailjet.InfoMessagesV31{
				From: &mailjet.RecipientV31{
					Email: os.Getenv("API_MAIL_FROM"),
					Name:  "GM API",
				},
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: email,
						Name:  " ",
					},
				},
				TemplateID:       int(templateAPI),
				TemplateLanguage: true,
				Subject:          "Votre code de confirmation",
				Variables:        map[string]interface{}{"code": code},
			},
		}
		messages := mailjet.MessagesV31{Info: messagesInfo}
		_, err = m.SendMailV31(&messages)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
