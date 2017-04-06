## 识别原理与效果

### 原理

对图像的 HSV 空间模型进行贝叶斯分类，首先针对如下示例背景做数据处理与分析：

' | ' | '
--- | --- | ---
![](../source/01.jpg) | ![](../source/01.jpg) | ![](../source/01.jpg)

得到的数据做 H-V 和 H-S 散点图如下所示

H-V | H-S
--- | ---
![](HV.png) | ![](HS.png)

由于要求运算精度并不高以及处理速度的要求，得出的趋势线方程如下：

> 5 * V - 13 * H = 600
> 5 * S - 17 * H = 500

所以确定合适的 HSV 阈值，就可以很好地识别背景

### 效果

No. | SOURCE | RESULT
--- | ------ | ------
1   | ![](../source/02.jpg) | ![](../dist/02.jpg)
2   | ![](../source/03.jpg) | ![](../dist/03.jpg)
3   | ![](../source/04.jpg) | ![](../dist/04.jpg)
4   | ![](../source/05.jpg) | ![](../dist/05.jpg)
5   | ![](../source/06.jpg) | ![](../dist/06.jpg)
6   | ![](../source/07.jpg) | ![](../dist/07.jpg)
7   | ![](../source/08.jpg) | ![](../dist/08.jpg)
8   | ![](../source/09.jpg) | ![](../dist/09.jpg)
9   | ![](../source/10.jpg) | ![](../dist/10.jpg)
10   | ![](../source/11.jpg) | ![](../dist/11.jpg)
11   | ![](../source/12.jpg) | ![](../dist/12.jpg)
12   | ![](../source/13.jpg) | ![](../dist/13.jpg)

### 说明

可以看到标号为 10 的图片识别效果不太好，对于这种肉眼都看不出来的图，我只想说：

# 我能怎么办，我 TM 也很无奈啊