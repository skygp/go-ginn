package regex

import (
	cRegex "github.com/mingrammer/commonregex"
)

func GetDateWithString(text string) []string{
	return cRegex.Date(text)
}

func GetDateWithByte(byte []byte) []string{
	return cRegex.Date(string(byte))
}

func GetTimeWithString(text string) []string{
	return cRegex.Time(text)
}

func GetTimeWithByte(byte []byte) []string{
	return cRegex.Time(string(byte))
}

func GetLinksWithString(text string) []string{
	return cRegex.Links(text)
}

func GetLinksWithByte(byte []byte) []string{
	return cRegex.Links(string(byte))
}

func GetPhonesWithString(text string) []string{
	return cRegex.PhonesWithExts(text)
}

func GetPhonesWithByte(byte []byte) []string{
	return cRegex.PhonesWithExts(string(byte))
}

func GetEmailsWithString(text string) []string{
	return cRegex.Emails(text)
}

func GetEmailsWithByte(byte []byte) []string {
	return cRegex.Emails(string(byte))
}

func GetMD5HexesWithString(text string) []string{
	return cRegex.Emails(text)
}

func GetMD5HexesWithByte(byte []byte) []string {
	return cRegex.Emails(string(byte))
}

func GetIPsWithString(text string) []string{
	return cRegex.IPs(text)
}

func GetIPsWithByte(byte []byte) []string {
	return cRegex.IPs(string(byte))
}

func GetIPv4sWithString(text string) []string{
	return cRegex.IPv4s(text)
}

func GetIPv4sWithByte(byte []byte) []string {
	return cRegex.IPv4s(string(byte))
}

func GetIPv6sWithString(text string) []string{
	return cRegex.IPv6s(text)
}

func GetIPv6sWithByte(byte []byte) []string {
	return cRegex.IPv6s(string(byte))
}
