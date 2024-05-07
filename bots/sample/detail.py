#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys

def description():
	
	language = os.getenv('LANG', 'en')	
	command  = os.getenv('COMMAND')
	
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
			'en': 'Send message to chat',
			'fr': 'Envoyer un message au chat',
			'de': 'Nachricht an den Chat senden',
			'es': 'Enviar mensaje al chat',
			'ru': 'Отправить сообщение в чат',
			'it': 'Invia messaggio alla chat',
			'pt': 'Enviar mensagem para o chat',
			'zh': '发送消息到聊天',
			'ja': 'チャットにメッセージを送信',
			'ar': 'إرسال رسالة إلى الدردشة'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'send_message_html':
		localizations = {
			'en': 'Send message to chat',
			'fr': 'Envoyer un message au chat',
			'de': 'Nachricht an den Chat senden',
			'es': 'Enviar mensaje al chat',
			'ru': 'Отправить сообщение в чат',
			'it': 'Invia messaggio alla chat',
			'pt': 'Enviar mensagem para o chat',
			'zh': '发送消息到聊天',
			'ja': 'チャットにメッセージを送信',
			'ar': 'إرسال رسالة إلى الدردشة'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'send_picture_url':
		localizations = {
			'en': 'Send picture from URL to chat',
			'fr': 'Envoyer une image depuis une URL au chat',
			'de': 'Bild über URL an den Chat senden',
			'es': 'Enviar imagen desde URL al chat',
			'ru': 'Отправить изображение из URL в чат',
			'it': 'Invia immagine da URL alla chat',
			'pt': 'Enviar imagem de URL para o chat',
			'zh': '从URL发送图片到聊天',
			'ja': 'URLからチャットに画像を送信',
			'ar': 'إرسال صورة من URL إلى الدردشة'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'check_subscribe':
		localizations = {
			'en': 'Check users subscribed to channel',
			'fr': 'Vérifier les utilisateurs abonnés au canal',
			'de': 'Überprüfen Sie die Abonnenten des Kanals',
			'es': 'Comprobar los usuarios suscritos al canal',
			'ru': 'Проверить подписчиков канала',
			'it': 'Controllare gli utenti iscritti al canale',
			'pt': 'Verificar usuários inscritos no canal',
			'zh': '检查订阅频道的用户',
			'ja': 'チャンネルに登録されているユーザーを確認',
			'ar': 'تحقق من المستخدمين المشتركين في القناة'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'set_value':
		localizations = {
			'en': 'Set a value in the storage, enter keyName = value',
			'fr': 'Définir une valeur dans le stockage, entrez keyName = value',
			'de': 'Setzen Sie einen Wert im Speicher, eingeben keyName = value',
			'es': 'Establece un valor en el almacenamiento, introduzca keyName = value',
			'ru': 'Устанавливает значение в хранилище, введите keyName = value',
			'it': 'Imposta un valore nello storage, inserire keyName = value',
			'pt': 'Definir um valor no armazenamento, digite keyName = value',
			'zh': '在存储中设置值，请输入 keyName = value',
			'ja': 'ストレージに値を設定する、keyName = value を入力してください',
			'ar': 'تعيين قيمة في التخزين، أدخل keyName = value'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'get_value':
		localizations = {
			'en': 'Reads a value from the storage, enter keyName',
			'fr': 'Lit une valeur du stockage, entrez keyName',
			'de': 'Liest einen Wert aus dem Speicher, geben Sie keyName ein',
			'es': 'Lee un valor del almacenamiento, introduzca keyName',
			'ru': 'Читает значение из хранилища, введите keyName',
			'it': 'Legge un valore dall’archiviazione, inserire keyName',
			'pt': 'Lê um valor do armazenamento, digite keyName',
			'zh': '从存储中读取值，请输入 keyName',
			'ja': 'ストレージから値を読み取る、keyName を入力してください',
			'ar': 'يقرأ قيمة من التخزين، أدخل keyName'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'remove_value':
		localizations = {
			'en': 'Removes a value from the storage, enter keyName',
			'fr': 'Supprime une valeur du stockage, entrez keyName',
			'de': 'Entfernt einen Wert aus dem Speicher, geben Sie keyName ein',
			'es': 'Elimina un valor del almacenamiento, introduzca keyName',
			'ru': 'Удаляет значение из хранилища, введите keyName',
			'it': 'Rimuove un valore dall’archiviazione, inserire keyName',
			'pt': 'Remove um valor do armazenamento, digite keyName',
			'zh': '从存储中删除值，请输入 keyName',
			'ja': 'ストレージから値を削除します、keyName を入力してください',
			'ar': 'يزيل قيمة من التخزين، أدخل keyName'
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'do_command':
		localizations = {
			"en": "Executes any bot command",
			"fr": "Exécute n'importe quelle commande du bot",
			"de": "Führt beliebigen Bot-Befehl aus",
			"es": "Ejecuta cualquier comando del bot",
			"ru": "Выполняет любую команду бота",
			"it": "Esegue qualsiasi comando del bot",
			"pt": "Executa qualquer comando do bot",
			"zh": "执行任何机器人命令",
			"ja": "任意のボットコマンドを実行します",
			"ar": "ينفذ أي أمر للروبوت",
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'im_admin':
		localizations = {
			"en": "Checks for admin access",
			"fr": "Vérifie l'accès administrateur",
			"de": "Überprüft den Administratorzugang",
			"es": "Verifica el acceso de administrador",
			"ru": "Проверяет есть ли админский доступ",
			"it": "Controlla l'accesso amministrativo",
			"pt": "Verifica o acesso de administrador",
			"zh": "检查管理员访问权限",
			"ja": "管理者アクセスを確認します",
			"ar": "يتحقق من وجود حق الوصول الإداري",
		}
		return localizations.get(language, localizations.get('en'))

	if command == 'im_superuser':
		localizations = {
			"en": "Checks if superuser access is available",
			"fr": "Vérifie si l'accès superutilisateur est disponible",
			"de": "Überprüft, ob der Superuser-Zugriff verfügbar ist",
			"es": "Comprueba si hay acceso de superusuario disponible",
			"ru": "Проверяет, доступен ли суперпользователь",
			"it": "Controlla se è disponibile l'accesso come superutente",
			"pt": "Verifica se o acesso de superusuário está disponível",
			"zh": "检查是否有超级用户访问权限",
			"ja": "スーパーユーザーアクセスが利用可能かどうかを確認します",
			"ar": "تحقق مما إذا كان الوصول كمستخدم جذر متاحًا",
		}
		return localizations.get(language, localizations.get('en'))

	return ''

sys.stdout.write(description())
sys.stdout.flush()