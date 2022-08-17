package lib

import (
	"os/exec"
	"testing"
)

func TestFilter(t *testing.T) {
	var sampleFolderData []string = []string{
		"gcs://sample-bucket/config/Code/logs",
		"gcs://sample-bucket/config/Code/logs/20220816T195842",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/exthost1",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/exthost1/output_logging_20220816T195848",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/exthost1/output_logging_20220816T200336",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/output_1_20220816T195844",
		"gcs://sample-bucket/config/Code/logs/20220816T195842/output_1_20220816T200333",
		"gcs://sample-bucket/config/Code/logs/20220525T201131",
		"gcs://sample-bucket/config/Code/logs/20220525T201131/exthost1",
		"gcs://sample-bucket/config/Code/logs/20220525T201131/exthost1/output_logging_20220525T201139",
		"gcs://sample-bucket/config/Code/logs/20220525T201131/output_1_20220525T201135",
		"gcs://sample-bucket/config/Code/logs/20220514T230127",
		"gcs://sample-bucket/config/Code/logs/20220514T230127/exthost1",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/c8a89dc5a4dcee8d4646c4b8474ae324ddaac427/c66e042f-72ec-4292-b6a8-cccbea394816/index-dir",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/98e3acaec81ab4b1d4493535e1030a26de1af26a",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/98e3acaec81ab4b1d4493535e1030a26de1af26a/70e7ca36-8902-4d9f-b6cb-b380bf578efc",
		"gcs://sample-bucket/config/Code/Service Worker/CacheStorage/98e3acaec81ab4b1d4493535e1030a26de1af26a/70e7ca36-8902-4d9f-b6cb-b380bf578efc/index-dir",
		"gcs://sample-bucket/config/Code/Backups",
		"gcs://sample-bucket/config/Code/Backups/1660697528120",
		"gcs://sample-bucket/config/Code/Backups/1660697528120/userdata",
		"gcs://sample-bucket/config/Code/Backups/1660697528120/file",
		"gcs://sample-bucket/config/Code/CachedData/c3511e6c69bb39013c4a4b7b9566ec1ca73fc4d5/chrome/wasm/index-dir",
		"gcs://sample-bucket/config/Code/CachedData/c3511e6c69bb39013c4a4b7b9566ec1ca73fc4d5/chrome/js",
		"gcs://sample-bucket/config/Code/CachedData/c3511e6c69bb39013c4a4b7b9566ec1ca73fc4d5/chrome/js/index-dir",
	}

	var standardGcsFolderRegex string = "gcs:.+"
	var folderToFind string = "exthost1"
	var completeFolderRegex string = standardGcsFolderRegex + folderToFind

	var filteredFolders []string = Filter(completeFolderRegex, sampleFolderData)

	if len(filteredFolders) != 6 {
		t.Errorf("Filter did not find files, got %d, want %d", len(filteredFolders), 6)
	}

	var unhappyFolder string = "xxxxx"
	var unhappyCompleteFolderRegex string = standardGcsFolderRegex + unhappyFolder

	var unhappyFilteredFolders []string = Filter(unhappyCompleteFolderRegex, sampleFolderData)

	if len(unhappyFilteredFolders) != 0 {
		t.Errorf("Filter did not filter out unwanted files, got %d, want %d", len(unhappyFilteredFolders), 0)
	}
}

func TestParseCmdOutput(t *testing.T) {
	cmd, _ := exec.Command("ls", "/").Output()

	var result []string = ParseCmdOutput(cmd)

	if result[0] != "bin" || result[5] != "home" {
		t.Errorf("ParseCmdOutput did not correctly format command output, got %s and %s, want bin and home", result[0], result[5])
	}

}
