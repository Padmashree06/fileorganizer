package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main(){
	//take the source directory input and store it in dirpath 
	fmt.Println("Enter the path of the directory: ")
	var dirpath string 
	fmt.Scanln(&dirpath)

	//checks if the directory exists 
	_, err := os.Stat(dirpath)
	if os.IsNotExist(err) {
		fmt.Println("Directory does not exist.")
		return
	} else if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	//walk through the directory recursively
	err = filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil{
			fmt.Println("Error path:", err)
			return err
		}

		if !info.IsDir() {
			// Extract the file extension
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if ext == "" {
				ext = "no_extension" // Handle files without extensions
			} else {
				ext = ext[1:] // Remove the dot from the extension
			}

			// Create a folder for the extension if it doesn't exist
			extFolder := filepath.Join(dirpath, ext)
			os.MkdirAll(extFolder, os.ModePerm)

			// Move the file into the corresponding folder
			newPath := filepath.Join(extFolder, info.Name())
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Printf("Error moving file %s: %s\n", info.Name(), err)
			} else {
				fmt.Printf("Moved: %s -> %s\n", path, newPath)
			}
		}
		return nil
		})
		

		}



