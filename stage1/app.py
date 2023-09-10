from flask import Flask, request, jsonify
import datetime
import os

app = Flask(__name__)
app.json.sort_keys=False

@app.route('/api', methods=['GET'])
def get_info():
    slack_name = request.args.get('Ezemba Marvellous', default='Ezemba Marvellous')
    track = request.args.get('Backend', default='Backend')

    current_day = datetime.datetime.utcnow().strftime('%A')
    current_utc_time = datetime.datetime.utcnow().strftime('%Y-%m-%-d %H:%M:%S')

    github_url_file = "https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP/blob/master/stage1/app.py"
    github_url_source = "https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP"

    github_url_file = 'https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP/blob/master/stage1/app.py'
    github_url_source = 'https://github.com/marviigrey/HNG-BACKEND-INTERNSHIP'

    status_code = 200

    response = {
            'slack_name': slack_name,
            'current_day': current_day,
            'current_utc_time': current_utc_time,
            'track': track,
            'github_url_file': github_url_file,
            'github_url_source': github_url_source,
            'status_code': status_code

            }
    return jsonify(response)
if __name__=='__main__':
    app.run(host='0.0.0.0')
