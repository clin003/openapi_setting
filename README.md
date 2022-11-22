openapi_setting 程序打包指南

	找到一个 png 的图片，放在项目根目录

###	编译步骤
```bash
// 步骤一
// 安装fyne工具
go get fyne.io/fyne/cmd/fyne

// 步骤二
// 将静态资源编译为 go 文件
fyne bundle fav.png >> bundled.go

// 最后打包，使用 fyne package 而不是 go build
// fyne package 设置了可执行文件的图标，并且隐藏了CMD框，但是它默认的编译没有添加压缩参数。
fyne package -os windows -icon fav.png
```

至此，可执行文件，窗口图标，任务栏图标都有了。

###	跨平台交叉编译
```bash
// 步骤一
// 安装fyne cross工具
go install gitee.com/lyhuilin/fyne_cross

// 最后打包，使用 cross windows，参考 _build.sh 中命令（依赖docker）
sudo cross windows -arch=*
```