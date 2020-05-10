#!/bin/bash
read -p "Commit description: " desc
git add Program.cs 
git add main.go
git add filelist.ps1
git add filelist.py
git add git.sh
git commit -m "$desc" 
git push -u origin master
read -p "Press any Key" aa