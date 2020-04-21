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
    field.Fquery     = "select per_value from permissions where per_uniqueid='delete' and per_param='permisionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('delete','permisionGroup','admin')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='post' and per_param='permisionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('post','permisionGroup','admin')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='put' and per_param='permisionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('put','permisionGroup','admin')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='get' and per_param='permisionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('get','permisionGroup','admin')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='admin' and per_param='permisionGroup' and per_value='admin'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('admin','permisionGroup','admin')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControl' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControl','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddServer' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddServer','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllServers' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllServers','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServer' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServer','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStap' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStap','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStap' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStap','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStap' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStap','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapServer' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapServer','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServer' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServer','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingServerStap' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServerStap','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapServer' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapServer','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapServerput' and per_param='permisionGroup' and per_value='ChangeControl'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapServerput','permisionGroup','ChangeControl')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='permisionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','permisionGroup','Collector')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopCollector' and per_param='permisionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopCollector','permisionGroup','Collector')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowCollector' and per_param='permisionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowCollector','permisionGroup','Collector')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayMasterCollector' and per_param='permisionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayMasterCollector','permisionGroup','Collector')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopMasterCollector' and per_param='permisionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopMasterCollector','permisionGroup','Collector')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowMasterCollector' and per_param='permisionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowMasterCollector','permisionGroup','Collector')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='permisionGroup' and per_value='Collector'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','permisionGroup','Collector')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddStapServer' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddStapServer','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllStapServers' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllStapServers','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStapServer' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStapServer','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetStapStatus' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetStapStatus','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapService' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapService','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapService' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapService','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapService' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapService','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingStapServer' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingStapServer','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapService' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapService','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapService' and per_param='permisionGroup' and per_value='Stap'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapService','permisionGroup','Stap')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRulesetsBySearch' and per_param='permisionGroup' and per_value='Search'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRulesetsBySearch','permisionGroup','Search')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SchedulerTask' and per_param='permisionGroup' and per_value='Scheduler'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SchedulerTask','permisionGroup','Scheduler')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopTask' and per_param='permisionGroup' and per_value='Scheduler'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopTask','permisionGroup','Scheduler')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetLog' and per_param='permisionGroup' and per_value='Scheduler'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetLog','permisionGroup','Scheduler')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetFileContent' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetFileContent','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContent' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContent','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPlugins' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPlugins','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingFlow' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingFlow','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePluginStatus' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePluginStatus','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveStapInterface' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveStapInterface','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowStatus' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowStatus','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNetworkInterface' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNetworkInterface','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateMasterNetworkInterface' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateMasterNetworkInterface','permisionGroup','Master')"
    field.Fname      = "permissions - group"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadMasterNetworkValuesSelected' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadMasterNetworkValuesSelected','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingServiceMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServiceMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployServiceMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployServiceMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginServiceMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginServiceMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteServiceMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteServiceMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValuesMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValuesMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateMasterStapInterface' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateMasterStapInterface','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetBPF' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetBPF','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapServiceMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapServiceMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidents' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidents','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveZeekValues' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveZeekValues','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsMaster' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsMaster','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetPathFileContent' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetPathFileContent','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFilePathContent' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFilePathContent','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddUser' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUser','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllUsers' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllUsers','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUser' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUser','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupUsers' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupUsers','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRole' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRole','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolesForUser' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolesForUser','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetGroupsForUser' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetGroupsForUser','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddUsersTo' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUsersTo','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangePassword' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangePassword','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUserRole' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUserRole','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteUserGroup' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteUserGroup','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRoles' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRoles','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRole' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRole','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRole' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRole','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllUserGroups' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllUserGroups','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditUserGroup' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditUserGroup','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolesForGroups' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolesForGroups','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRoleToGroup' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRoleToGroup','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRoleUser' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRoleUser','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRoleGroup' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRoleGroup','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroupUser' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroupUser','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteGroupRole' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteGroupRole','permisionGroup','Master')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolePermissions' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolePermissions','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRolePermissions' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRolePermissions','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewRole' and per_param='permisionGroup' and per_value='Master'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewRole','permisionGroup','Master')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateRulesetSource' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateRulesetSource','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateCustomRulesetSource' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateCustomRulesetSource','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRulesetSource' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRulesetSource','permisionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRulesetSource' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRulesetSource','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRulesetFile' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRulesetFile','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRulesetSource' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRulesetSource','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DownloadFile' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DownloadFile','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='OverwriteDownload' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('OverwriteDownload','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='CompareFiles' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CompareFiles','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewLinesToRuleset' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewLinesToRuleset','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetDetails' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetDetails','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetFileUUIDfromRulesetUUID' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetFileUUIDfromRulesetUUID','permisionGroup','RulesetSource')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='OverwriteRuleFile' and per_param='permisionGroup' and per_value='RulesetSource'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('OverwriteRuleFile','permisionGroup','RulesetSource')"
    field.Fname      = "permissions - group"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRules' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRules','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleSID' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleSID','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRulesets' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRulesets','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRulesetRules' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRulesetRules','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRuleSelected' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRuleSelected','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleSelected' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleSelected','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleName' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleName','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRulesetAction' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRulesetAction','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetRuleNote' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetRuleNote','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetRuleNote' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetRuleNote','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteRuleset' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteRuleset','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToAllNodes' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToAllNodes','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllRuleData' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllRuleData','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddNewRuleset' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddNewRuleset','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyRuleset' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyRuleset','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllCustomRulesets' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllCustomRulesets','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SynchronizeAllRulesets' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SynchronizeAllRulesets','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddRulesToCustomRuleset' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRulesToCustomRuleset','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReadRulesetData' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReadRulesetData','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveRulesetData' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveRulesetData','permisionGroup','Ruleset')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateRule' and per_param='permisionGroup' and per_value='Ruleset'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateRule','permisionGroup','Ruleset')"
    field.Fname      = "permissions - group"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='CreateNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CreateNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetSuricata' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetSuricata','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetZeek' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetZeek','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetWazuh' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetWazuh','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PutSuricataBPF' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PutSuricataBPF','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllNodes' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllNodes','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceStatus' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceStatus','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployService' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployService','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeFile' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeFile','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SetNodeFile' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetNodeFile','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllFiles' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllFiles','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunSuricata' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunSuricata','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricata' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricata','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunZeek' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunZeek','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopZeek' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopZeek','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='RunWazuh' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunWazuh','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopWazuh' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopWazuh','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPorts' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPorts','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ShowPorts' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowPorts','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeletePorts' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeletePorts','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteAllPorts' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteAllPorts','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeStatus' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeStatus','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetMainconfData' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMainconfData','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingAnalyzer' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingAnalyzer','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeAnalyzerStatus' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeAnalyzerStatus','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='Deploy' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Deploy','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='CheckDeploy' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('CheckDeploy','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowValues' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowValues','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDataflowValues' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDataflowValues','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValues' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValues','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValuesSelected' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValuesSelected','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNetworkInterface' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNetworkInterface','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetwork' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetwork','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNewLocal' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNewLocal','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveVxLAN' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveVxLAN','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SocketToNetworkList' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SocketToNetworkList','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetworkSelected' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetworkSelected','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteDataFlowValueSelected' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteDataFlowValueSelected','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetNodeMonitor' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNodeMonitor','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginService' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginService','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeServiceStatus' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeServiceStatus','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMainServiceStatus' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMainServiceStatus','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteService' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteService','permisionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSuricataInterface' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSuricataInterface','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSuricataInterface' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSuricataInterface','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapService' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapService','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceNode','permisionGroup','Node')"
    field.Fname      = "permissions - group"
    ok = CheckField(field)
	
	field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValues' and per_param='desc' and per_value='-'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValues','desc','-')"
    field.Fname      = "permissions - desc"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "permissions"
    field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValues' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValues','permisionGroup','Node')"
    field.Fname      = "permissions - group"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingWazuhFiles' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingWazuhFiles','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteWazuhFile' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteWazuhFile','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddWazuhFile' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddWazuhFile','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='LoadFileLastLines' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadFileLastLines','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContentWazuh' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContentWazuh','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadFilesData' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadFilesData','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddMonitorFile' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMonitorFile','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingMonitorFiles' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingMonitorFiles','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteMonitorFile' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteMonitorFile','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeZeekMode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeZeekMode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='AddClusterValue' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddClusterValue','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='PingCluster' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingCluster','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditClusterValue' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditClusterValue','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteClusterValue' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteClusterValue','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncCluster' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncCluster','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControlNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControlNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidentsNode' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidentsNode','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeSuricataTable' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeSuricataTable','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRulesetToAllGroupNodes' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRulesetToAllGroupNodes','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAnalyzerToAllGroupNodes' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAnalyzerToAllGroupNodes','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StartSuricataMainConf' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StartSuricataMainConf','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricataMainConf' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricataMainConf','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='KillSuricataMainConf' and per_param='permisionGroup'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('KillSuricataMainConf','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadSuricataMainConf' and per_param='permisionGroup'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadSuricataMainConf','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='LaunchZeekMainConf' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LaunchZeekMainConf','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeRotationStatus' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeRotationStatus','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='EditRotation' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRotation','permisionGroup','Node')"
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
    field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceCommands' and per_param='permisionGroup' and per_value='Node'"
    field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceCommands','permisionGroup','Node')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingAnalyzer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingAnalyzer','permisionGroup','Analyzer')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeAnalyzerStatus' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeAnalyzerStatus','permisionGroup','Analyzer')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncAnalyzer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncAnalyzer','permisionGroup','Analyzer')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddUserFromMaster' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUserFromMaster','permisionGroup','Autentication')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddRolesFromMaster' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddRolesFromMaster','permisionGroup','Autentication')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddGroupFromMaster' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddGroupFromMaster','permisionGroup','Autentication')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddUserGroupRolesFromMaster' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddUserGroupRolesFromMaster','permisionGroup','Autentication')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncRolePermissions' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncRolePermissions','permisionGroup','Autentication')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncPermissions' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncPermissions','permisionGroup','Autentication')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetChangeControlNode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetChangeControlNode','permisionGroup','ChangeControl')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PlayCollector' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PlayCollector','permisionGroup','ChangeControl')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopCollector' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopCollector','permisionGroup','ChangeControl')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ShowCollector' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowCollector','permisionGroup','ChangeControl')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeDataflowValues' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeDataflowValues','permisionGroup','DataFlow')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadDataflowValues' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadDataflowValues','permisionGroup','DataFlow')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetwork' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetwork','permisionGroup','DataFlow')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNewLocal' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNewLocal','permisionGroup','DataFlow')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveVxLAN' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveVxLAN','permisionGroup','DataFlow')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSocketToNetworkSelected' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSocketToNetworkSelected','permisionGroup','DataFlow')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteDataFlowValueSelected' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteDataFlowValueSelected','permisionGroup','DataFlow')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployNode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployNode','permisionGroup','Deploy')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='Deploy' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Deploy','permisionGroup','Deploy')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SendFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SendFile','permisionGroup','File')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFile','permisionGroup','File')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllFiles' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllFiles','permisionGroup','File')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServiceNode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServiceNode','permisionGroup','File')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncSuricataGroupValues' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncSuricataGroupValues','permisionGroup','Group')"
        field.Fname      = "permissions"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SuricataGroupService' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SuricataGroupService','permisionGroup','Group')"
        field.Fname      = "permissions"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetIncidentsNode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetIncidentsNode','permisionGroup','Incidents')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PutIncidentNode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PutIncidentNode','permisionGroup','Incidents')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddMonitorFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddMonitorFile','permisionGroup','Monitor')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingMonitorFiles' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingMonitorFiles','permisionGroup','Monitor')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteMonitorFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteMonitorFile','permisionGroup','Monitor')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeRotationStatus' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeRotationStatus','permisionGroup','Monitor')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditRotation' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditRotation','permisionGroup','Monitor')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetNetworkData' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetNetworkData','permisionGroup','Net')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadNetworkValuesSelected' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadNetworkValuesSelected','permisionGroup','Net')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNetworkInterface' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNetworkInterface','permisionGroup','Net')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='UpdateNodeData' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('UpdateNodeData','permisionGroup','Ping')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingService' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingService','permisionGroup','Ping')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployService' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployService','permisionGroup','Ping')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetMainconfData' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetMainconfData','permisionGroup','Ping')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingPluginsNode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPluginsNode','permisionGroup','Ping')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveNodeInformation' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveNodeInformation','permisionGroup','Ping')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteNode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteNode','permisionGroup','Ping')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeServiceStatus' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeServiceStatus','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMainServiceStatus' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMainServiceStatus','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteService' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteService','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddPluginService' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddPluginService','permisionGroup','Plugin')"
        field.Fname      = "permissions"
        ok = CheckField(field)
    
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSuricataInterface' and per_param='desc' and per_value='-'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSuricataInterface','desc','-')"
        field.Fname      = "permissions"
        ok = CheckField(field)
        if !ok {return false}
        field.Fconn      = "masterConn"
        field.Ftable     = "permissions"
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveSuricataInterface' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveSuricataInterface','permisionGroup','Plugin')"
        field.Fname      = "permissions"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeployStapService' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeployStapService','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapService' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapService','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ModifyStapValues' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ModifyStapValues','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeSuricataTable' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeSuricataTable','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetServiceCommands' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServiceCommands','permisionGroup','Plugin')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingPorts' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingPorts','permisionGroup','Ports')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ShowPorts' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ShowPorts','permisionGroup','Ports')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeMode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeMode','permisionGroup','Ports')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeStatus' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeStatus','permisionGroup','Ports')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeletePorts' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeletePorts','permisionGroup','Ports')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteAllPorts' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteAllPorts','permisionGroup','Ports')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetZeek' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetZeek','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='Set' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('Set','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunZeek' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunZeek','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopZeek' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopZeek','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='ChangeZeekMode' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ChangeZeekMode','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddClusterValue' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddClusterValue','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingCluster' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingCluster','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditClusterValue' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditClusterValue','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteClusterValue' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteClusterValue','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncCluster' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncCluster','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SavePolicyFiles' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SavePolicyFiles','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncClusterFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncClusterFile','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='LaunchZeekMainConf' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LaunchZeekMainConf','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SyncZeekValues' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SyncZeekValues','permisionGroup','Zeek')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetWazuh' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetWazuh','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunWazuh' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunWazuh','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopWazuh' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopWazuh','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingWazuhFiles' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingWazuhFiles','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteWazuhFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteWazuhFile','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddWazuhFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddWazuhFile','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='LoadFileLastLines' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('LoadFileLastLines','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveFileContentWazuh' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveFileContentWazuh','permisionGroup','Wazuh')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SetBPF' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SetBPF','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='SaveConfigFile' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('SaveConfigFile','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunSuricata' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunSuricata','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricata' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricata','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetSuricataServices' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetSuricataServices','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StartSuricataMainConf' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StartSuricataMainConf','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopSuricataMainConf' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopSuricataMainConf','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='KillSuricataMainConf' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('KillSuricataMainConf','permisionGroup','Suricata')"
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
        // field.Fquery     = "select per_value from permissions where per_uniqueid='ReloadSuricataMainConf' and per_param='permisionGroup''"
        // field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('ReloadSuricataMainConf','permisionGroup','Suricata')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='AddServer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('AddServer','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetAllServers' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetAllServers','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='GetServer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('GetServer','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingStap' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingStap','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunStap' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStap','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStap' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStap','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='RunStapServer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('RunStapServer','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='StopStapServer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('StopStapServer','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='PingServerStap' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('PingServerStap','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='DeleteStapServer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('DeleteStapServer','permisionGroup','Stap')"
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
        field.Fquery     = "select per_value from permissions where per_uniqueid='EditStapServer' and per_param='permisionGroup'"
        field.Finsert    = "insert into permissions (per_uniqueid,per_param,per_value) values ('EditStapServer','permisionGroup','Stap')"
        field.Fname      = "permissions"
        ok = CheckField(field)

    return true
}