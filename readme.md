# 说明

---

### 编程环境

- Ubuntu 20.04.5
- go 19.3
- fyne v2.2.4
- gcc 9.4.0

该项目用 ***fyne*** 构建的MarkDown，为一个简单的go语言练手项目，如有问题  
欢迎给我发送 <gzy_njust@njust.edu.cn> ,欢迎一起学习交流


##  使用须知 

---

### Windows
> 下载MyMarkDown.exe，直接可以用  

### Ubuntu
> 1. 先解压：`tar -zxvf  MyMarkDown.tar.xz`
> 2. 再运行：`./usr/local/bin/FileApp` 

## 安装说明

  如果你计划安装并学习 ***fyne*** ，下面我给你一个初略的教程

---

>- 首先确保你的电脑里有正确的gcc环境，如果没有的话，请自行搜索。 
  当你在命令行中键入`gcc --version`，说明你这步完成（ ~~ 不保证正确，
  在Windows下我试了许多方法安装gcc，gcc成功，但是运行fyne依然失败~~）
>- 创建一个项目（~~ 就是创建一个文件夹 ~~）命令行输入`go mod init`
>- 再输入`go get fyne.io/fyne/v2`
>- 创建一个go文件夹，记住要package main，之后依次输入`go mod tidy`和`go run .`即可

###### ***可以直接运行本项目！***

##### 如果你想要将你的fyne项目打包成可执行文件

>- `go get fyne.io/fyne/cmd/fyne`  
>- `fyne package -appVersion 1.0.0 -name MyMarkDown -release`  

 如果上述运行后，未打包成功，我这里给出两个我自己遇到的问题
1. 首先如果提示fyne不是命令，可以直接到`{GOPATH}/pkg/mod/fyne.io/fyne/cmd/fyne`或者
   直接`git clone fyne.io/fyne/cmd/fyne`，到fyne文件夹下，运行`get install`，如果
   `{GOPATH}/bin`路径下有 ***fyne.exe*** 重新打开命令行运行。
2. 如果提示缺少相应的参数，请根据具体提示添加
