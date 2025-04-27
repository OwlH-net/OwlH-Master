package ndb

import (
  "database/sql"
  "errors"
  "os"

  "github.com/OwlH-net/OwlH-Master/utils"
  "github.com/astaxie/beego/logs"
  _ "github.com/mattn/go-sqlite3"
)

var (
  Rdb *sql.DB
)

func RConn() {
  var err error
  path, err := utils.GetKeyValueString("rulesetConn", "path")
  if err != nil {
    logs.Error("RConn Error getting data from main.conf at master: " + err.Error())
  }
  cmd, err := utils.GetKeyValueString("rulesetConn", "cmd")
  if err != nil {
    logs.Error("RConn Error getting data from main.conf at master: " + err.Error())
  }

  _, err = os.Stat(path)
  if err != nil {
    panic("ruleset/Ruleset DB -- DB Open Failed: " + err.Error())
  }
  Rdb, err = sql.Open(cmd, path)
  if err != nil {
    logs.Error("ruleset/Ruleset DB -- SQL openning Error: " + err.Error())
  }
  logs.Info("ruleset/Ruleset DB -- DB -> sql.Open, DB Ready")
}

func RulesetSourceKeyInsert(nkey string, key string, value string) (err error) {
  if Rdb == nil {
    logs.Error("no access to database")
    return errors.New("no access to database")
  }
  stmt, err := Rdb.Prepare("insert into ruleset (ruleset_uniqueid, ruleset_param, ruleset_value) values (?,?,?);")
  if err != nil {
    logs.Error("Prepare -> %s", err.Error())
    return err
  }
  _, err = stmt.Exec(&nkey, &key, &value)
  if err != nil {
    logs.Error("Execute -> %s", err.Error())
    return err
  }
  return nil
}

func InsertRulesetSourceRules(nkey string, key string, value string) (err error) {
  if Rdb == nil {
    logs.Error("no access to database")
    return errors.New("no access to database")
  }
  stmt, err := Rdb.Prepare("insert into rule_files (rule_uniqueid, rule_param, rule_value) values(?,?,?)")
  if err != nil {
    logs.Error("Prepare -> %s", err.Error())
    return err
  }

  _, err = stmt.Exec(&nkey, &key, &value)
  if err != nil {
    logs.Error("Execute -> %s", err.Error())
    return err
  }
  return nil
}

func UpdateRuleset(uuid string, param string, value string) (err error) {
  updateRulesetNode, err := Rdb.Prepare("update ruleset set ruleset_value = ? where ruleset_uniqueid = ? and ruleset_param = ?;")
  if err != nil {
    logs.Error("UpdateRuleset UPDATE prepare error for update isDownloaded -- " + err.Error())
    return err
  }
  _, err = updateRulesetNode.Exec(&value, &uuid, &param)
  defer updateRulesetNode.Close()
  if err != nil {
    logs.Error("UpdateRuleset UPDATE error for update isDownloaded -- " + err.Error())
    return err
  }
  return nil
}

func UpdateRuleFiles(uuid string, param string, value string) (err error) {
  updateRulesetNode, err := Rdb.Prepare("update rule_files set rule_value = ? where rule_uniqueid = ? and rule_param = ?;")
  defer updateRulesetNode.Close()
  if err != nil {
    logs.Error("UpdateRuleFiles UPDATE prepare error for update isDownloaded -- " + err.Error())
    // defer updateRulesetNode.Close()
    return err
  }
  // defer updateRulesetNode.Close()
  _, err = updateRulesetNode.Exec(&value, &uuid, &param)
  if err != nil {
    logs.Error("UpdateRuleFiles UPDATE error for update isDownloaded -- " + err.Error())
    // defer updateRulesetNode.Close()
    return err
  }
  return nil
}

func GetRulesetSourceValue(uuid string, param string) (val string, err error) {
  var value string
  sql := "select ruleset_value from ruleset where ruleset_uniqueid='" + uuid + "' and ruleset_param = '" + param + "';"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("Rdb.Query Error : %s", err.Error())
    return "", err
  }
  for rows.Next() {
    if err = rows.Scan(&value); err != nil {
      logs.Error("GetRulesetSourcePath rows.Scan: %s", err.Error())
      return "", err
    }
  }
  return value, nil
}

func GetAllCustomRulesetDB() (path []string, err error) {
  var customData []string
  var uniqid string

  sql := "select ruleset_uniqueid from ruleset where ruleset_value='custom' and ruleset_param='sourceType'"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetAllCustomRuleset Rdb.Query Error : %s", err.Error())
    return nil, err
  }
  for rows.Next() {
    if err = rows.Scan(&uniqid); err != nil {
      logs.Error("GetAllCustomRuleset -- Query return error: %s", err.Error())
      return nil, err
    }
    customData = append(customData, uniqid)
  }
  return customData, nil
}

func GetAllLocalRulesetDB() (path []string, err error) {
  var customData []string
  var uniqid string

  sql := "select ruleset_uniqueid from ruleset where ruleset_value='local' and ruleset_param='type'"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetAllLocalRulesetDB Rdb.Query Error : %s", err.Error())
    return nil, err
  }
  for rows.Next() {
    if err = rows.Scan(&uniqid); err != nil {
      logs.Error("GetAllLocalRulesetDB -- Query return error: %s", err.Error())
      return nil, err
    }
    customData = append(customData, uniqid)
  }
  return customData, nil
}

func GetAllDataRulesetDB(uuid string) (data map[string]map[string]string, err error) {
  var customData = map[string]map[string]string{}
  var uniqid string
  var param string
  var value string

  sql := "select ruleset_uniqueid, ruleset_param, ruleset_value from ruleset where ruleset_uniqueid='" + uuid + "'"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetAllDataRulesetDB Rdb.Query Error : %s", err.Error())
    return nil, err
  }
  for rows.Next() {
    if err = rows.Scan(&uniqid, &param, &value); err != nil {
      logs.Error("GetAllDataRulesetDB -- Query return error: %s", err.Error())
      return nil, err
    }
    if customData[uniqid] == nil {
      customData[uniqid] = map[string]string{}
    }
    customData[uniqid][param] = value
  }
  return customData, nil
}

func GetAllRulesets() (data map[string]map[string]string, err error) {
  var customData = map[string]map[string]string{}
  var uniqid string
  var param string
  var value string

  sql := "select ruleset_uniqueid, ruleset_param, ruleset_value from ruleset"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetAllRulesets Rdb.Query Error : %s", err.Error())
    return nil, err
  }
  for rows.Next() {
    if err = rows.Scan(&uniqid, &param, &value); err != nil {
      logs.Error("GetAllRulesets -- Query return error: %s", err.Error())
      return nil, err
    }
    if customData[uniqid] == nil {
      customData[uniqid] = map[string]string{}
    }
    customData[uniqid][param] = value
  }
  return customData, nil
}

// Get a specific ruleset path
func GetRulesetPath(uuid string) (n string, err error) {
  var path string
  if Rdb != nil {
    row := Rdb.QueryRow("SELECT rule_value FROM rule_files WHERE rule_uniqueid=$1 and rule_param=\"path\";", uuid)
    err = row.Scan(&path)

    if err == sql.ErrNoRows {
      logs.Error("DB RULESET -> There is no ruleset with id %s", uuid)
      return "", errors.New("DB RULESET -> There is no ruleset with id " + uuid)
    }
    if err != nil {
      logs.Error("DB RULESET -> rows.Scan Error -> %s", err.Error())
      return "", errors.New("DB RULESET -> -> rows.Scan Error -> " + err.Error())
    }
    return path, nil
  } else {
    logs.Info("DB RULESET -> No access to database")
    return "", errors.New("DB RULESET -> No access to database")
  }
}

func GetRuleFilesValue(uuid string, param string) (path string, err error) {
  var value string
  sql := "select rule_value from rule_files where rule_uniqueid='" + uuid + "' and rule_param = '" + param + "';"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("Rdb.Query Error : %s", err.Error())
    return "", err
  }
  for rows.Next() {
    if err = rows.Scan(&value); err != nil {
      logs.Error("GetRulesetSourcePath rows.Scan: %s", err.Error())
      return "", err
    }
  }
  return value, nil
}

func GetRulesFromRuleset(uuid string) (data map[string]map[string]string, err error) {
  var allRuleDetails = map[string]map[string]string{}
  var uniqid string
  var param string
  var value string
  var uuidSource string
  if Rdb == nil {
    logs.Error("no access to database")
    return nil, errors.New("no access to database")
  }
  sqlUUID := "select rule_uniqueid from rule_files where rule_param='sourceUUID' and rule_value = '" + uuid + "';"
  uuidRows, err := Rdb.Query(sqlUUID)
  if err != nil {
    logs.Error("Rdb.Query Error checking uuid for take the uuid list for GetDetails: %s", err.Error())
    return nil, err
  }
  defer uuidRows.Close()
  for uuidRows.Next() {
    if err = uuidRows.Scan(&uuidSource); err != nil {
      logs.Error("GetDetails UUIDSource uuidRows.Scan: %s", err.Error())
      return nil, err
    }
    sql := "select rule_uniqueid, rule_param, rule_value from rule_files where rule_uniqueid='" + uuidSource + "';"
    rows, err := Rdb.Query(sql)
    if err != nil {
      logs.Error("Rdb.Query Error : %s", err.Error())
      return nil, err
    }
    defer rows.Close()
    for rows.Next() {
      if err = rows.Scan(&uniqid, &param, &value); err != nil {
        logs.Error("GetDetails rows.Scan: %s", err.Error())
        return nil, err
      }
      if allRuleDetails[uniqid] == nil {
        allRuleDetails[uniqid] = map[string]string{}
      }
      allRuleDetails[uniqid][param] = value
    }
  }
  return allRuleDetails, nil
}

func GetRuleFilesByUniqueid(uuid string) (data map[string]map[string]string, err error) {
  var allRuleDetails = map[string]map[string]string{}
  var uniqid string
  var param string
  var value string
  sql := "select rule_uniqueid, rule_param, rule_value from rule_files where rule_uniqueid='" + uuid + "';"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("Rdb.Query Error : %s", err.Error())
    return nil, err
  }
  defer rows.Close()
  for rows.Next() {
    if err = rows.Scan(&uniqid, &param, &value); err != nil {
      logs.Error("GetDetails rows.Scan: %s", err.Error())
      return nil, err
    }
    if allRuleDetails[uniqid] == nil {
      allRuleDetails[uniqid] = map[string]string{}
    }
    allRuleDetails[uniqid][param] = value
  }
  return allRuleDetails, nil
}

func DeleteRulesetByUniqueid(uuid string) (err error) {
  deleteRulesetQuery, err := Rdb.Prepare("delete from ruleset where ruleset_uniqueid = ?;")
  _, err = deleteRulesetQuery.Exec(&uuid)
  defer deleteRulesetQuery.Close()
  if err != nil {
    logs.Error("DB DeleteRulese/deleteRulesetQueryt -> ERROR on table Ruleset...")
    return errors.New("DB DeleteRuleset/deleteRulesetQuery -> ERROR on table Ruleset...")
  }
  return nil
}

func DeleteRulesetNodeByUniqueid(uuid string) (err error) {
  deleteRulesetNodeQuery, err := Rdb.Prepare("delete from ruleset_node where ruleset_uniqueid = ?;")
  _, err = deleteRulesetNodeQuery.Exec(&uuid)
  defer deleteRulesetNodeQuery.Close()
  if err != nil {
    logs.Error("DB DeleteRuleset/deleteRulesetNodeQuery -> ERROR on table Ruleset_node...")
    return errors.New("DB DeleteRuleset/deleteRulesetNodeQuery -> ERROR on table Ruleset_node...")
  }
  return nil
}

func DeleteRulesetNodeByNode(uuid string) (err error) {
  deleteRulesetNodeQuery, err := Rdb.Prepare("delete from ruleset_node where node_uniqueid = ?;")
  _, err = deleteRulesetNodeQuery.Exec(&uuid)
  defer deleteRulesetNodeQuery.Close()
  if err != nil {
    logs.Error("DB DeleteRuleset/deleteRulesetNodeQuery -> ERROR on table Ruleset_node...")
    return errors.New("DB DeleteRuleset/deleteRulesetNodeQuery -> ERROR on table Ruleset_node...")
  }
  return nil
}

func DeleteRuleFilesByUuid(uuid string) (err error) {
  DeleteUUIDPrepare, err := Rdb.Prepare("delete from rule_files where rule_uniqueid = ?")
  if err != nil {
    logs.Error("DeleteRulese Rdb.Query Error deleting by rule_uniqueid for rule_files: %s", err.Error())
    return err
  }
  _, err = DeleteUUIDPrepare.Exec(&uuid)
  if err != nil {
    logs.Error("DeleteRulese deleting a ruleset source -> %s", err.Error())
    return err
  }
  return nil
}

func DeleteRuleFileParamValue(param string, value string) (err error) {
  DeleteUUIDPrepare, err := Rdb.Prepare("delete from rule_files where rule_param = ? and rule_value = ?")
  if err != nil {
    logs.Error("DeleteRulese Rdb.Query Error deleting by rule_uniqueid for rule_files: %s", err.Error())
    return err
  }
  _, err = DeleteUUIDPrepare.Exec(&param, &value)
  if err != nil {
    logs.Error("DeleteRulese deleting a ruleset source -> %s", err.Error())
    return err
  }
  return nil
}

func GetAllRuleFiles() (data map[string]map[string]string, err error) {
  var allRuleDetails = map[string]map[string]string{}
  var uniqid string
  var param string
  var value string
  sql := "select rule_uniqueid, rule_param, rule_value from rule_files;"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetAllRuleFiles Rdb.Query Error: %s", err.Error())
    return nil, err
  }
  defer rows.Close()
  for rows.Next() {
    if err = rows.Scan(&uniqid, &param, &value); err != nil {
      logs.Error("GetAllRuleFiles rows.Scan error: %s", err.Error())
      return nil, err
    }
    if allRuleDetails[uniqid] == nil {
      allRuleDetails[uniqid] = map[string]string{}
    }
    allRuleDetails[uniqid][param] = value
  }
  return allRuleDetails, nil
}

func GetRulesetUUID(node string) (uuid string, err error) {
  sql := "select ruleset_uniqueid from ruleset_node where node_uniqueid='" + node + "';"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetRulesetUUID Rdb.Query Error: %s", err.Error())
    return "", err
  }
  defer rows.Close()
  if rows.Next() {
    if err = rows.Scan(&uuid); err != nil {
      logs.Error("GetRulesetUUID rows.Scan error: %s", err.Error())
      return "", err
    }
  }
  return uuid, nil
}

func GetNodeWithRulesetUUID(ruleset string) (data []string, err error) {
  var customData []string
  var uniqid string

  sql := "select node_uniqueid from ruleset_node where ruleset_uniqueid='" + ruleset + "'"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetNodeWithRulesetUUID Rdb.Query Error : %s", err.Error())
    return nil, err
  }
  for rows.Next() {
    if err = rows.Scan(&uniqid); err != nil {
      logs.Error("GetNodeWithRulesetUUID -- Query return error: %s", err.Error())
      return nil, err
    }
    customData = append(customData, uniqid)
  }
  return customData, nil
}

func InsertGroupRulesets(uuid string, param string, value string) (err error) {
  if Rdb == nil {
    logs.Error("no access to database")
    return err
  }
  insertValues, err := Rdb.Prepare("insert into grouprulesets(gr_uniqueid, gr_param, gr_value) values(?,?,?);")
  if err != nil {
    logs.Error("Prepare InsertGroupRulesets-> %s", err.Error())
    return err
  }

  _, err = insertValues.Exec(&uuid, &param, &value)
  if err != nil {
    logs.Error("Execute InsertGroupRulesets-> %s", err.Error())
    return err
  }

  return nil
}

func GetAllGroupRulesets() (groups map[string]map[string]string, err error) {
  var allgrouprulesets = map[string]map[string]string{}
  var uniqid string
  var param string
  var value string
  if Rdb == nil {
    logs.Error("GetAllGroupRulesets no access to database")
    return nil, err
  }

  sql := "select gr_uniqueid, gr_param, gr_value from grouprulesets;"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("GetAllGroupRulesets Rdb.Query Error : %s", err.Error())
    return nil, err
  }

  for rows.Next() {
    if err = rows.Scan(&uniqid, &param, &value); err != nil {
      logs.Error("GetAllGroupRulesets rows.Scan: %s", err.Error())
      return nil, err
    }

    if allgrouprulesets[uniqid] == nil {
      allgrouprulesets[uniqid] = map[string]string{}
    }
    allgrouprulesets[uniqid][param] = value
  }
  return allgrouprulesets, nil
}

func DeleteGroupRuleset(uuid string) (err error) {
  if Rdb == nil {
    logs.Error("DeleteGroupRuleset -- Can't acces to database")
    return err
  }

  stmt, err := Rdb.Prepare("delete from grouprulesets where gr_uniqueid = ?")
  if err != nil {
    logs.Error("Prepare DeleteGroupRuleset -> %s", err.Error())
    return err
  }

  _, err = stmt.Exec(&uuid)
  if err != nil {
    logs.Error("Execute DeleteGroupRuleset -> %s", err.Error())
    return err
  }

  return nil
}

func DeleteGroupRulesetByValue(param string, value string) (err error) {
  if Rdb == nil {
    logs.Error("DeleteGroupRulesetByValue -- Can't acces to database")
    return err
  }

  stmt, err := Rdb.Prepare("delete from grouprulesets where gr_param = ? and gr_value= ? ")
  if err != nil {
    logs.Error("Prepare DeleteGroupRulesetByValue -> %s", err.Error())
    return err
  }

  _, err = stmt.Exec(&param, &value)
  if err != nil {
    logs.Error("Execute DeleteGroupRulesetByValue -> %s", err.Error())
    return err
  }

  return nil
}

func GetRuleName(uuid string) (ruleset string, err error) {
  var nameRule string
  if Rdb == nil {
    logs.Warn("GetRuleName -> no access to database")
    return "", err
  }

  sql := "SELECT ruleset_value FROM ruleset WHERE ruleset_uniqueid = \"" + uuid + "\" and ruleset_param = \"name\";"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Warn("GetRuleName -> param or param doesn't exists")
    return "", err
  }

  for rows.Next() {
    if err = rows.Scan(&nameRule); err != nil {
      logs.Error("GetRuleName rows.Scan: %s", err.Error())
      return "", err
    }
  }
  return nameRule, nil
}

func GetRuleSelected(uuid string) (ruleset string, err error) {
  var ruleSelected string
  if Rdb != nil {
    row := Rdb.QueryRow("SELECT ruleset_uniqueid FROM ruleset_node WHERE node_uniqueid = \"" + uuid + "\";")
    err = row.Scan(&ruleSelected)
    if err != nil {
      logs.Warn("GetRuleSelected -> row.Scan error %s", err.Error())
      return "", err
    }

    return ruleSelected, nil
  } else {
    logs.Error("GetRuleSelected -> No access to database")
    return "", err
  }
}

func RulesetParamExists(uuid, param string) bool {
  sql := "select ruleset_uniqueid from ruleset where ruleset_param = '" + param + "' and ruleset_uniqueid = '" + uuid + "';"
  rows, err := Rdb.Query(sql)
  if err != nil {
    logs.Error("Rdb.Query Error : %s", err.Error())
    return false
  }

  for rows.Next() {
    logs.Info("field -> %v DOES exist for ruleset -> %v", param, uuid)
    rows.Close()
    return true
  }
  logs.Info("field -> %v DOES NOT exist for ruleset -> %v", param, uuid)
  return false
}

func GetGroupRulesets(guuid string) string {
  if Rdb == nil {
    logs.Warn("GetDefaultRuleset -> no access to database")
    return ""
  }
  sql := Rdb.QueryRow("select gr_uniqueid from grouprulesets where gr_param = 'groupid' and gr_value='" + guuid + "'")
  gruuid := ""
  sql.Scan(&gruuid)
  sql = Rdb.QueryRow("select gr_value from grouprulesets where gr_param = 'rulesetid' and gr_uniqueid='" + gruuid + "'")
  ruuid := ""
  sql.Scan(&ruuid)
  return ruuid
}

func GetDefaultRuleset() string {
  if Rdb == nil {
    logs.Warn("GetDefaultRuleset -> no access to database")
    return ""
  }
  sql := Rdb.QueryRow("select ruleset_uniqueid from ruleset where ruleset_param = 'default' and ruleset_value='true'")
  ruuid := ""
  sql.Scan(&ruuid)
  return ruuid
}
