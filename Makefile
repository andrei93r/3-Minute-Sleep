Version := 0.0.4
VersionI := 401


Program-Name := "3-Minute-Sleep"
Exec-Path := "./bin/3-Minute-Sleep.exe"


pre:
	gh release create $(Version) $(Exec-Path) --title "Release $(Version)" --prerelease --generate-notes

deps:
	go mod tidy
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest

win:
	set GOOS=windows
	set GOARCH=amd64
	@echo "Building windows version $(Version)"
	goversioninfo
	cd
	go  build -ldflags "-X 'main.Version=9.0.1' " -o $(Exec-Path)
	del resource.syso

test:
	$(Exec-Path)

last:
	#go  build -ldflags "-s -w -H windowsgui" -o $(Exec-Path)
