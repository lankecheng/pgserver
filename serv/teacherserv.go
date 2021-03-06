package serv

import (
	"github.com/lankecheng/pgserver/dao"
	"github.com/lankecheng/pgserver/pgpub"
)

func ShowTeachers() (teachers []map[string]interface{}, err error) {
	struFlds := []string{"Uid", "Uname", "Gender", "Occup", "Avatar", "Audio", "Country", "Description", "Interest"}
	users, err := dao.QueryUsersByType(pgpub.TypeTeacher, struFlds...)
	if err != nil {
		return
	}

	for _, user := range users {
		teachers = append(teachers, pgpub.ConvertStruct2Map(&user, struFlds...))
	}

	return
}