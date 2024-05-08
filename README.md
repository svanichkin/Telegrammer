# Telegrammer 

The bot features a streamlined interface designed for ease of operation. When sending a request to the bot, users have the option to include a parameter called type, which specifies the desired response format as either `ini` or `json`. The default response format is set to `json`.

It is important to note that the `bid` and `uid` parameter is mandatory for all bot methods, ensuring proper functionality. The bot is capable of processing both `GET` and `POST` requests.

In the event of an error, the bot will return `{"error":"error description"}`, along with the corresponding `HTTP status code`. For successful requests, the bot will return status `200`.

The service can return an error in any of the following formats: `json`, `jsonp`, `securejson`, `indentedjson`, `asciijson`, `purejson`, `xml`, `html`, `yaml`, `toml`, `protobuf`, `ini`. To do this, you need to specify the `type` parameter with any of these types.

The `Python` script can request text input from the user using `sys.stdin.read()`. It is also possible to pass parameters immediately following the command, for example, `/my_command parameter`.

The bot provides a mechanism for attaching additional identifiers called aliases `aid`. Multiple aliases can be added for each user. This may be necessary, for example, if your third-party service has its own user identifiers. In that case, you can easily link them and refer to them by the identifiers from the third-party service.

You can differentiate access rights and create your own user groups, which allows you to flexibly configure access for specific users and/or chats.

## Features

- Single file startup
- Low resource usage
- Command rules
- Simplicity and ease of use
- Response format option
- Support for GET and POST methods
- Aliases for identifiers
- User access configuration
- Logic in Python files

### API → Send

`/send?bid=sample_bot_id&uid=123&cid=234&text=Hello&type=ini`  
The method sends text to the user. Use `uid` (user id) or `aid` (alias id) if needed, `cid` (chat id) is optional.

`/send?bid=sample_bot_id&uid=123&cid=234&html=<b>Hello</b>&type=json`  
The method sends text to the user in `HTML` format. Use `uid` (user id) or `aid` (alias id) if needed, `cid` (chat id) is optional. Only simple tags are supported (see Telegram description).

`/send?bid=sample_bot_id&uid=123&cid=234&markdown=**Hello**&type=json`  
The method sends text to the user in `MARKDOWN` format. Use `uid` (user id) or `aid` (alias id) if needed, `cid` (chat id) is optional. Only simple tags are supported (see Telegram description).

`/send?bid=sample_bot_id&uid=123&cid=234&markdownv2=**Hello**&type=json`  
The method sends text to the user in `MARKDOWNV2` format. Use `uid` (user id) or `aid` (alias id) if needed, `cid` (chat id) is optional. Only simple tags are supported (see Telegram description).

`/send?bid=sample_bot_id&uid=123&cid=234&photo=Base64&type=ini`  
The method parses `Base64` and sends it as an image to the user. Use `uid` (user id) or `aid` (alias id) if needed, `cid` (chat id) is optional.

`/send?bid=sample_bot_id&uid=123&cid=234&photo=./sample/filename.png&type=json`  
The method sends the file specified in the local `path` to the user. Use `uid` (user id) or `aid` (alias id) if needed, `cid` (chat id) is optional.

`/send?bid=sample_bot_id&uid=123&cid=234&photo=http://myserver.com/filename.png&type=ini`  
The method sends the file specified in the `URL` to the user. Use `uid` (user id) or `aid` (alias id) if needed, `cid` (chat id) is optional.

### API → Check

`/check?bid=sample_bot_id&uid=123&cid=234&type=json`  
The method checks if the user is in the specified chat. To perform the check, the bot must be in the specified chat. Use `uid` (user id) or `aid` (alias id) if needed.

### API → Data

`/data?do=set&bid=sample_bot_id&uid=123&cid=234&key=mykey&value=myvalue&type=json`  
The method saves the key and value in the storage, and this data will be located in the config folder. Use `uid` or `aid` if needed.

`/data?do=get&bid=sample_bot_id&uid=123&cid=234&key=mykey&type=json`  
The method removed the saved value from the storage. Use `uid` or `aid` if needed.

`/data?do=remove&bid=sample_bot_id&uid=123&cid=234&key=mykey&type=json`  
The method retrieves the saved value from the storage. Use `uid` or `aid` if needed.

`/data?do=set&bid=sample_bot_id&uid=123&cid=234&keys=key1,key2&values=value1,value2&type=ini`  
The method saves keys and values in the storage, and this data will be located in the config folder. Use `uid` or `aid` if needed.

`/data?do=set&bid=sample_bot_id&uid=123&cid=234&keys=key1,key2&type=ini`  
The method removed the saved values from the storage. Use `uid` or `aid` if needed.

`/data?do=remove&bid=sample_bot_id&uid=123&cid=234&keys=key1,key2&type=ini`  
The method retrieves the saved values from the storage. Use `uid` or `aid` if needed.

### API → Command

`/command?bid=sample_bot_id&name=command_name&data=param_for_command&uid=123&cid=234&firstname=John&lastname=Jhonson&username=jj&type=ini`  
The method executes the bot command. Optional paramteres: `uid`, `cid`, `data`, `firstname`, `lastname`, `username`. Use `uid` or `aid` if needed.

### API → Alias

`/alias?bid=sample_bot_id&uid=123&aid=new_alias&do=add&type=ini`  
The method adds an alias for accessing the service, similar to how it's done with `uid`. Multiple aliases can be added.

`/alias?bid=sample_bot_id&uid=123&aid=new_alias&do=remove&type=ini`  
The method removes an alias for the specified `uid`.

`/alias?bid=sample_bot_id&aid=new_alias&do=find&type=ini`  
The method searches for a `uid` by alias.

### API → Restart

`/restart?bid=sample_bot_id`  
Restarts the bot for the specified config. Can be useful if the config file has been manually edited.

### API → Access

`/access?bid=sample_bot_id&uid=123&group=user&do=set&type=ini`
Adds a user to the list of `users` in the specified group. Use `uid` or `aid` if needed.

`/access?bid=sample_bot_id&cid=123&group=user&do=set&type=ini`
Adds a chat to the list of `chats` in the specified group.

`/access?bid=sample_bot_id&uid=123do=get&type=ini`
Retrieves the name of the group if the user is found in `users`. Use `uid` or `aid` if needed.

`/access?bid=sample_bot_id&cid=123&do=get&type=ini`
Retrieves the name of the group if the chat is found in `chats`.

`/access?bid=sample_bot_id&uid=123&group=user&do=remove&type=ini`
Removes a user from the list of `users` in the specified group. Use `uid` or `aid` if needed.

`/access?bid=sample_bot_id&cid=123&group=user&do=remove&type=ini`
Removes a chat from the list of `chats` in the specified group.

## Config

The main configuration file is located at `./config.ini` and it contains the port and host information for launching the `API`. Additional configuration files for specific bots can be found in the `./bots/*/config.ini` folder.

The config lists commands in the commands parameter, then each command is described in sections. Multiple parameters can be used in each section.

The scripting language can be any, it is important that the script interpreter is located in the system, as the service executes the script in the system's command line.

By default, bots are allowed to interact with everyone, but there's the option to impose restrictions and create user groups. You can restrict commands for specific `uid` by adding them to `users`. You can also limit the bot's interaction with specific chats by creating a list of `chats`. These functions are also available through the `API`.

When the command is invoked, the environment variable `ENV` will contain a list of `groups` for that `command`, such as `ACCESS_GROUPS=admin, user`, and `ACCESS_ACTION=true/false`, indicating to the script that executing the `command` is not necessary because the user does not belong to any of the `groups`. After processing by the script, you can write a message to the user informing them that they need to join the specified `groups`.

### Main config.ini

```ini
host = localhost
port = 8010
bots = ./bots
```

### Main parameters

`host = ipaddress`  
Specifies the `IP address` on which the `API` will operate to handle `HTTP` requests. Addresses can be specified as `ip`, for example, `127.0.0.1`.

`порт = number`  
Specifies the `port` on which the `API` will operate to handle `HTTP` requests. Ensure that the system allows the use of this `port` and that it is not occupied by another application.

`bots = ./bots`  
Specifies the `path` where the folder containing the bots will be located. Create the folder and place it at the specified `path`. Within the bots folder, create a folder for your new bot and name it, for example, `sample`. Place the `config.ini` file and necessary scripts for the bot's commands within it. The folder structure will be `./bots/sample/config.ini`.

### Sample config.ini for telegram bot

```ini
bid = sample_bot_id
key = 2467348235:ABGvJ45chUygOPzjdpQRGFXH2ZGb_APc2QU

commands = start, command1, command2
groups = admin, user

# Commands

[start]
showed = false
action = ./bots/sample/action.py
access = admin, user

[command1]
detail = ./bots/sample/detail.sh
action = ./bots/sample/action.py
access = admin, user

[command2]
detail = ./bots/sample/detail.sh
action = ./bots/sample/action.py
access = admin


# Access

[admin]
users = 123

[user]
users = 234, 423424, -3432232, 7676434
chats = 5545454

```
### Main parameters

`bid = sample_bot_id`  
This is the internal `identification` name for your bot. This parameter is used for passing in the `API` and other places. Since there may be many bots, it's necessary to distinguish commands coming to a specific bot from others.

`key = 2467348235:ABGvJ45chUygOPzjdpQRGFXH2ZGb_APc2QU`  
The key to access the bot's `API`. You can create it in Telegram. Refer to the documentation on creating bots in Telegram for more information.

`commands = start, command1, command2`  
Listing all the `commands` that the bot can use, both through the `API` and via the menu in Telegram.

`groups = admin, user`  
Creates user `groups`. User `groups` allow you to display specific menu `commands` for `groups`. Restrictions can occur based on both the `user id` and the `chat id`. This allows for flexible access rights configuration, including when using the `API`. This is an optional parameter.

### Command parameters

`showed = true/false` default is `menu = true`  
This parameter determines whether the command will be visible in the bot's command menu or will be hidden. For example, the `/start` command, which is the default command of any bot, definitely does not need to be in the menu. Commands can also be called through `API`, which allows for more flexible logic building. This is an optional parameter.

`detail = ./bots/detail.py`  
This is the description of the command that will be shown under the command in the bot’s menu. The text will be returned from the script. `ENV` variables such as `LANG, COMMAND, UID, BID, AID, FIRST_NAME, LAST_NAME, USER_NAME, DO_TYPE, DATA`, and others are set for the script, so that one script can handle any logic for any command.

`action = ./bots/action.py`  
When a command is called, a script will be executed which in turn will perform the necessary actions for the bot. `ENV` variables such as `LANG, COMMAND, UID, BID, AID, FIRST_NAME, LAST_NAME, USER_NAME, DO_TYPE, DATA`, and others are set for the script, so that one script can handle any logic for any command.

`access = admin, user`  
This parameter allows viewing and executing the `command` for the specified user `groups`. This is an optional parameter.

### Group parameters

`users = 234, 423424, -3432232, 7676434`  
Adds `users` to the list associated with a group. This is an optional parameter if `chats` are specified.

`chats = 5545454`  
Adds `chats` to the list associated with a group. This is an optional parameter if `users` are specified.