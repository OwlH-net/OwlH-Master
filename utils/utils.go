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
	"time"
)

//Read map data
//leer json del fichero para obtener el path del bpf
func GetConf(loadData map[string]map[string]string)(loadDataReturn map[string]map[string]string, err error) { 
    // confFilePath := "/etc/owlh/conf/main.conf"
    confFilePath := "./conf/main.conf"
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

func BackupFile(path string, fileName string) (err error) { 
    t := time.Now()
    newFile := fileName+"-"+strconv.FormatInt(t.Unix(), 10)
    srcFolder := path+fileName
    destFolder := path+newFile
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
func DownloadFile(filepath string, url string) error {
    // Get the data
    resp, err := http.Get(url)
    if err != nil {
		logs.Error("Error downloading file: "+err.Error())
        return err
    }
    defer resp.Body.Close()

	//check if the directory exists
	// ck := exists(filepath)
	// if ck == true { // path contains data
	// 	//remove from old_download
	// 	os.RemoveAll("/tmp/owlh/old_download")
	// 	//copy from new_download to old_download
	// 	err = CopyDir("/tmp/owlh/new_download/", "/tmp/owlh/old_download/")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	//remove from new_download
	// 	os.RemoveAll("/tmp/owlh/new_download")
	// 	//download new content
	// 	// out, err := os.Create(filepath)
	// 	// if err != nil {
	// 	// 	logs.Error("Error creating file after download: "+err.Error())
	// 	// 	return err
	// 	// }
	// 	// defer out.Close()
	// 	// _, err = io.Copy(out, resp.Body)
	// 	// if err != nil {
	// 	// 	logs.Error("Error Copying downloaded file: "+err.Error())
	// 	// 	return err
	// 	// }
	// }else{ //path is empty

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
	// }
	return nil
}

func ExtractTarGz(filepath string)(err error){
	file, err := os.Open(filepath)
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
			err := os.Mkdir(header.Name, 0755);
			if err != nil {
				logs.Error("TypeDir: "+err.Error())
				return err
            }
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
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

// exists returns whether the given file or directory exists
func exists(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { return false}
    if os.IsNotExist(err) { return true }
    return false
}

// func CopyDir(source string, dest string) (err error) {
// 	// get properties of source dir
// 	sourceinfo, err := os.Stat(source)
// 	if err != nil {
// 		return err
// 	}
// 	// create dest dir
// 	err = os.MkdirAll(dest, sourceinfo.Mode())
// 	if err != nil {
// 		return err
// 	}
// 	directory, _ := os.Open(source)
// 	objects, err := directory.Readdir(-1)
// 	for _, obj := range objects {
// 		sourcefilepointer := source + "/" + obj.Name()
// 		destinationfilepointer := dest + "/" + obj.Name()
// 		if obj.IsDir() {
// 			// create sub-directories - recursively
// 			err = CopyDir(sourcefilepointer, destinationfilepointer)
// 			if err != nil {
// 				logs.Error("Error copying folder recursively: "+err.Error())
// 			}
// 		} else {
// 			// perform copy
// 			err = CopyFile(sourcefilepointer, destinationfilepointer)
// 			if err != nil {
// 				logs.Error("Error copying file recursively: "+err.Error())
// 			}
// 		}
// 	}
// 	return
// }

// func CopyFile(source string, dest string) (err error) {
// 	sourcefile, err := os.Open(source)
// 	if err != nil {
// 		return err
// 	}
// 	defer sourcefile.Close()
// 	destfile, err := os.Create(dest)
// 	if err != nil {
// 		return err
// 	}
// 	defer destfile.Close()
// 	_, err = io.Copy(destfile, sourcefile)
// 	if err == nil {
// 		sourceinfo, err := os.Stat(source)
// 		if err != nil {
// 			err = os.Chmod(dest, sourceinfo.Mode())
// 		}
// 	}
// 	return
// }

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

func ReplaceLines(data map[string]string)(err error){
	logs.Error(data)
	saved := false
	f, err := os.Create("_creating-new-file.txt")
	defer f.Close()
	var validID = regexp.MustCompile(`sid:(\d+);`)

	fileOpened, err := os.Open("rules2/drop.rules")
	scanner := bufio.NewScanner(fileOpened)
	for scanner.Scan() {
		for x := range data{
			if data[x] == "N/A"{
				logs.Error(data[x])
				continue
			}
			sid := validID.FindStringSubmatch(scanner.Text())
			if (sid != nil) && (sid[1] == string(x)) {
				_, err = f.WriteString(string(data[x]))	
				_, err = f.WriteString("\n")	
				saved = true
			}
		}
		if !saved{
			_, err = f.WriteString(scanner.Text())
			_, err = f.WriteString("\n")	
		}
		saved = false
	}
	if err != nil {
		logs.Error("ReplaceLines error writting new lines: "+ err.Error())
		return err
	}
	// b, err := ioutil.ReadFile("_creating-new-file.txt")
	// _ = os.Remove("_creating-new-file.txt")
	return nil
}