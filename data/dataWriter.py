# Writes the clean JSON-data to a MySQL Database

import pymysql
import json
import datetime

# AUTH

with open('credentials/db_pw.txt', 'r') as f:
    pw = f.read()

# CONNECTION

mydb = pymysql.connect(host='localhost',
                       user='root',
                       passwd=pw,
                       db='reddit_go')

cursor = mydb.cursor()


# Insert users

def pushUsers():
    with open('json/cleanup/clean_users.json', 'r') as infile:
        file = json.load(infile)

    for object in file:
        cursor.execute("INSERT INTO user (Id, Username, Bio, Email) VALUES (%s,%s,%s, %s)",
                       (object["id"], object["username"], object["bio"], object["email"]))


# Insert subreddits

def pushCommunities():
    with open('json/cleanup/clean_subreddits.json', 'r') as infile:
        file = json.load(infile)

    for object in file:
        cursor.execute("INSERT INTO community (Id, Name, Description, Description_html) VALUES (%s,%s,%s, %s)",
                       (object["id"], object["name"], object["description"], object["description_html"]))


# Insert posts

def pushPosts():
    with open('json/cleanup/clean_posts.json', 'r') as infile:
        file = json.load(infile)

    for object in file:
        cursor.execute("INSERT INTO post (Id, Title, Content, Content_html, Timestamp, User_Id, Community_Id) VALUES (%s,%s,%s,%s,%s,%s,%s)", (
            object["id"], object["title"], object["content"], object["content_html"], datetime.datetime.fromtimestamp(object["timestamp"]), object["user_id"], object["community_id"]))


# Insert comments

def pushComments():
    with open('json/cleanup/clean_comments.json', 'r') as infile:
        file = json.load(infile)

    for object in file:
        cursor.execute("INSERT INTO comment (Id, Content, Timestamp, Post_Id, User_Id) VALUES (%s,%s,%s, %s, %s)", (
            object["id"], object["content"], datetime.datetime.fromtimestamp(object["timestamp"]), object["post_id"], object["user_id"]))

def pushUserCommunities():
    with open('json/cleanup/clean_usercommunities.json', 'r') as infile:
        file = json.load(infile)

    for object in file:
        cursor.execute("INSERT INTO usercommunity (Community_Id, User_Id) VALUES (%s,%s)", (
            object["community_id"], object["user_id"]))

# EXECUTE REGION

# pushUsers()

# pushCommunities()

# pushPosts()

# pushComments()

pushUserCommunities()


# close the connection to the database.
mydb.commit()
cursor.close()
print("Done")
