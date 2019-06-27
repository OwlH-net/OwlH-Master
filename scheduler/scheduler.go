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
		logs.Info("LOOP")
		RunScheduler()
		for {			
			time.Sleep(time.Second*60)
			_, currentMinutes, _ := time.Now().Clock()
			confMinutes,_ := strconv.Atoi(minutes)

			if currentMinutes % confMinutes == 0 {
				logs.Warn("Mod == 0")
				break
			}
		}
		logs.Info("Scheduler Running")
	}
}

//update task if their time is out
func RunScheduler() bool {	
	// location,_ := time.LoadLocation("Europe/Rome")
	location,_ := time.LoadLocation("Europe/Madrid")
	t1 := time.Now().In(location).Unix()
	t2 := strconv.FormatInt(t1, 10)
	logs.Debug(t2)


	t := time.Now().Unix()
	currentTime := strconv.FormatInt(t, 10)
	tasks,err := CheckTasks()	
	if err != nil {
		logs.Error("Error RunScheduler checking tasks: %s", err.Error())
	}
	for j,k := range tasks {
		logs.Info(k["nextEpoch"]+" -- "+currentTime)
		if k["nextEpoch"] <= currentTime{
			err = TaskUpdater(k)
			if err != nil {
				logs.Error("Error RunScheduler TaskUpdater: %s", err.Error())	
			}

			//calculate next epoch
			dbTime,_ := strconv.Atoi(k["period"])
			nextEpoch,_ := strconv.Atoi(k["nextEpoch"])
			s := strconv.Itoa(nextEpoch + dbTime)

			//update next epoch
			err = ndb.UpdateScheduler(j, "nextEpoch", s)

			logs.Notice("EPOCH updated")
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
	logs.Info(taskUUID)
	if taskUUID == "" {
		timeEpoch,err := utils.EpochTime(content["year"]+"-"+content["month"]+"-"+content["day"]+"T"+content["hour"]+":"+content["minute"]+":00.000Z")
		logs.Notice(strconv.FormatInt(timeEpoch, 10))
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
		// err = TaskUpdater(content)
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
			break
		}
		for y := range values{
			sourceFile,err := ndb.GetRuleFilesByUniqueid(values[y]["sourceFileUUID"])
			if err != nil {
				logs.Error("TimeSchedule Error GetRuleFilesByUniqueid sourceFile: %s", err)
				break
			}
			for z := range sourceFile{
				rulesetMap := make(map[string]string)
				sourceUUIDValue,err := ndb.GetRuleFilesValue(z,"sourceUUID")
				if err != nil {
					logs.Error("TimeSchedule Error GetRuleFilesValue sourceUUIDValue: %s", err)
					break
				}
				finalData,err := ndb.GetAllDataRulesetDB(sourceUUIDValue)
				if err != nil {
					logs.Error("TimeSchedule Error GetAllDataRulesetDB finalData: %s", err)
					break
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
						break
					}
				}else if rulesetMap["isDownloaded"] == "true"{
					err = rulesetSource.OverwriteDownload(rulesetMap)
					if err != nil {
						logs.Error("TimeSchedule Error Downloading: %s", err)
						break
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
				break
			}
		}else if content["update"] == "add-lines" {
			err = rulesetSource.AddNewLinesToRuleset(d)
			if err != nil {
				logs.Error("TimeSchedule Error AddNewLinesToRuleset ruleset: %s", err)
				break
			}
		}

	}

	//synchronize
	err = node.SyncRulesetToAllNodes(content)
	if err != nil {
		logs.Error("TimeSchedule Error synchronizing ruleset: %s", err)
	}

	logs.Notice("Ruleset synchronized "+content["uuid"])	
	return nil
}



// func TimeSchedule(content map[string]string)(err error) {
// 	secondsLeft,err := CheckTimeDiff(content)	
// 	if err != nil {
// 		logs.Error("TimeSchedule failed checking time difference: %s", err)
// 		return err
// 	}
// 	exists,err := CheckScheduleExists(content["uuid"])
// 	if err != nil {
// 		logs.Error("TimeSchedule failed retrieving data dataSchedule: %s", err)
// 		return err
// 	}
// 	if !exists {
// 		err = ndb.InsertRulesetSchedule(content["uuid"], "secondsLeft", secondsLeft)
// 		if err != nil {
// 			logs.Error("TimeSchedule failed inserting time difference: %s", err)
// 			return err
// 		}
// 		for key := range content {
// 			err = ndb.InsertRulesetSchedule(content["uuid"], key, content[key])
// 		}
// 	}else{
// 		err = UpdateSchedule(content["uuid"], "secondsLeft", secondsLeft)
// 		if err != nil {
// 			logs.Error("TimeSchedule failed inserting time difference: %s", err)
// 			return err
// 		}
// 	}

// 	for RunSchedule() {
// 		logs.Warn("Executing Schedule")
// 	}
		
// 	// //First update
// 	// if secondsLeft >= 0{
// 	// 	time.Sleep(time.Duration(secondsLeft) * time.Second)
// 	// 	err = TickerUpdate(dataSchedule)
// 	// 	if err != nil {
// 	// 		logs.Error("TimeSchedule failed At first update: %s", err)
// 	// 		return err
// 	// 	}
// 	// }else {
// 	// 	return errors.New("User's date and time is older than the current time. Can't update...")
// 	// }
	
// 	// t, err := strconv.Atoi(dataSchedule["schedule"])
// 	// // ticker = time.NewTicker(time.Duration(t) * time.Minute)
// 	// ticker = time.NewTicker(15 * time.Second)
	
// 	// go func() {
// 	// 	for e := range ticker.C {
// 	// 		logs.Debug(c)
// 	// 		logs.Info(dataSchedule)
// 	// 		err = TickerUpdate(dataSchedule)
// 	// 		if err != nil {
// 	// 			logs.Error("TimeSchedule Error Updating into a Ticker: %s", err)
// 	// 			break
// 	// 		}
// 	// 	}
// 	// }()
		
// 	return err
// }

// func RunSchedule()(status bool){
// 	ruleset := Ruleset{}
// 	identifier := UUID{}
// 	list := ListOfRulesetsSchedule{}
// 	content,err := ndb.GetScheduleByUniqueid(content["uuid"])
// 	for x := range content{
// 		if content[x]["status"] == "enabled" {
// 			identifier.uuid = x
// 			ruleset.hour = content[x]["hour"]
// 			ruleset.minute = content[x]["minute"]
// 			ruleset.day = content[x]["day"]
// 			ruleset.month = content[x]["month"]
// 			ruleset.year = content[x]["year"]
// 			ruleset.update = content[x]["update"]
// 			ruleset.schedule = content[x]["schedule"]
// 			ruleset.secondsLeft = content[x]["secondsLeft"]
// 			identifier.ruleset = ruleset
// 			list.uuid = identifier
// 		}
// 	}
// }

// func StopTimeSchedule(content map[string]string)(err error){
// 	logs.Error(content["uuid"])
//    	ticker.Stop()

// 	logs.Error("Timer stopped")
// 	logs.Error("Timer stopped")
// 	logs.Error("Timer stopped")
    
// 	return nil
// }



// func CheckTimeDiff(content map[string]string)(t int, err error){
// 	dt := time.Now()
// 	currentYear := dt.Year()
// 	currentMonth := dt.Month()
// 	currentDay := dt.Day()
// 	currentHour := dt.Hour() 
// 	currentMinute := dt.Minute()
// 	currentSecond := dt.Second()
	
// 	year,err := strconv.Atoi(content["year"])
// 	month,err := strconv.Atoi(content["month"])
// 	day,err := strconv.Atoi(content["day"])
// 	hour,err := strconv.Atoi(content["hour"])
// 	minute,err := strconv.Atoi(content["minute"])
// 	dUser := time.Date(year,time.Month(month),day,hour,minute, 0, 0, time.Local)
// 	dCurrent := time.Date(currentYear,time.Month(currentMonth),currentDay,currentHour,currentMinute, currentSecond, 0, time.Local)
// 	diff := dUser.Sub(dCurrent)
// 	secondsLeft := int(diff.Seconds())

// 	if err != nil {
// 		logs.Error("CheckTimeDiff Error strconv to Integer: %s", err)
// 		return 0,err
// 	}

// 	return secondsLeft,nil
// }