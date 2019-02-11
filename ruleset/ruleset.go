package ruleset

import(
    //"io/ioutil"
    "fmt"
    "github.com/astaxie/beego/logs"
    "bufio" //read line by line the doc
    "regexp"
    "os"
    //"strconv"
)

func ReadSID(sid string)( sidLine map[string]string ,err error){

    data, err := os.Open("/etc/owlh/ruleset/owlh.rules")
    if err != nil {
        fmt.Println("File reading error", err)
        return 
    }
    
    var validID = regexp.MustCompile(`sid:`+sid+`;`)
    scanner := bufio.NewScanner(data)
    for scanner.Scan(){
        if validID.MatchString(scanner.Text()){
            sidLine := make(map[string]string)
            sidLine["raw"] = scanner.Text()
            return sidLine,err
        }    
    }
    return nil,err
}

func Read()(rules map[string]map[string]string, err error) {
//leer fichero V
//dar formato a las reglas (json)
//enviar datos a ruleset.html para mostrarlos por pantalla

    //var rules map[string]string
    logs.Info ("Buscando el fichero desde ruleset/ruleset.go")
    //data, err := ioutil.ReadFile("/etc/owlh/ruleset/owlh.rules")
    data, err := os.Open("/etc/owlh/ruleset/owlh.rules")
    
    if err != nil {
        fmt.Println("File reading error", err)
        return 
    }

    var validID = regexp.MustCompile(`sid:(\d+);`)
    var ipfield = regexp.MustCompile(`^([^\(]+)\(`)
    var msgfield = regexp.MustCompile(`msg:([^;]+);`)
    var enablefield = regexp.MustCompile(`^#`)

    scanner := bufio.NewScanner(data)
    rules = make(map[string]map[string]string)
    for scanner.Scan(){
        if validID.MatchString(scanner.Text()){
            sid := validID.FindStringSubmatch(scanner.Text())
            msg := msgfield.FindStringSubmatch(scanner.Text())
            ip := ipfield.FindStringSubmatch(scanner.Text())
            rule := make(map[string]string)

            if enablefield.MatchString(scanner.Text()){
                rule["enabled"]="Disabled"
            }else{
                rule["enabled"]="Enabled"
            }

            rule["sid"]=sid[1]
            rule["msg"]=msg[1]
            rule["ip"]=ip[1]
            rule["raw"]=scanner.Text()
            rules[sid[1]]=rule

        }
    }
    return rules,err
}