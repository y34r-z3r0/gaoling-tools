# golang-tools

Hint for building executable files: [link](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

### log-generator
A simple test log generator that prints the current date and time. You can change seconds to minutes by changing `time.Second` to `time.Minute` and the dependent variable name `randomSecond`.

Launch hint:
```
sudo nohup ./log-generator &
```

Example:
```
2024.02.03 18:15:45
2024.02.03 18:15:46
2024.02.03 18:15:49
```

### hider
Especially for those who have something to hide.
Put builded binary file in the directory where you want to hide all the files and run it. Once executed, the script file will self-destruct. 

UNIX only!