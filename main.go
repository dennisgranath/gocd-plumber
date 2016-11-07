package main

import (
  "flag"
  log "github.com/Sirupsen/logrus"
  "github.com/go-ini/ini"
  "io/ioutil"
  "strings"
)

type Config struct {
  BaseUrl string
  Username string
  Password string
}
var config Config = getConfig()
var allTemplates map[string]interface{}
var allPipelines []interface{}
func getPipelines(filename string) {
  data, err := parseFile(filename)

  if err != nil {
    log.WithFields(log.Fields{
      "error":    err.Error(),
      "filename": filename,
    }).Error("Unable to parse pipeline file")
    return
  }
  templates := data.(map[string]interface{})["templates"]
  pipelines := data.(map[string]interface{})["pipelines"]

  if templates != nil {
    for k,v := range templates.(map[string]interface {}) {
      allTemplates[k] = v.(map[string]interface {})["pipeline"]
    }
  }
  for _, v := range pipelines.([]interface{}){
    allPipelines = append(allPipelines, v.(map[string]interface{}))
  }

}

func processPipelines() {
  for _, p := range allPipelines{
    processPipeline(p.(map[string]interface{}))
  }
}

func processPipeline(pipeline map[string]interface{}) {
  var template map[string]interface{}
  if pipeline["pipeline_template"] != nil {
    template = allTemplates[pipeline["pipeline_template"].(string)].(map[string]interface{})
  }
  pipelineConf := pipeline["pipeline"].(map[string]interface{})
  if template != nil {
    pipeline["pipeline"] = merge(template, pipelineConf)

  }
  createOrEditPipeline(pipeline)
}

func main() {

  files, _ := ioutil.ReadDir(".")
  for _, f := range files {

		if f.IsDir() {
			continue
		}
    if !strings.HasSuffix(f.Name(),".yml") {
      continue
    }
    getPipelines(f.Name())
  }
  processPipelines()

}

func getConfig() Config{
  config_file := flag.String("config", "root/etc/gocd-plumber/config.ini", "Config file path")
  debug_option := flag.Bool("debug", false, "Verbose debug output")
  flag.Parse()
  if *debug_option {
     log.SetLevel(log.DebugLevel)
  }
  inifile, err := ini.Load(*config_file)
  if err != nil {
    log.WithFields(log.Fields{
      "error": err.Error(),
    }).Fatal("Unable to open config file")
  }
  var config Config
  err = inifile.MapTo(&config)
  if err != nil {
    log.WithFields(log.Fields{
      "error": err.Error(),
    }).Fatal("Unable to parse config file")
  }
  return config
}

func init(){
  allTemplates = make(map[string]interface{})
}
