此项目通过决策树算法可以识别一张图片是否二维码

使用说明：将model.csv放到你自己某个目录中

import github.com/Mrfogg/QRPie/qrpie

func main(){

  var qr = NewQr("model.csv的地址")

  fmt.Println(qr.IsQrUrl("你的图片url"))

}

预测二维码正确率几乎100%，

召回率97%左右