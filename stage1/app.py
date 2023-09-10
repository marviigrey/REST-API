import json
import datetime

def lambda_handler(event, context):
    # Extract query parameters from the event
    slack_name = event.get('queryStringParameters', {}).get('slack_name', 'Ezemba Marvellous')
    
    # Get the current day of the week
    current_day = datetime.datetime.now().strftime('%A')
    
    # Get the current UTC time
    utc_time = datetime.datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')
    
    # Define other required information
    track = "backend"
    github_file_url = "https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP/blob/master/stage1/app.py"
    github_repo_url = "https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP"
    
    # Prepare the JSON response
    response = {
        "slack_name": slack_name,
        "current_day": current_day,
        "utc_time": utc_time,
        "track": track,
        "github_file_url": https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP/blob/master/stage1/app.py,
        "github_repo_url": https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP,
        "status_code": 200
    }
    
    return {
        "statusCode": 200,
        "body": json.dumps(response)
    }
