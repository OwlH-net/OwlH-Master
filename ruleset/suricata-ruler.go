package ruleset

import (
  // "bytes"
  // "encoding/binary"
  // "errors"
  // "log"
  // "net"
  // "sync"
  // "bufio"
  // "os"
  // "time"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "regexp"
  "strings"

  "github.com/OwlH-net/OwlH-Master/utils"
  "github.com/astaxie/beego/logs"
)

type ruleheader struct {
  Action    string
  Active    bool
  Proto     string
  Srcip     string
  Srcport   string
  Dstip     string
  Dstport   string
  Direction string
}

type config struct {
  keywordsFile string
  ruleFile     string
}

var rulerconfig config

type keyworddata struct {
  Validation string `json:"validation"`
  Must       bool   `json:"must"`
  Multiple   bool   `json:"multiple"`
}

type keyword map[string]keyworddata

type keywords map[string]keyword

var rulekeywords = keywords{}
var rule string
var mustkeywords = map[string]string{}

func loadkeywords() {
  byteValue, err := ioutil.ReadFile(rulerconfig.keywordsFile)
  if err != nil {
    logs.Error("readkeywords - error getting suricata rule keywords from file %s -> %s", rulerconfig.keywordsFile, err.Error())
    return
  }
  json.Unmarshal(byteValue, &rulekeywords)
  for kwgroup := range rulekeywords {
    for kword := range rulekeywords[kwgroup] {
      if rulekeywords[kwgroup][kword].Must {
        mustkeywords[kword] = ""
      }
    }
  }
}

func parsehead(head string) (success bool, report []string) {
  var rh ruleheader
  report = []string{}

  str := fmt.Sprintf("INFO -> rule header verification")
  report = append(report, str)

  str = fmt.Sprintf("INFO -> HEADER -> %s", head)
  report = append(report, str)
  success = true
  re := regexp.MustCompile(`^([^\s]+)\s+([^\s]+)\s+([^\s]+)\s+([^\s]+)\s+([^\s]+)\s+([^\s]+)\s+([^\s]+)\s+`)
  headitems := re.FindStringSubmatch(head)
  if headitems == nil {
    success = false
    str := fmt.Sprintf("ERROR -> bad rule header format -> %s", head)
    report = append(report, str)
    str = fmt.Sprintf("ERROR -> DETAIL - no enough parameters, expecting 7, we got %d, regex didn't match", len(headitems))
    report = append(report, str)
    return success, report
  }

  rh.Active = strings.Contains(headitems[1], "#")
  rh.Action = strings.Trim(headitems[1], "#")

  if !verifyAction(rh.Action) {
    success = false
    str := fmt.Sprintf("ERROR -> header ACTION -> should be pass, drop, reject or alert, but found '%s'", rh.Action)
    report = append(report, str)
  }
  rh.Proto = headitems[2]
  rh.Srcip = headitems[3]
  rh.Srcport = headitems[4]
  rh.Direction = headitems[5]
  if !verifyDirection(rh.Direction) {
    success = false
    str := fmt.Sprintf("ERROR -> header DIRECTION -> should be  '->' or '<>', but found '%s'", rh.Direction)
    report = append(report, str)
  }
  rh.Dstip = headitems[6]
  rh.Dstport = headitems[7]

  str = fmt.Sprintf("INFO -> rule header verification - DONE ")
  report = append(report, str)
  return success, report
}

func parsebody(body string) (success bool, report []string) {
  success = true
  report = []string{}
  str := fmt.Sprintf("INFO -> parsing rule body ")
  report = append(report, str)
  str = fmt.Sprintf("INFO -> BODY -> %s", body)
  report = append(report, str)

  keywordsitem := strings.Split(body, ";")
  if keywordsitem == nil {
    success = false
    str = fmt.Sprintf("ERROR -> can't find keywords-value couples")
    report = append(report, str)
    return success, report
  }

  re := regexp.MustCompile(`^([^:]+):(.*)`)
  rulekwords := map[string]string{}
  for kword := range keywordsitem {
    if len(keywordsitem[kword]) == 0 {
      continue
    }
    values := re.FindStringSubmatch(keywordsitem[kword])
    if len(values) == 3 {
      values[1] = strings.Trim(values[1], " ")
      val := strings.Trim(values[2], " ")
      rulekwords[values[1]] = val
    } else if len(values) == 2 {
      values[1] = strings.Trim(values[1], " ")
      rulekwords[values[1]] = ""
    } else {
      str = fmt.Sprintf("WARNING -> keyword-value seems to be bad formated %v", values)
      report = append(report, str)
    }
  }
  str = fmt.Sprintf("INFO -> Check Mandatory keywords")
  report = append(report, str)
  for mustkw := range mustkeywords {
    if _, ok := rulekwords[mustkw]; ok {
    } else {
      str := fmt.Sprintf("ERROR -> mandatory keyword %s is not present", mustkw)
      report = append(report, str)
      success = false
    }
  }
  str = fmt.Sprintf("INFO -> Check Mandatory keywords - DONE")
  report = append(report, str)

  str = fmt.Sprintf("INFO -> Check rule keywords against keyword dictionary")
  report = append(report, str)

  for kword := range rulekwords {
    if ok := keywordexists(kword); ok {
    } else {
      str := fmt.Sprintf("WARNING -> rule keyword '%s' is not present in current keyword map", kword)
      report = append(report, str)
      str = fmt.Sprintf("WARNING -> rule keyword '%s' value -> %s", kword, rulekwords[kword])
      report = append(report, str)
    }
  }
  str = fmt.Sprintf("INFO -> Check rule keywords against keyword dictionary - DONE")
  report = append(report, str)
  str = fmt.Sprintf("INFO -> parsing rule body - DONE")
  report = append(report, str)
  return success, report
}

func keywordexists(keyword string) (exists bool) {
  for kwgroup := range rulekeywords {
    if _, ok := rulekeywords[kwgroup][keyword]; ok {
      return true
    }
  }
  return false
}

func parserule(rule string) (success bool, report []string) {
  report = []string{}
  report = append(report, "INFO -> parsing rule -> "+rule)

  re := regexp.MustCompile(`^([^\(]+)\(([^\)]+)\)`)
  values := re.FindStringSubmatch(rule)
  if values == nil {
    report = append(report, "ERROR -> can't fine header and rule body")
    return false, report
  }

  success, localreport := parsehead(values[1])
  report = Uappend(report, localreport)
  if !success {
    return success, report
  }

  success, localreport = parsebody(values[2])
  report = Uappend(report, localreport)
  if !success {
    return success, report
  }
  return true, report
}

func Init() {
  logs.Info("Loading ruleset keywords...")
  var err error
  rulerconfig.keywordsFile, err = utils.GetKeyValueString("ruleset", "keywordsFile")
  if err != nil {
    logs.Error("Suricata-Ruler Error getting data from main.conf")
    return
  }

  loadkeywords()
}

func Uappend(out, in []string) (newout []string) {
  for str := range in {
    out = append(out, in[str])
  }
  return out
}

func verifyDirection(direction string) (ok bool) {
  switch direction {
  case "->", "<>":
    return true
  }
  return false
}

func verifyAction(action string) (ok bool) {
  switch action {
  case "alert", "pass", "drop", "reject":
    return true
  }
  return false
}
