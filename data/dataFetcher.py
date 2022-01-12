import requests
import json

BASE_URI = 'https://oauth.reddit.com'

# AUTH
with open('credentials/pw.txt', 'r') as f:
    pw = f.read()

with open('credentials/client_id.txt', 'r') as f:
    CLIENT_ID = f.read()

with open('credentials/secret_key.txt', 'r') as f:
    SECRET_KEY = f.read()

auth = requests.auth.HTTPBasicAuth(CLIENT_ID, SECRET_KEY)


data = {
    'grant_type': 'password',
    'username': 'Ibe_Verbeke',
    'password': pw,
}

headers = {'User-Agent': 'MyAPI/0.0.1'}

res = requests.post('https://www.reddit.com/api/v1/access_token',
                    auth=auth, data=data, headers=headers)


TOKEN = res.json()['access_token']

headers['Authorization'] = f'bearer {TOKEN}'

# GET - MY USER PROFILE

def getProfile():
    me = requests.get('https://oauth.reddit.com/api/v1/me', headers=headers)

    with open('json/me/me.json', 'w') as outfile:
        json.dump(me.json(), outfile)


# GET - SUBREDDITS

def getSubreddits():

    subredditsData = requests.get(f'{BASE_URI}/subreddits/popular', headers= headers, params={'limit': '50'})

    with open('json/subreddits/subreddits.json', 'w') as outfile:
        json.dump(subredditsData.json(), outfile)

# GET - POSTS PER SUBREDDIT

def getPosts():
    subreddits_titles = []

    with open('json/cleanup/clean_subreddits.json', 'r') as infile:
        file = json.load(infile)
        

    for subreddit in file:
        title = subreddit['name']
        subreddits_titles.append(title)

    for subreddit in subreddits_titles:
        res = requests.get(f'{BASE_URI}/r/{subreddit}/hot',
                        headers=headers, params={'limit': '50'})

        with open(f'json/posts/{subreddit}.json', 'w') as outfile:
            json.dump(res.json(), outfile)


#EXECUTE REGION

getPosts()