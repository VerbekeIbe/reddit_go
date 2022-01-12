import uuid, json, requests

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





#EXECUTE REGION

cleanup_subreddits()
    

