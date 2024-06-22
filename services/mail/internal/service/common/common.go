package common

import (
	"fmt"

	pb "github.com/gstones/platform/api/gen/mail"
)

func FilterMailsMapWithLanguage(mails map[int64]*pb.Mail, language string) (map[int64]*pb.Mail, error) {
	res := make(map[int64]*pb.Mail)
	for k, v := range mails {
		body, err := filterMailsWithLanguage(v.Body, language)
		if err != nil {
			return nil, err
		}
		v.Body = body
		title, err := filterMailsWithLanguage(v.Title, language)
		if err != nil {
			return nil, err
		}
		v.Title = title
		res[k] = v
	}
	return res, nil
}

func FilterMailsWithLanguage(mails []*pb.Mail, language string) ([]*pb.Mail, error) {
	res := make([]*pb.Mail, 0)
	for _, v := range mails {
		body, err := filterMailsWithLanguage(v.Body, language)
		if err != nil {
			return nil, err
		}
		v.Body = body
		title, err := filterMailsWithLanguage(v.Title, language)
		if err != nil {
			return nil, err
		}
		v.Title = title
		res = append(res, v)
	}
	return res, nil
}

func FilterMailsMapWithRegisterTime(mails map[int64]*pb.Mail, registerTime int64) map[int64]*pb.Mail {
	res := make(map[int64]*pb.Mail)
	for k, v := range mails {
		if v.Filters.RegisterTime >= registerTime {
			res[k] = v
		}
	}
	return res
}

func FilterMailsWithRegisterTime(mails []*pb.Mail, registerTime int64) []*pb.Mail {
	res := make([]*pb.Mail, 0)
	for _, v := range mails {
		if v.Filters.RegisterTime >= registerTime {
			res = append(res, v)
		}
	}
	return res
}

func filterMailsWithLanguage(contens map[string]string, language string) (map[string]string, error) {
	if len(contens) == 0 {
		return nil, fmt.Errorf("mail body is empty")
	}
	body := make(map[string]string)
	for k1, v1 := range contens {
		if k1 == language {
			body[k1] = v1
			break
		}
	}
	if len(body) == 0 {
		if contens["en"] != "" {
			body["en"] = contens["en"]
		} else {
			for k1, v1 := range contens {
				body[k1] = v1
				break
			}
		}
	}
	return body, nil
}
