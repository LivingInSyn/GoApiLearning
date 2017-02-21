#go run .\main.go .\LogEvent.go .\Routes.go .\Handlers.go
$command = "go run "
foreach($file in Get-ChildItem C:\users\jmill\Documents\GoApi -Filter *.go) {
    $command = $command + $file.Name + " "
}
iex $command