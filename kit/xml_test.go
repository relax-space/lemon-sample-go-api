package kit

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func Test_Xml(t *testing.T) {
	school := School{
		Name: "Eland",
		Students: []Student{
			Student{
				Name: "xiao.xinmiao",
				Age:  18,
			},
			Student{
				Name: "big mao",
				Age:  18,
			},
		},
	}
	b, _ := xml.MarshalIndent(school, "", " ")
	fmt.Println(string(b))
}
