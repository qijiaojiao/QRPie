package qrpie

import (
    "testing"
    "fmt"
    "encoding/csv"
    _"image/png"
    "os"
    "image"
)

func TestProcessImg(t *testing.T)  {
    processImg("1504662961.png")
}

func TestS(t *testing.T)  {
    _ ,err :=loadImage("/Users/coyte/Downloads/erweima/u=1360432999,376837849&fm=27&gp=0.jpg")
    if err != nil{
        fmt.Println(err.Error())
    }
}
func TestIsSim(t *testing.T)  {
    fmt.Println(isSim(45,9*3))
}

func TestGenerateTrainData(t *testing.T)  {
    generateTrainData("/Users/coyte/Downloads/erweima","/Users/coyte/Downloads/img","train_data.csv")
}

func TestCsv(t *testing.T){
    mm,_ := os.Create("aa.csv")
    writer := csv.NewWriter(mm)
    writer.Write([]string{"a","b","c","d"})
    writer.Flush()
}

func TestCSVRead(t *testing.T)  {
    model,_ := os.Open("model.csv")
    reader := csv.NewReader(model)
    for{
        record,err := reader.Read()
        if err != nil{
            break
        }
        fmt.Println(record)
    }
}

func TestIsQr(t *testing.T) {
    file,_ := os.Open("1504662961.png")
    img,_,_ := image.Decode(file)
    fmt.Println(IsQr(img))
    file,_ = os.Open("1.06253348869.jpg")
    img,_,_ = image.Decode(file)
    fmt.Println(IsQr(img))
    file.Close()
}