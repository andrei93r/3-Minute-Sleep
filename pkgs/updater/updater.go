package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type UpdaterContainer struct {
	wd                 string
	repoPartUri        string
	programName        string
	currentVersion     int
	lastCheckedVersion int
}

func NewUpdater(initializedVersion string) error {
	uc := new(UpdaterContainer)
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get the working directory for NewUpdater due to : %w", err)
	}

	uc.wd = wd

	uc.repoPartUri = "andrei93r/3-Minute-Sleep"
	uc.programName = "3-Minute-Sleep"
	cv, err := uc.parseVersionToInt(initializedVersion)
	if err != nil {
		return fmt.Errorf("current version of the app could not be set due in NewUpdater to %w", err)
	}
	uc.currentVersion = cv

	uc.repeatCheck()

	return err

}

func (uc *UpdaterContainer) parseVersionToInt(version string) (int, error) {

	versionArr := strings.Split(version, ".")
	major, err := strconv.Atoi(versionArr[0])
	if err != nil {
		return 0, fmt.Errorf("the major version could not be parsed in int on parseVersionFromInt")
	}
	minor, err := strconv.Atoi(versionArr[1])
	if err != nil {
		return 0, fmt.Errorf("the minor version could not be parsed in int on parseVersionFromInt")
	}
	patch, err := strconv.Atoi(versionArr[2])
	if err != nil {
		return 0, fmt.Errorf("the patch version could not be parsed in int on parseVersionFromInt")
	}

	var versionInt = 100000000000

	versionInt = versionInt + major*1000000
	versionInt = versionInt + minor*1000
	versionInt = versionInt + patch

	fmt.Println(versionInt, major, minor, patch)

	return versionInt, nil

}

func (uc *UpdaterContainer) checkForUpdates() (downloadLink string) {
	res, err := http.Get("https://api.github.com/repos/" + uc.repoPartUri + "/releases")
	defer res.Body.Close()
	if err != nil {
		return
	}

	var releases []struct {
		TagName string `json:"tag_name"`
		Assets  []struct {
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	err = json.NewDecoder(res.Body).Decode(&releases)
	if err != nil {
		return
	}

	var latestVersionIndex int
	var latestVersion int

	for i := 0; i < len(releases); i++ {
		versionInt, err := uc.parseVersionToInt(releases[i].TagName)
		// probably a bad idea
		if err != nil {
			continue
		}

		if versionInt > latestVersion {
			latestVersionIndex = i
			latestVersion = versionInt
		}

	}

	uc.lastCheckedVersion = latestVersion

	return releases[latestVersionIndex].Assets[0].BrowserDownloadURL
}

func (uc *UpdaterContainer) downloadUpdate(uri string) {
	res, err := http.Get(uri)
	defer res.Body.Close()
	if err != nil {
		return
	}

	diskLocation := uc.wd + "/" + uc.programName + ".exe"

	f, err := os.Create(diskLocation)
	defer f.Close()
	if err != nil {
		return
	}
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return
	}
}

func (uc *UpdaterContainer) update() {
	files, err := os.ReadDir(uc.wd)
	if err != nil {
		return
	}

	var _ []struct {
		modifiedDate time.Time
		path         string
	}

	for i := 0; i < len(files); i++ {

	}

}

func (uc *UpdaterContainer) repeatCheck() {
	for {
		downloadUrl := uc.checkForUpdates()

		fmt.Println("updated +%v", uc)

		fmt.Println("loop")

		if uc.lastCheckedVersion > uc.currentVersion {
			fmt.Println("file is updating")
			uc.downloadUpdate(downloadUrl)

		}

		time.Sleep(10 * time.Second)
	}
}
