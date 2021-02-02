package protoc

import (
	"errors"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli"
	"regexp"
	"strconv"
	"strings"
)

func Go2protoc(c *cli.Context) (err error) {
	j, err := clipboard.ReadAll()
	if err != nil {
		return
	}

	r, err := go2pb(j)
	if err != nil {
		return
	}
	return clipboard.WriteAll(r)
}

func go2pb(in string) (out string, err error) {
	result := strings.Split(in, "\n")

	typesMap := map[string]string{
		"int":               "int32",
		"int8":              "int32",
		"map[string]string": "map<string, string>",
	}

	message := "message "
	isEmpty := true

	for i, vs := range result {
		if vs == "" {
			continue
		}
		// 去掉多余空格
		r, _ := regexp.Compile(`\s+`)
		vs = r.ReplaceAllString(vs, " ")

		vs = strings.Trim(vs, " ")
		v := strings.Split(vs, " ")
		if len(v) < 2 {
			continue
		}
		// 跳过注释
		if v[0][0] == '/' {
			continue
		}

		if v[0] == "type" {
			message += v[1] + " {\n"
			isEmpty = false
			continue
		}
		fieldName := string(tuoFeng2SheXing([]byte(v[0])))
		types := v[1]

		if tm, ok := typesMap[types]; ok {
			types = tm
		}
		message += "    "

		// 判断数组
		if len(types) >= 2 && types[:2] == "[]" {
			types = types[2:]
			message += "repeated "
		}
		types = strings.Trim(types, "*")
		message += types + " " + fieldName + " = " + strconv.Itoa(i) + ";\n"
	}

	message += "}\n"

	if isEmpty {
		err = errors.New("empty input")
	} else {
		err = nil
	}
	return message, err
}

func tuoFeng2SheXing(src []byte) (out []byte) {
	l := len(src)
	out = []byte{}
	for i := 0; i < l; i = i + 1 {
		// 大写变小写
		if 97-32 <= src[i] && src[i] <= 122-32 {
			if i != 0 {
				out = append(out, 95)
			}
			out = append(out, src[i]+32)
		} else {
			out = append(out, src[i])
		}
	}
	return
}
