package scheduler

import (
	"time"
	"owlhmaster/utils"
	"owlhmaster/rulesetSource"
	"owlhmaster/node"
	"strconv"
	"owlhmaster/database"
	"github.com/astaxie/beego/logs"
)

func Init() {
	schedulerConf := map[string]map[string]string{}
	schedulerConf["scheduler"] = map[string]string{}
	schedulerConf["scheduler"]["minutes"] = ""
	schedulerConf["scheduler"]["status"] = ""
	schedulerConf,_ = utils.GetConf(schedulerConf)
	minutes := schedulerConf["scheduler"]["minutes"]
	status := schedulerConf["scheduler"]["status"]

	for status == "enabled"{
		RunScheduler()
		for {			
			time.Sleep(time.Second*60)
			_, currentMinutes, _ := time.Now().Clock()
			confMinutes,_ := strconv.Atoi(minutes)

			if currentMinutes % confMinutes == 0 {
				break
			}
		}
		logs.Info("Scheduler Running")
	}
}

//update task if their time is out
func RunScheduler() bool {	
	t := time.Now().Unix()
	currentTime := strconv.FormatInt(t, 10)
	tasks,err := CheckTasks()	
	if err != nil {
		logs.Error("Error RunScheduler checking tasks: %s", err.Error())
	}
	for j,k := range tasks {
		if k["nextEpoch"] <= currentTime{
			err = TaskUpdater(k)
			if err != nil {
				logs.Error("Error RunScheduler TaskUpdater: %s", err.Error())	
				continue
			}else{
			//calculate next epoch
			dbTime,_ := strconv.Atoi(k["period"])
			nextEpoch,_ := strconv.Atoi(k["nextEpoch"])
			s := strconv.Itoa(nextEpoch + dbTime)

			//update next epoch
			err = ndb.UpdateScheduler(j, "nextEpoch", s)
			if err != nil {
				logs.Error("Error RunScheduler UpdateScheduler updating next EPOCH time: %s", err.Error())	
				continue
			}
			logs.Notice("EPOCH updated")
			}
		}
	}
	return true
}

//return all the enabled tasks
func CheckTasks()(tasksEnabled map[string]map[string]string, err error){
	tasks,err := ndb.GetAllScheduler()
	if err != nil {
		logs.Error("Error CheckTasks GetAllScheduler: %s", err.Error())
		return nil,err
	}
	if len(tasks) == 0 { return nil,nil }
	enabledTasks := make(map[string]map[string]string)
	for x,y := range tasks{
		if tasks[x]["status"] == "enabled"{
			values := make(map[string]string)
			for y,_ := range y{
					values[y] = tasks[x][y]
			}
			enabledTasks[x] = values
		}
	}
	return enabledTasks,nil
}

func SchedulerTask(content map[string]string)(err error){
	taskUUID,err := ndb.GetSchedulerByValue(content["uuid"])
	if taskUUID == "" {
		timeEpoch,err := utils.EpochTime(content["year"]+"-"+content["month"]+"-"+content["day"]+"T"+content["hour"]+":"+content["minute"]+":00")
		logs.Warn(timeEpoch)
		if err != nil {
			logs.Error("Error RunScheduler epoch time: %s", err.Error())
			return err
		}
		newUUID := utils.Generate()
		err = ndb.InsertScheduler(newUUID, "type", content["type"])
		err = ndb.InsertScheduler(newUUID, "update", content["update"])
		err = ndb.InsertScheduler(newUUID, "period", content["period"])
		err = ndb.InsertScheduler(newUUID, "uuid", content["uuid"])
		err = ndb.InsertScheduler(newUUID, "nextEpoch", strconv.FormatInt(timeEpoch, 10))
		err = ndb.InsertScheduler(newUUID, "status", content["status"])		
		if err != nil {
			logs.Error("Error SchedulerTask TaskUpdater after first update: %s", err.Error())
			return err
		}
		logs.Notice("Task added")
	}else{
		err = ndb.UpdateScheduler(taskUUID, "status", "enabled")
		logs.Notice("Task updated")
	}
	return nil
}

func StopTask(content map[string]string)(err error){
	taskUUID,err := ndb.GetSchedulerByValue(content["uuid"])
	err = ndb.UpdateScheduler(taskUUID, "status", "disabled")
	if err != nil {
		logs.Error("Error StopTask UpdateScheduler: %s", err.Error())
		return err
	}
	return nil
}

func TaskUpdater(content map[string]string)(err error){
	data,err := ndb.GetRulesFromRuleset(content["uuid"])
	for x := range data{
		values,err := ndb.GetRuleFilesByUniqueid(x)
		if err != nil {
			logs.Error("TimeSchedule Error GetRuleFilesByUniqueid values: %s", err)
			// break
			return err
		}
		for y := range values{
			sourceFile,err := ndb.GetRuleFilesByUniqueid(values[y]["sourceFileUUID"])
			if err != nil {
				logs.Error("TimeSchedule Error GetRuleFilesByUniqueid sourceFile: %s", err)
				// break
				return err
			}
			for z := range sourceFile{
				rulesetMap := make(map[string]string)
				sourceUUIDValue,err := ndb.GetRuleFilesValue(z,"sourceUUID")
				if err != nil {
					logs.Error("TimeSchedule Error GetRuleFilesValue sourceUUIDValue: %s", err)
					// break
					return err
				}
				finalData,err := ndb.GetAllDataRulesetDB(sourceUUIDValue)
				if err != nil {
					logs.Error("TimeSchedule Error GetAllDataRulesetDB finalData: %s", err)
					// break
					return err
				}
				for a,b := range finalData{
					for b,_ := range b {
						rulesetMap[b] = finalData[a][b]
					}
				}
				if rulesetMap["isDownloaded"] == "false"{
					err = rulesetSource.DownloadFile(rulesetMap)
					if err != nil {
						logs.Error("TimeSchedule Error Downloading: %s", err)
						// break
						return err
					}
				}else if rulesetMap["isDownloaded"] == "true"{
					err = rulesetSource.OverwriteDownload(rulesetMap)
					if err != nil {
						logs.Error("TimeSchedule Error Downloading: %s", err)
						// break
						return err
					}							
				}	
			}
		}				
	}
	//overwrite files for this ruleset
	for d := range data{
		if content["update"] == "overwrite" {
			err = rulesetSource.OverwriteRuleFile(d)
			if err != nil {
				logs.Error("TimeSchedule Error OverwriteRuleFile ruleset: %s", err)
				// break
				return err
			}
		}else if content["update"] == "add-lines" {
			err = rulesetSource.AddNewLinesToRuleset(d)
			if err != nil {
				logs.Error("TimeSchedule Error AddNewLinesToRuleset ruleset: %s", err)
				// break
				return err
			}
		}

	}

	//synchronize
	err = node.SyncRulesetToAllNodes(content)
	if err != nil {
		logs.Error("TimeSchedule Error synchronizing ruleset: %s", err)
		return err
	}

	logs.Notice("Ruleset synchronized "+content["uuid"])	
	return nil
}