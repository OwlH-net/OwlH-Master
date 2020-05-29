package configuration

import (
    _ "github.com/mattn/go-sqlite3"
)

func checkPermissionsFields()(ok bool){

    var field Field

    //role groups
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='admin' and rg_param='group' and rg_value='admin'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('admin','group','admin')"
    field.Fname      = "roleGroups - admin group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='admin' and rg_param='desc' and rg_value='Admin privileges'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('admin','desc','Admin privileges')"
    field.Fname      = "roleGroups - admin group"
    ok = CheckField(field)
    if !ok {return false}
    
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='OpenRules' and rg_param='group' and rg_value='Open rules'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('OpenRules','group','Open rules')"
    field.Fname      = "roleGroups - OpenRules group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='OpenRules' and rg_param='desc' and rg_value='Open rules permissions'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('OpenRules','desc','Open rules permissions')"
    field.Fname      = "roleGroups - OpenRules group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='ChangeControl' and rg_param='group' and rg_value='Change control'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('ChangeControl','group','Change control')"
    field.Fname      = "roleGroups - ChangeControl group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='ChangeControl' and rg_param='desc' and rg_value='Change control information'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('ChangeControl','desc','Change control information')"
    field.Fname      = "roleGroups - ChangeControl group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Collector' and rg_param='group' and rg_value='Collector'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Collector','group','Collector')"
    field.Fname      = "roleGroups - Collector group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Collector' and rg_param='desc' and rg_value='Collector information'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Collector','desc','Collector information')"
    field.Fname      = "roleGroups - Collector group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Group' and rg_param='group' and rg_value='Group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Group','group','Group')"
    field.Fname      = "roleGroups - Group group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Group' and rg_param='desc' and rg_value='Group information'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Group','desc','Group information')"
    field.Fname      = "roleGroups - Group group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Master' and rg_param='group' and rg_value='Master'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Master','group','Master')"
    field.Fname      = "roleGroups - Master group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Master' and rg_param='desc' and rg_value='Master group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Master','desc','Master group')"
    field.Fname      = "roleGroups - Master group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Node' and rg_param='group' and rg_value='Node'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Node','group','Node')"
    field.Fname      = "roleGroups - Node group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Node' and rg_param='desc' and rg_value='Node group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Node','desc','Node group')"
    field.Fname      = "roleGroups - Node group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Ruleset' and rg_param='group' and rg_value='Ruleset'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Ruleset','group','Ruleset')"
    field.Fname      = "roleGroups - Ruleset group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Ruleset' and rg_param='desc' and rg_value='Ruleset group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Ruleset','desc','Ruleset group')"
    field.Fname      = "roleGroups - Ruleset group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='RulesetSource' and rg_param='group' and rg_value='RulesetSource'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('RulesetSource','group','RulesetSource')"
    field.Fname      = "roleGroups - RulesetSource group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='RulesetSource' and rg_param='desc' and rg_value='RulesetSource group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('RulesetSource','desc','RulesetSource group')"
    field.Fname      = "roleGroups - RulesetSource group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Scheduler' and rg_param='group' and rg_value='Scheduler'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Scheduler','group','Scheduler')"
    field.Fname      = "roleGroups - Scheduler group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Scheduler' and rg_param='desc' and rg_value='Scheduler group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Scheduler','desc','Scheduler group')"
    field.Fname      = "roleGroups - Scheduler group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Scheduler' and rg_param='group' and rg_value='Scheduler'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Scheduler','group','Scheduler')"
    field.Fname      = "roleGroups - Scheduler group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Scheduler' and rg_param='desc' and rg_value='Scheduler group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Scheduler','desc','Scheduler group')"
    field.Fname      = "roleGroups - Scheduler group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Search' and rg_param='group' and rg_value='Search'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Search','group','Search')"
    field.Fname      = "roleGroups - Search group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Search' and rg_param='desc' and rg_value='Search group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Search','desc','Search group')"
    field.Fname      = "roleGroups - Search group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Stap' and rg_param='group' and rg_value='Stap'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Stap','group','Stap')"
    field.Fname      = "roleGroups - Stap group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Stap' and rg_param='desc' and rg_value='Stap group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Stap','desc','Stap group')"
    field.Fname      = "roleGroups - Stap group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Analyzer' and rg_param='group' and rg_value='Analyzer'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Analyzer','group','Analyzer')"
    field.Fname      = "roleGroups - Analyzer group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Analyzer' and rg_param='desc' and rg_value='Analyzer group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Analyzer','desc','Analyzer group')"
    field.Fname      = "roleGroups - Analyzer group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Autentication' and rg_param='group' and rg_value='Autentication'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Autentication','group','Autentication')"
    field.Fname      = "roleGroups - Autentication group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Autentication' and rg_param='desc' and rg_value='Autentication group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Autentication','desc','Autentication group')"
    field.Fname      = "roleGroups - Autentication group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Autentication' and rg_param='group' and rg_value='Autentication'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Autentication','group','Autentication')"
    field.Fname      = "roleGroups - Autentication group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Autentication' and rg_param='desc' and rg_value='Autentication group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Autentication','desc','Autentication group')"
    field.Fname      = "roleGroups - Autentication group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='DataFlow' and rg_param='group' and rg_value='DataFlow'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('DataFlow','group','DataFlow')"
    field.Fname      = "roleGroups - DataFlow group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='DataFlow' and rg_param='desc' and rg_value='DataFlow group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('DataFlow','desc','DataFlow group')"
    field.Fname      = "roleGroups - DataFlow group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Deploy' and rg_param='group' and rg_value='Deploy'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Deploy','group','Deploy')"
    field.Fname      = "roleGroups - Deploy group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Deploy' and rg_param='desc' and rg_value='Deploy group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Deploy','desc','Deploy group')"
    field.Fname      = "roleGroups - Deploy group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='File' and rg_param='group' and rg_value='File'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('File','group','File')"
    field.Fname      = "roleGroups - File group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='File' and rg_param='desc' and rg_value='File group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('File','desc','File group')"
    field.Fname      = "roleGroups - File group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Incidents' and rg_param='group' and rg_value='Incidents'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Incidents','group','Incidents')"
    field.Fname      = "roleGroups - Incidents group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Incidents' and rg_param='desc' and rg_value='Incidents group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Incidents','desc','Incidents group')"
    field.Fname      = "roleGroups - Incidents group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Monitor' and rg_param='group' and rg_value='Monitor'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Monitor','group','Monitor')"
    field.Fname      = "roleGroups - Monitor group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Monitor' and rg_param='desc' and rg_value='Monitor group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Monitor','desc','Monitor group')"
    field.Fname      = "roleGroups - Monitor group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Net' and rg_param='group' and rg_value='Net'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Net','group','Net')"
    field.Fname      = "roleGroups - Net group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Net' and rg_param='desc' and rg_value='Net group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Net','desc','Net group')"
    field.Fname      = "roleGroups - Net group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Ping' and rg_param='group' and rg_value='Ping'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Ping','group','Ping')"
    field.Fname      = "roleGroups - Ping group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Ping' and rg_param='desc' and rg_value='Ping group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Ping','desc','Ping group')"
    field.Fname      = "roleGroups - Ping group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Plugin' and rg_param='group' and rg_value='Plugin'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Plugin','group','Plugin')"
    field.Fname      = "roleGroups - Plugin group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Plugin' and rg_param='desc' and rg_value='Plugin group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Plugin','desc','Plugin group')"
    field.Fname      = "roleGroups - Plugin group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Ports' and rg_param='group' and rg_value='Ports'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Ports','group','Ports')"
    field.Fname      = "roleGroups - Ports group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Ports' and rg_param='desc' and rg_value='Ports group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Ports','desc','Ports group')"
    field.Fname      = "roleGroups - Ports group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Suricata' and rg_param='group' and rg_value='Suricata'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Suricata','group','Suricata')"
    field.Fname      = "roleGroups - Suricata group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Suricata' and rg_param='desc' and rg_value='Suricata group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Suricata','desc','Suricata group')"
    field.Fname      = "roleGroups - Suricata group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Wazuh' and rg_param='group' and rg_value='Wazuh'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Wazuh','group','Wazuh')"
    field.Fname      = "roleGroups - Wazuh group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Wazuh' and rg_param='desc' and rg_value='Wazuh group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Wazuh','desc','Wazuh group')"
    field.Fname      = "roleGroups - Wazuh group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Zeek' and rg_param='group' and rg_value='Zeek'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Zeek','group','Zeek')"
    field.Fname      = "roleGroups - Zeek group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "roleGroups"
    field.Fquery     = "select rg_value from roleGroups where rg_uniqueid='Zeek' and rg_param='desc' and rg_value='Zeek group'"
    field.Finsert    = "insert into roleGroups (rg_uniqueid,rg_param,rg_value) values ('Zeek','desc','Zeek group')"
    field.Fname      = "roleGroups - Zeek group"
    ok = CheckField(field)
    if !ok {return false}


	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
	//add user permissions
    //add user permissions
    
	//add hwaddmng
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddMacIp' and per_param='desc' and per_value='PCAP add MAC IP'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMacIp','desc','PCAP add MAC IP')"
    field.Fname      = "permissions - AddMacIp desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddMacIp' and per_param='permissionGroup' and per_value='HwAddMng'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMacIp','permissionGroup','HwAddMng')"
    field.Fname      = "permissions - AddMacIp group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadConfig' and per_param='desc' and per_value='Load PCAP MAC Configuration'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadConfig','desc','Load PCAP MAC Configuration')"
    field.Fname      = "permissions - LoadConfig desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadConfig' and per_param='permissionGroup' and per_value='HwAddMng'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadConfig','permissionGroup','HwAddMng')"
    field.Fname      = "permissions - LoadConfig group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ConfigPost' and per_param='desc' and per_value='Configuration PCAP MAC POST'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ConfigPost','desc','Configuration PCAP MAC POST')"
    field.Fname      = "permissions - ConfigPost desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ConfigPost' and per_param='permissionGroup' and per_value='HwAddMng'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ConfigPost','permissionGroup','HwAddMng')"
    field.Fname      = "permissions - ConfigPost group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='Db' and per_param='desc' and per_value='Configuration DB PCAP MAC'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Db','desc','Configuration DB PCAP MAC')"
    field.Fname      = "permissions - Db desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='Db' and per_param='permissionGroup' and per_value='HwAddMng'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Db','permissionGroup','HwAddMng')"
    field.Fname      = "permissions - Db group"
    ok = CheckField(field)
    if !ok {return false}

    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ConfigGet' and per_param='desc' and per_value='Configuration PCAP MAC GET'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ConfigGet','desc','Configuration PCAP MAC GET')"
    field.Fname      = "permissions - ConfigGet desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ConfigGet' and per_param='permissionGroup' and per_value='HwAddMng'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ConfigGet','permissionGroup','HwAddMng')"
    field.Fname      = "permissions - ConfigGet group"
    ok = CheckField(field)
    if !ok {return false}

	//add user permissions
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='delete' and per_param='desc' and per_value='http delete request'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('delete','desc','http delete request')"
    field.Fname      = "permissions - delete desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='delete' and per_param='permissionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('delete','permissionGroup','admin')"
    field.Fname      = "permissions - delete group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='post' and per_param='desc' and per_value='http post request'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('post','desc','http post request')"
    field.Fname      = "permissions - post desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='post' and per_param='permissionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('post','permissionGroup','admin')"
    field.Fname      = "permissions - post group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='put' and per_param='desc' and per_value='http put request'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('put','desc','http put request')"
    field.Fname      = "permissions - put desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='put' and per_param='permissionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('put','permissionGroup','admin')"
    field.Fname      = "permissions - put group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='get' and per_param='desc' and per_value='http get request'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('get','desc','http get request')"
    field.Fname      = "permissions - get desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='get' and per_param='permissionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('get','permissionGroup','admin')"
    field.Fname      = "permissions - get group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='admin' and per_param='desc' and per_value='Get everything'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('admin','desc','Get everything')"
    field.Fname      = "permissions - admin desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='admin' and per_param='permissionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('admin','permissionGroup','admin')"
    field.Fname      = "permissions - admin group"
	ok = CheckField(field)
	
	//ChangeControl requests
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControl' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControl','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControl' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControl','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddServer' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddServer','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllServers' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllServers','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllServers' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllServers','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServer' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServer','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStap' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStap','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStap' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStap','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStap' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStap','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStap' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStap','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStap' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStap','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStap' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStap','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapServer' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapServer','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServer' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServer','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingServerStap' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServerStap','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingServerStap' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServerStap','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapServer' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapServer','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapServerput' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapServerput','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapServerput' and per_param='permissionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapServerput','permissionGroup','ChangeControl')"
    field.Fname      = "permissions - group"
	ok = CheckField(field)
    
    //Collector
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='permissionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','permissionGroup','Collector')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopCollector' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopCollector','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopCollector' and per_param='permissionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopCollector','permissionGroup','Collector')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowCollector' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowCollector','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowCollector' and per_param='permissionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowCollector','permissionGroup','Collector')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayMasterCollector' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayMasterCollector','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayMasterCollector' and per_param='permissionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayMasterCollector','permissionGroup','Collector')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopMasterCollector' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopMasterCollector','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopMasterCollector' and per_param='permissionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopMasterCollector','permissionGroup','Collector')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowMasterCollector' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowMasterCollector','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowMasterCollector' and per_param='permissionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowMasterCollector','permissionGroup','Collector')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
    
    //Group
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='permissionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','permissionGroup','Collector')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
    //Stap
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddStapServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddStapServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddStapServer' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddStapServer','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllStapServers' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllStapServers','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllStapServers' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllStapServers','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStapServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStapServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStapServer' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStapServer','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStapStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStapStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStapStatus' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStapStatus','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapService' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapService','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapService' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapService','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapService' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapService','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
		
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingStapServer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingStapServer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingStapServer' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingStapServer','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapService' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapService','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapService' and per_param='permissionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapService','permissionGroup','Stap')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
    
    //Search
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRulesetsBySearch' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRulesetsBySearch','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRulesetsBySearch' and per_param='permissionGroup' and per_value='Search'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRulesetsBySearch','permissionGroup','Search')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
    //Scheduler
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SchedulerTask' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SchedulerTask','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SchedulerTask' and per_param='permissionGroup' and per_value='Scheduler'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SchedulerTask','permissionGroup','Scheduler')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopTask' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopTask','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopTask' and per_param='permissionGroup' and per_value='Scheduler'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopTask','permissionGroup','Scheduler')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetLog' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetLog','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetLog' and per_param='permissionGroup' and per_value='Scheduler'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetLog','permissionGroup','Scheduler')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

    //Master
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetFileContent' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetFileContent','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetFileContent' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetFileContent','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContent' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContent','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContent' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContent','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPlugins' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPlugins','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPlugins' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPlugins','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingFlow' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingFlow','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingFlow' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingFlow','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePluginStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePluginStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePluginStatus' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePluginStatus','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveStapInterface' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveStapInterface','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveStapInterface' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveStapInterface','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowStatus' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowStatus','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNetworkInterface' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNetworkInterface','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNetworkInterface' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNetworkInterface','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateMasterNetworkInterface' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateMasterNetworkInterface','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateMasterNetworkInterface' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateMasterNetworkInterface','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllGroupRulesetsForAllNodes' and per_param='desc' and per_value='Get all Groups and group rulesets for every node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllGroupRulesetsForAllNodes','desc','Get all Groups and group rulesets for every node')"
    field.Fname      = "permissions - GetAllGroupRulesetsForAllNodes desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllGroupRulesetsForAllNodes' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllGroupRulesetsForAllNodes','permissionGroup','Master')"
    field.Fname      = "permissions - GetAllGroupRulesetsForAllNodes group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadMasterNetworkValuesSelected' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadMasterNetworkValuesSelected','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadMasterNetworkValuesSelected' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadMasterNetworkValuesSelected','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingServiceMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServiceMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingServiceMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServiceMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployServiceMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployServiceMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployServiceMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployServiceMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginServiceMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginServiceMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginServiceMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginServiceMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteServiceMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteServiceMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteServiceMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteServiceMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValuesMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValuesMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValuesMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValuesMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateMasterStapInterface' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateMasterStapInterface','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateMasterStapInterface' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateMasterStapInterface','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetBPF' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetBPF','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetBPF' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetBPF','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapServiceMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapServiceMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapServiceMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapServiceMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidents' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidents','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidents' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidents','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveZeekValues' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveZeekValues','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveZeekValues' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveZeekValues','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsMaster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsMaster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsMaster' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsMaster','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetPathFileContent' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetPathFileContent','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetPathFileContent' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetPathFileContent','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFilePathContent' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFilePathContent','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFilePathContent' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFilePathContent','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddUser' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUser','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddUser' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUser','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllUsers' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllUsers','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllUsers' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllUsers','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUser' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUser','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUser' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUser','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupUsers' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupUsers','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupUsers' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupUsers','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRole' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRole','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRole' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRole','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolesForUser' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolesForUser','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolesForUser' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolesForUser','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetGroupsForUser' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetGroupsForUser','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetGroupsForUser' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetGroupsForUser','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddUsersTo' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUsersTo','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddUsersTo' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUsersTo','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePassword' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePassword','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePassword' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePassword','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUserRole' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUserRole','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUserRole' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUserRole','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUserGroup' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUserGroup','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUserGroup' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUserGroup','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRoles' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRoles','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRoles' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRoles','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRole' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRole','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRole' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRole','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRole' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRole','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRole' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRole','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllUserGroups' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllUserGroups','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllUserGroups' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllUserGroups','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditUserGroup' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditUserGroup','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditUserGroup' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditUserGroup','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolesForGroups' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolesForGroups','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolesForGroups' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolesForGroups','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRoleToGroup' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRoleToGroup','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRoleToGroup' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRoleToGroup','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRoleUser' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRoleUser','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRoleUser' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRoleUser','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRoleGroup' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRoleGroup','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRoleGroup' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRoleGroup','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroupUser' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroupUser','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroupUser' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroupUser','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroupRole' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroupRole','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroupRole' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroupRole','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroup' and per_param='desc' and per_value='Delete master group'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroup','desc','Delete master group')"
    field.Fname      = "permissions - DeleteGroup description"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroup' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroup','permissionGroup','Master')"
    field.Fname      = "permissions - DeleteGroup group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetPermissionsByRole' and per_param='desc' and per_value='Get all permissions for specific role'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetPermissionsByRole','desc','Get all permissions for specific role')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetPermissionsByRole' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetPermissionsByRole','permissionGroup','Master')"
    field.Fname      = "permissions - GetPermissionsByRole"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetPermissions' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetPermissions','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetPermissions' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetPermissions','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewRole' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewRole','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewRole' and per_param='permissionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewRole','permissionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
    
    //RulesetSource
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateRulesetSource' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateRulesetSource','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateRulesetSource' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateRulesetSource','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateCustomRulesetSource' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateCustomRulesetSource','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateCustomRulesetSource' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateCustomRulesetSource','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRulesetSource' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRulesetSource','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRulesetSource' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRulesetSource','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ViewAllRulesetSource' and per_param='desc' and per_value='View full rulesetSource details'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ViewAllRulesetSource','desc','View full rulesetSource details')"
    field.Fname      = "permissions - ViewAllRulesetSource desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ViewAllRulesetSource' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ViewAllRulesetSource','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - ViewAllRulesetSource group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRulesetSource' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRulesetSource','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRulesetSource' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRulesetSource','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRulesetFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRulesetFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRulesetFile' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRulesetFile','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRulesetSource' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRulesetSource','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRulesetSource' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRulesetSource','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DownloadFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DownloadFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DownloadFile' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DownloadFile','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='OverwriteDownload' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('OverwriteDownload','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='OverwriteDownload' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('OverwriteDownload','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CompareFiles' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CompareFiles','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CompareFiles' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CompareFiles','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewLinesToRuleset' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewLinesToRuleset','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewLinesToRuleset' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewLinesToRuleset','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetDetails' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetDetails','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetDetails' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetDetails','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetFileUUIDfromRulesetUUID' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetFileUUIDfromRulesetUUID','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetFileUUIDfromRulesetUUID' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetFileUUIDfromRulesetUUID','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='OverwriteRuleFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('OverwriteRuleFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='OverwriteRuleFile' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('OverwriteRuleFile','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDefaultRulesets' and per_param='desc' and per_value='Get all default rulesets'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDefaultRulesets','desc','Get all default rulesets')"
    field.Fname      = "permissions - LoadDefaultRulesets desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDefaultRulesets' and per_param='permissionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDefaultRulesets','permissionGroup','RulesetSource')"
    field.Fname      = "permissions - LoadDefaultRulesets group"
    ok = CheckField(field)
    
    //Ruleset
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRules' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRules','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRules' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRules','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleSID' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleSID','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleSID' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleSID','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRulesets' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRulesets','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRulesets' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRulesets','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRulesetRules' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRulesetRules','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRulesetRules' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRulesetRules','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRuleSelected' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRuleSelected','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRuleSelected' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRuleSelected','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleSelected' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleSelected','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleSelected' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleSelected','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleName' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleName','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleName' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleName','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRulesetAction' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRulesetAction','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRulesetAction' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRulesetAction','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleNote' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleNote','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleNote' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleNote','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRuleNote' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRuleNote','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRuleNote' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRuleNote','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRuleset' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRuleset','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRuleset' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRuleset','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToAllNodes' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToAllNodes','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToAllNodes' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToAllNodes','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRuleData' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRuleData','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRuleData' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRuleData','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewRuleset' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewRuleset','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewRuleset' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewRuleset','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyRuleset' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyRuleset','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyRuleset' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyRuleset','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllCustomRulesets' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllCustomRulesets','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllCustomRulesets' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllCustomRulesets','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SynchronizeAllRulesets' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SynchronizeAllRulesets','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SynchronizeAllRulesets' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SynchronizeAllRulesets','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRulesToCustomRuleset' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRulesToCustomRuleset','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRulesToCustomRuleset' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRulesToCustomRuleset','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReadRulesetData' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReadRulesetData','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReadRulesetData' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReadRulesetData','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveRulesetData' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveRulesetData','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveRulesetData' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveRulesetData','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateRule' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateRule','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateRule' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateRule','permissionGroup','Ruleset')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncToAll' and per_param='desc' and per_value='Sync ruleset to all nodes and all groups'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncToAll','desc','Sync ruleset to all nodes and all groups')"
    field.Fname      = "permissions - SyncToAll desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncToAll' and per_param='permissionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncToAll','permissionGroup','Ruleset')"
    field.Fname      = "permissions - SyncToAll group"
    ok = CheckField(field)
    
    //Node
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetSuricata' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetSuricata','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetSuricata' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetSuricata','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetZeek' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetZeek','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetZeek' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetZeek','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetWazuh' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetWazuh','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetWazuh' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetWazuh','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PutSuricataBPF' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PutSuricataBPF','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PutSuricataBPF' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PutSuricataBPF','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllNodes' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllNodes','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllNodes' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllNodes','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceStatus' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceStatus','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployService' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployService','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeFile' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeFile','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetNodeFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetNodeFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetNodeFile' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetNodeFile','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllFiles' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllFiles','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllFiles' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllFiles','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunSuricata' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunSuricata','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunSuricata' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunSuricata','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricata' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricata','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricata' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricata','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunZeek' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunZeek','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunZeek' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunZeek','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopZeek' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopZeek','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopZeek' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopZeek','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunWazuh' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunWazuh','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunWazuh' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunWazuh','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopWazuh' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopWazuh','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopWazuh' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopWazuh','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPorts' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPorts','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPorts' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPorts','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowPorts' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowPorts','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowPorts' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowPorts','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeletePorts' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeletePorts','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeletePorts' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeletePorts','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteAllPorts' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteAllPorts','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteAllPorts' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteAllPorts','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeStatus' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeStatus','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetMainconfData' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMainconfData','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetMainconfData' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMainconfData','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingAnalyzer' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingAnalyzer','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingAnalyzer' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingAnalyzer','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeAnalyzerStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeAnalyzerStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeAnalyzerStatus' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeAnalyzerStatus','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='Deploy' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Deploy','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='Deploy' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Deploy','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CheckDeploy' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CheckDeploy','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='CheckDeploy' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CheckDeploy','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowValues' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowValues','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowValues' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowValues','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDataflowValues' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDataflowValues','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDataflowValues' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDataflowValues','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValues' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValues','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValues' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValues','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValuesSelected' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValuesSelected','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValuesSelected' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValuesSelected','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNetworkInterface' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNetworkInterface','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNetworkInterface' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNetworkInterface','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetwork' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetwork','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetwork' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetwork','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNewLocal' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNewLocal','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNewLocal' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNewLocal','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveVxLAN' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveVxLAN','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveVxLAN' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveVxLAN','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SocketToNetworkList' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SocketToNetworkList','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SocketToNetworkList' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SocketToNetworkList','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetworkSelected' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetworkSelected','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetworkSelected' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetworkSelected','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteDataFlowValueSelected' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteDataFlowValueSelected','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteDataFlowValueSelected' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteDataFlowValueSelected','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeMonitor' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeMonitor','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeMonitor' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeMonitor','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginService' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginService','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeServiceStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeServiceStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeServiceStatus' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeServiceStatus','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMainServiceStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMainServiceStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMainServiceStatus' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMainServiceStatus','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteService' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteService','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
		
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapService' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapService','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapService' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapService','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyNodeOptionValues' and per_param='desc' and per_value='Modify node option values'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyNodeOptionValues','desc','Modify node option values')"
    field.Fname      = "permissions - ModifyNodeOptionValues desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyNodeOptionValues' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyNodeOptionValues','permissionGroup','Node')"
    field.Fname      = "permissions - ModifyNodeOptionValues group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSurictaRulesetSelected' and per_param='desc' and per_value='Modify ruleset for specific Suricata'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSurictaRulesetSelected','desc','Modify ruleset for specific Suricata')"
    field.Fname      = "permissions - SaveSurictaRulesetSelected desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSurictaRulesetSelected' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSurictaRulesetSelected','permissionGroup','Node')"
    field.Fname      = "permissions - SaveSurictaRulesetSelected group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RegisterNode' and per_param='desc' and per_value='Register pending node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RegisterNode','desc','Register pending node')"
    field.Fname      = "permissions - RegisterNode desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='RegisterNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RegisterNode','permissionGroup','Node')"
    field.Fname      = "permissions - RegisterNode group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingWazuhFiles' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingWazuhFiles','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingWazuhFiles' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingWazuhFiles','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteWazuhFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteWazuhFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteWazuhFile' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteWazuhFile','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddWazuhFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddWazuhFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddWazuhFile' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddWazuhFile','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadFileLastLines' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadFileLastLines','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadFileLastLines' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadFileLastLines','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContentWazuh' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContentWazuh','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContentWazuh' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContentWazuh','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadFilesData' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadFilesData','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadFilesData' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadFilesData','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddMonitorFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMonitorFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddMonitorFile' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMonitorFile','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingMonitorFiles' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingMonitorFiles','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingMonitorFiles' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingMonitorFiles','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteMonitorFile' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteMonitorFile','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteMonitorFile' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteMonitorFile','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeZeekMode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeZeekMode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeZeekMode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeZeekMode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddClusterValue' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddClusterValue','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddClusterValue' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddClusterValue','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingCluster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingCluster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingCluster' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingCluster','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditClusterValue' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditClusterValue','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditClusterValue' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditClusterValue','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteClusterValue' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteClusterValue','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteClusterValue' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteClusterValue','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncCluster' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncCluster','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncCluster' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncCluster','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControlNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControlNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControlNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControlNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidentsNode' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidentsNode','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidentsNode' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidentsNode','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeSuricataTable' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeSuricataTable','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeSuricataTable' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeSuricataTable','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToAllGroupNodes' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToAllGroupNodes','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToAllGroupNodes' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToAllGroupNodes','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAnalyzerToAllGroupNodes' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAnalyzerToAllGroupNodes','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAnalyzerToAllGroupNodes' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAnalyzerToAllGroupNodes','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StartSuricataMainConf' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StartSuricataMainConf','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StartSuricataMainConf' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StartSuricataMainConf','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricataMainConf' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricataMainConf','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricataMainConf' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricataMainConf','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='KillSuricataMainConf' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('KillSuricataMainConf','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='KillSuricataMainConf' and per_param='permissionGroup'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('KillSuricataMainConf','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadSuricataMainConf' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadSuricataMainConf','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadSuricataMainConf' and per_param='permissionGroup'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadSuricataMainConf','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LaunchZeekMainConf' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LaunchZeekMainConf','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='LaunchZeekMainConf' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LaunchZeekMainConf','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeRotationStatus' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeRotationStatus','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeRotationStatus' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeRotationStatus','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRotation' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRotation','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRotation' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRotation','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceCommands' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceCommands','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceCommands' and per_param='permissionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceCommands','permissionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)

    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////
    //////////////////////////////////////////////

    	//Analyzer
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingAnalyzer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingAnalyzer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingAnalyzer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingAnalyzer','permissionGroup','Analyzer')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeAnalyzerStatus' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeAnalyzerStatus','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeAnalyzerStatus' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeAnalyzerStatus','permissionGroup','Analyzer')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAnalyzer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAnalyzer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAnalyzer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAnalyzer','permissionGroup','Analyzer')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Autentication
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddUserFromMaster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUserFromMaster','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddUserFromMaster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUserFromMaster','permissionGroup','Autentication')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddRolesFromMaster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRolesFromMaster','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddRolesFromMaster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRolesFromMaster','permissionGroup','Autentication')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupFromMaster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupFromMaster','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupFromMaster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupFromMaster','permissionGroup','Autentication')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddUserGroupRolesFromMaster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUserGroupRolesFromMaster','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddUserGroupRolesFromMaster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUserGroupRolesFromMaster','permissionGroup','Autentication')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRolePermissions' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRolePermissions','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRolePermissions' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRolePermissions','permissionGroup','Autentication')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncPermissions' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncPermissions','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncPermissions' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncPermissions','permissionGroup','Autentication')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //ChangeControl
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControlNode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControlNode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControlNode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControlNode','permissionGroup','ChangeControl')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //ChangeControl
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','permissionGroup','ChangeControl')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopCollector' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopCollector','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopCollector' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopCollector','permissionGroup','ChangeControl')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ShowCollector' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowCollector','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ShowCollector' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowCollector','permissionGroup','ChangeControl')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //DataFlow
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowValues' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowValues','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowValues' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowValues','permissionGroup','DataFlow')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDataflowValues' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDataflowValues','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDataflowValues' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDataflowValues','permissionGroup','DataFlow')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetwork' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetwork','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetwork' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetwork','permissionGroup','DataFlow')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNewLocal' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNewLocal','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNewLocal' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNewLocal','permissionGroup','DataFlow')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveVxLAN' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveVxLAN','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveVxLAN' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveVxLAN','permissionGroup','DataFlow')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetworkSelected' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetworkSelected','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetworkSelected' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetworkSelected','permissionGroup','DataFlow')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteDataFlowValueSelected' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteDataFlowValueSelected','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteDataFlowValueSelected' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteDataFlowValueSelected','permissionGroup','DataFlow')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Deploy
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployNode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployNode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployNode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployNode','permissionGroup','Deploy')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='Deploy' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Deploy','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='Deploy' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Deploy','permissionGroup','Deploy')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //File
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SendFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SendFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SendFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SendFile','permissionGroup','File')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFile','permissionGroup','File')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllFiles' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllFiles','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllFiles' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllFiles','permissionGroup','File')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceNode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceNode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceNode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceNode','permissionGroup','File')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Group
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncSuricataGroupValues' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncSuricataGroupValues','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncSuricataGroupValues' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncSuricataGroupValues','permissionGroup','Group')"
        field.Fname      = "permissions"
        ok = CheckField(field)

        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetGroupSelectedRulesets' and per_param='desc' and per_value='Select rulesets for expert suricata in group'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetGroupSelectedRulesets','desc','Select rulesets for expert suricata in group')"
        field.Fname      = "permissions - GetGroupSelectedRulesets description"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetGroupSelectedRulesets' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetGroupSelectedRulesets','permissionGroup','Group')"
        field.Fname      = "permissions - GetGroupSelectedRulesets group"
        ok = CheckField(field)

        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddRulesetsToGroup' and per_param='desc' and per_value='Add rulesets to Groups for Suricata expert mode'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRulesetsToGroup','desc','Add rulesets to Groups for Suricata expert mode')"
        field.Fname      = "permissions - AddRulesetsToGroup description"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddRulesetsToGroup' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRulesetsToGroup','permissionGroup','Group')"
        field.Fname      = "permissions - AddRulesetsToGroup group"
        ok = CheckField(field)

        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteExpertGroupRuleset' and per_param='desc' and per_value='Delete ruleset from Suricata expert mode at groups'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteExpertGroupRuleset','desc','Delete ruleset from Suricata expert mode at groups')"
        field.Fname      = "permissions - DeleteExpertGroupRuleset description"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteExpertGroupRuleset' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteExpertGroupRuleset','permissionGroup','Group')"
        field.Fname      = "permissions - DeleteExpertGroupRuleset group"
        ok = CheckField(field)

        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncGroupRuleset' and per_param='desc' and per_value='Synchronize group ruleset to all nodes in this group'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncGroupRuleset','desc','Synchronize group ruleset to all nodes in this group')"
        field.Fname      = "SyncGroupRuleset permissions - description "
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncGroupRuleset' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncGroupRuleset','permissionGroup','Group')"
        field.Fname      = "SyncGroupRuleset permissions - group permission"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SuricataGroupService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SuricataGroupService','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SuricataGroupService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SuricataGroupService','permissionGroup','Group')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetMD5files' and per_param='desc' and per_value='Verify Master files Synchronized to all nodes'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMD5files','desc','Verify Master files Synchronized to all nodes')"
        field.Fname      = "permissions - GetMD5files description"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetMD5files' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMD5files','permissionGroup','Group')"
        field.Fname      = "permissions - GetMD5files permission group"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllGroups' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllGroups','desc','-')"
        field.Fname      = "permissions - GetAllGroups"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllGroups' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllGroups','permissionGroup','Group')"
        field.Fname      = "permissions - GetAllGroups"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddCluster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddCluster','desc','-')"
        field.Fname      = "permissions - AddCluster"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddCluster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddCluster','permissionGroup','Group')"
        field.Fname      = "permissions - AddCluster"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditGroup' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditGroup','desc','-')"
        field.Fname      = "permissions - EditGroup"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditGroup' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditGroup','permissionGroup','Group')"
        field.Fname      = "permissions - EditGroup"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupElement' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupElement','desc','-')"
        field.Fname      = "permissions - AddGroupElement"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupElement' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupElement','permissionGroup','Group')"
        field.Fname      = "permissions - AddGroupElement"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllNodesGroup' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllNodesGroup','desc','-')"
        field.Fname      = "permissions - GetAllNodesGroup"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllNodesGroup' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllNodesGroup','permissionGroup','Group')"
        field.Fname      = "permissions - GetAllNodesGroup"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupNodes' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupNodes','desc','-')"
        field.Fname      = "permissions - AddGroupNodes"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupNodes' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupNodes','permissionGroup','Group')"
        field.Fname      = "permissions - AddGroupNodes"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingGroupNodes' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingGroupNodes','desc','-')"
        field.Fname      = "permissions - PingGroupNodes"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingGroupNodes' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingGroupNodes','permissionGroup','Group')"
        field.Fname      = "permissions - PingGroupNodes"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeValues' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeValues','desc','-')"
        field.Fname      = "permissions - GetNodeValues"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeValues' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeValues','permissionGroup','Group')"
        field.Fname      = "permissions - GetNodeValues"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNodeGroup' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNodeGroup','desc','-')"
        field.Fname      = "permissions - DeleteNodeGroup"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNodeGroup' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNodeGroup','permissionGroup','Group')"
        field.Fname      = "permissions - DeleteNodeGroup"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeGroupRuleset' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeGroupRuleset','desc','-')"
        field.Fname      = "permissions - ChangeGroupRuleset"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeGroupRuleset' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeGroupRuleset','permissionGroup','Group')"
        field.Fname      = "permissions - ChangeGroupRuleset"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePathsGroups' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePathsGroups','desc','-')"
        field.Fname      = "permissions - ChangePathsGroups"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePathsGroups' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePathsGroups','permissionGroup','Group')"
        field.Fname      = "permissions - ChangePathsGroups"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncPathGroup' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncPathGroup','desc','-')"
        field.Fname      = "permissions - SyncPathGroup"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncPathGroup' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncPathGroup','permissionGroup','Group')"
        field.Fname      = "permissions - SyncPathGroup"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateGroupService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateGroupService','desc','-')"
        field.Fname      = "permissions - UpdateGroupService"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateGroupService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateGroupService','permissionGroup','Group')"
        field.Fname      = "permissions - UpdateGroupService"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAll' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAll','desc','-')"
        field.Fname      = "permissions - SyncAll"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAll' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAll','permissionGroup','Group')"
        field.Fname      = "permissions - SyncAll"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetClusterFiles' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetClusterFiles','desc','-')"
        field.Fname      = "permissions - GetClusterFiles"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetClusterFiles' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetClusterFiles','permissionGroup','Group')"
        field.Fname      = "permissions - GetClusterFiles"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteCluster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteCluster','desc','-')"
        field.Fname      = "permissions - DeleteCluster"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteCluster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteCluster','permissionGroup','Group')"
        field.Fname      = "permissions - DeleteCluster"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeClusterValue' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeClusterValue','desc','-')"
        field.Fname      = "permissions - ChangeClusterValue"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeClusterValue' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeClusterValue','permissionGroup','Group')"
        field.Fname      = "permissions - ChangeClusterValue"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetClusterFileContent' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetClusterFileContent','desc','-')"
        field.Fname      = "permissions - GetClusterFileContent"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetClusterFileContent' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetClusterFileContent','permissionGroup','Group')"
        field.Fname      = "permissions - GetClusterFileContent"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveClusterFileContent' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveClusterFileContent','desc','-')"
        field.Fname      = "permissions - SaveClusterFileContent"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveClusterFileContent' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveClusterFileContent','permissionGroup','Group')"
        field.Fname      = "permissions - SaveClusterFileContent"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncClusterFileGroup' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncClusterFileGroup','desc','-')"
        field.Fname      = "permissions - SyncClusterFileGroup"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncClusterFileGroup' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncClusterFileGroup','permissionGroup','Group')"
        field.Fname      = "permissions - SyncClusterFileGroup"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAllGroupCluster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAllGroupCluster','desc','-')"
        field.Fname      = "permissions - SyncAllGroupCluster"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAllGroupCluster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAllGroupCluster','permissionGroup','Group')"
        field.Fname      = "permissions - SyncAllGroupCluster"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAllSuricataGroup' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAllSuricataGroup','desc','-')"
        field.Fname      = "permissions - SyncAllSuricataGroup"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAllSuricataGroup' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAllSuricataGroup','permissionGroup','Group')"
        field.Fname      = "permissions - SyncAllSuricataGroup"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SuricataNodesStatus' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SuricataNodesStatus','desc','-')"
        field.Fname      = "permissions - SuricataNodesStatus"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SuricataNodesStatus' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SuricataNodesStatus','permissionGroup','Group')"
        field.Fname      = "permissions - SuricataNodesStatus"
        ok = CheckField(field)
    
        //Incidents
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidentsNode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidentsNode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidentsNode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidentsNode','permissionGroup','Incidents')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PutIncidentNode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PutIncidentNode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PutIncidentNode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PutIncidentNode','permissionGroup','Incidents')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Monitor
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddMonitorFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMonitorFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddMonitorFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMonitorFile','permissionGroup','Monitor')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingMonitorFiles' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingMonitorFiles','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingMonitorFiles' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingMonitorFiles','permissionGroup','Monitor')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteMonitorFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteMonitorFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteMonitorFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteMonitorFile','permissionGroup','Monitor')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeRotationStatus' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeRotationStatus','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeRotationStatus' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeRotationStatus','permissionGroup','Monitor')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditRotation' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRotation','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditRotation' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRotation','permissionGroup','Monitor')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Net
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetNetworkData' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNetworkData','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetNetworkData' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNetworkData','permissionGroup','Net')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValuesSelected' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValuesSelected','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValuesSelected' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValuesSelected','permissionGroup','Net')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNetworkInterface' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNetworkInterface','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNetworkInterface' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNetworkInterface','permissionGroup','Net')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Ping
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNodeData' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNodeData','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNodeData' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNodeData','permissionGroup','Ping')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingService','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingService','permissionGroup','Ping')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployService','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployService','permissionGroup','Ping')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetMainconfData' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMainconfData','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetMainconfData' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMainconfData','permissionGroup','Ping')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsNode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsNode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsNode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsNode','permissionGroup','Ping')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNodeInformation' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNodeInformation','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNodeInformation' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNodeInformation','permissionGroup','Ping')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNode','permissionGroup','Ping')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Plugin
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeServiceStatus' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeServiceStatus','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeServiceStatus' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeServiceStatus','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMainServiceStatus' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMainServiceStatus','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMainServiceStatus' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMainServiceStatus','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteService','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteService','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginService','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginService','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateSuricataValue' and per_param='desc' and per_value='Update suricata value'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateSuricataValue','desc','Update suricata value')"
        field.Fname      = "permissions - UpdateSuricataValue description"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateSuricataValue' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateSuricataValue','permissionGroup','Plugin')"
        field.Fname      = "permissions - UpdateSuricataValue group"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapService','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapService','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapService' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapService','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapService' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapService','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValues' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValues','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValues' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValues','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeSuricataTable' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeSuricataTable','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeSuricataTable' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeSuricataTable','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceCommands' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceCommands','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceCommands' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceCommands','permissionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Ports
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingPorts' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPorts','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingPorts' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPorts','permissionGroup','Ports')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ShowPorts' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowPorts','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ShowPorts' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowPorts','permissionGroup','Ports')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMode','permissionGroup','Ports')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeStatus' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeStatus','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeStatus' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeStatus','permissionGroup','Ports')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeletePorts' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeletePorts','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeletePorts' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeletePorts','permissionGroup','Ports')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteAllPorts' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteAllPorts','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteAllPorts' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteAllPorts','permissionGroup','Ports')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Zeek
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetZeek' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetZeek','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetZeek' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetZeek','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='Set' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Set','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='Set' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Set','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunZeek' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunZeek','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunZeek' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunZeek','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopZeek' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopZeek','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopZeek' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopZeek','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeZeekMode' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeZeekMode','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeZeekMode' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeZeekMode','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddClusterValue' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddClusterValue','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddClusterValue' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddClusterValue','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingCluster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingCluster','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingCluster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingCluster','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditClusterValue' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditClusterValue','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditClusterValue' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditClusterValue','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteClusterValue' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteClusterValue','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteClusterValue' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteClusterValue','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncCluster' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncCluster','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncCluster' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncCluster','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SavePolicyFiles' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SavePolicyFiles','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SavePolicyFiles' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SavePolicyFiles','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncClusterFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncClusterFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncClusterFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncClusterFile','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LaunchZeekMainConf' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LaunchZeekMainConf','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LaunchZeekMainConf' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LaunchZeekMainConf','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','permissionGroup','Zeek')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Wazuh
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetWazuh' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetWazuh','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetWazuh' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetWazuh','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunWazuh' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunWazuh','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunWazuh' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunWazuh','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopWazuh' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopWazuh','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopWazuh' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopWazuh','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingWazuhFiles' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingWazuhFiles','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingWazuhFiles' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingWazuhFiles','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteWazuhFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteWazuhFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteWazuhFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteWazuhFile','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddWazuhFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddWazuhFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddWazuhFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddWazuhFile','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadFileLastLines' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadFileLastLines','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadFileLastLines' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadFileLastLines','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContentWazuh' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContentWazuh','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContentWazuh' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContentWazuh','permissionGroup','Wazuh')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        //Suricata
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SetBPF' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetBPF','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SetBPF' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetBPF','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveConfigFile' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveConfigFile','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveConfigFile' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveConfigFile','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunSuricata' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunSuricata','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunSuricata' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunSuricata','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricata' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricata','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricata' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricata','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetSuricataServices' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetSuricataServices','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetSuricataServices' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetSuricataServices','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StartSuricataMainConf' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StartSuricataMainConf','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StartSuricataMainConf' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StartSuricataMainConf','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricataMainConf' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricataMainConf','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricataMainConf' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricataMainConf','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='KillSuricataMainConf' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('KillSuricataMainConf','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='KillSuricataMainConf' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('KillSuricataMainConf','permissionGroup','Suricata')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        // field.Fconn      = "masterConn"
        // field.Ftable     = "permissions"
        // field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadSuricataMainConf' and per_param='desc' and per_value='-'"
        // field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadSuricataMainConf','desc','-')"
        // field.Fname      = "permissions"
        // ok = CheckField(field)
        // if !ok {return false}
        // field.Fconn      = "masterConn"
        // field.Ftable     = "permissions"
        // field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadSuricataMainConf' and per_param='permissionGroup''"
        // field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadSuricataMainConf','permissionGroup','Suricata')"
        // field.Fname      = "permissions"
        // ok = CheckField(field)
    
        //Stap
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddServer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddServer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddServer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddServer','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllServers' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllServers','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllServers' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllServers','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetServer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetServer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServer','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingStap' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingStap','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingStap' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingStap','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunStap' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStap','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunStap' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStap','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStap' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStap','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStap' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStap','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapServer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapServer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapServer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapServer','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServer','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingServerStap' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServerStap','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingServerStap' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServerStap','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapServer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapServer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapServer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapServer','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapServer' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapServer','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapServer' and per_param='permissionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapServer','permissionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)

    return true
}