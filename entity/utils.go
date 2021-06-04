package entity

import "fmt"

func SetMediaUrl(s string) string {
	url := fmt.Sprintf("%v/%v", "http://159.89.29.83/media", s)
	return url
}
