import sys, json
array = json.load(sys.stdin)["items"][0]["metrics"]
for item in array:
    print item 


#print array["load"]

