package validation

import (
    "errors"
    "github.com/astaxie/beego/logs"
    auth "owlhmaster/go-ad-auth.v3"
    "owlhmaster/utils"
    "strconv"
)

type Ldapconfig struct {
    Enabled    bool   `json:"enabled"`
    Server     string `json:"server"`
    Port       int    `json:"port"`
    DN         string `json:"dn"`
    SkipVerify bool   `json:"skipverify"`
}

var configLdap Ldapconfig

func userAuthentication(user string, password string) (check bool, err error) {
    config := &auth.Config{Server: configLdap.Server, Port: configLdap.Port, BaseDN: configLdap.DN, Security: auth.SecurityStartTLS}
    status, err := auth.Authenticate(config, user, password, configLdap.SkipVerify)
    if err != nil { //connection problems
        logs.Error("LDAP Connection problems: " + err.Error())
        return false, err
    }
    if !status { //user/pass problems
        logs.Error("LDAP connection refused: Incorrect user or password")
        //handle failed authentication
        return false, errors.New("LDAP connection refused: Incorrect user or password")
    }
    return true, nil
}

func readLdapConfig() (err error) {
    ldapEnabled, err := utils.GetKeyValueString("ldap", "enabled")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting mainconf data: " + err.Error())
        return err
    }
    ldapServer, err := utils.GetKeyValueString("ldap", "server")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting mainconf data: " + err.Error())
        return err
    }
    ldapPort, err := utils.GetKeyValueString("ldap", "port")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting mainconf data: " + err.Error())
        return err
    }
    ldapDN, err := utils.GetKeyValueString("ldap", "DN")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting mainconf data: " + err.Error())
        return err
    }
    ldapVerify, err := utils.GetKeyValueString("ldap", "skipverify")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting mainconf data: " + err.Error())
        return err
    }

    isEnabled, err := strconv.ParseBool(ldapEnabled)
    configLdap.Enabled = isEnabled
    configLdap.Server = ldapServer
    port, err := strconv.Atoi(ldapPort)
    configLdap.Port = port
    configLdap.DN = ldapDN
    isVerified, err := strconv.ParseBool(ldapVerify)
    configLdap.SkipVerify = isVerified

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
    //      logs.Error(err)
    //         return
    //     }
    //     defer configFile.Close()
    //
    //     byteValue, _ := ioutil.ReadAll(configFile)
    //     json.Unmarshal(byteValue, &configLdap)
}

func CheckLdap(user string, password string) (check bool, err error) {
    err = readLdapConfig()
    if err != nil {
        return false, err
    }

    check, err = userAuthentication(user, password)
    if err != nil {
        return false, err
    }

    return true, nil
}
