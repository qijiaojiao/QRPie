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
	fmt.Println(IsQrUrl("https://imag00&sec=1505131357841&di=a9c45e9f308e14142eb19c99.jpg"))
}

func TestIsQrPath(t *testing.T) {
	fmt.Println(IsQrPath("1.06253348869.jpg"))
}