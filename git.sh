#!/bin/bash
read -p "Commit description: " desc
git add "CSharp/FileList/Program.cs" 
git add "GoLang/main.go"
git add "PowerShell/filelist.ps1"
git add "/Python/filelist.py"
git add "git.sh"
git commit -m "$desc" 
git push -u origin master
read -p "Press any Key to terminate" aa