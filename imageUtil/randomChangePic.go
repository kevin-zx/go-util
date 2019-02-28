package imageUtil

import (
	"github.com/disintegration/gift"
	"github.com/kevin-zx/go-util/randomUtil"
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
)

func LoadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func SaveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}

// random Resize 都是3:2的 图片
var widthSelects = []int{
	2048,
	1920,
	1632,
	1280,
	1238,
	1098,
	1024,
	977,
	948,
	800,
	798,
	768,
	670,
	600,
	504,
	450,
	425,
}

const initRandomRate = 0.3

func RandomResize(img image.Image) image.Image {
	//获取初始宽高
	originWith := img.Bounds().Size().X
	originHeight := img.Bounds().Size().Y
	rateDecr := initRandomRate/float64(len(widthSelects)) - 0.01
	resizeWidth := 0
	resizeHeight := 0
	var filter gift.Filter
	var selectFlag = false
	for !selectFlag {
		for i, ws := range widthSelects {
			if ws <= originWith && rand.Float64() < initRandomRate-(rateDecr*float64(i)) {
				resizeWidth = ws
				height := resizeWidth * 2 / 3
				if originHeight > height {
					resizeHeight = height
					filter = gift.Resize(resizeWidth, resizeHeight, gift.LanczosResampling)
				} else {
					resizeHeight = originHeight
					resizeWidth = resizeHeight * 3 / 2
					filter = gift.CropToSize(resizeWidth, resizeHeight, gift.LeftAnchor)
				}
				selectFlag = true
				break
			}
		}
	}
	img = useFilter(filter, img)
	return img
}

// random filter 随机加滤镜

func RandomFilter(img image.Image) image.Image {
	if checkRandom(0.8) {
		img = randomBrightness(img)
	}
	//对比度
	if checkRandom(0.5) {
		img = randomContrast(img)
	}
	//饱和度
	if checkRandom(0.6) {
		img = randomSaturation(img)
	}
	//Gamma值
	if checkRandom(0.3) {
		img = randomGamma(img)
	}
	// 高斯模糊
	if checkRandom(0.3) {
		img = randomGaussianBlur(img)
	}
	//反锐化处理
	if checkRandom(0.3) {
		img = randomUnsharpMask(img)
	}
	// 也是饱和度处理
	if checkRandom(0.2) {
		img = randomSigmoid(img)
	}

	//// 像素化
	//
	//if checkRandom(0.1) {
	//	img = randomPixelate(img)
	//}
	return img
}

func randomBrightness(img image.Image) image.Image {
	rb, _ := randomUtil.GetRandomInt(-30, 30)
	img = useFilter(gift.Brightness(float32(rb)), img)
	return img
}

func randomContrast(img image.Image) image.Image {
	ri, _ := randomUtil.GetRandomInt(-30, 30)
	img = useFilter(gift.Contrast(float32(ri)), img)
	return img
}

func randomSaturation(img image.Image) image.Image {
	ri, _ := randomUtil.GetRandomInt(-30, 30)
	img = useFilter(gift.Saturation(float32(ri)), img)
	return img
}

func randomGamma(img image.Image) image.Image {
	ri, _ := randomUtil.GetRandomInt(10, 15)
	img = useFilter(gift.Gamma(float32(ri)/10.00), img)
	return img
}

func randomGaussianBlur(img image.Image) image.Image {
	ri, _ := randomUtil.GetRandomInt(0, 15)
	img = useFilter(gift.GaussianBlur(float32(ri)/10.00), img)
	return img
}

func randomUnsharpMask(img image.Image) image.Image {
	sigma, _ := randomUtil.GetRandomInt(0, 10)
	amount, _ := randomUtil.GetRandomInt(5, 15)
	f := gift.UnsharpMask(float32(sigma)/10.00, float32(amount)/10.00, 0)
	img = useFilter(f, img)
	return img
}
func randomSigmoid(img image.Image) image.Image {
	factor, _ := randomUtil.GetRandomInt(-10, 10)
	f := gift.Sigmoid(0.5, float32(factor))
	img = useFilter(f, img)
	return img
}
func randomPixelate(img image.Image) image.Image {
	f := gift.Pixelate(2)
	img = useFilter(f, img)
	return img
}

func checkRandom(f float64) bool {
	return rand.Float64() < f
}

func useFilter(filter gift.Filter, img image.Image) image.Image {
	g := gift.New(filter)
	dst := image.NewNRGBA(g.Bounds(img.Bounds()))
	g.Draw(dst, img)
	img = dst
	return img
}
