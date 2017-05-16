//Copyright 2016-2017 Beate Ottenwälder
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package service

import (
	"github.com/fsnotify/fsnotify"
	"github.com/ottenwbe/golook/broker/models"
	log "github.com/sirupsen/logrus"
	"sync"
)

/*
FileMonitor monitors files/folders and reports any change in the file system regarding those files/folders.
For reporting a Reporter callback has to be registered with the FileMonitor.
*/
type FileMonitor struct {
	watcher  *fsnotify.Watcher
	once     sync.Once
	done     chan bool
	Reporter func(string)
}

/*
Open implements
*/
func (fm *FileMonitor) Open(monitoredFiles map[string]map[string]*models.File) {
	fm.once.Do(func() {
		var err error
		fm.watcher, err = fsnotify.NewWatcher()
		if err != nil {
			log.WithError(err).Fatal("Cannot start file monitor.")
		}

		fm.done = make(chan bool)
		go cMonitor(fm)
	})

	if monitoredFiles != nil {
		for _, folders := range monitoredFiles {
			for _, file := range folders {
				fm.Monitor(file.Name)
			}
		}
	} else {
		log.Warn("No initial set of files to be monitored.")
	}
}

/*
Close ensures that files are no longer monitored.
*/
func (fm *FileMonitor) Close() {
	if fm.done != nil {
		fm.done <- true
	}
	if fm.watcher != nil {
		fm.watcher.Close()
	}
	fm.once = sync.Once{}
}

/*
cMonitor implements a handler which reacts to file system events on behalf of the file monitor fm
*/
func cMonitor(fm *FileMonitor) {
	var stop = false
	for !stop {
		select {
		case event := <-fm.watcher.Events:
			if event.Name != "" {
				log.Infof("Event %s triggered report", event.String())
				if fm.Reporter != nil {
					fm.Reporter(event.Name)
				} else {
					log.Error("Not reporting monitored file change; Reporter is nil!")
				}
			}
		case err := <-fm.watcher.Errors:
			log.WithError(err).Error("Error from file watcher")
		case stop = <-fm.done:
			log.WithField("stop", stop).Info("Stopping file monitor")
		}
	}
}

/*
Monitor registers paths to files or folders with the FileMonitor. The FileMonitor can then report changes to the fies,
respectively files in the folders.
*/
func (fm *FileMonitor) Monitor(file string) bool {
	err := fm.watcher.Add(file)
	if err != nil {
		log.WithField("file", file).WithError(err).Error("Will not monitor file or folder for changes.")
		return false
	}
	return true
}

/*
RemoveMonitored removes paths to files or folders with the FileMonitor. Changes are no longer reported.
*/
func (fm *FileMonitor) RemoveMonitored(file string) bool {
	err := fm.watcher.Remove(file)
	if err != nil {
		log.WithField("file", file).WithError(err).Error("Cannot stop monitoring file or folder.")
		return false
	}
	return true
}
