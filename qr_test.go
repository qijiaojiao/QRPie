package qrpie

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"testing"
)

func TestIsSim(t *testing.T) {
	fmt.Println(isSim(45, 9*3))
}

func TestGenerateTrainData(t *testing.T) {
	GenerateTrainData("/Users/coyte/Downloads/erweima", "/Users/coyte/Downloads/img", "train_data.csv")
}

func TestIsQr(t *testing.T) {
	file, _ := os.Open("1504662961.png")
	img, _, _ := image.Decode(file)
	fmt.Println(IsQr(img))
	file, _ = os.Open("1.06253348869.jpg")
	img, _, _ = image.Decode(file)
	fmt.Println(IsQr(img))
	file.Close()
}

func TestDownLoadImg(t *testing.T) {
	downLoadImg("http://www.yigeshaozi.com/qiniu/5599/image/fbc36af155da17d17d512ea25b4ca874.png?imageView2/2/w/399")
}

func TestIsQrUrl(t *testing.T) {
	fmt.Println(IsQrUrl("http://wx3.sinaimg.cn/large/7ae2aa57ly1fiwde73v8kg20dw0af7wh.gif"))
}

func TestIsQrPath(t *testing.T) {
	fmt.Println(IsQrPath("1.06253348869.jpg"))
}