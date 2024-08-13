package main

//Note: パッケージをimportするために以下の実行が必要
// go mod init exerciese-slices
// go get golang.org/x/tour/pic
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	img := make([][]uint8, dy)

	for idx, _ := range img {
		img[idx] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img[y][x] = uint8(x+y) / uint8(2)
		}
	}
	return img
}

func main() {
	// Note: ターミナルからだと画像が表示されないので、
	pic.Show(Pic)
}
