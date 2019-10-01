package main

import (
	"fmt"
	"log"
	"time"

	// Copyright © 2019 IBM Corporation and others.
	//
	// Licensed under the Apache License, Version 2.0 (the "License");
	// you may not use this file except in compliance with the License.
	// You may obtain a copy of the License at
	//
	//     http://www.apache.org/licenses/LICENSE-2.0
	//
	// Unless required by applicable law or agreed to in writing, software
	// distributed under the License is distributed on an "AS IS" BASIS,
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	// See the License for the specific language governing permissions and
	// limitations under the License.

	//The following additions are  available under the same license as the rest of the code.

	"github.com/appsody/watcher"
)

func main() {
	w := watcher.New()

	go func() {
		for {
			select {
			case event := <-w.Event:
				// Print the event.
				fmt.Println(event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive("../test_folder"); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}
	fmt.Println()

	go func() {
		w.Wait()
		// Ignore ../test_folder/test_folder_recursive and ../test_folder/.dotfile
		if err := w.Ignore("../test_folder/test_folder_recursive", "../test_folder/.dotfile"); err != nil {
			log.Fatalln(err)
		}
		// Print a list of all of the files and folders currently being watched
		// and their paths after adding files and folders to the ignore list.
		for path, f := range w.WatchedFiles() {
			fmt.Printf("%s: %s\n", path, f.Name())
		}
		fmt.Println()
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
