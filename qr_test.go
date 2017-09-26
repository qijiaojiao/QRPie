package qrpie

import (
	"fmt"
	_ "image/png"

	"image"
	"math/rand"
	"os"
	"testing"
)

var qr = NewQr("model.csv")

func TestIsSim(t *testing.T) {
	fmt.Println(isSim(45, 9*3))
}

func TestGenerateTrainData(t *testing.T) {
	GenerateTrainData("/Users/coyte/Downloads/erweima", "/Users/coyte/Downloads/img", "train_data.csv")
}

func TestIsQr(t *testing.T) {

}

func TestDownLoadImg(t *testing.T) {
	downLoadImg("http://www.yigeshaozi.com/qiniu/5599/image/fbc36af155da17d17d512ea25b4ca874.png?imageView2/2/w/399")
}

func TestIsQrUrl(t *testing.T) {
	fmt.Println(qr.IsQrUrl("outPG/Art/view_formula_2x.png"))
}

func TestIsQrPath(t *testing.T) {
	for {
		fmt.Println(rand.Intn(2))
	}

}

func TestFeatures(t *testing.T) {
	qr := NewQr("model.csv")
	file, _ := os.Open("view_formula_2x.png")
	img, _, _ := image.Decode(file)
	fmt.Println(qr.IsQr(img))
}

func TestGrayMean(t *testing.T) {
	file, _ := os.Open("view_formula_2x.png")
	img, _, _ := image.Decode(file)
	fmt.Println(grayMean(img))
}
