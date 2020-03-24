package utils


import (
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
    "sort"
    "path/filepath"
)

//**********token global variables**********
var TokenMasterValidated string
var TokenMasterUser string
var TokenMasterUuid string
var NodeToken string
//**********token global variables**********

func Generate()(uuid string)  {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        logs.Info(err)
    }
    uuid = fmt.Sprintf("%x-%x-%x-%x-%x",b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    return uuid
}

//create conection through http.
func NewRequestHTTP(order string, url string, values io.Reader)(resp *http.Response, err error){
    req, err := http.NewRequest(order, url, values)
    req.Header.Set("token", TokenMasterValidated)
    req.Header.Set("user", TokenMasterUser)
    if err != nil {
        logs.Error("Error Executing HTTP new request")
    }
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, DisableKeepAlives: true,}
    client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
    resp, err = client.Do(req)
    if err != nil {
        logs.Error("Error Retrieving response from client HTTP new request")
    }
    return resp, err
}

//create a backup of selected file
func BackupFile(path string, fileName string) (err error) {
    backupFolder, err := GetKeyValueString("files", "backupPath")
    if err != nil {logs.Error("Error BackupFile Creating backup: "+err.Error()); return err}

    // check if folder exists
    if _, err := os.Stat(backupFolder); os.IsNotExist(err) {
        err = os.MkdirAll(backupFolder, 0755)
        if err != nil{logs.Error("utils.BackupFile Error creating main backup folder: "+err.Error()); return err}
    }

    //get older backup file
    listOfFiles,err := FilelistPathByFile(backupFolder, fileName)
    if err != nil{logs.Error("utils.BackupFile Error walking through backup folder: "+err.Error()); return err}
    count := 0
    previousBck := ""
    for x := range listOfFiles{
        count++
        if previousBck == "" {
            previousBck = x
            continue
        }else if previousBck > x{
            previousBck = x
        }
    }

    //delete older bck file if there are 5 bck files
    if count == 5 {
        err = os.Remove(backupFolder+previousBck)
        if err != nil{logs.Error("utils.BackupFile Error deleting older backup file: "+err.Error())}
    }

    //create backup
    t := time.Now()
    newFile := fileName+"-"+strconv.FormatInt(t.Unix(), 10)
    srcFolder := path+fileName
    destFolder := backupFolder+newFile

    copy, err := GetKeyValueString("execute", "copy")
    if err != nil {logs.Error("BackupFile Error getting data from main.conf"); return err}

    //check if file exist
    if _, err := os.Stat(srcFolder); os.IsNotExist(err) {
        return errors.New("utils.BackupFile error: Source file doesn't exists")
    }else{
        cpCmd := exec.Command(copy, srcFolder, destFolder)
        err = cpCmd.Run()
        if err != nil{logs.Error("utils.BackupFile Error exec cmd command: "+err.Error()); return err}
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
func ExtractFile(tarGzFile string, pathDownloads string)(err error){
    base := filepath.Base(tarGzFile)
    fileType := strings.Split(base, ".")

    wget, err := GetKeyValueString("execute", "command")
    if err != nil {logs.Error("ExtractFile Error getting data from main.conf"); return err}

    if fileType[len(fileType)-1] == "rules"{
        cmd := exec.Command(wget, tarGzFile, "-O", pathDownloads)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()

    // }else if fileType[len(fileType)-1] == "gz"{
    }else{
        // if fileType[len(fileType)-2] == "tar"{
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
        // }
    // }else if fileType[len(fileType)-1] == "tgz"{
    }

    return nil
}

//create a hashmap from file
func MapFromFile(path string)(mapData map[string]map[string]string, err error){
    var mapFile = make(map[string]map[string]string)
    var validID = regexp.MustCompile(`sid:\s?(\d+);`)
    var enablefield = regexp.MustCompile(`^#`)

    file, err := os.Open(path)
    if err != nil {logs.Error("utils/MapFromFile Error openning file for export to map: "+ err.Error()); return nil, err}

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

//merge some files thought their path
func MergeAllFiles(files []string)(content []byte, err error){
    allFiles := make(map[string]map[string]string)
    for x := range files {
        //only enabled lines
        lines,err := MapFromFile(files[x])
        if err != nil {logs.Error("MergeAllFiles/MapFromFile error creating map from file: "+err.Error()); return nil,err}
        for y := range lines {
            // exists := false
            if lines[y]["Enabled"] == "Enabled" {
                if allFiles[y] == nil { allFiles[y] = map[string]string{}}
                allFiles[y] = lines[y]
                // for z := range allFiles {
                //     if y == z {
                //         exists = true
                //     }
                // }
                // if exists {allFiles[y] = lines[y]}
            }
        }
    }
    for r := range allFiles{
        content = append(content, []byte(allFiles[r]["Line"])...)
        content = append(content, []byte("\n")...)
    }
    return content, nil
}

//replace lines between 2 files selected
func ReplaceLines(data map[string]string)(err error){
    pathDownloaded, err := GetKeyValueString("ruleset", "sourceDownload")
    if err != nil {logs.Error("ReplaceLines error loading data from main.conf: "+ err.Error());return err}
    ruleFile, err := GetKeyValueString("ruleset", "ruleFile")
    if err != nil {logs.Error("ReplaceLines error loading data from main.conf: "+ err.Error());return err}

    //split path
    splitPath := strings.Split(data["path"], "/")
    pathSelected := splitPath[len(splitPath)-2]

    saved := false
    rulesFile, err := os.Create("_creating-new-file.txt")
    defer rulesFile.Close()
    var validID = regexp.MustCompile(`sid:\s?(\d+);`)

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
    // err = ioutil.WriteFile("rules/drop.rules", input, 0644)
    err = ioutil.WriteFile(ruleFile, input, 0644)

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
    copy, err := GetKeyValueString("execute", "copy")
    if err != nil {logs.Error("BackupFullPath Error getting data from main.conf"); return err}

    t := time.Now()
    destFolder := path+"-"+strconv.FormatInt(t.Unix(), 10)
    cpCmd := exec.Command(copy, path, destFolder)
    err = cpCmd.Run()
    if err != nil{
        logs.Error("utils.BackupFullPath Error exec cmd command: "+err.Error())
        return err
    }
    return nil
}

func WriteNewDataOnFile(path string, data []byte)(err error){
    err = ioutil.WriteFile(path, data, 0644)
    if err != nil {logs.Error("Error WriteNewData"); return err}

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

func SortHashMap(data map[string]map[string]string)(dataSorted map[string]map[string]string){
    var val []string
    sortedValues := make(map[string]map[string]string)
    for x:=range data {
        val = append(val, strings.ToLower(data[x]["name"]))
    }
    sort.Strings(val)

    for z := range val{
        for y := range data {
            if strings.ToLower(val[z]) == strings.ToLower(data[y]["name"]) {
                if sortedValues[y] == nil { sortedValues[y] = map[string]string{}}
                sortedValues[y] = data[y]
            }
        }
    }

    return sortedValues
}

func ListFilepath(path string)(files map[string][]byte, err error){
    pathMap:= make(map[string][]byte)
    err = filepath.Walk(path,
        func(file string, info os.FileInfo, err error) error {
        if err != nil {return err}

        if !info.IsDir() {
            pathSplit := strings.Split(file, "/")
            content, err := ioutil.ReadFile(file)
            if err != nil {logs.Error("Error filepath walk: "+err.Error()); return err}
            pathMap[pathSplit[len(pathSplit)-1]] = content
        }
        return nil
    })
    if err != nil {logs.Error("Error filepath walk finish: "+err.Error()); return nil, err}

    return pathMap, nil
}

func FilelistPathByFile(path string, fileToSearch string)(files map[string][]byte, err error){
    pathMap:= make(map[string][]byte)
    err = filepath.Walk(path,
        func(file string, info os.FileInfo, err error) error {
        if err != nil {return err}

        if !info.IsDir() {
            pathSplit := strings.Split(file, "/")
            if strings.Contains(pathSplit[len(pathSplit)-1], fileToSearch){
                content, err := ioutil.ReadFile(file)
                if err != nil {logs.Error("Error filepath walk: "+err.Error()); return err}
                pathMap[pathSplit[len(pathSplit)-1]] = content
            }
        }
        return nil
    })
    if err != nil {logs.Error("Error filepath walk finish: "+err.Error()); return nil, err}

    return pathMap, nil
}