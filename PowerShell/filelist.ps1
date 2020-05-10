#==============================================================================
#
#   Created by      :   Wriju Ghosh
#   Created on      :   10-May-2020
#   Updated on      :   10-May-2020-
#   
#==============================================================================
$arr = @()
$rootFolder = ""
$a = Get-Date
$fileDate = ($a.Day.ToString() + $a.Month.ToString() + $a.Year.ToString() + $a.Hour.ToString() + $a.Minute.ToString() + $a.Second.ToString())
$outputFile = "c:\temp\FileList_" + $fileDate +".csv"
$fileExtension = ".pdf"
 
Get-ChildItem -Path $rootFolder -Filter $fileExtension -recurse | ? {$_.PSIsContainer -eq $False} | % {
    $obj = New-Object PSObject
    
    $obj | Add-Member NoteProperty Name $_.Name
    $obj | Add-Member NoteProperty Directory $_.DirectoryName.Replace($rootFolder + "\","")
    
    $arr += $obj
}
 
$arr | Export-CSV -Path $outputFile -noTypeInformation -Force
 
Invoke-Item $outputFile