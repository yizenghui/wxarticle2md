package wxarticle2md

import (
	"testing"
)

func TestToArticle(t *testing.T) {

	// url := `https://mp.weixin.qq.com/s/NiZ5iszTKEo2dxYo8mbRZg`
	// url := `https://mp.weixin.qq.com/s?__biz=MzIzMzA3NDg0Ng==&amp;mid=2649598185&amp;idx=1&amp;sn=74008c983b57acbb1671c3d9552ab7fc&amp;chksm=f0923118c7e5b80ef4949451f91bfa5a1758043575453466a322cfcb8bb0dc8b81feffe44706&scene=27#wechat_redirect`

	url := `https://mp.weixin.qq.com/s?__biz=MzA5OTA0NDIyMQ==&mid=2653896336&idx=4&sn=94964f5f65fbf978e700fbf5b29507e7&chksm=8b53c082bc24499421e546430587c878ca527116468f9248c7aaf7ce88e45376b2d8bbbd8f37`

	h, _ := GetHtml(url)
	a, _ := ToAticle(h)
	// t.Fatal(a)
	mk := Convert(a)
	t.Fatal(mk)
}
