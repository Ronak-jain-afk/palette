package export

import clipboard "github.com/atotto/clipboard"

func CopyToClipboard(text string) error {
	return clipboard.WriteAll(text)
}
