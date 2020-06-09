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
    logs.Debug("LDAP - AUTH - config -> %+v", config)
    logs.Debug("user -> %s", user)
    logs.Debug("pass -> %s", password)
    logs.Debug("LDAP - AUTH - Call in progress")
    logs.Debug("SkipVerify -> %t", configLdap.SkipVerify)
    status, err1 := auth.Authenticate(config, user, password, configLdap.SkipVerify)
    logs.Debug("LDAP - AUTH - Call done - status is %t", status)

    if err1 != nil { //connection problems
        logs.Error("LDAP Connection problems: " + err1.Error())
        return false, err
    }
    if !status { //user/pass problems
        logs.Error("LDAP - AUTH - user/pass error")
        return false, errors.New("Error - user can't get in")
    }
    logs.Debug("LDAP - auth DONE - user can get in")
    return true, err1
}

func readLdapConfig() (err error) {
    ldapEnabled, err := utils.GetKeyValueString("ldap", "enabled")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting ldap - enabled data: " + err.Error())
        return err
    }
    ldapServer, err := utils.GetKeyValueString("ldap", "server")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting ldap - server data: " + err.Error())
        return err
    }
    ldapPort, err := utils.GetKeyValueString("ldap", "port")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting ldap - port data: " + err.Error())
        return err
    }
    ldapDN, err := utils.GetKeyValueString("ldap", "DN")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting ldap - DN data: " + err.Error())
        return err
    }
    ldapVerify, err := utils.GetKeyValueString("ldap", "skipverify")
    if err != nil {
        logs.Error("ldap/readLdapConfig -- Error getting ldap - skipverify data: " + err.Error())
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

    return nil
}

func CheckLdap(user string, password string) (check bool, err error) {
    err = readLdapConfig()
    if err != nil {
        logs.Error("LDAP - Error reading LDAP configuration, can't continue user validation")
        return false, err
    }

    logs.Debug("LDAP - Let's verify User -> %s", user)
    check, err = userAuthentication(user, password)
    if err != nil {
        logs.Error("LDAP - Auth - Error - %s", err.Error())
        return false, err
    }
    logs.Debug("LDAP - User -> %s, can get in", user)

    return true, nil
}
