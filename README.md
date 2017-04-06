# ImageColorProcess
图像单色处理，使用语言：Golang

## 如下效果

No. | localImage | afterProcess
---|---|---
1 | ![](pics/f2.jpg) | ![](pics/f2t.jpg)
2 | ![](pics/car.jpg) | ![](pics/cart.jpg)
3 | ![](pics/nba.jpg) | ![](pics/nbat.jpg)
4 | ![](pics/jingwu.jpg) | ![](pics/jingwut.jpg)

## Tips

使用了 HSV 空间模式

## Main

> V <= 50 || H >= 50 || (5*int(V)-13*int(H) >= 600) || (5*int(S)-17*int(H) >= 500)