package validation

import (
    auth "gopkg.in/korylprince/go-ad-auth.v3"
	"github.com/astaxie/beego/logs"
	"owlhmaster/utils"
	"errors"
	"strconv"
)

type Ldapconfig struct {
    Enabled         bool      	`json:"enabled"`
    Server          string      `json:"server"`
    Port            int         `json:"port"`
    DN              string      `json:"dn"`
    SkipVerify      bool        `json:"skipverify"`
}
var configLdap      Ldapconfig

func userAuthentication(user string, password string)(check bool, err error) {
    config := &auth.Config{Server:configLdap.Server, Port:configLdap.Port, BaseDN:configLdap.DN, Security:auth.SecurityStartTLS}
	status, err := auth.Authenticate(config, user, password, configLdap.SkipVerify)
    if err != nil { //connection problems
        logs.Error("LDAP Connection problems: "+err.Error())
        return false,err
    }
    if !status { //user/pass problems
        logs.Error("LDAP connection refused: "+err.Error())
        //handle failed authentication
        return false,err 
	}
	return true,nil
}

func readLdapConfig()(err error) {
	mainconfLdap := map[string]map[string]string{}
	mainconfLdap["ldap"] = map[string]string{}
    mainconfLdap["ldap"]["enabled"] = ""
    mainconfLdap["ldap"]["server"] = ""
    mainconfLdap["ldap"]["port"] = ""
    mainconfLdap["ldap"]["DN"] = ""
    mainconfLdap["ldap"]["skipverify"] = ""
    mainconfLdap,err = utils.GetConf(mainconfLdap)
    if err != nil {logs.Error("ldap/readLdapConfig -- Error getting mainconf data: "+err.Error()); return err}
	isEnabled, err := strconv.ParseBool(mainconfLdap["ldap"]["enabled"]); configLdap.Enabled = isEnabled
	configLdap.Server =  mainconfLdap["ldap"]["server"]
	port, err := strconv.Atoi(mainconfLdap["ldap"]["port"]); configLdap.Port = port
	configLdap.DN =  mainconfLdap["ldap"]["DN"]
	isVerified, err := strconv.ParseBool(mainconfLdap["ldap"]["skipverify"]); configLdap.SkipVerify = isVerified

	if !isVerified {
		return errors.New("This user has not enough permissions for validate using LDAP")
	}

	return nil
/////////////////////////////////////////////////////////////////////	
			//READ FROM main.conf
			//save into Ldapconfig struct
			//if main.conf ldap.enabled == "disabled" {return err}
/////////////////////////////////////////////////////////////////////	

//     file := "config.json"
// 
//     configFile, err := os.Open(file)
//     if err != nil {
// 		logs.Error(err)
//         return
//     }
//     defer configFile.Close()
// 
//     byteValue, _ := ioutil.ReadAll(configFile)
//     json.Unmarshal(byteValue, &configLdap)
}

func CheckLdap(user string, password string)(check bool, err error){
	err = readLdapConfig()
	if err != nil {return false, err}
	
	check,err = userAuthentication(user, password);
	if err != nil {return false, err}
	
	return true, nil
}