package phone

import (
	_ "embed"
	"encoding/json"
	"regexp"
	"strings"
)

type CountryInfo struct {
	EnName      string `json:"en_name"`
	CnName      string `json:"cn_name"`
	CountryCode string `json:"country_code"`
	PhoneCode   string `json:"phone_code"`
}

//go:embed phone_rule.json
var countryInfos string

var (
	MaxPhoneCodeLength = 4 // 区号最长为4位
	countryInfoList    []*CountryInfo
)

func init() {
	err := json.Unmarshal([]byte(countryInfos), &countryInfoList)
	if err != nil {
		panic(err)
	}
}

// clearText 清洗输入的手机号
func clearText(text string) string {
	// 替换空白字符
	regBlank := regexp.MustCompile(`\s`)
	text = string(regBlank.ReplaceAll([]byte(text), []byte("")))
	// 替换 非数字
	regNotNumbered := regexp.MustCompile(`[^0-9]`)
	text = string(regNotNumbered.ReplaceAll([]byte(text), []byte("")))
	// 替换掉前导 00
	if strings.HasPrefix(text, "00") {
		text = strings.Replace(text, "00", "", 1)
	}
	if len(text) > MaxPhoneCodeLength {
		text = text[0:MaxPhoneCodeLength]
	}

	return text
}

// getCountryInfo 通过区号获取国家信息
func getCountryInfo(text string) *CountryInfo {
	for _, countryInfo := range countryInfoList {
		if text == countryInfo.PhoneCode {
			return countryInfo
		}
	}
	return nil
}

// GetPhoneCode 通过手机号获取 区号
func GetPhoneCode(phone string) (phoneCode string) {
	text := clearText(phone)
	for i := len(text) - 1; i >= 0; i-- {
		countryInfo := getCountryInfo(text)
		if countryInfo != nil {
			phoneCode = countryInfo.PhoneCode
			break
		}
		text = text[:i]
	}
	return
}
