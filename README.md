此项目通过决策树算法可以识别一张图片是否二维码

使用说明：

import github.com/Mrfogg/QRPie/qrpie

func main(){

	file,_ := os.Open("1504662961.png")

	img,_,_ := image.Decode(file)

	fmt.Println(qrpie.IsQr(img)) //IsQr返回是否为二维码，是的话返回true

}
