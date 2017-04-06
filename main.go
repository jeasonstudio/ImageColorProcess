package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func uint32Sqrt(n uint32) uint8 {
	res := int(n)
	bNum := res
	sNum := 0
	for {
		thisNum := (bNum + sNum) / 2
		ii := thisNum * thisNum
		ipi := (thisNum + 1) * (thisNum + 1)
		isi := (thisNum - 1) * (thisNum - 1)
		if ii == res || (ii < res && ipi >= res) || (ii > res && isi <= res) {
			return uint8(thisNum)
		} else if ii > res {
			bNum = thisNum
		} else if ii < res {
			sNum = thisNum
		}
	}
}

func uint16Sqrt(n uint16) uint {
	res := int(n)
	bNum := res
	sNum := 0
	for {
		thisNum := (bNum + sNum) / 2
		ii := thisNum * thisNum
		ipi := (thisNum + 1) * (thisNum + 1)
		isi := (thisNum - 1) * (thisNum - 1)
		if ii == res || (ii < res && ipi >= res) || (ii > res && isi <= res) {
			return uint(thisNum)
		} else if ii > res {
			bNum = thisNum
		} else if ii < res {
			sNum = thisNum
		}
	}
}

// func uint32ToUint8(n uint32) uint {
// 	return uint16Sqrt(uint32Sqrt(n))
// }

func main() {
	// fmt.Println(uint16Sqrt(uint32Sqrt(math.MaxUint32)))

	// file, err := os.Open("./colors/red.jpg")
	// file, err := os.Open("./colors/green.jpg")
	file, err := os.Open("./moshiimg/source/11.jpg")
	// file, err := os.Open("./flowers.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Create("11.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	img, _ := jpeg.Decode(file)
	xWidth := img.Bounds().Dx()
	yHeight := img.Bounds().Dy()

	jpg := image.NewRGBA64(img.Bounds())
	for i := 0; i < xWidth; i++ {
		for j := 0; j < yHeight; j++ {
			H, S, V := RGBAToHSV(img.At(i, j))
			thisR, thisG, thisB, thisA := img.At(i, j).RGBA()
			if V <= 50 || H >= 50 || (5*int(V)-13*int(H) >= 600) || (5*int(S)-17*int(H) >= 500) {
				fmt.Println(H, S, V)
				n := &color.RGBA64{uint16(thisR), uint16(thisG), uint16(thisB), uint16(thisA)}
				// n := &color.RGBA64{65535, 65535, 65535, 65535}
				jpg.SetRGBA64(i, j, *n)
			} else {
				m := RGBAToGray(img.At(i, j))
				n := &color.RGBA64{m, m, m, m}
				jpg.SetRGBA64(i, j, *n)
			}

			// fmt.Println(string(H) + "," + string(S) + "," + string(V))
			// fmt.Printf("%d,%d,%d\n", H, S, V)
		}
	}
	// draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)
}

func myAbs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

// golang max int
func getMax(first uint8, args ...uint8) uint8 {
	for _, v := range args {
		if first < v {
			first = v
		}
	}
	return first
}

func getMin(first uint8, args ...uint8) uint8 {
	for _, v := range args {
		if first > v {
			first = v
		}
	}
	return first
}

// RGBAToGray change RGB to Gray
func RGBAToGray(color color.Color) uint16 {
	thisR, thisG, thisB, _ := color.RGBA()
	return uint16((thisR*299 + thisG*587 + thisB*114 + 500) / 1000)
}

// RGBAToHSV H:色相, S:饱和度, V:明度(亮度)
func RGBAToHSV(c color.Color) (int, uint8, uint8) {
	// c input type color.Color
	// H(int) output range = 0 ~ 360
	// S and V output range = 0 ~ 255
	thisR, thisG, thisB, _ := c.RGBA()
	R := uint32Sqrt(thisR)
	G := uint32Sqrt(thisG)
	B := uint32Sqrt(thisB)

	max := getMax(R, G, B)
	min := getMin(R, G, B)

	// V := uint8((uint16(max) + uint16(min)) / 2)
	var H, S int

	if max == min {
		H = 0
	} else if max == R && G >= B {
		H = 60 * int(G-B) / int(max-min)
	} else if max == R && G < B {
		H = 60*(int(G-B)/int(max-min)) + 360
	} else if max == G {
		H = 60*(int(B-R)/int(max-min)) + 120
	} else if max == B {
		H = 60*(int(R-G)/int(max-min)) + 240
	}

	for {
		if H >= 0 && H <= 360 {
			break
		} else if H < 0 {
			H = H + 360
		} else if H > 360 {
			H = H - 360
		}
	}

	// L := int(V)
	derta := int(max-min) * 255

	if max == 0 {
		S = 0
	} else {
		S = int(derta / int(max))
	}
	return H, uint8(S), max
}
