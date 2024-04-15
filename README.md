# Telegrammer 

The bot features a streamlined interface designed for ease of operation. When sending a request to the bot, users have the option to include a parameter called type, which specifies the desired response format as either `ini` or `json`. The default response format is set to `json`.

It is important to note that the `bid` and `uid` parameter is mandatory for all bot methods, ensuring proper functionality. The bot is capable of processing both `GET` and `POST` requests.

In the event of an error, the bot will return `{"error":"error description"}`, along with the corresponding `HTTP status code`. For successful requests, the bot will return `{"key_name":"value for key"}` as part of the response.

### API → Send

`/send?bid=my_sample_bot&uid=123&text=Hello&type=ini`  
The method sends text to the user.

`/send?bid=my_sample_bot&uid=123&html=<b>Hello</b>&type=json`  
The method sends text to the user in HTML format. Only simple tags are supported (see Telegram description).

`/send?bid=my_sample_bot&uid=123&photo=Base64&type=ini`  
The method parses Base64 and sends it as an image to the user.

`/send?bid=my_sample_bot&uid=123&photo=./sample/filename.png&type=json`  
The method sends the file specified in the local path to the user.

`/send?bid=my_sample_bot&uid=123&photo=http://myserver.com/filename.png&type=ini`  
The method sends the file specified in the URL to the user.

### API → Check

`/check?bid=my_sample_bot&uid=123&chat=12345&type=json`  
The method checks if the user is in the specified chat. To perform the check, the bot must be in the specified chat.

### API → Data

`/set?bid=my_sample_bot&uid=123&key=mykey&value=myvalue&type=json`  
The method saves the key and value in the storage, and this data will be located in the same file of the user created at the start.

`/get?bid=my_sample_bot&uid=123&key=mykey&type=json`  
The method retrieves the saved value from the storage.

`/set?bid=my_sample_bot&uid=123&keys=key1,key2&values=value1,value2&type=ini`  
The method saves keys and values in the storage, and this data will be located in the same file of the user created at the start.

`/get?bid=my_sample_bot&uid=123&keys=key1,key2&type=ini`  
The method retrieves the saved values from the storage.

### API → Command

`/command?bid=my_sample_bot&uid=123&do=command_name`  
The method executes the bot command.

### API → Config

`/config?bid=id_of_new_service&set=url_encoded_text`  
The method saves or overwrites the configuration in the config folder, using the text string provided. The text will be read, validated, and then placed in the config folder.

`/config?bid=id_of_new_service&set=./path/to/config/new_config.txt`  
The method saves or overwrites the configuration in the config folder, using the specified local path. The file will be read, validated, and then placed in the config folder.

`/config?bid=id_of_new_service&set=http://server.com/new_config.ini`  
The method saves or overwrites the configuration in the config folder, using the specified URL. The file will be read, validated, and then placed in the config folder. After the new configuration is set, the bot automatically applies all changes.

`/config?bid=my_sample_bot`  
The method returns the requested configuration.

### API → Restart

`/restart?bid=id_of_new_service`
Restarts the bot for the specified config. Can be useful if the config file has been manually edited.

`/restart`
Restarts the API, can be useful if we have changed the main config.

## Config

There is a main config file located in the `./configs/config.ini` folder that only contains the port and host on which the API will be launched. Other configs will also be located in the same folder.
The config lists commands in the commands parameter, then each command is described in sections. Multiple parameters can be used in each section.

### Main config
```
# API server
host = localhost
port = 8010
```

### Sample config for telegram bot
```
# Bot id
bid = my_sample_bot
# Telegram bot id
key = 2467348235:ABGvJ45chUygOPzjdpQRGFXH2ZGb_APc2QU

# Commands
commands = start, command1, command2

[start]
# Not showed in menu
showed = false
# Do command
action = "./scripts/action.py"

[command1]
# Get description for command
description = "./scripts/description.py"
# Do command
action = "./scripts/action.py"

[command2]
# Get description for command
description = "./scripts/description.py"
# Do command
action = "./scripts/action.py"
```

### Command parameters
`showed = true/false` по умолчанию `showed = true`  
This parameter determines whether the command will be visible in the bot's command menu or will be hidden. For example, the `/start` command, which is the default command of any bot, definitely does not need to be in the menu. Commands can also be called through `API`, which allows for more flexible logic building.

`description = "./scripts/description.py"`  
This is the description of the command that will be shown under the command in the bot’s menu. The text will be returned from the script. `ENV` variables such as `language, command, uid, id`, and others are set for the script, so that one script can handle any logic for any command.

`action = "./scripts/action.py"`  
When a command is called, a script will be executed which in turn will perform the necessary actions for the bot. `ENV` variables such as `language, command, uid, id`, and others are set for the script, so that one script can handle any logic for any command.