package utilities

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"

	http "github.com/bogdanfinn/fhttp"
)

func SetHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func ExtractPtrFromHref(href string) (string, error) {
	// Look for href="/search/view/ptr/<UUID>/"
	re := regexp.MustCompile(`href="/search/view/ptr/([a-zA-Z0-9\-]+)/"`)
	matches := re.FindStringSubmatch(href)
	if len(matches) != 2 {
		return "", fmt.Errorf("expected 2 matches, got %d", len(matches))
	}
	return matches[1], nil
}

func ExtractMiddlewareToken(html string) (string, error) {
	re := regexp.MustCompile(`<input type="hidden" name="csrfmiddlewaretoken" value="(.*)">`)
	matches := re.FindStringSubmatch(html)
	if len(matches) != 2 {
		return "", fmt.Errorf("expected 2 matches, got %d", len(matches))
	}
	return matches[1], nil
}

func URLencode(qp interface{}) (string, error) {
	v := url.Values{}
	reflectType := reflect.TypeOf(qp)
	reflectValue := reflect.ValueOf(qp)

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		tag := field.Tag.Get("json")
		value := reflectValue.Field(i).Interface()

		switch reflect.TypeOf(value).Kind() {
		case reflect.Int:
			v.Set(tag, strconv.Itoa(value.(int)))
		case reflect.Bool:
			v.Set(tag, strconv.FormatBool(value.(bool)))
		case reflect.Slice:
			slice := reflect.ValueOf(value)
			for j := 0; j < slice.Len(); j++ {
				elem := slice.Index(j)
				if elem.Kind() == reflect.Int {
					v.Add(tag, strconv.Itoa(int(elem.Int())))
				}
			}
		case reflect.String:
			v.Set(tag, value.(string))
		default:
			return "", fmt.Errorf("unhandled type: %s", reflect.TypeOf(value).Kind())
		}
	}

	return v.Encode(), nil
}
