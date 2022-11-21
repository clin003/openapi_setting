编译与压缩，隐藏CMD框

go build -ldflags="-H windowsgui -w -s"


设置图标
找到一个 png 的图片，放在项目根目录


// 步骤一
// 安装fyne工具
go get fyne.io/fyne/cmd/fyne

// 步骤二
// 将静态资源编译为 go 文件
fyne bundle fav.png >> bundled.go
// 如果是要追加资源
// fyne bundle -append image2.jpg >> bundled.go

// 步骤三
// 打开 bundled.go 文件会看到一个变量 resourceFavPng
// 修改代码，给 app 设置图标
// 这样，窗口图标，任务栏图标就设置好了
a := app.New()
a.SetIcon(resourceFavPng)

// 最后打包，使用 fyne package 而不是 go build
// fyne package 设置了可执行文件的图标，并且隐藏了CMD框，但是它默认的编译没有添加压缩参数。
fyne package -os windows -icon fav.png

至此，可执行文件，窗口图标，任务栏图标都有了。遗憾的是没法添加压缩参数 -ldflags="-H windowsgui -w -s


CC=x86_64-w64-mingw32-gcc GOARCH=amd64 GOOS=windows CGO_ENABLED=1  fyne package -os windows -icon fav.png

GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags="-H windowsgui -w -s"

