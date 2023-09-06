#!/env/bin/python3

import json
import time 
import requests as r

confile = open("config.json",'r')
config = json.load(confile)
confile.close()
keys={}
exptime = int(config["expire"])

def rmfile(site,key,fileid):
    """
    删除过期文件
    """
    url = "https://{}/api/drive/files/delete".format(site)
    file_json = {"i": key, "fileId": fileid}
    res=r.post(url, json=file_json)
    return res.text
    



for site in config["site"]:
    keys[site["name"]] = site["key"]

logfile = open("files.log",'r')
log = logfile.read()

logfile.close()

nowtime=int(time.time())
newlines=[]
logs=log.splitlines()
for line in logs:
    xline = line.split(" ")
    print(xline)
    if nowtime -exptime > int(xline[0]):
        print("\n+++++++++++++++++++++")
        print(f"文件:{xline[1]}({xline[0]})来自{xline[2]} 已过期,应当删除。")
        res=rmfile(xline[2],keys[xline[2]],xline[1])
        print(res)
        print("+++++++++++++++++++++\n")
    else :
        print("\n+++++++++++++++++++++")
        print(f"文件:{xline[1]}({xline[0]})来自{xline[2]} 未过期,不应当删除。")
        newlines.append(line)
        print("+++++++++++++++++++++\n")
        # newlines.append(line)
logfile = open("files.log",'w')
logfile.writelines([f"{line}\n" for line in newlines])
logfile.close()





