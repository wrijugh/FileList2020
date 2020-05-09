# File List 2020 

## Description
Makes a list of files from your local folder in Windows OS. This works for all subfolders recurrsively. Gives serial number to individual files under a specific folder. Optionally you can pass two other parameters to the entry function 
* Search pattern to filter the file extenstion 
* If you don't want reccursive subfolder search

Since they are optional - the default values are already provided.  


Currently available in C#. Plan is to have version in Golang and Python as well. 

## How to run the C# code
In the folder you will find a file name *Program.cs* 

Assuming that you have the latest *.netcore* installed

Then run the below command 

```CLI
dotnet new console
````
This will create a new blank Console application. Replace the *Program.cs* with the *Program.cs* from here

Then run 

````CLI
dotnet run
````
This would then compile and run the application. A parameter will be asked. Please enter the full path. 