package levenshtein

import "testing"

func TestSentence_Distance(t *testing.T) {
	a := NewSentence("圏央道　外回は友田ＴＮと青梅ＩＣの間で工事のため走行車線が規制されています")
	println(a.Distance("東北道　下りは安代ＩＣと鹿角八幡平ＩＣの間にある石通ＴＮ付近で、阿闍羅ＰＡと大鰐ＢＳの間で工事のため走行車線が規制されています"))
}
