I started this to make a handwritten Redis clone in Go and see where it takes me. 

Good sample commands to run in PowerShell (assuming that the server here is running): 
SET: $c = New-Object System.Net.Sockets.TcpClient("localhost",6380); $s=$c.GetStream(); $w=New-Object IO.StreamWriter($s); $w.AutoFlush=$true; $r=New-Object IO.StreamReader($s); $w.Write("*3`r`n`$3`r`nSET`r`n`$4`r`nname`r`n`$6`r`nxander`r`n"); $r.ReadLine()
GET: $c = New-Object System.Net.Sockets.TcpClient("localhost",6380); $s=$c.GetStream(); $w=New-Object IO.StreamWriter($s); $w.AutoFlush=$true; $r=New-Object IO.StreamReader($s); $w.Write("*2`r`n`$3`r`nGET`r`n`$6`r`nxander`r`n"); $r.ReadLine()

Notes on bytes format (how * and $ are used):
*3 = a command is coming with 3 words
$3 = next word is 3 letters long

Should be noted that, in true Redis, every value is wrapped in a structure called a Redis Object (robj). It tracks metadata about the value, including its type and its encoding. I could implement that eventually, but I don't do that quite yet.