'''
=================================================================================
This Python script would fetch all the files and folders from a Windows directory

Created by      :   Wriju Ghosh
Created on      :   10-May-2020
Updated on      :   10-May-2020
Python version  :   3.8.1

How to run      :   
Output          : The list is saved to a file named List.txt in the same executing folder 
=================================================================================
'''
import os
from os.path import join, getsize
import sys
# 

def GetAllSubFoldersAndFiles(folderPath):
    fileToWrite = open("List.txt", mode="wt", encoding="utf-8")
    for root, dirs, files in os.walk(folderPath):
        
        # print (root, "consumes", end=" ")
        # print(sum(getsize(join(root, name)) for name in files), end=" \r\n")
        if (len(files) != 0): # remove empty directories             
            fileToWrite.writelines("\n" + root + "\n")
            fileToWrite.writelines("===================\n")
            i = 1
            for file in files:
                # print the files
                fileToWrite.writelines("x" + ". " + file + "\n")
                i += 1
    fileToWrite.close()
    print("file list.txt is written")
if __name__ == '__main__':
    # can use sys.args[1] for the argument which takes path
    print(sys.argv[1])
    GetAllSubFoldersAndFiles(folderPath = sys.argv[1])