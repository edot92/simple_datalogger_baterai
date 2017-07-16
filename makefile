#simple build to all platform and copy conf :)
#for linux you must set permision to excecute
all: build

build: app

app:
	#env  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o ./build/darwin_amd64/appukon
	#cp -f -R  conf build/darwin_amd64 
	#cp -f -R  swagger build/darwin_amd64 
	#cp -f -R  static build/darwin_amd64 
	#env  CGO_ENABLED=0 GOOS=linux GOARCH=386 go build  -o ./build/linux_386/appukon
	#cp -f -R  conf build/darwin_amd64 
	#cp -f -R  swagger build/darwin_amd64 
	#cp -f -R  static build/darwin_amd64 
	#env  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o  ./build/linux_amd64/appukon
	#cp -f -R conf build/linux_amd64
	#cp -f -R swagger build/linux_amd64
	#cp -f -R static build/linux_amd64
	env  CGO_ENABLED=0 GOOS=linux CC=arm-linux-gnueabihf-gcc GOARCH=arm go build  -o ./build/linux_arm/appukon
	cp -f -R conf build/linux_arm
	cp -f -R swagger build/linux_arm
	#cp -f -R static build/linux_arm
	#env  CGO_ENABLED=0 GOOS=windows GOARCH=386 go build  -o ./build/windows_386/appukon.exe
	#cp -f -R conf build/windows_386
	#		cp -f -R swagger build/windows_386
	#cp -f -R static build/windows_386
	#env  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/windows_amd64/appukon.exe
	#cp -f -R conf build/windows_amd64
	#	cp -f -R conf build/windows_amd64
	#		cp -f -R swagger build/windows_amd64
	#cp -f -R static build/windows_amd64
