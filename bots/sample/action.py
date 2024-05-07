#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys

try: # Python 3
	from urllib.parse import quote
	from urllib.request import urlopen
	from urllib.error import URLError, HTTPError
except: # Python 2
	from urllib import quote
	from urllib2 import urlopen
	from urllib2 import URLError, HTTPError

def action():
	
	language  = os.getenv('LANG', 'en')	
	command   = os.getenv('COMMAND')
	bid       = os.getenv('BID')
	uid       = os.getenv('UID')
	cid       = os.getenv('CID')
	firstname = os.getenv('FIRST_NAME')
	lastname  = os.getenv('LAST_NAME')
	username  = os.getenv('USER_NAME')
	data      = os.getenv('DATA')
	server    = "http://localhost:8010"
	groups    = os.getenv('ACCESS_GROUPS')
	access    = os.getenv('ACCESS')

	if access == 'false':
		localizations = {
			'en': 'Access denied',
			'fr': 'Accès refusé',
			'de': 'Zugriff verweigert',
			'es': 'Acceso denegado',
			'ru': 'Доступ запрещен',
			'it': 'Accesso negato',
			'pt': 'Acesso negado',
			'zh': '访问被拒绝',
			'ja': 'アクセスが拒否されました',
			'ar': 'الوصول مرفوض'
		}
		return localizations.get(language, localizations.get('en'))
	
	if command == 'start':
		return ''

	if command == 'return_string':
		localizations = {
			'en': 'Return string from script',
			'fr': 'Chaîne de retour du script',
			'de': 'Zeichenfolge aus Skript zurückgeben',
			'es': 'Cadena de retorno del script',
			'ru': 'Строка, возвращаемая из скрипта',
			'it': 'Stringa di ritorno dello script',
			'pt': 'String de retorno do script',
			'zh': '脚本返回字符串',
			'ja': 'スクリプトからの戻り文字列',
			'ar': 'السلسلة المعادة من السكريبت'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'send_message':
		localizations = {
			'en': 'hello!',
			'fr': 'salut!',
			'de': 'hallo!',
			'es': '¡hola!',
			'ru': 'привет!',
			'it': 'ciao!',
			'pt': 'olá!',
			'zh': '你好！',
			'ja': 'こんにちは！',
			'ar': 'مرحباً!'
		}
		returnedString = quote(firstname + " " + lastname + " " + username + " " + localizations.get(language, localizations.get('en')))
		urlopen(server + '/send?bid=' + bid + '&uid=' + uid + '&cid=' + cid + '&text=' + returnedString + '&type=ini')

	if command == 'send_message_html':
		localizations = {
			'en': 'hello!',
			'fr': 'salut!',
			'de': 'hallo!',
			'es': '¡hola!',
			'ru': 'привет!',
			'it': 'ciao!',
			'pt': 'olá!',
			'zh': '你好！',
			'ja': 'こんにちは！',
			'ar': 'مرحباً!'
		}
		returnedString = quote("<b>" + firstname + "</b> " + localizations.get(language, localizations.get('en')))
		urlopen(server + '/send?bid=' + bid + '&uid=' + uid + '&cid=' + cid  + '&html=' + returnedString + '&type=ini')
		
	if command == 'send_picture_url':
		returnedString = quote("https://upload.wikimedia.org/wikipedia/en/thumb/3/3b/SpongeBob_SquarePants_character.svg/640px-SpongeBob_SquarePants_character.svg.png")
		urlopen(server + '/send?bid=' + bid + '&uid=' + uid + '&cid=' + cid  + '&photo=' + returnedString + '&type=ini')

	if command == 'check_subscribe':
		localizationsFalse = {
			'en': 'You are not subscribed to the channel!',
			'fr': 'Vous n\'êtes pas abonné à la chaîne!',
			'de': 'Du bist nicht auf den Kanal abonniert!',
			'es': '¡No estás suscrito al canal!',
			'ru': 'ты не подписан на канал!',
			'it': 'Non sei iscritto al canale!',
			'pt': 'Você não está inscrito no canal!',
			'zh': '你没有订阅该频道！',
			'ja': 'チャンネルに登録していません！',
			'ar': 'أنت غير مشترك في القناة!'
		}
		localizationsTrue = {
			'en': 'You are subscribed to the channel!',
			'fr': 'Vous êtes abonné à la chaîne!',
			'de': 'Du bist auf den Kanal abonniert!',
			'es': '¡Estás suscrito al canal!',
			'ru': 'подписан на канал!',
			'it': 'Sei iscritto al canale!',
			'pt': 'Você está inscrito no canal!',
			'zh': '你已订阅该频道！',
			'ja': 'チャンネルに登録しています！',
			'ar': 'أنت مشترك في القناة!'
		}
		returnedStringFalse = firstname + " " + localizationsFalse.get(language, localizationsFalse.get('en'))
		returnedStringTrue = firstname + " " + localizationsTrue.get(language, localizationsTrue.get('en'))
		url = server + '/check?bid=' + bid + '&uid=' + uid + '&cid=' + cid + '&type=ini'
		try:
			urlopen(url)
			return returnedStringTrue
		except HTTPError as e:
			return returnedStringFalse
			# return e.read().decode('utf-8')
		except URLError as e:
			return str(e)
		except Exception as e:
			return str(e)

	if command == 'set_value':
		parts = data.split('=', 1)
		if len(parts) != 2:
			localizations = {
				'en': 'Enter keyName = value',
				'fr': 'Entrez keyName = value',
				'de': 'Geben Sie keyName = value ein',
				'es': 'Introduzca keyName = value',
				'ru': 'Введите keyName = value',
				'it': 'Inserire keyName = value',
				'pt': 'Digite keyName = value',
				'zh': '输入 keyName = value',
				'ja': 'keyName = value を入力してください',
				'ar': 'أدخل keyName = value'
			}
			returnedString = quote(localizations.get(language, localizations.get('en')))
			urlopen(server + '/send?bid=' + bid + '&uid=' + uid + '&cid=' + cid  + '&text=' + returnedString + '&type=ini')
			data = sys.stdin.read()
			parts = data.split('=', 1)
			if len(parts) == 2:
				key = parts[0].strip()
				value = parts[1].strip()
				url = server + '/data?do=set&bid=' + bid + '&uid=' + uid + '&key=' + key + '&value=' + value + '&type=ini'
				try:
					urlopen(url)
					return key + ' = ' + value
				except HTTPError as e:
					return e.read().decode('utf-8')
				except URLError as e:
					return str(e)
				except Exception as e:
					return str(e)
		else:
			key = parts[0].strip()
			value = parts[1].strip()
			url = server + '/data?do=set&bid=' + bid + '&uid=' + uid + '&key=' + key + '&value=' + value + '&type=ini'
			try:
				urlopen(url)
				return key + ' = ' + value
			except HTTPError as e:
				return e.read().decode('utf-8')
			except URLError as e:
				return str(e)
			except Exception as e:
				return str(e)

	if command == 'get_value':
		if len(data) == 0:
			localizations = {
				'en': 'Enter keyName',
				'fr': 'Entrez keyName',
				'de': 'Geben Sie keyName ein',
				'es': 'Introduzca keyName',
				'ru': 'Введите keyName',
				'it': 'Inserire keyName',
				'pt': 'Digite keyName',
				'zh': '输入 keyName',
				'ja': 'keyName を入力してください',
				'ar': 'أدخل keyName'
			}
			returnedString = quote(localizations.get(language, localizations.get('en')))
			urlopen(server + '/send?bid=' + bid + '&uid=' + uid + '&cid=' + cid + '&text=' + returnedString + '&type=ini')
			data = sys.stdin.read()
			if len(data) > 0:
				url = server + '/data?do=get&bid=' + bid + '&uid=' + uid + '&key=' + data + '&type=ini'
				try:
					return urlopen(url).read()
				except HTTPError as e:
					return e.read().decode('utf-8')
				except URLError as e:
					return str(e)
				except Exception as e:
					return str(e)
		else:
			url = server + '/data?do=get&bid=' + bid + '&uid=' + uid + '&key=' + data + '&type=ini'
			try:
				return urlopen(url).read()
			except HTTPError as e:
				return e.read().decode('utf-8')
			except URLError as e:
				return str(e)
			except Exception as e:
				return str(e)

	if command == 'do_command':
		if len(data) == 0:
			localizations = {
				'en': 'Enter command',
				'fr': 'Entrez command',
				'de': 'Geben Sie command ein',
				'es': 'Introduzca command',
				'ru': 'Введите command',
				'it': 'Inserire command',
				'pt': 'Digite command',
				'zh': '输入 command',
				'ja': 'command を入力してください',
				'ar': 'أدخل command'
			}
			returnedString = quote(localizations.get(language, localizations.get('en')))
			urlopen(server + '/send?bid=' + bid + '&uid=' + uid + '&cid=' + cid  + '&text=' + returnedString + '&type=ini')
			data = sys.stdin.read()
			if len(data) != 0:
				url = server + '/command?bid=' + bid + '&uid=' + uid + '&name=' + data + '&type=ini'
				try:
					return urlopen(url).read()
				except HTTPError as e:
					return e.read().decode('utf-8')
				except URLError as e:
					return str(e)
				except Exception as e:
					return str(e)
		else:
			url = server + '/command?bid=' + bid + '&uid=' + uid + '&name=' + data + '&type=ini'
			try:
				return urlopen(url).read()
			except HTTPError as e:
				return e.read().decode('utf-8')
			except URLError as e:
				return str(e)
			except Exception as e:
				return str(e)

	if command == 'im_admin':
		return 'admin'

	if command == 'im_superuser':
		return 'superuser'

	return ''

sys.stdout.write(action())
sys.stdout.flush()