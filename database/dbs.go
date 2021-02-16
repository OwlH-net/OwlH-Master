package ndb

import (
	"database/sql"
	"errors"
	"os"
	"owlhmaster/utils"

	"github.com/astaxie/beego/logs"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Db *sql.DB
)

func Close() {
	Db.Close()
	Rdb.Close()
	Gdb.Close()
	RSdb.Close()
	Mdb.Close()
}

func Conn() {
	var err error
	path, err := utils.GetKeyValueString("dbsConn", "path")
	if err != nil {
		logs.Error("Conn Error getting data from main.conf at master: " + err.Error())
	}
	cmd, err := utils.GetKeyValueString("dbsConn", "cmd")
	if err != nil {
		logs.Error("Conn Error getting data from main.conf at master: " + err.Error())
	}

	_, err = os.Stat(path)
	if err != nil {
		panic("Error: dbs/node DB -- DB Open Failed: " + err.Error())
	}
	Db, err = sql.Open(cmd, path+"?cache=shared&mode=rwc&_busy_timeout=5000")
	if err != nil {
		logs.Error("dbs/node DB -- DB Open Failed: " + err.Error())
	} else {
		logs.Info("dbs/node DB -- DB -> sql.Open, DB Ready")
	}
}

func DeleteNode(uuid string) (err error) {
	deleteNodeQuery, err := Db.Prepare("delete from nodes where node_uniqueid = ?;")
	_, err = deleteNodeQuery.Exec(&uuid)
	defer deleteNodeQuery.Close()
	if err != nil {
		logs.Error("DB DeleteRulese/deleteNodeQuery -> ERROR on table Ruleset...")
		return errors.New("DB DeleteRuleset/deleteNodeQuery -> ERROR on table Ruleset...")
	}
	return nil
}

func NodeKeyExists(nodekey string, key string) (id int, err error) {
	if Db == nil {
		logs.Error("no access to database")
		return 0, errors.New("no access to database")
	}
	sql := "SELECT node_id FROM nodes where node_uniqueid = '" + nodekey + "' and node_param = '" + key + "';"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error(err.Error())
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&id); err == nil {
			return id, err
		}
	}
	return 0, nil
}

func InsertNodeKey(nkey string, key string, value string) (err error) {
	if Db == nil {
		logs.Error("no access to database")
		return errors.New("no access to database")
	}
	stmt, err := Db.Prepare("insert into nodes (node_uniqueid, node_param, node_value) values(?,?,?);")
	if err != nil {
		logs.Error("InsertNodeKey Prepare -> %s", err.Error())
		return err
	}

	_, err = stmt.Exec(&nkey, &key, &value)
	if err != nil {
		logs.Error("InsertNodeKey Execute -> %s", err.Error())
		return err
	}
	return nil
}

func UpdateNode(uuid string, param string, value string) (err error) {
	updateNode, err := Db.Prepare("update nodes set node_value = ? where node_uniqueid = ? and node_param = ?;")
	if err != nil {
		logs.Error("updateNode UPDATE prepare error for update-- " + err.Error())
		return err
	}

	_, err = updateNode.Exec(&value, &uuid, &param)
	defer updateNode.Close()
	if err != nil {
		logs.Error("updateNode UPDATE error -- " + err.Error())
		return err
	}
	return nil
}

func GetNodeIpbyName(n string) (ip string, err error) {
	if Db == nil {
		logs.Error("no access to database")
		return "", errors.New("no access to database")
	}
	sql := "select node_value from nodes where node_uniqueid like '%" + n + "%' and node_param = 'ip';"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&ip); err == nil {
			return ip, err
		}
	}
	return "", errors.New("There is no IP for given node name")
}

func GetNodeTags() (tags map[string]map[string]string, err error) {
	var allNodeTags = map[string]map[string]string{}
	var uniqid string
	var param string
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return nil, err
	}

	sql := "select nt_uniqueid, nt_param, nt_value from nodeTags;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetNodeTags Db.Query Error : %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetNodeTags rows.Scan: %s", err.Error())
			return nil, err
		}

		if allNodeTags[uniqid] == nil {
			allNodeTags[uniqid] = map[string]string{}
		}
		allNodeTags[uniqid][param] = value
	}
	return allNodeTags, nil
}

func GetAllOrganizations() (orgs map[string]map[string]string, err error) {
	var allOrgs = map[string]map[string]string{}
	var uniqid string
	var param string
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return nil, err
	}

	sql := "select org_uniqueid, org_param, org_value from organizations;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetAllOrganizations Db.Query Error : %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetAllOrganizations rows.Scan: %s", err.Error())
			return nil, err
		}

		if allOrgs[uniqid] == nil {
			allOrgs[uniqid] = map[string]string{}
		}
		allOrgs[uniqid][param] = value
	}
	return allOrgs, nil
}

func GetNodeOrgs() (orgs map[string]map[string]string, err error) {
	var allOrgs = map[string]map[string]string{}
	var uniqid string
	var param string
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return nil, err
	}

	sql := "select no_uniqueid, no_param, no_value from nodeOrgs;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetNodeOrgs Db.Query Error : %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetNodeOrgs rows.Scan: %s", err.Error())
			return nil, err
		}

		if allOrgs[uniqid] == nil {
			allOrgs[uniqid] = map[string]string{}
		}
		allOrgs[uniqid][param] = value
	}
	return allOrgs, nil
}

func GetAllTags() (tags map[string]map[string]string, err error) {
	var allTags = map[string]map[string]string{}
	var uniqid string
	var param string
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return nil, err
	}

	sql := "select tag_uniqueid, tag_param, tag_value from tags;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetAllTags Db.Query Error : %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetAllTags rows.Scan: %s", err.Error())
			return nil, err
		}

		if allTags[uniqid] == nil {
			allTags[uniqid] = map[string]string{}
		}
		allTags[uniqid][param] = value
	}
	return allTags, nil
}

func InsertTag(nkey string, key string, value string) (err error) {
	if Db == nil {
		logs.Error("no access to database")
		return errors.New("no access to database")
	}
	stmt, err := Db.Prepare("insert into tags (tag_uniqueid, tag_param, tag_value) values(?,?,?);")
	if err != nil {
		logs.Error("InsertTag Prepare -> %s", err.Error())
		return err
	}

	_, err = stmt.Exec(&nkey, &key, &value)
	if err != nil {
		logs.Error("InsertTag Execute -> %s", err.Error())
		return err
	}
	return nil
}

func GetTagIDByName(name string) (tag string, err error) {
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return "", err
	}

	sql := "select tag_uniqueid from tags where tag_param = 'tagName' and tag_value = '" + name + "' ;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetTagIDByName Db.Query Error : %s", err.Error())
		return "", err
	}

	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			logs.Error("GetTagIDByName rows.Scan: %s", err.Error())
			return "", err
		}
	}
	return value, nil
}

func GetTagsByID(uuid string) (tag string, err error) {
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return "", err
	}

	sql := "select tag_value from tags where tag_param = 'tagName' and tag_uniqueid = '" + uuid + "' ;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetTagsByID Db.Query Error : %s", err.Error())
		return "", err
	}

	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			logs.Error("GetTagsByID rows.Scan: %s", err.Error())
			return "", err
		}
	}
	return value, nil
}

func GetOrgsByID(uuid string) (tag string, err error) {
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return "", err
	}

	sql := "select org_value from organizations where org_param = 'name' and org_uniqueid = '" + uuid + "' ;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetOrgsByID Db.Query Error : %s", err.Error())
		return "", err
	}

	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			logs.Error("GetOrgsByID rows.Scan: %s", err.Error())
			return "", err
		}
	}
	return value, nil
}

func GetOrgByID(uuid string) (tag string, err error) {
	var value string
	if Db == nil {
		logs.Error("no access to database")
		return "", err
	}

	sql := "select org_value from organizations where org_param = 'name' and org_uniqueid = '" + uuid + "' ;"
	rows, err := Db.Query(sql)
	if err != nil {
		logs.Error("GetOrgByID Db.Query Error : %s", err.Error())
		return "", err
	}

	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			logs.Error("GetOrgByID rows.Scan: %s", err.Error())
			return "", err
		}
	}
	return value, nil
}

func DeleteNodeTag(uuid string) (err error) {
	deleteNodeQuery, err := Db.Prepare("delete from nodeTags where nt_uniqueid = ?;")
	_, err = deleteNodeQuery.Exec(&uuid)
	defer deleteNodeQuery.Close()
	if err != nil {
		logs.Error("DeleteNodeTag -> ERROR on table nodeTags...")
		return errors.New("DeleteNodeTag -> ERROR on table nodeTags...")
	}
	return nil
}

func DeleteNodeOrg(uuid string) (err error) {
	deleteNodeQuery, err := Db.Prepare("delete from nodeOrgs where no_uniqueid = ?;")
	_, err = deleteNodeQuery.Exec(&uuid)
	defer deleteNodeQuery.Close()
	if err != nil {
		logs.Error("DeleteNodeOrg -> ERROR on table nodeOrgs...")
		return errors.New("DeleteNodeOrg -> ERROR on table nodeOrgs...")
	}
	return nil
}

func DeleteOrganization(uuid string) (err error) {
	deleteNodeQuery, err := Db.Prepare("delete from organizations where org_uniqueid = ?;")
	_, err = deleteNodeQuery.Exec(&uuid)
	defer deleteNodeQuery.Close()
	if err != nil {
		logs.Error("DeleteOrganization -> ERROR on table organizations...")
		return errors.New("DeleteOrganization -> ERROR on table organizations...")
	}
	return nil
}

func InsertNodeTag(nkey string, key string, value string) (err error) {
	if Db == nil {
		logs.Error("no access to database")
		return errors.New("no access to database")
	}
	stmt, err := Db.Prepare("insert into nodeTags (nt_uniqueid, nt_param, nt_value) values(?,?,?);")
	if err != nil {
		logs.Error("InsertNodeTag Prepare -> %s", err.Error())
		return err
	}

	_, err = stmt.Exec(&nkey, &key, &value)
	if err != nil {
		logs.Error("InsertNodeTag Execute -> %s", err.Error())
		return err
	}
	return nil
}

func InsertNodeOrgs(nkey string, key string, value string) (err error) {
	if Db == nil {
		logs.Error("no access to database")
		return errors.New("no access to database")
	}
	stmt, err := Db.Prepare("insert into nodeOrgs (no_uniqueid, no_param, no_value) values(?,?,?);")
	if err != nil {
		logs.Error("InsertNodeOrgs Prepare -> %s", err.Error())
		return err
	}

	_, err = stmt.Exec(&nkey, &key, &value)
	if err != nil {
		logs.Error("InsertNodeOrgs Execute -> %s", err.Error())
		return err
	}
	return nil
}

func EditOrganization(uuid string, param string, value string) (err error) {
	EditOrganization, err := Db.Prepare("update organizations set org_value = ? where org_uniqueid = ? and org_param = ?;")
	if err != nil {
		logs.Error("EditOrganization UPDATE prepare error for update-- " + err.Error())
		return err
	}

	_, err = EditOrganization.Exec(&value, &uuid, &param)
	defer EditOrganization.Close()
	if err != nil {
		logs.Error("EditOrganization UPDATE error -- " + err.Error())
		return err
	}
	return nil
}

func InsertOrganization(nkey string, key string, value string) (err error) {
	if Db == nil {
		logs.Error("no access to database")
		return errors.New("no access to database")
	}
	stmt, err := Db.Prepare("insert into organizations (org_uniqueid, org_param, org_value) values(?,?,?);")
	if err != nil {
		logs.Error("InsertOrganization Prepare -> %s", err.Error())
		return err
	}

	_, err = stmt.Exec(&nkey, &key, &value)
	if err != nil {
		logs.Error("InsertOrganization Execute -> %s", err.Error())
		return err
	}
	return nil
}
