package actioninfo

import "fmt"

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		dataString, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(dataString)
	}
}
