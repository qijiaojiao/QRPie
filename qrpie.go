package qrpie

import (
	"encoding/csv"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"

	"code.google.com/p/graphics-go/graphics"

	"truxing/commons/log"
)

const (
	imgWidth  = 30
	vecLen    = imgWidth*imgWidth + 2 //最后两个是提取的特征，前边900个是像素点
	Threshold = 0.2
)

var (
	once  sync.Once
	model [][]string
)

func initEnv() {
	once.Do(func() {
		cs, _ := os.Open("model.csv")
		reader := csv.NewReader(cs)
		model, _ = reader.ReadAll()
	})
}

func loadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

func extractFeature(img image.Image) []float64 {
	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y
	f1 := 0
	f2 := 0
	cc := true
	x := 0
	y := 0
	features := make([]float64, 0, vecLen)
	grayImg := image.NewGray(img.Bounds())
	for h := 0; h < height; h++ {
		list := make([]int, 5, 5)
		p := 0
		for w := 0; w < width; w++ {
			c := color.GrayModel.Convert(img.At(w, h))
			if c.(color.Gray).Y < 127 {
				if cc {
					p = (p + 1) % 5
					if isDemandBiLy(list, p) {
						if x == w || y == h {
							f2++
						}
						x = w
						y = h
						f1++
					}
					list[p] = 0
				}
				cc = false
				list[p] = list[p] + 1
				grayImg.SetGray(w, h, color.Gray{Y: 0})

			} else {
				if !cc {
					p = (p + 1) % 5
					if isDemandBiLy(list, p) {
						if x == w || y == h {
							f2++
						}
						x = w
						y = h
						f1++
					}
					list[p] = 0
				}
				cc = true
				list[p] = list[p] + 1
				grayImg.SetGray(w, h, color.Gray{Y: 255})
			}
		}
	}

	sImg, err := scale(grayImg)
	if err != nil {
		log.Errorf("scale image fail:%s", err.Error())
	}
	sum := 0
	for i := 0; i < sImg.Bounds().Dy(); i++ {
		for j := 0; j < sImg.Bounds().Dx(); j++ {
			c := color.GrayModel.Convert(sImg.At(j, i))
			if c.(color.Gray).Y == 0 {
				features = append(features, 0)
			} else {
				features = append(features, 1)
				sum++
			}

		}
	}
	features = append(features, float64(f1)/math.Log(float64(height)))
	features = append(features, float64(f2)/math.Log(float64(height)))
	return features
}

func isDemandBiLy(list []int, p int) bool {
	p = p % 5
	if isSim(list[(p+1)%5], list[p]) && isSim(list[(p+1)%5], list[(p+3)%5]) && isSim(list[(p+1)%5], list[(p+4)%5]) && isSim(list[(p+2)%5], list[(p+3)%5]*3) {
		return true
	} else {
		return false
	}
}

func isSim(x int, y int) bool {
	if y == 0 {
		return false
	}
	if math.Abs((float64(x)/float64(y) - 1)) < 0.3 {
		return true
	} else {
		return false
	}
}

func scale(img image.Image) (d image.Image, err error) {
	dst := image.NewGray(image.Rect(0, 0, imgWidth, imgWidth))
	err = graphics.Scale(dst, img)
	return dst, err
}

//此方法是用来产生训练数据的
//qrpath:放二维码图片的文件夹地址
//other: 放非二维码图片的文件夹地址
//name : 产生的训练数据文件的文件名
func GenerateTrainData(qrPath string, other string, name string) (err error) {
	dirs := []string{qrPath, other}
	file, _ := os.Create(name)
	writer := csv.NewWriter(file)
	header := make([]string, 0, vecLen)
	header = append(header, "/")
	for i := 0; i < imgWidth*imgWidth; i++ {
		header = append(header, "pix."+strconv.Itoa(i))
	}
	header = append(header, "f1")
	header = append(header, "f2")
	header = append(header, "y")
	writer.Write(header)
	fail := 0

	for i, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			} else {
				img, err := loadImage(dir + "/" + file.Name())
				if err != nil {
					log.Debugf("load img fail error msg is %s,fileName is %s", err.Error(), file.Name())
					fail++
					continue
				}
				features := extractFeature(img)
				record := make([]string, 0, vecLen+1)
				record = append(record, file.Name())
				for _, s := range features {
					record = append(record, strconv.FormatFloat(s, 'f', -1, 64))
				}
				if i == 0 {
					record = append(record, strconv.Itoa(1))
				} else {
					record = append(record, strconv.Itoa(0))
				}
				writer.Write(record)
			}
		}
	}

	writer.Flush()
	file.Close()
	return
}

func predict(features []float64) bool {
	initEnv()
	fm := make(map[string]float64)
	for i := 0; i < imgWidth*imgWidth; i++ {
		key := "pix." + strconv.Itoa(i)
		fm[key] = features[i]
	}
	fm["f1"] = features[900]
	fm["f2"] = features[901]

	i := 0
	var gain float64
	nextNode := "0-0"
	for _, record := range model {
		if i == 0 {
			i = 1
			continue
		}
		if nextNode != record[2] {
			continue
		}
		if record[3] != "Leaf" {
			split, _ := strconv.ParseFloat(record[4], 64)
			if fm[record[3]] < split {
				nextNode = record[5]
			} else {
				nextNode = record[6]
			}

		} else {
			g, _ := strconv.ParseFloat(record[8], 64)
			gain += g
		}
	}
	if math.Exp(gain)/(1+math.Exp(gain)) > Threshold {
		return true
	}
	return false
}

func IsQr(img image.Image) (bool, error) {
	features := extractFeature(img)
	return predict(features), nil
}

func downLoadImg(url string) (image.Image, error) {
	response, e := http.Get(url)
	if e != nil {
		log.Errorf("download image error:%s", e.Error())
		return nil, e
	}
	defer response.Body.Close()
	img, _, err := image.Decode(response.Body)
	if err != nil {
		log.Debug(err.Error())
	}
	return img, err
}

func IsQrUrl(url string) (bool, error) {
	img, err := downLoadImg(url)
	if err == nil {
		return IsQr(img)
	} else {
		return false, err
	}
}

func IsQrPath(path string) (bool, error) {
	img, err := loadImage(path)
	if err == nil {
		return IsQr(img)
	} else {
		return false, err
	}
}
