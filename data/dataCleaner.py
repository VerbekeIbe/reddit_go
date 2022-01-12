import uuid, json, requests, random

# HELPER FUNCTIONS

def getRandomUser():
    userIndex = random.randint(0,99)

    with open('json/cleanup/clean_users.json', 'r') as infile:
        file = json.load(infile)

        random_user = file[userIndex]['id']
        return random_user

def getCommunityId(title):
    with open('json/cleanup/clean_subreddits.json', 'r') as infile:
        file = json.load(infile)

        community = next(item for item in file if item["name"] == title)

    return community['id']

def posts_amount():
    with open('json/cleanup/clean_posts.json', 'r') as infile:
        file = json.load(infile)
        amount = len(file)
        print(amount)

    # 2563 posts




# CLEANUP SUBREDDIT DATA
def cleanup_subreddits(): 
    with open('json/subreddits/subreddits.json', 'r') as infile:
        file = json.load(infile)
        subreddits = file['data']['children']

        data = []
    
        for subreddit in subreddits:
            subdata = {
                'id': str(uuid.uuid4()),
                'name' : subreddit['data']['display_name'],
                'description': subreddit['data']['description'],
                'description_html': subreddit['data']['description_html']
            }
            data.append(subdata)

    with open('json/cleanup/clean_subreddits.json', 'w') as outfile:
        json.dump(data, outfile)




# CLEANUP POST DATA
def cleanup_posts():
    data = []

    # Determine subreddit titles
    subreddits_titles = []

    with open('json/cleanup/clean_subreddits.json', 'r') as infile:
        file = json.load(infile)
        

    for subreddit in file:
        title = subreddit['name']
        subreddits_titles.append(title)
    

    # Get data from json files

    for title in subreddits_titles:
        with open(f'json/posts/{title}.json', 'r') as infile:
            file = json.load(infile)
            posts = file['data']['children']

            for post in posts:
                # assign community
                community_id = getCommunityId(title)

                # assign a random user id

                user_id = getRandomUser()


                postData = {
                    'id':  str(uuid.uuid4()),
                    'community_id': community_id,
                    'user_id': user_id,
                    'title': post['data']['title'],
                    'content': post['data']['selftext'],
                    'content_html': post['data']['selftext_html'],
                    'timestamp': post['data']['created_utc']
                }
                data.append(postData)




    with open('json/cleanup/clean_posts.json', 'w') as outfile:
        json.dump(data, outfile)






#EXECUTE REGION

# cleanup_subreddits()

# cleanup_posts()

# posts_amount()


    

