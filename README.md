此项目通过决策树算法可以识别一张图片是否二维码

使用说明：

import github.com/Mrfogg/QRPie/qrpie

func main(){
  var qr = NewQr("model.csv")
  fmt.Println(qr.IsQrUrl("你的图片url"))
}
