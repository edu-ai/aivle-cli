package operations

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/xuri/excelize/v2"
	"net/http"
)

func UploadWhitelist(apiRoot string, token string) {
	client := &http.Client{}
	// read the student list Excel file
	var fileName string
	err := survey.AskOne(&survey.Input{Message: "LumiNUS student list file location (.xlsx)"}, &fileName)
	if err != nil {
		panic(err)
	}
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	rows, err := f.GetRows("Results")
	if err != nil {
		panic(err)
	}
	var emailList []string
	for i, row := range rows {
		if i <= 2 {
			continue
		}
		emailList = append(emailList, row[1])
	}
	// select one of the courses
	selectedCourse, err := selectOneCourse(client, apiRoot, token)
	if err != nil {
		panic(err)
	}
	// upload the whitelist
	for _, email := range emailList {
		err = addWhitelist(client, apiRoot, token, selectedCourse.Id, email)
		if err != nil {
			panic(err)
		}
	}
}
