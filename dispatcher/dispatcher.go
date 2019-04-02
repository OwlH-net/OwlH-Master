package dispatcher

import (
    "time"
    "os"
    "io"
    "io/ioutil"
    "encoding/json"
    "errors"
    "github.com/astaxie/beego/logs"
)

type listOfNodesAndFolders struct {
    Nodes      []Nodes
    Folders    []Folders
}

type Folders struct {
    Fname      string
    Ftype      string
    Fpath      string
}

type Nodes struct {
    Name       string
    Ntype      string
    Pcappath   string
}

var latest int
var timeToken time.Time

func Init() {
	logs.Info ("Starting Dispatcher ...")
    latest = 0
    for runDispatcher() {
        nodesAndPcaps, err  := loadNodesAndPcaps()
        if err!= nil {
            return
        }
        dispatch(nodesAndPcaps)
    }
}

func loadNodesAndPcaps()(nodes listOfNodesAndFolders, err error) {
    nodesFile, err := os.Open("owlhmaster/conf/nodes.json")
    if err != nil {
        logs.Error(err)
    }
    defer nodesFile.Close()
    t := listOfNodesAndFolders{}
    byteValue, _ := ioutil.ReadAll(nodesFile)
    json.Unmarshal([]byte(byteValue), &t)
    return t, nil
}

func noTokenFile() bool {
    if _, err := os.Stat("owlhmaster/conf/stopdispatcher"); os.IsNotExist(err) {
        return true
    }
    return false
}

func runDispatcher() bool {
    if noTokenFile() {
        return true
    }
    return false
}


func copyFileToNode(dstfolder string, srcfolder string, file string, BUFFERSIZE int64) (err error) {
    sourceFileStat, err := os.Stat(srcfolder+file)
    if err != nil {
        logs.Error("Error -> " + err.Error())
        return err
    }
    
    timeToken := time.Now().Add(time.Second*-60)	
    if !sourceFileStat.ModTime().Before(timeToken) {
        return errors.New("Newer file")    
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
            logs.Error("Error Write File =-> "+err.Error())
            return err
        }
    }
    return err
}

func getFileFromSrcFolders (folder string)(files []os.FileInfo, nofile bool) {
    files, err := ioutil.ReadDir(folder)
    if err != nil {
        logs.Error("Error --> "+err.Error())
        return nil, false
    }
    if len(files) == 0 {
        return files, false
    }
    return files, true
}

func dispatch(theList listOfNodesAndFolders) {
    var bufferSize int64
    bufferSize = 10000
    
    var pool []string
    var alone []string
    for j:=0; j < len(theList.Nodes); j++ {
        if theList.Nodes[j].Ntype == "pool" {
            pool = append(pool, theList.Nodes[j].Pcappath)
        }
        if theList.Nodes[j].Ntype == "alone" {
            alone = append(alone, theList.Nodes[j].Pcappath)
        }
    }
    for i:=0; i < len(theList.Folders);i++ {
        files, areFiles := getFileFromSrcFolders(theList.Folders[i].Fpath)
        if !areFiles{
            logs.Info("...waiting Files...")
            time.Sleep(time.Second*10)
            continue
        }
        for _, file := range files {
            if latest == len(pool) {
                latest = 0
            }
            err := copyFileToNode(pool[latest], theList.Folders[i].Fpath,file.Name(), bufferSize)
            if err != nil {
                continue
            }
            for k:=0; k < len(alone); k++ {
                err = copyFileToNode(alone[k], theList.Folders[i].Fpath,file.Name(), bufferSize)
            }
            if err != nil {
                continue
            }
            err = os.Remove(theList.Folders[i].Fpath+file.Name())
            if err != nil {
                logs.Info("Error Removing =-> "+err.Error())
                return
            }
            latest += 1
        }
    }
    return
}
