package utils

import (
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "io/ioutil"
	"io"
	"os"
	"net/http"
	"crypto/tls"
	"archive/tar"
	"compress/gzip"
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

func ExtractTarGz(filePath string)(err error){
	os.Remove("/tmp/owlh/unzipped")
	file, err := os.Open(filePath)
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