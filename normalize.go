package pctenc

import (
	"unicode/utf8"
)

// Normalize normalizes the percent-encoded parts of a string,
//
// (Note that Normalize does no percent-encode anything that is not already percent-encoded,
// and does not percent-decode anything that is already percent-encoded.
// It just focuses on whether the hexadecimal encoding is using upper-case letters 'A','B','C','D','E','F' or not.)
//
// Percent-Encoding Normalization is important, for example, if you want to determine if two percent-encoded strings are equal or not.
//
// For example, the following string:
//
//	"⚡ Hello world! 😈"
//
// Can be percent-encoded in all these ways:
//
//	"%E2%9A%A1%20Hello%20world%21%20%F0%9F%98%88" // <--- normalized version
//	"%E2%9A%A1%20Hello%20world%21%20%F0%9f%98%88"
//	"%E2%9A%A1%20Hello%20world%21%20%f0%9F%98%88"
//	"%E2%9A%A1%20Hello%20world%21%20%f0%9f%98%88"
//	"%E2%9A%a1%20Hello%20world%21%20%F0%9F%98%88"
//	"%E2%9A%a1%20Hello%20world%21%20%F0%9f%98%88"
//	"%E2%9A%a1%20Hello%20world%21%20%f0%9F%98%88"
//	"%E2%9A%a1%20Hello%20world%21%20%f0%9f%98%88"
//	"%E2%9a%A1%20Hello%20world%21%20%F0%9F%98%88"
//	"%E2%9a%A1%20Hello%20world%21%20%F0%9f%98%88"
//	"%E2%9A%a1%20Hello%20world%21%20%f0%9F%98%88"
//	"%E2%9A%a1%20Hello%20world%21%20%f0%9f%98%88"
//	"%E2%9a%a1%20Hello%20world%21%20%F0%9F%98%88"
//	"%E2%9a%a1%20Hello%20world%21%20%F0%9f%98%88"
//	"%E2%9a%a1%20Hello%20world%21%20%f0%9F%98%88"
//	"%E2%9a%a1%20Hello%20world%21%20%f0%9f%98%88"
//	"%e2%9A%A1%20Hello%20world%21%20%F0%9F%98%88"
//	"%e2%9A%A1%20Hello%20world%21%20%F0%9f%98%88"
//	"%e2%9A%A1%20Hello%20world%21%20%f0%9F%98%88"
//	"%e2%9A%A1%20Hello%20world%21%20%f0%9f%98%88"
//	"%e2%9A%a1%20Hello%20world%21%20%F0%9F%98%88"
//	"%e2%9A%a1%20Hello%20world%21%20%F0%9f%98%88"
//	"%e2%9A%a1%20Hello%20world%21%20%f0%9F%98%88"
//	"%e2%9A%a1%20Hello%20world%21%20%f0%9f%98%88"
//	"%e2%9a%A1%20Hello%20world%21%20%F0%9F%98%88"
//	"%e2%9a%A1%20Hello%20world%21%20%F0%9f%98%88"
//	"%e2%9A%a1%20Hello%20world%21%20%f0%9F%98%88"
//	"%e2%9A%a1%20Hello%20world%21%20%f0%9f%98%88"
//	"%e2%9a%a1%20Hello%20world%21%20%F0%9F%98%88"
//	"%e2%9a%a1%20Hello%20world%21%20%F0%9f%98%88"
//	"%e2%9a%a1%20Hello%20world%21%20%f0%9F%98%88"
//	"%e2%9a%a1%20Hello%20world%21%20%f0%9f%98%88"
//
// Plus more!
//
// Percent-Encoding Normalization picks a single unique percent-encoding for a each string.
//
// Normalize does not validate the string.
// It just applies the normalization algorithm to the string.
func Normalize(percentEncoded string) string {
	var buffer [128]byte
	var p []byte = buffer[0:0]

	var prev0 rune
	var prev1 rune

	for _, r := range percentEncoded {
		switch {
		case '%' == prev0 || '%' == prev1:
			r = upperHexDig(r)
		}
		p = utf8.AppendRune(p, r)

		prev1 = prev0
		prev0 = r
	}

	return string(p)
}

func upperHexDig(r rune) rune {
	switch {
	case 'a' <= r && r <= 'f':
		return r + ('A' - 'a')
	default:
		return r
	}
}
