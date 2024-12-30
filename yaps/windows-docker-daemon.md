# copying and go-testing files from a Linux runtime into a Windows container

## -it mode is disabled (we will have to keep the container running manually) 

running 
```
docker run -it --rm -v "$(pwd):C:/derived-output-event-system" -w C:/app mcr.microsoft.com/windows/servercore:ltsc2022 cmd
```
will give you:
```
the input device is not a TTY
Error: Process completed with exit code 1.
```

will need to run in the foreground or detached (see reason (a))

Trying foreground:
```
Run docker run --name go-for-windows mcr.microsoft.com/windows/servercore:ltsc2019 powershell -NoExit -Command "echo Hello"
Hello
PS C:\> CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
Error response from daemon: Container ebfcde66e504b59e2acdf26df560351184ac0637ed28a741e5f1a8e8717d99e4 is not running
```
Therefore, we will need, for now, since we will need to also copy some files from our host machine via `docker cp`, 
detached mode and `echo Hello; Start-Sleep -Seconds 100000` to keep our container running,

for debugging, we can temporarily use
```
docker stop go-for-windows
docker rm go-for-windows
```
to resolve the duplicate container conflict.

this then results in:
```
container 5c6d7ad6b94d98479b83e8867c47e97d8bb3ae55a3f22f33ddb1bde819b1f342 encountered an error during hcsshim::
System::CreateProcess: failure in a Windows system call: The system cannot find the file specified
```

However, I did notice that this message appears in the interactive argo terminal! x2 in fact, consistently seemingly.

and fwiw `powershell` is the correct shell to use here, since this works:
```
kubectl exec -it does-backfill-filesync-64c6c5d4db-55wpk -- powershell
```

oh ! It's just the `exec` command I have under it that isn't working. The container is running!

### onto copying...

```
docker cp /runner/_work/derived-output-event-system/derived-output-event-system go-for-windows:C:/derived-output-event-system/
```

but will need to create the directory first or else `Error response from daemon: No relative path specified in "C:"`

```
docker exec go-for-windows powershell mkdir C:\derived-output-event-system
```

### onto exec...

```
docker exec -it go-for-windows powershell
```
will result in:
```
the input device is not a TTY
```

so the sleep is still needed to make sure containers don't exit.

Installing Go...
```
docker exec go-for-windows powershell -Command \
            "& { Invoke-WebRequest -Uri 'https://go.dev/dl/go1.22.10.windows-amd64.msi' -OutFile 'go.msi'; \
            Start-Process -FilePath 'msiexec.exe' -ArgumentList '/i', 'go.msi', '/quiet', '/norestart' -NoNewWindow -Wait; \
            Remove-Item -Path 'go.msi' }"
```

Running tests...
```
docker exec go-for-windows powershell cd C:\derived-output-event-system; go test -v ./...
```

I tried `docker logs`, and this, albeit printed my test `echo Hello;`, did not print my test output.  

Ah! You will need `-NoExit`, see reason (b)

## In conclusion, a working block of docker commands look like:

```
docker run -d --name go-for-windows mcr.microsoft.com/windows/servercore:ltsc2019 powershell -NoExit -Command "echo Hello; Start-Sleep -Seconds 3600"
docker exec go-for-windows powershell mkdir C:\derived-output-event-system
docker cp /runner/_work/derived-output-event-system/derived-output-event-system go-for-windows:C:/derived-output-event-system/
docker exec go-for-windows powershell -Command \
            "& { Invoke-WebRequest -Uri 'https://go.dev/dl/go1.22.10.windows-amd64.msi' -OutFile 'go.msi'; \
            Start-Process -FilePath 'msiexec.exe' -ArgumentList '/i', 'go.msi', '/quiet', '/norestart' -NoNewWindow -Wait; \
docker exec go-for-windows powershell -Command "{ cd C:\derived-output-event-system; & 'C:\Program Files\Go\bin\go.exe' test -v ./... }"

docker stop go-for-windows
docker rm go-for-windows
```

this didn't work btw:
```
docker start -ai go-for-windows powershell -Command "echo Hello; Start-Sleep -Seconds 3600"
```
resulted in
```
unknown shorthand flag: 'C' in -Command
```

I decided against using `start` because it gave me problems with running background commands...


*appendix*

## reasons
(a)
> https://stackoverflow.com/questions/62786853/docker-run-command-without-it-option
When you run docker run without -it it's still running the container but you've not given it a command, so it finishes and exits.
-i is saying keep it alive and work within in the terminal (allow it to be interactive), but if you type exit, you're done and the container stops.
-t is showing the terminal of within the docker container (see: What are pseudo terminals (pty/tty)?)
-it allows you to see the terminal in the docker instance and interact with it.
Additionally you can use -d to run it in the background and then get to it afterwards.

## reasons
(b)
> 

## to verify the daemon *can indeed* run *a* container...

```
docker run hello-world
```
works:

```
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
fc1cdf365373: Pulling fs layer
8192e4d745e6: Pulling fs layer
...
fc1cdf365373: Verifying Checksum
fc1cdf365373: Download complete
Hello from Docker!
This message shows that your installation appears to be working correctly.
```

