using System;
using System.IO;
using System.Text;

namespace FileListCS
{
    class Program
    {
        static void Main(string[] args)
        {
            //Console.WriteLine("Please enter the path:");
           
            var _folderPath = @"D:\OneDrive - Microsoft\Wriju_Personal\E_books_OneDriveBus\Bengali";
            GetAllSubDirectories(_folderPath);            
        }

        static void WriteToFile(string fileName, string textToWrite){
            //Write to the current path - the safe option. 
            //Else can also write to Documents folder 
            var fileWrite = File.CreateText(fileName);
            fileWrite.Write(textToWrite);
            fileWrite.Close();

            //Then Open File
            System.Console.WriteLine("Written to : " + fileName);
        }

        //int iCount = 1;
        static string GetAllSubDirectories(string folderPath, 
                string fileSearchPattern = "*.*", bool isRecurse = true)
        {
            string allData = string.Empty;

            EnumerationOptions recurse = new EnumerationOptions();
            recurse.RecurseSubdirectories = true;
            var directories = Directory.GetDirectories(folderPath, "*", recurse);
            //var directories = Directory.GetDirectories(folderPath);
            //Files in Root
            allData = allData + GetFiles(folderPath);
            //Files in Sub directories 
            for (int i=0; i < directories.Length; i++)
            {
                allData = allData + GetFiles(directories[i]);                
            }
            WriteToFile("temp.txt", allData);
            return allData;
        }

        static string GetFiles(string directoryPath, string fileSearchPattern = "*.*"){
            //Take the last part of the path as Folder Name to be displayed
            string folderName = directoryPath.Substring(directoryPath.LastIndexOf("\\")+1);
            string data = "\n\n" 
                            + folderName + "\n" 
                            + GenerateDash(folderName.Length) ;

            var files = Directory.GetFiles(directoryPath,"*.pdf");
            for(int i=0; i< files.Length; i++){
                data = data + "\n" + (i+1).ToString("D2") + ". " + Path.GetFileName(files[i]);                
            }

            Console.OutputEncoding = System.Text.Encoding.UTF8;
            //Console.WriteLine(data);
            return data;
        }

        
        static string GenerateDash(int len)
        {
            string retVal = "=";
            for(int i = 0; i<len; i++){
                retVal = retVal + "="; 
            }   

            return retVal;     
        }
    }


}
