package utils


import (
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "io/ioutil"
	"io"
	"os"
	"os/exec"
	"net/http"
	"crypto/tls"
	"archive/tar"
	"compress/gzip"
	"regexp"
	"strconv"
	"bufio"
	"errors"
	"time"
	"strings"
	"crypto/md5"
	"fmt"
	"crypto/rand"
	"encoding/hex"
)

func Generate()(uuid string)  {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logs.Info(err)
	}
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x",b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

//Read main.conf and return a map data
func GetConf(loadData map[string]map[string]string)(loadDataReturn map[string]map[string]string, err error) { 
    confFilePath := "conf/main.conf"
    jsonPathBpf, err := ioutil.ReadFile(confFilePath)
    if err != nil {
        logs.Error("utils/GetConf -> can't open Conf file -> " + confFilePath)
        return nil, err
    }

    var anode map[string]map[string]string
    json.Unmarshal(jsonPathBpf, &anode)

    for k,y := range loadData { 
        for y,_ := range y {
            if v, ok := anode[k][y]; ok {
                loadData[k][y] = v
            }else{
                loadData[k][y] = "None"
            }
        }
    }
    return loadData, nil
}

//create conection through http.
func NewRequestHTTP(order string, url string, values io.Reader)(resp *http.Response, err error){
	req, err := http.NewRequest(order, url, values)
	if err != nil {
		logs.Error("Error Executing HTTP new request")
	}
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, DisableKeepAlives: true,}
	client := &http.Client{Transport: tr}
	resp, err = client.Do(req)
	if err != nil {
		logs.Error("Error Retrieving response from client HTTP new request")
	}
	return resp, err
}

//create a backup of selected file
func BackupFile(path string, fileName string) (err error) { 
	loadData := map[string]map[string]string{}
	loadData["files"] = map[string]string{}
	loadData["files"]["backupPath"] = ""
	loadData,err = GetConf(loadData)
	backupPath := loadData["files"]["backupPath"]
	if err != nil {
		logs.Error("Error BackupFile Creating backup: "+err.Error())
		return err
	}

    t := time.Now()
    newFile := fileName+"-"+strconv.FormatInt(t.Unix(), 10)
    srcFolder := path+fileName
    destFolder := backupPath+newFile
    cpCmd := exec.Command("cp", srcFolder, destFolder)
    err = cpCmd.Run()
    if err != nil{
        logs.Error("BackupFile Error exec cmd command: "+err.Error())
        return err
    }
    return nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string)(err error){
	//Get the data	
    resp, err := http.Get(url)
    if err != nil {
		logs.Error("Error downloading file: "+err.Error())
        return err
    }
    defer resp.Body.Close()
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		logs.Error("Error creating file after download: "+err.Error())
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logs.Error("Error Copying downloaded file: "+err.Error())
		return err
	}
	return nil
}

//extract tar.gz files
// func ExtractTarGz(tarGzFile string, pathDownloads string, folder string)(err error){
func ExtractTarGz(tarGzFile string, pathDownloads string)(err error){
	file, err := os.Open(tarGzFile)
	defer file.Close()
	if err != nil {
        return err
	}

	uncompressedStream, err := gzip.NewReader(file)
	if err != nil {
        return err
	}

	tarReader := tar.NewReader(uncompressedStream)
	for true {
		header, err := tarReader.Next()
        if err == io.EOF {
            break
        }
        if err != nil {
			return err
        }

        switch header.Typeflag {
		case tar.TypeDir:
			err := os.MkdirAll(pathDownloads+"/"+header.Name, 0755);
			if err != nil {
				logs.Error("TypeDir: "+err.Error())
				return err
            }
		case tar.TypeReg:
			outFile, err := os.Create(pathDownloads+"/"+header.Name)
			_, err = io.Copy(outFile, tarReader)
            if err != nil {
				logs.Error("TypeReg: "+err.Error())
				return err
            }
        default:
            logs.Error(
                "ExtractTarGz: uknown type: %s in %s",
                header.Typeflag,
                header.Name)
        }
    }
	return nil
}

//create a hashmap from file
func MapFromFile(path string)(mapData map[string]map[string]string, err error){
	var mapFile = make(map[string]map[string]string)
	var validID = regexp.MustCompile(`sid:(\d+);`)
	var enablefield = regexp.MustCompile(`^#`)
	
	file, err := os.Open(path)
	if err != nil {
		logs.Error("Openning File for export to map: "+ err.Error())
		return nil, err
	}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sid := validID.FindStringSubmatch(scanner.Text())
		if sid != nil {
			lineData := make(map[string]string)
			if enablefield.MatchString(scanner.Text()){
                lineData["Enabled"]="Disabled"
            }else{
                lineData["Enabled"]="Enabled"
            }
			lineData["Line"] = scanner.Text()
			mapFile[sid[1]] = lineData
		}
	}
	return mapFile, nil
}

//replace lines between 2 files selected
func ReplaceLines(data map[string]string)(err error){
	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["sourceDownload"] = ""
	sourceDownload,err = GetConf(sourceDownload)
	pathDownloaded := sourceDownload["ruleset"]["sourceDownload"]
	if err != nil {
		logs.Error("ReplaceLines error loading data from main.conf: "+ err.Error())
		return err
	}
	
	//split path 
	splitPath := strings.Split(data["path"], "/")
	pathSelected := splitPath[len(splitPath)-2]

	saved := false
	rulesFile, err := os.Create("_creating-new-file.txt")
	defer rulesFile.Close()
	var validID = regexp.MustCompile(`sid:(\d+);`)

	newFileDownloaded, err := os.Open(pathDownloaded + pathSelected + "/rules/" + "drop.rules")

	scanner := bufio.NewScanner(newFileDownloaded)
	for scanner.Scan() {
		for x := range data{
			sid := validID.FindStringSubmatch(scanner.Text())
			if (sid != nil) && (sid[1] == string(x)) {
				if data[x] == "N/A"{
					saved = true
					continue
				}else{
					_, err = rulesFile.WriteString(string(data[x]))	
					_, err = rulesFile.WriteString("\n")	
					saved = true
					continue
				}
			}
		}
		if !saved{
			_, err = rulesFile.WriteString(scanner.Text())
			_, err = rulesFile.WriteString("\n")	
		}
		saved = false
	}

	input, err := ioutil.ReadFile("_creating-new-file.txt")
	err = ioutil.WriteFile("rules/drop.rules", input, 0644)

	_ = os.Remove("_creating-new-file.txt")

	if err != nil {
		logs.Error("ReplaceLines error writting new lines: "+ err.Error())
		return err
	}
	return nil
}

func CalculateMD5(path string)(md5Data string, err error){
	file, err := os.Open(path)
	if err != nil {
		logs.Error("Error calculating md5: %s", err.Error())
		return "",err
	}
	defer file.Close()
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		logs.Error("Error copying md5: %s", err.Error())
		return "",err
	}

	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String := hex.EncodeToString(hashInBytes)

	return returnMD5String,nil
}

func VerifyPathExists(path string)(stauts string){
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "false"
	}else{
		return "true"
	}
}

func EpochTime(date string)(epoch int64, err error){
	t1 := time.Now() 
	t2,_ := time.ParseInLocation("2006-01-02T15:04:05", date, t1.Location())

	return t2.Unix(), nil
}

func HumanTime(epoch int64)(date string){
	return time.Unix(epoch , 0).String()
}

func BackupFullPath(path string) (err error) { 
    t := time.Now()
    destFolder := path+"-"+strconv.FormatInt(t.Unix(), 10)
    cpCmd := exec.Command("cp", path, destFolder)
    err = cpCmd.Run()
    if err != nil{
        logs.Error("utils.BackupFullPath Error exec cmd command: "+err.Error())
        return err
    }
    return nil
}

func WriteNewDataOnFile(path string, data []byte)(err error){
    err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
        logs.Error("Error WriteNewData")
		return err
	}
    return nil
}

func CopyFile(dstfolder string, srcfolder string, file string, BUFFERSIZE int64) (err error) {
	if BUFFERSIZE == 0{
		BUFFERSIZE = 1000
	}
	sourceFileStat, err := os.Stat(srcfolder+file)
    if err != nil {
        logs.Error("Error checking file at CopyFile function" + err.Error())
        return err
	}
    if !sourceFileStat.Mode().IsRegular() {
        logs.Error("%s is not a regular file.", sourceFileStat)
        return errors.New(sourceFileStat.Name()+" is not a regular file.")
    } 
    source, err := os.Open(srcfolder+file)
    if err != nil {
        return err
    }
    defer source.Close()
    _, err = os.Stat(dstfolder+file)
    if err == nil {
        return errors.New("File "+dstfolder+file+" already exists.")
    }
    destination, err := os.Create(dstfolder+file)
    if err != nil {
        logs.Error("Error Create =-> "+err.Error())
        return err
    }
    defer destination.Close()
    logs.Info("copy file -> "+srcfolder+file)
    logs.Info("to file -> "+dstfolder+file)
    buf := make([]byte, BUFFERSIZE)
    for {
        n, err := source.Read(buf)
        if err != nil && err != io.EOF {
            logs.Error("Error no EOF=-> "+err.Error())
            return err
        }
        if n == 0 {
            break
        }
        if _, err := destination.Write(buf[:n]); err != nil {
            logs.Error("Error Writing File: "+err.Error())
            return err
        }
    }
    return err
}